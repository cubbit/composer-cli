package cmd_gateway

import (
	"bytes"
	"testing"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	servicegateway "github.com/cubbit/composer-cli/src/service/gateway"
)

func TestGatewaySubCmd_List_Output_JSON(t *testing.T) {
	mockCfg := newTestGatewayConfig()

	mockGatewayAPI := &api.MockGatewayAPI{
		ListGatewaysV5Func: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			_ ...api.ListGatewaysV5Option,
		) (*api.GenericPaginatedResponse[api.GatewayV5ListItemResponse], error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
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
				},
				NextPage: nil,
				Count:    2,
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
		"--output", "json",
	})

	err := gatewayCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `[
  {
    "id": "gw-001",
    "name": "prod-gateway",
    "slug": "prod-gateway",
    "type": "singlecluster",
    "number_of_swarms": 3,
    "number_of_tenants": 5,
    "status": "ready"
  },
  {
    "id": "gw-002",
    "name": "dev-gateway",
    "slug": "dev-gateway",
    "type": "manual",
    "number_of_swarms": 1,
    "number_of_tenants": 2,
    "status": "not-ready"
  }
]
`
	if commandOutput.String() != expected {
		t.Errorf("JSON snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

func TestGatewaySubCmd_List_Output_YAML(t *testing.T) {
	mockCfg := newTestGatewayConfig()

	mockGatewayAPI := &api.MockGatewayAPI{
		ListGatewaysV5Func: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			_ ...api.ListGatewaysV5Option,
		) (*api.GenericPaginatedResponse[api.GatewayV5ListItemResponse], error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
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
				},
				NextPage: nil,
				Count:    2,
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
		"--output", "yaml",
	})

	err := gatewayCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `- id: gw-001
  name: prod-gateway
  slug: prod-gateway
  type: singlecluster
  number_of_swarms: 3
  number_of_tenants: 5
  status: ready
- id: gw-002
  name: dev-gateway
  slug: dev-gateway
  type: manual
  number_of_swarms: 1
  number_of_tenants: 2
  status: not-ready
`
	if commandOutput.String() != expected {
		t.Errorf("YAML snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

func TestGatewaySubCmd_List_Output_JSON_FromProfile(t *testing.T) {
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
		ListGatewaysV5Func: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			_ ...api.ListGatewaysV5Option,
		) (*api.GenericPaginatedResponse[api.GatewayV5ListItemResponse], error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
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
				},
				NextPage: nil,
				Count:    2,
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

	expected := `[
  {
    "id": "gw-001",
    "name": "prod-gateway",
    "slug": "prod-gateway",
    "type": "singlecluster",
    "number_of_swarms": 3,
    "number_of_tenants": 5,
    "status": "ready"
  },
  {
    "id": "gw-002",
    "name": "dev-gateway",
    "slug": "dev-gateway",
    "type": "manual",
    "number_of_swarms": 1,
    "number_of_tenants": 2,
    "status": "not-ready"
  }
]
`
	if commandOutput.String() != expected {
		t.Errorf("JSON snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

func TestGatewaySubCmd_List_Output_YAML_FromProfile(t *testing.T) {
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
		ListGatewaysV5Func: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			_ ...api.ListGatewaysV5Option,
		) (*api.GenericPaginatedResponse[api.GatewayV5ListItemResponse], error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
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
				},
				NextPage: nil,
				Count:    2,
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

	expected := `- id: gw-001
  name: prod-gateway
  slug: prod-gateway
  type: singlecluster
  number_of_swarms: 3
  number_of_tenants: 5
  status: ready
- id: gw-002
  name: dev-gateway
  slug: dev-gateway
  type: manual
  number_of_swarms: 1
  number_of_tenants: 2
  status: not-ready
`
	if commandOutput.String() != expected {
		t.Errorf("YAML snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}
