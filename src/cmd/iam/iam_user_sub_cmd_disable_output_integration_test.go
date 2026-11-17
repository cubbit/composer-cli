package cmd_iam

import (
	"testing"
	"time"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"

	"bytes"

	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"

	"github.com/cubbit/composer-cli/src/service/user"
)

func TestIAMUserSubCmd_Disable_Output_JSON(t *testing.T) {
	frozen := time.Date(2026, 6, 24, 20, 24, 13, 0, time.UTC)
	orgID := "test-org-id"
	organizationName := "test-org"

	mockUserAPI := &api.MockUserAPI{
		UpdateIAMUserFunc: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			userID string,
			request *api.UpdateIAMUserRequestBody,
		) (*api.IAMUser, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}
			if userID != "550e8400-e29b-41d4-a716-446655440000" {
				t.Fatalf("Expected user ID to be propagated, got %q", userID)
			}
			if request.Enabled == nil || *request.Enabled != false {
				t.Fatalf("Expected enabled=false, got %+v", request.Enabled)
			}

			return &api.IAMUser{
				ID:        userID,
				Username:  "alice.wonder",
				FirstName: "Alice",
				LastName:  "Wonder",
				Enabled:   false,
				Status:    "inactive",
				CreatedAt: frozen,
				Emails: []api.IAMUserEmail{
					{Email: "alice@example.com", Default: true},
				},
				Policies:         []api.IAMUserPolicy{},
				OrganizationID:   &orgID,
				OrganizationName: &organizationName,
			}, nil
		},
	}

	iamCmd, commandOutput := setupIAMUserIntegrationCommand(nil, mockUserAPI)
	iamCmd.SetArgs([]string{
		"user", "disable",
		"--user-id", "550e8400-e29b-41d4-a716-446655440000",
		"--output", "json",
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "username": "alice.wonder",
  "first_name": "Alice",
  "last_name": "Wonder",
  "email": "",
  "emails": [
    {
      "id": "",
      "email": "alice@example.com",
      "verified": false,
      "default": true,
      "created_at": "0001-01-01T00:00:00Z",
      "organization_id": null
    }
  ],
  "enabled": false,
  "internal": false,
  "banned": false,
  "is_root": false,
  "two_factor_enabled": false,
  "status": "inactive",
  "created_at": "2026-06-24T20:24:13Z",
  "deleted_at": null,
  "last_activity_at": null,
  "max_allowed_projects": 0,
  "policies_count": 0,
  "policies": [],
  "organization_id": "test-org-id",
  "organization_name": "test-org"
}
`
	if commandOutput.String() != expected {
		t.Errorf("JSON snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

func TestIAMUserSubCmd_Disable_Output_YAML(t *testing.T) {
	frozen := time.Date(2026, 6, 24, 20, 24, 13, 0, time.UTC)
	orgID := "test-org-id"
	organizationName := "test-org"

	mockUserAPI := &api.MockUserAPI{
		UpdateIAMUserFunc: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			userID string,
			request *api.UpdateIAMUserRequestBody,
		) (*api.IAMUser, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}
			if userID != "550e8400-e29b-41d4-a716-446655440001" {
				t.Fatalf("Expected user ID to be propagated, got %q", userID)
			}
			if request.Enabled == nil || *request.Enabled != false {
				t.Fatalf("Expected enabled=false, got %+v", request.Enabled)
			}

			return &api.IAMUser{
				ID:        userID,
				Username:  "alice.wonder",
				FirstName: "Alice",
				LastName:  "Wonder",
				Enabled:   false,
				Status:    "inactive",
				CreatedAt: frozen,
				Emails: []api.IAMUserEmail{
					{Email: "alice@example.com", Default: true},
				},
				Policies:         []api.IAMUserPolicy{},
				OrganizationID:   &orgID,
				OrganizationName: &organizationName,
			}, nil
		},
	}

	iamCmd, commandOutput := setupIAMUserIntegrationCommand(nil, mockUserAPI)
	iamCmd.SetArgs([]string{
		"user", "disable",
		"--user-id", "550e8400-e29b-41d4-a716-446655440001",
		"--output", "yaml",
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `id: 550e8400-e29b-41d4-a716-446655440001
username: alice.wonder
first_name: Alice
last_name: Wonder
email: ""
emails:
    - id: ""
      email: alice@example.com
      verified: false
      default: true
      created_at: 0001-01-01T00:00:00Z
      organization_id: null
enabled: false
internal: false
banned: false
is_root: false
two_factor_enabled: false
status: inactive
created_at: 2026-06-24T20:24:13Z
deleted_at: null
last_activity_at: null
max_allowed_projects: 0
policies_count: 0
policies: []
organization_id: test-org-id
organization_name: test-org
`
	if commandOutput.String() != expected {
		t.Errorf("YAML snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

func TestIAMUserSubCmd_Disable_Output_JSON_FromProfile(t *testing.T) {
	frozen := time.Date(2026, 6, 24, 20, 24, 13, 0, time.UTC)
	orgID := "test-org-id"
	organizationName := "test-org"

	mockUserAPI := &api.MockUserAPI{
		UpdateIAMUserFunc: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			userID string,
			request *api.UpdateIAMUserRequestBody,
		) (*api.IAMUser, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}
			if userID != "550e8400-e29b-41d4-a716-446655440000" {
				t.Fatalf("Expected user ID to be propagated, got %q", userID)
			}
			if request.Enabled == nil || *request.Enabled != false {
				t.Fatalf("Expected enabled=false, got %+v", request.Enabled)
			}

			return &api.IAMUser{
				ID:        userID,
				Username:  "alice.wonder",
				FirstName: "Alice",
				LastName:  "Wonder",
				Enabled:   false,
				Status:    "inactive",
				CreatedAt: frozen,
				Emails: []api.IAMUserEmail{
					{Email: "alice@example.com", Default: true},
				},
				Policies:         []api.IAMUserPolicy{},
				OrganizationID:   &orgID,
				OrganizationName: &organizationName,
			}, nil
		},
	}
	mockCfg := configuration_handler.NewMockConfigurationHandler()
	mockCfg.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{
			APIKey:         "test-api-key",
			OrganizationID: "test-org-id",
			Output:         "json",
			Endpoints: configuration_models.EndpointsV2{
				IAM: "https://iam.example.com",
			},
		}, nil
	}

	userService := user.NewUserService(mockCfg, nil, mockUserAPI)
	iamCmd := NewIAMCmd(userService)
	iamCmd.PersistentFlags().String("profile", "", "Profile")
	iamCmd.PersistentFlags().String("output", "human", "Output format")
	iamCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	iamCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	commandOutput := new(bytes.Buffer)
	iamCmd.SetOut(commandOutput)
	iamCmd.SetErr(commandOutput)
	iamCmd.SetArgs([]string{
		"user", "disable",
		"--user-id", "550e8400-e29b-41d4-a716-446655440000",
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "username": "alice.wonder",
  "first_name": "Alice",
  "last_name": "Wonder",
  "email": "",
  "emails": [
    {
      "id": "",
      "email": "alice@example.com",
      "verified": false,
      "default": true,
      "created_at": "0001-01-01T00:00:00Z",
      "organization_id": null
    }
  ],
  "enabled": false,
  "internal": false,
  "banned": false,
  "is_root": false,
  "two_factor_enabled": false,
  "status": "inactive",
  "created_at": "2026-06-24T20:24:13Z",
  "deleted_at": null,
  "last_activity_at": null,
  "max_allowed_projects": 0,
  "policies_count": 0,
  "policies": [],
  "organization_id": "test-org-id",
  "organization_name": "test-org"
}
`
	if commandOutput.String() != expected {
		t.Errorf("JSON snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

func TestIAMUserSubCmd_Disable_Output_YAML_FromProfile(t *testing.T) {
	frozen := time.Date(2026, 6, 24, 20, 24, 13, 0, time.UTC)
	orgID := "test-org-id"
	organizationName := "test-org"

	mockUserAPI := &api.MockUserAPI{
		UpdateIAMUserFunc: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			userID string,
			request *api.UpdateIAMUserRequestBody,
		) (*api.IAMUser, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}
			if userID != "550e8400-e29b-41d4-a716-446655440001" {
				t.Fatalf("Expected user ID to be propagated, got %q", userID)
			}
			if request.Enabled == nil || *request.Enabled != false {
				t.Fatalf("Expected enabled=false, got %+v", request.Enabled)
			}

			return &api.IAMUser{
				ID:        userID,
				Username:  "alice.wonder",
				FirstName: "Alice",
				LastName:  "Wonder",
				Enabled:   false,
				Status:    "inactive",
				CreatedAt: frozen,
				Emails: []api.IAMUserEmail{
					{Email: "alice@example.com", Default: true},
				},
				Policies:         []api.IAMUserPolicy{},
				OrganizationID:   &orgID,
				OrganizationName: &organizationName,
			}, nil
		},
	}
	mockCfg := configuration_handler.NewMockConfigurationHandler()
	mockCfg.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{
			APIKey:         "test-api-key",
			OrganizationID: "test-org-id",
			Output:         "yaml",
			Endpoints: configuration_models.EndpointsV2{
				IAM: "https://iam.example.com",
			},
		}, nil
	}

	userService := user.NewUserService(mockCfg, nil, mockUserAPI)
	iamCmd := NewIAMCmd(userService)
	iamCmd.PersistentFlags().String("profile", "", "Profile")
	iamCmd.PersistentFlags().String("output", "human", "Output format")
	iamCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	iamCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	commandOutput := new(bytes.Buffer)
	iamCmd.SetOut(commandOutput)
	iamCmd.SetErr(commandOutput)
	iamCmd.SetArgs([]string{
		"user", "disable",
		"--user-id", "550e8400-e29b-41d4-a716-446655440001",
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `id: 550e8400-e29b-41d4-a716-446655440001
username: alice.wonder
first_name: Alice
last_name: Wonder
email: ""
emails:
    - id: ""
      email: alice@example.com
      verified: false
      default: true
      created_at: 0001-01-01T00:00:00Z
      organization_id: null
enabled: false
internal: false
banned: false
is_root: false
two_factor_enabled: false
status: inactive
created_at: 2026-06-24T20:24:13Z
deleted_at: null
last_activity_at: null
max_allowed_projects: 0
policies_count: 0
policies: []
organization_id: test-org-id
organization_name: test-org
`
	if commandOutput.String() != expected {
		t.Errorf("YAML snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}
