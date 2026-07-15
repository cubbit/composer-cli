package service

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/spf13/cobra"
)

func TestUserService_ImportUsers_SampleSkipsProfileLoading(t *testing.T) {
	mockCfg := configuration_handler.NewMockConfigurationHandler()
	mockCfg.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		t.Fatal("GetActiveProfile should not be called when generating an import sample")
		return configuration_models.ProfileV2{}, nil
	}

	cmd := &cobra.Command{}
	cmd.SetOut(new(bytes.Buffer))
	cmd.Flags().String("sample", "", "Sample format")
	cmd.Flags().String("file", "", "Users file")
	cmd.Flags().Set("sample", "json")

	service := NewUserService(mockCfg, nil, nil)
	err := service.ImportUsers(cmd, nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !strings.Contains(cmd.OutOrStdout().(*bytes.Buffer).String(), `"users": [`) {
		t.Fatalf("Expected sample JSON output, got %q", cmd.OutOrStdout().(*bytes.Buffer).String())
	}
}

func TestUserService_GetProfileOrganizationName(t *testing.T) {
	organizationName := "test-org"
	userAPI := &api.MockUserAPI{
		GetIAMUserSelfFunc: func(endpoints configuration_models.EndpointsV2, accessToken string, apiKey string) (*api.IAMUser, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected API key to be propagated, got %q", apiKey)
			}
			return &api.IAMUser{
				OrganizationName: &organizationName,
			}, nil
		},
	}
	service := NewUserService(nil, nil, userAPI)

	actualOrganizationName, err := service.getProfileOrganizationName(configuration_models.ProfileV2{
		APIKey: "test-api-key",
		Endpoints: configuration_models.EndpointsV2{
			IAM: "https://iam.example.com",
		},
	})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if actualOrganizationName != organizationName {
		t.Fatalf("Expected organization name %q, got %q", organizationName, actualOrganizationName)
	}
}

func TestUserService_GetProfileOrganizationName_Missing(t *testing.T) {
	userAPI := &api.MockUserAPI{
		GetIAMUserSelfFunc: func(endpoints configuration_models.EndpointsV2, accessToken string, apiKey string) (*api.IAMUser, error) {
			return &api.IAMUser{}, nil
		},
	}
	service := NewUserService(nil, nil, userAPI)

	_, err := service.getProfileOrganizationName(configuration_models.ProfileV2{})
	if err == nil {
		t.Fatal("Expected missing organization name error, got nil")
	}
	if !strings.Contains(err.Error(), "current IAM user does not expose an organization name") {
		t.Fatalf("Expected missing organization name error, got %v", err)
	}
}

func setupListTestCommand() *cobra.Command {
	cmd := setupTestCommand()
	cmd.Flags().String("enabled", "", "Filter enabled")
	cmd.Flags().String("search", "", "Search")
	cmd.Flags().Int("page", 1, "Page")
	cmd.Flags().Int("items", 100, "Items")
	cmd.Flags().String("sort-key", "", "Sort key")
	cmd.Flags().String("sort-order", "", "Sort order")
	return cmd
}

func TestUserService_ListUsers_Human(t *testing.T) {
	mockCfg := configuration_handler.NewMockConfigurationHandler()
	mockCfg.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{
			APIKey:         "test-api-key",
			OrganizationID: "test-org-id",
			Output:         configuration_models.OutputHuman,
			Endpoints: configuration_models.EndpointsV2{
				IAM: "https://iam.example.com",
			},
		}, nil
	}

	createdAt := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)
	firstName := "Alice"
	lastName := "Wonder"
	mockUserAPI := &api.MockUserAPI{
		ListIAMUsersFunc: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			enabled *bool,
			search string,
			page int,
			items int,
			sortKey string,
			sortOrder string,
		) (*api.GenericPaginatedResponse[api.IAMUserListItem], error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}
			if page != 1 || items != 100 {
				t.Fatalf("Expected page=1 items=100, got page=%d items=%d", page, items)
			}
			if enabled != nil {
				t.Fatalf("Expected enabled to be nil when not set, got %v", *enabled)
			}

			return &api.GenericPaginatedResponse[api.IAMUserListItem]{
				Data: []api.IAMUserListItem{
					{
						ID:        "user-001",
						Username:  "alice.wonder",
						FirstName: &firstName,
						LastName:  &lastName,
						Enabled:   true,
						CreatedAt: createdAt,
						Emails: []api.IAMUserEmail{
							{Email: "alice@example.com", Default: true},
						},
						Status: "active",
					},
				},
				NextPage: nil,
				Count:    1,
			}, nil
		},
	}

	service := NewUserService(mockCfg, nil, mockUserAPI)
	cmd := setupListTestCommand()

	err := service.ListUsers(cmd, nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	output := cmd.OutOrStdout().(*bytes.Buffer).String()
	if !strings.Contains(output, "alice.wonder") ||
		!strings.Contains(output, "user-001") ||
		!strings.Contains(output, "alice@example.com") ||
		!strings.Contains(output, "Username") {
		t.Fatalf("Expected human output to contain user list with headers, got %q", output)
	}
}

func TestUserService_ListUsers_WithFilters(t *testing.T) {
	mockCfg := configuration_handler.NewMockConfigurationHandler()
	mockCfg.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{
			APIKey:         "test-api-key",
			OrganizationID: "test-org-id",
			Output:         configuration_models.OutputHuman,
			Endpoints:      configuration_models.EndpointsV2{IAM: "https://iam.example.com"},
		}, nil
	}

	mockUserAPI := &api.MockUserAPI{
		ListIAMUsersFunc: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			enabled *bool,
			search string,
			page int,
			items int,
			sortKey string,
			sortOrder string,
		) (*api.GenericPaginatedResponse[api.IAMUserListItem], error) {
			if enabled == nil || *enabled != true {
				t.Fatalf("Expected enabled to be true, got %v", enabled)
			}
			if search != "bob" {
				t.Fatalf("Expected search to be 'bob', got %q", search)
			}
			if sortKey != "username" {
				t.Fatalf("Expected sortKey to be 'username', got %q", sortKey)
			}
			if sortOrder != "desc" {
				t.Fatalf("Expected sortOrder to be 'desc', got %q", sortOrder)
			}

			return &api.GenericPaginatedResponse[api.IAMUserListItem]{
				Data:     []api.IAMUserListItem{},
				NextPage: nil,
				Count:    0,
			}, nil
		},
	}

	service := NewUserService(mockCfg, nil, mockUserAPI)
	cmd := setupListTestCommand()
	cmd.Flags().Set("enabled", "true")
	cmd.Flags().Set("search", "bob")
	cmd.Flags().Set("sort-key", "username")
	cmd.Flags().Set("sort-order", "desc")

	err := service.ListUsers(cmd, nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	output := cmd.OutOrStdout().(*bytes.Buffer).String()
	if !strings.Contains(output, "No IAM users found") {
		t.Fatalf("Expected empty list message, got %q", output)
	}
}

func TestUserService_ListUsers_Pagination(t *testing.T) {
	mockCfg := configuration_handler.NewMockConfigurationHandler()
	mockCfg.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{
			APIKey:         "test-api-key",
			OrganizationID: "test-org-id",
			Output:         configuration_models.OutputHuman,
			Endpoints:      configuration_models.EndpointsV2{IAM: "https://iam.example.com"},
		}, nil
	}

	pageTwo := 2
	callCount := 0
	mockUserAPI := &api.MockUserAPI{
		ListIAMUsersFunc: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			enabled *bool,
			search string,
			page int,
			items int,
			sortKey string,
			sortOrder string,
		) (*api.GenericPaginatedResponse[api.IAMUserListItem], error) {
			callCount++
			if page == 1 {
				return &api.GenericPaginatedResponse[api.IAMUserListItem]{
					Data:     []api.IAMUserListItem{{ID: "user-001", Username: "page1-user"}},
					NextPage: &pageTwo,
					Count:    2,
				}, nil
			}
			return &api.GenericPaginatedResponse[api.IAMUserListItem]{
				Data:     []api.IAMUserListItem{{ID: "user-002", Username: "page2-user"}},
				NextPage: nil,
				Count:    2,
			}, nil
		},
	}

	service := NewUserService(mockCfg, nil, mockUserAPI)
	cmd := setupListTestCommand()

	err := service.ListUsers(cmd, nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if callCount != 2 {
		t.Fatalf("Expected 2 API calls for pagination, got %d", callCount)
	}

	output := cmd.OutOrStdout().(*bytes.Buffer).String()
	if !strings.Contains(output, "page1-user") || !strings.Contains(output, "page2-user") {
		t.Fatalf("Expected output to contain users from both pages, got %q", output)
	}
}

func TestUserService_ListUsers_SinglePage(t *testing.T) {
	mockCfg := configuration_handler.NewMockConfigurationHandler()
	mockCfg.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{
			APIKey:         "test-api-key",
			OrganizationID: "test-org-id",
			Output:         configuration_models.OutputHuman,
			Endpoints:      configuration_models.EndpointsV2{IAM: "https://iam.example.com"},
		}, nil
	}

	callCount := 0
	mockUserAPI := &api.MockUserAPI{
		ListIAMUsersFunc: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			enabled *bool,
			search string,
			page int,
			items int,
			sortKey string,
			sortOrder string,
		) (*api.GenericPaginatedResponse[api.IAMUserListItem], error) {
			callCount++
			if page != 3 {
				t.Fatalf("Expected page=3, got %d", page)
			}
			if items != 50 {
				t.Fatalf("Expected items=50, got %d", items)
			}
			return &api.GenericPaginatedResponse[api.IAMUserListItem]{
				Data:     []api.IAMUserListItem{{ID: "user-003", Username: "single-page-user"}},
				NextPage: nil,
				Count:    1,
			}, nil
		},
	}

	service := NewUserService(mockCfg, nil, mockUserAPI)
	cmd := setupListTestCommand()
	cmd.Flags().Set("page", "3")
	cmd.Flags().Set("items", "50")

	err := service.ListUsers(cmd, nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if callCount != 1 {
		t.Fatalf("Expected 1 API call for single page mode, got %d", callCount)
	}
}

func TestUserService_ListUsers_InvalidEnabled(t *testing.T) {
	mockCfg := configuration_handler.NewMockConfigurationHandler()
	mockCfg.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{
			APIKey:         "test-api-key",
			OrganizationID: "test-org-id",
			Output:         configuration_models.OutputHuman,
			Endpoints:      configuration_models.EndpointsV2{IAM: "https://iam.example.com"},
		}, nil
	}

	service := NewUserService(mockCfg, nil, &api.MockUserAPI{})
	cmd := setupListTestCommand()
	cmd.Flags().Set("enabled", "not-a-bool")

	err := service.ListUsers(cmd, nil)
	if err == nil {
		t.Fatal("Expected error for invalid enabled value, got nil")
	}
	if !strings.Contains(err.Error(), "invalid value for --enabled") {
		t.Fatalf("Expected invalid enabled error, got %v", err)
	}
}
