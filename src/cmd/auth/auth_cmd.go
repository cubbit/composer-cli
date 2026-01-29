package cmd_auth

import (
	"github.com/cubbit/composer-cli/src/service"
	"github.com/spf13/cobra"
)

func NewAuthCmd(
	authService service.AuthServiceInterface,
) *cobra.Command {
	var authCmd = &cobra.Command{
		Use:   "auth",
		Short: "Execute commands in auth sections",
	}

	authLoginSubCmd := NewAuthSubCmdLogin(authService)
	authCmd.AddCommand(authLoginSubCmd)

	authLogoutSubCmd := NewAuthSubCmdLogout(authService)
	authCmd.AddCommand(authLogoutSubCmd)
	return authCmd
}
