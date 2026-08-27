package describe

import (
	"fmt"

	"github.com/cubbit/composer-cli/constants"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/src/service/api_key/shared"
	"github.com/cubbit/composer-cli/utils/printer"
	"github.com/spf13/cobra"
)

func DescribeAPIKey(deps shared.Dependencies, cmd *cobra.Command, handler configuration_handler.ConfigurationHandlerInterface, profile configuration_models.ProfileV2) error {
	apiKeyID, err := cmd.Flags().GetString("id")
	if err != nil {
		return fmt.Errorf("%s id: %w", constants.ErrorRetrievingField, err)
	}

	operatorID, err := shared.ResolveAPIKeyTargetOperatorID(deps, cmd, profile)
	if err != nil {
		return err
	}

	apiKey, err := deps.UserAPI.GetIAMAPIKeyByID(profile.Endpoints, profile.APIKey, profile.OrganizationID, operatorID, apiKeyID)
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDescribingIAMAPIKeyRequest, err)
	}

	return printer.ComposeStructured(cmd, handler, apiKey,
		func() error { return shared.PrintAPIKeyDetails(cmd, handler, *apiKey) },
	)
}
