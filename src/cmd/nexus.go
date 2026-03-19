// Package cmd provides CLI commands for managing nexuses.
package cmd

import (
	"github.com/cubbit/composer-cli/src/action"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

var nexusCmd = &cobra.Command{
	Use:   "nexus",
	Short: "Execute commands in nexus sections",
}

var createNexusSubCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new nexus",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("swarm-id")
		cmd.MarkFlagRequired("name")
		cmd.MarkFlagRequired("location")
		cmd.MarkFlagRequired("provider-id")
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.CreateNexus(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

var editNexusSubCmd = &cobra.Command{
	Use:   "edit",
	Short: "edit a nexus",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("nexus-id")
		cmd.MarkFlagRequired("swarm-id")
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.EditNexus(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

var removeNexusSubCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove a nexus",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("nexus-id")
		cmd.MarkFlagRequired("swarm-id")
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.RemoveNexus(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

var listNexusesSubCmd = &cobra.Command{
	Use:   "list",
	Short: "list nexuses",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("swarm-id")
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.ListNexuses(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

var describeNexusSubCmd = &cobra.Command{
	Use:   "describe",
	Short: "describe a nexus",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("nexus-id")
		cmd.MarkFlagRequired("swarm-id")
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.DescribeNexus(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

var deployNexusSubCmd = &cobra.Command{
	Use:   "deploy",
	Short: "generate deploy files for a nexus",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("swarm-id")
		cmd.MarkFlagRequired("nexus-id")
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.GenerateNexusDeployFiles(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

func init() {
	nexusCmd.AddCommand(createNexusSubCmd)
	createNexusSubCmd.Flags().String("name", "", "Name of the nexus")
	createNexusSubCmd.Flags().String("description", "", "Description of the nexus")
	createNexusSubCmd.Flags().String("location", "", "Location of the nexus")
	createNexusSubCmd.Flags().String("provider-id", "", "Provider ID of the nexus")

	nexusCmd.AddCommand(editNexusSubCmd)
	editNexusSubCmd.Flags().String("nexus-id", "", "ID of the nexus")
	editNexusSubCmd.Flags().String("name", "", "Name of the nexus")
	editNexusSubCmd.Flags().String("description", "", "Description of the nexus")
	editNexusSubCmd.Flags().String("location", "", "Location of the nexus")

	nexusCmd.AddCommand(removeNexusSubCmd)
	removeNexusSubCmd.Flags().String("nexus-id", "", "ID of the nexus to remove")

	nexusCmd.AddCommand(listNexusesSubCmd)
	listNexusesSubCmd.Flags().String("sort", "", "Sorts the output based on the given field")
	listNexusesSubCmd.Flags().String("filter", "", "Filters the output based on the given field, allowed format is key:value")

	nexusCmd.AddCommand(describeNexusSubCmd)
	describeNexusSubCmd.Flags().String("nexus-id", "", "ID of the nexus")

	nexusCmd.AddCommand(deployNexusSubCmd)
	deployNexusSubCmd.Flags().String("nexus-id", "", "ID of the nexus")
	deployNexusSubCmd.Flags().String("output-dir", ".", "Directory to save the deployment files (default: current directory)")

	rootCmd.AddCommand(nexusCmd)
	nexusCmd.PersistentFlags().String("swarm-id", "", "ID of the swarm")
}
