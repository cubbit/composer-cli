package cmd_gateway

import (
	servicegateway "github.com/cubbit/composer-cli/src/service/gateway"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func NewGatewaySubCmdCreate(
	gatewayService servicegateway.GatewayServiceInterface,
) *cobra.Command {
	var gatewayCreateCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a new gateway using the Gateway V5 API",
		Long: `Create a new gateway using the Gateway V5 API.

This command submits a gateway creation request and returns immediately with a gateway ID.

Swarm-RC Format:
   Each --swarm-rc flag specifies a swarm, redundancy class, and default flag in the format: swarm-id:rc-id:is-default

Examples:
   # Create a gateway with required flags
   cubbit gateway create \
     --name "my-gateway" \
     --slug "my-gateway" \
     --cluster-id "<cluster-uuid>" \
     --ingress-type "manual" \
     --swarm-rc "<swarm-uuid>:<rc-uuid>:true"

   # Create a gateway with optional description
   cubbit gateway create \
     --name "my-gateway" \
     --slug "my-gateway" \
     --cluster-id "<cluster-uuid>" \
     --description "Production gateway" \
     --ingress-type "manual" \
     --swarm-rc "<swarm-uuid>:<rc-uuid>:true" \
     --swarm-rc "<swarm-uuid-2>:<rc-uuid-2>:false"

   # Create a gateway in interactive mode
   cubbit gateway create --interactive`,
		PreRun: func(cmd *cobra.Command, args []string) {
			interactive, _ := cmd.Flags().GetBool("interactive")
			if !interactive {
				cmd.MarkFlagRequired("name")
				cmd.MarkFlagRequired("slug")
				cmd.MarkFlagRequired("cluster-id")
				cmd.MarkFlagRequired("ingress-type")
				cmd.MarkFlagRequired("swarm-rc")
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			if err := gatewayService.Create(cmd, args); err != nil {
				utils.PrintErrorWithWriter(cmd.ErrOrStderr(), err)
			}
		},
	}

	gatewayCreateCmd.Flags().Bool("interactive", false, "Run in interactive mode")
	gatewayCreateCmd.Flags().String("name", "", "Name of the gateway (required)")
	gatewayCreateCmd.Flags().String("slug", "", "URL-friendly slug for the gateway (required)")
	gatewayCreateCmd.Flags().String("cluster-id", "", "UUID of the cluster where the gateway will be created (required)")
	gatewayCreateCmd.Flags().String("description", "", "Optional description of the gateway")
	gatewayCreateCmd.Flags().String("ingress-type", "", "Type of ingress configuration (required)")
	gatewayCreateCmd.Flags().StringArray("swarm-rc", []string{}, "Swarm and redundancy class pair in format swarm-id:rc-id:is-default. Can be specified multiple times. (required)")

	return gatewayCreateCmd
}
