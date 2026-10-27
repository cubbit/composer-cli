package create

import (
	"encoding/json"
	"fmt"

	api "github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/spf13/cobra"
)

type Dependencies struct {
	GatewayAPI         api.GatewayAPIInterface
	SwarmAPI           api.SwarmAPIInterface
	RedundancyClassAPI api.RedundancyClassAPIInterface
	ProcessAPI         api.ProcessAPIInterface
	LocationAPI        api.LocationAPIInterface
}

func Create(deps Dependencies, cmd *cobra.Command, handler configuration_handler.ConfigurationHandlerInterface, profile configuration_models.ProfileV2, interactiveMode bool) error {
	if err := checkRunningGatewayProcess(deps, profile.Endpoints, profile.APIKey, profile.OrganizationID); err != nil {
		return err
	}

	if interactiveMode {
		return createInteractive(deps, cmd, profile.Endpoints, profile.APIKey, profile.OrganizationID)
	}

	return createInline(deps, cmd, handler, profile.Endpoints, profile.APIKey, profile.OrganizationID)
}

func checkRunningGatewayProcess(deps Dependencies, endpoints configuration_models.EndpointsV2, apiKey string, organizationID string) error {
	processes, err := deps.ProcessAPI.ListProcesses(
		endpoints,
		apiKey,
		organizationID,
		api.WithProcessType(api.ProcessTypeGatewayCreation),
		api.WithProcessStatus(api.ProcessStatusRunning),
	)
	if err != nil {
		return fmt.Errorf("failed to check for existing gateway creation process: %w", err)
	}

	for _, p := range processes {
		if p.Type != api.ProcessTypeGatewayCreation {
			continue
		}

		var data api.GatewayCreationProcessData
		if err := json.Unmarshal(p.Data, &data); err != nil {
			continue
		}

		if data.ID != "" {
			return fmt.Errorf("a gateway creation is already in progress for gateway %s", data.ID)
		}
	}

	return nil
}
