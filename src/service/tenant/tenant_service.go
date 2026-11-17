package tenant

import (
	"fmt"

	"github.com/cubbit/composer-cli/constants"
	api "github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/src/service/tenant/create"
	"github.com/cubbit/composer-cli/src/service/tenant/describe"
	"github.com/cubbit/composer-cli/src/service/tenant/list"
	listconnections "github.com/cubbit/composer-cli/src/service/tenant/list_connections"
	verifyconnection "github.com/cubbit/composer-cli/src/service/tenant/verify_connection"
	"github.com/spf13/cobra"
)

type TenantServiceInterface interface {
	Create(cmd *cobra.Command, args []string) error
	List(cmd *cobra.Command, args []string) error
	Describe(cmd *cobra.Command, args []string) error
	ListConnections(cmd *cobra.Command, args []string) error
	VerifyConnection(cmd *cobra.Command, args []string) error
}

type TenantService struct {
	configuration configuration_handler.ConfigurationHandlerInterface
	tenantAPI     api.TenantAPIInterface
	domainAPI     api.DomainAPIInterface
	gatewayAPI    api.GatewayAPIInterface
	processAPI    api.ProcessAPIInterface
	connectionAPI api.ConnectionAPIInterface
}

func NewTenantService(
	configuration configuration_handler.ConfigurationHandlerInterface,
	tenantAPI api.TenantAPIInterface,
	domainAPI api.DomainAPIInterface,
	gatewayAPI api.GatewayAPIInterface,
	processAPI api.ProcessAPIInterface,
	connectionAPI api.ConnectionAPIInterface,
) TenantService {
	return TenantService{
		configuration: configuration,
		tenantAPI:     tenantAPI,
		domainAPI:     domainAPI,
		gatewayAPI:    gatewayAPI,
		processAPI:    processAPI,
		connectionAPI: connectionAPI,
	}
}

func (s TenantService) List(cmd *cobra.Command, args []string) error {
	profile, err := s.configuration.GetActiveProfile()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	deps := list.Dependencies{
		TenantAPI: s.tenantAPI,
	}

	return list.List(deps, cmd, s.configuration, profile)
}

func (s TenantService) Describe(cmd *cobra.Command, args []string) error {
	profile, err := s.configuration.GetActiveProfile()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	deps := describe.Dependencies{
		TenantAPI: s.tenantAPI,
	}

	return describe.Describe(deps, cmd, s.configuration, profile, args)
}

func (s TenantService) Create(cmd *cobra.Command, args []string) error {
	profile, err := s.configuration.GetActiveProfile()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	interactiveMode, err := cmd.Flags().GetBool("interactive")
	if err != nil {
		return fmt.Errorf("%s interactive: %w", constants.ErrorRetrievingField, err)
	}

	deps := create.Dependencies{
		TenantAPI:  s.tenantAPI,
		DomainAPI:  s.domainAPI,
		GatewayAPI: s.gatewayAPI,
		ProcessAPI: s.processAPI,
	}

	return create.Create(deps, cmd, s.configuration, profile, interactiveMode)
}

func (s TenantService) ListConnections(cmd *cobra.Command, args []string) error {
	profile, err := s.configuration.GetActiveProfile()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	deps := listconnections.Dependencies{
		ConnectionAPI: s.connectionAPI,
		TenantAPI:     s.tenantAPI,
	}

	return listconnections.ListConnections(deps, cmd, s.configuration, profile, args)
}

func (s TenantService) VerifyConnection(cmd *cobra.Command, args []string) error {
	profile, err := s.configuration.GetActiveProfile()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	deps := verifyconnection.Dependencies{
		ConnectionAPI: s.connectionAPI,
		TenantAPI:     s.tenantAPI,
	}

	return verifyconnection.VerifyConnection(deps, cmd, s.configuration, profile, args)
}


