package cmd_gateway

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration"
	servicegateway "github.com/cubbit/composer-cli/src/service/gateway"
)

func TestGatewaySubCmd_List_Integration_Success(t *testing.T) {
	mockCfg := api.NewMockConfig(configuration.ProfileTypeComposer, "test-api-key", "test-org-id")

	mockGatewayAPI := &api.MockGatewayAPI{
		ListGatewaysV5Func: func(
			urlConfig configuration.URLs,
			apiKey string,
			organizationID string,
			opts ...api.ListGatewaysV5Option,
		) (*api.GenericPaginatedResponse[api.GatewayV5ListItemResponse], error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}

			options := &api.ListGatewaysV5Options{Page: 1, Items: 100}
			for _, opt := range opts {
				opt(options)
			}

			if options.Page != 1 {
				t.Fatalf("Expected page 1, got %d", options.Page)
			}
			if options.Items != 100 {
				t.Fatalf("Expected items 100, got %d", options.Items)
			}

			return &api.GenericPaginatedResponse[api.GatewayV5ListItemResponse]{
				Data: []api.GatewayV5ListItemResponse{
					{
						ID:              "gw-001",
						Name:            "prod-gateway",
						Slug:            "prod-gateway",
						Type:            "singlecluster",
						NumberOfSwarms:  3,
						NumberOfTenants: 5,
						Status:          api.GatewayV5StatusReady,
					},
					{
						ID:              "gw-002",
						Name:            "dev-gateway",
						Slug:            "dev-gateway",
						Type:            "manual",
						NumberOfSwarms:  1,
						NumberOfTenants: 2,
						Status:          api.GatewayV5StatusNotReady,
					},
					{
						ID:              "gw-003",
						Name:            "staging-gateway",
						Slug:            "staging-gateway",
						Type:            "multicluster_worker",
						NumberOfSwarms:  5,
						NumberOfTenants: 10,
						Status:          api.GatewayV5StatusReady,
					},
					{
						ID:              "gw-004",
						Name:            "dr-gateway",
						Slug:            "dr-gateway",
						Type:            "multicluster_controller",
						NumberOfSwarms:  2,
						NumberOfTenants: 0,
						Status:          api.GatewayV5StatusNotReady,
					},
				},
				NextPage: nil,
				Count:    4,
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
		"list",
	})

	err := gatewayCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedResult := strings.TrimSpace(`
╭─────────────────┬─────────────────┬─────────────────────────┬────────┬─────────┬───────────╮
│ Name            │ Slug            │ Type                    │ Swarms │ Tenants │ Status    │
├─────────────────┼─────────────────┼─────────────────────────┼────────┼─────────┼───────────┤
│ prod-gateway    │ prod-gateway    │ singlecluster           │ 3      │ 5       │ ready     │
│ dev-gateway     │ dev-gateway     │ manual                  │ 1      │ 2       │ not-ready │
│ staging-gateway │ staging-gateway │ multicluster_worker     │ 5      │ 10      │ ready     │
│ dr-gateway      │ dr-gateway      │ multicluster_controller │ 2      │ 0       │ not-ready │
╰─────────────────┴─────────────────┴─────────────────────────┴────────┴─────────┴───────────╯
`)

	actualResult := strings.TrimSpace(commandOutput.String())
	if actualResult != expectedResult {
		t.Errorf("Expected gateway list output does not match actual output\nExpected:\n%s\nActual:\n%s", expectedResult, actualResult)
	}
}

func TestGatewaySubCmd_List_Integration_WithPagination(t *testing.T) {
	mockCfg := api.NewMockConfig(configuration.ProfileTypeComposer, "test-api-key", "test-org-id")

	mockGatewayAPI := &api.MockGatewayAPI{
		ListGatewaysV5Func: func(
			urlConfig configuration.URLs,
			apiKey string,
			organizationID string,
			opts ...api.ListGatewaysV5Option,
		) (*api.GenericPaginatedResponse[api.GatewayV5ListItemResponse], error) {
			options := &api.ListGatewaysV5Options{}
			for _, opt := range opts {
				opt(options)
			}

			if options.Page != 2 {
				t.Fatalf("Expected page 2, got %d", options.Page)
			}
			if options.Items != 50 {
				t.Fatalf("Expected items 50, got %d", options.Items)
			}

			return &api.GenericPaginatedResponse[api.GatewayV5ListItemResponse]{
				Data: []api.GatewayV5ListItemResponse{
					{
						ID:              "gw-paged",
						Name:            "paged-gateway",
						Slug:            "paged-gateway",
						Type:            "manual",
						NumberOfSwarms:  2,
						NumberOfTenants: 3,
						Status:          api.GatewayV5StatusReady,
					},
				},
				NextPage: nil,
				Count:    1,
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
		"list",
		"--page", "2",
		"--items", "50",
	})

	err := gatewayCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedResult := strings.TrimSpace(`
╭───────────────┬───────────────┬────────┬────────┬─────────┬────────╮
│ Name          │ Slug          │ Type   │ Swarms │ Tenants │ Status │
├───────────────┼───────────────┼────────┼────────┼─────────┼────────┤
│ paged-gateway │ paged-gateway │ manual │ 2      │ 3       │ ready  │
╰───────────────┴───────────────┴────────┴────────┴─────────┴────────╯
`)

	actualResult := strings.TrimSpace(commandOutput.String())
	if actualResult != expectedResult {
		t.Errorf("Expected gateway list output does not match actual output\nExpected:\n%s\nActual:\n%s", expectedResult, actualResult)
	}
}

func TestGatewaySubCmd_List_Integration_WithFilter(t *testing.T) {
	mockCfg := api.NewMockConfig(configuration.ProfileTypeComposer, "test-api-key", "test-org-id")

	mockGatewayAPI := &api.MockGatewayAPI{
		ListGatewaysV5Func: func(
			urlConfig configuration.URLs,
			apiKey string,
			organizationID string,
			opts ...api.ListGatewaysV5Option,
		) (*api.GenericPaginatedResponse[api.GatewayV5ListItemResponse], error) {
			options := &api.ListGatewaysV5Options{}
			for _, opt := range opts {
				opt(options)
			}

			if options.Filter != "name:eq(test)" {
				t.Fatalf("Expected filter 'name:eq(test)', got %q", options.Filter)
			}

			return &api.GenericPaginatedResponse[api.GatewayV5ListItemResponse]{
				Data: []api.GatewayV5ListItemResponse{
					{
						ID:              "gw-filtered",
						Name:            "test",
						Slug:            "test",
						Type:            "manual",
						NumberOfSwarms:  1,
						NumberOfTenants: 1,
						Status:          api.GatewayV5StatusReady,
					},
				},
				NextPage: nil,
				Count:    1,
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
		"list",
		"--query", "name:eq(test)",
	})

	err := gatewayCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedResult := strings.TrimSpace(`
╭──────┬──────┬────────┬────────┬─────────┬────────╮
│ Name │ Slug │ Type   │ Swarms │ Tenants │ Status │
├──────┼──────┼────────┼────────┼─────────┼────────┤
│ test │ test │ manual │ 1      │ 1       │ ready  │
╰──────┴──────┴────────┴────────┴─────────┴────────╯
`)

	actualResult := strings.TrimSpace(commandOutput.String())
	if actualResult != expectedResult {
		t.Errorf("Expected gateway list output does not match actual output\nExpected:\n%s\nActual:\n%s", expectedResult, actualResult)
	}
}

func TestGatewaySubCmd_List_Integration_Empty(t *testing.T) {
	mockCfg := api.NewMockConfig(configuration.ProfileTypeComposer, "test-api-key", "test-org-id")

	mockGatewayAPI := &api.MockGatewayAPI{
		ListGatewaysV5Func: func(
			urlConfig configuration.URLs,
			apiKey string,
			organizationID string,
			opts ...api.ListGatewaysV5Option,
		) (*api.GenericPaginatedResponse[api.GatewayV5ListItemResponse], error) {
			return &api.GenericPaginatedResponse[api.GatewayV5ListItemResponse]{
				Data:     []api.GatewayV5ListItemResponse{},
				NextPage: nil,
				Count:    0,
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
		"list",
	})

	err := gatewayCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedResult := "No gateways found.\n"
	actualResult := commandOutput.String()
	if actualResult != expectedResult {
		t.Errorf("Expected output %q, got %q", expectedResult, actualResult)
	}
}

func TestGatewaySubCmd_List_Integration_Error(t *testing.T) {
	mockCfg := api.NewMockConfig(configuration.ProfileTypeComposer, "test-api-key", "test-org-id")

	mockGatewayAPI := &api.MockGatewayAPI{
		ListGatewaysV5Func: func(
			urlConfig configuration.URLs,
			apiKey string,
			organizationID string,
			opts ...api.ListGatewaysV5Option,
		) (*api.GenericPaginatedResponse[api.GatewayV5ListItemResponse], error) {
			return nil, errors.New("network error")
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
		"list",
	})

	err := gatewayCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error from Execute (Run doesn't return errors), got %v", err)
	}

	output := commandOutput.String()
	if !strings.Contains(output, "failed to list gateways") {
		t.Errorf("Expected output to contain error message, got %q", output)
	}
}
