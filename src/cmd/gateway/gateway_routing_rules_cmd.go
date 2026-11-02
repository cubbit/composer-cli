package cmd_gateway

import (
	servicegateway "github.com/cubbit/composer-cli/src/service/gateway"
	"github.com/spf13/cobra"
)

func NewGatewayRoutingRulesCmd(
	gatewayService servicegateway.GatewayServiceInterface,
) *cobra.Command {
	routingRulesCmd := &cobra.Command{
		Use:     "routing-rules",
		Aliases: []string{"rr"},
		Short:   "Manage gateway routing rules",
		Long: `Manage the smart data placement routing rules of a gateway.

A gateway can have zero or more routing rules, evaluated in order. Each rule
has an id, a name, an optional description, an enabled flag, the input storage
class it applies to, optional conditions and an action.

Use "list" to discover the rule IDs of a gateway and "get --output json" to
capture a full rule document. Redirect that output to a file and pass it back
with "update --file-path" to round-trip the document, including fields this
CLI does not model.`,
	}

	routingRulesCmd.AddCommand(NewGatewayRoutingRulesSubCmdList(gatewayService))
	routingRulesCmd.AddCommand(NewGatewayRoutingRulesSubCmdGet(gatewayService))
	routingRulesCmd.AddCommand(NewGatewayRoutingRulesSubCmdCreate(gatewayService))
	routingRulesCmd.AddCommand(NewGatewayRoutingRulesSubCmdUpdate(gatewayService))
	routingRulesCmd.AddCommand(NewGatewayRoutingRulesSubCmdDelete(gatewayService))

	return routingRulesCmd
}
