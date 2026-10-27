package service

import (
	"fmt"
	"strings"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/utils/printer"
	"github.com/cubbit/composer-cli/utils/printer/table"
	"github.com/cubbit/composer-cli/utils/printer/utils"
	"github.com/spf13/cobra"
)

func PrintDomainList(cmd *cobra.Command, handler configuration_handler.ConfigurationHandlerInterface, domains []api.DomainDTO) error {
	if len(domains) == 0 {
		return printer.PrintText(cmd, handler, "No domains found.\n")
	}

	noHeaders, err := cmd.Flags().GetBool("no-headers")
	if err != nil {
		return fmt.Errorf("failed to read no-headers flag: %w", err)
	}

	tableColumns := []table.Column[api.DomainDTO]{
		{Title: "Domain Name"},
		{Title: "ID"},
		{Title: "Verified"},
		{Title: "Created On"},
		{Title: "Challenge"},
	}

	rowMapper := func(v api.DomainDTO) []string {
		verified := "No"
		if v.VerifiedAt != nil {
			verified = "Yes"
		}

		return []string{
			v.DomainName,
			v.ID,
			verified,
			utils.FormatTime(v.CreatedAt),
			v.Challenge,
		}
	}

	return printer.PrintTable(
		cmd,
		handler,
		domains,
		table.WithColumns(tableColumns),
		table.WithRowMapper(rowMapper),
		table.WithShowHeader[api.DomainDTO](!noHeaders),
		table.WithSuffix[api.DomainDTO]("\n"),
	)
}

func PrintDomainDetails(cmd *cobra.Command, handler configuration_handler.ConfigurationHandlerInterface, domain api.DomainDTO) error {
	return printer.PrintText(cmd, handler, buildDomainDetailsOutput(domain))
}

func buildDomainDetailsOutput(domain api.DomainDTO) string {
	verified := "No"
	verifiedAt := "N/A"
	if domain.VerifiedAt != nil {
		verified = "Yes"
		verifiedAt = utils.FormatTime(*domain.VerifiedAt)
	}

	deleted := "No"
	if domain.DeletedAt != nil {
		deleted = "Yes"
	}

	shared := "No"
	if domain.IsShared {
		shared = "Yes"
	}

	lines := []string{
		fmt.Sprintf("Domain: %s", domain.DomainName),
		fmt.Sprintf("ID: %s", domain.ID),
		fmt.Sprintf("Organization: %s", domain.OrganizationID),
		fmt.Sprintf("Challenge: %s", domain.Challenge),
		fmt.Sprintf("Verified: %s", verified),
		fmt.Sprintf("Verified At: %s", verifiedAt),
		fmt.Sprintf("Deleted: %s", deleted),
		fmt.Sprintf("Shared: %s", shared),
		fmt.Sprintf("Created At: %s", utils.FormatTime(domain.CreatedAt)),
	}

	return strings.Join(lines, "\n") + "\n"
}
