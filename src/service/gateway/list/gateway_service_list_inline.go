package list

import (
	"fmt"
	"strings"

	"github.com/cubbit/composer-cli/constants"
	api "github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration"
	"github.com/cubbit/composer-cli/src/service/gateway/shared"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func ListInline(deps Dependencies, cmd *cobra.Command, resolvedProfile configuration.ResolvedProfile, urls configuration.URLs) error {
	filters, err := cmd.Flags().GetStringArray("query")
	if err != nil {
		return fmt.Errorf("%s filter: %w", constants.ErrorRetrievingField, err)
	}

	page, err := cmd.Flags().GetInt("page")
	if err != nil {
		return fmt.Errorf("%s page: %w", constants.ErrorRetrievingField, err)
	}

	items, err := cmd.Flags().GetInt("items")
	if err != nil {
		return fmt.Errorf("%s items: %w", constants.ErrorRetrievingField, err)
	}

	filter := strings.Join(filters, ",")

	var gateways []api.GatewayV5ListItemResponse

	if cmd.Flags().Changed("page") || cmd.Flags().Changed("items") {
		response, err := deps.GatewayAPI.ListGatewaysV5(
			urls,
			resolvedProfile.APIKey,
			resolvedProfile.OrganizationID,
			api.WithPage(page),
			api.WithItems(items),
			api.WithFilter(filter),
		)
		if err != nil {
			return fmt.Errorf("failed to list gateways: %w", err)
		}

		gateways = response.Data
	} else {
		gateways, err = fetchAllGateways(deps, urls, resolvedProfile.APIKey, resolvedProfile.OrganizationID, filter)
		if err != nil {
			return fmt.Errorf("failed to list gateways: %w", err)
		}
	}

	output, err := shared.ResolveCommandOutput(cmd, resolvedProfile.Output)
	if err != nil {
		return err
	}

	if output == string(configuration.OutputHuman) {
		return PrintGatewayList(cmd, gateways)
	}

	utils.PrintFormattedData(cmd.OutOrStdout(), gateways, output)
	return nil
}

func fetchAllGateways(deps Dependencies, urls configuration.URLs, apiKey string, organizationID string, filter string) ([]api.GatewayV5ListItemResponse, error) {
	page := 1
	itemsPerPage := 100
	var all []api.GatewayV5ListItemResponse

	for {
		response, err := deps.GatewayAPI.ListGatewaysV5(
			urls,
			apiKey,
			organizationID,
			api.WithPage(page),
			api.WithItems(itemsPerPage),
			api.WithFilter(filter),
		)
		if err != nil {
			return nil, err
		}

		all = append(all, response.Data...)

		if response.NextPage == nil {
			break
		}

		page = *response.NextPage
	}

	return all, nil
}
