package cmd_gateway

import (
	"errors"
	"strings"
	"testing"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
)

func TestGatewayStorageClassList_Integration_HumanOutput(t *testing.T) {
	mockGatewayAPI := &api.MockGatewayAPI{
		ListGatewayStorageClassesV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
		) (*api.GatewayStorageClassesV5ListResponse, error) {
			if apiKey != "test-api-key" || organizationID != "test-org-id" {
				t.Fatalf("Expected credentials to be propagated")
			}
			if gatewayID != "gw-001" {
				t.Fatalf("Expected gateway ID gw-001, got %q", gatewayID)
			}

			return &api.GatewayStorageClassesV5ListResponse{
				Data: []api.GatewayStorageClassV5{
					{
						ID:              "sc-default-001",
						Name:            "default",
						Pool:            true,
						Status:          api.GatewayStorageClassStatusV5Active,
						ReferencedRules: []api.GatewayStorageClassV5ReferencedRule{},
					},
					{
						ID:     "sc-001",
						Name:   "cold-data",
						Status: api.GatewayStorageClassStatusV5FallbackOnly,
						ReferencedRules: []api.GatewayStorageClassV5ReferencedRule{
							{ID: "rule-001", Name: "cold-rule"},
						},
					},
				},
			}, nil
		},
	}

	commandOutput, execute := setupGatewayStorageClassIntegrationCommand(newGatewayStorageClassService(mockGatewayAPI))

	if err := execute([]string{"storage-class", "list", "--gateway-id", "gw-001"}); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := `╭───────────┬────────────────┬───────────────┬─────────┬──────────────────╮
│ Name      │ ID             │ Status        │ Default │ Referenced Rules │
├───────────┼────────────────┼───────────────┼─────────┼──────────────────┤
│ default   │ sc-default-001 │ Active        │ yes     │ 0                │
│ cold-data │ sc-001         │ Fallback only │ -       │ 1                │
╰───────────┴────────────────┴───────────────┴─────────┴──────────────────╯
`
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestGatewayStorageClassList_Integration_JSONOutput(t *testing.T) {
	mockGatewayAPI := &api.MockGatewayAPI{
		ListGatewayStorageClassesV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
		) (*api.GatewayStorageClassesV5ListResponse, error) {
			return &api.GatewayStorageClassesV5ListResponse{
				Data: []api.GatewayStorageClassV5{
					{
						ID:     "sc-001",
						Name:   "cold-data",
						Status: api.GatewayStorageClassStatusV5Active,
						ReferencedRules: []api.GatewayStorageClassV5ReferencedRule{
							{ID: "rule-001", Name: "cold-rule"},
						},
					},
				},
			}, nil
		},
	}

	commandOutput, execute := setupGatewayStorageClassIntegrationCommand(newGatewayStorageClassService(mockGatewayAPI))

	if err := execute([]string{"storage-class", "list", "--gateway-id", "gw-001", "--output", "json"}); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := `[
  {
    "id": "sc-001",
    "name": "cold-data",
    "pool": false,
    "status": "active",
    "referenced_rules": [
      {
        "id": "rule-001",
        "name": "cold-rule"
      }
    ]
  }
]
`
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestGatewayStorageClassList_Integration_YAMLOutput(t *testing.T) {
	mockGatewayAPI := &api.MockGatewayAPI{
		ListGatewayStorageClassesV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
		) (*api.GatewayStorageClassesV5ListResponse, error) {
			return &api.GatewayStorageClassesV5ListResponse{
				Data: []api.GatewayStorageClassV5{
					{
						ID:              "sc-default-001",
						Name:            "default",
						Pool:            true,
						Status:          api.GatewayStorageClassStatusV5Active,
						ReferencedRules: []api.GatewayStorageClassV5ReferencedRule{},
					},
					{
						ID:     "sc-001",
						Name:   "cold-data",
						Status: api.GatewayStorageClassStatusV5FallbackOnly,
						ReferencedRules: []api.GatewayStorageClassV5ReferencedRule{
							{ID: "rule-001", Name: "cold-rule"},
						},
					},
				},
			}, nil
		},
	}

	commandOutput, execute := setupGatewayStorageClassIntegrationCommand(newGatewayStorageClassService(mockGatewayAPI))

	if err := execute([]string{"storage-class", "list", "--gateway-id", "gw-001", "--output", "yaml"}); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := `- id: sc-default-001
  name: default
  pool: true
  status: active
  referenced_rules: []
- id: sc-001
  name: cold-data
  pool: false
  status: fallback_only
  referenced_rules:
    - id: rule-001
      name: cold-rule
`
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestGatewayStorageClassList_Integration_PropagatesAPIError(t *testing.T) {
	mockGatewayAPI := &api.MockGatewayAPI{
		ListGatewayStorageClassesV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
		) (*api.GatewayStorageClassesV5ListResponse, error) {
			return nil, errors.New("code status expected 200, but received 500 instead")
		},
	}

	commandOutput, execute := setupGatewayStorageClassIntegrationCommand(newGatewayStorageClassService(mockGatewayAPI))

	err := execute([]string{"storage-class", "list", "--gateway-id", "gw-001"})
	if err == nil {
		t.Fatal("Expected a command-level error, got nil")
	}

	if !containsAll(err.Error(), "failed to list storage classes", "received 500") {
		t.Fatalf("Expected wrapped API error, got %q", err.Error())
	}
	if !containsAll(commandOutput.String(), "Error: failed to list storage classes", "received 500") {
		t.Fatalf("Expected failure output, got %q", commandOutput.String())
	}
	if strings.Contains(commandOutput.String(), "Referenced Rules") {
		t.Fatalf("Expected no success output on failure, got %q", commandOutput.String())
	}
}
