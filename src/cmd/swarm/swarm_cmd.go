package cmd_swarm

import (
	"github.com/cubbit/composer-cli/src/service"
	"github.com/spf13/cobra"
)

func NewSwarmCmd(
	swarmService service.SwarmServiceInterface,
) *cobra.Command {
	var swarmCmd = &cobra.Command{
		Use:   "swarm",
		Short: "Execute commands in swarm sections",
	}

	swarmCreateCmd := NewSwarmSubCmdCreate(swarmService)
	swarmDescribeCmd := NewSwarmSubCmdDescribe(swarmService)
	swarmListCmd := NewSwarmSubCmdList(swarmService)
	swarmCmd.AddCommand(swarmCreateCmd, swarmDescribeCmd, swarmListCmd)

	return swarmCmd
}
