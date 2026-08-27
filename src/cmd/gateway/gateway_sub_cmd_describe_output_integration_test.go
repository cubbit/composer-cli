package cmd_gateway

import (
	"bytes"
	"testing"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	servicegateway "github.com/cubbit/composer-cli/src/service/gateway"
)

func TestGatewaySubCmd_Describe_Output_JSON(t *testing.T) {
	mockCfg := newTestGatewayConfig()

	mockGatewayAPI := &api.MockGatewayAPI{
		GetGatewayV5Func: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
		) (*api.GatewayV5GetResponse, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}
			if gatewayID != "gw-json" {
				t.Fatalf("Expected gateway ID gw-json, got %q", gatewayID)
			}

			return &api.GatewayV5GetResponse{
				ID:     "gw-json",
				Name:   "json-test-gateway",
				Slug:   "json-test-gateway",
				Type:   api.IngressTypeManual,
				Status: api.GatewayV5StatusReady,
				RedundancyClasses: []api.GatewayV5GetRedundancyClass{
					{ID: "rc-001", Name: "redundancy-alpha"},
					{ID: "rc-002", Name: "redundancy-beta"},
				},
			}, nil
		},
	}

	gatewayService := servicegateway.NewGatewayService(
		mockCfg,
		mockGatewayAPI,
		&api.MockSwarmAPI{},
		&api.MockRedundancyClassAPI{},
		&api.MockProcessAPI{},
		&api.MockLocationAPI{},
	)

	gatewayCmd := NewGatewayCmd(gatewayService)
	gatewayCmd.PersistentFlags().String("profile", "", "Profile")
	gatewayCmd.PersistentFlags().String("output", "human", "Output format")
	gatewayCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	gatewayCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	commandOutput := new(bytes.Buffer)
	gatewayCmd.SetOut(commandOutput)
	gatewayCmd.SetErr(commandOutput)
	gatewayCmd.SetArgs([]string{
		"describe",
		"gw-json",
		"--output", "json",
	})

	err := gatewayCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `{
  "id": "gw-json",
  "name": "json-test-gateway",
  "slug": "json-test-gateway",
  "type": "manual",
  "redundancy_classes": [
    {
      "id": "rc-001",
      "name": "redundancy-alpha"
    },
    {
      "id": "rc-002",
      "name": "redundancy-beta"
    }
  ],
  "status": "ready"
}
`
	if commandOutput.String() != expected {
		t.Errorf("JSON snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

func TestGatewaySubCmd_Describe_Output_YAML(t *testing.T) {
	mockCfg := newTestGatewayConfig()

	mockGatewayAPI := &api.MockGatewayAPI{
		GetGatewayV5Func: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
		) (*api.GatewayV5GetResponse, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}
			if gatewayID != "gw-yaml" {
				t.Fatalf("Expected gateway ID gw-yaml, got %q", gatewayID)
			}

			return &api.GatewayV5GetResponse{
				ID:     "gw-yaml",
				Name:   "yaml-test-gateway",
				Slug:   "yaml-test-gateway",
				Type:   api.IngressTypeManual,
				Status: api.GatewayV5StatusReady,
				RedundancyClasses: []api.GatewayV5GetRedundancyClass{
					{ID: "rc-003", Name: "redundancy-gamma"},
					{ID: "rc-004", Name: "redundancy-delta"},
				},
			}, nil
		},
	}

	gatewayService := servicegateway.NewGatewayService(
		mockCfg,
		mockGatewayAPI,
		&api.MockSwarmAPI{},
		&api.MockRedundancyClassAPI{},
		&api.MockProcessAPI{},
		&api.MockLocationAPI{},
	)

	gatewayCmd := NewGatewayCmd(gatewayService)
	gatewayCmd.PersistentFlags().String("profile", "", "Profile")
	gatewayCmd.PersistentFlags().String("output", "human", "Output format")
	gatewayCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	gatewayCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	commandOutput := new(bytes.Buffer)
	gatewayCmd.SetOut(commandOutput)
	gatewayCmd.SetErr(commandOutput)
	gatewayCmd.SetArgs([]string{
		"describe",
		"gw-yaml",
		"--output", "yaml",
	})

	err := gatewayCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `id: gw-yaml
name: yaml-test-gateway
slug: yaml-test-gateway
type: manual
redundancy_classes:
    - id: rc-003
      name: redundancy-gamma
    - id: rc-004
      name: redundancy-delta
status: ready
`
	if commandOutput.String() != expected {
		t.Errorf("YAML snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

func TestGatewaySubCmd_Describe_Output_JSON_FromProfile(t *testing.T) {
	mockCfg := newTestGatewayConfig()

	mockCfg.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{
			APIKey:         "test-api-key",
			OrganizationID: "test-org-id",
			Output:         "json",
			Endpoints:      configuration_models.EndpointsV2{},
		}, nil
	}

	mockGatewayAPI := &api.MockGatewayAPI{
		GetGatewayV5Func: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
		) (*api.GatewayV5GetResponse, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}
			if gatewayID != "gw-json" {
				t.Fatalf("Expected gateway ID gw-json, got %q", gatewayID)
			}

			return &api.GatewayV5GetResponse{
				ID:     "gw-json",
				Name:   "json-test-gateway",
				Slug:   "json-test-gateway",
				Type:   api.IngressTypeManual,
				Status: api.GatewayV5StatusReady,
				RedundancyClasses: []api.GatewayV5GetRedundancyClass{
					{ID: "rc-001", Name: "redundancy-alpha"},
					{ID: "rc-002", Name: "redundancy-beta"},
				},
			}, nil
		},
	}

	gatewayService := servicegateway.NewGatewayService(
		mockCfg,
		mockGatewayAPI,
		&api.MockSwarmAPI{},
		&api.MockRedundancyClassAPI{},
		&api.MockProcessAPI{},
		&api.MockLocationAPI{},
	)

	gatewayCmd := NewGatewayCmd(gatewayService)
	gatewayCmd.PersistentFlags().String("profile", "", "Profile")
	gatewayCmd.PersistentFlags().String("output", "human", "Output format")
	gatewayCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	gatewayCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	commandOutput := new(bytes.Buffer)
	gatewayCmd.SetOut(commandOutput)
	gatewayCmd.SetErr(commandOutput)
	gatewayCmd.SetArgs([]string{
		"describe",
		"gw-json",
	})

	err := gatewayCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `{
  "id": "gw-json",
  "name": "json-test-gateway",
  "slug": "json-test-gateway",
  "type": "manual",
  "redundancy_classes": [
    {
      "id": "rc-001",
      "name": "redundancy-alpha"
    },
    {
      "id": "rc-002",
      "name": "redundancy-beta"
    }
  ],
  "status": "ready"
}
`
	if commandOutput.String() != expected {
		t.Errorf("JSON snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

func TestGatewaySubCmd_Describe_Output_YAML_FromProfile(t *testing.T) {
	mockCfg := newTestGatewayConfig()

	mockCfg.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{
			APIKey:         "test-api-key",
			OrganizationID: "test-org-id",
			Output:         "yaml",
			Endpoints:      configuration_models.EndpointsV2{},
		}, nil
	}

	mockGatewayAPI := &api.MockGatewayAPI{
		GetGatewayV5Func: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
		) (*api.GatewayV5GetResponse, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}
			if gatewayID != "gw-yaml" {
				t.Fatalf("Expected gateway ID gw-yaml, got %q", gatewayID)
			}

			return &api.GatewayV5GetResponse{
				ID:     "gw-yaml",
				Name:   "yaml-test-gateway",
				Slug:   "yaml-test-gateway",
				Type:   api.IngressTypeManual,
				Status: api.GatewayV5StatusReady,
				RedundancyClasses: []api.GatewayV5GetRedundancyClass{
					{ID: "rc-003", Name: "redundancy-gamma"},
					{ID: "rc-004", Name: "redundancy-delta"},
				},
			}, nil
		},
	}

	gatewayService := servicegateway.NewGatewayService(
		mockCfg,
		mockGatewayAPI,
		&api.MockSwarmAPI{},
		&api.MockRedundancyClassAPI{},
		&api.MockProcessAPI{},
		&api.MockLocationAPI{},
	)

	gatewayCmd := NewGatewayCmd(gatewayService)
	gatewayCmd.PersistentFlags().String("profile", "", "Profile")
	gatewayCmd.PersistentFlags().String("output", "human", "Output format")
	gatewayCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	gatewayCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	commandOutput := new(bytes.Buffer)
	gatewayCmd.SetOut(commandOutput)
	gatewayCmd.SetErr(commandOutput)
	gatewayCmd.SetArgs([]string{
		"describe",
		"gw-yaml",
	})

	err := gatewayCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `id: gw-yaml
name: yaml-test-gateway
slug: yaml-test-gateway
type: manual
redundancy_classes:
    - id: rc-003
      name: redundancy-gamma
    - id: rc-004
      name: redundancy-delta
status: ready
`
	if commandOutput.String() != expected {
		t.Errorf("YAML snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}
