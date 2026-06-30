package cmd_tenant

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	servicetenant "github.com/cubbit/composer-cli/src/service/tenant"
	"github.com/spf13/cobra"
)

func newTestTenantConfig() *configuration_handler.MockConfigurationHandler {
	mockCfg := configuration_handler.NewMockConfigurationHandler()
	mockCfg.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{
			Output:         configuration_models.OutputHuman,
			APIKey:         "test-api-key",
			OrganizationID: "test-org-id",
			Endpoints:      configuration_models.EndpointsV2{},
		}, nil
	}
	return mockCfg
}

func mockTenantCreationProcess(tenantID string) *api.Process {
	data, _ := json.Marshal(api.TenantCreationProcessData{TenantID: tenantID})
	return &api.Process{
		ID:     "process-id",
		Type:   api.ProcessTypeTenantCreation,
		Status: api.ProcessStatusSuccess,
		Data:   data,
	}
}

func setupTenantCreateInlineCommand(tenantService servicetenant.TenantServiceInterface) (*cobra.Command, *bytes.Buffer) {
	tenantCmd := NewTenantCmd(tenantService)
	tenantCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	commandOutput := new(bytes.Buffer)
	tenantCmd.SetOut(commandOutput)
	tenantCmd.SetErr(commandOutput)

	return tenantCmd, commandOutput
}

func TestTenantSubCmd_Create_Integration_Inline_Success(t *testing.T) {
	mockCfg := newTestTenantConfig()

	mockTenantAPI := &api.MockTenantAPI{
		CreateTenantV5Func: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			request *api.CreateTenantV5Request,
		) (*api.GenericIDResponseModel, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}
			if request.Name != "Tenant Cubbit 1" {
				t.Fatalf("Expected tenant name 'Tenant Cubbit 1', got %q", request.Name)
			}
			if request.Slug != "cubbit-1" {
				t.Fatalf("Expected tenant slug 'cubbit-1', got %q", request.Slug)
			}
			if len(request.Connections) != 1 {
				t.Fatalf("Expected 1 connection, got %d", len(request.Connections))
			}
			if request.Connections[0].DomainID != "domainID" {
				t.Fatalf("Expected domain ID 'domainID', got %q", request.Connections[0].DomainID)
			}
			if len(request.Connections[0].GatewayIDs) != 2 {
				t.Fatalf("Expected 2 gateway IDs, got %d", len(request.Connections[0].GatewayIDs))
			}
			if request.Connections[0].GatewayIDs[0] != "gatewayID1" {
				t.Fatalf("Expected gateway ID 'gatewayID1', got %q", request.Connections[0].GatewayIDs[0])
			}

			return &api.GenericIDResponseModel{ID: "test-tenant-id"}, nil
		},
	}

	mockDomainAPI := &api.MockDomainAPI{}
	mockGatewayAPI := &api.MockGatewayAPI{}
	mockProcessAPI := &api.MockProcessAPI{
		GetProcessFunc: func(
			_ configuration_models.EndpointsV2,
			_ string,
			_ string,
			processID string,
		) (*api.Process, error) {
			if processID != "test-tenant-id" {
				t.Fatalf("Expected process ID 'test-tenant-id', got %q", processID)
			}
			return mockTenantCreationProcess("test-tenant-id"), nil
		},
	}

	tenantService := servicetenant.NewTenantService(mockCfg, mockTenantAPI, mockDomainAPI, mockGatewayAPI, mockProcessAPI)
	tenantCmd, commandOutput := setupTenantCreateInlineCommand(tenantService)
	tenantCmd.SetArgs([]string{
		"create",
		"--name", "Tenant Cubbit 1",
		"--slug", "cubbit-1",
		"--connection", "domainID:gatewayID1,gatewayID2:mysub",
	})

	err := tenantCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := "Tenant creation started — Tenant ID: test-tenant-id\n"
	actual := commandOutput.String()
	if actual != expected {
		t.Fatalf("Expected output %q, got %q", expected, actual)
	}
}

func TestTenantSubCmd_Create_Integration_Inline_WithMultipleConnections(t *testing.T) {
	mockCfg := newTestTenantConfig()

	mockTenantAPI := &api.MockTenantAPI{
		CreateTenantV5Func: func(
			_ configuration_models.EndpointsV2,
			_ string,
			_ string,
			request *api.CreateTenantV5Request,
		) (*api.GenericIDResponseModel, error) {
			if len(request.Connections) != 2 {
				t.Fatalf("Expected 2 connections, got %d", len(request.Connections))
			}
			if request.Connections[0].DomainID != "domainID" {
				t.Fatalf("Expected first domain ID 'domainID', got %q", request.Connections[0].DomainID)
			}
			if request.Connections[1].DomainID != "domainID2" {
				t.Fatalf("Expected second domain ID 'domainID2', got %q", request.Connections[1].DomainID)
			}

			return &api.GenericIDResponseModel{ID: "multi-conn-tenant-id"}, nil
		},
	}

	mockDomainAPI := &api.MockDomainAPI{}
	mockGatewayAPI := &api.MockGatewayAPI{}
	mockProcessAPI := &api.MockProcessAPI{
		GetProcessFunc: func(
			_ configuration_models.EndpointsV2,
			_ string,
			_ string,
			processID string,
		) (*api.Process, error) {
			if processID != "multi-conn-tenant-id" {
				t.Fatalf("Expected process ID 'multi-conn-tenant-id', got %q", processID)
			}
			return mockTenantCreationProcess("multi-conn-tenant-id"), nil
		},
	}

	tenantService := servicetenant.NewTenantService(mockCfg, mockTenantAPI, mockDomainAPI, mockGatewayAPI, mockProcessAPI)
	tenantCmd, commandOutput := setupTenantCreateInlineCommand(tenantService)
	tenantCmd.SetArgs([]string{
		"create",
		"--name", "Tenant Cubbit 1",
		"--slug", "cubbit-1",
		"--connection", "domainID:gatewayID1,gatewayID2",
		"--connection", "domainID2:gatewayID3",
	})

	err := tenantCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := "Tenant creation started — Tenant ID: multi-conn-tenant-id\n"
	actual := commandOutput.String()
	if actual != expected {
		t.Fatalf("Expected output %q, got %q", expected, actual)
	}
}

func TestTenantSubCmd_Create_Integration_Inline_WithOptionalDescription(t *testing.T) {
	mockCfg := newTestTenantConfig()

	mockTenantAPI := &api.MockTenantAPI{
		CreateTenantV5Func: func(
			_ configuration_models.EndpointsV2,
			_ string,
			_ string,
			request *api.CreateTenantV5Request,
		) (*api.GenericIDResponseModel, error) {
			if request.Description == nil || *request.Description != "Production tenant" {
				t.Fatalf("Expected description 'Production tenant', got %v", request.Description)
			}
			return &api.GenericIDResponseModel{ID: "tenant-with-desc"}, nil
		},
	}

	mockDomainAPI := &api.MockDomainAPI{}
	mockGatewayAPI := &api.MockGatewayAPI{}
	mockProcessAPI := &api.MockProcessAPI{
		GetProcessFunc: func(
			_ configuration_models.EndpointsV2,
			_ string,
			_ string,
			processID string,
		) (*api.Process, error) {
			return mockTenantCreationProcess("tenant-with-desc"), nil
		},
	}

	tenantService := servicetenant.NewTenantService(mockCfg, mockTenantAPI, mockDomainAPI, mockGatewayAPI, mockProcessAPI)
	tenantCmd, commandOutput := setupTenantCreateInlineCommand(tenantService)
	tenantCmd.SetArgs([]string{
		"create",
		"--name", "Tenant Cubbit 1",
		"--slug", "cubbit-1",
		"--description", "Production tenant",
		"--connection", "domainID:gatewayID1,gatewayID2:mysub",
	})

	err := tenantCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := "Tenant creation started — Tenant ID: tenant-with-desc\n"
	actual := commandOutput.String()
	if actual != expected {
		t.Fatalf("Expected output %q, got %q", expected, actual)
	}
}

func TestTenantSubCmd_Create_Integration_Inline_WithoutDescription(t *testing.T) {
	mockCfg := newTestTenantConfig()

	mockTenantAPI := &api.MockTenantAPI{
		CreateTenantV5Func: func(
			_ configuration_models.EndpointsV2,
			_ string,
			_ string,
			request *api.CreateTenantV5Request,
		) (*api.GenericIDResponseModel, error) {
			if request.Description != nil {
				t.Fatalf("Expected nil description when not provided, got %v", *request.Description)
			}
			return &api.GenericIDResponseModel{ID: "tenant-no-desc"}, nil
		},
	}

	mockDomainAPI := &api.MockDomainAPI{}
	mockGatewayAPI := &api.MockGatewayAPI{}
	mockProcessAPI := &api.MockProcessAPI{
		GetProcessFunc: func(
			_ configuration_models.EndpointsV2,
			_ string,
			_ string,
			processID string,
		) (*api.Process, error) {
			return mockTenantCreationProcess("tenant-no-desc"), nil
		},
	}

	tenantService := servicetenant.NewTenantService(mockCfg, mockTenantAPI, mockDomainAPI, mockGatewayAPI, mockProcessAPI)
	tenantCmd, commandOutput := setupTenantCreateInlineCommand(tenantService)
	tenantCmd.SetArgs([]string{
		"create",
		"--name", "Tenant Cubbit 1",
		"--slug", "cubbit-1",
		"--connection", "domainID:gatewayID1,gatewayID2:mysub",
	})

	err := tenantCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := "Tenant creation started — Tenant ID: tenant-no-desc\n"
	actual := commandOutput.String()
	if actual != expected {
		t.Fatalf("Expected output %q, got %q", expected, actual)
	}
}

func TestTenantSubCmd_Create_Integration_Inline_APIError(t *testing.T) {
	mockCfg := newTestTenantConfig()

	mockTenantAPI := &api.MockTenantAPI{
		CreateTenantV5Func: func(
			_ configuration_models.EndpointsV2,
			_ string,
			_ string,
			_ *api.CreateTenantV5Request,
		) (*api.GenericIDResponseModel, error) {
			return nil, fmt.Errorf("internal server error")
		},
	}

	mockDomainAPI := &api.MockDomainAPI{}
	mockGatewayAPI := &api.MockGatewayAPI{}
	mockProcessAPI := &api.MockProcessAPI{}

	tenantService := servicetenant.NewTenantService(mockCfg, mockTenantAPI, mockDomainAPI, mockGatewayAPI, mockProcessAPI)
	tenantCmd, commandOutput := setupTenantCreateInlineCommand(tenantService)
	tenantCmd.SetArgs([]string{
		"create",
		"--name", "Tenant Cubbit 1",
		"--slug", "cubbit-1",
		"--connection", "domainID:gatewayID1",
	})

	err := tenantCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error from Execute (Run prints errors, doesn't return them), got %v", err)
	}

	output := commandOutput.String()
	if !strings.Contains(output, "failed request to create the tenant") {
		t.Fatalf("Expected output to contain 'failed request to create the tenant', got %q", output)
	}
	if !strings.Contains(output, "internal server error") {
		t.Fatalf("Expected output to contain 'internal server error', got %q", output)
	}
}

func TestTenantSubCmd_Create_Integration_Inline_ConfigError(t *testing.T) {
	mockCfg := configuration_handler.NewMockConfigurationHandler()
	mockCfg.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{}, fmt.Errorf("error while loading file path configuration")
	}

	tenantService := servicetenant.NewTenantService(mockCfg, &api.MockTenantAPI{}, &api.MockDomainAPI{}, &api.MockGatewayAPI{}, &api.MockProcessAPI{})
	tenantCmd, commandOutput := setupTenantCreateInlineCommand(tenantService)
	tenantCmd.SetArgs([]string{
		"create",
		"--name", "Tenant Cubbit 1",
		"--slug", "cubbit-1",
		"--connection", "domainID:gatewayID1",
	})

	err := tenantCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error from Execute (Run prints errors), got %v", err)
	}

	output := commandOutput.String()
	if !strings.Contains(output, "error while loading file path configuration") {
		t.Fatalf("Expected output to contain 'error while loading file path configuration', got %q", output)
	}
}
