package cmd_config

import (
	"github.com/cubbit/composer-cli/src/service"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func NewConfigSubCmdSwitchProfile(
	configService service.ConfigServiceInterface,
) *cobra.Command {
	var configSwitchProfileCmd = &cobra.Command{
		Use:   "switch-profile <profile-name>",
		Short: "Switch to a different profile",
		Long:  `Switch the active profile. This will set the default profile for subsequent commands.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if err := configService.SwitchProfile(cmd, args); err != nil {
				utils.PrintError(err)
			}
		},
	}

	return configSwitchProfileCmd
}
