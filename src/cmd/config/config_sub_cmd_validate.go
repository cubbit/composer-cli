package cmd_config

import (
	"github.com/cubbit/composer-cli/src/service"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func NewConfigSubCmdValidate(
	configService service.ConfigServiceInterface,
) *cobra.Command {
	var configValidateCmd = &cobra.Command{
		Use:   "validate",
		Short: "Validate configuration",
		Long:  `Validate the configuration file and all profiles for correctness.`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := configService.Validate(cmd, args); err != nil {
				utils.PrintError(err)
			}
		},
	}

	return configValidateCmd
}
