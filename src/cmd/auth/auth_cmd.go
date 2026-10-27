package cmd_auth

import (
	"github.com/cubbit/composer-cli/src/service"
	"github.com/spf13/cobra"
)

func NewAuthCmd(
	authServiceFn service.AuthServiceInterface,
) *cobra.Command {
	var authCmd = &cobra.Command{
		Use:   "auth",
		Short: "Execute commands in auth sections",
	}

	authActivateSubCmd := NewAuthSubCmdActivate(authServiceFn)
	authCmd.AddCommand(authActivateSubCmd)

	authLoginSubCmd := NewAuthSubCmdLogin(authServiceFn)
	authCmd.AddCommand(authLoginSubCmd)

	authLogoutSubCmd := NewAuthSubCmdLogout(authServiceFn)
	authCmd.AddCommand(authLogoutSubCmd)

	authSignUpSubCmd := NewAuthSubCmdSignUp(authServiceFn)
	authCmd.AddCommand(authSignUpSubCmd)

	return authCmd
}
