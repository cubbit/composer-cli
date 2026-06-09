package cmd_auth

import (
	"github.com/cubbit/composer-cli/src/service"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func NewAuthSubCmdLogout(
	authService service.AuthServiceInterface,
) *cobra.Command {
	var authLogoutCmd = &cobra.Command{
		Use:   "logout",
		Short: "Log out the user",
		Run: func(cmd *cobra.Command, args []string) {
			if err := authService.Logout(cmd, args); err != nil {
				utils.PrintError(err)
			}
		},
	}

	authLogoutCmd.Flags().String("profile", "", "Profile to use for logout (default: use active profile)")
	authLogoutCmd.Flags().Bool("all", false, "Logout from all profiles")

	return authLogoutCmd
}
