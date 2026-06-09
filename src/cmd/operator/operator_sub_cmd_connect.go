package cmd_operator

import (
	"github.com/cubbit/composer-cli/src/service"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func NewOperatorSubCmdConnect(
	operatorService service.OperatorServiceInterface,
) *cobra.Command {
	var operatorConnectSubCmd = &cobra.Command{
		Use:     "generate-connect-command",
		Short:   "Print the command to install operator on cluster (creates cluster location)",
		Aliases: []string{""},
		PreRun: func(cmd *cobra.Command, args []string) {
		},
		Run: func(cmd *cobra.Command, args []string) {
			if err := operatorService.Connect(cmd, args); err != nil {
				utils.PrintErrorWithWriter(cmd.ErrOrStderr(), err)
			}
		},
	}

	operatorConnectSubCmd.Flags().StringP("profile", "P", "", "Profile to use for login (default: use active profile)")

	return operatorConnectSubCmd
}
