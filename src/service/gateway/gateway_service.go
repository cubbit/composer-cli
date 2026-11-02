package gateway

import (
	"fmt"

	"github.com/cubbit/composer-cli/constants"
	api "github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/src/service/gateway/connection"
	"github.com/cubbit/composer-cli/src/service/gateway/create"
	"github.com/cubbit/composer-cli/src/service/gateway/describe"
	"github.com/cubbit/composer-cli/src/service/gateway/list"
	"github.com/cubbit/composer-cli/src/service/gateway/routing_rules"
	"github.com/cubbit/composer-cli/src/service/gateway/storageclass"
	"github.com/cubbit/composer-cli/src/service/gateway/update"
	"github.com/spf13/cobra"
)

type GatewayServiceInterface interface {
	Create(cmd *cobra.Command, args []string) error
	List(cmd *cobra.Command, args []string) error
	Describe(cmd *cobra.Command, args []string) error
	Update(cmd *cobra.Command, args []string) error
	ListStorageClasses(cmd *cobra.Command, args []string) error
	CreateStorageClass(cmd *cobra.Command, args []string) error
	DeleteStorageClass(cmd *cobra.Command, args []string) error
	ListConnections(cmd *cobra.Command, args []string) error
	CreateConnections(cmd *cobra.Command, args []string) error
	DeleteConnection(cmd *cobra.Command, args []string) error
	ListRoutingRules(cmd *cobra.Command, args []string) error
	GetRoutingRule(cmd *cobra.Command, args []string) error
	CreateRoutingRule(cmd *cobra.Command, args []string) error
	UpdateRoutingRule(cmd *cobra.Command, args []string) error
	DeleteRoutingRule(cmd *cobra.Command, args []string) error
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

	return list.ListInline(deps, cmd, s.configuration, profile)
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

	return create.Create(deps, cmd, s.configuration, profile, interactiveMode)
}

func (s GatewayService) Describe(cmd *cobra.Command, args []string) error {
	profile, err := s.configuration.GetActiveProfile()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	deps := describe.Dependencies{
		GatewayAPI: s.gatewayAPI,
	}

	return describe.Describe(deps, cmd, s.configuration, profile, args)
}

func (s GatewayService) Update(cmd *cobra.Command, args []string) error {
	profile, err := s.configuration.GetActiveProfile()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	deps := update.Dependencies{
		GatewayAPI: s.gatewayAPI,
	}

	return update.Update(deps, cmd, s.configuration, profile, args)
}

func (s GatewayService) ListStorageClasses(cmd *cobra.Command, args []string) error {
	profile, err := s.configuration.GetActiveProfile()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	deps := storageclass.Dependencies{
		GatewayAPI: s.gatewayAPI,
	}

	return storageclass.List(deps, cmd, s.configuration, profile, args)
}

func (s GatewayService) CreateStorageClass(cmd *cobra.Command, args []string) error {
	profile, err := s.configuration.GetActiveProfile()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	deps := storageclass.Dependencies{
		GatewayAPI: s.gatewayAPI,
	}

	return storageclass.Create(deps, cmd, s.configuration, profile, args)
}

func (s GatewayService) DeleteStorageClass(cmd *cobra.Command, args []string) error {
	profile, err := s.configuration.GetActiveProfile()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	deps := storageclass.Dependencies{
		GatewayAPI: s.gatewayAPI,
	}

	return storageclass.Delete(deps, cmd, s.configuration, profile, args)
}

func (s GatewayService) ListConnections(cmd *cobra.Command, args []string) error {
	profile, err := s.configuration.GetActiveProfile()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	deps := connection.Dependencies{
		GatewayAPI: s.gatewayAPI,
	}

	return connection.List(deps, cmd, s.configuration, profile, args)
}

func (s GatewayService) CreateConnections(cmd *cobra.Command, args []string) error {
	profile, err := s.configuration.GetActiveProfile()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	deps := connection.Dependencies{
		GatewayAPI: s.gatewayAPI,
	}

	return connection.Create(deps, cmd, s.configuration, profile, args)
}

func (s GatewayService) DeleteConnection(cmd *cobra.Command, args []string) error {
	profile, err := s.configuration.GetActiveProfile()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	deps := connection.Dependencies{
		GatewayAPI: s.gatewayAPI,
	}

	return connection.Delete(deps, cmd, s.configuration, profile, args)
}

func (s GatewayService) ListRoutingRules(cmd *cobra.Command, args []string) error {
	profile, err := s.configuration.GetActiveProfile()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	deps := routing_rules.Dependencies{
		GatewayAPI: s.gatewayAPI,
	}

	return routing_rules.List(deps, cmd, s.configuration, profile, args)
}

func (s GatewayService) GetRoutingRule(cmd *cobra.Command, args []string) error {
	profile, err := s.configuration.GetActiveProfile()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	deps := routing_rules.Dependencies{
		GatewayAPI: s.gatewayAPI,
	}

	return routing_rules.Get(deps, cmd, s.configuration, profile, args)
}

func (s GatewayService) CreateRoutingRule(cmd *cobra.Command, args []string) error {
	profile, err := s.configuration.GetActiveProfile()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	deps := routing_rules.Dependencies{
		GatewayAPI: s.gatewayAPI,
	}

	return routing_rules.Create(deps, cmd, s.configuration, profile, args)
}

func (s GatewayService) UpdateRoutingRule(cmd *cobra.Command, args []string) error {
	profile, err := s.configuration.GetActiveProfile()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	deps := routing_rules.Dependencies{
		GatewayAPI: s.gatewayAPI,
	}

	return routing_rules.Update(deps, cmd, s.configuration, profile, args)
}

func (s GatewayService) DeleteRoutingRule(cmd *cobra.Command, args []string) error {
	profile, err := s.configuration.GetActiveProfile()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	deps := routing_rules.Dependencies{
		GatewayAPI: s.gatewayAPI,
	}

	return routing_rules.Delete(deps, cmd, s.configuration, profile, args)
}
