package update

import (
	"fmt"

	"github.com/cubbit/composer-cli/constants"
	"github.com/cubbit/composer-cli/src/api"
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
	AuthAPI api.AuthAPIInterface
}

func EditUser(
	deps Dependencies,
	cmd *cobra.Command,
	profile configuration_models.ProfileV2,
	args []string,
) error {
	userID, err := shared.ResolveUserID(deps.UserAPI, cmd, args, profile)
	if err != nil {
		return err
	}

	request, err := buildEditIAMUserRequest(cmd)
	if err != nil {
		return err
	}

	return updateUser(deps, cmd, profile, userID, request)
}

func EnableUser(
	deps Dependencies,
	cmd *cobra.Command,
	profile configuration_models.ProfileV2,
	args []string,
) error {
	userID, err := shared.ResolveUserID(deps.UserAPI, cmd, args, profile)
	if err != nil {
		return err
	}

	enabled := true
	return updateUser(deps, cmd, profile, userID, &api.UpdateIAMUserRequestBody{Enabled: &enabled})
}

func DisableUser(
	deps Dependencies,
	cmd *cobra.Command,
	profile configuration_models.ProfileV2,
	args []string,
) error {
	userID, err := shared.ResolveUserID(deps.UserAPI, cmd, args, profile)
	if err != nil {
		return err
	}

	enabled := false
	return updateUser(deps, cmd, profile, userID, &api.UpdateIAMUserRequestBody{Enabled: &enabled})
}

func updateUser(
	deps Dependencies,
	cmd *cobra.Command,
	profile configuration_models.ProfileV2,
	userID string,
	request *api.UpdateIAMUserRequestBody,
) error {
	updatedUser, err := deps.UserAPI.UpdateIAMUser(
		profile.Endpoints,
		profile.APIKey,
		profile.OrganizationID,
		userID,
		request,
	)
	if err != nil {
		return err
	}

	output, err := shared.ResolveCommandOutput(cmd, profile.Output)
	if err != nil {
		return err
	}
	quiet, err := cmd.Flags().GetBool("quiet")
	if err != nil {
		return fmt.Errorf("%s quiet: %w", constants.ErrorRetrievingField, err)
	}

	if output != string(configuration_models.OutputHuman) && !quiet {
		utils.PrintFormattedData(cmd.OutOrStdout(), updatedUser, output)
		return nil
	}

	if quiet {
		utils.PrintQuiet(cmd.OutOrStdout(), updatedUser.Username, updatedUser.ID, shared.DefaultIAMUserEmail(updatedUser.Emails), updatedUser.Status, fmt.Sprintf("%t", updatedUser.Enabled))
		return nil
	}

	return PrintIAMUserUpdate(cmd, *updatedUser)
}

func buildEditIAMUserRequest(cmd *cobra.Command) (*api.UpdateIAMUserRequestBody, error) {
	request := &api.UpdateIAMUserRequestBody{}
	updates := 0

	if cmd.Flags().Changed("first-name") {
		firstName, err := cmd.Flags().GetString("first-name")
		if err != nil {
			return nil, fmt.Errorf("%s first-name: %w", constants.ErrorRetrievingField, err)
		}
		request.FirstName = &firstName
		updates++
	}

	if cmd.Flags().Changed("last-name") {
		lastName, err := cmd.Flags().GetString("last-name")
		if err != nil {
			return nil, fmt.Errorf("%s last-name: %w", constants.ErrorRetrievingField, err)
		}
		request.LastName = &lastName
		updates++
	}

	if cmd.Flags().Changed("email") {
		email, err := cmd.Flags().GetString("email")
		if err != nil {
			return nil, fmt.Errorf("%s email: %w", constants.ErrorRetrievingField, err)
		}
		request.Email = &email
		updates++
	}

	if updates == 0 {
		return nil, fmt.Errorf("specify at least one field to update")
	}

	return request, nil
}

func PrintIAMUserUpdate(cmd *cobra.Command, user api.IAMUser) error {
	noHeaders, err := cmd.Flags().GetBool("no-headers")
	if err != nil {
		return fmt.Errorf("%s no-headers: %w", constants.ErrorRetrievingField, err)
	}

	tableColumns := []table.Column[api.IAMUser]{
		{Title: "Username"},
		{Title: "ID"},
		{Title: "Email"},
		{Title: "First Name"},
		{Title: "Last Name"},
		{Title: "Status"},
		{Title: "Enabled"},
		{Title: "Created At"},
	}

	rowMapper := func(user api.IAMUser) []string {
		return []string{
			user.Username,
			user.ID,
			shared.DefaultIAMUserEmail(user.Emails),
			user.FirstName,
			user.LastName,
			user.Status,
			fmt.Sprintf("%t", user.Enabled),
			printerutils.FormatTime(user.CreatedAt),
		}
	}

	return printer.CreateTable(
		cmd,
		[]api.IAMUser{user},
		table.WithColumns(tableColumns),
		table.WithRowMapper(rowMapper),
		table.WithShowHeader[api.IAMUser](!noHeaders),
		table.WithSuffix[api.IAMUser]("\n"),
	)
}
