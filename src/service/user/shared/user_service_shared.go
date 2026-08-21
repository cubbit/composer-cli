package shared

import (
	"fmt"
	"strings"

	"github.com/cubbit/composer-cli/constants"
	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/spf13/cobra"
)

func ResolveCommandOutput(cmd *cobra.Command, defaultOutput configuration_models.OutputFormat) (string, error) {
	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return "", fmt.Errorf("%s output: %w", constants.ErrorRetrievingField, err)
	}

	if defaultOutput != "" &&
		!cmd.Flags().Changed("output") &&
		!cmd.Flags().Changed("quiet") {
		output = string(defaultOutput)
	}

	return output, nil
}

func ResolveUserID(
	userAPI api.UserAPIInterface,
	cmd *cobra.Command,
	args []string,
	profile configuration_models.ProfileV2,
) (string, error) {
	identifiers := 0

	userIDFlag, err := cmd.Flags().GetString("user-id")
	if err != nil {
		return "", fmt.Errorf("%s user-id: %w", constants.ErrorRetrievingField, err)
	}
	if userIDFlag != "" {
		identifiers++
	}

	usernameFlag, err := cmd.Flags().GetString("username")
	if err != nil {
		return "", fmt.Errorf("%s username: %w", constants.ErrorRetrievingField, err)
	}
	if usernameFlag != "" {
		identifiers++
	}

	userIDPositional := ""
	if len(args) > 0 {
		userIDPositional = strings.TrimSpace(args[0])
		if userIDPositional != "" {
			identifiers++
		}
	}

	if identifiers != 1 {
		return "", fmt.Errorf("specify exactly one of USER_ID, --user-id or --username")
	}

	if userIDPositional != "" {
		return userIDPositional, nil
	}

	if userIDFlag != "" {
		return userIDFlag, nil
	}

	return ResolveUserIDByUsername(userAPI, profile, usernameFlag)
}

func ResolveUserIDByUsername(
	userAPI api.UserAPIInterface,
	profile configuration_models.ProfileV2,
	username string,
) (string, error) {
	page := 1
	const itemsPerPage = 1000

	for {
		response, err := userAPI.ListIAMUsers(
			profile.Endpoints,
			profile.APIKey,
			profile.OrganizationID,
			nil,
			username,
			page,
			itemsPerPage,
			"",
			"",
		)
		if err != nil {
			return "", fmt.Errorf("failed to resolve username '%s': %w", username, err)
		}

		for _, item := range response.Data {
			if item.Username == username {
				return item.ID, nil
			}
		}

		if response.NextPage == nil {
			break
		}

		page = *response.NextPage
	}

	return "", fmt.Errorf("IAM user with username '%s' not found", username)
}

func DefaultIAMUserEmail(emails []api.IAMUserEmail) string {
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
