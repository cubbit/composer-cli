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
	var configPath, profile string

	if _, err = tui.TextInputs("", true, tui.Input{Placeholder: "Enter the config file to load (default: ./)", Value: &configPath}, tui.Input{Placeholder: "Enter the configuration profile (default: default)", Value: &profile}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if configPath == "" {
		configPath = constants.DefaultFilePath
	}

	if profile == "" {
		profile = constants.DefaultProfile
	}

	var conf = configuration.NewConfig(profile, configuration.Url{}, "")
	if err = conf.StoreSession(configPath); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorStoringSession, err)
	}

	utils.PrintSuccess(fmt.Sprintf("configuration %s signed out successfully", profile))

	return nil
}
