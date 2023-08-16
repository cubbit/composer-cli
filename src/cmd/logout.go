package cmd

import (
	"fmt"

	"github.com/cubbit/cubbit/client/cli/src/action"
	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Log out the operator",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = action.SignOutOperator(cmd); err != nil {
				fmt.Println(err)
			}
		} else {
			if err := action.SignOutOperatorInteractive(cmd); err != nil {
				fmt.Println(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}
