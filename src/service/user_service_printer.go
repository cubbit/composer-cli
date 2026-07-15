package service

import (
	"fmt"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/utils/printer"
	"github.com/cubbit/composer-cli/utils/printer/table"
	"github.com/cubbit/composer-cli/utils/printer/utils"
	"github.com/spf13/cobra"
)

func PrintIAMUserList(cmd *cobra.Command, users []api.IAMUserListItem) error {
	if len(users) == 0 {
		return printer.PrintText(cmd, "No IAM users found.\n")
	}

	noHeaders, err := cmd.Flags().GetBool("no-headers")
	if err != nil {
		return fmt.Errorf("failed to read no-headers flag: %w", err)
	}

	tableColumns := []table.Column[api.IAMUserListItem]{
		{Title: "Username"},
		{Title: "ID"},
		{Title: "Email"},
		{Title: "First Name"},
		{Title: "Last Name"},
		{Title: "Status"},
		{Title: "Enabled"},
		{Title: "Created At"},
	}

	rowMapper := func(item api.IAMUserListItem) []string {
		firstName := ""
		if item.FirstName != nil {
			firstName = *item.FirstName
		}
		lastName := ""
		if item.LastName != nil {
			lastName = *item.LastName
		}

		return []string{
			item.Username,
			item.ID,
			defaultIAMUserEmail(item.Emails),
			firstName,
			lastName,
			item.Status,
			fmt.Sprintf("%t", item.Enabled),
			utils.FormatTime(item.CreatedAt),
		}
	}

	return printer.CreateTable(
		cmd,
		users,
		table.WithColumns(tableColumns),
		table.WithRowMapper(rowMapper),
		table.WithShowHeader[api.IAMUserListItem](!noHeaders),
		table.WithSuffix[api.IAMUserListItem]("\n"),
	)
}

func defaultIAMUserEmail(emails []api.IAMUserEmail) string {
	for _, email := range emails {
		if email.Default {
			return email.Email
		}
	}
	if len(emails) > 0 {
		return emails[0].Email
	}
	return ""
}
