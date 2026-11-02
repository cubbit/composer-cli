package cmd_gateway

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/cubbit/composer-cli/src/api"
	servicegateway "github.com/cubbit/composer-cli/src/service/gateway"
)

const (
	testRoutingRuleID                = "907fcf49-ec63-431b-975d-814919e72a21"
	testRoutingRuleGatewayID         = "gw-001"
	testRoutingRuleOtherID           = "d34c2f25-56b4-4e6a-9e07-0d0d6f4d9b6a"
	testRoutingRuleRedundancyClassID = "33333333-3333-4333-8333-333333333333"
	testRoutingRuleFileName          = "rule.json"
)

func setupGatewayRoutingRulesIntegrationCommand(gatewayService servicegateway.GatewayServiceInterface) (*bytes.Buffer, func([]string) error) {
	return setupGatewayConnectionIntegrationCommand(gatewayService)
}

// setupGatewayRoutingRulesIntegrationCommandWithStdin builds a gateway command
// with a stdin feed, for the create/update flows that read the document from
// stdin when --file-path is "-".
func setupGatewayRoutingRulesIntegrationCommandWithStdin(
	gatewayService servicegateway.GatewayServiceInterface,
	stdin string,
) (*bytes.Buffer, func([]string) error) {
	gatewayCmd := NewGatewayCmd(gatewayService)
	gatewayCmd.PersistentFlags().String("profile", "", "Profile")
	gatewayCmd.PersistentFlags().String("output", "human", "Output format")
	gatewayCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	gatewayCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	commandOutput := new(bytes.Buffer)
	gatewayCmd.SetOut(commandOutput)
	gatewayCmd.SetErr(commandOutput)
	gatewayCmd.SetIn(strings.NewReader(stdin))

	return commandOutput, func(args []string) error {
		gatewayCmd.SetArgs(args)
		return gatewayCmd.Execute()
	}
}

func newGatewayRoutingRulesService(gatewayAPI api.GatewayAPIInterface) servicegateway.GatewayService {
	return newGatewayConnectionService(gatewayAPI)
}

func testGatewayRoutingRule() api.GatewayRoutingRuleV5 {
	return api.GatewayRoutingRuleV5{
		ID:                testRoutingRuleID,
		Name:              "cold-data",
		Description:       "Route cold objects to the cold storage class",
		Enabled:           true,
		InputStorageClass: "COLD",
		Conditions: &api.GatewayRoutingRuleV5Condition{
			Operator: "and",
			Conditions: []*api.GatewayRoutingRuleV5Condition{
				{
					Field:         "object.key",
					Operator:      "starts_with",
					Value:         "logs/",
					CaseSensitive: true,
				},
			},
		},
		Action: api.GatewayRoutingRuleV5Action{
			Type: "route",
			Target: &api.GatewayRoutingRuleV5Target{
				Type:              "fixed",
				RedundancyClassID: testRoutingRuleRedundancyClassID,
			},
		},
	}
}

// testMarshalGatewayRoutingRule serialises a rule the same way the mocked API
// does, so tests can assert on the exact documents the commands print.
func testMarshalGatewayRoutingRule(rule api.GatewayRoutingRuleV5) ([]byte, error) {
	return json.Marshal(rule)
}

func mustMarshalGatewayRoutingRule(t *testing.T, rule api.GatewayRoutingRuleV5) []byte {
	t.Helper()

	document, err := testMarshalGatewayRoutingRule(rule)
	if err != nil {
		t.Fatalf("Expected the rule to marshal, got %v", err)
	}

	return document
}

// writeGatewayRoutingRuleFixture writes the given rule as a policy JSON
// document to a temporary file and returns its path.
func writeGatewayRoutingRuleFixture(t *testing.T, rule api.GatewayRoutingRuleV5) string {
	t.Helper()

	document, err := testMarshalGatewayRoutingRule(rule)
	if err != nil {
		t.Fatalf("Expected the rule to marshal, got %v", err)
	}

	filePath := filepath.Join(t.TempDir(), testRoutingRuleFileName)
	if err := os.WriteFile(filePath, document, 0o600); err != nil {
		t.Fatalf("Expected the fixture to be written, got %v", err)
	}

	return filePath
}

// reindentJSON mirrors the printer's two-space JSON indentation so the
// expected document matches what writeJSON produces.
func reindentJSON(document []byte) (string, error) {
	var buffer bytes.Buffer
	if err := json.Indent(&buffer, document, "", "  "); err != nil {
		return "", err
	}

	return buffer.String() + "\n", nil
}

func rawGatewayRoutingRuleWithUnknownField() json.RawMessage {
	return json.RawMessage(`{
  "id": "` + testRoutingRuleID + `",
  "name": "cold-data",
  "description": "Route cold objects to the cold storage class",
  "enabled": true,
  "inputStorageClass": "COLD",
  "conditions": null,
  "action": {
    "type": "route",
    "target": {
      "type": "fixed",
      "redundancyClassId": "` + testRoutingRuleRedundancyClassID + `"
    }
  },
  "futurePolicy": {
    "revision": 7,
    "labels": ["beta", "sdp"]
  }
}`)
}

// writeGatewayRoutingRuleRawFixture writes a raw JSON document to a temporary
// file and returns its path.
func writeGatewayRoutingRuleRawFixture(t *testing.T, document []byte) string {
	t.Helper()

	filePath := filepath.Join(t.TempDir(), testRoutingRuleFileName)
	if err := os.WriteFile(filePath, document, 0o600); err != nil {
		t.Fatalf("Expected the fixture to be written, got %v", err)
	}

	return filePath
}
