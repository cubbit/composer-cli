package cmd_tenant

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	servicetenant "github.com/cubbit/composer-cli/src/service/tenant"
	"github.com/cubbit/composer-cli/utils/interactive/interactive_tester"
)

func TestTenantSubCmd_Create_Interactive_Success(t *testing.T) {
	mockCfg := newTestTenantConfig()

	var capturedRequest *api.CreateTenantV5Request

	mockTenantAPI := &api.MockTenantAPI{
		CreateTenantV5Func: func(
			_ configuration_models.EndpointsV2, _ string, _ string, request *api.CreateTenantV5Request,
		) (*api.GenericIDResponseModel, error) {
			capturedRequest = request
			return &api.GenericIDResponseModel{ID: "test-process-id"}, nil
		},
	}

	mockDomainAPI := &api.MockDomainAPI{
		ListFunc: func(
			_ configuration_models.EndpointsV2, _ string, _ string, _ int, _ int,
		) (*api.GenericPaginatedResponse[api.DomainDTO], error) {
			now := time.Now()
			return &api.GenericPaginatedResponse[api.DomainDTO]{
				Data: []api.DomainDTO{
					{
						ID:             "domain-001",
						DomainName:     "example.com",
						CreatedAt:      now,
						VerifiedAt:     &now,
						OrganizationID: "test-org-id",
					},
					{
						ID:             "domain-002",
						DomainName:     "test.org",
						CreatedAt:      now,
						VerifiedAt:     &now,
						OrganizationID: "test-org-id",
					},
					{
						ID:             "domain-003",
						DomainName:     "unverified.dev",
						CreatedAt:      now,
						VerifiedAt:     nil,
						OrganizationID: "test-org-id",
					},
				},
			}, nil
		},
	}

	mockGatewayAPI := &api.MockGatewayAPI{
		ListGatewaysV5Func: func(
			_ configuration_models.EndpointsV2, _ string, _ string, _ ...api.ListGatewaysV5Option,
		) (*api.GenericPaginatedResponse[api.GatewayV5ListItemResponse], error) {
			return &api.GenericPaginatedResponse[api.GatewayV5ListItemResponse]{
				Data: []api.GatewayV5ListItemResponse{
					{ID: "gw-eu-001", Name: "gw-eu-west", Slug: "gw-eu-west"},
					{ID: "gw-us-002", Name: "gw-us-east", Slug: "gw-us-east"},
				},
			}, nil
		},
	}

	mockProcessAPI := &api.MockProcessAPI{
		GetProcessFunc: func(
			_ configuration_models.EndpointsV2, _ string, _ string, processID string,
		) (*api.Process, error) {
			data, _ := json.Marshal(api.TenantCreationProcessData{TenantID: "tenant-created-001"})
			return &api.Process{
				ID:     processID,
				Step:   api.ProcessStepCompleted,
				Status: api.ProcessStatusSuccess,
				Data:   data,
			}, nil
		},
	}

	mockConnectionAPI := &api.MockConnectionAPI{}
	tenantService := servicetenant.NewTenantService(mockCfg, mockTenantAPI, mockDomainAPI, mockGatewayAPI, mockProcessAPI, mockConnectionAPI)
	tenantCmd := NewTenantCmd(tenantService)
	h, err := interactive_tester.New(tenantCmd, []string{"create", "--interactive"})
	if err != nil {
		t.Fatalf("could not create PTY harness: %v", err)
	}
	h.Start()
	defer h.Close()

	h.ExpectMultipleT(t, []string{"Tenant name", "required"}, interactive_tester.WithTimeout(30*time.Second))
	h.WriteLineT(t, "test-tenant")

	h.ExpectMultipleT(t, []string{"Tenant slug", "required"})
	h.WriteLineT(t, "test-tenant-slug")

	h.ExpectT(t, "Description")
	h.WriteDataT(t, h.Enter)

	h.ExpectMultipleT(t, []string{
		"Select domains",
		"example.com (domain-001)",
		"test.org (domain-002)",
	})
	h.WriteDataT(t, h.Enter)
	h.ExpectT(t, "at least one option must be selected")
	h.WriteKeysT(t, h.Space, h.Down, h.Space, h.Enter)

	h.ExpectMultipleT(t, []string{
		"Select gateways for domain example.com",
		"gw-eu-west (gw-eu-001)",
		"gw-us-east (gw-us-002)",
	})
	h.WriteKeysT(t, h.Space, h.Enter)

	h.ExpectMultipleT(t, []string{
		"Subdomain for domain example.com",
	})
	h.WriteDataT(t, h.Enter)

	h.ExpectMultipleT(t, []string{
		"Select gateways for domain test.org",
		"gw-eu-west (gw-eu-001)",
		"gw-us-east (gw-us-002)",
	})
	h.WriteKeysT(t, h.Down, h.Space, h.Enter)

	h.ExpectMultipleT(t, []string{
		"Subdomain for domain test.org",
	})
	h.WriteDataT(t, h.Enter)

	h.ExpectT(t, "Tenant created successfully — Tenant ID: tenant-created-001", interactive_tester.WithTimeout(30*time.Second))

	if capturedRequest == nil {
		t.Fatal("expected tenant creation to be called")
	}
	if capturedRequest.Name != "test-tenant" {
		t.Fatalf("expected name 'test-tenant', got %q", capturedRequest.Name)
	}
	if capturedRequest.Slug != "test-tenant-slug" {
		t.Fatalf("expected slug 'test-tenant-slug', got %q", capturedRequest.Slug)
	}
	if capturedRequest.Description != nil {
		t.Fatalf("expected nil description, got %v", *capturedRequest.Description)
	}
	if len(capturedRequest.Connections) != 2 {
		t.Fatalf("expected 2 connections, got %d", len(capturedRequest.Connections))
	}
	if capturedRequest.Connections[0].DomainID != "domain-001" {
		t.Fatalf("expected first domain 'domain-001', got %q", capturedRequest.Connections[0].DomainID)
	}
	if len(capturedRequest.Connections[0].GatewayIDs) != 1 || capturedRequest.Connections[0].GatewayIDs[0] != "gw-eu-001" {
		t.Fatalf("expected first connection gateway 'gw-eu-001', got %v", capturedRequest.Connections[0].GatewayIDs)
	}
	if capturedRequest.Connections[1].DomainID != "domain-002" {
		t.Fatalf("expected second domain 'domain-002', got %q", capturedRequest.Connections[1].DomainID)
	}
	if len(capturedRequest.Connections[1].GatewayIDs) != 1 || capturedRequest.Connections[1].GatewayIDs[0] != "gw-us-002" {
		t.Fatalf("expected second connection gateway 'gw-us-002', got %v", capturedRequest.Connections[1].GatewayIDs)
	}
}

func TestTenantSubCmd_Create_Interactive_WithDescription(t *testing.T) {
	mockCfg := newTestTenantConfig()

	mockTenantAPI := &api.MockTenantAPI{
		CreateTenantV5Func: func(
			_ configuration_models.EndpointsV2, _ string, _ string, request *api.CreateTenantV5Request,
		) (*api.GenericIDResponseModel, error) {
			return &api.GenericIDResponseModel{ID: "test-process-id"}, nil
		},
	}

	mockDomainAPI := &api.MockDomainAPI{
		ListFunc: func(
			_ configuration_models.EndpointsV2, _ string, _ string, _ int, _ int,
		) (*api.GenericPaginatedResponse[api.DomainDTO], error) {
			now := time.Now()
			return &api.GenericPaginatedResponse[api.DomainDTO]{
				Data: []api.DomainDTO{
					{
						ID:             "domain-001",
						DomainName:     "example.com",
						CreatedAt:      now,
						VerifiedAt:     &now,
						OrganizationID: "test-org-id",
					},
				},
			}, nil
		},
	}

	mockGatewayAPI := &api.MockGatewayAPI{
		ListGatewaysV5Func: func(
			_ configuration_models.EndpointsV2, _ string, _ string, _ ...api.ListGatewaysV5Option,
		) (*api.GenericPaginatedResponse[api.GatewayV5ListItemResponse], error) {
			return &api.GenericPaginatedResponse[api.GatewayV5ListItemResponse]{
				Data: []api.GatewayV5ListItemResponse{
					{ID: "gw-eu-001", Name: "gw-eu-west", Slug: "gw-eu-west"},
				},
			}, nil
		},
	}

	mockProcessAPI := &api.MockProcessAPI{
		GetProcessFunc: func(
			_ configuration_models.EndpointsV2, _ string, _ string, processID string,
		) (*api.Process, error) {
			data, _ := json.Marshal(api.TenantCreationProcessData{TenantID: "tenant-with-desc"})
			return &api.Process{
				ID:     processID,
				Step:   api.ProcessStepCompleted,
				Status: api.ProcessStatusSuccess,
				Data:   data,
			}, nil
		},
	}

	mockConnectionAPI := &api.MockConnectionAPI{}
	tenantService := servicetenant.NewTenantService(mockCfg, mockTenantAPI, mockDomainAPI, mockGatewayAPI, mockProcessAPI, mockConnectionAPI)
	tenantCmd := NewTenantCmd(tenantService)
	h, err := interactive_tester.New(tenantCmd, []string{"create", "--interactive"})
	if err != nil {
		t.Fatalf("could not create PTY harness: %v", err)
	}
	h.Start()
	defer h.Close()

	h.ExpectMultipleT(t, []string{"Tenant name", "required"}, interactive_tester.WithTimeout(30*time.Second))
	h.WriteLineT(t, "test-tenant")

	h.ExpectMultipleT(t, []string{"Tenant slug", "required"})
	h.WriteLineT(t, "test-tenant-slug")

	h.ExpectT(t, "Description")
	h.WriteLineT(t, "Production tenant")

	h.ExpectMultipleT(t, []string{
		"Select domains",
		"example.com (domain-001)",
	})
	h.WriteKeysT(t, h.Space, h.Enter)

	h.ExpectMultipleT(t, []string{
		"Select gateways for domain example.com",
		"gw-eu-west (gw-eu-001)",
	})
	h.WriteKeysT(t, h.Space, h.Enter)

	h.ExpectMultipleT(t, []string{
		"Subdomain for domain example.com",
	})
	h.WriteDataT(t, h.Enter)

	h.ExpectT(t, "Tenant created successfully — Tenant ID: tenant-with-desc", interactive_tester.WithTimeout(30*time.Second))
}

func TestTenantSubCmd_Create_Interactive_Subdomain(t *testing.T) {
	mockCfg := newTestTenantConfig()

	var capturedRequest *api.CreateTenantV5Request

	mockTenantAPI := &api.MockTenantAPI{
		CreateTenantV5Func: func(
			_ configuration_models.EndpointsV2, _ string, _ string, request *api.CreateTenantV5Request,
		) (*api.GenericIDResponseModel, error) {
			capturedRequest = request
			return &api.GenericIDResponseModel{ID: "test-process-id"}, nil
		},
	}

	mockDomainAPI := &api.MockDomainAPI{
		ListFunc: func(
			_ configuration_models.EndpointsV2, _ string, _ string, _ int, _ int,
		) (*api.GenericPaginatedResponse[api.DomainDTO], error) {
			now := time.Now()
			return &api.GenericPaginatedResponse[api.DomainDTO]{
				Data: []api.DomainDTO{
					{
						ID:             "domain-001",
						DomainName:     "example.com",
						CreatedAt:      now,
						VerifiedAt:     &now,
						OrganizationID: "test-org-id",
					},
				},
			}, nil
		},
	}

	mockGatewayAPI := &api.MockGatewayAPI{
		ListGatewaysV5Func: func(
			_ configuration_models.EndpointsV2, _ string, _ string, _ ...api.ListGatewaysV5Option,
		) (*api.GenericPaginatedResponse[api.GatewayV5ListItemResponse], error) {
			return &api.GenericPaginatedResponse[api.GatewayV5ListItemResponse]{
				Data: []api.GatewayV5ListItemResponse{
					{ID: "gw-eu-001", Name: "gw-eu-west", Slug: "gw-eu-west"},
				},
			}, nil
		},
	}

	mockProcessAPI := &api.MockProcessAPI{
		GetProcessFunc: func(
			_ configuration_models.EndpointsV2, _ string, _ string, processID string,
		) (*api.Process, error) {
			data, _ := json.Marshal(api.TenantCreationProcessData{TenantID: "tenant-with-sub"})
			return &api.Process{
				ID:     processID,
				Step:   api.ProcessStepCompleted,
				Status: api.ProcessStatusSuccess,
				Data:   data,
			}, nil
		},
	}

	mockConnectionAPI := &api.MockConnectionAPI{}
	tenantService := servicetenant.NewTenantService(mockCfg, mockTenantAPI, mockDomainAPI, mockGatewayAPI, mockProcessAPI, mockConnectionAPI)
	tenantCmd := NewTenantCmd(tenantService)
	h, err := interactive_tester.New(tenantCmd, []string{"create", "--interactive"})
	if err != nil {
		t.Fatalf("could not create PTY harness: %v", err)
	}
	h.Start()
	defer h.Close()

	h.ExpectMultipleT(t, []string{"Tenant name", "required"}, interactive_tester.WithTimeout(30*time.Second))
	h.WriteLineT(t, "test-tenant")

	h.ExpectMultipleT(t, []string{"Tenant slug", "required"})
	h.WriteLineT(t, "test-tenant-slug")

	h.ExpectT(t, "Description")
	h.WriteDataT(t, h.Enter)

	h.ExpectMultipleT(t, []string{
		"Select domains",
		"example.com (domain-001)",
	})
	h.WriteKeysT(t, h.Space, h.Enter)

	h.ExpectMultipleT(t, []string{
		"Select gateways for domain example.com",
		"gw-eu-west (gw-eu-001)",
	})
	h.WriteKeysT(t, h.Space, h.Enter)

	h.ExpectMultipleT(t, []string{
		"Subdomain for domain example.com",
	})
	h.WriteLineT(t, "my-subdomain")

	h.ExpectT(t, "Tenant created successfully — Tenant ID: tenant-with-sub", interactive_tester.WithTimeout(30*time.Second))

	if capturedRequest == nil {
		t.Fatal("expected tenant creation to be called")
	}
	if capturedRequest.Connections[0].Subdomain == nil || *capturedRequest.Connections[0].Subdomain != "my-subdomain" {
		t.Fatalf("expected subdomain 'my-subdomain', got %v", capturedRequest.Connections[0].Subdomain)
	}
}

func TestTenantSubCmd_Create_Interactive_DomainAPIError(t *testing.T) {
	mockCfg := newTestTenantConfig()

	mockTenantAPI := &api.MockTenantAPI{}
	mockDomainAPI := &api.MockDomainAPI{
		ListFunc: func(
			_ configuration_models.EndpointsV2, _ string, _ string, _ int, _ int,
		) (*api.GenericPaginatedResponse[api.DomainDTO], error) {
			return nil, fmt.Errorf("connection refused")
		},
	}
	mockGatewayAPI := &api.MockGatewayAPI{}
	mockProcessAPI := &api.MockProcessAPI{}
	mockConnectionAPI := &api.MockConnectionAPI{}

	tenantService := servicetenant.NewTenantService(mockCfg, mockTenantAPI, mockDomainAPI, mockGatewayAPI, mockProcessAPI, mockConnectionAPI)
	tenantCmd := NewTenantCmd(tenantService)
	h, err := interactive_tester.New(tenantCmd, []string{"create", "--interactive"})
	if err != nil {
		t.Fatalf("could not create PTY harness: %v", err)
	}
	h.Start()
	defer h.Close()

	h.ExpectMultipleT(t, []string{"Tenant name", "required"}, interactive_tester.WithTimeout(30*time.Second))
	h.WriteLineT(t, "test-tenant")

	h.ExpectMultipleT(t, []string{"Tenant slug", "required"})
	h.WriteLineT(t, "test-tenant-slug")

	h.ExpectT(t, "Description")
	h.WriteDataT(t, h.Enter)

	h.ExpectT(t, "connection refused", interactive_tester.WithTimeout(15*time.Second))
}

func TestTenantSubCmd_Create_Interactive_NoVerifiedDomains(t *testing.T) {
	mockCfg := newTestTenantConfig()

	mockTenantAPI := &api.MockTenantAPI{}
	mockDomainAPI := &api.MockDomainAPI{
		ListFunc: func(
			_ configuration_models.EndpointsV2, _ string, _ string, _ int, _ int,
		) (*api.GenericPaginatedResponse[api.DomainDTO], error) {
			now := time.Now()
			return &api.GenericPaginatedResponse[api.DomainDTO]{
				Data: []api.DomainDTO{
					{
						ID:             "domain-001",
						DomainName:     "unverified.dev",
						CreatedAt:      now,
						VerifiedAt:     nil,
						OrganizationID: "test-org-id",
					},
				},
			}, nil
		},
	}
	mockGatewayAPI := &api.MockGatewayAPI{}
	mockProcessAPI := &api.MockProcessAPI{}
	mockConnectionAPI := &api.MockConnectionAPI{}

	tenantService := servicetenant.NewTenantService(mockCfg, mockTenantAPI, mockDomainAPI, mockGatewayAPI, mockProcessAPI, mockConnectionAPI)
	tenantCmd := NewTenantCmd(tenantService)
	h, err := interactive_tester.New(tenantCmd, []string{"create", "--interactive"})
	if err != nil {
		t.Fatalf("could not create PTY harness: %v", err)
	}
	h.Start()
	defer h.Close()

	h.ExpectMultipleT(t, []string{"Tenant name", "required"}, interactive_tester.WithTimeout(30*time.Second))
	h.WriteLineT(t, "test-tenant")

	h.ExpectMultipleT(t, []string{"Tenant slug", "required"})
	h.WriteLineT(t, "test-tenant-slug")

	h.ExpectT(t, "Description")
	h.WriteDataT(t, h.Enter)

	h.ExpectT(t, "no verified domains available", interactive_tester.WithTimeout(15*time.Second))
}

func TestTenantSubCmd_Create_Interactive_GatewayAPIError(t *testing.T) {
	mockCfg := newTestTenantConfig()

	mockTenantAPI := &api.MockTenantAPI{}
	mockDomainAPI := &api.MockDomainAPI{
		ListFunc: func(
			_ configuration_models.EndpointsV2, _ string, _ string, _ int, _ int,
		) (*api.GenericPaginatedResponse[api.DomainDTO], error) {
			now := time.Now()
			return &api.GenericPaginatedResponse[api.DomainDTO]{
				Data: []api.DomainDTO{
					{
						ID:             "domain-001",
						DomainName:     "example.com",
						CreatedAt:      now,
						VerifiedAt:     &now,
						OrganizationID: "test-org-id",
					},
				},
			}, nil
		},
	}
	mockGatewayAPI := &api.MockGatewayAPI{
		ListGatewaysV5Func: func(
			_ configuration_models.EndpointsV2, _ string, _ string, _ ...api.ListGatewaysV5Option,
		) (*api.GenericPaginatedResponse[api.GatewayV5ListItemResponse], error) {
			return nil, fmt.Errorf("internal server error")
		},
	}
	mockProcessAPI := &api.MockProcessAPI{}
	mockConnectionAPI := &api.MockConnectionAPI{}

	tenantService := servicetenant.NewTenantService(mockCfg, mockTenantAPI, mockDomainAPI, mockGatewayAPI, mockProcessAPI, mockConnectionAPI)
	tenantCmd := NewTenantCmd(tenantService)
	h, err := interactive_tester.New(tenantCmd, []string{"create", "--interactive"})
	if err != nil {
		t.Fatalf("could not create PTY harness: %v", err)
	}
	h.Start()
	defer h.Close()

	h.ExpectMultipleT(t, []string{"Tenant name", "required"}, interactive_tester.WithTimeout(30*time.Second))
	h.WriteLineT(t, "test-tenant")

	h.ExpectMultipleT(t, []string{"Tenant slug", "required"})
	h.WriteLineT(t, "test-tenant-slug")

	h.ExpectT(t, "Description")
	h.WriteDataT(t, h.Enter)

	h.ExpectMultipleT(t, []string{
		"Select domains",
		"example.com (domain-001)",
	})
	h.WriteKeysT(t, h.Space, h.Enter)

	h.ExpectT(t, "internal server error", interactive_tester.WithTimeout(15*time.Second))
}

func TestTenantSubCmd_Create_Interactive_NoGateways(t *testing.T) {
	mockCfg := newTestTenantConfig()

	mockTenantAPI := &api.MockTenantAPI{}
	mockDomainAPI := &api.MockDomainAPI{
		ListFunc: func(
			_ configuration_models.EndpointsV2, _ string, _ string, _ int, _ int,
		) (*api.GenericPaginatedResponse[api.DomainDTO], error) {
			now := time.Now()
			return &api.GenericPaginatedResponse[api.DomainDTO]{
				Data: []api.DomainDTO{
					{
						ID:             "domain-001",
						DomainName:     "example.com",
						CreatedAt:      now,
						VerifiedAt:     &now,
						OrganizationID: "test-org-id",
					},
				},
			}, nil
		},
	}
	mockGatewayAPI := &api.MockGatewayAPI{
		ListGatewaysV5Func: func(
			_ configuration_models.EndpointsV2, _ string, _ string, _ ...api.ListGatewaysV5Option,
		) (*api.GenericPaginatedResponse[api.GatewayV5ListItemResponse], error) {
			return &api.GenericPaginatedResponse[api.GatewayV5ListItemResponse]{
				Data: []api.GatewayV5ListItemResponse{},
			}, nil
		},
	}
	mockProcessAPI := &api.MockProcessAPI{}
	mockConnectionAPI := &api.MockConnectionAPI{}

	tenantService := servicetenant.NewTenantService(mockCfg, mockTenantAPI, mockDomainAPI, mockGatewayAPI, mockProcessAPI, mockConnectionAPI)
	tenantCmd := NewTenantCmd(tenantService)
	h, err := interactive_tester.New(tenantCmd, []string{"create", "--interactive"})
	if err != nil {
		t.Fatalf("could not create PTY harness: %v", err)
	}
	h.Start()
	defer h.Close()

	h.ExpectMultipleT(t, []string{"Tenant name", "required"}, interactive_tester.WithTimeout(30*time.Second))
	h.WriteLineT(t, "test-tenant")

	h.ExpectMultipleT(t, []string{"Tenant slug", "required"})
	h.WriteLineT(t, "test-tenant-slug")

	h.ExpectT(t, "Description")
	h.WriteDataT(t, h.Enter)

	h.ExpectMultipleT(t, []string{
		"Select domains",
		"example.com (domain-001)",
	})
	h.WriteKeysT(t, h.Space, h.Enter)

	h.ExpectT(t, "no gateways available", interactive_tester.WithTimeout(15*time.Second))
}

func TestTenantSubCmd_Create_Interactive_TenantAPIFailure(t *testing.T) {
	mockCfg := newTestTenantConfig()

	mockTenantAPI := &api.MockTenantAPI{
		CreateTenantV5Func: func(
			_ configuration_models.EndpointsV2, _ string, _ string, _ *api.CreateTenantV5Request,
		) (*api.GenericIDResponseModel, error) {
			return nil, fmt.Errorf("rate limit exceeded")
		},
	}

	mockDomainAPI := &api.MockDomainAPI{
		ListFunc: func(
			_ configuration_models.EndpointsV2, _ string, _ string, _ int, _ int,
		) (*api.GenericPaginatedResponse[api.DomainDTO], error) {
			now := time.Now()
			return &api.GenericPaginatedResponse[api.DomainDTO]{
				Data: []api.DomainDTO{
					{
						ID:             "domain-001",
						DomainName:     "example.com",
						CreatedAt:      now,
						VerifiedAt:     &now,
						OrganizationID: "test-org-id",
					},
				},
			}, nil
		},
	}
	mockGatewayAPI := &api.MockGatewayAPI{
		ListGatewaysV5Func: func(
			_ configuration_models.EndpointsV2, _ string, _ string, _ ...api.ListGatewaysV5Option,
		) (*api.GenericPaginatedResponse[api.GatewayV5ListItemResponse], error) {
			return &api.GenericPaginatedResponse[api.GatewayV5ListItemResponse]{
				Data: []api.GatewayV5ListItemResponse{
					{ID: "gw-eu-001", Name: "gw-eu-west", Slug: "gw-eu-west"},
				},
			}, nil
		},
	}
	mockProcessAPI := &api.MockProcessAPI{}
	mockConnectionAPI := &api.MockConnectionAPI{}

	tenantService := servicetenant.NewTenantService(mockCfg, mockTenantAPI, mockDomainAPI, mockGatewayAPI, mockProcessAPI, mockConnectionAPI)
	tenantCmd := NewTenantCmd(tenantService)
	h, err := interactive_tester.New(tenantCmd, []string{"create", "--interactive"})
	if err != nil {
		t.Fatalf("could not create PTY harness: %v", err)
	}
	h.Start()
	defer h.Close()

	h.ExpectMultipleT(t, []string{"Tenant name", "required"}, interactive_tester.WithTimeout(30*time.Second))
	h.WriteLineT(t, "test-tenant")

	h.ExpectMultipleT(t, []string{"Tenant slug", "required"})
	h.WriteLineT(t, "test-tenant-slug")

	h.ExpectT(t, "Description")
	h.WriteDataT(t, h.Enter)

	h.ExpectMultipleT(t, []string{
		"Select domains",
		"example.com (domain-001)",
	})
	h.WriteKeysT(t, h.Space, h.Enter)

	h.ExpectMultipleT(t, []string{
		"Select gateways for domain example.com",
		"gw-eu-west (gw-eu-001)",
	})
	h.WriteKeysT(t, h.Space, h.Enter)

	h.ExpectMultipleT(t, []string{
		"Subdomain for domain example.com",
	})
	h.WriteDataT(t, h.Enter)

	h.ExpectT(t, "rate limit exceeded", interactive_tester.WithTimeout(15*time.Second))
}

func TestTenantSubCmd_Create_Interactive_DeploymentFailed(t *testing.T) {
	mockCfg := newTestTenantConfig()

	mockTenantAPI := &api.MockTenantAPI{
		CreateTenantV5Func: func(
			_ configuration_models.EndpointsV2, _ string, _ string, _ *api.CreateTenantV5Request,
		) (*api.GenericIDResponseModel, error) {
			return &api.GenericIDResponseModel{ID: "proc-fail"}, nil
		},
	}

	mockDomainAPI := &api.MockDomainAPI{
		ListFunc: func(
			_ configuration_models.EndpointsV2, _ string, _ string, _ int, _ int,
		) (*api.GenericPaginatedResponse[api.DomainDTO], error) {
			now := time.Now()
			return &api.GenericPaginatedResponse[api.DomainDTO]{
				Data: []api.DomainDTO{
					{
						ID:             "domain-001",
						DomainName:     "example.com",
						CreatedAt:      now,
						VerifiedAt:     &now,
						OrganizationID: "test-org-id",
					},
				},
			}, nil
		},
	}

	mockGatewayAPI := &api.MockGatewayAPI{
		ListGatewaysV5Func: func(
			_ configuration_models.EndpointsV2, _ string, _ string, _ ...api.ListGatewaysV5Option,
		) (*api.GenericPaginatedResponse[api.GatewayV5ListItemResponse], error) {
			return &api.GenericPaginatedResponse[api.GatewayV5ListItemResponse]{
				Data: []api.GatewayV5ListItemResponse{
					{ID: "gw-eu-001", Name: "gw-eu-west", Slug: "gw-eu-west"},
				},
			}, nil
		},
	}

	mockProcessAPI := &api.MockProcessAPI{
		GetProcessFunc: func(
			_ configuration_models.EndpointsV2, _ string, _ string, processID string,
		) (*api.Process, error) {
			data, _ := json.Marshal(api.TenantCreationProcessData{
				TenantID: "tenant-fail-id",
				Error: &api.ProcessError{
					Code:    "TENANT_ERR",
					Message: "gateway creation timed out",
				},
			})
			return &api.Process{
				ID:     processID,
				Step:   api.ProcessStepCreatingTenantGateway,
				Status: api.ProcessStatusFailed,
				Data:   data,
			}, nil
		},
	}

	mockConnectionAPI := &api.MockConnectionAPI{}
	tenantService := servicetenant.NewTenantService(mockCfg, mockTenantAPI, mockDomainAPI, mockGatewayAPI, mockProcessAPI, mockConnectionAPI)
	tenantCmd := NewTenantCmd(tenantService)
	h, err := interactive_tester.New(tenantCmd, []string{"create", "--interactive"})
	if err != nil {
		t.Fatalf("could not create PTY harness: %v", err)
	}
	h.Start()
	defer h.Close()

	h.ExpectMultipleT(t, []string{"Tenant name", "required"}, interactive_tester.WithTimeout(30*time.Second))
	h.WriteLineT(t, "test-tenant")

	h.ExpectMultipleT(t, []string{"Tenant slug", "required"})
	h.WriteLineT(t, "test-tenant-slug")

	h.ExpectT(t, "Description")
	h.WriteDataT(t, h.Enter)

	h.ExpectMultipleT(t, []string{
		"Select domains",
		"example.com (domain-001)",
	})
	h.WriteKeysT(t, h.Space, h.Enter)

	h.ExpectMultipleT(t, []string{
		"Select gateways for domain example.com",
		"gw-eu-west (gw-eu-001)",
	})
	h.WriteKeysT(t, h.Space, h.Enter)

	h.ExpectMultipleT(t, []string{
		"Subdomain for domain example.com",
	})
	h.WriteDataT(t, h.Enter)

	h.ExpectT(t, ", Error TENANT_ERR: gateway creation timed out", interactive_tester.WithTimeout(15*time.Second))
}
