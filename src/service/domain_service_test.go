package service

import (
	"bytes"
	"strings"
	"testing"
	"time"

	api "github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
)

func TestDomainService_Create_Human(t *testing.T) {
	mockCfg := configuration_handler.NewMockConfigurationHandler()
	mockCfg.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{
			APIKey:         "test-api-key",
			OrganizationID: "test-org-id",
			Output:         configuration_models.OutputHuman,
			Endpoints: configuration_models.EndpointsV2{
				IAM:  "https://iam.example.com",
				Dash: "https://dash.example.com",
				CH:   "https://ch.example.com",
			},
		}, nil
	}

	createdAt := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)
	mockDomainAPI := &api.MockDomainAPI{
		CreateFunc: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, request *api.CreateDomainRequestBody) (*api.DomainDTO, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}
			if request.DomainName != "example.com" {
				t.Fatalf("Expected domain name 'example.com', got %q", request.DomainName)
			}

			return &api.DomainDTO{
				ID:             "domain-123",
				DomainName:     "example.com",
				CreatedAt:      createdAt,
				Challenge:      "challenge-token-abc",
				OrganizationID: "test-org-id",
				IsShared:       false,
			}, nil
		},
	}

	service := NewDomainService(mockCfg, mockDomainAPI, &api.UserAPI{})
	cmd := setupTestCommand()
	cmd.Flags().String("domain-name", "", "Domain name")
	cmd.Flags().Set("domain-name", "example.com")

	err := service.Create(cmd, nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	output := cmd.OutOrStdout().(*bytes.Buffer).String()
	if !strings.Contains(output, "Domain: example.com") || !strings.Contains(output, "ID: domain-123") {
		t.Fatalf("Expected human output to contain domain details, got %q", output)
	}
}

func TestDomainService_Create_JSON(t *testing.T) {
	t.Skip("skipping JSON test")
	mockCfg := configuration_handler.NewMockConfigurationHandler()
	mockCfg.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{
			APIKey:         "test-api-key",
			OrganizationID: "test-org-id",
			Output:         "json",
			Endpoints: configuration_models.EndpointsV2{
				IAM:  "https://iam.example.com",
				Dash: "https://dash.example.com",
				CH:   "https://ch.example.com",
			},
		}, nil
	}

	createdAt := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)
	mockDomainAPI := &api.MockDomainAPI{
		CreateFunc: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, request *api.CreateDomainRequestBody) (*api.DomainDTO, error) {
			return &api.DomainDTO{
				ID:             "domain-456",
				DomainName:     "json-domain.com",
				CreatedAt:      createdAt,
				Challenge:      "challenge-json",
				OrganizationID: "test-org-id",
				IsShared:       false,
			}, nil
		},
	}

	service := NewDomainService(mockCfg, mockDomainAPI, &api.UserAPI{})
	cmd := setupTestCommand()
	cmd.Flags().String("domain-name", "", "Domain name")
	cmd.Flags().Set("domain-name", "json-domain.com")

	err := service.Create(cmd, nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	output := cmd.OutOrStdout().(*bytes.Buffer).String()
	if !strings.Contains(output, `"id": "domain-456"`) {
		t.Fatalf("Expected json output to contain domain id, got %q", output)
	}
}

func TestDomainService_Describe_Human(t *testing.T) {
	mockCfg := configuration_handler.NewMockConfigurationHandler()
	mockCfg.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{
			APIKey:         "test-api-key",
			OrganizationID: "test-org-id",
			Output:         configuration_models.OutputHuman,
			Endpoints: configuration_models.EndpointsV2{
				IAM:  "https://iam.example.com",
				Dash: "https://dash.example.com",
				CH:   "https://ch.example.com",
			},
		}, nil
	}

	createdAt := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)
	verifiedAt := time.Date(2024, 2, 20, 14, 0, 0, 0, time.UTC)
	mockDomainAPI := &api.MockDomainAPI{
		GetFunc: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, domainID string) (*api.DomainDTO, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}
			if domainID != "domain-123" {
				t.Fatalf("Expected domain ID domain-123, got %q", domainID)
			}

			return &api.DomainDTO{
				ID:             "domain-123",
				DomainName:     "verified-example.com",
				CreatedAt:      createdAt,
				VerifiedAt:     &verifiedAt,
				Challenge:      "verify-token",
				OrganizationID: "test-org-id",
				IsShared:       true,
			}, nil
		},
	}

	service := NewDomainService(mockCfg, mockDomainAPI, &api.UserAPI{})
	cmd := setupTestCommand()
	cmd.Flags().String("domain-id", "", "Domain ID")
	cmd.Flags().Set("domain-id", "domain-123")

	err := service.Describe(cmd, nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	output := cmd.OutOrStdout().(*bytes.Buffer).String()
	if !strings.Contains(output, "Domain: verified-example.com") ||
		!strings.Contains(output, "ID: domain-123") ||
		!strings.Contains(output, "Verified: Yes") ||
		!strings.Contains(output, "Shared: Yes") {
		t.Fatalf("Expected human output to contain verified domain details, got %q", output)
	}
}

func TestDomainService_Describe_JSON(t *testing.T) {
	t.Skip("skipping JSON test")
	mockCfg := configuration_handler.NewMockConfigurationHandler()
	mockCfg.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{
			APIKey:         "test-api-key",
			OrganizationID: "test-org-id",
			Output:         "json",
			Endpoints: configuration_models.EndpointsV2{
				IAM:  "https://iam.example.com",
				Dash: "https://dash.example.com",
				CH:   "https://ch.example.com",
			},
		}, nil
	}

	mockDomainAPI := &api.MockDomainAPI{
		GetFunc: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, domainID string) (*api.DomainDTO, error) {
			return &api.DomainDTO{
				ID:             "domain-789",
				DomainName:     "json-describe.com",
				Challenge:      "token-json",
				OrganizationID: "test-org-id",
			}, nil
		},
	}

	service := NewDomainService(mockCfg, mockDomainAPI, &api.UserAPI{})
	cmd := setupTestCommand()
	cmd.Flags().String("domain-id", "", "Domain ID")
	cmd.Flags().Set("domain-id", "domain-789")

	err := service.Describe(cmd, nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	output := cmd.OutOrStdout().(*bytes.Buffer).String()
	if !strings.Contains(output, `"id": "domain-789"`) ||
		!strings.Contains(output, `"domain_name": "json-describe.com"`) {
		t.Fatalf("Expected json output to contain domain data, got %q", output)
	}
}

func TestDomainService_List_Human(t *testing.T) {
	mockCfg := configuration_handler.NewMockConfigurationHandler()
	mockCfg.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{
			APIKey:         "test-api-key",
			OrganizationID: "test-org-id",
			Output:         configuration_models.OutputHuman,
			Endpoints: configuration_models.EndpointsV2{
				IAM:  "https://iam.example.com",
				Dash: "https://dash.example.com",
				CH:   "https://ch.example.com",
			},
		}, nil
	}

	createdAt := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)
	mockDomainAPI := &api.MockDomainAPI{
		ListFunc: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, page int, itemsPerPage int) (*api.GenericPaginatedResponse[api.DomainDTO], error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}
			if page != 1 || itemsPerPage != 100 {
				t.Fatalf("Expected page=1 items=100, got page=%d items=%d", page, itemsPerPage)
			}

			return &api.GenericPaginatedResponse[api.DomainDTO]{
				Data: []api.DomainDTO{
					{
						ID:         "domain-001",
						DomainName: "first-domain.com",
						CreatedAt:  createdAt,
						Challenge:  "challenge-001",
					},
					{
						ID:         "domain-002",
						DomainName: "second-domain.com",
						CreatedAt:  createdAt,
						Challenge:  "challenge-002",
						VerifiedAt: &createdAt,
					},
				},
				NextPage: nil,
				Count:    2,
			}, nil
		},
	}

	service := NewDomainService(mockCfg, mockDomainAPI, &api.UserAPI{})
	cmd := setupTestCommand()

	err := service.List(cmd, nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	output := cmd.OutOrStdout().(*bytes.Buffer).String()
	if !strings.Contains(output, "first-domain.com") ||
		!strings.Contains(output, "second-domain.com") ||
		!strings.Contains(output, "Domain Name") {
		t.Fatalf("Expected human output to contain domain list with headers, got %q", output)
	}
}

func TestDomainService_List_JSON(t *testing.T) {
	t.Skip("skipping JSON test")
	mockCfg := configuration_handler.NewMockConfigurationHandler()
	mockCfg.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{
			APIKey:         "test-api-key",
			OrganizationID: "test-org-id",
			Output:         "json",
			Endpoints: configuration_models.EndpointsV2{
				IAM:  "https://iam.example.com",
				Dash: "https://dash.example.com",
				CH:   "https://ch.example.com",
			},
		}, nil
	}

	createdAt := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)
	mockDomainAPI := &api.MockDomainAPI{
		ListFunc: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, page int, itemsPerPage int) (*api.GenericPaginatedResponse[api.DomainDTO], error) {
			return &api.GenericPaginatedResponse[api.DomainDTO]{
				Data: []api.DomainDTO{
					{
						ID:         "domain-001",
						DomainName: "json-list.com",
						CreatedAt:  createdAt,
						Challenge:  "challenge-json-list",
					},
				},
				NextPage: nil,
				Count:    1,
			}, nil
		},
	}

	service := NewDomainService(mockCfg, mockDomainAPI, &api.UserAPI{})
	cmd := setupTestCommand()

	err := service.List(cmd, nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	output := cmd.OutOrStdout().(*bytes.Buffer).String()
	if !strings.Contains(output, `"id": "domain-001"`) ||
		!strings.Contains(output, `"domain_name": "json-list.com"`) {
		t.Fatalf("Expected json output to contain domain list data, got %q", output)
	}
}

func TestDomainService_List_Pagination(t *testing.T) {
	t.Skip("skipping JSON test")
	mockCfg := configuration_handler.NewMockConfigurationHandler()
	mockCfg.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{
			APIKey:         "test-api-key",
			OrganizationID: "test-org-id",
			Output:         "json",
			Endpoints: configuration_models.EndpointsV2{
				IAM:  "https://iam.example.com",
				Dash: "https://dash.example.com",
				CH:   "https://ch.example.com",
			},
		}, nil
	}

	callCount := 0
	mockDomainAPI := &api.MockDomainAPI{
		ListFunc: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, page int, itemsPerPage int) (*api.GenericPaginatedResponse[api.DomainDTO], error) {
			callCount++
			if itemsPerPage != 100 {
				t.Fatalf("Expected items per page 100, got %d", itemsPerPage)
			}

			switch page {
			case 1:
				return &api.GenericPaginatedResponse[api.DomainDTO]{
					Data: []api.DomainDTO{
						{
							ID:         "domain-page1",
							DomainName: "first-page-domain.com",
						},
					},
					NextPage: func() *int { next := 2; return &next }(),
					Count:    3,
				}, nil
			case 2:
				return &api.GenericPaginatedResponse[api.DomainDTO]{
					Data: []api.DomainDTO{
						{
							ID:         "domain-page2",
							DomainName: "second-page-domain.com",
						},
						{
							ID:         "domain-page3",
							DomainName: "third-domain.com",
						},
					},
					NextPage: nil,
					Count:    2,
				}, nil
			default:
				t.Fatalf("Unexpected page requested: %d", page)
				return nil, nil
			}
		},
	}

	service := NewDomainService(mockCfg, mockDomainAPI, &api.UserAPI{})
	cmd := setupTestCommand()

	err := service.List(cmd, nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if callCount != 2 {
		t.Fatalf("Expected 2 API calls for 2 pages, got %d", callCount)
	}

	output := cmd.OutOrStdout().(*bytes.Buffer).String()
	if !strings.Contains(output, "domain-page1") ||
		!strings.Contains(output, "domain-page2") ||
		!strings.Contains(output, "domain-page3") {
		t.Fatalf("Expected all paginated results in output, got %q", output)
	}
}

func TestDomainService_Delete(t *testing.T) {
	mockCfg := configuration_handler.NewMockConfigurationHandler()
	mockCfg.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{
			APIKey:         "test-api-key",
			OrganizationID: "test-org-id",
			Output:         configuration_models.OutputHuman,
			Endpoints: configuration_models.EndpointsV2{
				IAM:  "https://iam.example.com",
				Dash: "https://dash.example.com",
				CH:   "https://ch.example.com",
			},
		}, nil
	}

	mockDomainAPI := &api.MockDomainAPI{
		DeleteFunc: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, domainID string) error {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}
			if domainID != "domain-to-delete" {
				t.Fatalf("Expected domain ID 'domain-to-delete', got %q", domainID)
			}
			return nil
		},
	}

	service := NewDomainService(mockCfg, mockDomainAPI, &api.UserAPI{})
	cmd := setupTestCommand()
	cmd.Flags().String("domain-id", "", "Domain ID")
	cmd.Flags().Set("domain-id", "domain-to-delete")

	err := service.Delete(cmd, nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	output := cmd.OutOrStdout().(*bytes.Buffer).String()
	if !strings.Contains(output, "Domain domain-to-delete deleted successfully") {
		t.Fatalf("Expected deletion success message, got %q", output)
	}
}

func TestDomainService_Verify_Verified_Human(t *testing.T) {
	mockCfg := configuration_handler.NewMockConfigurationHandler()
	mockCfg.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{
			APIKey:         "test-api-key",
			OrganizationID: "test-org-id",
			Output:         configuration_models.OutputHuman,
			Endpoints: configuration_models.EndpointsV2{
				IAM:  "https://iam.example.com",
				Dash: "https://dash.example.com",
				CH:   "https://ch.example.com",
			},
		}, nil
	}

	mockDomainAPI := &api.MockDomainAPI{
		VerifyFunc: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, domainID string) (*api.DomainVerifyResult, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}
			if domainID != "domain-verified" {
				t.Fatalf("Expected domain ID 'domain-verified', got %q", domainID)
			}
			return &api.DomainVerifyResult{Verified: true}, nil
		},
	}

	service := NewDomainService(mockCfg, mockDomainAPI, &api.UserAPI{})
	cmd := setupTestCommand()
	cmd.Flags().String("domain-id", "", "Domain ID")
	cmd.Flags().Set("domain-id", "domain-verified")

	err := service.Verify(cmd, nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	output := cmd.OutOrStdout().(*bytes.Buffer).String()
	if !strings.Contains(output, "Domain domain-verified is verified") {
		t.Fatalf("Expected verified message, got %q", output)
	}
}

func TestDomainService_Verify_NotVerified_Human(t *testing.T) {
	mockCfg := configuration_handler.NewMockConfigurationHandler()
	mockCfg.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{
			APIKey:         "test-api-key",
			OrganizationID: "test-org-id",
			Output:         configuration_models.OutputHuman,
			Endpoints: configuration_models.EndpointsV2{
				IAM:  "https://iam.example.com",
				Dash: "https://dash.example.com",
				CH:   "https://ch.example.com",
			},
		}, nil
	}

	mockDomainAPI := &api.MockDomainAPI{
		VerifyFunc: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, domainID string) (*api.DomainVerifyResult, error) {
			return &api.DomainVerifyResult{Verified: false}, nil
		},
	}

	service := NewDomainService(mockCfg, mockDomainAPI, &api.UserAPI{})
	cmd := setupTestCommand()
	cmd.Flags().String("domain-id", "", "Domain ID")
	cmd.Flags().Set("domain-id", "domain-not-verified")

	err := service.Verify(cmd, nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	output := cmd.OutOrStdout().(*bytes.Buffer).String()
	if !strings.Contains(output, "Domain domain-not-verified is not verified") {
		t.Fatalf("Expected not verified message, got %q", output)
	}
}

func TestDomainService_Verify_JSON(t *testing.T) {
	t.Skip("skipping JSON test")
	mockCfg := configuration_handler.NewMockConfigurationHandler()
	mockCfg.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{
			APIKey:         "test-api-key",
			OrganizationID: "test-org-id",
			Output:         "json",
			Endpoints: configuration_models.EndpointsV2{
				IAM:  "https://iam.example.com",
				Dash: "https://dash.example.com",
				CH:   "https://ch.example.com",
			},
		}, nil
	}

	mockDomainAPI := &api.MockDomainAPI{
		VerifyFunc: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, domainID string) (*api.DomainVerifyResult, error) {
			return &api.DomainVerifyResult{Verified: true}, nil
		},
	}

	service := NewDomainService(mockCfg, mockDomainAPI, &api.UserAPI{})
	cmd := setupTestCommand()
	cmd.Flags().String("domain-id", "", "Domain ID")
	cmd.Flags().Set("domain-id", "domain-verify-json")

	err := service.Verify(cmd, nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	output := cmd.OutOrStdout().(*bytes.Buffer).String()
	if !strings.Contains(output, `"verified": true`) {
		t.Fatalf("Expected json output to contain verified result, got %q", output)
	}
}

func TestDomainService_InterfaceCompliance(t *testing.T) {
	var _ DomainServiceInterface = DomainService{}
}
