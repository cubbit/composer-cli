package cmd_agent

import (
	"github.com/cubbit/composer-cli/src/service"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func NewAgentSubCmdEdit(
	newAgentService service.AgentServiceInterface,
) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit",
		Short: "edit an agent",
		PreRun: func(cmd *cobra.Command, args []string) {
			cmd.MarkFlagRequired("nexus-id")
			cmd.MarkFlagRequired("node-id")
			cmd.MarkFlagRequired("agent-id")
		},
		Run: func(cmd *cobra.Command, args []string) {
			if err := newAgentService.EditAgent(cmd, args); err != nil {
				utils.PrintError(err)
			}
		},
	}

	cmd.Flags().String("agent-id", "", "ID of the agent")
	cmd.Flags().Int("agent-port", 0, "Port of the agent")
	cmd.Flags().String("agent-disk", "", "Disk of the agent")
	cmd.Flags().String("agent-mount-point", "", "Mount point of the agent")
	cmd.Flags().String("agent-features", "", "Features of the agent")

	return cmd
}
