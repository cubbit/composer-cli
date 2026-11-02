package cmd_gateway

import (
	"errors"
	"strings"
	"testing"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
)

func TestGatewayStorageClassCreate_Integration_HumanOutput(t *testing.T) {
	mockGatewayAPI := &api.MockGatewayAPI{
		CreateGatewayStorageClassV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			request *api.CreateGatewayStorageClassV5Request,
		) (*api.GatewayStorageClassV5, error) {
			if gatewayID != "gw-001" {
				t.Fatalf("Expected gateway ID gw-001, got %q", gatewayID)
			}
			if request.Name != "cold-data" {
				t.Fatalf("Expected name cold-data, got %q", request.Name)
			}

			return &api.GatewayStorageClassV5{
				ID:              "sc-001",
				Name:            request.Name,
				Status:          api.GatewayStorageClassStatusV5Active,
				ReferencedRules: []api.GatewayStorageClassV5ReferencedRule{},
			}, nil
		},
	}

	commandOutput, execute := setupGatewayStorageClassIntegrationCommand(newGatewayStorageClassService(mockGatewayAPI))

	if err := execute([]string{"storage-class", "create", "--gateway-id", "gw-001", "--name", "cold-data"}); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := `Storage class cold-data created successfully.

Storage Class:
  Name: cold-data
  ID: sc-001
  Status: Active
  Default: -
  Referenced Rules: 0`
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestGatewayStorageClassCreate_Integration_JSONOutput(t *testing.T) {
	mockGatewayAPI := &api.MockGatewayAPI{
		CreateGatewayStorageClassV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			request *api.CreateGatewayStorageClassV5Request,
		) (*api.GatewayStorageClassV5, error) {
			return &api.GatewayStorageClassV5{
				ID:              "sc-001",
				Name:            request.Name,
				Status:          api.GatewayStorageClassStatusV5Active,
				ReferencedRules: []api.GatewayStorageClassV5ReferencedRule{},
			}, nil
		},
	}

	commandOutput, execute := setupGatewayStorageClassIntegrationCommand(newGatewayStorageClassService(mockGatewayAPI))

	if err := execute([]string{"storage-class", "create", "--gateway-id", "gw-001", "--name", "cold-data", "--output", "json"}); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := `{
  "id": "sc-001",
  "name": "cold-data",
  "pool": false,
  "status": "active",
  "referenced_rules": []
}
`
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestGatewayStorageClassCreate_Integration_YAMLOutput(t *testing.T) {
	mockGatewayAPI := &api.MockGatewayAPI{
		CreateGatewayStorageClassV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			request *api.CreateGatewayStorageClassV5Request,
		) (*api.GatewayStorageClassV5, error) {
			return &api.GatewayStorageClassV5{
				ID:     "sc-001",
				Name:   request.Name,
				Status: api.GatewayStorageClassStatusV5Active,
				ReferencedRules: []api.GatewayStorageClassV5ReferencedRule{
					{ID: "rule-001", Name: "cold-rule"},
				},
			}, nil
		},
	}

	commandOutput, execute := setupGatewayStorageClassIntegrationCommand(newGatewayStorageClassService(mockGatewayAPI))

	if err := execute([]string{"storage-class", "create", "--gateway-id", "gw-001", "--name", "cold-data", "--output", "yaml"}); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := `id: sc-001
name: cold-data
pool: false
status: active
referenced_rules:
    - id: rule-001
      name: cold-rule
`
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestGatewayStorageClassCreate_Integration_PropagatesConflict(t *testing.T) {
	mockGatewayAPI := &api.MockGatewayAPI{
		CreateGatewayStorageClassV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			request *api.CreateGatewayStorageClassV5Request,
		) (*api.GatewayStorageClassV5, error) {
			return nil, errors.New("code status expected 201, but received 409 instead\nstorage class already exists")
		},
	}

	commandOutput, execute := setupGatewayStorageClassIntegrationCommand(newGatewayStorageClassService(mockGatewayAPI))

	err := execute([]string{"storage-class", "create", "--gateway-id", "gw-001", "--name", "cold-data"})
	if err == nil {
		t.Fatal("Expected a command-level error, got nil")
	}

	if !containsAll(err.Error(), "failed to create storage class", "received 409") {
		t.Fatalf("Expected conflict error, got %q", err.Error())
	}
	if !containsAll(commandOutput.String(), "Error: failed to create storage class", "received 409") {
		t.Fatalf("Expected failure output, got %q", commandOutput.String())
	}
	if strings.Contains(commandOutput.String(), "created successfully") {
		t.Fatalf("Expected no success output on failure, got %q", commandOutput.String())
	}
}
