package cmd

import (
	"github.com/cubbit/cubbit/client/cli/src/action"
	"github.com/cubbit/cubbit/client/cli/utils"
	"github.com/spf13/cobra"
)

var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "Execute commands in account sections",
}

var accountSignupSubCmd = &cobra.Command{
	Use:   "signup",
	Short: "Create a new account",
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			cmd.MarkFlagRequired("email")
			cmd.MarkFlagRequired("password")
			cmd.MarkFlagRequired("first-name")
			cmd.MarkFlagRequired("last-name")
			cmd.MarkFlagRequired("tenant-id")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		action.SetupOutput(cmd)

		if !interactive {
			if err = action.CreateAccount(cmd, args); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.CreateAccountInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var accountLoginSubCmd = &cobra.Command{
	Use:     "login",
	Short:   "Login the account",
	Aliases: []string{"signin"},
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			cmd.MarkFlagRequired("email")
			cmd.MarkFlagRequired("password")
			cmd.MarkFlagRequired("tenant-id")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		action.SetupOutput(cmd)

		if !interactive {
			if err = action.SignInAccount(cmd, args); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.SignInAccountInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var accountLogoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Log out the operator",
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		action.SetupOutput(cmd)

		if !interactive {
			if err = action.SignOutAccount(cmd, args); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err := action.SignOutAccountInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

func init() {
	accountCmd.AddCommand(accountSignupSubCmd)
	accountSignupSubCmd.Flags().String("api-server-url", "https://api.cubbit.eu/iam", "Api server URL")
	accountSignupSubCmd.Flags().String("email", "", "Email Address")
	accountSignupSubCmd.Flags().String("password", "", "Password")
	accountSignupSubCmd.Flags().String("first-name", "", "First Name")
	accountSignupSubCmd.Flags().String("last-name", "", "Last Name")
	accountSignupSubCmd.Flags().String("tenant-id", "", "Tenant id")

	accountCmd.AddCommand(accountLoginSubCmd)
	accountLoginSubCmd.Flags().String("api-server-url", "https://api.cubbit.eu/iam", "Api server url")
	accountLoginSubCmd.Flags().StringP("email", "e", "", "Email Address")
	accountLoginSubCmd.Flags().StringP("password", "p", "", "")
	accountLoginSubCmd.Flags().String("code", "", "Two factor authentication code")
	accountLoginSubCmd.Flags().String("tenant-id", "", "Tenant id")

	accountCmd.AddCommand(accountLogoutCmd)

	if ENABLE_ACCOUNT_SECTION {
		rootCmd.AddCommand(accountCmd)
	}
}
