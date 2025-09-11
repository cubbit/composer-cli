// Package cmd provides CLI commands for managing authentication.
package cmd

import (
	"github.com/cubbit/cubbit/client/cli/src/action"
	"github.com/cubbit/cubbit/client/cli/utils"
	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Execute commands in auth sections",
}

var authLoginSubCmd = &cobra.Command{
	Use:     "login",
	Short:   "Login the user",
	Aliases: []string{"signin"},
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("profile")
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.SignInComposer(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

var authLogoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Log out the user",
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.Logout(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

func init() {

	authCmd.AddCommand(authLoginSubCmd)
	authLoginSubCmd.Flags().String("endpoint", "", "Endpoint to connect to (default: use configured endpoint)")
	authLoginSubCmd.Flags().StringP("profile", "P", "", "Profile to use for login (default: use active profile)")

	authCmd.AddCommand(authLogoutCmd)
	authLogoutCmd.Flags().String("profile", "", "Profile to use for logout (default: use active profile)")
	authLogoutCmd.Flags().Bool("all", false, "Logout from all profiles")

	rootCmd.AddCommand(authCmd)
}
