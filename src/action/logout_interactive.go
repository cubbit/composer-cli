package action

import (
	"fmt"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/src/tui"
	"github.com/cubbit/cubbit/client/cli/utils"
	"github.com/spf13/cobra"
)

func SignOutOperatorInteractive(cmd *cobra.Command) error {
	var err error
	var configPath, profile, defaultConfigPath string

	if defaultConfigPath, err = configuration.GetDefaultConfigPath(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if _, err = tui.TextInputs("", true, tui.Input{Placeholder: fmt.Sprintf("Enter the config file path to load (default: %s)", defaultConfigPath), Value: &configPath}, tui.Input{Placeholder: "Enter the configuration profile (default: default)", Value: &profile}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if configPath == "" {
		configPath = defaultConfigPath
	}

	if profile == "" {
		profile = constants.DefaultProfile
	}

	var conf = configuration.NewConfig(configuration.SessionTypeOperator, profile, configuration.Url{}, "")
	if err = conf.StoreSession(configPath); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorStoringSession, err)
	}

	utils.PrintSuccess(fmt.Sprintf("operator %s signed out successfully", profile))

	return nil
}

func SignOutAccountInteractive(cmd *cobra.Command) error {
	var err error
	var configPath, profile, defaultConfigPath string

	if defaultConfigPath, err = configuration.GetDefaultConfigPath(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if _, err = tui.TextInputs(
		"",
		true,
		tui.Input{Placeholder: fmt.Sprintf("Enter the config file path to load (default: %s)", defaultConfigPath), Value: &configPath},
		tui.Input{Placeholder: "Enter the configuration profile (default: default)", Value: &profile}); err != nil {

		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if configPath == "" {
		configPath = defaultConfigPath
	}

	if profile == "" {
		profile = constants.DefaultProfile
	}

	var conf = configuration.NewConfig(configuration.SessionTypeAccount, profile, configuration.Url{}, "")
	if err = conf.StoreSession(configPath); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorStoringSession, err)
	}

	utils.PrintSuccess(fmt.Sprintf("account %s signed out successfully", profile))

	return nil
}
