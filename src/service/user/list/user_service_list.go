package list

import (
	"fmt"
	"strconv"

	"github.com/cubbit/composer-cli/constants"
	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/src/service/user/shared"
	"github.com/cubbit/composer-cli/utils"
	"github.com/cubbit/composer-cli/utils/printer"
	"github.com/cubbit/composer-cli/utils/printer/table"
	printerutils "github.com/cubbit/composer-cli/utils/printer/utils"
	"github.com/spf13/cobra"
)

type Dependencies struct {
	UserAPI api.UserAPIInterface
}

func ListUsers(deps Dependencies, cmd *cobra.Command, handler configuration_handler.ConfigurationHandlerInterface, profile configuration_models.ProfileV2) error {
	enabledStr, err := cmd.Flags().GetString("enabled")
	if err != nil {
		return fmt.Errorf("%s enabled: %w", constants.ErrorRetrievingField, err)
	}
	search, err := cmd.Flags().GetString("search")
	if err != nil {
		return fmt.Errorf("%s search: %w", constants.ErrorRetrievingField, err)
	}
	page, err := cmd.Flags().GetInt("page")
	if err != nil {
		return fmt.Errorf("%s page: %w", constants.ErrorRetrievingField, err)
	}
	items, err := cmd.Flags().GetInt("items")
	if err != nil {
		return fmt.Errorf("%s items: %w", constants.ErrorRetrievingField, err)
	}
	sortKey, err := cmd.Flags().GetString("sort-key")
	if err != nil {
		return fmt.Errorf("%s sort-key: %w", constants.ErrorRetrievingField, err)
	}
	sortOrder, err := cmd.Flags().GetString("sort-order")
	if err != nil {
		return fmt.Errorf("%s sort-order: %w", constants.ErrorRetrievingField, err)
	}

	var enabled *bool
	if enabledStr != "" {
		parsed, err := strconv.ParseBool(enabledStr)
		if err != nil {
			return fmt.Errorf("invalid value for --enabled: expected true or false, got %q", enabledStr)
		}
		enabled = &parsed
	}

	if items <= 0 {
		items = 100
	}

	var allUsers []api.IAMUserListItem
	if cmd.Flags().Changed("page") {
		response, err := deps.UserAPI.ListIAMUsers(
			profile.Endpoints, profile.APIKey, profile.OrganizationID,
			enabled, search, page, items, sortKey, sortOrder,
		)
		if err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingIAMUsersRequest, err)
		}
		allUsers = response.Data
	} else {
		page = 1
		for {
			response, err := deps.UserAPI.ListIAMUsers(
				profile.Endpoints, profile.APIKey, profile.OrganizationID,
				enabled, search, page, items, sortKey, sortOrder,
			)
			if err != nil {
				return fmt.Errorf("%s: %w", constants.ErrorListingIAMUsersRequest, err)
			}
			allUsers = append(allUsers, response.Data...)
			if response.NextPage == nil {
				break
			}
			page = *response.NextPage
		}
	}

	quiet, err := cmd.Flags().GetBool("quiet")
	if err != nil {
		return fmt.Errorf("%s quiet: %w", constants.ErrorRetrievingField, err)
	}

	if quiet {
		for _, item := range allUsers {
			firstName := ""
			if item.FirstName != nil {
				firstName = *item.FirstName
			}
			lastName := ""
			if item.LastName != nil {
				lastName = *item.LastName
			}
			utils.PrintQuiet(
				cmd.OutOrStdout(),
				item.Username,
				item.ID,
				shared.DefaultIAMUserEmail(item.Emails),
				firstName,
				lastName,
				item.Status,
				strconv.FormatBool(item.Enabled),
			)
		}
		return nil
	}

	return printer.ComposeStructured(cmd, handler, allUsers,
		func() error { return PrintIAMUserList(cmd, handler, allUsers) },
	)
}

func PrintIAMUserList(cmd *cobra.Command, handler configuration_handler.ConfigurationHandlerInterface, users []api.IAMUserListItem) error {
	if len(users) == 0 {
		return printer.PrintText(cmd, handler, "No IAM users found.\n")
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
			shared.DefaultIAMUserEmail(item.Emails),
			firstName,
			lastName,
			item.Status,
			fmt.Sprintf("%t", item.Enabled),
			printerutils.FormatTime(item.CreatedAt),
		}
	}

	return printer.PrintTable(
		cmd,
		handler,
		users,
		table.WithColumns(tableColumns),
		table.WithRowMapper(rowMapper),
		table.WithShowHeader[api.IAMUserListItem](!noHeaders),
		table.WithSuffix[api.IAMUserListItem]("\n"),
	)
}
