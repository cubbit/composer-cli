package service

import (
	"fmt"

	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/utils/printer"
	"github.com/cubbit/composer-cli/utils/printer/table"
	"github.com/spf13/cobra"
)

type profileRow struct {
	Name           string
	Output         string
	OrganizationID string
	UpdatedAt      string
	Active         string
}

func PrintProfiles(cmd *cobra.Command, handler configuration_handler.ConfigurationHandlerInterface, profiles map[string]configuration_models.ProfileV2, activeProfile string) error {
	if len(profiles) == 0 {
		return printer.PrintText(
			cmd,
			handler,
			"No profiles found\n",
		)
	}

	noHeaders, err := cmd.Flags().GetBool("no-headers")
	if err != nil {
		return fmt.Errorf("failed to read no-headers flag: %w", err)
	}

	tableColumns := []table.Column[profileRow]{
		{Title: "Name"},
		{Title: "Output"},
		{Title: "Organization ID"},
		{Title: "Updated At"},
		{Title: "Active"},
	}

	rowMapper := func(row profileRow) []string {
		return []string{
			row.Name,
			row.Output,
			row.OrganizationID,
			row.UpdatedAt,
			row.Active,
		}
	}

	rows := make([]profileRow, 0, len(profiles))
	for name, profile := range profiles {
		isActive := "No"
		if name == activeProfile {
			isActive = "Yes"
		}

		rows = append(rows, profileRow{
			Name:           name,
			Output:         string(profile.Output),
			OrganizationID: profile.OrganizationID,
			UpdatedAt:      profile.UpdatedAt.Format("2006-01-02 15:04:05"),
			Active:         isActive,
		})
	}

	return printer.PrintTable(cmd, handler, rows,
		table.WithColumns(tableColumns),
		table.WithRowMapper(rowMapper),
		table.WithShowHeader[profileRow](!noHeaders),
		table.WithSuffix[profileRow]("\n"),
	)
}
