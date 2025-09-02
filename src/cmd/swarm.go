// Package cmd provides CLI commands for managing swarms.
package cmd

import (
	"fmt"
	"os"

	"github.com/cubbit/cubbit/client/cli/src/action"
	"github.com/cubbit/cubbit/client/cli/utils"
	"github.com/spf13/cobra"
)

var swarmCmd = &cobra.Command{
	Use:   "swarm",
	Short: "Execute commands in swarm sections",
}

var createSwarmSubCmd = &cobra.Command{
	Use:     "create",
	Short:   "create a new swarm",
	Aliases: []string{"new"},
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("name")
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.CreateSwarm(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

var describeSwarmSubCmd = &cobra.Command{
	Use:   "describe",
	Short: "describe a swarm",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("swarm-id")
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.DescribeSwarm(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

var editSwarmSubCmd = &cobra.Command{
	Use:   "edit",
	Short: "edit a swarm",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("swarm-id")

		swarmName, _ := cmd.Flags().GetString("name")
		swarmDescription, _ := cmd.Flags().GetString("description")

		if swarmName == "" && swarmDescription == "" {
			fmt.Println("Error: at least one of the two required flags --name or --description should be provided.")
			cmd.Usage()
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.EditSwarm(cmd, args...); err != nil {
			utils.PrintError(err)
		}
	},
}

var checkSwarmStatusSubCmd = &cobra.Command{
	Use:   "status",
	Short: "check the status of a swarm",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("swarm-id")
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.CheckSwarmStatus(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

var listSwarmSubCmd = &cobra.Command{
	Use:     "list",
	Short:   "list swarms",
	Aliases: []string{"ls"},
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.ListSwarms(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

var removeSwarmSubCmd = &cobra.Command{
	Use:     "remove",
	Short:   "remove a swarm",
	Aliases: []string{"rm"},
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("swarm-id")
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.RemoveSwarm(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

func init() {
	swarmCmd.AddCommand(createSwarmSubCmd)
	createSwarmSubCmd.Flags().String("name", "", "Name of the swarm")
	createSwarmSubCmd.Flags().String("description", "", "Description of the swarm")

	swarmCmd.AddCommand(describeSwarmSubCmd)
	describeSwarmSubCmd.Flags().String("swarm-id", "", "ID of the swarm")

	swarmCmd.AddCommand(editSwarmSubCmd)
	editSwarmSubCmd.Flags().String("swarm-id", "", "ID of the swarm")
	editSwarmSubCmd.Flags().String("name", "", "Name of the swarm")
	editSwarmSubCmd.Flags().String("description", "", "Description of the swarm")

	swarmCmd.AddCommand(checkSwarmStatusSubCmd)
	checkSwarmStatusSubCmd.Flags().String("swarm-id", "", "ID of the swarm")

	swarmCmd.AddCommand(listSwarmSubCmd)
	listSwarmSubCmd.Flags().String("sort", "", "Sorts the output based on the given field")
	listSwarmSubCmd.Flags().String("filter", "", "Filters the output based on the given field")

	swarmCmd.AddCommand(removeSwarmSubCmd)
	removeSwarmSubCmd.Flags().String("swarm-id", "", "ID of the swarm")

	rootCmd.AddCommand(swarmCmd)
}
