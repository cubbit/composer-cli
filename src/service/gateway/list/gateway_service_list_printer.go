package list

import (
	"fmt"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/utils/printer"
	"github.com/cubbit/composer-cli/utils/printer/table"
	"github.com/spf13/cobra"
)

func PrintGatewayList(cmd *cobra.Command, gateways []api.GatewayV5ListItemResponse) error {
	if len(gateways) == 0 {
		return printer.PrintText(cmd, "No gateways found.\n")
	}

	noHeaders, err := cmd.Flags().GetBool("no-headers")
	if err != nil {
		return fmt.Errorf("failed to read no-headers flag: %w", err)
	}

	tableColumns := []table.Column[api.GatewayV5ListItemResponse]{
		{Title: "Name"},
		{Title: "Slug"},
		{Title: "Type"},
		{Title: "Swarms"},
		{Title: "Tenants"},
		{Title: "Status"},
	}

	rowMapper := func(v api.GatewayV5ListItemResponse) []string {
		return []string{
			v.Name,
			v.Slug,
			v.Type,
			fmt.Sprintf("%d", v.NumberOfSwarms),
			fmt.Sprintf("%d", v.NumberOfTenants),
			formatGatewayStatus(v.Status),
		}
	}

	return printer.CreateTable(
		cmd,
		gateways,
		table.WithColumns(tableColumns),
		table.WithRowMapper(rowMapper),
		table.WithShowHeader[api.GatewayV5ListItemResponse](!noHeaders),
		table.WithSuffix[api.GatewayV5ListItemResponse]("\n"),
	)
}

func formatGatewayStatus(status api.GatewayV5Status) string {
	if status == api.GatewayV5StatusReady {
		return "ready"
	}
	return "not-ready"
}
