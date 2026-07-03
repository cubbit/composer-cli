package service

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/BurntSushi/toml"
	"github.com/cubbit/composer-cli/constants"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	utils "github.com/cubbit/composer-cli/utils"
	"github.com/cubbit/composer-cli/utils/printer"
	"github.com/spf13/cobra"
)

type ConfigServiceInterface interface {
	InitConfiguration(cmd *cobra.Command, args []string) error
	View(cmd *cobra.Command, args []string) error
	Edit(cmd *cobra.Command, args []string) error
	Profiles(cmd *cobra.Command, args []string) error
	SwitchProfile(cmd *cobra.Command, args []string) error
	Validate(cmd *cobra.Command, args []string) error
}

type ConfigService struct {
	configuration *configuration_handler.ConfigurationHandler
}

func NewConfigService(
	configuration *configuration_handler.ConfigurationHandler,
) *ConfigService {
	return &ConfigService{
		configuration: configuration,
	}
}

func (s *ConfigService) View(cmd *cobra.Command, args []string) error {
	config, err := s.configuration.GetConfigFile()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	encoder := toml.NewEncoder(cmd.OutOrStdout())
	if err = encoder.Encode(config); err != nil {
		return fmt.Errorf("failed to encode config: %w", err)
	}

	return nil
}

func (s *ConfigService) InitConfiguration(cmd *cobra.Command, args []string) error {
	configPath, err := s.configuration.GetConfigFilePath()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		configFile := configuration_models.CreateEmptyConfigV2()
		err := configFile.Save(configPath)
		if err != nil {
			return fmt.Errorf("failed to create config file at %s: %w", configPath, err)
		}

		return printer.PrintText(
			cmd,
			fmt.Sprintf("Created configuration file template in %s\n", configPath),
		)
	} else if err != nil {
		return fmt.Errorf("failed to check if config file exists: %w", err)
	}

	return printer.PrintText(
		cmd,
		fmt.Sprintf("Configuration file already exists at %s\n", configPath),
	)
}

func (s *ConfigService) Edit(cmd *cobra.Command, args []string) error {
	configPath, err := s.configuration.GetConfigFilePath()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		configFile := configuration_models.CreateEmptyConfigV2()
		err := configFile.Save(configPath)
		if err != nil {
			return fmt.Errorf("failed to create config file at %s: %w", configPath, err)
		}
	}

	editor := ""
	editorFromFlag, err := cmd.Flags().GetString("editor")
	if err != nil {
		return fmt.Errorf("failed to get 'editor' flag: %w", err)
	}

	if editorFromFlag != "" {
		editor = editorFromFlag
	}

	if editor == "" {
		editor = os.Getenv("EDITOR")
	}

	if editor == "" {
		editor = "nano"
	}

	editorCmd := exec.Command(editor, configPath)
	editorCmd.Stdin = os.Stdin
	editorCmd.Stdout = os.Stdout
	editorCmd.Stderr = os.Stderr

	if err := editorCmd.Run(); err != nil {
		return fmt.Errorf("failed to open editor: %w", err)
	}

	return nil
}

func (s *ConfigService) Profiles(cmd *cobra.Command, args []string) error {
	profiles, err := s.configuration.GetProfiles()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	activeProfileName, err := s.configuration.GetActiveProfileName()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	return PrintProfiles(cmd, profiles, activeProfileName)
}

func (s *ConfigService) SwitchProfile(cmd *cobra.Command, args []string) error {
	profileName := args[0]

	err := s.configuration.SwitchProfile(profileName)
	if err != nil {
		return fmt.Errorf("failed to switch profile: %w", err)
	}

	return nil
}

func (s *ConfigService) Validate(cmd *cobra.Command, args []string) error {
	configPath, err := s.configuration.GetConfigFilePath()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	_, err = configuration_models.ParseAndValidateConfigV2(configPath)
	if err != nil {
		utils.PrintErrorWithWriter(cmd.ErrOrStderr(), err)
		return nil
	}

	return printer.PrintText(
		cmd,
		fmt.Sprintf("Configuration file at %s is valid", configPath),
	)
}
