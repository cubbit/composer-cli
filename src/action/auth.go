// Package action provides CLI actions for managing auth.
package action

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/api"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/google/uuid"
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
)

const (
	pollInterval    = 2 * time.Second
	maxPollDuration = 10 * time.Minute
)

var boldStyle = lipgloss.NewStyle().Bold(true)

func SignInComposer(cmd *cobra.Command, args []string) error {
	var err error
	var profile string
	var conf *configuration.Config
	var urls *configuration.URLs
	var endpoint string

	if profile, err = cmd.Flags().GetString("profile"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if endpoint, err = cmd.Flags().GetString("endpoint"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	conf, err = configuration.LoadConfig()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	resolvedEndpoint := endpoint
	if resolvedEndpoint == "" {
		if existingProfile, err := conf.ResolveProfile(profile); err == nil {
			resolvedEndpoint = existingProfile.Endpoint
		} else {
			resolvedEndpoint = conf.Default.Endpoint
		}
	}

	urls, err = configuration.ConfigureAPIServerURL(resolvedEndpoint)
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorConfiguringAPIURL, err)
	}

	return performBrowserLogin(*urls, conf, profile, resolvedEndpoint)
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

func Logout(cmd *cobra.Command, args []string) error {
	var err error
	var profile string
	var allProfiles bool

	if profile, err = cmd.Flags().GetString("profile"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if allProfiles, err = cmd.Flags().GetBool("all"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	config, err := configuration.LoadConfig()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if allProfiles {
		config.Profile = make(map[string]*configuration.Profile)
	} else {
		if profile == "" {
		} else {
			if err = config.DeleteProfile(profile); err != nil {
				return fmt.Errorf("failed to logout from profile '%s': %w", profile, err)
			}
		}
	}

	if err = config.SaveConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorSavingConfig, err)
	}

	return nil
}
