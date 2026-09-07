package listconnections

import (
	"fmt"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/utils/printer"
	"github.com/cubbit/composer-cli/utils/printer/tree"
	"github.com/spf13/cobra"
)

func stateSuffix(s *api.AggregatedVerificationStateV5DTO) string {
	if s == nil {
		return ""
	}
	return fmt.Sprintf(" (%s)", s.Status)
}

func PrintConnectionsList(cmd *cobra.Command, handler configuration_handler.ConfigurationHandlerInterface, connections []api.ConnectionV5DTO) error {
	if len(connections) == 0 {
		return printer.PrintText(cmd, handler, "No connections found.\n")
	}

	nodes := make([]tree.TreeNode, 0, len(connections))
	for _, conn := range connections {
		nodes = append(nodes, buildConnectionNode(conn))
	}

	return printer.PrintTree(cmd, handler, nodes, func(n []tree.TreeNode) []tree.TreeNode { return n })
}

func buildConnectionNode(conn api.ConnectionV5DTO) tree.TreeNode {
	domainName := conn.Domain.DomainName
	if conn.Subdomain != nil && *conn.Subdomain != "" {
		domainName = fmt.Sprintf("%s.%s", *conn.Subdomain, domainName)
	}
	domainName += stateSuffix(conn.AggregatedVerificationState)

	children := []tree.TreeNode{
		{Value: fmt.Sprintf("ID: %s", conn.ID)},
	}

	if len(conn.Gateways) > 0 {
		gatewayNodes := make([]tree.TreeNode, 0, len(conn.Gateways))
		for _, gw := range conn.Gateways {
			gatewayNodes = append(gatewayNodes, buildGatewayNode(gw, conn.DomainVerificationState))
		}
		children = append(children, tree.TreeNode{
			Value:    "Gateways",
			Children: gatewayNodes,
		})
	}

	return tree.TreeNode{
		Value:    domainName,
		Children: children,
	}
}

func buildGatewayNode(gw api.GatewayConnectionV5DTO, domainVerificationState *[]api.DomainVerificationStateV5DTO) tree.TreeNode {
	var children []tree.TreeNode

	if gw.VerificationState != nil {
		gwLabel := "Gateway verification"
		gwLabel += stateSuffix(gw.VerificationState.AggregatedVerificationState)
		children = append(children, tree.TreeNode{
			Value:    gwLabel,
			Children: buildRecordStatusNodes(gw.VerificationState),
		})
	}

	domainVS := findDomainVerificationStateForGateway(domainVerificationState, gw.GatewayID)
	if domainVS != nil {
		domainLabel := "Domain verification"
		domainLabel += stateSuffix(domainVS.AggregatedVerificationState)
		children = append(children, tree.TreeNode{
			Value:    domainLabel,
			Children: buildRecordStatusNodes(domainVS),
		})
	}

	return tree.TreeNode{
		Value:    gw.Name,
		Children: children,
	}
}

func findDomainVerificationStateForGateway(vs *[]api.DomainVerificationStateV5DTO, gatewayID string) *api.VerificationStateV5DTO {
	if vs == nil {
		return nil
	}
	for _, item := range *vs {
		if item.GatewayID == gatewayID {
			return item.VerificationState
		}
	}
	return nil
}

func buildRecordStatusNodes(vs *api.VerificationStateV5DTO) []tree.TreeNode {
	if vs == nil {
		return []tree.TreeNode{{Value: "N/A"}}
	}
	return []tree.TreeNode{
		{Value: fmt.Sprintf("Console: %s", vs.Console.Status)},
		{Value: fmt.Sprintf("S3: %s", vs.S3.Status)},
		{Value: fmt.Sprintf("Wildcard S3: %s", vs.WildcardS3.Status)},
	}
}
