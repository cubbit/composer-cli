// Package cmd provides CLI commands for managing agents.
package cmd

import (
	"fmt"

	"github.com/cubbit/composer-cli/src/action"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

var agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "Execute commands in agent sections",
}

var createAgentSubCmd = &cobra.Command{
	Use:     "create",
	Short:   "create a new agent or multiple agents from a batch file",
	Example: "cubbit agent create --swarm-id <swarm-id> --nexus-id <nexus-id> --node-id <node-id> --batch --file ./batch.json\ncubbit agent create --swarm-id <swarm-id> --nexus-id <nexus-id> --node-id <node-id> --agent-port 8080 --agent-disk /dev/sdb --agent-mount-point /mnt/agent",
	PreRun: func(cmd *cobra.Command, args []string) {

		cmd.MarkFlagRequired("swarm-id")
		cmd.MarkFlagRequired("nexus-id")
		cmd.MarkFlagRequired("node-id")

		cmd.MarkFlagsOneRequired("batch", "agent-port")
		cmd.MarkFlagsRequiredTogether("batch", "file")
		cmd.MarkFlagsRequiredTogether("agent-port", "agent-disk", "agent-mount-point")

		cmd.MarkFlagsMutuallyExclusive("batch", "agent-port")
		cmd.MarkFlagsMutuallyExclusive("batch", "agent-disk")
		cmd.MarkFlagsMutuallyExclusive("batch", "agent-mount-point")
		cmd.MarkFlagsMutuallyExclusive("file", "agent-port")
		cmd.MarkFlagsMutuallyExclusive("file", "agent-disk")
		cmd.MarkFlagsMutuallyExclusive("file", "agent-mount-point")
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
		cmd.MarkFlagRequired("swarm-id")
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
		cmd.MarkFlagRequired("swarm-id")
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
	Use:     "list",
	Short:   "list agents in a node or redundancy class",
	Example: "cubbit agent list --swarm-id <swarm-id> --nexus-id <nexus-id> --node-id <node-id>\ncubbit agent list --swarm-id <swarm-id> --redundancy-class-id <redundancy-class-id>",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("swarm-id")
		cmd.MarkFlagsOneRequired("redundancy-class-id", "node-id")

		cmd.MarkFlagsRequiredTogether("nexus-id", "node-id")
		cmd.MarkFlagsMutuallyExclusive("nexus-id", "redundancy-class-id")
		cmd.MarkFlagsMutuallyExclusive("node-id", "redundancy-class-id")
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		allowedSortingKeys := []string{"id", "node_id", "port", "created_at"}
		sort, _ := cmd.Flags().GetString("sort")
		if sort != "" && !utils.Contains(allowedSortingKeys, sort) {
			msg := "Error: invalid sort key provided, allowed keys are: id, node_id port created_at"
			fmt.Println(msg)
			return fmt.Errorf(msg)
		}

		filter, _ := cmd.Flags().GetString("filter")
		if filter != "" && !utils.IsValidFilter(filter) {
			msg := "Error: invalid filter provided, allowed format is: key:value key:value ..."
			fmt.Println(msg)
			return fmt.Errorf(msg)
		}

		if err := action.ListAgents(cmd, args); err != nil {
			utils.PrintError(err)
			return nil
		}

		return nil
	},
}

var removeAgentSubCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove an agent",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("swarm-id")
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
		cmd.MarkFlagRequired("swarm-id")
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
	listAgentsSubCmd.Flags().String("filter", "", "Filters the output based on the given field, allowed format is key:value")
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
	agentCmd.PersistentFlags().String("nexus-id", "", "ID of the nexus")
	agentCmd.PersistentFlags().String("node-id", "", "ID of the node")
}
