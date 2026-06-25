package tenant

import (
	"fmt"

	"github.com/cubbit/composer-cli/constants"
	api "github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/src/service/tenant/create"
	"github.com/spf13/cobra"
)

type TenantServiceInterface interface {
	Create(cmd *cobra.Command, args []string) error
}

type TenantService struct {
	configuration configuration_handler.ConfigurationHandlerInterface
	tenantAPI     api.TenantAPIInterface
	domainAPI     api.DomainAPIInterface
	gatewayAPI    api.GatewayAPIInterface
	processAPI    api.ProcessAPIInterface
}

func NewTenantService(
	configuration configuration_handler.ConfigurationHandlerInterface,
	tenantAPI api.TenantAPIInterface,
	domainAPI api.DomainAPIInterface,
	gatewayAPI api.GatewayAPIInterface,
	processAPI api.ProcessAPIInterface,
) TenantService {
	return TenantService{
		configuration: configuration,
		tenantAPI:     tenantAPI,
		domainAPI:     domainAPI,
		gatewayAPI:    gatewayAPI,
		processAPI:    processAPI,
	}
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

	return create.Create(deps, cmd, profile, interactiveMode)
}
