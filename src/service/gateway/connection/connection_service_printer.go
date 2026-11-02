package connection

import (
	"fmt"
	"strconv"

	api "github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/utils/printer"
	"github.com/cubbit/composer-cli/utils/printer/table"
	"github.com/spf13/cobra"
)

func PrintConnectionList(cmd *cobra.Command, handler configuration_handler.ConfigurationHandlerInterface, connections []api.GatewayConnectionV5) error {
	if len(connections) == 0 {
		return printer.PrintText(cmd, handler, "No connections found.\n")
	}

	noHeaders, err := cmd.Flags().GetBool("no-headers")
	if err != nil {
		return fmt.Errorf("failed to read no-headers flag: %w", err)
	}

	tableColumns := []table.Column[api.GatewayConnectionV5]{
		{Title: "ID"},
		{Title: "Swarm"},
		{Title: "Redundancy Class"},
		{Title: "Default"},
		{Title: "Status"},
		{Title: "Impacted Rules"},
	}

	rowMapper := func(connection api.GatewayConnectionV5) []string {
		return []string{
			connection.ID,
			connection.Swarm.Name,
			connection.RedundancyClass.Name,
			formatConnectionDefault(connection.IsDefault),
			formatConnectionStatus(connection.Status),
			strconv.Itoa(len(connection.ImpactedRules)),
		}
	}

	return printer.PrintTable(
		cmd,
		handler,
		connections,
		table.WithColumns(tableColumns),
		table.WithRowMapper(rowMapper),
		table.WithShowHeader[api.GatewayConnectionV5](!noHeaders),
		table.WithSuffix[api.GatewayConnectionV5]("\n"),
	)
}

func PrintConnectionsCreated(cmd *cobra.Command, handler configuration_handler.ConfigurationHandlerInterface, connections []api.GatewayConnectionV5) error {
	if err := printer.PrintText(cmd, handler, buildConnectionsCreatedOutput(len(connections))); err != nil {
		return err
	}

	return PrintConnectionList(cmd, handler, connections)
}

func buildConnectionsCreatedOutput(count int) string {
	if count == 1 {
		return "1 connection created successfully.\n\n"
	}

	return fmt.Sprintf("%d connections created successfully.\n\n", count)
}

func buildConnectionDeletedOutput(connectionID string) string {
	return fmt.Sprintf("Connection %s deleted successfully.\n", connectionID)
}

func formatConnectionDefault(isDefault bool) string {
	if isDefault {
		return "Yes"
	}

	return "-"
}

func formatConnectionStatus(status api.GatewayConnectionStatusV5) string {
	switch status {
	case api.GatewayConnectionStatusV5Active:
		return "Active"
	case api.GatewayConnectionStatusV5NotActive:
		return "Not active"
	default:
		return string(status)
	}
}
