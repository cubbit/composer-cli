package cmd_config

import (
	"github.com/cubbit/composer-cli/src/service"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func NewConfigSubCmdView(
	configService service.ConfigServiceInterface,
) *cobra.Command {
	var configViewCmd = &cobra.Command{
		Use:   "view",
		Short: "Display current configuration",
		Long:  `Display the current configuration. Can show entire config or specific profile.`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := configService.View(cmd, args); err != nil {
				utils.PrintError(err)
			}
		},
	}

	return configViewCmd
}
