package cmd_config

import (
	"github.com/cubbit/composer-cli/src/service"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func NewConfigSubCmdEdit(
	configService service.ConfigServiceInterface,
) *cobra.Command {
	var configEditCmd = &cobra.Command{
		Use:   "edit",
		Short: "Edit configuration file",
		Long:  `Open the configuration file in your default editor ($EDITOR).`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := configService.Edit(cmd, args); err != nil {
				utils.PrintError(err)
			}
		},
	}

	return configEditCmd
}
