package service

import (
	"fmt"
	"strings"

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

func PrintIAMUserDetails(cmd *cobra.Command, user api.IAMUser) error {
	return printer.PrintText(cmd, buildIAMUserDetailsOutput(user))
}

func buildIAMUserDetailsOutput(user api.IAMUser) string {
	userName := strings.TrimSpace(user.FirstName + " " + user.LastName)
	if userName == "" {
		userName = user.Username
	}

	twoFactor := "Disabled"
	if user.TwoFactorEnabled {
		twoFactor = "Enabled"
	}

	orgID := "N/A"
	if user.OrganizationID != nil {
		orgID = *user.OrganizationID
	}

	orgName := "N/A"
	if user.OrganizationName != nil {
		orgName = *user.OrganizationName
	}

	deletedAt := "N/A"
	if user.DeletedAt != nil {
		deletedAt = utils.FormatTime(*user.DeletedAt)
	}

	var policyNames []string
	for _, p := range user.Policies {
		policyNames = append(policyNames, p.Name)
	}
	policiesStr := "N/A"
	if len(policyNames) > 0 {
		policiesStr = strings.Join(policyNames, ", ")
	}

	root := "No"
	if user.IsRoot {
		root = "Yes"
	}

	lines := []string{
		fmt.Sprintf("User: %s", userName),
		fmt.Sprintf("Root: %s", root),
		fmt.Sprintf("Username: %s", user.Username),
		fmt.Sprintf("Status: %s", formatIAMUserStatus(user.Status)),
		"",
		"Organization:",
		fmt.Sprintf("  Name: %s", orgName),
		fmt.Sprintf("  ID: %s", orgID),
		"",
		"Policies:",
		fmt.Sprintf("  %s", policiesStr),
	}

	if len(user.Emails) > 0 {
		lines = append(lines, "", "Emails:")
		for _, e := range user.Emails {
			label := e.Email
			if e.Default {
				label += " (default)"
			}
			lines = append(lines, fmt.Sprintf("  %s", label))
		}
	}

	lines = append(lines,
		"",
		"Metadata:",
		fmt.Sprintf("  ID: %s", user.ID),
		fmt.Sprintf("  Two-Factor: %s", twoFactor),
		fmt.Sprintf("  Created At: %s", utils.FormatTime(user.CreatedAt)),
		fmt.Sprintf("  Deleted At: %s", deletedAt),
	)

	return strings.Join(lines, "\n") + "\n"
}

func formatIAMUserStatus(status string) string {
	if status == "" {
		return "N/A"
	}

	if len(status) == 1 {
		return "● " + strings.ToUpper(status)
	}

	return "● " + strings.ToUpper(status[:1]) + status[1:]
}
