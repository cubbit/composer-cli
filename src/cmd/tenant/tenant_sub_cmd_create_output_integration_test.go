package cmd_tenant

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	servicetenant "github.com/cubbit/composer-cli/src/service/tenant"
)

func TestTenantSubCmd_Create_Output_Inline_JSON(t *testing.T) {
	mockCfg := newTestTenantConfig()

	mockTenantAPI := &api.MockTenantAPI{
		CreateTenantV5Func: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			_ *api.CreateTenantV5Request,
		) (*api.GenericIDResponseModel, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}

			return &api.GenericIDResponseModel{ID: "response-id"}, nil
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

			data, _ := json.Marshal(api.TenantCreationProcessData{TenantID: "tnt-created"})
			return &api.Process{
				ID:     "process-id",
				Type:   api.ProcessTypeTenantCreation,
				Status: api.ProcessStatusSuccess,
				Data:   data,
			}, nil
		},
	}

	tenantService := 	servicetenant.NewTenantService(
		mockCfg,
		mockTenantAPI,
		&api.MockDomainAPI{},
		&api.MockGatewayAPI{},
		mockProcessAPI,
		&api.MockConnectionAPI{},
	)

	tenantCmd := NewTenantCmd(tenantService)
	tenantCmd.PersistentFlags().String("profile", "", "Profile")
	tenantCmd.PersistentFlags().String("output", "human", "Output format")
	tenantCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	tenantCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	commandOutput := new(bytes.Buffer)
	tenantCmd.SetOut(commandOutput)
	tenantCmd.SetErr(commandOutput)
	tenantCmd.SetArgs([]string{
		"create",
		"--name", "Test",
		"--slug", "test",
		"--connection", "dom:gw",
		"--output", "json",
	})

	err := tenantCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `{
  "message": "Tenant creation started — Tenant ID: tnt-created\n"
}
`
	if commandOutput.String() != expected {
		t.Errorf("JSON snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

func TestTenantSubCmd_Create_Output_Inline_YAML(t *testing.T) {
	mockCfg := newTestTenantConfig()

	mockTenantAPI := &api.MockTenantAPI{
		CreateTenantV5Func: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			_ *api.CreateTenantV5Request,
		) (*api.GenericIDResponseModel, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}

			return &api.GenericIDResponseModel{ID: "response-id"}, nil
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

			data, _ := json.Marshal(api.TenantCreationProcessData{TenantID: "tnt-created"})
			return &api.Process{
				ID:     "process-id",
				Type:   api.ProcessTypeTenantCreation,
				Status: api.ProcessStatusSuccess,
				Data:   data,
			}, nil
		},
	}

	tenantService := 	servicetenant.NewTenantService(
		mockCfg,
		mockTenantAPI,
		&api.MockDomainAPI{},
		&api.MockGatewayAPI{},
		mockProcessAPI,
		&api.MockConnectionAPI{},
	)

	tenantCmd := NewTenantCmd(tenantService)
	tenantCmd.PersistentFlags().String("profile", "", "Profile")
	tenantCmd.PersistentFlags().String("output", "human", "Output format")
	tenantCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	tenantCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	commandOutput := new(bytes.Buffer)
	tenantCmd.SetOut(commandOutput)
	tenantCmd.SetErr(commandOutput)
	tenantCmd.SetArgs([]string{
		"create",
		"--name", "Test",
		"--slug", "test",
		"--connection", "dom:gw",
		"--output", "yaml",
	})

	err := tenantCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `message: |
    Tenant creation started — Tenant ID: tnt-created
`
	if commandOutput.String() != expected {
		t.Errorf("YAML snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

func TestTenantSubCmd_Create_InteractiveAndOutputConflict(t *testing.T) {
	mockCfg := newTestTenantConfig()

	mockTenantAPI := &api.MockTenantAPI{}

	tenantService := servicetenant.NewTenantService(
		mockCfg,
		mockTenantAPI,
		&api.MockDomainAPI{},
		&api.MockGatewayAPI{},
		&api.MockProcessAPI{},
		&api.MockConnectionAPI{},
	)

	tenantCmd := NewTenantCmd(tenantService)
	tenantCmd.PersistentFlags().String("profile", "", "Profile")
	tenantCmd.PersistentFlags().String("output", "human", "Output format")
	tenantCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	tenantCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	tenantCmd.SetArgs([]string{
		"create",
		"--interactive",
		"--output", "json",
	})

	err := tenantCmd.Execute()
	if err == nil {
		t.Fatal("Expected error for conflicting --interactive and --output flags, got nil")
	}

	if !strings.Contains(err.Error(), "mutually exclusive") {
		t.Errorf("Expected error message to mention flags are mutually exclusive, got %q", err.Error())
	}
}

func TestTenantSubCmd_Create_Output_Inline_JSON_FromProfile(t *testing.T) {
	mockCfg := newTestTenantConfig()

	mockCfg.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{
			APIKey:         "test-api-key",
			OrganizationID: "test-org-id",
			Output:         "json",
				Endpoints: configuration_models.EndpointsV2{},
		}, nil
	}

	mockTenantAPI := &api.MockTenantAPI{
		CreateTenantV5Func: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			_ *api.CreateTenantV5Request,
		) (*api.GenericIDResponseModel, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}

			return &api.GenericIDResponseModel{ID: "response-id"}, nil
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

			data, _ := json.Marshal(api.TenantCreationProcessData{TenantID: "tnt-created"})
			return &api.Process{
				ID:     "process-id",
				Type:   api.ProcessTypeTenantCreation,
				Status: api.ProcessStatusSuccess,
				Data:   data,
			}, nil
		},
	}

	tenantService := 	servicetenant.NewTenantService(
		mockCfg,
		mockTenantAPI,
		&api.MockDomainAPI{},
		&api.MockGatewayAPI{},
		mockProcessAPI,
		&api.MockConnectionAPI{},
	)

	tenantCmd := NewTenantCmd(tenantService)
	tenantCmd.PersistentFlags().String("profile", "", "Profile")
	tenantCmd.PersistentFlags().String("output", "human", "Output format")
	tenantCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	tenantCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	commandOutput := new(bytes.Buffer)
	tenantCmd.SetOut(commandOutput)
	tenantCmd.SetErr(commandOutput)
	tenantCmd.SetArgs([]string{
		"create",
		"--name", "Test",
		"--slug", "test",
		"--connection", "dom:gw",
	})

	err := tenantCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `{
  "message": "Tenant creation started — Tenant ID: tnt-created\n"
}
`
	if commandOutput.String() != expected {
		t.Errorf("JSON snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

func TestTenantSubCmd_Create_Output_Inline_YAML_FromProfile(t *testing.T) {
	mockCfg := newTestTenantConfig()

	mockCfg.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{
			APIKey:         "test-api-key",
			OrganizationID: "test-org-id",
			Output:         "yaml",
				Endpoints: configuration_models.EndpointsV2{},
		}, nil
	}

	mockTenantAPI := &api.MockTenantAPI{
		CreateTenantV5Func: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			_ *api.CreateTenantV5Request,
		) (*api.GenericIDResponseModel, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}

			return &api.GenericIDResponseModel{ID: "response-id"}, nil
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

			data, _ := json.Marshal(api.TenantCreationProcessData{TenantID: "tnt-created"})
			return &api.Process{
				ID:     "process-id",
				Type:   api.ProcessTypeTenantCreation,
				Status: api.ProcessStatusSuccess,
				Data:   data,
			}, nil
		},
	}

	tenantService := 	servicetenant.NewTenantService(
		mockCfg,
		mockTenantAPI,
		&api.MockDomainAPI{},
		&api.MockGatewayAPI{},
		mockProcessAPI,
		&api.MockConnectionAPI{},
	)

	tenantCmd := NewTenantCmd(tenantService)
	tenantCmd.PersistentFlags().String("profile", "", "Profile")
	tenantCmd.PersistentFlags().String("output", "human", "Output format")
	tenantCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	tenantCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	commandOutput := new(bytes.Buffer)
	tenantCmd.SetOut(commandOutput)
	tenantCmd.SetErr(commandOutput)
	tenantCmd.SetArgs([]string{
		"create",
		"--name", "Test",
		"--slug", "test",
		"--connection", "dom:gw",
	})

	err := tenantCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `message: |
    Tenant creation started — Tenant ID: tnt-created
`
	if commandOutput.String() != expected {
		t.Errorf("YAML snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}
