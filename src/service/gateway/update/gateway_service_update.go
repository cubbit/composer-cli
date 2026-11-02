package update

import (
	"fmt"

	api "github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/utils"
	"github.com/cubbit/composer-cli/utils/printer"
	"github.com/spf13/cobra"
)

type Dependencies struct {
	GatewayAPI api.GatewayAPIInterface
}

func Update(deps Dependencies, cmd *cobra.Command, handler configuration_handler.ConfigurationHandlerInterface, profile configuration_models.ProfileV2, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("GATEWAY_ID is required")
	}

	// Validate the update fields before calling the API so invalid input does
	// not trigger an unnecessary API call or get masked by a 403/404.
	request, err := buildUpdateGatewayV5Request(cmd)
	if err != nil {
		return err
	}

	gatewayID := args[0]

	gateway, err := deps.GatewayAPI.UpdateGatewayV5(
		profile.Endpoints,
		profile.APIKey,
		profile.OrganizationID,
		gatewayID,
		request,
	)
	if err != nil {
		return fmt.Errorf("failed to update gateway: %w", err)
	}

	return printer.ComposeStructured(cmd, handler, gateway,
		func() error { return PrintGatewayUpdate(cmd, handler, *gateway) },
	)
}

func buildUpdateGatewayV5Request(cmd *cobra.Command) (*api.UpdateGatewayV5Request, error) {
	name, err := utils.GetOptionalStringFlag(cmd, "name")
	if err != nil {
		return nil, err
	}

	description, err := utils.GetOptionalStringFlag(cmd, "description")
	if err != nil {
		return nil, err
	}

	location, err := utils.GetOptionalStringFlag(cmd, "location")
	if err != nil {
		return nil, err
	}

	if name == nil && description == nil && location == nil {
		return nil, fmt.Errorf("specify at least one of --name, --description or --location")
	}

	return &api.UpdateGatewayV5Request{
		Name:        name,
		Description: description,
		Location:    location,
	}, nil
}
