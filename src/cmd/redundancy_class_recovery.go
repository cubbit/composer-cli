// Package cmd provides CLI commands for managing redundancy class recovery.
package cmd

import (
	"github.com/cubbit/cubbit/client/cli/src/action"
	"github.com/cubbit/cubbit/client/cli/utils"
	"github.com/spf13/cobra"
)

var redundancyClassRecoveryCmd = &cobra.Command{
	Use:   "recovery",
	Short: "Execute commands in redundancy class recovery sections",
}

var RecoverRedundancyClassSubCmd = &cobra.Command{
	Use:   "start",
	Short: "start redundancy class recovery",
	PreRun: func(cmd *cobra.Command, args []string) {
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.RecoverRedundancyClass(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

var CheckRedundancyClassRecoveryStatusSubCmd = &cobra.Command{
	Use:   "status",
	Short: "check redundancy class recovery status",
	PreRun: func(cmd *cobra.Command, args []string) {
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.CheckRedundancyClassRecoveryStatus(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

func init() {
	redundancyClassRecoveryCmd.AddCommand(RecoverRedundancyClassSubCmd)
	RecoverRedundancyClassSubCmd.Flags().Bool("dry-run", false, "Dry run mode")

	redundancyClassRecoveryCmd.AddCommand(CheckRedundancyClassRecoveryStatusSubCmd)

	redundancyClassCmd.AddCommand(redundancyClassRecoveryCmd)
	redundancyClassRecoveryCmd.PersistentFlags().String("rc-id", "", "ID of the redundancy class")
	redundancyClassRecoveryCmd.MarkPersistentFlagRequired("rc-id")
}
