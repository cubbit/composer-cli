package action

import (
	"github.com/cubbit/cubbit/client/cli/utils"
	"github.com/spf13/cobra"
)

func SetupOutput(cmd *cobra.Command) {
	humanMode, _ := cmd.Flags().GetBool("human")
	interactive, _ := cmd.Flags().GetBool("interactive")

	if humanMode || interactive {
		utils.SetOutputMode(utils.OutputHuman)
	} else {
		utils.SetOutputMode(utils.OutputQuiet)
	}
}
