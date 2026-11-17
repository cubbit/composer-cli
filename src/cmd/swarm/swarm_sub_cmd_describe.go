package cmd_swarm

import (
	"github.com/cubbit/composer-cli/src/service"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func NewSwarmSubCmdDescribe(
	swarmService service.SwarmServiceInterface,
) *cobra.Command {
	var swarmID string
	var swarmName string

	var swarmDescribeCmd = &cobra.Command{
		Use:     "describe [SWARM_ID]",
		Aliases: []string{"info", "show"},
		Short:   "Describe a swarm using the v5 API",
		Args:    cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if err := swarmService.Describe(cmd, args); err != nil {
				utils.PrintErrorWithWriter(cmd.ErrOrStderr(), err)
			}
		},
	}

	swarmDescribeCmd.Flags().StringVar(&swarmID, "swarm-id", "", "Swarm ID")
	swarmDescribeCmd.Flags().StringVar(&swarmName, "swarm-name", "", "Swarm name")

	return swarmDescribeCmd
}
