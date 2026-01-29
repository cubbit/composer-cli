package cmd_agent

import (
	"fmt"

	"github.com/cubbit/composer-cli/src/service"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func NewAgentSubCmdList(
	newAgentService service.AgentServiceInterface,
) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list agents",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			cmd.MarkFlagsMutuallyExclusive("nexus-id", "redundancy-class-id")
			cmd.MarkFlagsMutuallyExclusive("node-id", "redundancy-class-id")

			allowedSortingKeys := []string{"id", "node_id", "port", "created_at"}
			sort, _ := cmd.Flags().GetString("sort")

			if sort != "" && !utils.Contains(allowedSortingKeys, sort) {
				cmd.Usage()
				return fmt.Errorf("error: invalid sort key provided, allowed keys are: id, node_id, port, created_at")
			}

			filter, _ := cmd.Flags().GetString("filter")
			if filter != "" {
				if !utils.IsValidFilter(filter) {
					cmd.Usage()
					return fmt.Errorf("error: invalid filter provided, allowed format is: key:value key:value")
				}
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			if err := newAgentService.ListAgents(cmd, args); err != nil {
				utils.PrintError(err)
			}
		},
	}

	cmd.Flags().String("sort", "", "Sorts the output based on the given field")
	cmd.Flags().String("filter", "", "Filters the output based on the given field, allowed format is key:value")
	cmd.Flags().String("redundancy-class-id", "", "ID of the redundancy class")

	return cmd
}
