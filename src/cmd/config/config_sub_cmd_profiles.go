package cmd_config

import (
	"github.com/cubbit/composer-cli/src/service"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func NewConfigSubCmdProfiles(
	configService service.ConfigServiceInterface,
) *cobra.Command {
	var configProfilesCmd = &cobra.Command{
		Use:   "profiles",
		Short: "List all configuration profiles",
		Long:  `List all available configuration profiles with their details.`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := configService.Profiles(cmd, args); err != nil {
				utils.PrintError(err)
			}
		},
	}

	return configProfilesCmd
}
