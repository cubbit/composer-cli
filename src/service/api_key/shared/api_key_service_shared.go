package shared

import (
	"fmt"
	"time"

	"github.com/cubbit/composer-cli/constants"
	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	userShared "github.com/cubbit/composer-cli/src/service/user/shared"
	"github.com/spf13/cobra"
)

type Dependencies struct {
	UserAPI api.UserAPIInterface
}

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

func ParseOptionalExpiresAt(cmd *cobra.Command) (*time.Time, error) {
	expiresAt, err := cmd.Flags().GetString("expires-at")
	if err != nil {
		return nil, fmt.Errorf("%s expires-at: %w", constants.ErrorRetrievingField, err)
	}
	if expiresAt == "" {
		return nil, nil
	}

	parsed, err := time.Parse(time.RFC3339, expiresAt)
	if err != nil {
		return nil, fmt.Errorf("invalid value for --expires-at: expected RFC3339 timestamp, got %q", expiresAt)
	}

	return &parsed, nil
}

func ResolveAPIKeyTargetOperatorID(deps Dependencies, cmd *cobra.Command, profile configuration_models.ProfileV2) (string, error) {
	userID, err := getOptionalStringFlag(cmd, "user-id")
	if err != nil {
		return "", fmt.Errorf("%s user-id: %w", constants.ErrorRetrievingField, err)
	}
	username, err := getOptionalStringFlag(cmd, "username")
	if err != nil {
		return "", fmt.Errorf("%s username: %w", constants.ErrorRetrievingField, err)
	}

	if userID != "" && username != "" {
		return "", fmt.Errorf("specify at most one of --user-id or --username")
	}
	if userID != "" {
		return ResolveUserIDByID(deps.UserAPI, profile, userID)
	}
	if username != "" {
		return userShared.ResolveUserIDByUsername(deps.UserAPI, profile, username)
	}

	operator, _, err := ResolveCurrentOperator(deps, profile)
	if err != nil {
		return "", err
	}

	return operator.ID, nil
}

func ResolveUserIDByID(
	userAPI api.UserAPIInterface,
	profile configuration_models.ProfileV2,
	userID string,
) (string, error) {
	user, err := userAPI.GetIAMUserByID(profile.Endpoints, profile.APIKey, profile.OrganizationID, userID)
	if err != nil {
		return "", fmt.Errorf("failed to resolve user-id '%s': %w", userID, err)
	}
	if user == nil || user.ID != userID {
		return "", fmt.Errorf("IAM user with ID '%s' not found", userID)
	}
	if user.OrganizationID != nil && *user.OrganizationID != profile.OrganizationID {
		return "", fmt.Errorf("IAM user with ID '%s' not found in organization '%s'", userID, profile.OrganizationID)
	}

	return user.ID, nil
}

func getOptionalStringFlag(cmd *cobra.Command, name string) (string, error) {
	if cmd.Flags().Lookup(name) == nil {
		return "", nil
	}

	return cmd.Flags().GetString(name)
}

func ResolveCurrentOperator(deps Dependencies, profile configuration_models.ProfileV2) (*api.IAMUser, string, error) {
	operator, err := deps.UserAPI.GetIAMUserSelf(profile.Endpoints, "", profile.APIKey)
	if err != nil {
		return nil, "", fmt.Errorf("failed to retrieve current IAM user: %w", err)
	}
	if operator.ID == "" {
		return nil, "", fmt.Errorf("current IAM user does not expose an ID")
	}
	if operator.OrganizationName == nil || *operator.OrganizationName == "" {
		return nil, "", fmt.Errorf("current IAM user does not expose an organization name")
	}

	return operator, *operator.OrganizationName, nil
}
