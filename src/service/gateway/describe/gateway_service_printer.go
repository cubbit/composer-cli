package describe

import (
	"fmt"
	"strings"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/utils/printer"
	"github.com/spf13/cobra"
)

func PrintGatewayDetails(cmd *cobra.Command, gateway api.GatewayV5GetResponse) error {
	return printer.PrintText(cmd, buildGatewayDetailsOutput(gateway))
}

func buildGatewayDetailsOutput(gateway api.GatewayV5GetResponse) string {
	lines := []string{
		fmt.Sprintf("Gateway: %s", gateway.Name),
		fmt.Sprintf("Status: %s", formatGatewayStatus(string(gateway.Status))),
		"",
		"Metadata:",
		fmt.Sprintf("  ID: %s", gateway.ID),
		fmt.Sprintf("  Slug: %s", gateway.Slug),
		fmt.Sprintf("  Type: %s", gateway.Type),
		"",
		"Redundancy Classes:",
	}

	if len(gateway.RedundancyClasses) == 0 {
		lines = append(lines, "  N/A")
	} else {
		for _, redundancyClass := range gateway.RedundancyClasses {
			lines = append(lines, fmt.Sprintf("  - %s (%s)", redundancyClass.Name, redundancyClass.ID))
		}
	}

	return strings.Join(lines, "\n")
}

func formatGatewayStatus(status string) string {
	if status == "" {
		return "N/A"
	}

	if len(status) == 1 {
		return "● " + strings.ToUpper(status)
	}

	return "● " + strings.ToUpper(status[:1]) + status[1:]
}
