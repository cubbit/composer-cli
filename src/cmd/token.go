package cmd

import (
	"github.com/cubbit/cubbit/client/cli/src/action"
	"github.com/cubbit/cubbit/client/cli/src/tui"
	"github.com/cubbit/cubbit/client/cli/utils"
	"github.com/spf13/cobra"
)

var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Generate access token",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = tui.Send(cmd, args, action.GenerateAccessToken); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err := action.GenerateAccessTokenInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(tokenCmd)
}
