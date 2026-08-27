package cmd_tenant

import (
	"bytes"
	"testing"
	"time"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	servicetenant "github.com/cubbit/composer-cli/src/service/tenant"
)

func TestTenantSubCmd_List_Output_JSON(t *testing.T) {
	mockCfg := newTestTenantConfig()
	frozen := time.Date(2026, 6, 25, 13, 54, 6, 0, time.UTC)

	mockTenantAPI := &api.MockTenantAPI{
		ListTenantsV5Func: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			_ ...api.ListTenantsV5Option,
		) (*api.GenericPaginatedResponse[api.TenantV5DTO], error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}

			return &api.GenericPaginatedResponse[api.TenantV5DTO]{
				Data: []api.TenantV5DTO{
					{
						ID:          "tnt-001",
						Name:        "prod-tenant",
						Slug:        "prod-tenant",
						Description: strPtr("Production tenant"),
						CreatedAt:   frozen,
						Storage:     &api.UsageDTO{Consumed: 1073741824, Reserved: 2147483648},
						Bandwidth:   &api.UsageDTO{Consumed: 524288000, Reserved: 1073741824},
				},
					{
						ID:        "tnt-002",
						Name:      "dev-tenant",
						Slug:      "dev-tenant",
						CreatedAt: frozen,
						Storage:   &api.UsageDTO{Consumed: 0, Reserved: 1073741824},
						Bandwidth: &api.UsageDTO{Consumed: 0, Reserved: 536870912},
				},
					{
						ID:        "tnt-003",
						Name:      "suspended-tenant",
						Slug:      "suspended-tenant",
						CreatedAt: frozen,
				},
				},
				NextPage: nil,
				Count:    3,
			}, nil
		},
	}

	tenantService := servicetenant.NewTenantService(
		mockCfg,
		mockTenantAPI,
		&api.MockDomainAPI{},
		&api.MockGatewayAPI{},
		&api.MockProcessAPI{},
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
		"list",
		"--output", "json",
	})

	err := tenantCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `[
  {
    "id": "tnt-001",
    "name": "prod-tenant",
    "slug": "prod-tenant",
    "description": "Production tenant",
    "created_at": "2026-06-25T13:54:06Z",
    "storage": {
      "consumed": 1073741824,
      "reserved": 2147483648,
      "percentage": 0
    },
    "bandwidth": {
      "consumed": 524288000,
      "reserved": 1073741824,
      "percentage": 0
    },
    "settings": null,
    "zk_enabled": false
  },
  {
    "id": "tnt-002",
    "name": "dev-tenant",
    "slug": "dev-tenant",
    "description": null,
    "created_at": "2026-06-25T13:54:06Z",
    "storage": {
      "consumed": 0,
      "reserved": 1073741824,
      "percentage": 0
    },
    "bandwidth": {
      "consumed": 0,
      "reserved": 536870912,
      "percentage": 0
    },
    "settings": null,
    "zk_enabled": false
  },
  {
    "id": "tnt-003",
    "name": "suspended-tenant",
    "slug": "suspended-tenant",
    "description": null,
    "created_at": "2026-06-25T13:54:06Z",
    "storage": null,
    "bandwidth": null,
    "settings": null,
    "zk_enabled": false
  }
]
`
	if commandOutput.String() != expected {
		t.Errorf("JSON snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

func TestTenantSubCmd_List_Output_YAML(t *testing.T) {
	mockCfg := newTestTenantConfig()
	frozen := time.Date(2026, 6, 25, 13, 54, 6, 0, time.UTC)

	mockTenantAPI := &api.MockTenantAPI{
		ListTenantsV5Func: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			_ ...api.ListTenantsV5Option,
		) (*api.GenericPaginatedResponse[api.TenantV5DTO], error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}

			return &api.GenericPaginatedResponse[api.TenantV5DTO]{
				Data: []api.TenantV5DTO{
					{
						ID:          "tnt-001",
						Name:        "prod-tenant",
						Slug:        "prod-tenant",
						Description: strPtr("Production tenant"),
						CreatedAt:   frozen,
						Storage:     &api.UsageDTO{Consumed: 1073741824, Reserved: 2147483648},
						Bandwidth:   &api.UsageDTO{Consumed: 524288000, Reserved: 1073741824},
				},
					{
						ID:        "tnt-002",
						Name:      "dev-tenant",
						Slug:      "dev-tenant",
						CreatedAt: frozen,
						Storage:   &api.UsageDTO{Consumed: 0, Reserved: 1073741824},
						Bandwidth: &api.UsageDTO{Consumed: 0, Reserved: 536870912},
				},
					{
						ID:        "tnt-003",
						Name:      "suspended-tenant",
						Slug:      "suspended-tenant",
						CreatedAt: frozen,
				},
				},
				NextPage: nil,
				Count:    3,
			}, nil
		},
	}

	tenantService := servicetenant.NewTenantService(
		mockCfg,
		mockTenantAPI,
		&api.MockDomainAPI{},
		&api.MockGatewayAPI{},
		&api.MockProcessAPI{},
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
		"list",
		"--output", "yaml",
	})

	err := tenantCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `- id: tnt-001
  name: prod-tenant
  slug: prod-tenant
  description: Production tenant
  created_at: 2026-06-25T13:54:06Z
  storage:
    consumed: 1073741824
    reserved: 2147483648
    percentage: 0
  bandwidth:
    consumed: 524288000
    reserved: 1073741824
    percentage: 0
  settings: null
  zk_enabled: false
- id: tnt-002
  name: dev-tenant
  slug: dev-tenant
  description: null
  created_at: 2026-06-25T13:54:06Z
  storage:
    consumed: 0
    reserved: 1073741824
    percentage: 0
  bandwidth:
    consumed: 0
    reserved: 536870912
    percentage: 0
  settings: null
  zk_enabled: false
- id: tnt-003
  name: suspended-tenant
  slug: suspended-tenant
  description: null
  created_at: 2026-06-25T13:54:06Z
  storage: null
  bandwidth: null
  settings: null
  zk_enabled: false
`
	if commandOutput.String() != expected {
		t.Errorf("YAML snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

func TestTenantSubCmd_List_Output_JSON_FromProfile(t *testing.T) {
	mockCfg := newTestTenantConfig()

	mockCfg.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{
			APIKey:         "test-api-key",
			OrganizationID: "test-org-id",
			Output:         "json",
			Endpoints: configuration_models.EndpointsV2{},
			}, nil
	}
	frozen := time.Date(2026, 6, 25, 13, 54, 6, 0, time.UTC)

	mockTenantAPI := &api.MockTenantAPI{
		ListTenantsV5Func: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			_ ...api.ListTenantsV5Option,
		) (*api.GenericPaginatedResponse[api.TenantV5DTO], error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}

			return &api.GenericPaginatedResponse[api.TenantV5DTO]{
				Data: []api.TenantV5DTO{
					{
						ID:          "tnt-001",
						Name:        "prod-tenant",
						Slug:        "prod-tenant",
						Description: strPtr("Production tenant"),
						CreatedAt:   frozen,
						Storage:     &api.UsageDTO{Consumed: 1073741824, Reserved: 2147483648},
						Bandwidth:   &api.UsageDTO{Consumed: 524288000, Reserved: 1073741824},
				},
					{
						ID:        "tnt-002",
						Name:      "dev-tenant",
						Slug:      "dev-tenant",
						CreatedAt: frozen,
						Storage:   &api.UsageDTO{Consumed: 0, Reserved: 1073741824},
						Bandwidth: &api.UsageDTO{Consumed: 0, Reserved: 536870912},
				},
					{
						ID:        "tnt-003",
						Name:      "suspended-tenant",
						Slug:      "suspended-tenant",
						CreatedAt: frozen,
				},
				},
				NextPage: nil,
				Count:    3,
			}, nil
		},
	}

	tenantService := servicetenant.NewTenantService(
		mockCfg,
		mockTenantAPI,
		&api.MockDomainAPI{},
		&api.MockGatewayAPI{},
		&api.MockProcessAPI{},
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
		"list",
	})

	err := tenantCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `[
  {
    "id": "tnt-001",
    "name": "prod-tenant",
    "slug": "prod-tenant",
    "description": "Production tenant",
    "created_at": "2026-06-25T13:54:06Z",
    "storage": {
      "consumed": 1073741824,
      "reserved": 2147483648,
      "percentage": 0
    },
    "bandwidth": {
      "consumed": 524288000,
      "reserved": 1073741824,
      "percentage": 0
    },
    "settings": null,
    "zk_enabled": false
  },
  {
    "id": "tnt-002",
    "name": "dev-tenant",
    "slug": "dev-tenant",
    "description": null,
    "created_at": "2026-06-25T13:54:06Z",
    "storage": {
      "consumed": 0,
      "reserved": 1073741824,
      "percentage": 0
    },
    "bandwidth": {
      "consumed": 0,
      "reserved": 536870912,
      "percentage": 0
    },
    "settings": null,
    "zk_enabled": false
  },
  {
    "id": "tnt-003",
    "name": "suspended-tenant",
    "slug": "suspended-tenant",
    "description": null,
    "created_at": "2026-06-25T13:54:06Z",
    "storage": null,
    "bandwidth": null,
    "settings": null,
    "zk_enabled": false
  }
]
`
	if commandOutput.String() != expected {
		t.Errorf("JSON snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

func TestTenantSubCmd_List_Output_YAML_FromProfile(t *testing.T) {
	mockCfg := newTestTenantConfig()

	mockCfg.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{
			APIKey:         "test-api-key",
			OrganizationID: "test-org-id",
			Output:         "yaml",
			Endpoints: configuration_models.EndpointsV2{},
			}, nil
	}
	frozen := time.Date(2026, 6, 25, 13, 54, 6, 0, time.UTC)

	mockTenantAPI := &api.MockTenantAPI{
		ListTenantsV5Func: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			_ ...api.ListTenantsV5Option,
		) (*api.GenericPaginatedResponse[api.TenantV5DTO], error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}

			return &api.GenericPaginatedResponse[api.TenantV5DTO]{
				Data: []api.TenantV5DTO{
					{
						ID:          "tnt-001",
						Name:        "prod-tenant",
						Slug:        "prod-tenant",
						Description: strPtr("Production tenant"),
						CreatedAt:   frozen,
						Storage:     &api.UsageDTO{Consumed: 1073741824, Reserved: 2147483648},
						Bandwidth:   &api.UsageDTO{Consumed: 524288000, Reserved: 1073741824},
				},
					{
						ID:        "tnt-002",
						Name:      "dev-tenant",
						Slug:      "dev-tenant",
						CreatedAt: frozen,
						Storage:   &api.UsageDTO{Consumed: 0, Reserved: 1073741824},
						Bandwidth: &api.UsageDTO{Consumed: 0, Reserved: 536870912},
				},
					{
						ID:        "tnt-003",
						Name:      "suspended-tenant",
						Slug:      "suspended-tenant",
						CreatedAt: frozen,
				},
				},
				NextPage: nil,
				Count:    3,
			}, nil
		},
	}

	tenantService := servicetenant.NewTenantService(
		mockCfg,
		mockTenantAPI,
		&api.MockDomainAPI{},
		&api.MockGatewayAPI{},
		&api.MockProcessAPI{},
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
		"list",
	})

	err := tenantCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `- id: tnt-001
  name: prod-tenant
  slug: prod-tenant
  description: Production tenant
  created_at: 2026-06-25T13:54:06Z
  storage:
    consumed: 1073741824
    reserved: 2147483648
    percentage: 0
  bandwidth:
    consumed: 524288000
    reserved: 1073741824
    percentage: 0
  settings: null
  zk_enabled: false
- id: tnt-002
  name: dev-tenant
  slug: dev-tenant
  description: null
  created_at: 2026-06-25T13:54:06Z
  storage:
    consumed: 0
    reserved: 1073741824
    percentage: 0
  bandwidth:
    consumed: 0
    reserved: 536870912
    percentage: 0
  settings: null
  zk_enabled: false
- id: tnt-003
  name: suspended-tenant
  slug: suspended-tenant
  description: null
  created_at: 2026-06-25T13:54:06Z
  storage: null
  bandwidth: null
  settings: null
  zk_enabled: false
`
	if commandOutput.String() != expected {
		t.Errorf("YAML snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}
