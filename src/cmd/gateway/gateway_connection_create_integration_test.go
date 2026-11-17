package cmd_gateway

import (
	"errors"
	"strings"
	"testing"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
)

func TestGatewayConnectionCreate_Integration_HumanOutput(t *testing.T) {
	mockGatewayAPI := &api.MockGatewayAPI{
		CreateGatewayConnectionsV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			request *api.CreateGatewayConnectionsV5Request,
		) (*api.GatewayConnectionsV5ListResponse, error) {
			if gatewayID != "gw-001" {
				t.Fatalf("Expected gateway ID gw-001, got %q", gatewayID)
			}
			if len(request.Connections) != 1 {
				t.Fatalf("Expected one grouped swarm, got %+v", request.Connections)
			}
			if request.Connections[0].SwarmID != testConnectionSwarmID {
				t.Fatalf("Expected swarm %s, got %q", testConnectionSwarmID, request.Connections[0].SwarmID)
			}
			if strings.Join(request.Connections[0].RedundancyClassIDs, ",") != testConnectionRCID+","+testConnectionRCID2 {
				t.Fatalf("Expected two redundancy classes, got %v", request.Connections[0].RedundancyClassIDs)
			}

			return &api.GatewayConnectionsV5ListResponse{
				Data: []api.GatewayConnectionV5{
					{
						ID:              testConnectionID,
						Swarm:           api.GatewayConnectionV5Swarm{ID: testConnectionSwarmID, Name: "swarm-1"},
						RedundancyClass: api.GatewayConnectionV5RedundancyClass{ID: testConnectionRCID, Name: "default"},
						Status:          api.GatewayConnectionStatusV5NotActive,
						ImpactedRules:   []api.GatewayConnectionV5ImpactedRule{},
					},
					{
						ID:              "77777777-7777-4777-8777-777777777777",
						Swarm:           api.GatewayConnectionV5Swarm{ID: testConnectionSwarmID, Name: "swarm-1"},
						RedundancyClass: api.GatewayConnectionV5RedundancyClass{ID: testConnectionRCID2, Name: "replicated"},
						Status:          api.GatewayConnectionStatusV5NotActive,
						ImpactedRules:   []api.GatewayConnectionV5ImpactedRule{},
					},
				},
			}, nil
		},
	}

	commandOutput, execute := setupGatewayConnectionIntegrationCommand(newGatewayConnectionService(mockGatewayAPI))

	if err := execute([]string{
		"connection", "create",
		"--gateway-id", "gw-001",
		"--connection", testConnectionSwarmID + ":" + testConnectionRCID + "," + testConnectionRCID2,
	}); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := `2 connections created successfully.

╭──────────────────────────────────────┬─────────┬──────────────────┬─────────┬────────────┬────────────────╮
│ ID                                   │ Swarm   │ Redundancy Class │ Default │ Status     │ Impacted Rules │
├──────────────────────────────────────┼─────────┼──────────────────┼─────────┼────────────┼────────────────┤
│ 66666666-6666-4666-8666-666666666666 │ swarm-1 │ default          │ -       │ Not active │ 0              │
│ 77777777-7777-4777-8777-777777777777 │ swarm-1 │ replicated       │ -       │ Not active │ 0              │
╰──────────────────────────────────────┴─────────┴──────────────────┴─────────┴────────────┴────────────────╯
`
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestGatewayConnectionCreate_Integration_SingularHumanOutput(t *testing.T) {
	mockGatewayAPI := &api.MockGatewayAPI{
		CreateGatewayConnectionsV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			request *api.CreateGatewayConnectionsV5Request,
		) (*api.GatewayConnectionsV5ListResponse, error) {
			return &api.GatewayConnectionsV5ListResponse{
				Data: []api.GatewayConnectionV5{
					{
						ID:              testConnectionID,
						Swarm:           api.GatewayConnectionV5Swarm{ID: testConnectionSwarmID, Name: "swarm-1"},
						RedundancyClass: api.GatewayConnectionV5RedundancyClass{ID: testConnectionRCID, Name: "default"},
						IsDefault:       false,
						Status:          api.GatewayConnectionStatusV5NotActive,
						ImpactedRules:   []api.GatewayConnectionV5ImpactedRule{},
					},
				},
			}, nil
		},
	}

	commandOutput, execute := setupGatewayConnectionIntegrationCommand(newGatewayConnectionService(mockGatewayAPI))

	if err := execute([]string{
		"connection", "create",
		"--gateway-id", "gw-001",
		"--connection", testConnectionSwarmID + ":" + testConnectionRCID,
	}); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := `1 connection created successfully.

╭──────────────────────────────────────┬─────────┬──────────────────┬─────────┬────────────┬────────────────╮
│ ID                                   │ Swarm   │ Redundancy Class │ Default │ Status     │ Impacted Rules │
├──────────────────────────────────────┼─────────┼──────────────────┼─────────┼────────────┼────────────────┤
│ 66666666-6666-4666-8666-666666666666 │ swarm-1 │ default          │ -       │ Not active │ 0              │
╰──────────────────────────────────────┴─────────┴──────────────────┴─────────┴────────────┴────────────────╯
`
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestGatewayConnectionCreate_Integration_JSONOutput(t *testing.T) {
	mockGatewayAPI := &api.MockGatewayAPI{
		CreateGatewayConnectionsV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			request *api.CreateGatewayConnectionsV5Request,
		) (*api.GatewayConnectionsV5ListResponse, error) {
			return &api.GatewayConnectionsV5ListResponse{
				Data: []api.GatewayConnectionV5{
					{
						ID:              testConnectionID,
						Swarm:           api.GatewayConnectionV5Swarm{ID: testConnectionSwarmID, Name: "swarm-1"},
						RedundancyClass: api.GatewayConnectionV5RedundancyClass{ID: testConnectionRCID, Name: "default"},
						Status:          api.GatewayConnectionStatusV5NotActive,
						ImpactedRules:   []api.GatewayConnectionV5ImpactedRule{},
					},
				},
			}, nil
		},
	}

	commandOutput, execute := setupGatewayConnectionIntegrationCommand(newGatewayConnectionService(mockGatewayAPI))

	if err := execute([]string{
		"connection", "create",
		"--gateway-id", "gw-001",
		"--connection", testConnectionSwarmID + ":" + testConnectionRCID,
		"--output", "json",
	}); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := `[
  {
    "id": "66666666-6666-4666-8666-666666666666",
    "swarm": {
      "id": "11111111-1111-4111-8111-111111111111",
      "name": "swarm-1"
    },
    "redundancy_class": {
      "id": "33333333-3333-4333-8333-333333333333",
      "name": "default"
    },
    "is_default": false,
    "status": "not_active",
    "impacted_rules": []
  }
]
`
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestGatewayConnectionCreate_Integration_GatewayIDFromFlag(t *testing.T) {
	var requestedGatewayID string

	mockGatewayAPI := &api.MockGatewayAPI{
		CreateGatewayConnectionsV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			request *api.CreateGatewayConnectionsV5Request,
		) (*api.GatewayConnectionsV5ListResponse, error) {
			requestedGatewayID = gatewayID
			return &api.GatewayConnectionsV5ListResponse{}, nil
		},
	}

	_, execute := setupGatewayConnectionIntegrationCommand(newGatewayConnectionService(mockGatewayAPI))

	if err := execute([]string{
		"connection", "create",
		"--gateway-id", "gw-003",
		"--connection", testConnectionSwarmID + ":" + testConnectionRCID,
	}); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if requestedGatewayID != "gw-003" {
		t.Fatalf("Expected gateway ID gw-003, got %q", requestedGatewayID)
	}
}

func TestGatewayConnectionCreate_Integration_PropagatesConflict(t *testing.T) {
	mockGatewayAPI := &api.MockGatewayAPI{
		CreateGatewayConnectionsV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			request *api.CreateGatewayConnectionsV5Request,
		) (*api.GatewayConnectionsV5ListResponse, error) {
			return nil, errors.New("code status expected 201, but received 409 instead\nconnection already exists")
		},
	}

	commandOutput, execute := setupGatewayConnectionIntegrationCommand(newGatewayConnectionService(mockGatewayAPI))

	err := execute([]string{
		"connection", "create",
		"--gateway-id", "gw-001",
		"--connection", testConnectionSwarmID + ":" + testConnectionRCID,
	})
	if err == nil {
		t.Fatal("Expected a command-level error, got nil")
	}

	if !strings.Contains(err.Error(), "failed to create connections") ||
		!strings.Contains(err.Error(), "received 409") {
		t.Fatalf("Expected conflict error, got %q", err.Error())
	}
	if !strings.Contains(commandOutput.String(), "Error: failed to create connections") {
		t.Fatalf("Expected failure output, got %q", commandOutput.String())
	}
	if strings.Contains(commandOutput.String(), "created successfully") {
		t.Fatalf("Expected no success output on failure, got %q", commandOutput.String())
	}
}
