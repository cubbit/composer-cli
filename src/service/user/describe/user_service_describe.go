package describe

import (
	"fmt"
	"strings"

	"github.com/cubbit/composer-cli/constants"
	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/src/service/user/shared"
	"github.com/cubbit/composer-cli/utils"
	"github.com/cubbit/composer-cli/utils/printer"
	printerutils "github.com/cubbit/composer-cli/utils/printer/utils"
	"github.com/spf13/cobra"
)

type Dependencies struct {
	UserAPI api.UserAPIInterface
}

func DescribeUser(deps Dependencies, cmd *cobra.Command, profile configuration_models.ProfileV2, args []string) error {
	user, err := resolveUser(deps, cmd, profile, args)
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDescribingIAMUserRequest, err)
	}

	user.OrganizationID = &profile.OrganizationID

	output, err := shared.ResolveCommandOutput(cmd, profile.Output)
	if err != nil {
		return err
	}

	if output == string(configuration_models.OutputHuman) {
		return PrintIAMUserDetails(cmd, *user)
	}

	utils.PrintFormattedData(cmd.OutOrStdout(), user, output)
	return nil
}

func resolveUser(deps Dependencies, cmd *cobra.Command, profile configuration_models.ProfileV2, args []string) (*api.IAMUser, error) {
	self, err := shouldDescribeCurrentUser(cmd, args)
	if err != nil {
		return nil, err
	}
	if self {
		return deps.UserAPI.GetIAMUserSelfV3(profile.Endpoints, profile.APIKey, profile.OrganizationID)
	}

	userID, err := shared.ResolveUserID(deps.UserAPI, cmd, args, profile)
	if err != nil {
		return nil, err
	}

	return deps.UserAPI.GetIAMUserByID(profile.Endpoints, profile.APIKey, profile.OrganizationID, userID)
}

func shouldDescribeCurrentUser(cmd *cobra.Command, args []string) (bool, error) {
	selfFlag, err := cmd.Flags().GetBool("self")
	if err != nil {
		return false, fmt.Errorf("%s self: %w", constants.ErrorRetrievingField, err)
	}

	if !selfFlag {
		return false, nil
	}

	userIDFlag, err := cmd.Flags().GetString("user-id")
	if err != nil {
		return false, fmt.Errorf("%s user-id: %w", constants.ErrorRetrievingField, err)
	}

	usernameFlag, err := cmd.Flags().GetString("username")
	if err != nil {
		return false, fmt.Errorf("%s username: %w", constants.ErrorRetrievingField, err)
	}

	if userIDFlag != "" || usernameFlag != "" || len(args) > 0 {
		return false, fmt.Errorf("specify exactly one of USER_ID, --user-id, --username or --self")
	}

	return true, nil
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
		deletedAt = printerutils.FormatTime(*user.DeletedAt)
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
		fmt.Sprintf("  Created At: %s", printerutils.FormatTime(user.CreatedAt)),
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
