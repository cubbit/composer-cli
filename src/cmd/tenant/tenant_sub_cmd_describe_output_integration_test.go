package cmd_tenant

import (
	"bytes"
	"testing"
	"time"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	servicetenant "github.com/cubbit/composer-cli/src/service/tenant"
)

func TestTenantSubCmd_Describe_Output_JSON(t *testing.T) {
	mockCfg := newTestTenantConfig()

	frozen := time.Date(2026, 6, 24, 20, 24, 13, 0, time.UTC)
	mockTenantAPI := &api.MockTenantAPI{
		GetTenantV5Func: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			tenantID string,
		) (*api.TenantV5DTO, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}
			if tenantID != "tnt-out" {
				t.Fatalf("Expected tenant ID tnt-out, got %q", tenantID)
			}

			return &api.TenantV5DTO{
				ID:          "tnt-out",
				Name:        "output-test-tenant",
				Slug:        "output-test-tenant",
				Description: strPtr("JSON/YAML output test"),
				CreatedAt:   frozen,
				Storage:     &api.UsageDTO{Consumed: 1073741824, Reserved: 2147483648, Percentage: 0.5},
				Bandwidth:   &api.UsageDTO{Consumed: 524288000, Reserved: 1073741824, Percentage: 0.25},
				Settings: &api.TenantSettings{
					ConsoleUrl:        strPtr("https://console.example.com"),
					GatewayUrl:        strPtr("https://gw.example.com"),
					SignupDisabled:    boolPtr(false),
					WhitelabelEnabled: boolPtr(true),
				},
			}, nil
		},
	}

	tenantService := 	servicetenant.NewTenantService(
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

	commandOutput := new(bytes.Buffer)
	tenantCmd.SetOut(commandOutput)
	tenantCmd.SetErr(commandOutput)
	tenantCmd.SetArgs([]string{
		"describe",
		"tnt-out",
		"--output", "json",
	})

	err := tenantCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `{
  "id": "tnt-out",
  "name": "output-test-tenant",
  "slug": "output-test-tenant",
  "description": "JSON/YAML output test",
  "created_at": "2026-06-24T20:24:13Z",
  "storage": {
    "consumed": 1073741824,
    "reserved": 2147483648,
    "percentage": 0.5
  },
  "bandwidth": {
    "consumed": 524288000,
    "reserved": 1073741824,
    "percentage": 0.25
  },
  "settings": {
    "console_url": "https://console.example.com",
    "gateway_url": "https://gw.example.com",
    "project": null,
    "account": null,
    "allowed_domains": null,
    "blocked_domains": null,
    "notifications": null,
    "signup_disabled": false,
    "support_link": null,
    "display_name": null,
    "whitelabel_enabled": true,
    "whitelabel_assets": null,
    "white_label": null
  },
  "zk_enabled": false
}
`
	if commandOutput.String() != expected {
		t.Errorf("JSON snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

func TestTenantSubCmd_Describe_Output_YAML(t *testing.T) {
	mockCfg := newTestTenantConfig()

	frozen := time.Date(2026, 6, 24, 20, 24, 13, 0, time.UTC)
	mockTenantAPI := &api.MockTenantAPI{
		GetTenantV5Func: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			tenantID string,
		) (*api.TenantV5DTO, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}
			if tenantID != "tnt-yaml" {
				t.Fatalf("Expected tenant ID tnt-yaml, got %q", tenantID)
			}

			return &api.TenantV5DTO{
				ID:          "tnt-yaml",
				Name:        "yaml-test-tenant",
				Slug:        "yaml-test-tenant",
				Description: strPtr("YAML output test"),
				CreatedAt:   frozen,
				Storage:     &api.UsageDTO{Consumed: 1073741824, Reserved: 2147483648, Percentage: 0.5},
				Bandwidth:   &api.UsageDTO{Consumed: 524288000, Reserved: 1073741824, Percentage: 0.25},
				Settings: &api.TenantSettings{
					ConsoleUrl:        strPtr("https://console.example.com"),
					GatewayUrl:        strPtr("https://gw.example.com"),
					SignupDisabled:    boolPtr(false),
					WhitelabelEnabled: boolPtr(true),
				},
			}, nil
		},
	}

	tenantService := 	servicetenant.NewTenantService(
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

	commandOutput := new(bytes.Buffer)
	tenantCmd.SetOut(commandOutput)
	tenantCmd.SetErr(commandOutput)
	tenantCmd.SetArgs([]string{
		"describe",
		"tnt-yaml",
		"--output", "yaml",
	})

	err := tenantCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `id: tnt-yaml
name: yaml-test-tenant
slug: yaml-test-tenant
description: YAML output test
created_at: 2026-06-24T20:24:13Z
storage:
    consumed: 1073741824
    reserved: 2147483648
    percentage: 0.5
bandwidth:
    consumed: 524288000
    reserved: 1073741824
    percentage: 0.25
settings:
    console_url: https://console.example.com
    gateway_url: https://gw.example.com
    project: null
    account: null
    allowed_domains: null
    blocked_domains: null
    notifications: null
    signup_disabled: false
    support_link: null
    display_name: null
    whitelabel_enabled: true
    whitelabel_assets: null
    white_label: null
zk_enabled: false
`
	if commandOutput.String() != expected {
		t.Errorf("YAML snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

func TestTenantSubCmd_Describe_Output_JSON_FromProfile(t *testing.T) {
	mockCfg := newTestTenantConfig()

	mockCfg.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{
			APIKey:         "test-api-key",
			OrganizationID: "test-org-id",
			Output:         "json",
			Endpoints: configuration_models.EndpointsV2{},
			}, nil
	}

	frozen := time.Date(2026, 6, 24, 20, 24, 13, 0, time.UTC)
	mockTenantAPI := &api.MockTenantAPI{
		GetTenantV5Func: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			tenantID string,
		) (*api.TenantV5DTO, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}
			if tenantID != "tnt-out" {
				t.Fatalf("Expected tenant ID tnt-out, got %q", tenantID)
			}

			return &api.TenantV5DTO{
				ID:          "tnt-out",
				Name:        "output-test-tenant",
				Slug:        "output-test-tenant",
				Description: strPtr("JSON/YAML output test"),
				CreatedAt:   frozen,
				Storage:     &api.UsageDTO{Consumed: 1073741824, Reserved: 2147483648, Percentage: 0.5},
				Bandwidth:   &api.UsageDTO{Consumed: 524288000, Reserved: 1073741824, Percentage: 0.25},
				Settings: &api.TenantSettings{
					ConsoleUrl:        strPtr("https://console.example.com"),
					GatewayUrl:        strPtr("https://gw.example.com"),
					SignupDisabled:    boolPtr(false),
					WhitelabelEnabled: boolPtr(true),
				},
			}, nil
		},
	}

	tenantService := 	servicetenant.NewTenantService(
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

	commandOutput := new(bytes.Buffer)
	tenantCmd.SetOut(commandOutput)
	tenantCmd.SetErr(commandOutput)
	tenantCmd.SetArgs([]string{
		"describe",
		"tnt-out",
	})

	err := tenantCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `{
  "id": "tnt-out",
  "name": "output-test-tenant",
  "slug": "output-test-tenant",
  "description": "JSON/YAML output test",
  "created_at": "2026-06-24T20:24:13Z",
  "storage": {
    "consumed": 1073741824,
    "reserved": 2147483648,
    "percentage": 0.5
  },
  "bandwidth": {
    "consumed": 524288000,
    "reserved": 1073741824,
    "percentage": 0.25
  },
  "settings": {
    "console_url": "https://console.example.com",
    "gateway_url": "https://gw.example.com",
    "project": null,
    "account": null,
    "allowed_domains": null,
    "blocked_domains": null,
    "notifications": null,
    "signup_disabled": false,
    "support_link": null,
    "display_name": null,
    "whitelabel_enabled": true,
    "whitelabel_assets": null,
    "white_label": null
  },
  "zk_enabled": false
}
`
	if commandOutput.String() != expected {
		t.Errorf("JSON snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

func TestTenantSubCmd_Describe_Output_YAML_FromProfile(t *testing.T) {
	mockCfg := newTestTenantConfig()

	mockCfg.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{
			APIKey:         "test-api-key",
			OrganizationID: "test-org-id",
			Output:         "yaml",
			Endpoints: configuration_models.EndpointsV2{},
			}, nil
	}

	frozen := time.Date(2026, 6, 24, 20, 24, 13, 0, time.UTC)
	mockTenantAPI := &api.MockTenantAPI{
		GetTenantV5Func: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			tenantID string,
		) (*api.TenantV5DTO, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}
			if tenantID != "tnt-yaml" {
				t.Fatalf("Expected tenant ID tnt-yaml, got %q", tenantID)
			}

			return &api.TenantV5DTO{
				ID:          "tnt-yaml",
				Name:        "yaml-test-tenant",
				Slug:        "yaml-test-tenant",
				Description: strPtr("YAML output test"),
				CreatedAt:   frozen,
				Storage:     &api.UsageDTO{Consumed: 1073741824, Reserved: 2147483648, Percentage: 0.5},
				Bandwidth:   &api.UsageDTO{Consumed: 524288000, Reserved: 1073741824, Percentage: 0.25},
				Settings: &api.TenantSettings{
					ConsoleUrl:        strPtr("https://console.example.com"),
					GatewayUrl:        strPtr("https://gw.example.com"),
					SignupDisabled:    boolPtr(false),
					WhitelabelEnabled: boolPtr(true),
				},
			}, nil
		},
	}

	tenantService := 	servicetenant.NewTenantService(
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

	commandOutput := new(bytes.Buffer)
	tenantCmd.SetOut(commandOutput)
	tenantCmd.SetErr(commandOutput)
	tenantCmd.SetArgs([]string{
		"describe",
		"tnt-yaml",
	})

	err := tenantCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `id: tnt-yaml
name: yaml-test-tenant
slug: yaml-test-tenant
description: YAML output test
created_at: 2026-06-24T20:24:13Z
storage:
    consumed: 1073741824
    reserved: 2147483648
    percentage: 0.5
bandwidth:
    consumed: 524288000
    reserved: 1073741824
    percentage: 0.25
settings:
    console_url: https://console.example.com
    gateway_url: https://gw.example.com
    project: null
    account: null
    allowed_domains: null
    blocked_domains: null
    notifications: null
    signup_disabled: false
    support_link: null
    display_name: null
    whitelabel_enabled: true
    whitelabel_assets: null
    white_label: null
zk_enabled: false
`
	if commandOutput.String() != expected {
		t.Errorf("YAML snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}
