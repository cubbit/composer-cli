package verifyconnection

import (
	"fmt"

	"github.com/cubbit/composer-cli/constants"
	api "github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/src/service/gateway/shared"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

type Dependencies struct {
	ConnectionAPI api.ConnectionAPIInterface
	TenantAPI     api.TenantAPIInterface
}

func VerifyConnection(deps Dependencies, cmd *cobra.Command, handler configuration_handler.ConfigurationHandlerInterface, profile configuration_models.ProfileV2, args []string) error {
	tenantID, err := resolveTenantID(deps, cmd, profile)
	if err != nil {
		return err
	}

	connectionID, err := cmd.Flags().GetString("connection-id")
	if err != nil {
		return fmt.Errorf("%s connection-id: %w", constants.ErrorRetrievingField, err)
	}
	if connectionID == "" {
		return fmt.Errorf("--connection-id is required")
	}

	verifyResponse, err := deps.ConnectionAPI.VerifyConnectionV5(
		profile.Endpoints,
		profile.APIKey,
		profile.OrganizationID,
		tenantID,
		connectionID,
	)
	if err != nil {
		return fmt.Errorf("failed to verify connection: %w", err)
	}

	output, err := shared.ResolveCommandOutput(cmd, profile.Output)
	if err != nil {
		return err
	}

	if output == string(configuration_models.OutputHuman) {
		return PrintConnectionVerification(cmd, handler, verifyResponse)
	}

	utils.PrintFormattedData(cmd.OutOrStdout(), verifyResponse, output)
	return nil
}

func resolveTenantID(deps Dependencies, cmd *cobra.Command, profile configuration_models.ProfileV2) (string, error) {
	tenantIDFlag, err := cmd.Flags().GetString("tenant-id")
	if err != nil {
		return "", fmt.Errorf("%s tenant-id: %w", constants.ErrorRetrievingField, err)
	}

	tenantNameFlag, err := cmd.Flags().GetString("tenant-name")
	if err != nil {
		return "", fmt.Errorf("%s tenant-name: %w", constants.ErrorRetrievingField, err)
	}

	if tenantIDFlag != "" && tenantNameFlag != "" {
		return "", fmt.Errorf("specify either --tenant-id or --tenant-name, not both")
	}

	if tenantIDFlag != "" {
		return tenantIDFlag, nil
	}

	if tenantNameFlag != "" {
		return resolveTenantIDByName(deps, profile, tenantNameFlag)
	}

	return "", fmt.Errorf("specify either --tenant-id or --tenant-name")
}

func resolveTenantIDByName(deps Dependencies, profile configuration_models.ProfileV2, tenantName string) (string, error) {
	const itemsPerPage = 1000
	page := 1

	for {
		response, err := deps.TenantAPI.ListTenantsV5(
			profile.Endpoints,
			profile.APIKey,
			profile.OrganizationID,
			api.WithTenantPage(page),
			api.WithTenantItems(itemsPerPage),
		)
		if err != nil {
			return "", fmt.Errorf("failed to resolve tenant name '%s': %w", tenantName, err)
		}

		for _, t := range response.Data {
			if t.Name == tenantName {
				return t.ID, nil
			}
		}

		if response.NextPage == nil {
			break
		}
		page = *response.NextPage
	}

	return "", fmt.Errorf("tenant with name '%s' not found", tenantName)
}
