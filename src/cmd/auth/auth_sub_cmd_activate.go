package cmd_auth

import (
	"github.com/cubbit/composer-cli/src/service"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func NewAuthSubCmdActivate(
	authServiceFn service.AuthServiceInterface,
) *cobra.Command {
	var authActivateCmd = &cobra.Command{
		Use:   "activate",
		Short: "Activate an operator account",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			cmd.MarkFlagRequired("token")

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			if err := authServiceFn.Activate(cmd, args); err != nil {
				utils.PrintErrorWithWriter(cmd.ErrOrStderr(), err)
			}
		},
	}

	authActivateCmd.Flags().String("token", "", "Activation token sent via email (required)")
	authActivateCmd.Flags().String("endpoints", "", "Path to endpoints file (.toml)")
	authActivateCmd.RegisterFlagCompletionFunc("endpoints", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"toml"}, cobra.ShellCompDirectiveFilterFileExt
	})

	return authActivateCmd
}
