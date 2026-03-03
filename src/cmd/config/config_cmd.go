package cmd_config

import (
	"github.com/cubbit/composer-cli/src/service"
	"github.com/spf13/cobra"
)

func NewConfigCmd(
	configService service.ConfigServiceInterface,
) *cobra.Command {
	var configCmd = &cobra.Command{
		Use:   "config",
		Short: "Manage Cubbit CLI configuration",
		Long:  `Manage Cubbit CLI configuration including profiles, endpoints, and authentication settings.`,
	}

	configInitSubCmd := NewConfigSubCmdInit()
	configCmd.AddCommand(configInitSubCmd)

	configViewSubCmd := NewConfigSubCmdView(configService)
	configCmd.AddCommand(configViewSubCmd)

	configEditSubCmd := NewConfigSubCmdEdit(configService)
	configCmd.AddCommand(configEditSubCmd)

	configProfilesSubCmd := NewConfigSubCmdProfiles(configService)
	configCmd.AddCommand(configProfilesSubCmd)

	configSwitchProfileSubCmd := NewConfigSubCmdSwitchProfile(configService)
	configCmd.AddCommand(configSwitchProfileSubCmd)

	configValidateSubCmd := NewConfigSubCmdValidate(configService)
	configCmd.AddCommand(configValidateSubCmd)

	return configCmd
}
