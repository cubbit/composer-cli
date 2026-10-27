package cmd_gateway

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	servicegateway "github.com/cubbit/composer-cli/src/service/gateway"
)

func TestGatewaySubCmd_Create_Output_Inline_JSON(t *testing.T) {
	mockCfg := newTestGatewayConfig()

	mockGatewayAPI := &api.MockGatewayAPI{
		CreateGatewayV5Func: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			_ *api.CreateGatewayV5Request,
		) (*api.CreateGatewayV5Response, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}

			return &api.CreateGatewayV5Response{ID: "response-id"}, nil
		},
	}

	mockProcessAPI := &api.MockProcessAPI{
		GetProcessFunc: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			_ string,
		) (*api.Process, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}

			data, _ := json.Marshal(api.GatewayCreationProcessData{ID: "gw-created"})
			return &api.Process{
				ID:     "process-id",
				Status: api.ProcessStatusRunning,
				Data:   data,
			}, nil
		},
	}

	gatewayService := servicegateway.NewGatewayService(
		mockCfg,
		mockGatewayAPI,
		&api.MockSwarmAPI{},
		&api.MockRedundancyClassAPI{},
		mockProcessAPI,
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
		"create",
		"--name", "Test",
		"--slug", "test",
		"--cluster-id", "550e8400-e29b-41d4-a716-446655440000",
		"--ingress-type", "manual",
		"--swarm-rc", "swarm-1:rc-1:true",
		"--output", "json",
	})

	err := gatewayCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `{
  "message": "Gateway creation started — Gateway ID: gw-created\n"
}
`
	if commandOutput.String() != expected {
		t.Errorf("JSON snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

func TestGatewaySubCmd_Create_Output_Inline_YAML(t *testing.T) {
	mockCfg := newTestGatewayConfig()

	mockGatewayAPI := &api.MockGatewayAPI{
		CreateGatewayV5Func: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			_ *api.CreateGatewayV5Request,
		) (*api.CreateGatewayV5Response, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}

			return &api.CreateGatewayV5Response{ID: "response-id"}, nil
		},
	}

	mockProcessAPI := &api.MockProcessAPI{
		GetProcessFunc: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			_ string,
		) (*api.Process, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}

			data, _ := json.Marshal(api.GatewayCreationProcessData{ID: "gw-created"})
			return &api.Process{
				ID:     "process-id",
				Status: api.ProcessStatusRunning,
				Data:   data,
			}, nil
		},
	}

	gatewayService := servicegateway.NewGatewayService(
		mockCfg,
		mockGatewayAPI,
		&api.MockSwarmAPI{},
		&api.MockRedundancyClassAPI{},
		mockProcessAPI,
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
		"create",
		"--name", "Test",
		"--slug", "test",
		"--cluster-id", "550e8400-e29b-41d4-a716-446655440000",
		"--ingress-type", "manual",
		"--swarm-rc", "swarm-1:rc-1:true",
		"--output", "yaml",
	})

	err := gatewayCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `message: |
    Gateway creation started — Gateway ID: gw-created
`
	if commandOutput.String() != expected {
		t.Errorf("YAML snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

func TestGatewaySubCmd_Create_InteractiveAndOutputConflict(t *testing.T) {
	mockCfg := newTestGatewayConfig()

	mockGatewayAPI := &api.MockGatewayAPI{}

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

	gatewayCmd.SetArgs([]string{
		"create",
		"--interactive",
		"--output", "json",
	})

	err := gatewayCmd.Execute()
	if err == nil {
		t.Fatal("Expected error for conflicting --interactive and --output flags, got nil")
	}

	if !strings.Contains(err.Error(), "mutually exclusive") {
		t.Errorf("Expected error message to mention flags are mutually exclusive, got %q", err.Error())
	}
}

func TestGatewaySubCmd_Create_Output_Inline_JSON_FromProfile(t *testing.T) {
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
		CreateGatewayV5Func: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			_ *api.CreateGatewayV5Request,
		) (*api.CreateGatewayV5Response, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}

			return &api.CreateGatewayV5Response{ID: "response-id"}, nil
		},
	}

	mockProcessAPI := &api.MockProcessAPI{
		GetProcessFunc: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			_ string,
		) (*api.Process, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}

			data, _ := json.Marshal(api.GatewayCreationProcessData{ID: "gw-created"})
			return &api.Process{
				ID:     "process-id",
				Status: api.ProcessStatusRunning,
				Data:   data,
			}, nil
		},
	}

	gatewayService := servicegateway.NewGatewayService(
		mockCfg,
		mockGatewayAPI,
		&api.MockSwarmAPI{},
		&api.MockRedundancyClassAPI{},
		mockProcessAPI,
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
		"create",
		"--name", "Test",
		"--slug", "test",
		"--cluster-id", "550e8400-e29b-41d4-a716-446655440000",
		"--ingress-type", "manual",
		"--swarm-rc", "swarm-1:rc-1:true",
	})

	err := gatewayCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `{
  "message": "Gateway creation started — Gateway ID: gw-created\n"
}
`
	if commandOutput.String() != expected {
		t.Errorf("JSON snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

func TestGatewaySubCmd_Create_Output_Inline_YAML_FromProfile(t *testing.T) {
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
		CreateGatewayV5Func: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			_ *api.CreateGatewayV5Request,
		) (*api.CreateGatewayV5Response, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}

			return &api.CreateGatewayV5Response{ID: "response-id"}, nil
		},
	}

	mockProcessAPI := &api.MockProcessAPI{
		GetProcessFunc: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			_ string,
		) (*api.Process, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}

			data, _ := json.Marshal(api.GatewayCreationProcessData{ID: "gw-created"})
			return &api.Process{
				ID:     "process-id",
				Status: api.ProcessStatusRunning,
				Data:   data,
			}, nil
		},
	}

	gatewayService := servicegateway.NewGatewayService(
		mockCfg,
		mockGatewayAPI,
		&api.MockSwarmAPI{},
		&api.MockRedundancyClassAPI{},
		mockProcessAPI,
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
		"create",
		"--name", "Test",
		"--slug", "test",
		"--cluster-id", "550e8400-e29b-41d4-a716-446655440000",
		"--ingress-type", "manual",
		"--swarm-rc", "swarm-1:rc-1:true",
	})

	err := gatewayCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `message: |
    Gateway creation started — Gateway ID: gw-created
`
	if commandOutput.String() != expected {
		t.Errorf("YAML snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}
