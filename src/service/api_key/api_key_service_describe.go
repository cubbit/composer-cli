package apikey

import (
	"fmt"

	"github.com/cubbit/composer-cli/constants"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func DescribeAPIKey(deps Dependencies, cmd *cobra.Command, profile configuration_models.ProfileV2) error {
	apiKeyID, err := cmd.Flags().GetString("id")
	if err != nil {
		return fmt.Errorf("%s id: %w", constants.ErrorRetrievingField, err)
	}

	operator, _, err := resolveCurrentOperator(deps, profile)
	if err != nil {
		return err
	}

	apiKey, err := deps.UserAPI.GetIAMAPIKeyByID(profile.Endpoints, profile.APIKey, profile.OrganizationID, operator.ID, apiKeyID)
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDescribingIAMAPIKeyRequest, err)
	}

	output, err := resolveOutput(cmd, profile.Output)
	if err != nil {
		return err
	}
	if output == string(configuration_models.OutputHuman) {
		return PrintAPIKeyDetails(cmd, *apiKey)
	}

	utils.PrintFormattedData(cmd.OutOrStdout(), apiKey, output)
	return nil
}
