package cmd_swarm

import (
	"github.com/cubbit/composer-cli/src/service"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func NewSwarmSubCmdList(
	swarmService service.SwarmServiceInterface,
) *cobra.Command {
	var swarmListCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List all swarms using the v5 API",
		Run: func(cmd *cobra.Command, args []string) {
			if err := swarmService.List(cmd, args); err != nil {
				utils.PrintErrorWithWriter(cmd.ErrOrStderr(), err)
			}
		},
	}

	return swarmListCmd
}
