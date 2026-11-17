package cmd_gateway

import (
	"errors"
	"strings"
	"testing"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
)

func TestGatewayRoutingRulesDelete_Integration_Success(t *testing.T) {
	var requestedRuleID string
	mockGatewayAPI := &api.MockGatewayAPI{
		DeleteGatewayRoutingRuleV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			ruleID string,
		) error {
			if gatewayID != testRoutingRuleGatewayID {
				t.Fatalf("Expected gateway ID %s, got %q", testRoutingRuleGatewayID, gatewayID)
			}
			requestedRuleID = ruleID
			return nil
		},
	}

	commandOutput, execute := setupGatewayRoutingRulesIntegrationCommand(newGatewayRoutingRulesService(mockGatewayAPI))

	if err := execute([]string{
		"routing-rules", "delete",
		"--gateway-id", testRoutingRuleGatewayID,
		"--id", testRoutingRuleID,
	}); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if requestedRuleID != testRoutingRuleID {
		t.Fatalf("Expected rule ID %s, got %q", testRoutingRuleID, requestedRuleID)
	}
	expectedResult := "Routing rule " + testRoutingRuleID + " deleted successfully."
	actualResult := strings.TrimSpace(commandOutput.String())
	if actualResult != expectedResult {
		t.Fatalf("Expected delete message, got %q", actualResult)
	}
}

func TestGatewayRoutingRulesDelete_Integration_PropagatesAPIError(t *testing.T) {
	mockGatewayAPI := &api.MockGatewayAPI{
		DeleteGatewayRoutingRuleV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			ruleID string,
		) error {
			return errors.New("code status expected 204, but received 404 instead")
		},
	}

	_, execute := setupGatewayRoutingRulesIntegrationCommand(newGatewayRoutingRulesService(mockGatewayAPI))

	err := execute([]string{
		"routing-rules", "delete",
		"--gateway-id", testRoutingRuleGatewayID,
		"--id", testRoutingRuleID,
	})
	if err == nil {
		t.Fatal("Expected a command-level error, got nil")
	}
	if !strings.Contains(err.Error(), "failed to delete routing rule") {
		t.Fatalf("Expected wrapped API error, got %q", err.Error())
	}
}
