package cmd_gateway

import (
	"bytes"
	"encoding/json"
	"errors"
	"path/filepath"
	"strings"
	"testing"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
)

func TestGatewayRoutingRulesCreate_Integration_FromFile(t *testing.T) {
	rule := testGatewayRoutingRule()
	filePath := writeGatewayRoutingRuleFixture(t, rule)

	var received json.RawMessage
	mockGatewayAPI := &api.MockGatewayAPI{
		CreateGatewayRoutingRuleV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			body json.RawMessage,
		) (json.RawMessage, error) {
			if gatewayID != testRoutingRuleGatewayID {
				t.Fatalf("Expected gateway ID %s, got %q", testRoutingRuleGatewayID, gatewayID)
			}
			received = body
			return mustMarshalGatewayRoutingRule(t, rule), nil
		},
	}

	commandOutput, execute := setupGatewayRoutingRulesIntegrationCommand(newGatewayRoutingRulesService(mockGatewayAPI))

	if err := execute([]string{
		"routing-rules", "create",
		"--gateway-id", testRoutingRuleGatewayID,
		"--file-path", filePath,
	}); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedRequest := mustMarshalGatewayRoutingRule(t, rule)
	if !bytes.Equal(received, expectedRequest) {
		t.Fatalf("Expected raw body %s, got %s", expectedRequest, received)
	}

	expectedResult := "Routing rule " + testRoutingRuleID + " created successfully."
	actualResult := strings.TrimSpace(commandOutput.String())
	if actualResult != expectedResult {
		t.Fatalf("Expected output %q, got %q", expectedResult, actualResult)
	}
}

func TestGatewayRoutingRulesCreate_Integration_FromStdin(t *testing.T) {
	rule := testGatewayRoutingRule()
	document := string(mustMarshalGatewayRoutingRule(t, rule))

	mockGatewayAPI := &api.MockGatewayAPI{
		CreateGatewayRoutingRuleV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			body json.RawMessage,
		) (json.RawMessage, error) {
			return mustMarshalGatewayRoutingRule(t, rule), nil
		},
	}

	commandOutput, execute := setupGatewayRoutingRulesIntegrationCommandWithStdin(
		newGatewayRoutingRulesService(mockGatewayAPI),
		document,
	)

	if err := execute([]string{
		"routing-rules", "create",
		"--gateway-id", testRoutingRuleGatewayID,
		"--file-path", "-",
	}); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedResult := "Routing rule " + testRoutingRuleID + " created successfully."
	actualResult := strings.TrimSpace(commandOutput.String())
	if actualResult != expectedResult {
		t.Fatalf("Expected success message, got %q", actualResult)
	}
}

func TestGatewayRoutingRulesCreate_Integration_IgnoresJSONOutput(t *testing.T) {
	rule := testGatewayRoutingRule()

	mockGatewayAPI := &api.MockGatewayAPI{
		CreateGatewayRoutingRuleV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			body json.RawMessage,
		) (json.RawMessage, error) {
			return mustMarshalGatewayRoutingRule(t, rule), nil
		},
	}

	commandOutput, execute := setupGatewayRoutingRulesIntegrationCommand(newGatewayRoutingRulesService(mockGatewayAPI))

	if err := execute([]string{
		"routing-rules", "create",
		"--gateway-id", testRoutingRuleGatewayID,
		"--file-path", writeGatewayRoutingRuleFixture(t, rule),
		"--output", "json",
	}); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// --output json only affects the message envelope: the created document is
	// never printed by create.
	output := commandOutput.String()
	if !strings.Contains(output, "Routing rule "+testRoutingRuleID+" created successfully.") {
		t.Fatalf("Expected the success message, got %q", output)
	}
	if strings.Contains(output, "inputStorageClass") || strings.Contains(output, "redundancyClassId") {
		t.Fatalf("Expected no document body in create output, got %q", output)
	}
}

func TestGatewayRoutingRulesCreate_Integration_MissingFile(t *testing.T) {
	apiCalled := false
	mockGatewayAPI := &api.MockGatewayAPI{
		CreateGatewayRoutingRuleV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			body json.RawMessage,
		) (json.RawMessage, error) {
			apiCalled = true
			return nil, nil
		},
	}

	_, execute := setupGatewayRoutingRulesIntegrationCommand(newGatewayRoutingRulesService(mockGatewayAPI))

	err := execute([]string{
		"routing-rules", "create",
		"--gateway-id", testRoutingRuleGatewayID,
		"--file-path", filepath.Join(t.TempDir(), "missing.json"),
	})
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if !strings.Contains(err.Error(), "failed to read file") {
		t.Fatalf("Expected file read error, got %q", err.Error())
	}
	if apiCalled {
		t.Fatal("CreateGatewayRoutingRuleV5 should not be called when the file cannot be read")
	}
}

func TestGatewayRoutingRulesCreate_Integration_RejectsTrailingData(t *testing.T) {
	rule := testGatewayRoutingRule()
	document := append(mustMarshalGatewayRoutingRule(t, rule), []byte(`{"extra":true}`)...)

	mockGatewayAPI := &api.MockGatewayAPI{
		CreateGatewayRoutingRuleV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			body json.RawMessage,
		) (json.RawMessage, error) {
			t.Fatal("CreateGatewayRoutingRuleV5 should not be called for trailing data")
			return nil, nil
		},
	}

	commandOutput, execute := setupGatewayRoutingRulesIntegrationCommandWithStdin(
		newGatewayRoutingRulesService(mockGatewayAPI),
		string(document),
	)

	err := execute([]string{
		"routing-rules", "create",
		"--gateway-id", testRoutingRuleGatewayID,
		"--file-path", "-",
	})
	if err == nil {
		t.Fatal("Expected a command-level error, got nil")
	}
	if !strings.Contains(commandOutput.String(), "unexpected trailing data") {
		t.Fatalf("Expected trailing data error, got %q", commandOutput.String())
	}
}

func TestGatewayRoutingRulesCreate_Integration_Surfaces422Message(t *testing.T) {
	rule := testGatewayRoutingRule()

	mockGatewayAPI := &api.MockGatewayAPI{
		CreateGatewayRoutingRuleV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			body json.RawMessage,
		) (json.RawMessage, error) {
			return nil, errors.New(
				"code status expected 201, but received 422 instead\n" +
					"validation failed: input storage class \"x\" is not part of the gateway vocabulary",
			)
		},
	}

	commandOutput, execute := setupGatewayRoutingRulesIntegrationCommand(newGatewayRoutingRulesService(mockGatewayAPI))

	err := execute([]string{
		"routing-rules", "create",
		"--gateway-id", testRoutingRuleGatewayID,
		"--file-path", writeGatewayRoutingRuleFixture(t, rule),
	})
	if err == nil {
		t.Fatal("Expected a command-level error, got nil")
	}

	output := commandOutput.String()
	for _, expected := range []string{
		"failed to create routing rule",
		"received 422",
		"input storage class",
		"not part of the gateway vocabulary",
	} {
		if !strings.Contains(output, expected) {
			t.Fatalf("Expected error output to contain %q, got %q", expected, output)
		}
	}
}
