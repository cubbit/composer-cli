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

func TestIAMUserSubCmd_List_Output_JSON(t *testing.T) {
	frozen := time.Date(2026, 6, 24, 20, 24, 13, 0, time.UTC)
	firstName := "Alice"
	lastName := "Wonder"
	disabledName := "Bob"

	mockUserAPI := &api.MockUserAPI{
		ListIAMUsersFunc: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			_ *bool,
			_ string,
			_ int,
			_ int,
			_ string,
			_ string,
		) (*api.GenericPaginatedResponse[api.IAMUserListItem], error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}

			return &api.GenericPaginatedResponse[api.IAMUserListItem]{
				Data: []api.IAMUserListItem{
					{
						ID:        "user-001",
						Username:  "alice.wonder",
						FirstName: &firstName,
						LastName:  &lastName,
						Enabled:   true,
						CreatedAt: frozen,
						Emails: []api.IAMUserEmail{
							{Email: "alice@example.com", Default: true},
					},
						Status: "active",
				},
					{
						ID:        "user-002",
						Username:  "bob.builder",
						FirstName: &disabledName,
						Enabled:   false,
						CreatedAt: frozen.Add(time.Hour),
						Emails: []api.IAMUserEmail{
							{Email: "bob@example.com"},
					},
						Status: "pending",
				},
				},
				NextPage: nil,
				Count:    2,
			}, nil
		},
	}

	iamCmd, commandOutput := setupIAMUserIntegrationCommand(iamUserCommandChallengeAPI{}, mockUserAPI)
	iamCmd.SetArgs([]string{
		"user", "list",
		"--output", "json",
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `[
  {
    "id": "user-001",
    "username": "alice.wonder",
    "first_name": "Alice",
    "last_name": "Wonder",
    "enabled": true,
    "created_at": "2026-06-24T20:24:13Z",
    "deleted_at": null,
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
    "policies_count": 0,
    "status": "active",
    "is_root": false,
    "two_factor_enabled": false,
    "organization_id": "",
    "last_activity_at": null
  },
  {
    "id": "user-002",
    "username": "bob.builder",
    "first_name": "Bob",
    "last_name": null,
    "enabled": false,
    "created_at": "2026-06-24T21:24:13Z",
    "deleted_at": null,
    "emails": [
      {
        "id": "",
        "email": "bob@example.com",
        "verified": false,
        "default": false,
        "created_at": "0001-01-01T00:00:00Z",
        "organization_id": null
      }
    ],
    "policies_count": 0,
    "status": "pending",
    "is_root": false,
    "two_factor_enabled": false,
    "organization_id": "",
    "last_activity_at": null
  }
]
`
	if commandOutput.String() != expected {
		t.Errorf("JSON snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

func TestIAMUserSubCmd_List_Output_YAML(t *testing.T) {
	frozen := time.Date(2026, 6, 24, 20, 24, 13, 0, time.UTC)
	firstName := "Alice"
	lastName := "Wonder"
	disabledName := "Bob"

	mockUserAPI := &api.MockUserAPI{
		ListIAMUsersFunc: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			_ *bool,
			_ string,
			_ int,
			_ int,
			_ string,
			_ string,
		) (*api.GenericPaginatedResponse[api.IAMUserListItem], error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}

			return &api.GenericPaginatedResponse[api.IAMUserListItem]{
				Data: []api.IAMUserListItem{
					{
						ID:        "user-001",
						Username:  "alice.wonder",
						FirstName: &firstName,
						LastName:  &lastName,
						Enabled:   true,
						CreatedAt: frozen,
						Emails: []api.IAMUserEmail{
							{Email: "alice@example.com", Default: true},
					},
						Status: "active",
				},
					{
						ID:        "user-002",
						Username:  "bob.builder",
						FirstName: &disabledName,
						Enabled:   false,
						CreatedAt: frozen.Add(time.Hour),
						Emails: []api.IAMUserEmail{
							{Email: "bob@example.com"},
					},
						Status: "pending",
				},
				},
				NextPage: nil,
				Count:    2,
			}, nil
		},
	}

	iamCmd, commandOutput := setupIAMUserIntegrationCommand(iamUserCommandChallengeAPI{}, mockUserAPI)
	iamCmd.SetArgs([]string{
		"user", "list",
		"--output", "yaml",
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `- id: user-001
  username: alice.wonder
  first_name: Alice
  last_name: Wonder
  enabled: true
  created_at: 2026-06-24T20:24:13Z
  deleted_at: null
  emails:
    - id: ""
      email: alice@example.com
      verified: false
      default: true
      created_at: 0001-01-01T00:00:00Z
      organization_id: null
  policies_count: 0
  status: active
  is_root: false
  two_factor_enabled: false
  organization_id: ""
  last_activity_at: null
- id: user-002
  username: bob.builder
  first_name: Bob
  last_name: null
  enabled: false
  created_at: 2026-06-24T21:24:13Z
  deleted_at: null
  emails:
    - id: ""
      email: bob@example.com
      verified: false
      default: false
      created_at: 0001-01-01T00:00:00Z
      organization_id: null
  policies_count: 0
  status: pending
  is_root: false
  two_factor_enabled: false
  organization_id: ""
  last_activity_at: null
`
	if commandOutput.String() != expected {
		t.Errorf("YAML snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

func TestIAMUserSubCmd_List_Output_JSON_FromProfile(t *testing.T) {
	frozen := time.Date(2026, 6, 24, 20, 24, 13, 0, time.UTC)
	firstName := "Alice"
	lastName := "Wonder"
	disabledName := "Bob"

	mockUserAPI := &api.MockUserAPI{
		ListIAMUsersFunc: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			_ *bool,
			_ string,
			_ int,
			_ int,
			_ string,
			_ string,
		) (*api.GenericPaginatedResponse[api.IAMUserListItem], error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}

			return &api.GenericPaginatedResponse[api.IAMUserListItem]{
				Data: []api.IAMUserListItem{
					{
						ID:        "user-001",
						Username:  "alice.wonder",
						FirstName: &firstName,
						LastName:  &lastName,
						Enabled:   true,
						CreatedAt: frozen,
						Emails: []api.IAMUserEmail{
							{Email: "alice@example.com", Default: true},
					},
						Status: "active",
				},
					{
						ID:        "user-002",
						Username:  "bob.builder",
						FirstName: &disabledName,
						Enabled:   false,
						CreatedAt: frozen.Add(time.Hour),
						Emails: []api.IAMUserEmail{
							{Email: "bob@example.com"},
					},
						Status: "pending",
				},
				},
				NextPage: nil,
				Count:    2,
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
		"user", "list",
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `[
  {
    "id": "user-001",
    "username": "alice.wonder",
    "first_name": "Alice",
    "last_name": "Wonder",
    "enabled": true,
    "created_at": "2026-06-24T20:24:13Z",
    "deleted_at": null,
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
    "policies_count": 0,
    "status": "active",
    "is_root": false,
    "two_factor_enabled": false,
    "organization_id": "",
    "last_activity_at": null
  },
  {
    "id": "user-002",
    "username": "bob.builder",
    "first_name": "Bob",
    "last_name": null,
    "enabled": false,
    "created_at": "2026-06-24T21:24:13Z",
    "deleted_at": null,
    "emails": [
      {
        "id": "",
        "email": "bob@example.com",
        "verified": false,
        "default": false,
        "created_at": "0001-01-01T00:00:00Z",
        "organization_id": null
      }
    ],
    "policies_count": 0,
    "status": "pending",
    "is_root": false,
    "two_factor_enabled": false,
    "organization_id": "",
    "last_activity_at": null
  }
]
`
	if commandOutput.String() != expected {
		t.Errorf("JSON snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

func TestIAMUserSubCmd_List_Output_YAML_FromProfile(t *testing.T) {
	frozen := time.Date(2026, 6, 24, 20, 24, 13, 0, time.UTC)
	firstName := "Alice"
	lastName := "Wonder"
	disabledName := "Bob"

	mockUserAPI := &api.MockUserAPI{
		ListIAMUsersFunc: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			_ *bool,
			_ string,
			_ int,
			_ int,
			_ string,
			_ string,
		) (*api.GenericPaginatedResponse[api.IAMUserListItem], error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}

			return &api.GenericPaginatedResponse[api.IAMUserListItem]{
				Data: []api.IAMUserListItem{
					{
						ID:        "user-001",
						Username:  "alice.wonder",
						FirstName: &firstName,
						LastName:  &lastName,
						Enabled:   true,
						CreatedAt: frozen,
						Emails: []api.IAMUserEmail{
							{Email: "alice@example.com", Default: true},
					},
						Status: "active",
				},
					{
						ID:        "user-002",
						Username:  "bob.builder",
						FirstName: &disabledName,
						Enabled:   false,
						CreatedAt: frozen.Add(time.Hour),
						Emails: []api.IAMUserEmail{
							{Email: "bob@example.com"},
					},
						Status: "pending",
				},
				},
				NextPage: nil,
				Count:    2,
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
		"user", "list",
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `- id: user-001
  username: alice.wonder
  first_name: Alice
  last_name: Wonder
  enabled: true
  created_at: 2026-06-24T20:24:13Z
  deleted_at: null
  emails:
    - id: ""
      email: alice@example.com
      verified: false
      default: true
      created_at: 0001-01-01T00:00:00Z
      organization_id: null
  policies_count: 0
  status: active
  is_root: false
  two_factor_enabled: false
  organization_id: ""
  last_activity_at: null
- id: user-002
  username: bob.builder
  first_name: Bob
  last_name: null
  enabled: false
  created_at: 2026-06-24T21:24:13Z
  deleted_at: null
  emails:
    - id: ""
      email: bob@example.com
      verified: false
      default: false
      created_at: 0001-01-01T00:00:00Z
      organization_id: null
  policies_count: 0
  status: pending
  is_root: false
  two_factor_enabled: false
  organization_id: ""
  last_activity_at: null
`
	if commandOutput.String() != expected {
		t.Errorf("YAML snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}
