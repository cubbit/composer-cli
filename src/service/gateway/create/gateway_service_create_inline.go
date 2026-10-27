package create

import (
	"fmt"
	"strings"

	"github.com/cubbit/composer-cli/constants"
	api "github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/utils"
	"github.com/cubbit/composer-cli/utils/printer"
	"github.com/spf13/cobra"
)

func createInline(deps Dependencies, cmd *cobra.Command, handler configuration_handler.ConfigurationHandlerInterface, endpoints configuration_models.EndpointsV2, apiKey string, organizationID string) error {
	createRequest, err := collectGatewayCreateFlags(cmd)
	if err != nil {
		return err
	}

	response, err := deps.GatewayAPI.CreateGatewayV5(endpoints, apiKey, organizationID, createRequest)
	if err != nil {
		return fmt.Errorf("failed to create gateway: %w", err)
	}

	process, err := deps.ProcessAPI.GetProcess(
		endpoints,
		apiKey,
		organizationID,
		response.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to retrieve gateway creation process: %w", err)
	}

	gatewayCreationProcess, ok := process.CastToGatewayCreationProcess()
	if !ok {
		return fmt.Errorf("unexpected process type: expected gateway creation process, got %s", process.Type)
	}

	return printer.PrintText(cmd, handler, fmt.Sprintf("Gateway creation started — Gateway ID: %s\n", gatewayCreationProcess.Data.ID))
}

func collectGatewayCreateFlags(cmd *cobra.Command) (*api.CreateGatewayV5Request, error) {
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		return nil, fmt.Errorf("%s name: %w", constants.ErrorRetrievingField, err)
	}

	slug, err := cmd.Flags().GetString("slug")
	if err != nil {
		return nil, fmt.Errorf("%s slug: %w", constants.ErrorRetrievingField, err)
	}

	clusterID, err := cmd.Flags().GetString("cluster-id")
	if err != nil {
		return nil, fmt.Errorf("%s cluster-id: %w", constants.ErrorRetrievingField, err)
	}

	description, err := utils.GetOptionalStringFlag(cmd, "description")
	if err != nil {
		return nil, fmt.Errorf("%s description: %w", constants.ErrorRetrievingField, err)
	}

	ingressTypeStr, err := cmd.Flags().GetString("ingress-type")
	if err != nil {
		return nil, fmt.Errorf("%s ingress-type: %w", constants.ErrorRetrievingField, err)
	}

	swarmRCFlags, err := cmd.Flags().GetStringArray("swarm-rc")
	if err != nil {
		return nil, fmt.Errorf("%s swarm-rc: %w", constants.ErrorRetrievingField, err)
	}

	swarmsAndRC, err := parseSwarmRCFlags(swarmRCFlags)
	if err != nil {
		return nil, err
	}

	ingress := api.CubbitIngress{
		Type: api.CubbitIngressType(ingressTypeStr),
	}

	return &api.CreateGatewayV5Request{
		ClusterID:                clusterID,
		Name:                     name,
		Slug:                     slug,
		Description:              description,
		SwarmsAndRedundancyClass: swarmsAndRC,
		CubbitIngress:            ingress,
	}, nil
}

func parseSwarmRCFlags(flags []string) ([]api.SwarmAndRedundancyClassV5, error) {
	if len(flags) == 0 {
		return nil, fmt.Errorf("at least one --swarm-rc flag is required")
	}

	result := make([]api.SwarmAndRedundancyClassV5, 0, len(flags))
	hasDefault := false

	for i, flag := range flags {
		parts := strings.SplitN(flag, ":", 3)
		if len(parts) != 3 {
			return nil, fmt.Errorf("invalid --swarm-rc format at position %d: expected 'swarm-id:rc-id:is-default', got '%s'", i, flag)
		}

		swarmID := strings.TrimSpace(parts[0])
		rcID := strings.TrimSpace(parts[1])
		isDefault := strings.TrimSpace(parts[2]) == "true"

		if swarmID == "" {
			return nil, fmt.Errorf("swarm ID cannot be empty at position %d", i)
		}
		if rcID == "" {
			return nil, fmt.Errorf("redundancy class ID cannot be empty at position %d", i)
		}

		if isDefault {
			if hasDefault {
				return nil, fmt.Errorf("only one swarm-rc pair can be set as default, found multiple at positions")
			}
			hasDefault = true
		}

		result = append(result, api.SwarmAndRedundancyClassV5{
			SwarmID:           swarmID,
			RedundancyClassID: rcID,
			IsDefault:         isDefault,
		})
	}

	if !hasDefault {
		return nil, fmt.Errorf("exactly one --swarm-rc must be marked as default (set is-default to 'true')")
	}

	return result, nil
}
