package cmd_domain

import (
	"bytes"
	"strings"
	"testing"
	"time"

	api "github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/src/service"
	"github.com/spf13/cobra"
)

func createTestDomainCmd(mockAPI *api.MockDomainAPI) (*cobra.Command, *bytes.Buffer) {
	mockCfg := configuration_handler.NewMockConfigurationHandler()
	mockCfg.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{
			Output:         configuration_models.OutputHuman,
			APIKey:         "test-api-key",
			OrganizationID: "test-org-id",
		}, nil
	}
	svc := service.NewDomainService(mockCfg, mockAPI, &api.UserAPI{})

	rootCmd := &cobra.Command{Use: "cubbit"}
	rootCmd.PersistentFlags().String("profile", "", "Profile Configuration")
	rootCmd.PersistentFlags().String("output", "human", "Output format: human (default), json, yaml")
	rootCmd.PersistentFlags().Bool("no-headers", false, "Suppress table headers in human output")
	rootCmd.PersistentFlags().Bool("quiet", false, "Minimize stdout for CI/CD workflows")
	rootCmd.PersistentFlags().Bool("silent", false, "Redirect all output to /dev/null")

	domainCmd := NewDomainCmd(svc)
	rootCmd.AddCommand(domainCmd)

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	return rootCmd, buf
}

func TestDomainIntegration_Create_HumanOutput(t *testing.T) {
	createdAt := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)

	mockAPI := &api.MockDomainAPI{
		CreateFunc: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, request *api.CreateDomainRequestBody) (*api.DomainDTO, error) {
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

	cmd, buf := createTestDomainCmd(mockAPI)
	cmd.SetArgs([]string{"domain", "create", "--domain-name", "example.com"})

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := strings.TrimSpace(`
Domain: example.com
ID: domain-123
Organization: test-org-id
Challenge: challenge-token-abc
Verified: No
Verified At: N/A
Deleted: No
Shared: No
Created At: 2024-01-15 10:30:00
`)

	if strings.TrimSpace(buf.String()) != expected {
		t.Fatalf("Expected output:\n%s\n\nactual:\n%s", expected, strings.TrimSpace(buf.String()))
	}
}

func TestDomainIntegration_Create_JSONOutput(t *testing.T) {
	t.Skip("enable after printer migration")
	createdAt := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)

	mockAPI := &api.MockDomainAPI{
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

	cmd, buf := createTestDomainCmd(mockAPI)
	cmd.SetArgs([]string{"domain", "create", "--domain-name", "json-domain.com", "--output", "json"})

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	output := buf.String()
	if !strings.Contains(output, `"id": "domain-456"`) {
		t.Fatalf("Expected json output to contain domain id, got %q", output)
	}
	if !strings.Contains(output, `"domain_name": "json-domain.com"`) {
		t.Fatalf("Expected json output to contain domain name, got %q", output)
	}
}

func TestDomainIntegration_Describe_HumanOutput(t *testing.T) {
	createdAt := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)
	verifiedAt := time.Date(2024, 2, 20, 14, 0, 0, 0, time.UTC)

	mockAPI := &api.MockDomainAPI{
		GetFunc: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, domainID string) (*api.DomainDTO, error) {
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

	cmd, buf := createTestDomainCmd(mockAPI)
	cmd.SetArgs([]string{"domain", "describe", "--domain-id", "domain-123"})

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := strings.TrimSpace(`
Domain: verified-example.com
ID: domain-123
Organization: test-org-id
Challenge: verify-token
Verified: Yes
Verified At: 2024-02-20 14:00:00
Deleted: No
Shared: Yes
Created At: 2024-01-15 10:30:00
`)

	if strings.TrimSpace(buf.String()) != expected {
		t.Fatalf("Expected output:\n%s\n\nactual:\n%s", expected, strings.TrimSpace(buf.String()))
	}
}

func TestDomainIntegration_Describe_HumanOutput_NotVerified(t *testing.T) {
	createdAt := time.Date(2024, 3, 10, 8, 0, 0, 0, time.UTC)

	mockAPI := &api.MockDomainAPI{
		GetFunc: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, domainID string) (*api.DomainDTO, error) {
			return &api.DomainDTO{
				ID:             "domain-unver",
				DomainName:     "unverified-domain.com",
				CreatedAt:      createdAt,
				Challenge:      "unverified-challenge",
				OrganizationID: "test-org-id",
				IsShared:       false,
			}, nil
		},
	}

	cmd, buf := createTestDomainCmd(mockAPI)
	cmd.SetArgs([]string{"domain", "describe", "--domain-id", "domain-unver"})

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := strings.TrimSpace(`
Domain: unverified-domain.com
ID: domain-unver
Organization: test-org-id
Challenge: unverified-challenge
Verified: No
Verified At: N/A
Deleted: No
Shared: No
Created At: 2024-03-10 08:00:00
`)

	if strings.TrimSpace(buf.String()) != expected {
		t.Fatalf("Expected output:\n%s\n\nactual:\n%s", expected, strings.TrimSpace(buf.String()))
	}
}

func TestDomainIntegration_Describe_JSONOutput(t *testing.T) {
	t.Skip("enable after printer refactor")
	mockAPI := &api.MockDomainAPI{
		GetFunc: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, domainID string) (*api.DomainDTO, error) {
			return &api.DomainDTO{
				ID:             "domain-789",
				DomainName:     "json-describe.com",
				Challenge:      "token-json",
				OrganizationID: "test-org-id",
			}, nil
		},
	}

	cmd, buf := createTestDomainCmd(mockAPI)
	cmd.SetArgs([]string{"domain", "describe", "--domain-id", "domain-789", "--output", "json"})

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	output := buf.String()
	if !strings.Contains(output, `"id": "domain-789"`) {
		t.Fatalf("Expected json output to contain domain id, got %q", output)
	}
	if !strings.Contains(output, `"domain_name": "json-describe.com"`) {
		t.Fatalf("Expected json output to contain domain name, got %q", output)
	}
}

func TestDomainIntegration_List_HumanOutput(t *testing.T) {
	createdAt := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)

	mockAPI := &api.MockDomainAPI{
		ListFunc: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, page int, itemsPerPage int) (*api.GenericPaginatedResponse[api.DomainDTO], error) {
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

	cmd, buf := createTestDomainCmd(mockAPI)
	cmd.SetArgs([]string{"domain", "list"})

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := strings.TrimSpace(`
╭───────────────────┬────────────┬──────────┬─────────────────────┬───────────────╮
│ Domain Name       │ ID         │ Verified │ Created On          │ Challenge     │
├───────────────────┼────────────┼──────────┼─────────────────────┼───────────────┤
│ first-domain.com  │ domain-001 │ No       │ 2024-01-15 10:30:00 │ challenge-001 │
│ second-domain.com │ domain-002 │ Yes      │ 2024-01-15 10:30:00 │ challenge-002 │
╰───────────────────┴────────────┴──────────┴─────────────────────┴───────────────╯
`)

	if strings.TrimSpace(buf.String()) != expected {
		t.Fatalf("Expected output:\n%s\n\nactual:\n%s", expected, strings.TrimSpace(buf.String()))
	}
}

func TestDomainIntegration_List_HumanOutput_Empty(t *testing.T) {
	mockAPI := &api.MockDomainAPI{
		ListFunc: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, page int, itemsPerPage int) (*api.GenericPaginatedResponse[api.DomainDTO], error) {
			return &api.GenericPaginatedResponse[api.DomainDTO]{
				Data:     []api.DomainDTO{},
				NextPage: nil,
				Count:    0,
			}, nil
		},
	}

	cmd, buf := createTestDomainCmd(mockAPI)
	cmd.SetArgs([]string{"domain", "list"})

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	output := buf.String()
	if !strings.Contains(output, "No domains found") {
		t.Fatalf("Expected empty state message, got %q", output)
	}
}

func TestDomainIntegration_List_JSONOutput(t *testing.T) {
	t.Skip("enable after printer migration")
	createdAt := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)

	mockAPI := &api.MockDomainAPI{
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

	cmd, buf := createTestDomainCmd(mockAPI)
	cmd.SetArgs([]string{"domain", "list", "--output", "json"})

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	output := buf.String()
	if !strings.Contains(output, `"id": "domain-001"`) {
		t.Fatalf("Expected json output to contain domain id, got %q", output)
	}
	if !strings.Contains(output, `"domain_name": "json-list.com"`) {
		t.Fatalf("Expected json output to contain domain name, got %q", output)
	}
}

func TestDomainIntegration_List_Paginated(t *testing.T) {
	mockAPI := &api.MockDomainAPI{
		ListFunc: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, page int, itemsPerPage int) (*api.GenericPaginatedResponse[api.DomainDTO], error) {
			switch page {
			case 1:
				return &api.GenericPaginatedResponse[api.DomainDTO]{
					Data: []api.DomainDTO{
						{ID: "domain-page1", DomainName: "first-page.com"},
					},
					NextPage: func() *int { next := 2; return &next }(),
					Count:    3,
				}, nil
			case 2:
				return &api.GenericPaginatedResponse[api.DomainDTO]{
					Data: []api.DomainDTO{
						{ID: "domain-page2", DomainName: "second-page.com"},
						{ID: "domain-page3", DomainName: "third-domain.com"},
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

	cmd, buf := createTestDomainCmd(mockAPI)
	cmd.SetArgs([]string{"domain", "list", "--output", "json"})

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	output := buf.String()
	if !strings.Contains(output, "domain-page1") ||
		!strings.Contains(output, "domain-page2") ||
		!strings.Contains(output, "domain-page3") {
		t.Fatalf("Expected all paginated results in output, got %q", output)
	}
}

func TestDomainIntegration_List_AliasLS(t *testing.T) {
	mockAPI := &api.MockDomainAPI{
		ListFunc: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, page int, itemsPerPage int) (*api.GenericPaginatedResponse[api.DomainDTO], error) {
			return &api.GenericPaginatedResponse[api.DomainDTO]{
				Data:     []api.DomainDTO{},
				NextPage: nil,
				Count:    0,
			}, nil
		},
	}

	cmd, buf := createTestDomainCmd(mockAPI)
	cmd.SetArgs([]string{"domain", "ls"})

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	output := buf.String()
	if !strings.Contains(output, "No domains found") {
		t.Fatalf("Expected output via ls alias, got %q", output)
	}
}

func TestDomainIntegration_Delete_HumanOutput(t *testing.T) {
	mockAPI := &api.MockDomainAPI{
		DeleteFunc: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, domainID string) error {
			return nil
		},
	}

	cmd, buf := createTestDomainCmd(mockAPI)
	cmd.SetArgs([]string{"domain", "delete", "--domain-id", "domain-to-delete"})

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := "Domain domain-to-delete deleted successfully\n"
	if buf.String() != expected {
		t.Fatalf("Expected output %q, got %q", expected, buf.String())
	}
}

func TestDomainIntegration_Delete_AliasRM(t *testing.T) {
	mockAPI := &api.MockDomainAPI{
		DeleteFunc: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, domainID string) error {
			return nil
		},
	}

	cmd, buf := createTestDomainCmd(mockAPI)
	cmd.SetArgs([]string{"domain", "rm", "--domain-id", "domain-to-rm"})

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := "Domain domain-to-rm deleted successfully\n"
	if buf.String() != expected {
		t.Fatalf("Expected output via rm alias, got %q", buf.String())
	}
}

func TestDomainIntegration_Verify_HumanOutput_Verified(t *testing.T) {
	mockAPI := &api.MockDomainAPI{
		VerifyFunc: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, domainID string) (*api.DomainVerifyResult, error) {
			return &api.DomainVerifyResult{Verified: true}, nil
		},
	}

	cmd, buf := createTestDomainCmd(mockAPI)
	cmd.SetArgs([]string{"domain", "verify", "--domain-id", "domain-verified"})

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := "Domain domain-verified is verified\n"
	if buf.String() != expected {
		t.Fatalf("Expected output %q, got %q", expected, buf.String())
	}
}

func TestDomainIntegration_Verify_HumanOutput_NotVerified(t *testing.T) {
	mockAPI := &api.MockDomainAPI{
		VerifyFunc: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, domainID string) (*api.DomainVerifyResult, error) {
			return &api.DomainVerifyResult{Verified: false}, nil
		},
	}

	cmd, buf := createTestDomainCmd(mockAPI)
	cmd.SetArgs([]string{"domain", "verify", "--domain-id", "domain-not-verified"})

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := "Domain domain-not-verified is not verified\n"
	if buf.String() != expected {
		t.Fatalf("Expected output %q, got %q", expected, buf.String())
	}
}

func TestDomainIntegration_Verify_JSONOutput(t *testing.T) {
	t.Skip("enable after printer migration")
	mockAPI := &api.MockDomainAPI{
		VerifyFunc: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, domainID string) (*api.DomainVerifyResult, error) {
			return &api.DomainVerifyResult{Verified: true}, nil
		},
	}

	cmd, buf := createTestDomainCmd(mockAPI)
	cmd.SetArgs([]string{"domain", "verify", "--domain-id", "domain-verify-json", "--output", "json"})

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	output := buf.String()
	if !strings.Contains(output, `"verified": true`) {
		t.Fatalf("Expected json output to contain verified result, got %q", output)
	}
}

func TestDomainIntegration_Describe_AliasInfo(t *testing.T) {
	mockAPI := &api.MockDomainAPI{
		GetFunc: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, domainID string) (*api.DomainDTO, error) {
			return &api.DomainDTO{
				ID:             "domain-alias",
				DomainName:     "alias-domain.com",
				OrganizationID: "test-org-id",
			}, nil
		},
	}

	cmd, buf := createTestDomainCmd(mockAPI)
	cmd.SetArgs([]string{"domain", "info", "--domain-id", "domain-alias"})

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !strings.Contains(buf.String(), "alias-domain.com") {
		t.Fatalf("Expected output containing domain name via info alias, got %q", buf.String())
	}
}
