package cmd_gateway

import (
	"errors"
	"strings"
	"testing"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
)

func TestGatewayConnectionList_Integration_HumanOutput(t *testing.T) {
	mockGatewayAPI := &api.MockGatewayAPI{
		ListGatewayConnectionsV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
		) (*api.GatewayConnectionsV5ListResponse, error) {
			if apiKey != "test-api-key" || organizationID != "test-org-id" {
				t.Fatalf("Expected credentials to be propagated")
			}
			if gatewayID != "gw-001" {
				t.Fatalf("Expected gateway ID gw-001, got %q", gatewayID)
			}

			return &api.GatewayConnectionsV5ListResponse{
				Data: []api.GatewayConnectionV5{
					{
						ID:              testConnectionID,
						Swarm:           api.GatewayConnectionV5Swarm{ID: testConnectionSwarmID, Name: "swarm-1"},
						RedundancyClass: api.GatewayConnectionV5RedundancyClass{ID: testConnectionRCID, Name: "default"},
						IsDefault:       true,
						Status:          api.GatewayConnectionStatusV5Active,
						ImpactedRules: []api.GatewayConnectionV5ImpactedRule{
							{ID: "rule-001", Name: "cold-rule", Enabled: true},
						},
					},
				},
			}, nil
		},
	}

	commandOutput, execute := setupGatewayConnectionIntegrationCommand(newGatewayConnectionService(mockGatewayAPI))

	if err := execute([]string{"connection", "list", "--gateway-id", "gw-001"}); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := `╭──────────────────────────────────────┬─────────┬──────────────────┬─────────┬────────┬────────────────╮
│ ID                                   │ Swarm   │ Redundancy Class │ Default │ Status │ Impacted Rules │
├──────────────────────────────────────┼─────────┼──────────────────┼─────────┼────────┼────────────────┤
│ 66666666-6666-4666-8666-666666666666 │ swarm-1 │ default          │ Yes     │ Active │ 1              │
╰──────────────────────────────────────┴─────────┴──────────────────┴─────────┴────────┴────────────────╯
`
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestGatewayConnectionList_Integration_JSONOutput(t *testing.T) {
	mockGatewayAPI := &api.MockGatewayAPI{
		ListGatewayConnectionsV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
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

	if err := execute([]string{"connection", "list", "--gateway-id", "gw-001", "--output", "json"}); err != nil {
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

func TestGatewayConnectionList_Integration_YAMLOutput(t *testing.T) {
	mockGatewayAPI := &api.MockGatewayAPI{
		ListGatewayConnectionsV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
		) (*api.GatewayConnectionsV5ListResponse, error) {
			return &api.GatewayConnectionsV5ListResponse{
				Data: []api.GatewayConnectionV5{
					{
						ID:              testConnectionID,
						Swarm:           api.GatewayConnectionV5Swarm{ID: testConnectionSwarmID, Name: "swarm-1"},
						RedundancyClass: api.GatewayConnectionV5RedundancyClass{ID: testConnectionRCID, Name: "default"},
						IsDefault:       true,
						Status:          api.GatewayConnectionStatusV5Active,
						ImpactedRules:   []api.GatewayConnectionV5ImpactedRule{},
					},
				},
			}, nil
		},
	}

	commandOutput, execute := setupGatewayConnectionIntegrationCommand(newGatewayConnectionService(mockGatewayAPI))

	if err := execute([]string{"connection", "list", "--gateway-id", "gw-001", "--output", "yaml"}); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := `- id: 66666666-6666-4666-8666-666666666666
  swarm:
    id: 11111111-1111-4111-8111-111111111111
    name: swarm-1
  redundancy_class:
    id: 33333333-3333-4333-8333-333333333333
    name: default
  is_default: true
  status: active
  impacted_rules: []
`
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestGatewayConnectionList_Integration_GatewayIDFromFlag(t *testing.T) {
	var requestedGatewayID string

	mockGatewayAPI := &api.MockGatewayAPI{
		ListGatewayConnectionsV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
		) (*api.GatewayConnectionsV5ListResponse, error) {
			requestedGatewayID = gatewayID
			return &api.GatewayConnectionsV5ListResponse{Data: []api.GatewayConnectionV5{}}, nil
		},
	}

	_, execute := setupGatewayConnectionIntegrationCommand(newGatewayConnectionService(mockGatewayAPI))

	if err := execute([]string{"connection", "list", "--gateway-id", "gw-003"}); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if requestedGatewayID != "gw-003" {
		t.Fatalf("Expected gateway ID gw-003, got %q", requestedGatewayID)
	}
}

func TestGatewayConnectionList_Integration_PropagatesAPIError(t *testing.T) {
	mockGatewayAPI := &api.MockGatewayAPI{
		ListGatewayConnectionsV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
		) (*api.GatewayConnectionsV5ListResponse, error) {
			return nil, errors.New("code status expected 200, but received 500 instead")
		},
	}

	commandOutput, execute := setupGatewayConnectionIntegrationCommand(newGatewayConnectionService(mockGatewayAPI))

	err := execute([]string{"connection", "list", "--gateway-id", "gw-001"})
	if err == nil {
		t.Fatal("Expected a command-level error, got nil")
	}

	if !strings.Contains(err.Error(), "failed to list connections") ||
		!strings.Contains(err.Error(), "received 500") {
		t.Fatalf("Expected wrapped API error, got %q", err.Error())
	}
	if !strings.Contains(commandOutput.String(), "Error: failed to list connections") {
		t.Fatalf("Expected failure output, got %q", commandOutput.String())
	}
	if strings.Contains(commandOutput.String(), "Impacted Rules") {
		t.Fatalf("Expected no success output on failure, got %q", commandOutput.String())
	}
}
