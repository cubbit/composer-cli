// Package cmd provides CLI commands for managing redundancy classes.
package cmd

import (
	"fmt"
	"os"

	"github.com/cubbit/cubbit/client/cli/src/action"
	"github.com/cubbit/cubbit/client/cli/utils"
	"github.com/spf13/cobra"
)

var redundancyClassCmd = &cobra.Command{
	Use:   "rc",
	Short: "Execute commands in redundancy class section",
}

var CreateRedundancyClassSubCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new redundancy class",
	PreRun: func(cmd *cobra.Command, args []string) {

		cmd.MarkFlagsRequiredTogether("name", "inner-k", "inner-n", "outer-k", "outer-n", "anti-affinity-group", "nexuses")

		outerK, _ := cmd.Flags().GetInt("outer-k")
		outerN, _ := cmd.Flags().GetInt("outer-n")
		nexuses, _ := cmd.Flags().GetStringSlice("nexuses")
		if len(nexuses) == 0 || len(nexuses) != outerK+outerN {
			fmt.Println("Error: invalid number of nexuses provided, expected outer-k + outer-n nexuses")
			cmd.Usage()
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.CreateRedundancyClass(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

var DescribeRedundancyClassesSubCmd = &cobra.Command{
	Use:   "describe",
	Short: "describe a redundancy class",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("rc-id")
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.DescribeRedundancyClass(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

var listRedundancyClassesSubCmd = &cobra.Command{
	Use:   "list",
	Short: "list redundancy classes",
	PreRun: func(cmd *cobra.Command, args []string) {
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.ListRedundancyClasses(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

var CheckRedundancyClassStatusSubCmd = &cobra.Command{
	Use:   "status",
	Short: "check redundancy class status",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("rc-id")
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.CheckRedundancyClassStatus(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

var ExpandRedundancyClassSubCmd = &cobra.Command{
	Use:   "expand",
	Short: "expand a redundancy class",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("rc-id")
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.ExpandRedundancyClass(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

func init() {
	redundancyClassCmd.AddCommand(CreateRedundancyClassSubCmd)
	CreateRedundancyClassSubCmd.Flags().String("name", "", "Name of the redundancy class")
	CreateRedundancyClassSubCmd.Flags().String("description", "", "Description of the redundancy class")
	CreateRedundancyClassSubCmd.Flags().Int("inner-n", 1, "Inner N")
	CreateRedundancyClassSubCmd.Flags().Int("inner-k", 0, "Inner K")
	CreateRedundancyClassSubCmd.Flags().Int("outer-n", 1, "Outer N")
	CreateRedundancyClassSubCmd.Flags().Int("outer-k", 0, "Outer K")
	CreateRedundancyClassSubCmd.Flags().Int("anti-affinity-group", 1, "Anti affinity group")
	CreateRedundancyClassSubCmd.Flags().StringSlice("nexuses", []string{}, "List of nexuses IDs")

	redundancyClassCmd.AddCommand(DescribeRedundancyClassesSubCmd)
	DescribeRedundancyClassesSubCmd.Flags().String("rc-id", "", "ID of the redundancy class")

	redundancyClassCmd.AddCommand(listRedundancyClassesSubCmd)
	listRedundancyClassesSubCmd.Flags().String("sort", "", "Sorts the output based on the given field")
	listRedundancyClassesSubCmd.Flags().String("filter", "", "Filters the output based on the given field")

	redundancyClassCmd.AddCommand(CheckRedundancyClassStatusSubCmd)
	CheckRedundancyClassStatusSubCmd.Flags().String("rc-id", "", "ID of the redundancy class")

	redundancyClassCmd.AddCommand(ExpandRedundancyClassSubCmd)
	ExpandRedundancyClassSubCmd.Flags().String("rc-id", "", "ID of the redundancy class")
	ExpandRedundancyClassSubCmd.Flags().Bool("dry-run", false, "Perform a dry run without making changes")

	rootCmd.AddCommand(redundancyClassCmd)
	redundancyClassCmd.PersistentFlags().String("swarm-id", "", "ID of the swarm")
	redundancyClassCmd.MarkPersistentFlagRequired("swarm-id")
}
