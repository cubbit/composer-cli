package cmd

import (
	"github.com/cubbit/cubbit/client/cli/src/action"
	"github.com/cubbit/cubbit/client/cli/src/tui"
	"github.com/cubbit/cubbit/client/cli/utils"
	"github.com/spf13/cobra"
)

var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Execute commands in project sections",
}

var createProjectSubCmd = &cobra.Command{
	Use:     "create",
	Short:   "create a new project",
	Aliases: []string{"new"},
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			cmd.MarkFlagRequired("name")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = tui.Send(cmd, args, action.CreateProject); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.CreateProjectInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

func init() {
	projectCmd.AddCommand(createProjectSubCmd)
	createProjectSubCmd.Flags().String("name", "", "Name of the tenant")
	createProjectSubCmd.Flags().String("description", "", "Description of the tenant")
	createProjectSubCmd.Flags().String("image-url", "", "Image URL of the tenant")

	if ENABLE_ACCOUNT_SECTION {
		rootCmd.AddCommand(projectCmd)
	}
}
