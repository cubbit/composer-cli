package cmd_auth

import (
	"github.com/cubbit/composer-cli/src/service"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func NewAuthSubCmdActivate(
	authService service.AuthServiceInterface,
) *cobra.Command {
	var authActivateCmd = &cobra.Command{
		Use:   "activate",
		Short: "Activate an operator account",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			cmd.MarkFlagRequired("token")

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			if err := authService.Activate(cmd, args); err != nil {
				utils.PrintError(err)
			}
		},
	}

	authActivateCmd.Flags().String("token", "", "Activation token sent via email (required)")

	return authActivateCmd
}
