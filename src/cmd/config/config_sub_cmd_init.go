package cmd_config

import (
	"github.com/cubbit/composer-cli/src/configuration"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func NewConfigSubCmdInit() *cobra.Command {
	var configInitCmd = &cobra.Command{
		Use:   "init",
		Short: "Initialize configuration file",
		Run: func(cmd *cobra.Command, args []string) {
			config := configuration.NewConfig()
			if err := config.SaveConfig(); err != nil {
				utils.PrintError(err)
				return
			}
		},
	}

	return configInitCmd
}
