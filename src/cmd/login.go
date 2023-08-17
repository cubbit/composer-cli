package cmd

import (
	"github.com/cubbit/cubbit/client/cli/src/action"
	"github.com/cubbit/cubbit/client/cli/src/tui"
	"github.com/cubbit/cubbit/client/cli/utils"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login the operator",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = tui.Send(cmd, action.SignInOperator); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.SignInOperatorInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().String("api-server-url", "https://api.cubbit.eu/iam", "Api server url")
	loginCmd.Flags().StringP("email", "e", "", "Email Address")
	loginCmd.Flags().StringP("password", "p", "", "")
	loginCmd.Flags().String("code", "", "Two factor authentication code")
}
