package create

import (
	"fmt"
	"strings"

	"github.com/cubbit/composer-cli/constants"
	api "github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/utils"
	"github.com/cubbit/composer-cli/utils/printer"
	"github.com/spf13/cobra"
)

func createInline(deps Dependencies, cmd *cobra.Command, handler configuration_handler.ConfigurationHandlerInterface, endpoints configuration_models.EndpointsV2, apiKey string, organizationID string) error {
	createRequest, err := collectTenantCreateFlags(cmd)
	if err != nil {
		return err
	}

	response, err := deps.TenantAPI.CreateTenantV5(endpoints, apiKey, organizationID, createRequest)
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCreatingTenantRequest, err)
	}

	process, err := deps.ProcessAPI.GetProcess(
		endpoints,
		apiKey,
		organizationID,
		response.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to retrieve tenant creation process: %w", err)
	}

	tenantCreationProcess, ok := process.CastToTenantCreationProcess()
	if !ok {
		return fmt.Errorf("unexpected process type: expected tenant creation process, got %s", process.Type)
	}

	return printer.PrintText(cmd, handler, fmt.Sprintf("Tenant creation started — Tenant ID: %s\n", tenantCreationProcess.Data.TenantID))
}

func collectTenantCreateFlags(cmd *cobra.Command) (*api.CreateTenantV5Request, error) {
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		return nil, fmt.Errorf("%s name: %w", constants.ErrorRetrievingField, err)
	}

	slug, err := cmd.Flags().GetString("slug")
	if err != nil {
		return nil, fmt.Errorf("%s slug: %w", constants.ErrorRetrievingField, err)
	}

	description, err := utils.GetOptionalStringFlag(cmd, "description")
	if err != nil {
		return nil, fmt.Errorf("%s description: %w", constants.ErrorRetrievingField, err)
	}

	connectionFlags, err := cmd.Flags().GetStringArray("connection")
	if err != nil {
		return nil, fmt.Errorf("%s connection: %w", constants.ErrorRetrievingField, err)
	}

	connections, err := parseConnectionFlags(connectionFlags)
	if err != nil {
		return nil, err
	}

	return &api.CreateTenantV5Request{
		Name:        name,
		Slug:        slug,
		Description: description,
		Connections: connections,
	}, nil
}

func parseConnectionFlags(flags []string) ([]api.CreateTenantV5Connection, error) {
	if len(flags) == 0 {
		return nil, fmt.Errorf("at least one --connection flag is required")
	}

	connections := make([]api.CreateTenantV5Connection, 0, len(flags))

	for i, flag := range flags {
		parts := strings.SplitN(flag, ":", 3)

		if len(parts) < 2 {
			return nil, fmt.Errorf("invalid --connection format at position %d: expected 'domainID:gatewayID[,gatewayID...][:subdomain]', got '%s'", i, flag)
		}

		domainID := strings.TrimSpace(parts[0])
		if domainID == "" {
			return nil, fmt.Errorf("domain ID cannot be empty at position %d", i)
		}

		gatewayIDsStr := strings.TrimSpace(parts[1])
		if gatewayIDsStr == "" {
			return nil, fmt.Errorf("gateway IDs cannot be empty at position %d", i)
		}

		gatewayParts := strings.Split(gatewayIDsStr, ",")
		gatewayIDs := make([]string, 0, len(gatewayParts))
		for _, g := range gatewayParts {
			gwID := strings.TrimSpace(g)
			if gwID != "" {
				gatewayIDs = append(gatewayIDs, gwID)
			}
		}

		if len(gatewayIDs) == 0 {
			return nil, fmt.Errorf("at least one gateway ID is required at position %d", i)
		}

		var subdomain *string
		if len(parts) >= 3 && strings.TrimSpace(parts[2]) != "" {
			s := strings.TrimSpace(parts[2])
			subdomain = &s
		}

		connections = append(connections, api.CreateTenantV5Connection{
			DomainID:   domainID,
			GatewayIDs: gatewayIDs,
			Subdomain:  subdomain,
		})
	}

	return connections, nil
}
