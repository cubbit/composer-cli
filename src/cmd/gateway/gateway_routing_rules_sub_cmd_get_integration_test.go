package cmd_gateway

import (
	"encoding/json"
	"errors"
	"strings"
	"testing"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"gopkg.in/yaml.v3"
)

func TestGatewayRoutingRulesGet_Integration_HumanOutput(t *testing.T) {
	rawRule := rawGatewayRoutingRuleWithUnknownField()

	mockGatewayAPI := &api.MockGatewayAPI{
		GetGatewayRoutingRuleV5RawFunc: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			ruleID string,
		) (json.RawMessage, error) {
			if apiKey != "test-api-key" || organizationID != "test-org-id" {
				t.Fatalf("Expected credentials to be propagated")
			}
			if gatewayID != testRoutingRuleGatewayID {
				t.Fatalf("Expected gateway ID %s, got %q", testRoutingRuleGatewayID, gatewayID)
			}
			if ruleID != testRoutingRuleID {
				t.Fatalf("Expected rule ID %s, got %q", testRoutingRuleID, ruleID)
			}

			return rawRule, nil
		},
	}

	commandOutput, execute := setupGatewayRoutingRulesIntegrationCommand(newGatewayRoutingRulesService(mockGatewayAPI))

	if err := execute([]string{
		"routing-rules", "get",
		"--gateway-id", testRoutingRuleGatewayID,
		"--id", testRoutingRuleID,
	}); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedResult, err := reindentJSON(rawRule)
	if err != nil {
		t.Fatalf("Expected indented JSON, got %v", err)
	}
	actualResult := strings.TrimSpace(commandOutput.String())
	if actualResult != strings.TrimSpace(expectedResult) {
		t.Fatalf("Expected output %q, got %q", expectedResult, actualResult)
	}
}

func TestGatewayRoutingRulesGet_Integration_JSONPreservesUnmodeledField(t *testing.T) {
	rawRule := rawGatewayRoutingRuleWithUnknownField()

	mockGatewayAPI := &api.MockGatewayAPI{
		GetGatewayRoutingRuleV5RawFunc: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			ruleID string,
		) (json.RawMessage, error) {
			return rawRule, nil
		},
	}

	commandOutput, execute := setupGatewayRoutingRulesIntegrationCommand(newGatewayRoutingRulesService(mockGatewayAPI))

	if err := execute([]string{
		"routing-rules", "get",
		"--gateway-id", testRoutingRuleGatewayID,
		"--id", testRoutingRuleID,
		"--output", "json",
	}); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	for _, expected := range []string{`"futurePolicy"`, `"revision": 7`, `"beta"`} {
		if !strings.Contains(commandOutput.String(), expected) {
			t.Fatalf("Expected JSON output to preserve %q, got %q", expected, commandOutput.String())
		}
	}

	expectedOutput, err := reindentJSON(rawRule)
	if err != nil {
		t.Fatalf("Expected indented JSON, got %v", err)
	}
	actualResult := strings.TrimSpace(commandOutput.String())
	if actualResult != strings.TrimSpace(expectedOutput) {
		t.Fatalf("Expected raw passthrough %q, got %q", expectedOutput, actualResult)
	}
}

func TestGatewayRoutingRulesGet_Integration_YAMLOutput(t *testing.T) {
	rawRule := rawGatewayRoutingRuleWithUnknownField()

	mockGatewayAPI := &api.MockGatewayAPI{
		GetGatewayRoutingRuleV5RawFunc: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			ruleID string,
		) (json.RawMessage, error) {
			return rawRule, nil
		},
	}

	commandOutput, execute := setupGatewayRoutingRulesIntegrationCommand(newGatewayRoutingRulesService(mockGatewayAPI))

	if err := execute([]string{
		"routing-rules", "get",
		"--gateway-id", testRoutingRuleGatewayID,
		"--id", testRoutingRuleID,
		"--output", "yaml",
	}); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	var decoded any
	if err := json.Unmarshal(rawRule, &decoded); err != nil {
		t.Fatalf("Expected the rule document to parse, got %v", err)
	}

	expectedOutput, err := yaml.Marshal(decoded)
	if err != nil {
		t.Fatalf("Expected the rule to marshal, got %v", err)
	}

	actualResult := strings.TrimSpace(commandOutput.String())
	expectedResult := strings.TrimSpace(string(expectedOutput))
	if actualResult != expectedResult {
		t.Fatalf("Expected output %q, got %q", expectedResult, actualResult)
	}
}

func TestGatewayRoutingRulesGet_Integration_PropagatesAPIError(t *testing.T) {
	mockGatewayAPI := &api.MockGatewayAPI{
		GetGatewayRoutingRuleV5RawFunc: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			ruleID string,
		) (json.RawMessage, error) {
			return nil, errors.New("code status expected 200, but received 404 instead")
		},
	}

	_, execute := setupGatewayRoutingRulesIntegrationCommand(newGatewayRoutingRulesService(mockGatewayAPI))

	err := execute([]string{
		"routing-rules", "get",
		"--gateway-id", testRoutingRuleGatewayID,
		"--id", testRoutingRuleID,
	})
	if err == nil {
		t.Fatal("Expected a command-level error, got nil")
	}
	if !strings.Contains(err.Error(), "failed to get routing rule") {
		t.Fatalf("Expected wrapped API error, got %q", err.Error())
	}
}
