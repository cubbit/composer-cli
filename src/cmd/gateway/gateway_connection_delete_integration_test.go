package cmd_gateway

import (
	"errors"
	"strings"
	"testing"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
)

func TestGatewayConnectionDelete_Integration_HumanOutput(t *testing.T) {
	var (
		requestedGatewayID  string
		deletedConnectionID string
	)

	mockGatewayAPI := &api.MockGatewayAPI{
		DeleteGatewayConnectionV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			connectionID string,
		) error {
			if gatewayID != "gw-001" {
				t.Fatalf("Expected gateway ID gw-001, got %q", gatewayID)
			}
			requestedGatewayID = gatewayID
			deletedConnectionID = connectionID
			return nil
		},
	}

	commandOutput, execute := setupGatewayConnectionIntegrationCommand(newGatewayConnectionService(mockGatewayAPI))

	if err := execute([]string{
		"connection", "delete",
		"--gateway-id", "gw-001",
		"--id", testConnectionID,
	}); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if requestedGatewayID != "gw-001" {
		t.Fatalf("Expected gateway ID gw-001, got %q", requestedGatewayID)
	}
	if deletedConnectionID != testConnectionID {
		t.Fatalf("Expected connection ID %s, got %q", testConnectionID, deletedConnectionID)
	}

	expectedOutput := "Connection " + testConnectionID + " deleted successfully.\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestGatewayConnectionDelete_Integration_JSONOutput(t *testing.T) {
	commandOutput, execute := setupGatewayConnectionIntegrationCommand(newGatewayConnectionService(&api.MockGatewayAPI{}))

	if err := execute([]string{
		"connection", "delete",
		"--gateway-id", "gw-001",
		"--id", testConnectionID,
		"--output", "json",
	}); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "{\n  \"message\": \"Connection " + testConnectionID + " deleted successfully.\\n\"\n}\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestGatewayConnectionDelete_Integration_YAMLOutput(t *testing.T) {
	commandOutput, execute := setupGatewayConnectionIntegrationCommand(newGatewayConnectionService(&api.MockGatewayAPI{}))

	if err := execute([]string{
		"connection", "delete",
		"--gateway-id", "gw-001",
		"--id", testConnectionID,
		"--output", "yaml",
	}); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "message: |\n    Connection " + testConnectionID + " deleted successfully.\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestGatewayConnectionDelete_Integration_DefaultConflictIsForwarded(t *testing.T) {
	mockGatewayAPI := &api.MockGatewayAPI{
		DeleteGatewayConnectionV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			connectionID string,
		) error {
			return errors.New("code status expected 204, but received 409 instead\nthe default connection cannot be deleted")
		},
	}

	commandOutput, execute := setupGatewayConnectionIntegrationCommand(newGatewayConnectionService(mockGatewayAPI))

	err := execute([]string{
		"connection", "delete",
		"--gateway-id", "gw-001",
		"--id", testConnectionID,
	})
	if err == nil {
		t.Fatal("Expected a command-level error, got nil")
	}

	expected := "failed to delete connection: code status expected 204, but received 409 instead\nthe default connection cannot be deleted"
	if err.Error() != expected {
		t.Fatalf("Expected the API error to be forwarded unchanged, got %q", err.Error())
	}
	if !strings.Contains(commandOutput.String(), "the default connection cannot be deleted") {
		t.Fatalf("Expected default connection conflict error output, got %q", commandOutput.String())
	}
}

func TestGatewayConnectionDelete_Integration_PropagatesNotFound(t *testing.T) {
	mockGatewayAPI := &api.MockGatewayAPI{
		DeleteGatewayConnectionV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			connectionID string,
		) error {
			return errors.New("code status expected 204, but received 404 instead\nconnection not found")
		},
	}

	commandOutput, execute := setupGatewayConnectionIntegrationCommand(newGatewayConnectionService(mockGatewayAPI))

	err := execute([]string{
		"connection", "delete",
		"--gateway-id", "gw-001",
		"--id", testConnectionID,
	})
	if err == nil {
		t.Fatal("Expected a command-level error, got nil")
	}

	if !strings.Contains(err.Error(), "failed to delete connection") ||
		!strings.Contains(err.Error(), "received 404") {
		t.Fatalf("Expected not-found error, got %q", err.Error())
	}
	if !strings.Contains(commandOutput.String(), "Error: failed to delete connection") {
		t.Fatalf("Expected failure output, got %q", commandOutput.String())
	}
	if strings.Contains(commandOutput.String(), "deleted successfully") {
		t.Fatalf("Expected no success output on failure, got %q", commandOutput.String())
	}
}
