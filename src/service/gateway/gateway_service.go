package gateway

import (
	"fmt"

	"github.com/cubbit/composer-cli/constants"
	api "github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration"
	"github.com/cubbit/composer-cli/src/service/gateway/create"
	"github.com/cubbit/composer-cli/src/service/gateway/list"
	"github.com/spf13/cobra"
)

type GatewayServiceInterface interface {
	Create(cmd *cobra.Command, args []string) error
	List(cmd *cobra.Command, args []string) error
}

type GatewayService struct {
	configuration      configuration.ConfigInterface
	gatewayAPI         api.GatewayAPIInterface
	swarmAPI           api.SwarmAPIInterface
	redundancyClassAPI api.RedundancyClassAPIInterface
	processAPI         api.ProcessAPIInterface
	locationAPI        api.LocationAPIInterface
}

func NewGatewayService(
	configuration configuration.ConfigInterface,
	gatewayAPI api.GatewayAPIInterface,
	swarmAPI api.SwarmAPIInterface,
	redundancyClassAPI api.RedundancyClassAPIInterface,
	processAPI api.ProcessAPIInterface,
	locationAPI api.LocationAPIInterface,
) GatewayService {
	return GatewayService{
		configuration:      configuration,
		gatewayAPI:         gatewayAPI,
		swarmAPI:           swarmAPI,
		redundancyClassAPI: redundancyClassAPI,
		processAPI:         processAPI,
		locationAPI:        locationAPI,
	}
}

func (s GatewayService) List(cmd *cobra.Command, args []string) error {
	resolvedProfile, urls, err := s.configuration.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer)
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	deps := list.Dependencies{
		GatewayAPI: s.gatewayAPI,
	}

	return list.ListInline(deps, cmd, *resolvedProfile, *urls)
}

func (s GatewayService) Create(cmd *cobra.Command, args []string) error {
	resolvedProfile, urls, err := s.configuration.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer)
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	interactiveMode, err := cmd.Flags().GetBool("interactive")
	if err != nil {
		return fmt.Errorf("%s interactive: %w", constants.ErrorRetrievingField, err)
	}

	deps := create.Dependencies{
		GatewayAPI:         s.gatewayAPI,
		SwarmAPI:           s.swarmAPI,
		RedundancyClassAPI: s.redundancyClassAPI,
		ProcessAPI:         s.processAPI,
		LocationAPI:        s.locationAPI,
	}

	return create.Create(deps, cmd, *resolvedProfile, *urls, interactiveMode)
}
