package cmd_swarm

import (
	"github.com/cubbit/composer-cli/src/service"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func NewSwarmSubCmdCreate(
	swarmService service.SwarmServiceInterface,
) *cobra.Command {
	var swarmCreateCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a new swarm using Swarm Creation V5 API",
		Long: `Create a new swarm using the asynchronous Swarm Creation V5 API.

This command submits a swarm creation request and returns immediately with a process ID.
The swarm creation happens in the background.

Nexus Format:
   Each --nexus flag specifies a cluster and its nodes in the format: cluster-id:node-id1,node-id2

   The CLI automatically fetches cluster details from the backend and includes all volumes that are not yet used.
   (disks) from each specified node. For virtual clusters, use virtual node IDs.

Examples:
   # Create a swarm with a physical cluster (all volumes on specified nodes are included)
   cubbit swarm create \
     --name "my-swarm" \
     --description "Production swarm" \
     --nexus <cluster-id>:<node-id1> \
     --redundancy-class '{"name":"rc-1","outer_n":1,"outer_k":0,"inner_n":4,"inner_k":2,"anti_affinity_group":1,"cluster_ids":["<cluster-id>"]}'

   # Create a swarm with multiple nodes from the same cluster
   cubbit swarm create \
     --name "my-swarm" \
     --nexus <cluster-id>:<node-1>,<node-2>,<node-3> \
     --redundancy-class '{"name":"rc-1","outer_n":1,"outer_k":0,"inner_n":4,"inner_k":2,"anti_affinity_group":1,"cluster_ids":["<cluster-id>"]}'

   # Create a swarm with multiple clusters
   cubbit swarm create \
     --name "multi-cluster-swarm" \
     --nexus <cluster-1>:<node-1>,<node-2> \
     --nexus <cluster-2>:<node-3> \
     --redundancy-class '{"name":"rc-1","outer_n":1,"outer_k":1,"inner_n":4,"inner_k":2,"anti_affinity_group":1,"cluster_ids":["<cluster-1>","<cluster-2>"]}'`,
		PreRun: func(cmd *cobra.Command, args []string) {
			cmd.MarkFlagRequired("name")
			cmd.MarkFlagRequired("nexus")
		},
		Run: func(cmd *cobra.Command, args []string) {
			if err := swarmService.Create(cmd, args); err != nil {
				utils.PrintErrorWithWriter(cmd.ErrOrStderr(), err)
			}
		},
	}

	swarmCreateCmd.Flags().String("name", "", "Name of the swarm (required, 3-63 characters)")
	swarmCreateCmd.Flags().String("description", "", "Optional description of the swarm")
	swarmCreateCmd.Flags().StringArray("nexus", []string{}, "Nexus specification in format cluster-id:node-id1,node-id2. Can be specified multiple times. All volumes on specified nodes are automatically included.")
	swarmCreateCmd.Flags().StringArray("redundancy-class", []string{}, "Redundancy class configuration in JSON format. Can be specified multiple times. Format: {\"name\":\"...\",\"outer_n\":4,\"outer_k\":2,\"inner_n\":4,\"inner_k\":2,\"cluster_ids\":[...]}")
	swarmCreateCmd.Flags().String("redundancy-class-file", "", "Path to a JSON file containing an array of redundancy class configurations")

	return swarmCreateCmd
}
