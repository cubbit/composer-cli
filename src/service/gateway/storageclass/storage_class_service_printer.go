package storageclass

import (
	"fmt"
	"strconv"
	"strings"

	api "github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/utils/printer"
	"github.com/cubbit/composer-cli/utils/printer/table"
	"github.com/spf13/cobra"
)

func PrintStorageClassList(cmd *cobra.Command, handler configuration_handler.ConfigurationHandlerInterface, storageClasses []api.GatewayStorageClassV5) error {
	if len(storageClasses) == 0 {
		return printer.PrintText(cmd, handler, "No storage classes found.\n")
	}

	noHeaders, err := cmd.Flags().GetBool("no-headers")
	if err != nil {
		return fmt.Errorf("failed to read no-headers flag: %w", err)
	}

	tableColumns := []table.Column[api.GatewayStorageClassV5]{
		{Title: "Name"},
		{Title: "ID"},
		{Title: "Status"},
		{Title: "Default"},
		{Title: "Referenced Rules"},
	}

	rowMapper := func(storageClass api.GatewayStorageClassV5) []string {
		return []string{
			storageClass.Name,
			storageClass.ID,
			formatStorageClassStatus(storageClass.Status),
			formatStorageClassDefault(storageClass.Pool),
			strconv.Itoa(len(storageClass.ReferencedRules)),
		}
	}

	return printer.PrintTable(
		cmd,
		handler,
		storageClasses,
		table.WithColumns(tableColumns),
		table.WithRowMapper(rowMapper),
		table.WithShowHeader[api.GatewayStorageClassV5](!noHeaders),
		table.WithSuffix[api.GatewayStorageClassV5]("\n"),
	)
}

func PrintStorageClassCreated(cmd *cobra.Command, handler configuration_handler.ConfigurationHandlerInterface, storageClass api.GatewayStorageClassV5) error {
	if err := printer.PrintText(cmd, handler, fmt.Sprintf("Storage class %s created successfully.\n\n", storageClass.Name)); err != nil {
		return err
	}

	return printer.PrintText(cmd, handler, buildStorageClassDetailsOutput(storageClass))
}

func buildStorageClassDetailsOutput(storageClass api.GatewayStorageClassV5) string {
	lines := []string{
		"Storage Class:",
		fmt.Sprintf("  Name: %s", storageClass.Name),
		fmt.Sprintf("  ID: %s", storageClass.ID),
		fmt.Sprintf("  Status: %s", formatStorageClassStatus(storageClass.Status)),
		fmt.Sprintf("  Default: %s", formatStorageClassDefault(storageClass.Pool)),
		fmt.Sprintf("  Referenced Rules: %d", len(storageClass.ReferencedRules)),
	}

	return strings.Join(lines, "\n")
}

func buildStorageClassDeletedOutput(storageClassID string) string {
	return fmt.Sprintf("Storage class %s deleted successfully.\n", storageClassID)
}

func formatStorageClassStatus(status api.GatewayStorageClassStatusV5) string {
	switch status {
	case api.GatewayStorageClassStatusV5Active:
		return "Active"
	case api.GatewayStorageClassStatusV5FallbackOnly:
		return "Fallback only"
	default:
		return string(status)
	}
}

func formatStorageClassDefault(isDefault bool) string {
	if isDefault {
		return "yes"
	}

	return "-"
}
