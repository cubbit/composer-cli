package describe

import (
	"fmt"
	"strings"

	"github.com/cubbit/composer-cli/constants"
	api "github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration"
	"github.com/cubbit/composer-cli/src/service/gateway/shared"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

type Dependencies struct {
	GatewayAPI api.GatewayAPIInterface
}

func Describe(deps Dependencies, cmd *cobra.Command, resolvedProfile configuration.ResolvedProfile, urls configuration.URLs, args []string) error {
	gatewayID, err := resolveGatewayID(deps, cmd, args, resolvedProfile, urls)
	if err != nil {
		return err
	}

	gateway, err := deps.GatewayAPI.GetGatewayV5(urls, resolvedProfile.APIKey, resolvedProfile.OrganizationID, gatewayID)
	if err != nil {
		return fmt.Errorf("failed to describe gateway: %w", err)
	}

	output, err := shared.ResolveCommandOutput(cmd, resolvedProfile.Output)
	if err != nil {
		return err
	}

	if output == string(configuration.OutputHuman) {
		return PrintGatewayDetails(cmd, *gateway)
	}

	utils.PrintFormattedData(cmd.OutOrStdout(), gateway, output)
	return nil
}

func resolveGatewayID(deps Dependencies, cmd *cobra.Command, args []string, resolvedProfile configuration.ResolvedProfile, urls configuration.URLs) (string, error) {
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

	gatewayID, err := resolveGatewayIDByName(deps, urls, resolvedProfile, gatewayNameFlag)
	if err != nil {
		return "", err
	}
	return gatewayID, nil
}

func resolveGatewayIDByName(deps Dependencies, urls configuration.URLs, resolvedProfile configuration.ResolvedProfile, gatewayName string) (string, error) {
	page := 1
	const itemsPerPage = 1000

	for {
		response, err := deps.GatewayAPI.ListGatewaysV5(
			urls,
			resolvedProfile.APIKey,
			resolvedProfile.OrganizationID,
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
