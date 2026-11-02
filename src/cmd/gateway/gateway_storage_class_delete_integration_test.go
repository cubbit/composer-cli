package cmd_gateway

import (
	"errors"
	"strings"
	"testing"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
)

func TestGatewayStorageClassDelete_Integration_HumanOutput(t *testing.T) {
	var deletedStorageClassID string

	mockGatewayAPI := &api.MockGatewayAPI{
		DeleteGatewayStorageClassV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			storageClassID string,
		) error {
			if gatewayID != "gw-001" {
				t.Fatalf("Expected gateway ID gw-001, got %q", gatewayID)
			}
			deletedStorageClassID = storageClassID
			return nil
		},
	}

	commandOutput, execute := setupGatewayStorageClassIntegrationCommand(newGatewayStorageClassService(mockGatewayAPI))

	if err := execute([]string{"storage-class", "delete", "--gateway-id", "gw-001", "--id", "sc-001"}); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if deletedStorageClassID != "sc-001" {
		t.Fatalf("Expected storage class ID sc-001, got %q", deletedStorageClassID)
	}

	expectedOutput := "Storage class sc-001 deleted successfully.\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestGatewayStorageClassDelete_Integration_JSONOutput(t *testing.T) {
	commandOutput, execute := setupGatewayStorageClassIntegrationCommand(newGatewayStorageClassService(&api.MockGatewayAPI{}))

	if err := execute([]string{"storage-class", "delete", "--gateway-id", "gw-001", "--id", "sc-001", "--output", "json"}); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := `{
  "message": "Storage class sc-001 deleted successfully.\n"
}
`
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestGatewayStorageClassDelete_Integration_YAMLOutput(t *testing.T) {
	commandOutput, execute := setupGatewayStorageClassIntegrationCommand(newGatewayStorageClassService(&api.MockGatewayAPI{}))

	if err := execute([]string{"storage-class", "delete", "--gateway-id", "gw-001", "--id", "sc-001", "--output", "yaml"}); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "message: |\n    Storage class sc-001 deleted successfully.\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestGatewayStorageClassDelete_Integration_PropagatesNotFound(t *testing.T) {
	mockGatewayAPI := &api.MockGatewayAPI{
		DeleteGatewayStorageClassV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			storageClassID string,
		) error {
			return errors.New("code status expected 204, but received 404 instead")
		},
	}

	commandOutput, execute := setupGatewayStorageClassIntegrationCommand(newGatewayStorageClassService(mockGatewayAPI))

	err := execute([]string{"storage-class", "delete", "--gateway-id", "gw-001", "--id", "sc-001"})
	if err == nil {
		t.Fatal("Expected a command-level error, got nil")
	}

	if !containsAll(err.Error(), "failed to delete storage class", "received 404") {
		t.Fatalf("Expected not-found error, got %q", err.Error())
	}
	if !containsAll(commandOutput.String(), "Error: failed to delete storage class", "received 404") {
		t.Fatalf("Expected failure output, got %q", commandOutput.String())
	}
	if strings.Contains(commandOutput.String(), "deleted successfully") {
		t.Fatalf("Expected no success output on failure, got %q", commandOutput.String())
	}
}
