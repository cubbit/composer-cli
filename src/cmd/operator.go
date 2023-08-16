package cmd

import (
	"fmt"

	"github.com/cubbit/cubbit/client/cli/src/action"
	"github.com/spf13/cobra"
)

var operatorCmd = &cobra.Command{
	Use:   "operator",
	Short: "Execute commands in operator sections",
}

var signupSubCmd = &cobra.Command{
	Use:   "signup",
	Short: "Create a new operator",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = action.CreateOperator(cmd); err != nil {
				fmt.Println(err)
			}
		} else {
			if err = action.CreateOperatorInteractive(cmd); err != nil {
				fmt.Println(err)
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

	rootCmd.AddCommand(operatorCmd)
}
