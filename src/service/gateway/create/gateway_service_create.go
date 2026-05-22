package create

import (
	"encoding/json"
	"fmt"

	api "github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration"
	"github.com/spf13/cobra"
)

type Dependencies struct {
	GatewayAPI         api.GatewayAPIInterface
	SwarmAPI           api.SwarmAPIInterface
	RedundancyClassAPI api.RedundancyClassAPIInterface
	ProcessAPI         api.ProcessAPIInterface
	LocationAPI        api.LocationAPIInterface
}

func Create(deps Dependencies, cmd *cobra.Command, resolvedProfile configuration.ResolvedProfile, urls configuration.URLs, interactiveMode bool) error {
	if err := checkRunningGatewayProcess(deps, urls, resolvedProfile); err != nil {
		return err
	}

	if interactiveMode {
		return createInteractive(deps, cmd, resolvedProfile, urls)
	}

	return createInline(deps, cmd, resolvedProfile, urls)
}

func checkRunningGatewayProcess(deps Dependencies, urls configuration.URLs, resolvedProfile configuration.ResolvedProfile) error {
	processes, err := deps.ProcessAPI.ListProcesses(
		urls,
		resolvedProfile.APIKey,
		resolvedProfile.OrganizationID,
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
