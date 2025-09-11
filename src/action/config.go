package action

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/BurntSushi/toml"
	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/utils"
	"github.com/spf13/cobra"
)

type ProfileInfo struct {
	Name      string                    `json:"name"`
	Type      configuration.ProfileType `json:"type"`
	Endpoint  string                    `json:"endpoint"`
	HasAPIKey bool                      `json:"has_token"`
	UpdatedAt string                    `json:"updated_at"`
	IsDefault bool                      `json:"is_default"`
}

func ConfigView(cmd *cobra.Command, args []string) error {
	config, err := configuration.LoadConfig()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	encoder := toml.NewEncoder(os.Stdout)
	if err = encoder.Encode(config); err != nil {
		return fmt.Errorf("failed to encode config: %w", err)
	}
	if err = encoder.Encode(config.Profile); err != nil {
		return fmt.Errorf("failed to encode profiles: %w", err)
	}

	return nil
}

func ConfigEdit(cmd *cobra.Command, args []string) error {
	var err error
	var configPath string

	if configPath, err = configuration.GetDefaultConfigPath(); err != nil {
		return fmt.Errorf("failed to get default config path: %w", err)
	}

	configFile := configPath + "/config.toml"

	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		config := configuration.NewConfig()
		config.ConfigPath = configPath
		if err := config.SaveConfig(); err != nil {
			return fmt.Errorf("failed to create default config: %w", err)
		}
		utils.PrintInfo(fmt.Sprintf("Created default configuration at %s", configFile))
	}

	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "nano" // fallback to nano
	}

	editorCmd := exec.Command(editor, configFile)
	editorCmd.Stdin = os.Stdin
	editorCmd.Stdout = os.Stdout
	editorCmd.Stderr = os.Stderr

	if err := editorCmd.Run(); err != nil {
		return fmt.Errorf("failed to open editor: %w", err)
	}

	return nil
}

func ConfigProfiles(cmd *cobra.Command, args []string) error {
	config, err := configuration.LoadConfig()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	profiles := config.ListProfiles()
	profileInfos := listProfileInfos(config, profiles)

	return utils.PrintSmartOutput(
		cmd,
		profileInfos,
		func(info ProfileInfo) []string {
			return []string{
				info.Name,
				string(info.Type),
				info.Endpoint,
				fmt.Sprintf("%t", info.HasAPIKey),
			}
		},
		nil,
	)

}

func ConfigSwitchProfile(cmd *cobra.Command, args []string) error {
	profileName := args[0]

	config, err := configuration.LoadConfig()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if err := config.SetActiveProfile(profileName); err != nil {
		return fmt.Errorf("failed to switch to profile '%s': %w", profileName, err)
	}

	return nil
}

func ConfigValidate(cmd *cobra.Command, args []string) error {
	config, err := configuration.LoadConfig()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	return validateProfiles(config)
}

func listProfileInfos(config *configuration.Config, profiles []string) []ProfileInfo {
	var profileInfos []ProfileInfo

	for _, profileName := range profiles {
		resolved, err := config.ResolveProfile(profileName)
		if err != nil {
			continue
		}

		info := ProfileInfo{
			Name:      resolved.Name,
			Type:      resolved.Type,
			Endpoint:  resolved.Endpoint,
			HasAPIKey: resolved.APIKey != "",
		}

		if config.Active.Profile == resolved.Name {
			info.IsDefault = true
		}

		if !resolved.UpdatedAt.IsZero() {
			info.UpdatedAt = resolved.UpdatedAt.Format("2006-01-02T15:04:05Z")
		}

		profileInfos = append(profileInfos, info)
	}

	return profileInfos

}

func validateProfiles(config *configuration.Config) error {
	var hasErrors bool

	for profileName := range config.Profile {
		if err := config.ValidateProfile(profileName); err != nil {
			utils.PrintError(fmt.Errorf("profile '%s' validation failed: %w", profileName, err))

			hasErrors = true
		}
	}

	if hasErrors {
		return fmt.Errorf("configuration validation failed")
	}

	return nil
}
