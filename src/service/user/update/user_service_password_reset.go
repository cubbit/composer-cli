package update

import (
	"fmt"
	"os"

	"github.com/cubbit/composer-cli/constants"
	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/src/service/user/shared"
	"github.com/cubbit/composer-cli/src/tui"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

const iamPasswordChangeTokenType = "iam_password_change"

func ResetUserPassword(
	deps Dependencies,
	cmd *cobra.Command,
	profile configuration_models.ProfileV2,
	args []string,
) error {
	userID, err := shared.ResolveUserID(deps.UserAPI, cmd, args, profile)
	if err != nil {
		return err
	}

	newPassword, err := cmd.Flags().GetString("new-password")
	if err != nil {
		return fmt.Errorf("%s new-password: %w", constants.ErrorRetrievingField, err)
	}

	operatorPassword, err := cmd.Flags().GetString("password")
	if err != nil {
		return fmt.Errorf("%s password: %w", constants.ErrorRetrievingField, err)
	}
	if operatorPassword == "" {
		operatorPassword = os.Getenv("PASSWORD")
	}
	if operatorPassword == "" {
		if _, err = tui.TextInputs(
			"Please provide your login details",
			false,
			tui.Input{Placeholder: "password", IsPassword: true, Value: &operatorPassword},
		); err != nil {
			return fmt.Errorf("failed to read operator password: %w", err)
		}
	}

	tfaCode, err := cmd.Flags().GetString("tfa-code")
	if err != nil {
		return fmt.Errorf("%s tfa-code: %w", constants.ErrorRetrievingField, err)
	}

	currentUser, err := deps.UserAPI.GetIAMUserSelf(profile.Endpoints, "", profile.APIKey)
	if err != nil {
		return fmt.Errorf("failed to retrieve current IAM user: %w", err)
	}
	if currentUser.ID == "" {
		return fmt.Errorf("current IAM user does not expose an id")
	}
	if currentUser.Username == "" {
		return fmt.Errorf("current IAM user does not expose a username")
	}
	if currentUser.OrganizationName == nil || *currentUser.OrganizationName == "" {
		return fmt.Errorf("current IAM user does not expose an organization name")
	}

	targetUser, err := deps.UserAPI.GetIAMUserByID(profile.Endpoints, profile.APIKey, profile.OrganizationID, userID)
	if err != nil {
		return fmt.Errorf("failed to retrieve target IAM user: %w", err)
	}
	if targetUser.Username == "" {
		return fmt.Errorf("target IAM user does not expose a username")
	}

	tokens, err := deps.AuthAPI.SignIn(profile.Endpoints, currentUser.Username, *currentUser.OrganizationName, operatorPassword, tfaCode)
	if err != nil {
		return fmt.Errorf("failed to authenticate current IAM user: %w", err)
	}

	iamPasswordChangeToken, err := deps.AuthAPI.ForgeToken(
		profile.Endpoints,
		userID,
		currentUser.Username,
		*currentUser.OrganizationName,
		operatorPassword,
		tfaCode,
		iamPasswordChangeTokenType,
		tokens.AccessToken,
		tokens.RefreshToken,
	)
	if err != nil {
		return fmt.Errorf("failed to forge reset password token: %w", err)
	}

	saltsResponse, err := deps.UserAPI.BulkGenerateSalts(
		profile.Endpoints,
		profile.APIKey,
		profile.OrganizationID,
		&api.BulkGenerateSaltsRequestBody{
			Operators: []api.BulkGenerateSaltRequestBodyItem{
				{Username: targetUser.Username},
			},
		},
	)
	if err != nil {
		return fmt.Errorf("failed to generate password salt: %w", err)
	}
	if len(saltsResponse.Data) == 0 {
		return fmt.Errorf("salt not returned for user %q", targetUser.Username)
	}

	authenticationPublicKey, err := utils.AuthenticationPublicKeyFromPassword(newPassword, saltsResponse.Data[0].Salt)
	if err != nil {
		return fmt.Errorf("failed to generate authentication public key for user %q: %w", targetUser.Username, err)
	}

	if err := deps.UserAPI.ResetIAMUserPassword(
		profile.Endpoints,
		profile.APIKey,
		profile.OrganizationID,
		userID,
		&api.ResetIAMUserPasswordRequestBody{
			IAMPasswordChangeToken:  iamPasswordChangeToken,
			AuthenticationPublicKey: authenticationPublicKey,
		},
	); err != nil {
		return fmt.Errorf("failed to reset IAM user password: %w", err)
	}

	quiet, err := cmd.Flags().GetBool("quiet")
	if err != nil {
		return fmt.Errorf("%s quiet: %w", constants.ErrorRetrievingField, err)
	}
	if !quiet {
		cmd.Printf("IAM user %s password reset successfully\n", userID)
	}

	return nil
}
