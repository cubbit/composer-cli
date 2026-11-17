package create

import (
	api "github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/spf13/cobra"
)

type Dependencies struct {
	TenantAPI  api.TenantAPIInterface
	DomainAPI  api.DomainAPIInterface
	GatewayAPI api.GatewayAPIInterface
	ProcessAPI api.ProcessAPIInterface
}

func Create(deps Dependencies, cmd *cobra.Command, handler configuration_handler.ConfigurationHandlerInterface, profile configuration_models.ProfileV2, interactiveMode bool) error {
	if interactiveMode {
		return createInteractive(deps, cmd, profile.Endpoints, profile.APIKey, profile.OrganizationID)
	}

	return createInline(deps, cmd, handler, profile.Endpoints, profile.APIKey, profile.OrganizationID)
}
