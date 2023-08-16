package cmd

import (
	"fmt"

	"github.com/cubbit/cubbit/client/cli/src/action"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login the operator",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = action.SignInOperator(cmd); err != nil {
				fmt.Println(err)
			}
		} else {
			if err = action.SignInOperatorInteractive(cmd); err != nil {
				fmt.Println(err)
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
