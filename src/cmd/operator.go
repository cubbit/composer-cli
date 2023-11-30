package cmd

import (
	"github.com/cubbit/cubbit/client/cli/src/action"
	"github.com/cubbit/cubbit/client/cli/src/tui"
	"github.com/cubbit/cubbit/client/cli/utils"
	"github.com/spf13/cobra"
)

var operatorCmd = &cobra.Command{
	Use:   "operator",
	Short: "Execute commands in operator sections",
}

var signupSubCmd = &cobra.Command{
	Use:   "signup",
	Short: "Create a new operator",
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			cmd.MarkFlagRequired("email")
			cmd.MarkFlagRequired("password")
			cmd.MarkFlagRequired("first-name")
			cmd.MarkFlagRequired("last-name")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = tui.Send(cmd, args, action.CreateOperator); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.CreateOperatorInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var operatorLoginSubCmd = &cobra.Command{
	Use:     "login",
	Short:   "Login the operator",
	Aliases: []string{"signin"},
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			cmd.MarkFlagRequired("email")
			cmd.MarkFlagRequired("password")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = tui.Send(cmd, args, action.SignInOperator); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.SignInOperatorInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var operatorLogoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Log out the operator",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = tui.Send(cmd, args, action.SignOutOperator); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err := action.SignOutOperatorInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var tokenSubCmd = &cobra.Command{
	Use:   "token",
	Short: "Generate access token",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = tui.Send(cmd, args, action.GenerateOperatorAccessToken); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err := action.GenerateOperatorAccessTokenInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

func init() {
	operatorCmd.AddCommand(signupSubCmd)
	signupSubCmd.Flags().String("api-server-url", "https://api.cubbit.eu/iam", "Api server URL")
	signupSubCmd.Flags().String("email", "", "Email Address")
	signupSubCmd.Flags().String("password", "", "Password")
	signupSubCmd.Flags().String("first-name", "", "First Name")
	signupSubCmd.Flags().String("last-name", "", "Last Name")
	signupSubCmd.Flags().String("secret", "", "Secret")

	operatorCmd.AddCommand(operatorLoginSubCmd)
	operatorLoginSubCmd.Flags().String("api-server-url", "https://api.cubbit.eu/iam", "Api server url")
	operatorLoginSubCmd.Flags().StringP("email", "e", "", "Email Address")
	operatorLoginSubCmd.Flags().StringP("password", "p", "", "")
	operatorLoginSubCmd.Flags().String("code", "", "Two factor authentication code")

	operatorCmd.AddCommand(operatorLogoutCmd)

	operatorCmd.AddCommand(tokenSubCmd)

	rootCmd.AddCommand(operatorCmd)
}
