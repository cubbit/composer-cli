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
	userID, err := shared.ResolveUserID(deps.UserAPI, cmd, args, profile)
	if err != nil {
		return err
	}

	user, err := deps.UserAPI.GetIAMUserByID(profile.Endpoints, profile.APIKey, profile.OrganizationID, userID)
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDescribingIAMUserRequest, err)
	}

	user.OrganizationID = &profile.OrganizationID

	organizationName, err := getProfileOrganizationName(deps, profile)
	if err == nil {
		user.OrganizationName = &organizationName
	}

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

func getProfileOrganizationName(deps Dependencies, profile configuration_models.ProfileV2) (string, error) {
	user, err := deps.UserAPI.GetIAMUserSelf(profile.Endpoints, "", profile.APIKey)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve current IAM user: %w", err)
	}
	if user.OrganizationName == nil || *user.OrganizationName == "" {
		return "", fmt.Errorf("current IAM user does not expose an organization name")
	}

	return *user.OrganizationName, nil
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
