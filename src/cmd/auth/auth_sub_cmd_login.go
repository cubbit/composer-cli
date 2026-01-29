package cmd_auth

import (
	"github.com/cubbit/composer-cli/src/service"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func NewAuthSubCmdLogin(
	authService service.AuthServiceInterface,
) *cobra.Command {
	var authLoginSubCmd = &cobra.Command{
		Use:     "login",
		Short:   "Login the user",
		Aliases: []string{"signin"},
		PreRun: func(cmd *cobra.Command, args []string) {
			cmd.MarkFlagRequired("profile")
		},
		Run: func(cmd *cobra.Command, args []string) {
			if err := authService.Login(cmd, args); err != nil {
				utils.PrintError(err)
			}
		},
	}

	authLoginSubCmd.Flags().String("endpoint", "", "Endpoint to connect to (default: use configured endpoint)")
	authLoginSubCmd.Flags().StringP("profile", "P", "", "Profile to use for login (default: use active profile)")

	return authLoginSubCmd
}
