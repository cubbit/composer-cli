package list

import (
	"fmt"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/utils/printer"
	"github.com/cubbit/composer-cli/utils/printer/table"
	"github.com/spf13/cobra"
)

const descriptionMaxLen = 256

func PrintTenantList(cmd *cobra.Command, handler configuration_handler.ConfigurationHandlerInterface, tenants []api.TenantV5DTO) error {
	if len(tenants) == 0 {
		return printer.PrintText(cmd, handler, "No tenants found.\n")
	}

	noHeaders, err := cmd.Flags().GetBool("no-headers")
	if err != nil {
		return fmt.Errorf("failed to read no-headers flag: %w", err)
	}

	tableColumns := []table.Column[api.TenantV5DTO]{
		{Title: "ID"},
		{Title: "Name"},
		{Title: "Slug"},
		{Title: "Description"},
		{Title: "Created At"},
	}

	rowMapper := func(v api.TenantV5DTO) []string {
		return []string{
			v.ID,
			v.Name,
			v.Slug,
			formatDescription(v.Description),
			v.CreatedAt.Format("2006-01-02 15:04:05"),
		}
	}

	return printer.PrintTable(
		cmd,
		handler,
		tenants,
		table.WithColumns(tableColumns),
		table.WithRowMapper(rowMapper),
		table.WithShowHeader[api.TenantV5DTO](!noHeaders),
		table.WithSuffix[api.TenantV5DTO]("\n"),
	)
}

func formatDescription(desc *string) string {
	if desc == nil {
		return ""
	}
	if len(*desc) > descriptionMaxLen {
		return (*desc)[:descriptionMaxLen] + "..."
	}
	return *desc
}
