// Package cmd provides CLI commands for managing agents.
package cmd

import (
	"fmt"
	"os"

	"github.com/cubbit/cubbit/client/cli/src/action"
	"github.com/cubbit/cubbit/client/cli/utils"
	"github.com/spf13/cobra"
)

var agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "Execute commands in agent sections",
}

var createAgentSubCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new agent",
	PreRun: func(cmd *cobra.Command, args []string) {

		cmd.MarkFlagRequired("nexus-id")
		cmd.MarkFlagRequired("node-id")

		batch, _ := cmd.Flags().GetBool("batch")
		if batch {
			file, _ := cmd.Flags().GetString("file")
			if file == "" {
				fmt.Println("Error: --file flag is required when using --batch mode.")
				cmd.Usage()
				os.Exit(1)
			}

			if _, err := os.Stat(file); os.IsNotExist(err) {
				fmt.Println("Error: file does not exist:", file)
				os.Exit(1)
			}
			return
		}

		cmd.MarkFlagRequired("agent-port")
		cmd.MarkFlagRequired("agent-disk")
		cmd.MarkFlagRequired("agent-mount-point")

		agentPort := cmd.Flags().Lookup("agent-port")
		if agentPort != nil && !agentPort.Changed {
			fmt.Println("Error: --agent-port must be explicitly provided.")
			cmd.Usage()
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		batch, _ := cmd.Flags().GetBool("batch")
		if batch {
			if err := action.CreateAgentBatch(cmd, args); err != nil {
				utils.PrintError(err)
			}
			return
		}

		if err := action.CreateAgent(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

var editAgentSubCmd = &cobra.Command{
	Use:   "edit",
	Short: "edit an agent",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("nexus-id")
		cmd.MarkFlagRequired("node-id")
		cmd.MarkFlagRequired("agent-id")
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.EditAgent(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

var describeAgentSubCmd = &cobra.Command{
	Use:   "describe",
	Short: "describe an agent",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("nexus-id")
		cmd.MarkFlagRequired("node-id")
		cmd.MarkFlagRequired("agent-id")
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.DescribeAgent(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

var listAgentsSubCmd = &cobra.Command{
	Use:   "list",
	Short: "list agents",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagsMutuallyExclusive("nexus-id", "redundancy-class-id")
		cmd.MarkFlagsMutuallyExclusive("node-id", "redundancy-class-id")

		allowedSortingKeys := []string{"id", "node_id", "port", "created_at"}
		sort, _ := cmd.Flags().GetString("sort")

		if sort != "" && !utils.Contains(allowedSortingKeys, sort) {
			fmt.Println("Error: invalid sort key provided, allowed keys are: id, node_id", "port", "created_at")
			cmd.Usage()
			os.Exit(1)
		}

		filter, _ := cmd.Flags().GetString("filter")
		if filter != "" {
			if !utils.IsValidFilter(filter) {
				fmt.Println("Error: invalid filter provided, allowed format is: key:value key:value ...")
				cmd.Usage()
				os.Exit(1)
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.ListAgents(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

var removeAgentSubCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove an agent",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("nexus-id")
		cmd.MarkFlagRequired("node-id")
		cmd.MarkFlagRequired("agent-id")
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.RemoveAgent(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

var checkAgentStatusSubCmd = &cobra.Command{
	Use:   "status",
	Short: "check the status of an agent",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("nexus-id")
		cmd.MarkFlagRequired("node-id")
		cmd.MarkFlagRequired("agent-id")
	},

	Run: func(cmd *cobra.Command, args []string) {
		if err := action.CheckAgentStatus(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

func init() {
	agentCmd.AddCommand(createAgentSubCmd)
	createAgentSubCmd.Flags().Int("agent-port", 0, "Port of the agent")
	createAgentSubCmd.Flags().String("agent-disk", "", "Disk of the agent")
	createAgentSubCmd.Flags().String("agent-mount-point", "", "Mount point of the agent")
	createAgentSubCmd.Flags().String("agent-features", "", "Features of the agent")
	createAgentSubCmd.Flags().Bool("batch", false, "Create multiple agents from a batch file")
	createAgentSubCmd.Flags().String("file", "", "Path to the JSON file containing agent definitions")

	agentCmd.AddCommand(describeAgentSubCmd)
	describeAgentSubCmd.Flags().String("agent-id", "", "ID of the agent")

	agentCmd.AddCommand(listAgentsSubCmd)
	listAgentsSubCmd.Flags().String("sort", "", "Sorts the output based on the given field")
	listAgentsSubCmd.Flags().String("filter", "", "Filters the output based on the given field")
	listAgentsSubCmd.Flags().String("redundancy-class-id", "", "ID of the redundancy class")

	agentCmd.AddCommand(editAgentSubCmd)
	editAgentSubCmd.Flags().String("agent-id", "", "ID of the agent")
	editAgentSubCmd.Flags().Int("agent-port", 0, "Port of the agent")
	editAgentSubCmd.Flags().String("agent-disk", "", "Disk of the agent")
	editAgentSubCmd.Flags().String("agent-mount-point", "", "Mount point of the agent")
	editAgentSubCmd.Flags().String("agent-features", "", "Features of the agent")

	agentCmd.AddCommand(removeAgentSubCmd)
	removeAgentSubCmd.Flags().String("nexus-id", "", "ID of the nexus")
	removeAgentSubCmd.Flags().String("node-id", "", "ID of the node")
	removeAgentSubCmd.Flags().String("agent-id", "", "ID of the agent")

	agentCmd.AddCommand(checkAgentStatusSubCmd)
	checkAgentStatusSubCmd.Flags().String("nexus-id", "", "ID of the nexus")
	checkAgentStatusSubCmd.Flags().String("node-id", "", "ID of the node")
	checkAgentStatusSubCmd.Flags().String("agent-id", "", "ID of the agent")

	rootCmd.AddCommand(agentCmd)
	agentCmd.PersistentFlags().String("swarm-id", "", "ID of the swarm")
	agentCmd.MarkPersistentFlagRequired("swarm-id")
	agentCmd.PersistentFlags().String("nexus-id", "", "ID of the nexus")
	agentCmd.PersistentFlags().String("node-id", "", "ID of the node")

}
