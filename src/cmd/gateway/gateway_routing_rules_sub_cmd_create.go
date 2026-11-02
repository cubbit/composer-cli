package cmd_gateway

import (
	servicegateway "github.com/cubbit/composer-cli/src/service/gateway"
	"github.com/spf13/cobra"
)

func NewGatewayRoutingRulesSubCmdCreate(
	gatewayService servicegateway.GatewayServiceInterface,
) *cobra.Command {
	routingRuleCreateCmd := &cobra.Command{
		Use:   "create --gateway-id GATEWAY_ID --file-path PATH",
		Short: "Create a gateway routing rule using the V5 API",
		Long: `Create a routing rule on a gateway using the Gateway V5 API.

The rule document is a JSON object read from --file-path; use "-" to read it
from stdin. The document must contain at least a non-empty name, an
inputStorageClass and an action.type. Unknown fields are forwarded to the
server, so newer fields are preserved. On success only a confirmation message
is printed; use "get --id" to inspect the created document.

On a 422 the server returns the validator message (for example an unknown
storage class), which is printed as-is.

The gateway is identified by the required --gateway-id flag.

Examples:
  cubbit gateway routing-rules create --gateway-id <gateway-id> --file-path rule.json
  cat rule.json | cubbit gateway routing-rules create --gateway-id <gateway-id> --file-path -`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return gatewayService.CreateRoutingRule(cmd, args)
		},
	}

	routingRuleCreateCmd.Flags().String("gateway-id", "", "ID of the gateway (required)")
	_ = routingRuleCreateCmd.MarkFlagRequired("gateway-id")
	routingRuleCreateCmd.Flags().String("file-path", "", "Path to the routing rule JSON document, or - for stdin (required)")
	_ = routingRuleCreateCmd.MarkFlagRequired("file-path")

	return routingRuleCreateCmd
}
