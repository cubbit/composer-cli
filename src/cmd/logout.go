package cmd

import (
	"fmt"

	"github.com/cubbit/cubbit/client/cli/src/action"
	"github.com/cubbit/cubbit/client/cli/src/tui"
	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Log out the operator",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = tui.Send(cmd, action.SignOutOperator); err != nil {
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
