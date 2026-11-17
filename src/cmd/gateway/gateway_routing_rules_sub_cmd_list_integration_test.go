package cmd_gateway

import (
	"errors"
	"strings"
	"testing"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
)

func TestGatewayRoutingRulesList_Integration_EmptyState(t *testing.T) {
	mockGatewayAPI := &api.MockGatewayAPI{
		ListGatewayRoutingRulesV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
		) (*api.GatewayRoutingRulesV5ListResponse, error) {
			if gatewayID != testRoutingRuleGatewayID {
				t.Fatalf("Expected gateway ID %s, got %q", testRoutingRuleGatewayID, gatewayID)
			}
			return &api.GatewayRoutingRulesV5ListResponse{}, nil
		},
	}

	commandOutput, execute := setupGatewayRoutingRulesIntegrationCommand(newGatewayRoutingRulesService(mockGatewayAPI))

	if err := execute([]string{"routing-rules", "list", "--gateway-id", testRoutingRuleGatewayID}); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if commandOutput.String() != "No routing rules found.\n" {
		t.Fatalf("Expected empty state message, got %q", commandOutput.String())
	}
}

func TestGatewayRoutingRulesList_Integration_MetadataTable(t *testing.T) {
	first := testGatewayRoutingRule()
	second := first
	second.ID = testRoutingRuleOtherID
	second.Name = "hot-data"
	second.Enabled = false
	second.InputStorageClass = "HOT"
	second.Description = ""

	mockGatewayAPI := &api.MockGatewayAPI{
		ListGatewayRoutingRulesV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
		) (*api.GatewayRoutingRulesV5ListResponse, error) {
			return &api.GatewayRoutingRulesV5ListResponse{first, second}, nil
		},
	}

	commandOutput, execute := setupGatewayRoutingRulesIntegrationCommand(newGatewayRoutingRulesService(mockGatewayAPI))

	if err := execute([]string{"routing-rules", "list", "--gateway-id", testRoutingRuleGatewayID}); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	output := commandOutput.String()
	for _, expected := range []string{
		"ID",
		"Name",
		"Enabled",
		"Input Storage Class",
		"Description",
		first.ID,
		"cold-data",
		second.ID,
		"hot-data",
	} {
		if !strings.Contains(output, expected) {
			t.Fatalf("Expected list output to contain %q, got %q", expected, output)
		}
	}
	if strings.Contains(output, "Conditions") {
		t.Fatalf("Expected metadata-only list output, got %q", output)
	}
}

func TestGatewayRoutingRulesList_Integration_NoHeaders(t *testing.T) {
	rule := testGatewayRoutingRule()

	mockGatewayAPI := &api.MockGatewayAPI{
		ListGatewayRoutingRulesV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
		) (*api.GatewayRoutingRulesV5ListResponse, error) {
			return &api.GatewayRoutingRulesV5ListResponse{rule}, nil
		},
	}

	commandOutput, execute := setupGatewayRoutingRulesIntegrationCommand(newGatewayRoutingRulesService(mockGatewayAPI))

	if err := execute([]string{"routing-rules", "list", "--gateway-id", testRoutingRuleGatewayID, "--no-headers"}); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedResult := strings.TrimSpace(`
╭──────────────────────────────────────┬───────────┬─────┬──────┬──────────────────────────────────────────────╮
│ 907fcf49-ec63-431b-975d-814919e72a21 │ cold-data │ yes │ COLD │ Route cold objects to the cold storage class │
╰──────────────────────────────────────┴───────────┴─────┴──────┴──────────────────────────────────────────────╯
`)

	actualResult := strings.TrimSpace(commandOutput.String())
	if actualResult != expectedResult {
		t.Errorf("Expected list output without headers does not match actual output\nExpected:\n%s\nActual:\n%s", expectedResult, actualResult)
	}
}

func TestGatewayRoutingRulesList_Integration_JSONOutput(t *testing.T) {
	rule := testGatewayRoutingRule()

	mockGatewayAPI := &api.MockGatewayAPI{
		ListGatewayRoutingRulesV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
		) (*api.GatewayRoutingRulesV5ListResponse, error) {
			return &api.GatewayRoutingRulesV5ListResponse{rule}, nil
		},
	}

	commandOutput, execute := setupGatewayRoutingRulesIntegrationCommand(newGatewayRoutingRulesService(mockGatewayAPI))

	if err := execute([]string{"routing-rules", "list", "--gateway-id", testRoutingRuleGatewayID, "--output", "json"}); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	encoded, err := testMarshalGatewayRoutingRule(rule)
	if err != nil {
		t.Fatalf("Expected the rule to marshal, got %v", err)
	}
	expectedOutput, err := reindentJSON([]byte("[" + string(encoded) + "]"))
	if err != nil {
		t.Fatalf("Expected indented JSON, got %v", err)
	}
	actualResult := strings.TrimSpace(commandOutput.String())
	if actualResult != strings.TrimSpace(expectedOutput) {
		t.Fatalf("Expected output %q, got %q", expectedOutput, actualResult)
	}
}

func TestGatewayRoutingRulesList_Integration_PropagatesAPIError(t *testing.T) {
	mockGatewayAPI := &api.MockGatewayAPI{
		ListGatewayRoutingRulesV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
		) (*api.GatewayRoutingRulesV5ListResponse, error) {
			return nil, errors.New("code status expected 200, but received 500 instead")
		},
	}

	_, execute := setupGatewayRoutingRulesIntegrationCommand(newGatewayRoutingRulesService(mockGatewayAPI))

	err := execute([]string{"routing-rules", "list", "--gateway-id", testRoutingRuleGatewayID})
	if err == nil {
		t.Fatal("Expected a command-level error, got nil")
	}
	if !strings.Contains(err.Error(), "failed to list routing rules") {
		t.Fatalf("Expected wrapped API error, got %q", err.Error())
	}
}
