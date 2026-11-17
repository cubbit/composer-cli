package routing_rules

import (
	"bytes"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	api "github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/spf13/cobra"
)

const (
	testRoutingRuleGatewayID = "gateway-123"
	testRoutingRuleID        = "907fcf49-ec63-431b-975d-814919e72a21"
	testRoutingRuleOtherID   = "d34c2f25-56b4-4e6a-9e07-0d0d6f4d9b6a"
	testRoutingRuleRCID      = "33333333-3333-4333-8333-333333333333"
	testRoutingRuleOtherRCID = "44444444-4444-4444-8444-444444444444"
)

func setupRoutingRulesTestCommand() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.SetOut(new(bytes.Buffer))
	cmd.SetErr(new(bytes.Buffer))
	cmd.SetIn(strings.NewReader(""))
	cmd.Flags().String("gateway-id", "", "Gateway ID")
	_ = cmd.Flags().Set("gateway-id", testRoutingRuleGatewayID)
	cmd.Flags().String("output", "human", "Output format")
	cmd.Flags().Bool("quiet", false, "Quiet mode")
	cmd.Flags().Bool("no-headers", false, "No headers")
	cmd.Flags().String("id", "", "Routing rule ID")
	cmd.Flags().String("file-path", "", "Input file")
	return cmd
}

func routingRulesTestProfile() configuration_models.ProfileV2 {
	return configuration_models.ProfileV2{
		Output:         configuration_models.OutputHuman,
		APIKey:         "test-api-key",
		OrganizationID: "test-org-id",
	}
}

func testRoutingRule() api.GatewayRoutingRuleV5 {
	return api.GatewayRoutingRuleV5{
		ID:                testRoutingRuleID,
		Name:              "cold-data",
		Description:       "Route cold objects to the cold storage class",
		Enabled:           true,
		InputStorageClass: "COLD",
		Action: api.GatewayRoutingRuleV5Action{
			Type: "route",
			Target: &api.GatewayRoutingRuleV5Target{
				Type:              "fixed",
				RedundancyClassID: testRoutingRuleRCID,
			},
		},
	}
}

func testRoutingRuleWithConditions() api.GatewayRoutingRuleV5 {
	rule := testRoutingRule()
	rule.Conditions = &api.GatewayRoutingRuleV5Condition{
		Operator: "and",
		Conditions: []*api.GatewayRoutingRuleV5Condition{
			{
				Field:         "object.key",
				Operator:      "starts_with",
				Value:         "logs/",
				CaseSensitive: true,
			},
			{
				Field:    "object.size",
				Operator: "gt",
				Value:    float64(1073741824),
			},
		},
	}

	return rule
}

func mustMarshalRule(t *testing.T, rule api.GatewayRoutingRuleV5) []byte {
	t.Helper()

	document, err := json.Marshal(rule)
	if err != nil {
		t.Fatalf("expected the rule to marshal, got %v", err)
	}

	return document
}

func outputOf(cmd *cobra.Command) string {
	return cmd.OutOrStdout().(*bytes.Buffer).String()
}

func indentRawRule(t *testing.T, raw json.RawMessage) string {
	t.Helper()

	var buffer bytes.Buffer
	if err := json.Indent(&buffer, raw, "", "  "); err != nil {
		t.Fatalf("Expected the raw rule to indent, got %v", err)
	}

	return buffer.String() + "\n"
}

func testRawRoutingRuleWithUnknownField() json.RawMessage {
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
      "redundancyClassId": "` + testRoutingRuleRCID + `"
    }
  },
  "futurePolicy": {
    "revision": 7,
    "labels": ["beta", "sdp"]
  }
}`)
}

func TestList_PrintsMetadataTable(t *testing.T) {
	rule := testRoutingRule()

	mockGatewayAPI := &api.MockGatewayAPI{
		ListGatewayRoutingRulesV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
		) (*api.GatewayRoutingRulesV5ListResponse, error) {
			if apiKey != "test-api-key" || organizationID != "test-org-id" {
				t.Fatalf("Expected credentials to be propagated")
			}
			if gatewayID != testRoutingRuleGatewayID {
				t.Fatalf("Expected gateway ID %s, got %q", testRoutingRuleGatewayID, gatewayID)
			}

			return &api.GatewayRoutingRulesV5ListResponse{rule}, nil
		},
	}

	cmd := setupRoutingRulesTestCommand()

	if err := List(Dependencies{GatewayAPI: mockGatewayAPI}, cmd, nil, routingRulesTestProfile(), nil); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedResult := strings.TrimSpace(`
╭──────────────────────────────────────┬───────────┬─────────┬─────────────────────┬──────────────────────────────────────────────╮
│ ID                                   │ Name      │ Enabled │ Input Storage Class │ Description                                  │
├──────────────────────────────────────┼───────────┼─────────┼─────────────────────┼──────────────────────────────────────────────┤
│ 907fcf49-ec63-431b-975d-814919e72a21 │ cold-data │ yes     │ COLD                │ Route cold objects to the cold storage class │
╰──────────────────────────────────────┴───────────┴─────────┴─────────────────────┴──────────────────────────────────────────────╯
`)

	actualResult := strings.TrimSpace(outputOf(cmd))
	if actualResult != expectedResult {
		t.Errorf("Expected list output does not match actual output\nExpected:\n%s\nActual:\n%s", expectedResult, actualResult)
	}
}

func TestList_EmptyState(t *testing.T) {
	mockGatewayAPI := &api.MockGatewayAPI{
		ListGatewayRoutingRulesV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
		) (*api.GatewayRoutingRulesV5ListResponse, error) {
			return &api.GatewayRoutingRulesV5ListResponse{}, nil
		},
	}

	cmd := setupRoutingRulesTestCommand()

	if err := List(Dependencies{GatewayAPI: mockGatewayAPI}, cmd, nil, routingRulesTestProfile(), nil); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedResult := "No routing rules found."
	actualResult := strings.TrimSpace(outputOf(cmd))
	if actualResult != expectedResult {
		t.Fatalf("Expected empty state message, got %q", actualResult)
	}
}

func TestList_JSONOutput(t *testing.T) {
	rule := testRoutingRule()

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

	cmd := setupRoutingRulesTestCommand()
	_ = cmd.Flags().Set("output", "json")

	if err := List(Dependencies{GatewayAPI: mockGatewayAPI}, cmd, nil, routingRulesTestProfile(), nil); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	var decoded []api.GatewayRoutingRuleV5
	if err := json.Unmarshal([]byte(outputOf(cmd)), &decoded); err != nil {
		t.Fatalf("Expected JSON output to parse, got %v", err)
	}
	if len(decoded) != 1 || decoded[0].ID != rule.ID {
		t.Fatalf("Expected the rule in JSON output, got %+v", decoded)
	}
}

func TestGet_PrintsFullDocumentJSON(t *testing.T) {
	rawRule := testRawRoutingRuleWithUnknownField()

	mockGatewayAPI := &api.MockGatewayAPI{
		GetGatewayRoutingRuleV5RawFunc: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			ruleID string,
		) (json.RawMessage, error) {
			if gatewayID != testRoutingRuleGatewayID {
				t.Fatalf("Expected gateway ID %s, got %q", testRoutingRuleGatewayID, gatewayID)
			}
			if ruleID != testRoutingRuleID {
				t.Fatalf("Expected rule ID %s, got %q", testRoutingRuleID, ruleID)
			}
			return rawRule, nil
		},
	}

	cmd := setupRoutingRulesTestCommand()
	_ = cmd.Flags().Set("id", testRoutingRuleID)

	if err := Get(Dependencies{GatewayAPI: mockGatewayAPI}, cmd, nil, routingRulesTestProfile(), nil); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedResult := strings.TrimSpace(indentRawRule(t, rawRule))
	actualResult := strings.TrimSpace(outputOf(cmd))
	if actualResult != expectedResult {
		t.Fatalf("Expected full document %q, got %q", expectedResult, actualResult)
	}
	if !strings.Contains(outputOf(cmd), `"futurePolicy"`) {
		t.Fatalf("Expected unmodeled field in output, got %q", outputOf(cmd))
	}
}

func TestGet_ByIDCanonicalizesUUID(t *testing.T) {
	var requestedRuleID string
	rawRule := mustMarshalRule(t, testRoutingRule())

	mockGatewayAPI := &api.MockGatewayAPI{
		GetGatewayRoutingRuleV5RawFunc: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			ruleID string,
		) (json.RawMessage, error) {
			requestedRuleID = ruleID
			return rawRule, nil
		},
	}

	cmd := setupRoutingRulesTestCommand()
	_ = cmd.Flags().Set("id", strings.ToUpper(testRoutingRuleID))

	if err := Get(Dependencies{GatewayAPI: mockGatewayAPI}, cmd, nil, routingRulesTestProfile(), nil); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if requestedRuleID != testRoutingRuleID {
		t.Fatalf("Expected canonical rule ID %s, got %q", testRoutingRuleID, requestedRuleID)
	}
}

func TestGet_JSONOutputPreservesUnknownServerFields(t *testing.T) {
	rawRule := testRawRoutingRuleWithUnknownField()

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

	cmd := setupRoutingRulesTestCommand()
	_ = cmd.Flags().Set("id", testRoutingRuleID)
	_ = cmd.Flags().Set("output", "json")

	if err := Get(Dependencies{GatewayAPI: mockGatewayAPI}, cmd, nil, routingRulesTestProfile(), nil); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	output := outputOf(cmd)
	for _, expected := range []string{`"futurePolicy"`, `"revision": 7`, `"beta"`} {
		if !strings.Contains(output, expected) {
			t.Fatalf("Expected JSON output to preserve %q, got %q", expected, output)
		}
	}
}

func TestGet_YAMLOutputPreservesUnknownServerFields(t *testing.T) {
	rawRule := testRawRoutingRuleWithUnknownField()

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

	cmd := setupRoutingRulesTestCommand()
	_ = cmd.Flags().Set("id", testRoutingRuleID)
	_ = cmd.Flags().Set("output", "yaml")

	if err := Get(Dependencies{GatewayAPI: mockGatewayAPI}, cmd, nil, routingRulesTestProfile(), nil); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	output := outputOf(cmd)
	for _, expected := range []string{"futurePolicy:", "revision: 7", "- beta", "- sdp"} {
		if !strings.Contains(output, expected) {
			t.Fatalf("Expected YAML output to preserve %q, got %q", expected, output)
		}
	}
}

func TestGet_YAMLOutputPreservesLargeIntegers(t *testing.T) {
	// 2^53+1 is the first integer that float64 cannot represent exactly; a
	// naive any-decode would round it to 2^53.
	rawRule := json.RawMessage(`{
  "id": "` + testRoutingRuleID + `",
  "name": "cold-data",
  "description": "",
  "enabled": true,
  "inputStorageClass": "COLD",
  "action": {"type": "drop"},
  "futurePolicy": {
    "bigSigned": 9007199254740993,
    "bigUnsigned": 18446744073709551615,
    "ratio": 1.5,
    "revision": 7
  }
}`)

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

	cmd := setupRoutingRulesTestCommand()
	_ = cmd.Flags().Set("id", testRoutingRuleID)
	_ = cmd.Flags().Set("output", "yaml")

	if err := Get(Dependencies{GatewayAPI: mockGatewayAPI}, cmd, nil, routingRulesTestProfile(), nil); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	for _, expected := range []string{
		"bigSigned: 9007199254740993",
		"bigUnsigned: 18446744073709551615",
		"ratio: 1.5",
		"revision: 7",
	} {
		if !strings.Contains(outputOf(cmd), expected) {
			t.Fatalf("Expected YAML output to preserve %q, got %q", expected, outputOf(cmd))
		}
	}
}

func TestGet_JSONOutputByteEqualityNestedVariants(t *testing.T) {
	tests := []struct {
		name  string
		apply func(*api.GatewayRoutingRuleV5)
	}{
		{
			name: "field mapping target",
			apply: func(rule *api.GatewayRoutingRuleV5) {
				rule.Action.Target = &api.GatewayRoutingRuleV5Target{
					Type: "fieldMapping",
					FieldMapping: &api.GatewayRoutingRuleV5FieldMapping{
						Field: "object.tags.storageTier",
						Mapping: map[string]string{
							"archive": testRoutingRuleRCID,
							"cold":    testRoutingRuleOtherRCID,
						},
						DefaultRedundancyClassID: testRoutingRuleRCID,
					},
				}
			},
		},
		{
			name: "weighted load balancing target",
			apply: func(rule *api.GatewayRoutingRuleV5) {
				rule.Action.Target = &api.GatewayRoutingRuleV5Target{
					Type: "loadBalancing",
					LoadBalancing: &api.GatewayRoutingRuleV5LoadBalancing{
						RedundancyClassIDs: []string{testRoutingRuleRCID, testRoutingRuleOtherRCID},
						Algorithm:          "weighted",
						Weights: map[string]float32{
							testRoutingRuleRCID:      0.7,
							testRoutingRuleOtherRCID: 0.3,
						},
					},
				}
			},
		},
		{
			name: "multiple urls offloader target",
			apply: func(rule *api.GatewayRoutingRuleV5) {
				rule.Action.OffloaderTarget = &api.GatewayRoutingRuleV5OffloaderTarget{
					URLs:      []string{"https://offload-a.example.com", "https://offload-b.example.com"},
					Algorithm: "round_robin",
				}
			},
		},
		{
			name: "nested or in exists conditions",
			apply: func(rule *api.GatewayRoutingRuleV5) {
				rule.Conditions = &api.GatewayRoutingRuleV5Condition{
					Operator: "or",
					Conditions: []*api.GatewayRoutingRuleV5Condition{
						{
							Field:    "object.key",
							Operator: "in",
							Value:    []any{"logs/a", "logs/b"},
						},
						{
							Operator: "and",
							Conditions: []*api.GatewayRoutingRuleV5Condition{
								{
									Field:    "object.tags.retention",
									Operator: "exists",
								},
							},
						},
					},
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rule := testRoutingRule()
			test.apply(&rule)
			rawRule := mustMarshalRule(t, rule)

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

			cmd := setupRoutingRulesTestCommand()
			_ = cmd.Flags().Set("id", testRoutingRuleID)
			_ = cmd.Flags().Set("output", "json")

			if err := Get(Dependencies{GatewayAPI: mockGatewayAPI}, cmd, nil, routingRulesTestProfile(), nil); err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}

			var buffer bytes.Buffer
			if err := json.Indent(&buffer, rawRule, "", "  "); err != nil {
				t.Fatalf("Expected the rule to indent, got %v", err)
			}
			expectedResult := strings.TrimSpace(buffer.String())
			actualResult := strings.TrimSpace(outputOf(cmd))
			if actualResult != expectedResult {
				t.Fatalf("Expected raw passthrough %q, got %q", expectedResult, actualResult)
			}
		})
	}
}

func TestCreate_ForwardsRawDocument(t *testing.T) {
	rule := testRoutingRule()
	document, err := json.Marshal(rule)
	if err != nil {
		t.Fatalf("Expected the rule to marshal, got %v", err)
	}
	// Inject an unmodeled field: it must be forwarded verbatim.
	document = []byte(strings.TrimSuffix(string(document), "}") + `,"futurePolicy":{"revision":7}}`)

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
			return document, nil
		},
	}

	cmd := setupRoutingRulesTestCommand()
	_ = cmd.Flags().Set("file-path", "-")
	cmd.SetIn(bytes.NewReader(document))

	if err := Create(Dependencies{GatewayAPI: mockGatewayAPI}, cmd, nil, routingRulesTestProfile(), nil); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !bytes.Equal(received, document) {
		t.Fatalf("Expected the raw document to be forwarded.\nwant: %s\ngot:  %s", document, received)
	}
	expectedResult := "Routing rule " + testRoutingRuleID + " created successfully."
	actualResult := strings.TrimSpace(outputOf(cmd))
	if actualResult != expectedResult {
		t.Fatalf("Expected output %q, got %q", expectedResult, actualResult)
	}
}

func TestCreate_UnknownFieldTolerated(t *testing.T) {
	rule := testRoutingRule()
	document := []byte(`{"name":"cold-data","inputStorageClass":"COLD","action":{"type":"drop"},"futurePolicy":"x"}`)

	var received json.RawMessage
	mockGatewayAPI := &api.MockGatewayAPI{
		CreateGatewayRoutingRuleV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			body json.RawMessage,
		) (json.RawMessage, error) {
			received = body
			return mustMarshalRule(t, rule), nil
		},
	}

	cmd := setupRoutingRulesTestCommand()
	_ = cmd.Flags().Set("file-path", "-")
	cmd.SetIn(bytes.NewReader(document))

	if err := Create(Dependencies{GatewayAPI: mockGatewayAPI}, cmd, nil, routingRulesTestProfile(), nil); err != nil {
		t.Fatalf("Expected unknown fields to be tolerated, got %v", err)
	}
	if !bytes.Equal(received, document) {
		t.Fatalf("Expected the raw document to be forwarded, got %s", received)
	}
}

func TestCreate_ReadsFilePath(t *testing.T) {
	rule := testRoutingRule()
	filePath := filepath.Join(t.TempDir(), "rule.json")
	if err := os.WriteFile(filePath, mustMarshalRule(t, rule), 0o600); err != nil {
		t.Fatalf("Expected the fixture to be written, got %v", err)
	}

	var received json.RawMessage
	mockGatewayAPI := &api.MockGatewayAPI{
		CreateGatewayRoutingRuleV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			body json.RawMessage,
		) (json.RawMessage, error) {
			received = body
			return mustMarshalRule(t, rule), nil
		},
	}

	cmd := setupRoutingRulesTestCommand()
	_ = cmd.Flags().Set("file-path", filePath)

	if err := Create(Dependencies{GatewayAPI: mockGatewayAPI}, cmd, nil, routingRulesTestProfile(), nil); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedRequest, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Expected the fixture to be readable, got %v", err)
	}
	if !bytes.Equal(received, expectedRequest) {
		t.Fatalf("Expected the file contents to be forwarded.\nwant: %s\ngot:  %s", expectedRequest, received)
	}
}

func TestCreate_ReportsFileReadError(t *testing.T) {
	cmd := setupRoutingRulesTestCommand()
	_ = cmd.Flags().Set("file-path", filepath.Join(t.TempDir(), "missing.json"))

	err := Create(Dependencies{GatewayAPI: &api.MockGatewayAPI{}}, cmd, nil, routingRulesTestProfile(), nil)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if !strings.Contains(err.Error(), "failed to read file") {
		t.Fatalf("Expected file read error, got %q", err.Error())
	}
}

func TestCreate_ValidatesInputBeforeAPICall(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectError string
	}{
		{
			name:        "empty input",
			input:       "  \n",
			expectError: "routing rule input is empty",
		},
		{
			name:        "invalid json",
			input:       "{",
			expectError: "failed to parse routing rule JSON",
		},
		{
			name:        "not an object",
			input:       `[{"name":"cold"}]`,
			expectError: "document must be a JSON object",
		},
		{
			name:        "missing name",
			input:       `{"inputStorageClass":"COLD","action":{"type":"drop"}}`,
			expectError: "routing rule name is required",
		},
		{
			name:        "missing input storage class",
			input:       `{"name":"cold","action":{"type":"drop"}}`,
			expectError: "routing rule inputStorageClass is required",
		},
		{
			name:        "missing action type",
			input:       `{"name":"cold","inputStorageClass":"COLD","action":{}}`,
			expectError: "routing rule action.type is required",
		},
		{
			name:        "trailing json document",
			input:       `{"name":"cold","inputStorageClass":"COLD","action":{"type":"drop"}}{"name":"hot"}`,
			expectError: "unexpected trailing data",
		},
		{
			name:        "trailing garbage",
			input:       `{"name":"cold","inputStorageClass":"COLD","action":{"type":"drop"}} garbage`,
			expectError: "unexpected trailing data",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			createCalled := false

			mockGatewayAPI := &api.MockGatewayAPI{
				CreateGatewayRoutingRuleV5Func: func(
					endpoints configuration_models.EndpointsV2,
					apiKey string,
					organizationID string,
					gatewayID string,
					body json.RawMessage,
				) (json.RawMessage, error) {
					createCalled = true
					return json.RawMessage("{}"), nil
				},
			}

			cmd := setupRoutingRulesTestCommand()
			_ = cmd.Flags().Set("file-path", "-")
			cmd.SetIn(strings.NewReader(test.input))

			err := Create(Dependencies{GatewayAPI: mockGatewayAPI}, cmd, nil, routingRulesTestProfile(), nil)
			if err == nil {
				t.Fatal("Expected error, got nil")
			}
			if !strings.Contains(err.Error(), test.expectError) {
				t.Fatalf("Expected error containing %q, got %q", test.expectError, err.Error())
			}
			if createCalled {
				t.Fatal("CreateGatewayRoutingRuleV5 should not be called when input validation fails")
			}
		})
	}
}

func TestUpdate_ForwardsRawDocumentAndUnmodeledField(t *testing.T) {
	rule := testRoutingRule()
	document, err := json.Marshal(rule)
	if err != nil {
		t.Fatalf("Expected the rule to marshal, got %v", err)
	}
	document = []byte(strings.TrimSuffix(string(document), "}") + `,"futurePolicy":{"revision":7}}`)

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
			requestedRuleID = ruleID
			received = body
			return document, nil
		},
	}

	cmd := setupRoutingRulesTestCommand()
	_ = cmd.Flags().Set("id", testRoutingRuleID)
	_ = cmd.Flags().Set("file-path", "-")
	cmd.SetIn(bytes.NewReader(document))

	if err := Update(Dependencies{GatewayAPI: mockGatewayAPI}, cmd, nil, routingRulesTestProfile(), nil); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if requestedRuleID != testRoutingRuleID {
		t.Fatalf("Expected rule ID %s, got %q", testRoutingRuleID, requestedRuleID)
	}
	if !bytes.Equal(received, document) {
		t.Fatalf("Expected the raw document to be forwarded.\nwant: %s\ngot:  %s", document, received)
	}
	if !strings.Contains(string(received), `"futurePolicy"`) {
		t.Fatalf("Expected the unmodeled field to survive, got %s", received)
	}
	expectedResult := "Routing rule " + testRoutingRuleID + " updated successfully."
	actualResult := strings.TrimSpace(outputOf(cmd))
	if actualResult != expectedResult {
		t.Fatalf("Expected output %q, got %q", expectedResult, actualResult)
	}
}

func TestUpdate_ResolvesRuleIDPrecedence(t *testing.T) {
	tests := []struct {
		name        string
		documentID  string
		flagID      string
		expectedID  string
		expectError string
	}{
		{
			name:       "matching document and flag id",
			documentID: testRoutingRuleID,
			flagID:     testRoutingRuleID,
			expectedID: testRoutingRuleID,
		},
		{
			name:       "matching ids with different case",
			documentID: testRoutingRuleID,
			flagID:     strings.ToUpper(testRoutingRuleID),
			expectedID: testRoutingRuleID,
		},
		{
			name:       "flag id without document id",
			flagID:     testRoutingRuleID,
			expectedID: testRoutingRuleID,
		},
		{
			name:        "mismatched ids",
			documentID:  testRoutingRuleID,
			flagID:      testRoutingRuleOtherID,
			expectError: "does not match the input document id",
		},
		{
			name:        "invalid flag id",
			flagID:      "not-a-uuid",
			expectError: "routing rule id must be a valid UUID",
		},
		{
			name:        "missing flag id",
			expectError: "--id must not be empty",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var requestedRuleID string
			updateCalled := false

			rule := testRoutingRule()
			rule.ID = test.documentID

			mockGatewayAPI := &api.MockGatewayAPI{
				UpdateGatewayRoutingRuleV5Func: func(
					endpoints configuration_models.EndpointsV2,
					apiKey string,
					organizationID string,
					gatewayID string,
					ruleID string,
					body json.RawMessage,
				) (json.RawMessage, error) {
					updateCalled = true
					requestedRuleID = ruleID
					return mustMarshalRule(t, rule), nil
				},
			}

			cmd := setupRoutingRulesTestCommand()
			if test.flagID != "" {
				_ = cmd.Flags().Set("id", test.flagID)
			}
			_ = cmd.Flags().Set("file-path", "-")
			cmd.SetIn(bytes.NewReader(mustMarshalRule(t, rule)))

			err := Update(Dependencies{GatewayAPI: mockGatewayAPI}, cmd, nil, routingRulesTestProfile(), nil)

			if test.expectError != "" {
				if err == nil {
					t.Fatal("Expected error, got nil")
				}
				if !strings.Contains(err.Error(), test.expectError) {
					t.Fatalf("Expected error containing %q, got %q", test.expectError, err.Error())
				}
				if updateCalled {
					t.Fatal("UpdateGatewayRoutingRuleV5 should not be called when validation fails")
				}
				return
			}

			if err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}
			if requestedRuleID != test.expectedID {
				t.Fatalf("Expected rule ID %s, got %q", test.expectedID, requestedRuleID)
			}
		})
	}
}

func TestUpdate_PropagatesUnprocessableEntityMessage(t *testing.T) {
	rule := testRoutingRule()

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

	cmd := setupRoutingRulesTestCommand()
	_ = cmd.Flags().Set("id", testRoutingRuleID)
	_ = cmd.Flags().Set("file-path", "-")
	cmd.SetIn(bytes.NewReader(mustMarshalRule(t, rule)))

	err := Update(Dependencies{GatewayAPI: mockGatewayAPI}, cmd, nil, routingRulesTestProfile(), nil)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	for _, expected := range []string{
		"failed to update routing rule",
		"received 422",
		"input storage class",
		"not part of the gateway vocabulary",
	} {
		if !strings.Contains(err.Error(), expected) {
			t.Fatalf("Expected error to contain %q, got %q", expected, err.Error())
		}
	}
}

func TestDelete_PrintsSuccessMessage(t *testing.T) {
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

	cmd := setupRoutingRulesTestCommand()
	_ = cmd.Flags().Set("id", strings.ToUpper(testRoutingRuleID))

	if err := Delete(Dependencies{GatewayAPI: mockGatewayAPI}, cmd, nil, routingRulesTestProfile(), nil); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if requestedRuleID != testRoutingRuleID {
		t.Fatalf("Expected rule ID %s, got %q", testRoutingRuleID, requestedRuleID)
	}
	expectedResult := "Routing rule " + testRoutingRuleID + " deleted successfully."
	actualResult := strings.TrimSpace(outputOf(cmd))
	if actualResult != expectedResult {
		t.Fatalf("Expected delete message, got %q", actualResult)
	}
}

func TestGet_RequiresGatewayID(t *testing.T) {
	cmd := setupRoutingRulesTestCommand()
	_ = cmd.Flags().Set("gateway-id", "   ")
	_ = cmd.Flags().Set("id", testRoutingRuleID)

	err := Get(Dependencies{GatewayAPI: &api.MockGatewayAPI{}}, cmd, nil, routingRulesTestProfile(), nil)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	if err.Error() != "--gateway-id must not be empty" {
		t.Fatalf("Expected empty gateway-id error, got %q", err.Error())
	}
}

func TestGet_RequiresID(t *testing.T) {
	cmd := setupRoutingRulesTestCommand()

	err := Get(Dependencies{GatewayAPI: &api.MockGatewayAPI{}}, cmd, nil, routingRulesTestProfile(), nil)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	if err.Error() != "--id must not be empty" {
		t.Fatalf("Expected empty id error, got %q", err.Error())
	}
}

func TestUpdate_RequiresFilePath(t *testing.T) {
	cmd := setupRoutingRulesTestCommand()
	_ = cmd.Flags().Set("id", testRoutingRuleID)

	err := Update(Dependencies{GatewayAPI: &api.MockGatewayAPI{}}, cmd, nil, routingRulesTestProfile(), nil)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if !strings.Contains(err.Error(), "--file-path must not be empty") {
		t.Fatalf("Expected missing file-path error, got %q", err.Error())
	}
}

func TestCreate_PropagatesUnprocessableEntityMessage(t *testing.T) {
	rule := testRoutingRule()

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
					"validation failed: unknown storage class",
			)
		},
	}

	cmd := setupRoutingRulesTestCommand()
	_ = cmd.Flags().Set("file-path", "-")
	cmd.SetIn(bytes.NewReader(mustMarshalRule(t, rule)))

	err := Create(Dependencies{GatewayAPI: mockGatewayAPI}, cmd, nil, routingRulesTestProfile(), nil)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	for _, expected := range []string{"failed to create routing rule", "received 422", "unknown storage class"} {
		if !strings.Contains(err.Error(), expected) {
			t.Fatalf("Expected error to contain %q, got %q", expected, err.Error())
		}
	}
}

func TestUpdate_ReadsFilePath(t *testing.T) {
	rule := testRoutingRule()
	filePath := filepath.Join(t.TempDir(), "rule.json")
	if err := os.WriteFile(filePath, mustMarshalRule(t, rule), 0o600); err != nil {
		t.Fatalf("Expected the fixture to be written, got %v", err)
	}

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
			return mustMarshalRule(t, rule), nil
		},
	}

	cmd := setupRoutingRulesTestCommand()
	_ = cmd.Flags().Set("id", testRoutingRuleID)
	_ = cmd.Flags().Set("file-path", filePath)

	if err := Update(Dependencies{GatewayAPI: mockGatewayAPI}, cmd, nil, routingRulesTestProfile(), nil); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedRequest, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Expected the fixture to be readable, got %v", err)
	}
	if !bytes.Equal(received, expectedRequest) {
		t.Fatalf("Expected the file contents to be forwarded.\nwant: %s\ngot:  %s", expectedRequest, received)
	}
}

func TestUpdate_ReportsFileReadError(t *testing.T) {
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

	cmd := setupRoutingRulesTestCommand()
	_ = cmd.Flags().Set("id", testRoutingRuleID)
	_ = cmd.Flags().Set("file-path", filepath.Join(t.TempDir(), "missing.json"))

	err := Update(Dependencies{GatewayAPI: mockGatewayAPI}, cmd, nil, routingRulesTestProfile(), nil)
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
