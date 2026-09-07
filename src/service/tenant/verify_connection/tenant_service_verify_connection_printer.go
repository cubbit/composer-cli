package verifyconnection

import (
	"fmt"
	"strings"

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

func PrintConnectionVerification(cmd *cobra.Command, handler configuration_handler.ConfigurationHandlerInterface, conn *api.ConnectionV5DTO) error {
	if conn == nil {
		return printer.PrintText(cmd, handler, "No verification data returned.\n")
	}

	domainName := conn.Domain.DomainName
	if conn.Subdomain != nil && *conn.Subdomain != "" {
		domainName = fmt.Sprintf("%s.%s", *conn.Subdomain, domainName)
	}
	domainName += stateSuffix(conn.AggregatedVerificationState)

	children := make([]tree.TreeNode, 0, 1+len(conn.Gateways))

	if conn.DomainVerificationState != nil {
		for _, dvs := range *conn.DomainVerificationState {
			if dvs.VerificationState == nil {
				continue
			}
			vs := dvs.VerificationState
			sectionLabel := "Domain Records"
			sectionLabel += stateSuffix(conn.DomainAggregatedVerificationState)

			children = append(children, tree.TreeNode{
				Value: sectionLabel,
				Children: []tree.TreeNode{
					buildRecordNode("Console", vs.Console),
					buildRecordNode("S3", vs.S3),
					buildRecordNode("Wildcard S3", vs.WildcardS3),
				},
			})
			break
		}
	}

	for _, gw := range conn.Gateways {
		if gw.VerificationState == nil {
			continue
		}
		vs := gw.VerificationState
		gwLabel := fmt.Sprintf("Gateway: %s (%s)", gw.Name, gw.GatewayID)
		gwLabel += stateSuffix(vs.AggregatedVerificationState)

		gwChildren := []tree.TreeNode{
			buildRecordNode("Console", vs.Console),
			buildRecordNode("S3", vs.S3),
			buildRecordNode("Wildcard S3", vs.WildcardS3),
		}

		children = append(children, tree.TreeNode{
			Value:    gwLabel,
			Children: gwChildren,
		})
	}

	if len(children) == 0 {
		return printer.PrintText(cmd, handler, "No verification data available.\n")
	}

	nodes := []tree.TreeNode{
		{
			Value:    domainName,
			Children: children,
		},
	}

	return printer.PrintTree(cmd, handler, nodes, func(n []tree.TreeNode) []tree.TreeNode { return n })
}

func buildRecordNode(label string, record api.RecordInfoV5DTO) tree.TreeNode {
	recordLabel := fmt.Sprintf("%s (%s)", label, record.FQDN)

	children := []tree.TreeNode{{Value: fmt.Sprintf("Status: %s", record.Status)}}

	if record.Details != nil && *record.Details != "" {
		children = append(children, tree.TreeNode{Value: fmt.Sprintf("Details: %s", *record.Details)})
	}

	if record.Certificates != nil {
		children = append(children, buildCertificateNode(record.Certificates))
	}

	return tree.TreeNode{
		Value:    recordLabel,
		Children: children,
	}
}

func buildCertificateNode(cert *api.CertificatesV5DTO) tree.TreeNode {
	certChildren := []tree.TreeNode{
		{Value: fmt.Sprintf("Issuer: %s", cert.Issuer)},
		{Value: fmt.Sprintf("Subject: %s", cert.Subject)},
		{Value: fmt.Sprintf("Valid from: %s", cert.ValidFrom)},
		{Value: fmt.Sprintf("Valid to: %s (%d days remaining)", cert.ValidTo, cert.DaysRemaining)},
		{Value: fmt.Sprintf("Fingerprint SHA256: %s", cert.FingerprintSHA256)},
		{Value: fmt.Sprintf("SAN: %s", strings.Join(cert.San, ", "))},
	}

	return tree.TreeNode{
		Value:    "Certificate",
		Children: certChildren,
	}
}
