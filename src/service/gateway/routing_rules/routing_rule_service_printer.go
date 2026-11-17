package routing_rules

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	api "github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/utils/printer"
	"github.com/cubbit/composer-cli/utils/printer/table"
	"github.com/spf13/cobra"
)

// PrintRoutingRuleList prints a compact metadata table of the gateway routing
// rules. Use get (or structured output) for the full document.
func PrintRoutingRuleList(
	cmd *cobra.Command,
	handler configuration_handler.ConfigurationHandlerInterface,
	rules []api.GatewayRoutingRuleV5,
) error {
	if len(rules) == 0 {
		return printer.PrintText(cmd, handler, "No routing rules found.\n")
	}

	noHeaders, err := cmd.Flags().GetBool("no-headers")
	if err != nil {
		return fmt.Errorf("failed to read no-headers flag: %w", err)
	}

	tableColumns := []table.Column[api.GatewayRoutingRuleV5]{
		{Title: "ID"},
		{Title: "Name"},
		{Title: "Enabled"},
		{Title: "Input Storage Class"},
		{Title: "Description"},
	}

	rowMapper := func(rule api.GatewayRoutingRuleV5) []string {
		return []string{
			rule.ID,
			rule.Name,
			formatRoutingRuleBool(rule.Enabled),
			rule.InputStorageClass,
			formatRoutingRuleDescription(rule.Description),
		}
	}

	return printer.PrintTable(
		cmd,
		handler,
		rules,
		table.WithColumns(tableColumns),
		table.WithRowMapper(rowMapper),
		table.WithShowHeader[api.GatewayRoutingRuleV5](!noHeaders),
		table.WithSuffix[api.GatewayRoutingRuleV5]("\n"),
	)
}

// PrintRoutingRuleDocument prints the raw server document as indented JSON.
func PrintRoutingRuleDocument(
	cmd *cobra.Command,
	handler configuration_handler.ConfigurationHandlerInterface,
	raw json.RawMessage,
) error {
	var formatted bytes.Buffer
	if err := json.Indent(&formatted, raw, "", "  "); err != nil {
		return fmt.Errorf("failed to encode routing rule JSON: %w", err)
	}

	formatted.WriteByte('\n')

	return printer.PrintText(cmd, handler, formatted.String())
}

func PrintRoutingRuleCreated(
	cmd *cobra.Command,
	handler configuration_handler.ConfigurationHandlerInterface,
	ruleID string,
) error {
	return printer.PrintText(
		cmd,
		handler,
		fmt.Sprintf("Routing rule %s created successfully.\n", ruleID),
	)
}

func PrintRoutingRuleUpdated(
	cmd *cobra.Command,
	handler configuration_handler.ConfigurationHandlerInterface,
	ruleID string,
) error {
	return printer.PrintText(
		cmd,
		handler,
		fmt.Sprintf("Routing rule %s updated successfully.\n", ruleID),
	)
}

func buildRoutingRuleDeletedOutput(ruleID string) string {
	return fmt.Sprintf("Routing rule %s deleted successfully.\n", ruleID)
}

func formatRoutingRuleDescription(description string) string {
	description = strings.TrimSpace(description)
	if description == "" {
		return "-"
	}

	return description
}

func formatRoutingRuleBool(value bool) string {
	if value {
		return "yes"
	}

	return "no"
}
