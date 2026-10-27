package revoke

import (
	"fmt"

	"github.com/cubbit/composer-cli/constants"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/src/service/api_key/shared"
	"github.com/spf13/cobra"
)

func RevokeAPIKey(deps shared.Dependencies, cmd *cobra.Command, profile configuration_models.ProfileV2) error {
	apiKeyID, err := cmd.Flags().GetString("id")
	if err != nil {
		return fmt.Errorf("%s id: %w", constants.ErrorRetrievingField, err)
	}

	operatorID, err := shared.ResolveAPIKeyTargetOperatorID(deps, cmd, profile)
	if err != nil {
		return err
	}

	if err := deps.UserAPI.DeleteIAMAPIKey(profile.Endpoints, profile.APIKey, profile.OrganizationID, operatorID, apiKeyID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDeletingIAMAPIKeyRequest, err)
	}

	quiet, err := cmd.Flags().GetBool("quiet")
	if err != nil {
		return fmt.Errorf("%s quiet: %w", constants.ErrorRetrievingField, err)
	}
	if !quiet {
		cmd.Printf("IAM API key %s revoked successfully\n", apiKeyID)
	}
	return nil
}
