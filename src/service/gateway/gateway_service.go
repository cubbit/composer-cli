package gateway

import (
	"fmt"

	"github.com/cubbit/composer-cli/constants"
	api "github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/src/service/gateway/create"
	"github.com/cubbit/composer-cli/src/service/gateway/describe"
	"github.com/cubbit/composer-cli/src/service/gateway/list"
	"github.com/spf13/cobra"
)

type GatewayServiceInterface interface {
	Create(cmd *cobra.Command, args []string) error
	List(cmd *cobra.Command, args []string) error
	Describe(cmd *cobra.Command, args []string) error
}

type GatewayService struct {
	configuration      configuration_handler.ConfigurationHandlerInterface
	gatewayAPI         api.GatewayAPIInterface
	swarmAPI           api.SwarmAPIInterface
	redundancyClassAPI api.RedundancyClassAPIInterface
	processAPI         api.ProcessAPIInterface
	locationAPI        api.LocationAPIInterface
}

func NewGatewayService(
	configuration configuration_handler.ConfigurationHandlerInterface,
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
	profile, err := s.configuration.GetActiveProfile()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}
	deps := list.Dependencies{
		GatewayAPI: s.gatewayAPI,
	}

	return list.ListInline(deps, cmd, profile)
}

func (s GatewayService) Create(cmd *cobra.Command, args []string) error {
	profile, err := s.configuration.GetActiveProfile()
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

	return create.Create(deps, cmd, profile, interactiveMode)
}

func (s GatewayService) Describe(cmd *cobra.Command, args []string) error {
	profile, err := s.configuration.GetActiveProfile()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	deps := describe.Dependencies{
		GatewayAPI: s.gatewayAPI,
	}

	return describe.Describe(deps, cmd, profile, args)
}
