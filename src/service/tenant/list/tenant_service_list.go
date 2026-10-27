package list

import (
	"fmt"
	"strings"

	"github.com/cubbit/composer-cli/constants"
	api "github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/spf13/cobra"
)

type Dependencies struct {
	TenantAPI api.TenantAPIInterface
}

func List(deps Dependencies, cmd *cobra.Command, handler configuration_handler.ConfigurationHandlerInterface, profile configuration_models.ProfileV2) error {
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

	var tenants []api.TenantV5DTO

	if cmd.Flags().Changed("page") || cmd.Flags().Changed("items") {
		response, err := deps.TenantAPI.ListTenantsV5(
			profile.Endpoints,
			profile.APIKey,
			profile.OrganizationID,
			api.WithTenantPage(page),
			api.WithTenantItems(items),
			api.WithTenantFilter(filter),
		)
		if err != nil {
			return fmt.Errorf("failed to list tenants: %w", err)
		}

		tenants = response.Data
	} else {
		tenants, err = fetchAllTenants(deps, profile.Endpoints, profile.APIKey, profile.OrganizationID, filter)
		if err != nil {
			return fmt.Errorf("failed to list tenants: %w", err)
		}
	}

	return PrintTenantList(cmd, handler, tenants)
}

func fetchAllTenants(deps Dependencies, endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, filter string) ([]api.TenantV5DTO, error) {
	page := 1
	itemsPerPage := 100
	var all []api.TenantV5DTO

	for {
		response, err := deps.TenantAPI.ListTenantsV5(
			endpoints,
			apiKey,
			organizationID,
			api.WithTenantPage(page),
			api.WithTenantItems(itemsPerPage),
			api.WithTenantFilter(filter),
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
