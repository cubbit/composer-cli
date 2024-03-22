package cmd

import (
	"github.com/cubbit/cubbit/client/cli/src/action"
	"github.com/cubbit/cubbit/client/cli/utils"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of the CLI",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if err = action.GetCliVersion(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	if ENABLE_ACCOUNT_SECTION {
		rootCmd.AddCommand(projectCmd)
	}
}
