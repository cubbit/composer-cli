package apikey

import (
	"fmt"
	"strconv"

	"github.com/cubbit/composer-cli/constants"
	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func EditAPIKey(deps Dependencies, cmd *cobra.Command, profile configuration_models.ProfileV2) error {
	apiKeyID, err := cmd.Flags().GetString("id")
	if err != nil {
		return fmt.Errorf("%s id: %w", constants.ErrorRetrievingField, err)
	}

	request, err := buildEditAPIKeyRequest(cmd)
	if err != nil {
		return err
	}

	operator, _, err := resolveCurrentOperator(deps, profile)
	if err != nil {
		return err
	}

	apiKey, err := deps.UserAPI.UpdateIAMAPIKey(
		profile.Endpoints,
		profile.APIKey,
		profile.OrganizationID,
		operator.ID,
		apiKeyID,
		request,
	)
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorEditingIAMAPIKeyRequest, err)
	}

	output, err := resolveOutput(cmd, profile.Output)
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
		utils.PrintQuiet(cmd.OutOrStdout(), apiKey.ID, apiKey.Name, strconv.FormatBool(apiKey.Enabled))
		return nil
	}

	return PrintAPIKeyDetails(cmd, *apiKey)
}

func buildEditAPIKeyRequest(cmd *cobra.Command) (*api.UpdateIAMAPIKeyRequestBody, error) {
	request := &api.UpdateIAMAPIKeyRequestBody{}
	updates := 0

	if cmd.Flags().Changed("name") {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return nil, fmt.Errorf("%s name: %w", constants.ErrorRetrievingField, err)
		}
		request.Name = &name
		updates++
	}

	if cmd.Flags().Changed("expires-at") {
		expiresAt, err := parseOptionalExpiresAt(cmd)
		if err != nil {
			return nil, err
		}
		if expiresAt == nil {
			return nil, fmt.Errorf("invalid value for --expires-at: expected RFC3339 timestamp")
		}
		request.ExpiresAt = expiresAt
		updates++
	}

	if cmd.Flags().Changed("enabled") {
		enabled, err := cmd.Flags().GetBool("enabled")
		if err != nil {
			return nil, fmt.Errorf("%s enabled: %w", constants.ErrorRetrievingField, err)
		}
		request.Enabled = &enabled
		updates++
	}

	if updates == 0 {
		return nil, fmt.Errorf("specify at least one field to update")
	}

	return request, nil
}
