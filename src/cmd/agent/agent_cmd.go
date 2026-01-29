package cmd_agent

import (
	"github.com/cubbit/composer-cli/src/service"
	"github.com/spf13/cobra"
)

func NewAgentCmd(
	newAgentService service.AgentServiceInterface,
) *cobra.Command {
	agentRootCmd := &cobra.Command{
		Use:   "agent",
		Short: "Execute commands in agent sections",
	}

	agentRootCmd.PersistentFlags().String("node-id", "", "ID of the node")
	agentRootCmd.PersistentFlags().String("swarm-id", "", "ID of the swarm")
	agentRootCmd.MarkPersistentFlagRequired("swarm-id")
	agentRootCmd.PersistentFlags().String("nexus-id", "", "ID of the nexus")

	createAgentSubCmd := NewAgentSubCmdCreate(newAgentService)
	agentRootCmd.AddCommand(createAgentSubCmd)

	describeAgentSubCmd := NewAgentSubCmdDescribe(newAgentService)
	agentRootCmd.AddCommand(describeAgentSubCmd)

	editAgentSubCmd := NewAgentSubCmdEdit(newAgentService)
	agentRootCmd.AddCommand(editAgentSubCmd)

	listAgentSubCmd := NewAgentSubCmdList(newAgentService)
	agentRootCmd.AddCommand(listAgentSubCmd)

	removeAgentSubCmd := NewAgentSubCmdRemove(newAgentService)
	agentRootCmd.AddCommand(removeAgentSubCmd)

	statusAgentSubCmd := NewAgentSubCmdStatus(newAgentService)
	agentRootCmd.AddCommand(statusAgentSubCmd)

	return agentRootCmd
}
