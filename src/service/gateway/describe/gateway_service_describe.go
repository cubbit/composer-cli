package describe

import (
	"fmt"
	"strings"

	"github.com/cubbit/composer-cli/constants"
	api "github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/utils/printer"
	"github.com/spf13/cobra"
)

type Dependencies struct {
	GatewayAPI api.GatewayAPIInterface
}

func Describe(deps Dependencies, cmd *cobra.Command, handler configuration_handler.ConfigurationHandlerInterface, profile configuration_models.ProfileV2, args []string) error {
	gatewayID, err := resolveGatewayID(deps, cmd, args, profile)
	if err != nil {
		return err
	}

	gateway, err := deps.GatewayAPI.GetGatewayV5(profile.Endpoints, profile.APIKey, profile.OrganizationID, gatewayID)
	if err != nil {
		return fmt.Errorf("failed to describe gateway: %w", err)
	}

	return printer.ComposeStructured(cmd, handler, gateway,
		func() error { return PrintGatewayDetails(cmd, handler, *gateway) },
	)
}

func resolveGatewayID(deps Dependencies, cmd *cobra.Command, args []string, profile configuration_models.ProfileV2) (string, error) {
	identifiers := 0

	gatewayIDFlag, err := cmd.Flags().GetString("gateway-id")
	if err != nil {
		return "", fmt.Errorf("%s gateway-id: %w", constants.ErrorRetrievingField, err)
	}
	if gatewayIDFlag != "" {
		identifiers++
	}

	gatewayNameFlag, err := cmd.Flags().GetString("gateway-name")
	if err != nil {
		return "", fmt.Errorf("%s gateway-name: %w", constants.ErrorRetrievingField, err)
	}
	if gatewayNameFlag != "" {
		identifiers++
	}

	gatewayIDPositional := ""
	if len(args) > 0 {
		gatewayIDPositional = strings.TrimSpace(args[0])
		if gatewayIDPositional != "" {
			identifiers++
		}
	}

	if identifiers != 1 {
		return "", fmt.Errorf("specify exactly one of GATEWAY_ID, --gateway-id or --gateway-name")
	}

	if gatewayIDPositional != "" {
		return gatewayIDPositional, nil
	}

	if gatewayIDFlag != "" {
		return gatewayIDFlag, nil
	}

	gatewayID, err := resolveGatewayIDByName(deps, profile, gatewayNameFlag)
	if err != nil {
		return "", err
	}
	return gatewayID, nil
}

func resolveGatewayIDByName(deps Dependencies, profile configuration_models.ProfileV2, gatewayName string) (string, error) {
	page := 1
	const itemsPerPage = 1000

	for {
		response, err := deps.GatewayAPI.ListGatewaysV5(
			profile.Endpoints,
			profile.APIKey,
			profile.OrganizationID,
			api.WithPage(page),
			api.WithItems(itemsPerPage),
			api.WithFilter(fmt.Sprintf("name:eq(%s)", gatewayName)),
		)
		if err != nil {
			return "", fmt.Errorf("failed to resolve gateway name '%s': %w", gatewayName, err)
		}

		for _, gateway := range response.Data {
			if gateway.Name == gatewayName {
				return gateway.ID, nil
			}
		}

		if response.NextPage == nil {
			break
		}

		page = *response.NextPage
	}

	return "", fmt.Errorf("gateway with name '%s' not found", gatewayName)
}
