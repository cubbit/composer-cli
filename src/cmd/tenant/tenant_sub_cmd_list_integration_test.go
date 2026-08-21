package cmd_tenant

import (
	"bytes"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	servicetenant "github.com/cubbit/composer-cli/src/service/tenant"
)

func TestTenantSubCmd_List_Integration_Success(t *testing.T) {
	mockCfg := newTestTenantConfig()

	mockTenantAPI := &api.MockTenantAPI{
		ListTenantsV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			opts ...api.ListTenantsV5Option,
		) (*api.GenericPaginatedResponse[api.TenantV5DTO], error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}

			options := &api.ListTenantsV5Options{Page: 1, Items: 100}
			for _, opt := range opts {
				opt(options)
			}

			if options.Page != 1 {
				t.Fatalf("Expected page 1, got %d", options.Page)
			}
			if options.Items != 100 {
				t.Fatalf("Expected items 100, got %d", options.Items)
			}

			frozen := time.Date(2026, 6, 25, 13, 54, 6, 0, time.UTC)

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

	expectedResult := strings.TrimSpace(`
╭─────────┬──────────────────┬──────────────────┬───────────────────┬─────────────────────╮
│ ID      │ Name             │ Slug             │ Description       │ Created At          │
├─────────┼──────────────────┼──────────────────┼───────────────────┼─────────────────────┤
│ tnt-001 │ prod-tenant      │ prod-tenant      │ Production tenant │ 2026-06-25 13:54:06 │
│ tnt-002 │ dev-tenant       │ dev-tenant       │                   │ 2026-06-25 13:54:06 │
│ tnt-003 │ suspended-tenant │ suspended-tenant │                   │ 2026-06-25 13:54:06 │
╰─────────┴──────────────────┴──────────────────┴───────────────────┴─────────────────────╯
`)

	actualResult := strings.TrimSpace(commandOutput.String())
	if actualResult != expectedResult {
		t.Errorf("Expected tenant list output does not match actual output\nExpected:\n%s\nActual:\n%s", expectedResult, actualResult)
	}
}

func TestTenantSubCmd_List_Integration_WithPagination(t *testing.T) {
	mockCfg := newTestTenantConfig()

	mockTenantAPI := &api.MockTenantAPI{
		ListTenantsV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			opts ...api.ListTenantsV5Option,
		) (*api.GenericPaginatedResponse[api.TenantV5DTO], error) {
			options := &api.ListTenantsV5Options{}
			for _, opt := range opts {
				opt(options)
			}

			if options.Page != 2 {
				t.Fatalf("Expected page 2, got %d", options.Page)
			}
			if options.Items != 50 {
				t.Fatalf("Expected items 50, got %d", options.Items)
			}

			frozen := time.Date(2026, 6, 25, 13, 54, 6, 0, time.UTC)

			return &api.GenericPaginatedResponse[api.TenantV5DTO]{
				Data: []api.TenantV5DTO{
					{
						ID:        "tnt-paged",
						Name:      "paged-tenant",
						Slug:      "paged-tenant",
						CreatedAt: frozen,
					},
				},
				NextPage: nil,
				Count:    1,
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
		"--page", "2",
		"--items", "50",
	})

	err := tenantCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedResult := strings.TrimSpace(`
╭───────────┬──────────────┬──────────────┬─────────────┬─────────────────────╮
│ ID        │ Name         │ Slug         │ Description │ Created At          │
├───────────┼──────────────┼──────────────┼─────────────┼─────────────────────┤
│ tnt-paged │ paged-tenant │ paged-tenant │             │ 2026-06-25 13:54:06 │
╰───────────┴──────────────┴──────────────┴─────────────┴─────────────────────╯
`)

	actualResult := strings.TrimSpace(commandOutput.String())
	if actualResult != expectedResult {
		t.Errorf("Expected tenant list output does not match actual output\nExpected:\n%s\nActual:\n%s", expectedResult, actualResult)
	}
}

func TestTenantSubCmd_List_Integration_WithFilter(t *testing.T) {
	mockCfg := newTestTenantConfig()

	mockTenantAPI := &api.MockTenantAPI{
		ListTenantsV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			opts ...api.ListTenantsV5Option,
		) (*api.GenericPaginatedResponse[api.TenantV5DTO], error) {
			options := &api.ListTenantsV5Options{}
			for _, opt := range opts {
				opt(options)
			}

			if options.Filter != "name:eq(test)" {
				t.Fatalf("Expected filter 'name:eq(test)', got %q", options.Filter)
			}

			frozen := time.Date(2026, 6, 25, 13, 54, 6, 0, time.UTC)

			return &api.GenericPaginatedResponse[api.TenantV5DTO]{
				Data: []api.TenantV5DTO{
					{
						ID:        "tnt-filtered",
						Name:      "test",
						Slug:      "test",
						CreatedAt: frozen,
					},
				},
				NextPage: nil,
				Count:    1,
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
		"--query", "name:eq(test)",
	})

	err := tenantCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedResult := strings.TrimSpace(`
╭──────────────┬──────┬──────┬─────────────┬─────────────────────╮
│ ID           │ Name │ Slug │ Description │ Created At          │
├──────────────┼──────┼──────┼─────────────┼─────────────────────┤
│ tnt-filtered │ test │ test │             │ 2026-06-25 13:54:06 │
╰──────────────┴──────┴──────┴─────────────┴─────────────────────╯
`)

	actualResult := strings.TrimSpace(commandOutput.String())
	if actualResult != expectedResult {
		t.Errorf("Expected tenant list output does not match actual output\nExpected:\n%s\nActual:\n%s", expectedResult, actualResult)
	}
}

func TestTenantSubCmd_List_Integration_Empty(t *testing.T) {
	mockCfg := newTestTenantConfig()

	mockTenantAPI := &api.MockTenantAPI{
		ListTenantsV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			opts ...api.ListTenantsV5Option,
		) (*api.GenericPaginatedResponse[api.TenantV5DTO], error) {
			return &api.GenericPaginatedResponse[api.TenantV5DTO]{
				Data:     []api.TenantV5DTO{},
				NextPage: nil,
				Count:    0,
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

	expectedResult := "No tenants found.\n"
	actualResult := commandOutput.String()
	if actualResult != expectedResult {
		t.Errorf("Expected output %q, got %q", expectedResult, actualResult)
	}
}

func TestTenantSubCmd_List_Integration_Error(t *testing.T) {
	mockCfg := newTestTenantConfig()

	mockTenantAPI := &api.MockTenantAPI{
		ListTenantsV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			opts ...api.ListTenantsV5Option,
		) (*api.GenericPaginatedResponse[api.TenantV5DTO], error) {
			return nil, errors.New("network error")
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
		t.Fatalf("Expected no error from Execute (Run doesn't return errors), got %v", err)
	}

	output := commandOutput.String()
	if !strings.Contains(output, "failed to list tenants") {
		t.Errorf("Expected output to contain error message, got %q", output)
	}
}
