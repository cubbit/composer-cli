package service

import (
	"context"
	"crypto/sha256"
	b64 "encoding/base64"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/cubbit/composer-cli/constants"
	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration"
	"github.com/cubbit/composer-cli/src/tui"
	"github.com/cubbit/composer-cli/utils"
	"github.com/google/uuid"
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
)

const (
	pollInterval    = 2 * time.Second
	maxPollDuration = 10 * time.Minute
)

var boldStyle = lipgloss.NewStyle().Bold(true)

type AuthServiceInterface interface {
	Activate(cmd *cobra.Command, args []string) error
	SignUp(cmd *cobra.Command, args []string) error
	Login(cmd *cobra.Command, args []string) error
	Logout(cmd *cobra.Command, args []string) error
}

type AuthService struct {
	configuration *configuration.Config
	authAPI       api.AuthAPIInterface
}

func NewAuthService(
	configuration *configuration.Config,
	authAPI api.AuthAPIInterface,
) *AuthService {
	return &AuthService{
		configuration: configuration,
		authAPI:       authAPI,
	}
}

func (as *AuthService) Activate(cmd *cobra.Command, args []string) error {
	token, err := cmd.Flags().GetString("token")
	if err != nil {
		return fmt.Errorf("%s token: %w", constants.ErrorRetrievingField, err)
	}

	resolvedProfile, urls, err := as.configuration.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer)
	if err != nil {
		return fmt.Errorf("failed to resolve provide and urls: %w", err)
	}

	err = as.authAPI.Activate(*urls, token)
	if err != nil {
		return fmt.Errorf("failed during activation request: %w", err)
	}

	return utils.PrintSmartOutput(
		cmd,
		[]string{"Activation completed successfully. You can now log in."},
		func(s string) []string { return []string{s} },
		&utils.SmartOutputConfig[string]{
			SingleResource:              true,
			SingleResourceCompactOutput: true,
			DefaultOutput:               resolvedProfile.Output,
		},
	)
}

func (as *AuthService) SignUp(cmd *cobra.Command, args []string) error {
	email, err := cmd.Flags().GetString("email")
	if err != nil {
		return fmt.Errorf("%s email: %w", constants.ErrorRetrievingField, err)
	}

	username, err := cmd.Flags().GetString("username")
	if err != nil {
		return fmt.Errorf("%s username: %w", constants.ErrorRetrievingField, err)
	}

	firstName, err := utils.GetOptionalStringFlag(cmd, "first-name")
	if err != nil {
		return fmt.Errorf("%s first-name: %w", constants.ErrorRetrievingField, err)
	}

	lastName, err := utils.GetOptionalStringFlag(cmd, "last-name")
	if err != nil {
		return fmt.Errorf("%s last-name: %w", constants.ErrorRetrievingField, err)
	}

	password, err := utils.GetOptionalStringFlag(cmd, "password")
	if err != nil {
		return fmt.Errorf("%s password: %w", constants.ErrorRetrievingField, err)
	}

	orgName, err := cmd.Flags().GetString("organization")
	if err != nil {
		return fmt.Errorf("%s organization name: %w", constants.ErrorRetrievingField, err)
	}

	basePolicy, err := utils.JSONMapFromCommand(cmd, "base-policy")
	if err != nil {
		return fmt.Errorf("%s base-policy: %w", constants.ErrorRetrievingField, err)
	}

	settings, err := utils.JSONMapFromCommand(cmd, "settings")
	if err != nil {
		return fmt.Errorf("%s settings: %w", constants.ErrorRetrievingField, err)
	}

	resolvedProfile, urls, err := as.configuration.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer)
	if err != nil {
		return fmt.Errorf("failed to resolve provide and urls: %w", err)
	}

	var authenticationPublicKey *string
	if password != nil {
		challenge, err := as.authAPI.GenerateChallenge(
			*urls,
			nil,
			&username,
			&orgName,
		)

		if err != nil {
			return fmt.Errorf("failed to generate challenge: %w", err)
		}

		h := sha256.New()
		h.Write([]byte(*password + challenge.Salt))
		seed := h.Sum(nil)

		publicKey, _, err := utils.GenerateKeyPairFromSeed(seed)
		if err != nil {
			return fmt.Errorf("failed to generate key pair from seed: %w", err)
		}

		base64PublicKey := b64.StdEncoding.EncodeToString(publicKey)
		authenticationPublicKey = &base64PublicKey
	}

	if err := as.authAPI.SignUp(
		*urls,
		email,
		username,
		firstName,
		lastName,
		authenticationPublicKey,
		orgName,
		basePolicy,
		settings,
	); err != nil {
		return fmt.Errorf("failed during sign up request: %w", err)
	}

	return utils.PrintSmartOutput(
		cmd,
		[]string{"Sign up completed successfully. Please check your email to verify your account."},
		func(s string) []string { return []string{s} },
		&utils.SmartOutputConfig[string]{
			SingleResource:              true,
			SingleResourceCompactOutput: true,
			DefaultOutput:               resolvedProfile.Output,
		},
	)
}

func (as *AuthService) Login(cmd *cobra.Command, args []string) error {
	var err error
	var profile, username, orgName string
	var urls *configuration.URLs
	var endpoint string

	if profile, err = cmd.Flags().GetString("profile"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if endpoint, err = cmd.Flags().GetString("endpoint"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	resolvedEndpoint := endpoint
	if resolvedEndpoint == "" {
		if existingProfile, err := as.configuration.ResolveProfile(profile); err == nil {
			resolvedEndpoint = existingProfile.Endpoint
		} else {
			resolvedEndpoint = as.configuration.Default.Endpoint
		}
	}

	urls, err = configuration.ConfigureAPIServerURL(resolvedEndpoint)
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorConfiguringAPIURL, err)
	}

	if username, err = cmd.Flags().GetString("username"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if orgName, err = cmd.Flags().GetString("organization"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	apiKey, err := cmd.Flags().GetString("api-key")
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if apiKey == "" {
		value, exists := os.LookupEnv("API_KEY")
		if exists {
			apiKey = value
		}
	}

	if apiKey != "" {
		return as.configureAPIKey(cmd, *urls, as.configuration, profile, apiKey, resolvedEndpoint)
	}

	password, err := cmd.Flags().GetString("password")
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if password == "" {
		value, exists := os.LookupEnv("PASSWORD")
		if exists {
			password = value
		}
	}

	if password != "" && (username == "" || orgName == "") {
		return fmt.Errorf("when using the --password flag, both --username and --organization must also be provided; alternatively, omit --password to be prompted interactively")
	}

	if password != "" {
		return as.performInlineLogin(cmd, *urls, as.configuration, username, orgName, password, "", profile, resolvedEndpoint)
	}

	var choice string
	choices := []string{
		"Yes, open browser for authentication",
		"No, I want to login inline with username and password",
		"No, I want to provide an API key",
	}
	if choice, err = tui.ChooseOne("Can you currently access the composer dashboard", false, false, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorSignIn, err)
	}

	switch choice {
	case choices[0]:
		return performBrowserLogin(*urls, as.configuration, profile, resolvedEndpoint)
	case choices[1]:
		var password string
		var tfa string
		if _, err = tui.TextInputs(
			"Please provide your login details",
			false,
			tui.Input{Placeholder: "username", IsPassword: false, Value: &username},
			tui.Input{Placeholder: "organization", IsPassword: false, Value: &orgName},
			tui.Input{Placeholder: "password", IsPassword: true, Value: &password},
			tui.Input{Placeholder: "two-factor code (if enabled)", IsPassword: false, Value: &tfa},
		); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorSignIn, err)
		}
		return as.performInlineLogin(cmd, *urls, as.configuration, username, orgName, password, tfa, profile, resolvedEndpoint)
	case choices[2]:
		var apiKey string
		if _, err = tui.TextInputs(
			"Please provide your API key for authentication",
			false,
			tui.Input{Placeholder: "API key", IsPassword: true, Value: &apiKey},
		); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorSignIn, err)
		}

		return as.configureAPIKey(cmd, *urls, as.configuration, profile, apiKey, resolvedEndpoint)
	default:
		return fmt.Errorf("invalid choice: %s", choice)
	}
}

func (as *AuthService) configureAPIKey(cmd *cobra.Command, urls configuration.URLs, conf *configuration.Config, profile string, apiKey string, resolvedEndpoint string) error {
	if err := conf.CreateProfile(profile, configuration.ProfileTypeComposer, resolvedEndpoint, apiKey); err != nil {
		return fmt.Errorf("failed to create profile: %w", err)
	}

	if err := conf.SaveConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorSavingConfig, err)
	}

	fmt.Fprintf(cmd.OutOrStdout(), "✅ Authentication successful!\n")

	return nil
}

func (as *AuthService) performInlineLogin(cmd *cobra.Command, urls configuration.URLs, conf *configuration.Config, username, orgName, password, tfa, profile string, resolvedEndpoint string) error {
	var err error
	var tokens *api.SignInToken
	var apiKey string
	var hostname string
	fmt.Fprintf(cmd.OutOrStdout(), "🔐 Starting inline authentication...\n")

	if tokens, err = as.authAPI.SignIn(urls, username, orgName, password, tfa); err != nil {
		return fmt.Errorf("failed to perform sign in: %w", err)
	}

	rawHostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown-host"
	} else {
		hostname = strings.ReplaceAll(rawHostname, " ", "-")
	}
	date := time.Now().Format("20060102")
	nozzle := uuid.New().String()[:8]
	apiKeyName := fmt.Sprintf("composer-cli-%s-%s-%s-%s-%s", runtime.GOOS, runtime.GOARCH, hostname, date, nozzle)

	operator, err := api.GetOperatorSelf(urls, tokens.AccessToken, "")
	if err != nil {
		return fmt.Errorf("failed to retrieve operator information: %w", err)
	}

	if tfa != "" {
		if _, err = tui.TextInputs(
			"Please provide another two-factor code to forge an API key",
			false,
			tui.Input{Placeholder: "two-factor code", IsPassword: false, Value: &tfa},
		); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorSignIn, err)
		}
	}

	if len(operator.Emails) == 0 {
		return fmt.Errorf("operator email is missing, cannot forge API key")
	}

	operatorApiKeyToken, err := as.authAPI.ForgeToken(urls, operator.ID, operator.Emails[0].Email, password, tfa, "create_operator_api_key", tokens.AccessToken, tokens.RefreshToken)
	if err != nil {
		return fmt.Errorf("failed to forge token: %w", err)
	}

	if apiKey, err = as.authAPI.CreateApiKey(urls, operator.ID, apiKeyName, tokens.AccessToken, operatorApiKeyToken); err != nil {
		return fmt.Errorf("failed to create API key: %w", err)
	}

	if err = conf.CreateProfile(profile, configuration.ProfileTypeComposer, resolvedEndpoint, apiKey); err != nil {
		return fmt.Errorf("failed to create profile: %w", err)
	}

	if err = conf.SaveConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorSavingConfig, err)
	}

	fmt.Fprintf(cmd.OutOrStdout(), "✅ Authentication successful!\n")
	fmt.Fprintf(cmd.OutOrStdout(), "📜 Created API Key: %s\n", apiKeyName)

	return nil
}

func performBrowserLogin(urls configuration.URLs, conf *configuration.Config, profile string, resolvedEndpoint string) error {
	var err error
	var device *api.DeviceRegistrationResponse
	var apiKey string

	fmt.Println("🔐 Starting browser-based authentication...")

	uuid := uuid.New()

	// Register device and get device code
	if device, err = api.RegisterDevice(urls, uuid.String()); err != nil {
		return fmt.Errorf("failed to register device: %w", err)
	}

	fmt.Printf("🔑 First copy your one-time code: %s\n\n", boldStyle.Render(device.DeviceCode))

	// Build auth URL and prompt user to press Enter
	authURL := buildAuthURL(device.ClientURL)
	fmt.Printf("📋 Auth URL: %s\n\n", authURL)
	fmt.Println("Press Enter to open browser...")

	inputChan := make(chan bool, 1)
	apiKeyChan := make(chan string, 1)
	errChan := make(chan error, 1)

	go func() {
		fmt.Scanln()
		inputChan <- true
	}()

	go func() {
		key, err := pollForAPIKey(urls, uuid.String())
		if err != nil {
			errChan <- err
			return
		}
		apiKeyChan <- key
	}()

	select {
	case <-inputChan:
		fmt.Printf("🌐 Opening browser for authentication...\n")
		if err = browser.OpenURL(authURL); err != nil {
			fmt.Printf("⚠️ Failed to open browser automatically: %v\n", err)
			fmt.Printf("Please manually visit the URL above\n\n")
		} else {
			fmt.Println("👀 Waiting for authentication...")
		}

		fmt.Println()
		fmt.Println("📝 Complete the authentication in your browser")
		fmt.Println("⏳ This window will automatically continue once you're done")
		fmt.Println()

		select {
		case apiKey = <-apiKeyChan:
			fmt.Println("✅ Authentication successful!")
			fmt.Printf("📜 API Key received: %s\n", apiKey)
		case err = <-errChan:
			return fmt.Errorf("authentication failed: %w", err)
		}

	case apiKey = <-apiKeyChan:
		fmt.Println("✅ Authentication successful!")
		fmt.Printf("📜 API Key received: %s\n", apiKey)
	case err = <-errChan:
		return fmt.Errorf("authentication failed: %w", err)
	}

	if err = conf.CreateProfile(profile, configuration.ProfileTypeComposer, resolvedEndpoint, apiKey); err != nil {
		return fmt.Errorf("failed to create profile: %w", err)
	}

	if err = conf.SaveConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorSavingConfig, err)
	}

	return nil
}

func buildAuthURL(clientURL string) string {
	return fmt.Sprintf("%s/dashboard/auth/devices", clientURL)
}

func pollForAPIKey(urls configuration.URLs, deviceID string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), maxPollDuration)
	defer cancel()

	ticker := time.NewTicker(pollInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return "", fmt.Errorf("authentication timed out after %v", maxPollDuration)
		case <-ticker.C:
			apiKey, err := api.GetDeviceAPIKey(urls, deviceID)
			if err != nil {
				if strings.Contains(err.Error(), "not found") {
					return "", fmt.Errorf("device not found - registration may have expired")
				}
				return "", fmt.Errorf("failed to check authorization status: %w", err)
			}

			if apiKey != "" {
				return apiKey, nil
			}
			// continue polling if API key is not yet available
		}
	}
}

func (as *AuthService) Logout(cmd *cobra.Command, args []string) error {
	var err error
	var profile string
	var allProfiles bool

	if profile, err = cmd.Flags().GetString("profile"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if allProfiles, err = cmd.Flags().GetBool("all"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if allProfiles {
		as.configuration.Profile = make(map[string]*configuration.Profile)
	} else {
		if profile == "" {
		} else {
			if err = as.configuration.DeleteProfile(profile); err != nil {
				return fmt.Errorf("failed to logout from profile '%s': %w", profile, err)
			}
		}
	}

	if err = as.configuration.SaveConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorSavingConfig, err)
	}

	return nil
}
