package connection

import (
	"fmt"
	"strings"

	"github.com/cubbit/composer-cli/constants"
	api "github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/utils/printer"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

type Dependencies struct {
	GatewayAPI api.GatewayAPIInterface
}

// gatewayIDFromFlags reads and validates the required --gateway-id flag.
func gatewayIDFromFlags(cmd *cobra.Command) (string, error) {
	gatewayID, err := cmd.Flags().GetString("gateway-id")
	if err != nil {
		return "", fmt.Errorf("%s gateway-id: %w", constants.ErrorRetrievingField, err)
	}

	if strings.TrimSpace(gatewayID) == "" {
		return "", fmt.Errorf("--gateway-id must not be empty")
	}

	return strings.TrimSpace(gatewayID), nil
}

// List prints the swarm and redundancy class pairs a gateway can offload to.
func List(
	deps Dependencies,
	cmd *cobra.Command,
	handler configuration_handler.ConfigurationHandlerInterface,
	profile configuration_models.ProfileV2,
	args []string,
) error {
	gatewayID, err := gatewayIDFromFlags(cmd)
	if err != nil {
		return err
	}

	response, err := deps.GatewayAPI.ListGatewayConnectionsV5(
		profile.Endpoints,
		profile.APIKey,
		profile.OrganizationID,
		gatewayID,
	)
	if err != nil {
		return fmt.Errorf("failed to list connections: %w", err)
	}

	return printer.ComposeStructured(cmd, handler, response.Data,
		func() error { return PrintConnectionList(cmd, handler, response.Data) },
	)
}

// Create adds one or more swarm and redundancy class pairs to a gateway.
func Create(
	deps Dependencies,
	cmd *cobra.Command,
	handler configuration_handler.ConfigurationHandlerInterface,
	profile configuration_models.ProfileV2,
	args []string,
) error {
	connections, err := cmd.Flags().GetStringArray("connection")
	if err != nil {
		return fmt.Errorf("%s connection: %w", constants.ErrorRetrievingField, err)
	}

	request, err := buildCreateConnectionsRequest(connections)
	if err != nil {
		return err
	}

	gatewayID, err := gatewayIDFromFlags(cmd)
	if err != nil {
		return err
	}

	response, err := deps.GatewayAPI.CreateGatewayConnectionsV5(
		profile.Endpoints,
		profile.APIKey,
		profile.OrganizationID,
		gatewayID,
		request,
	)
	if err != nil {
		return fmt.Errorf("failed to create connections: %w", err)
	}

	return printer.ComposeStructured(cmd, handler, response.Data,
		func() error { return PrintConnectionsCreated(cmd, handler, response.Data) },
	)
}

// Delete removes a connection from a gateway. The default connection cannot
// be deleted because it is the smart data placement fallback.
func Delete(
	deps Dependencies,
	cmd *cobra.Command,
	handler configuration_handler.ConfigurationHandlerInterface,
	profile configuration_models.ProfileV2,
	args []string,
) error {
	connectionID, err := cmd.Flags().GetString("id")
	if err != nil {
		return fmt.Errorf("%s id: %w", constants.ErrorRetrievingField, err)
	}

	connectionID = strings.TrimSpace(connectionID)
	if connectionID == "" {
		return fmt.Errorf("--id must not be empty")
	}

	parsedConnectionID, err := uuid.Parse(connectionID)
	if err != nil {
		return fmt.Errorf("--id must be a valid UUID, got %q", connectionID)
	}
	connectionID = parsedConnectionID.String()

	gatewayID, err := gatewayIDFromFlags(cmd)
	if err != nil {
		return err
	}

	if err := deps.GatewayAPI.DeleteGatewayConnectionV5(
		profile.Endpoints,
		profile.APIKey,
		profile.OrganizationID,
		gatewayID,
		connectionID,
	); err != nil {
		return fmt.Errorf("failed to delete connection: %w", err)
	}

	return printer.PrintText(cmd, handler, buildConnectionDeletedOutput(connectionID))
}

// buildCreateConnectionsRequest parses repeated --connection flags. Each flag
// groups one swarm with one or more redundancy classes, in the format
// "<swarm-id>:<redundancy-class-id>[,<redundancy-class-id>...]". Flags that
// target the same swarm are merged into a single request item. Connections and
// redundancy classes are reported 1-based in validation errors.
func buildCreateConnectionsRequest(flags []string) (*api.CreateGatewayConnectionsV5Request, error) {
	if len(flags) == 0 {
		return nil, fmt.Errorf("at least one --connection flag is required")
	}

	items := make([]api.CreateGatewayConnectionV5RequestItem, 0, len(flags))
	itemIndexBySwarm := make(map[string]int)
	seenPairs := make(map[string]bool)

	for position, flag := range flags {
		connectionNumber := position + 1

		parts := strings.SplitN(strings.TrimSpace(flag), ":", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf(
				"invalid --connection format in connection %d: expected '<swarm-id>:<redundancy-class-id>[,<redundancy-class-id>...]', got %q",
				connectionNumber,
				flag,
			)
		}

		swarmID := strings.TrimSpace(parts[0])
		if swarmID == "" {
			return nil, fmt.Errorf("swarm ID cannot be empty in connection %d", connectionNumber)
		}

		parsedSwarmID, err := uuid.Parse(swarmID)
		if err != nil {
			return nil, fmt.Errorf("swarm ID in connection %d is not a valid UUID: %q", connectionNumber, swarmID)
		}
		swarmID = parsedSwarmID.String()

		redundancyClassIDs := strings.Split(parts[1], ",")

		itemIndex, ok := itemIndexBySwarm[swarmID]
		if !ok {
			itemIndex = len(items)
			itemIndexBySwarm[swarmID] = itemIndex
			items = append(items, api.CreateGatewayConnectionV5RequestItem{SwarmID: swarmID})
		}

		for classIndex, rawRedundancyClassID := range redundancyClassIDs {
			redundancyClassNumber := classIndex + 1

			redundancyClassID := strings.TrimSpace(rawRedundancyClassID)
			if redundancyClassID == "" {
				return nil, fmt.Errorf(
					"connection %d, redundancy class %d: redundancy class ID cannot be empty",
					connectionNumber,
					redundancyClassNumber,
				)
			}

			parsedRedundancyClassID, err := uuid.Parse(redundancyClassID)
			if err != nil {
				return nil, fmt.Errorf(
					"connection %d, redundancy class %d: redundancy class ID is not a valid UUID: %q",
					connectionNumber,
					redundancyClassNumber,
					redundancyClassID,
				)
			}
			redundancyClassID = parsedRedundancyClassID.String()

			pairKey := swarmID + ":" + redundancyClassID
			if seenPairs[pairKey] {
				return nil, fmt.Errorf(
					"connection %d, redundancy class %d: duplicate connection for swarm %s and redundancy class %s",
					connectionNumber,
					redundancyClassNumber,
					swarmID,
					redundancyClassID,
				)
			}
			seenPairs[pairKey] = true

			items[itemIndex].RedundancyClassIDs = append(
				items[itemIndex].RedundancyClassIDs,
				redundancyClassID,
			)
		}
	}

	return &api.CreateGatewayConnectionsV5Request{Connections: items}, nil
}
