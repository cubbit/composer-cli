package list

import (
	"fmt"
	"strconv"

	"github.com/cubbit/composer-cli/constants"
	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/src/service/api_key/shared"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func ListAPIKeys(deps shared.Dependencies, cmd *cobra.Command, handler configuration_handler.ConfigurationHandlerInterface, profile configuration_models.ProfileV2) error {
	operatorID, err := shared.ResolveAPIKeyTargetOperatorID(deps, cmd, profile)
	if err != nil {
		return err
	}

	page, err := cmd.Flags().GetInt("page")
	if err != nil {
		return fmt.Errorf("%s page: %w", constants.ErrorRetrievingField, err)
	}
	items, err := cmd.Flags().GetInt("items")
	if err != nil {
		return fmt.Errorf("%s items: %w", constants.ErrorRetrievingField, err)
	}
	sortKey, err := cmd.Flags().GetString("sort-key")
	if err != nil {
		return fmt.Errorf("%s sort-key: %w", constants.ErrorRetrievingField, err)
	}
	sortOrder, err := cmd.Flags().GetString("sort-order")
	if err != nil {
		return fmt.Errorf("%s sort-order: %w", constants.ErrorRetrievingField, err)
	}
	if items <= 0 {
		items = 100
	}

	var allAPIKeys []api.OperatorAPIKey
	if cmd.Flags().Changed("page") {
		response, err := deps.UserAPI.ListIAMAPIKeys(profile.Endpoints, profile.APIKey, profile.OrganizationID, operatorID, page, items, sortKey, sortOrder)
		if err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingIAMAPIKeysRequest, err)
		}
		allAPIKeys = response.Data
	} else {
		page = 1
		for {
			response, err := deps.UserAPI.ListIAMAPIKeys(profile.Endpoints, profile.APIKey, profile.OrganizationID, operatorID, page, items, sortKey, sortOrder)
			if err != nil {
				return fmt.Errorf("%s: %w", constants.ErrorListingIAMAPIKeysRequest, err)
			}
			allAPIKeys = append(allAPIKeys, response.Data...)
			if response.NextPage == nil {
				break
			}
			page = *response.NextPage
		}
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
		utils.PrintFormattedData(cmd.OutOrStdout(), allAPIKeys, output)
		return nil
	}
	if quiet {
		for _, item := range allAPIKeys {
			utils.PrintQuiet(cmd.OutOrStdout(), item.ID, item.Name, strconv.FormatBool(item.Enabled))
		}
		return nil
	}

	return shared.PrintAPIKeyList(cmd, handler, allAPIKeys)
}
