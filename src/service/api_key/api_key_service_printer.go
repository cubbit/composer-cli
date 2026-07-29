package apikey

import (
	"fmt"
	"strings"
	"time"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/utils/printer"
	"github.com/cubbit/composer-cli/utils/printer/table"
	printerutils "github.com/cubbit/composer-cli/utils/printer/utils"
	"github.com/spf13/cobra"
)

func PrintAPIKeyCreated(cmd *cobra.Command, apiKey api.OperatorAPIKey) error {
	return printer.PrintText(cmd, buildAPIKeyDetailsOutput(apiKey, true))
}

func PrintAPIKeyDetails(cmd *cobra.Command, apiKey api.OperatorAPIKey) error {
	return printer.PrintText(cmd, buildAPIKeyDetailsOutput(apiKey, false))
}

func PrintAPIKeyList(cmd *cobra.Command, apiKeys []api.OperatorAPIKey) error {
	if len(apiKeys) == 0 {
		return printer.PrintText(cmd, "No IAM API keys found.\n")
	}

	noHeaders, err := cmd.Flags().GetBool("no-headers")
	if err != nil {
		return fmt.Errorf("failed to read no-headers flag: %w", err)
	}

	tableColumns := []table.Column[api.OperatorAPIKey]{
		{Title: "Name"},
		{Title: "ID"},
		{Title: "Enabled"},
		{Title: "Created At"},
		{Title: "Expires At"},
	}

	rowMapper := func(item api.OperatorAPIKey) []string {
		return []string{
			item.Name,
			item.ID,
			fmt.Sprintf("%t", item.Enabled),
			printerutils.FormatTime(item.CreatedAt),
			formatOptionalTime(item.ExpiresAt),
		}
	}

	return printer.CreateTable(
		cmd,
		apiKeys,
		table.WithColumns(tableColumns),
		table.WithRowMapper(rowMapper),
		table.WithShowHeader[api.OperatorAPIKey](!noHeaders),
		table.WithSuffix[api.OperatorAPIKey]("\n"),
	)
}

func buildAPIKeyDetailsOutput(apiKey api.OperatorAPIKey, includeSecret bool) string {
	lines := []string{
		fmt.Sprintf("API Key: %s", apiKey.Name),
		fmt.Sprintf("Enabled: %t", apiKey.Enabled),
		"",
		"Metadata:",
		fmt.Sprintf("  ID: %s", apiKey.ID),
		fmt.Sprintf("  Operator ID: %s", apiKey.OperatorID),
		fmt.Sprintf("  Created At: %s", printerutils.FormatTime(apiKey.CreatedAt)),
		fmt.Sprintf("  Expires At: %s", formatOptionalTime(apiKey.ExpiresAt)),
		fmt.Sprintf("  Deleted At: %s", formatOptionalTime(apiKey.DeletedAt)),
		fmt.Sprintf("  Banned At: %s", formatOptionalTime(apiKey.BannedAt)),
	}

	if includeSecret {
		lines = append(lines, "", "Secret:", fmt.Sprintf("  Key: %s", apiKey.Key))
	}

	return strings.Join(lines, "\n") + "\n"
}

func formatOptionalTime(value *time.Time) string {
	if value == nil {
		return "N/A"
	}
	return printerutils.FormatTime(*value)
}
