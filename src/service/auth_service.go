package service

import (
	"context"
	"crypto/sha256"
	b64 "encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/cubbit/composer-cli/constants"
	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration"
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
			&email,
			nil,
			nil,
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
	var profile string
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

	return performBrowserLogin(*urls, as.configuration, profile, resolvedEndpoint)
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
