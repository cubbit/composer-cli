package cmd_gateway

import (
	"bytes"
	"encoding/json"
	"errors"
	"strings"
	"testing"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration"
	servicegateway "github.com/cubbit/composer-cli/src/service/gateway"
	"github.com/spf13/cobra"
)

func setupGatewayCreateInlineCommand(gatewayService servicegateway.GatewayServiceInterface) (*cobra.Command, *bytes.Buffer) {
	gatewayCmd := NewGatewayCmd(gatewayService)
	gatewayCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	commandOutput := new(bytes.Buffer)
	gatewayCmd.SetOut(commandOutput)
	gatewayCmd.SetErr(commandOutput)

	return gatewayCmd, commandOutput
}

func TestGatewaySubCmd_Create_Integration_Inline_Success(t *testing.T) {
	mockCfg := api.NewMockConfig(configuration.ProfileTypeComposer, "test-api-key", "test-org-id")

	mockGatewayAPI := &api.MockGatewayAPI{
		CreateGatewayV5Func: func(
			urlConfig configuration.URLs,
			apiKey string,
			organizationID string,
			request *api.CreateGatewayV5Request,
		) (*api.CreateGatewayV5Response, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}
			if request.Name != "test-gateway" {
				t.Fatalf("Expected gateway name 'test-gateway', got %q", request.Name)
			}
			if request.Slug != "test-gateway" {
				t.Fatalf("Expected gateway slug 'test-gateway', got %q", request.Slug)
			}
			if request.ClusterID != "550e8400-e29b-41d4-a716-446655440000" {
				t.Fatalf("Expected cluster ID, got %q", request.ClusterID)
			}
			if request.CubbitIngress.Type != api.CubbitIngressType("manual") {
				t.Fatalf("Expected ingress type 'manual', got %q", request.CubbitIngress.Type)
			}
			if len(request.SwarmsAndRedundancyClass) != 1 {
				t.Fatalf("Expected 1 swarm-rc pair, got %d", len(request.SwarmsAndRedundancyClass))
			}

			return &api.CreateGatewayV5Response{ID: "test-process-id-123"}, nil
		},
	}

	mockProcessAPI := &api.MockProcessAPI{
		GetProcessFunc: func(
			_ configuration.URLs, _ string, _ string, _ string,
		) (*api.Process, error) {
			data, _ := json.Marshal(api.GatewayCreationProcessData{ID: "gateway-id-001"})
			return &api.Process{
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

	gatewayCmd, commandOutput := setupGatewayCreateInlineCommand(gatewayService)
	gatewayCmd.SetArgs([]string{
		"create",
		"--name", "test-gateway",
		"--slug", "test-gateway",
		"--cluster-id", "550e8400-e29b-41d4-a716-446655440000",
		"--description", "Test gateway description",
		"--ingress-type", "manual",
		"--swarm-rc", "550e8400-e29b-41d4-a716-446655440001:550e8400-e29b-41d4-a716-446655440002:true",
	})

	err := gatewayCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := "Gateway creation started — Gateway ID: gateway-id-001\n"
	actual := commandOutput.String()
	if actual != expected {
		t.Fatalf("Expected output %q, got %q", expected, actual)
	}
}

func TestGatewaySubCmd_Create_Integration_Inline_WithOptionalDescription(t *testing.T) {
	mockCfg := api.NewMockConfig(configuration.ProfileTypeComposer, "test-api-key", "test-org-id")

	mockGatewayAPI := &api.MockGatewayAPI{
		CreateGatewayV5Func: func(
			urlConfig configuration.URLs,
			apiKey string,
			organizationID string,
			request *api.CreateGatewayV5Request,
		) (*api.CreateGatewayV5Response, error) {
			if request.Description == nil || *request.Description != "Optional description" {
				t.Fatalf("Expected description 'Optional description', got %v", request.Description)
			}
			return &api.CreateGatewayV5Response{ID: "process-with-desc"}, nil
		},
	}

	mockProcessAPI := &api.MockProcessAPI{
		GetProcessFunc: func(
			_ configuration.URLs, _ string, _ string, _ string,
		) (*api.Process, error) {
			data, _ := json.Marshal(api.GatewayCreationProcessData{ID: "gateway-desc-id"})
			return &api.Process{
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

	gatewayCmd, commandOutput := setupGatewayCreateInlineCommand(gatewayService)
	gatewayCmd.SetArgs([]string{
		"create",
		"--name", "gateway-with-desc",
		"--slug", "gateway-with-desc",
		"--cluster-id", "550e8400-e29b-41d4-a716-446655440000",
		"--description", "Optional description",
		"--ingress-type", "manual",
		"--swarm-rc", "swarm-1:rc-1:true",
	})

	err := gatewayCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := "Gateway creation started — Gateway ID: gateway-desc-id\n"
	actual := commandOutput.String()
	if actual != expected {
		t.Fatalf("Expected output %q, got %q", expected, actual)
	}
}

func TestGatewaySubCmd_Create_Integration_Inline_WithMultipleSwarmRC(t *testing.T) {
	mockCfg := api.NewMockConfig(configuration.ProfileTypeComposer, "test-api-key", "test-org-id")

	mockGatewayAPI := &api.MockGatewayAPI{
		CreateGatewayV5Func: func(
			urlConfig configuration.URLs,
			apiKey string,
			organizationID string,
			request *api.CreateGatewayV5Request,
		) (*api.CreateGatewayV5Response, error) {
			if len(request.SwarmsAndRedundancyClass) != 2 {
				t.Fatalf("Expected 2 swarm-rc pairs, got %d", len(request.SwarmsAndRedundancyClass))
			}
			if !request.SwarmsAndRedundancyClass[0].IsDefault {
				t.Fatal("Expected first swarm-rc to be marked as default")
			}
			if request.SwarmsAndRedundancyClass[1].IsDefault {
				t.Fatal("Expected second swarm-rc to not be marked as default")
			}
			return &api.CreateGatewayV5Response{ID: "process-multi-swarm"}, nil
		},
	}

	mockProcessAPI := &api.MockProcessAPI{
		GetProcessFunc: func(
			_ configuration.URLs, _ string, _ string, _ string,
		) (*api.Process, error) {
			data, _ := json.Marshal(api.GatewayCreationProcessData{ID: "gateway-multi-id"})
			return &api.Process{
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

	gatewayCmd, commandOutput := setupGatewayCreateInlineCommand(gatewayService)
	gatewayCmd.SetArgs([]string{
		"create",
		"--name", "multi-swarm-gateway",
		"--slug", "multi-swarm-gateway",
		"--cluster-id", "550e8400-e29b-41d4-a716-446655440000",
		"--ingress-type", "manual",
		"--swarm-rc", "swarm-1:rc-1:true",
		"--swarm-rc", "swarm-2:rc-2:false",
	})

	err := gatewayCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := "Gateway creation started — Gateway ID: gateway-multi-id\n"
	actual := commandOutput.String()
	if actual != expected {
		t.Fatalf("Expected output %q, got %q", expected, actual)
	}
}

func TestGatewaySubCmd_Create_Integration_Inline_WithoutDescription(t *testing.T) {
	mockCfg := api.NewMockConfig(configuration.ProfileTypeComposer, "test-api-key", "test-org-id")

	mockGatewayAPI := &api.MockGatewayAPI{
		CreateGatewayV5Func: func(
			urlConfig configuration.URLs,
			apiKey string,
			organizationID string,
			request *api.CreateGatewayV5Request,
		) (*api.CreateGatewayV5Response, error) {
			if request.Description != nil {
				t.Fatalf("Expected nil description when not provided, got %v", *request.Description)
			}
			return &api.CreateGatewayV5Response{ID: "process-no-desc"}, nil
		},
	}

	mockProcessAPI := &api.MockProcessAPI{
		GetProcessFunc: func(
			_ configuration.URLs, _ string, _ string, _ string,
		) (*api.Process, error) {
			data, _ := json.Marshal(api.GatewayCreationProcessData{ID: "gateway-no-desc-id"})
			return &api.Process{
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

	gatewayCmd, commandOutput := setupGatewayCreateInlineCommand(gatewayService)
	gatewayCmd.SetArgs([]string{
		"create",
		"--name", "no-desc-gateway",
		"--slug", "no-desc-gateway",
		"--cluster-id", "550e8400-e29b-41d4-a716-446655440000",
		"--ingress-type", "manual",
		"--swarm-rc", "swarm-1:rc-1:true",
	})

	err := gatewayCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := "Gateway creation started — Gateway ID: gateway-no-desc-id\n"
	actual := commandOutput.String()
	if actual != expected {
		t.Fatalf("Expected output %q, got %q", expected, actual)
	}
}

func TestGatewaySubCmd_Create_Integration_Inline_APIError(t *testing.T) {
	mockCfg := api.NewMockConfig(configuration.ProfileTypeComposer, "test-api-key", "test-org-id")

	mockGatewayAPI := &api.MockGatewayAPI{
		CreateGatewayV5Func: func(
			urlConfig configuration.URLs,
			apiKey string,
			organizationID string,
			request *api.CreateGatewayV5Request,
		) (*api.CreateGatewayV5Response, error) {
			return nil, errors.New("internal server error")
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

	gatewayCmd, commandOutput := setupGatewayCreateInlineCommand(gatewayService)
	gatewayCmd.SetArgs([]string{
		"create",
		"--name", "failing-gateway",
		"--slug", "failing-gateway",
		"--cluster-id", "550e8400-e29b-41d4-a716-446655440000",
		"--ingress-type", "manual",
		"--swarm-rc", "swarm-1:rc-1:true",
	})

	err := gatewayCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error from Execute (Run prints errors, doesn't return them), got %v", err)
	}

	output := commandOutput.String()
	if !strings.Contains(output, "failed to create gateway") {
		t.Fatalf("Expected output to contain 'failed to create gateway', got %q", output)
	}
	if !strings.Contains(output, "internal server error") {
		t.Fatalf("Expected output to contain 'internal server error', got %q", output)
	}
}

func TestGatewaySubCmd_Create_Integration_Inline_MissingName(t *testing.T) {
	mockCfg := api.NewMockConfig(configuration.ProfileTypeComposer, "test-api-key", "test-org-id")

	gatewayService := servicegateway.NewGatewayService(
		mockCfg,
		&api.MockGatewayAPI{},
		&api.MockSwarmAPI{},
		&api.MockRedundancyClassAPI{},
		&api.MockProcessAPI{},
		&api.MockLocationAPI{},
	)

	gatewayCmd, commandOutput := setupGatewayCreateInlineCommand(gatewayService)
	gatewayCmd.SetArgs([]string{
		"create",
		"--slug", "test-gateway",
		"--cluster-id", "550e8400-e29b-41d4-a716-446655440000",
		"--ingress-type", "manual",
		"--swarm-rc", "swarm-1:rc-1:true",
	})

	err := gatewayCmd.Execute()
	if err == nil {
		t.Fatal("Expected error when --name is missing, got nil")
	}

	if !strings.Contains(err.Error(), `required flag(s) "name" not set`) {
		t.Fatalf("Expected error message to indicate missing required flag 'name', got: %v", err)
	}

	_ = commandOutput // error is returned, not printed
}

func TestGatewaySubCmd_Create_Integration_Inline_MissingSwarmRC(t *testing.T) {
	mockCfg := api.NewMockConfig(configuration.ProfileTypeComposer, "test-api-key", "test-org-id")

	gatewayService := servicegateway.NewGatewayService(
		mockCfg,
		&api.MockGatewayAPI{},
		&api.MockSwarmAPI{},
		&api.MockRedundancyClassAPI{},
		&api.MockProcessAPI{},
		&api.MockLocationAPI{},
	)

	gatewayCmd, commandOutput := setupGatewayCreateInlineCommand(gatewayService)
	gatewayCmd.SetArgs([]string{
		"create",
		"--name", "test-gateway",
		"--slug", "test-gateway",
		"--cluster-id", "550e8400-e29b-41d4-a716-446655440000",
		"--ingress-type", "manual",
	})

	err := gatewayCmd.Execute()
	if err == nil {
		t.Fatal("Expected error when --swarm-rc is missing, got nil")
	}

	if !strings.Contains(err.Error(), `required flag(s) "swarm-rc" not set`) {
		t.Fatalf("Expected error message to indicate missing required flag 'swarm-rc', got: %v", err)
	}

	_ = commandOutput // error is returned, not printed
}

func TestGatewaySubCmd_Create_Integration_Inline_ConfigTypeMismatch(t *testing.T) {
	// Use Console profile type instead of Composer to trigger type mismatch
	mockCfg := api.NewMockConfig(configuration.ProfileTypeConsole, "test-api-key", "test-org-id")

	gatewayService := servicegateway.NewGatewayService(
		mockCfg,
		&api.MockGatewayAPI{},
		&api.MockSwarmAPI{},
		&api.MockRedundancyClassAPI{},
		&api.MockProcessAPI{},
		&api.MockLocationAPI{},
	)

	gatewayCmd, commandOutput := setupGatewayCreateInlineCommand(gatewayService)
	gatewayCmd.SetArgs([]string{
		"create",
		"--name", "test-gateway",
		"--slug", "test-gateway",
		"--cluster-id", "550e8400-e29b-41d4-a716-446655440000",
		"--ingress-type", "manual",
		"--swarm-rc", "swarm-1:rc-1:true",
	})

	err := gatewayCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error from Execute (Run prints errors), got %v", err)
	}

	output := commandOutput.String()
	if !strings.Contains(output, "error while loading file path configuration") {
		t.Fatalf("Expected output to contain 'error while loading file path configuration', got %q", output)
	}
}

func TestGatewaySubCmd_Create_Integration_Inline_ExistingGatewayProcess(t *testing.T) {
	mockCfg := api.NewMockConfig(configuration.ProfileTypeComposer, "test-api-key", "test-org-id")

	data, _ := json.Marshal(api.GatewayCreationProcessData{ID: "gateway-id-001"})

	mockGatewayAPI := &api.MockGatewayAPI{}
	mockSwarmAPI := &api.MockSwarmAPI{}
	mockRCApi := &api.MockRedundancyClassAPI{}
	mockLocationAPI := &api.MockLocationAPI{}
	mockProcessAPI := &api.MockProcessAPI{
		Processes: []api.Process{
			{
				Type:   api.ProcessTypeGatewayCreation,
				Status: api.ProcessStatusRunning,
				Data:   data,
			},
		},
	}

	gatewayService := servicegateway.NewGatewayService(
		mockCfg,
		mockGatewayAPI,
		mockSwarmAPI,
		mockRCApi,
		mockProcessAPI,
		mockLocationAPI,
	)

	gatewayCmd, commandOutput := setupGatewayCreateInlineCommand(gatewayService)
	gatewayCmd.SetArgs([]string{
		"create",
		"--name", "test-gateway",
		"--slug", "test-gateway",
		"--cluster-id", "cluster-eu-01",
		"--ingress-type", "manual",
		"--swarm-rc", "swarm-001:rc-001:true",
	})

	err := gatewayCmd.Execute()
	if err != nil {
		t.Fatalf("expected no command error, got %v", err)
	}

	output := commandOutput.String()
	expectedMsg := "a gateway creation is already in progress for gateway gateway-id-001"
	if !strings.Contains(output, expectedMsg) {
		t.Fatalf("expected output to contain %q, got %q", expectedMsg, output)
	}
}
