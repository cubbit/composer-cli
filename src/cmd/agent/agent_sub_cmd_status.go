package cmd_agent

import (
	"github.com/cubbit/composer-cli/src/service"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func NewAgentSubCmdStatus(
	newAgentService service.AgentServiceInterface,
) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "status",
		Short: "check the status of an agent",
		PreRun: func(cmd *cobra.Command, args []string) {
			cmd.MarkFlagRequired("nexus-id")
			cmd.MarkFlagRequired("node-id")
			cmd.MarkFlagRequired("agent-id")
		},

		Run: func(cmd *cobra.Command, args []string) {
			if err := newAgentService.CheckAgentStatus(cmd, args); err != nil {
				utils.PrintError(err)
			}
		},
	}

	cmd.Flags().String("nexus-id", "", "ID of the nexus")
	cmd.Flags().String("node-id", "", "ID of the node")
	cmd.Flags().String("agent-id", "", "ID of the agent")

	return cmd
}
