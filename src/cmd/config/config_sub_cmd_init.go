package cmd_config

import (
	"github.com/cubbit/composer-cli/src/service"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func NewConfigSubCmdInit(
	configService service.ConfigServiceInterface,
) *cobra.Command {
	var configInitCmd = &cobra.Command{
		Use:   "init",
		Short: "Initialize configuration file",
		Run: func(cmd *cobra.Command, args []string) {

			if err := configService.InitConfiguration(cmd, args); err != nil {
				utils.PrintError(err)
				return
			}
		},
	}

	return configInitCmd
}
