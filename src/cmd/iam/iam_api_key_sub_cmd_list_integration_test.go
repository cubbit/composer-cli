package cmd_iam

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	apikey "github.com/cubbit/composer-cli/src/service/api_key"
	"github.com/cubbit/composer-cli/src/service/user"
	"github.com/spf13/cobra"
)

func setupIAMAPIKeyIntegrationCommand(
	authAPI api.AuthAPIInterface,
	userAPI api.UserAPIInterface,
) (*cobra.Command, *bytes.Buffer) {
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

	userService := user.NewUserService(mockCfg, authAPI, userAPI)
	apiKeyService := apikey.NewAPIKeyService(mockCfg, userAPI)
	iamCmd := NewIAMCmd(userService, apiKeyService)
	iamCmd.PersistentFlags().String("profile", "", "Profile")
	iamCmd.PersistentFlags().String("output", "human", "Output format")
	iamCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	iamCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	commandOutput := new(bytes.Buffer)
	iamCmd.SetOut(commandOutput)
	iamCmd.SetErr(commandOutput)

	return iamCmd, commandOutput
}

func TestIAMAPIKeySubCmd_List_Integration_HumanOutput(t *testing.T) {
	createdAt := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)
	expiresAt := time.Date(2024, 12, 31, 23, 59, 0, 0, time.UTC)
	organizationName := "test-org"
	mockUserAPI := &api.MockUserAPI{
		GetIAMUserSelfFunc: func(endpoints configuration_models.EndpointsV2, accessToken string, apiKey string) (*api.IAMUser, error) {
			return &api.IAMUser{
				ID:               "operator-001",
				Username:         "admin",
				OrganizationName: &organizationName,
			}, nil
		},
		ListIAMAPIKeysFunc: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			operatorID string,
			page int,
			items int,
			sortKey string,
			sortOrder string,
		) (*api.GenericPaginatedResponse[api.OperatorAPIKey], error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}
			if operatorID != "operator-001" {
				t.Fatalf("Expected operator ID to be propagated, got %q", operatorID)
			}

			return &api.GenericPaginatedResponse[api.OperatorAPIKey]{
				Data: []api.OperatorAPIKey{
					{
						ID:         "key-001",
						Name:       "automation",
						OperatorID: "operator-001",
						CreatedAt:  createdAt,
						ExpiresAt:  &expiresAt,
						Enabled:    true,
					},
				},
			}, nil
		},
	}

	iamCmd, commandOutput := setupIAMAPIKeyIntegrationCommand(nil, mockUserAPI)
	iamCmd.SetArgs([]string{"api-key", "list"})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedResult := strings.TrimSpace(`
╭────────────┬─────────┬─────────┬─────────────────────┬─────────────────────╮
│ Name       │ ID      │ Enabled │ Created At          │ Expires At          │
├────────────┼─────────┼─────────┼─────────────────────┼─────────────────────┤
│ automation │ key-001 │ true    │ 2024-01-15 10:30:00 │ 2024-12-31 23:59:00 │
╰────────────┴─────────┴─────────┴─────────────────────┴─────────────────────╯
`)

	actualResult := strings.TrimSpace(commandOutput.String())
	if actualResult != expectedResult {
		t.Fatalf("Expected IAM API key list output does not match actual output\nExpected:\n%s\nActual:\n%s", expectedResult, actualResult)
	}
}

func TestIAMAPIKeySubCmd_List_Integration_WithPaginationAndSorting(t *testing.T) {
	organizationName := "test-org"
	mockUserAPI := &api.MockUserAPI{
		GetIAMUserSelfFunc: func(endpoints configuration_models.EndpointsV2, accessToken string, apiKey string) (*api.IAMUser, error) {
			return &api.IAMUser{
				ID:               "operator-001",
				Username:         "admin",
				OrganizationName: &organizationName,
			}, nil
		},
		ListIAMAPIKeysFunc: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			operatorID string,
			page int,
			items int,
			sortKey string,
			sortOrder string,
		) (*api.GenericPaginatedResponse[api.OperatorAPIKey], error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}
			if operatorID != "operator-001" {
				t.Fatalf("Expected operator ID to be propagated, got %q", operatorID)
			}
			if page != 2 || items != 50 {
				t.Fatalf("Expected page=2 items=50, got page=%d items=%d", page, items)
			}
			if sortKey != "created_at" {
				t.Fatalf("Expected sort-key created_at, got %q", sortKey)
			}
			if sortOrder != "desc" {
				t.Fatalf("Expected sort-order desc, got %q", sortOrder)
			}

			return &api.GenericPaginatedResponse[api.OperatorAPIKey]{
				Data:     []api.OperatorAPIKey{},
				NextPage: nil,
				Count:    0,
			}, nil
		},
	}

	iamCmd, commandOutput := setupIAMAPIKeyIntegrationCommand(nil, mockUserAPI)
	iamCmd.SetArgs([]string{
		"api-key",
		"list",
		"--page", "2",
		"--items", "50",
		"--sort-key", "created_at",
		"--sort-order", "desc",
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedResult := "No IAM API keys found.\n"
	actualResult := commandOutput.String()
	if actualResult != expectedResult {
		t.Fatalf("Expected output %q, got %q", expectedResult, actualResult)
	}
}

func TestIAMAPIKeySubCmd_List_Integration_WithUsername(t *testing.T) {
	mockUserAPI := &api.MockUserAPI{
		GetIAMUserSelfFunc: func(endpoints configuration_models.EndpointsV2, accessToken string, apiKey string) (*api.IAMUser, error) {
			t.Fatal("List with --username should not resolve the current IAM user")
			return nil, nil
		},
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
			if search != "alice" {
				t.Fatalf("Expected username search %q, got %q", "alice", search)
			}

			return &api.GenericPaginatedResponse[api.IAMUserListItem]{
				Data: []api.IAMUserListItem{
					{ID: "operator-002", Username: "alice"},
				},
			}, nil
		},
		ListIAMAPIKeysFunc: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			operatorID string,
			page int,
			items int,
			sortKey string,
			sortOrder string,
		) (*api.GenericPaginatedResponse[api.OperatorAPIKey], error) {
			if operatorID != "operator-002" {
				t.Fatalf("Expected target operator ID %q, got %q", "operator-002", operatorID)
			}

			return &api.GenericPaginatedResponse[api.OperatorAPIKey]{Data: []api.OperatorAPIKey{}}, nil
		},
	}

	iamCmd, commandOutput := setupIAMAPIKeyIntegrationCommand(nil, mockUserAPI)
	iamCmd.SetArgs([]string{"api-key", "list", "--username", "alice"})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedResult := "No IAM API keys found.\n"
	if commandOutput.String() != expectedResult {
		t.Fatalf("Expected output %q, got %q", expectedResult, commandOutput.String())
	}
}

func TestIAMAPIKeySubCmd_List_Integration_RejectsMultipleUserTargets(t *testing.T) {
	mockUserAPI := &api.MockUserAPI{
		GetIAMUserSelfFunc: func(endpoints configuration_models.EndpointsV2, accessToken string, apiKey string) (*api.IAMUser, error) {
			t.Fatal("List with invalid target flags should not resolve the current IAM user")
			return nil, nil
		},
		ListIAMAPIKeysFunc: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			operatorID string,
			page int,
			items int,
			sortKey string,
			sortOrder string,
		) (*api.GenericPaginatedResponse[api.OperatorAPIKey], error) {
			t.Fatal("List with invalid target flags should not list API keys")
			return nil, nil
		},
	}

	iamCmd, commandOutput := setupIAMAPIKeyIntegrationCommand(nil, mockUserAPI)
	iamCmd.SetArgs([]string{"api-key", "list", "--user-id", "operator-002", "--username", "alice"})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no command execution error, got %v", err)
	}
	if !strings.Contains(commandOutput.String(), "specify at most one of --user-id or --username") {
		t.Fatalf("Expected conflicting target flags error, got %q", commandOutput.String())
	}
}
