package action

import (
	"fmt"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/utils"
	"github.com/spf13/cobra"
)

func SignOutOperator(cmd *cobra.Command) error {
	var err error
	var profile, configPath string

	if profile, err = cmd.Flags().GetString("profile"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if configPath, err = cmd.Flags().GetString("config"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	var conf = configuration.NewConfig(profile, configuration.Url{}, "")

	if err = conf.StoreSession(configPath); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorStoringSession, err)
	}

	utils.PrintSuccess(fmt.Sprintf("configuration %s signed out successfully", profile))
	
	return nil
}
