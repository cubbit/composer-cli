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

func TestGatewayRoutingRulesUpdate_Integration_ForwardsRawUnmodeledField(t *testing.T) {
	rule := testGatewayRoutingRule()
	rule.Conditions = nil
	rawRule := rawGatewayRoutingRuleWithUnknownField()

	var received json.RawMessage
	var requestedRuleID string
	mockGatewayAPI := &api.MockGatewayAPI{
		UpdateGatewayRoutingRuleV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			ruleID string,
			body json.RawMessage,
		) (json.RawMessage, error) {
			if gatewayID != testRoutingRuleGatewayID {
				t.Fatalf("Expected gateway ID %s, got %q", testRoutingRuleGatewayID, gatewayID)
			}
			requestedRuleID = ruleID
			received = body
			return rawRule, nil
		},
	}

	commandOutput, execute := setupGatewayRoutingRulesIntegrationCommand(newGatewayRoutingRulesService(mockGatewayAPI))

	if err := execute([]string{
		"routing-rules", "update",
		"--gateway-id", testRoutingRuleGatewayID,
		"--id", testRoutingRuleID,
		"--file-path", writeGatewayRoutingRuleRawFixture(t, rawRule),
	}); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if requestedRuleID != testRoutingRuleID {
		t.Fatalf("Expected rule ID %s, got %q", testRoutingRuleID, requestedRuleID)
	}
	if !bytes.Equal(received, rawRule) {
		t.Fatalf("Expected raw PUT body.\nwant: %s\ngot:  %s", rawRule, received)
	}
	if !strings.Contains(string(received), `"futurePolicy"`) {
		t.Fatalf("Expected the unmodeled field to be forwarded, got %s", received)
	}

	expectedResult := "Routing rule " + testRoutingRuleID + " updated successfully."
	actualResult := strings.TrimSpace(commandOutput.String())
	if actualResult != expectedResult {
		t.Fatalf("Expected output %q, got %q", expectedResult, actualResult)
	}
}

func TestGatewayRoutingRulesUpdate_Integration_IDMismatch(t *testing.T) {
	rule := testGatewayRoutingRule()
	rule.ID = testRoutingRuleOtherID

	mockGatewayAPI := &api.MockGatewayAPI{
		UpdateGatewayRoutingRuleV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			ruleID string,
			body json.RawMessage,
		) (json.RawMessage, error) {
			t.Fatal("UpdateGatewayRoutingRuleV5 should not be called when ids mismatch")
			return nil, nil
		},
	}

	commandOutput, execute := setupGatewayRoutingRulesIntegrationCommand(newGatewayRoutingRulesService(mockGatewayAPI))

	err := execute([]string{
		"routing-rules", "update",
		"--gateway-id", testRoutingRuleGatewayID,
		"--id", testRoutingRuleID,
		"--file-path", writeGatewayRoutingRuleFixture(t, rule),
	})
	if err == nil {
		t.Fatal("Expected a command-level error, got nil")
	}
	if !strings.Contains(commandOutput.String(), "does not match the input document id") {
		t.Fatalf("Expected id mismatch error, got %q", commandOutput.String())
	}
}

func TestGatewayRoutingRulesUpdate_Integration_UnknownFieldTolerated(t *testing.T) {
	document := []byte(`{
	  "name":"cold-data",
	  "inputStorageClass":"COLD",
	  "action":{"type":"drop"},
	  "futurePolicy":{"revision":7}
	}`)

	var received json.RawMessage
	mockGatewayAPI := &api.MockGatewayAPI{
		UpdateGatewayRoutingRuleV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			ruleID string,
			body json.RawMessage,
		) (json.RawMessage, error) {
			received = body
			return document, nil
		},
	}

	commandOutput, execute := setupGatewayRoutingRulesIntegrationCommandWithStdin(
		newGatewayRoutingRulesService(mockGatewayAPI),
		string(document),
	)

	if err := execute([]string{
		"routing-rules", "update",
		"--gateway-id", testRoutingRuleGatewayID,
		"--id", testRoutingRuleID,
		"--file-path", "-",
	}); err != nil {
		t.Fatalf("Expected unknown fields to be tolerated, got %v", err)
	}

	if !bytes.Equal(received, document) {
		t.Fatalf("Expected the raw document to be forwarded, got %s", received)
	}
	expectedResult := "Routing rule " + testRoutingRuleID + " updated successfully."
	actualResult := strings.TrimSpace(commandOutput.String())
	if actualResult != expectedResult {
		t.Fatalf("Expected output %q, got %q", expectedResult, actualResult)
	}
}

func TestGatewayRoutingRulesUpdate_Integration_MissingFile(t *testing.T) {
	apiCalled := false
	mockGatewayAPI := &api.MockGatewayAPI{
		UpdateGatewayRoutingRuleV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			ruleID string,
			body json.RawMessage,
		) (json.RawMessage, error) {
			apiCalled = true
			return nil, nil
		},
	}

	_, execute := setupGatewayRoutingRulesIntegrationCommand(newGatewayRoutingRulesService(mockGatewayAPI))

	err := execute([]string{
		"routing-rules", "update",
		"--gateway-id", testRoutingRuleGatewayID,
		"--id", testRoutingRuleID,
		"--file-path", filepath.Join(t.TempDir(), "missing.json"),
	})
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if !strings.Contains(err.Error(), "failed to read file") {
		t.Fatalf("Expected file read error, got %q", err.Error())
	}
	if apiCalled {
		t.Fatal("UpdateGatewayRoutingRuleV5 should not be called when the file cannot be read")
	}
}

func TestGatewayRoutingRulesUpdate_Integration_Surfaces422Message(t *testing.T) {
	rule := testGatewayRoutingRule()

	mockGatewayAPI := &api.MockGatewayAPI{
		UpdateGatewayRoutingRuleV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			ruleID string,
			body json.RawMessage,
		) (json.RawMessage, error) {
			return nil, errors.New(
				"code status expected 200, but received 422 instead\n" +
					"validation failed: input storage class \"x\" is not part of the gateway vocabulary",
			)
		},
	}

	commandOutput, execute := setupGatewayRoutingRulesIntegrationCommand(newGatewayRoutingRulesService(mockGatewayAPI))

	err := execute([]string{
		"routing-rules", "update",
		"--gateway-id", testRoutingRuleGatewayID,
		"--id", testRoutingRuleID,
		"--file-path", writeGatewayRoutingRuleFixture(t, rule),
	})
	if err == nil {
		t.Fatal("Expected a command-level error, got nil")
	}

	output := commandOutput.String()
	for _, expected := range []string{
		"failed to update routing rule",
		"received 422",
		"input storage class",
		"not part of the gateway vocabulary",
	} {
		if !strings.Contains(output, expected) {
			t.Fatalf("Expected error output to contain %q, got %q", expected, output)
		}
	}
}
