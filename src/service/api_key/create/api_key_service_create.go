package create

import (
	"fmt"
	"strconv"

	"github.com/cubbit/composer-cli/constants"
	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/src/service/api_key/shared"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func CreateAPIKey(deps shared.Dependencies, cmd *cobra.Command, profile configuration_models.ProfileV2) error {
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		return fmt.Errorf("%s name: %w", constants.ErrorRetrievingField, err)
	}
	expiresAt, err := shared.ParseOptionalExpiresAt(cmd)
	if err != nil {
		return err
	}

	operatorID, err := shared.ResolveAPIKeyTargetOperatorID(deps, cmd, profile)
	if err != nil {
		return err
	}

	apiKey, err := deps.UserAPI.CreateIAMAPIKey(
		profile.Endpoints,
		profile.APIKey,
		profile.OrganizationID,
		operatorID,
		&api.CreateIAMAPIKeyRequestBody{
			Name:      name,
			ExpiresAt: expiresAt,
		},
	)
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCreatingIAMAPIKeyRequest, err)
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
		utils.PrintFormattedData(cmd.OutOrStdout(), apiKey, output)
		return nil
	}
	if quiet {
		utils.PrintQuiet(cmd.OutOrStdout(), apiKey.Key, apiKey.ID, apiKey.Name, strconv.FormatBool(apiKey.Enabled))
		return nil
	}

	return shared.PrintAPIKeyCreated(cmd, *apiKey)
}
