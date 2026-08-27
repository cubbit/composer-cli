package cmd_iam

import (
	"testing"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"

	"bytes"

	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"

	"github.com/cubbit/composer-cli/src/service/user"
)

func TestIAMUserSubCmd_Create_Output_JSON(t *testing.T) {
	email := "alice@example.com"
	mockUserAPI := &api.MockUserAPI{
		GetIAMUserSelfFunc: func(_ configuration_models.EndpointsV2, _ string, apiKey string) (*api.IAMUser, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			organizationName := "test-org"
			return &api.IAMUser{OrganizationName: &organizationName}, nil
		},
		BulkGenerateSaltsFunc: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			_ *api.BulkGenerateSaltsRequestBody,
		) (*api.BulkGenerateSaltsResponse, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}

			return &api.BulkGenerateSaltsResponse{
				Count: 1,
				Data: []api.BulkGenerateSaltResponseItem{
					{Username: "alice", Salt: "test-salt"},
				},
			}, nil
		},
		BulkCreateIAMUsersFunc: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			_ *api.BulkCreateIAMUsersRequestBody,
		) (*api.BulkCreateIAMUsersResponse, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}

			return &api.BulkCreateIAMUsersResponse{
				Count: 1,
				Data: []api.BulkCreateIAMUserResponseItem{
					{
						ID:       "user-001",
						Username: "alice",
						Email:    &email,
						Created:  true,
						Status:   "pending",
				},
				},
			}, nil
		},
	}

	iamCmd, commandOutput := setupIAMUserIntegrationCommand(iamUserCommandChallengeAPI{}, mockUserAPI)
	iamCmd.SetArgs([]string{
		"user", "create",
		"--username", "alice",
		"--password", "test-password",
		"--email", email,
		"--output", "json",
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `{
  "data": [
    {
      "id": "user-001",
      "username": "alice",
      "email": "alice@example.com",
      "created": true,
      "status": "pending"
    }
  ],
  "count": 1
}
`
	if commandOutput.String() != expected {
		t.Errorf("JSON snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

func TestIAMUserSubCmd_Create_Output_YAML(t *testing.T) {
	email := "alice@example.com"
	mockUserAPI := &api.MockUserAPI{
		GetIAMUserSelfFunc: func(_ configuration_models.EndpointsV2, _ string, apiKey string) (*api.IAMUser, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			organizationName := "test-org"
			return &api.IAMUser{OrganizationName: &organizationName}, nil
		},
		BulkGenerateSaltsFunc: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			_ *api.BulkGenerateSaltsRequestBody,
		) (*api.BulkGenerateSaltsResponse, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}

			return &api.BulkGenerateSaltsResponse{
				Count: 1,
				Data: []api.BulkGenerateSaltResponseItem{
					{Username: "alice", Salt: "test-salt"},
				},
			}, nil
		},
		BulkCreateIAMUsersFunc: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			_ *api.BulkCreateIAMUsersRequestBody,
		) (*api.BulkCreateIAMUsersResponse, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}

			return &api.BulkCreateIAMUsersResponse{
				Count: 2,
				Data: []api.BulkCreateIAMUserResponseItem{
					{
						ID:       "user-001",
						Username: "alice",
						Email:    &email,
						Created:  true,
						Status:   "pending",
				},
					{
						ID:       "user-002",
						Username: "bob",
						Created:  false,
						Status:   "active",
				},
				},
			}, nil
		},
	}

	iamCmd, commandOutput := setupIAMUserIntegrationCommand(iamUserCommandChallengeAPI{}, mockUserAPI)
	iamCmd.SetArgs([]string{
		"user", "create",
		"--username", "alice",
		"--password", "test-password",
		"--email", email,
		"--output", "yaml",
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `data:
    - id: user-001
      username: alice
      email: alice@example.com
      created: true
      status: pending
    - id: user-002
      username: bob
      email: null
      created: false
      status: active
count: 2
`
	if commandOutput.String() != expected {
		t.Errorf("YAML snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

func TestIAMUserSubCmd_Create_Output_JSON_FromProfile(t *testing.T) {
	email := "alice@example.com"
	mockUserAPI := &api.MockUserAPI{
		GetIAMUserSelfFunc: func(_ configuration_models.EndpointsV2, _ string, apiKey string) (*api.IAMUser, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			organizationName := "test-org"
			return &api.IAMUser{OrganizationName: &organizationName}, nil
		},
		BulkGenerateSaltsFunc: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			_ *api.BulkGenerateSaltsRequestBody,
		) (*api.BulkGenerateSaltsResponse, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}

			return &api.BulkGenerateSaltsResponse{
				Count: 1,
				Data: []api.BulkGenerateSaltResponseItem{
					{Username: "alice", Salt: "test-salt"},
				},
			}, nil
		},
		BulkCreateIAMUsersFunc: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			_ *api.BulkCreateIAMUsersRequestBody,
		) (*api.BulkCreateIAMUsersResponse, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}

			return &api.BulkCreateIAMUsersResponse{
				Count: 1,
				Data: []api.BulkCreateIAMUserResponseItem{
					{
						ID:       "user-001",
						Username: "alice",
						Email:    &email,
						Created:  true,
						Status:   "pending",
				},
				},
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
		"user", "create",
		"--username", "alice",
		"--password", "test-password",
		"--email", email,
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `{
  "data": [
    {
      "id": "user-001",
      "username": "alice",
      "email": "alice@example.com",
      "created": true,
      "status": "pending"
    }
  ],
  "count": 1
}
`
	if commandOutput.String() != expected {
		t.Errorf("JSON snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

func TestIAMUserSubCmd_Create_Output_YAML_FromProfile(t *testing.T) {
	email := "alice@example.com"
	mockUserAPI := &api.MockUserAPI{
		GetIAMUserSelfFunc: func(_ configuration_models.EndpointsV2, _ string, apiKey string) (*api.IAMUser, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			organizationName := "test-org"
			return &api.IAMUser{OrganizationName: &organizationName}, nil
		},
		BulkGenerateSaltsFunc: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			_ *api.BulkGenerateSaltsRequestBody,
		) (*api.BulkGenerateSaltsResponse, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}

			return &api.BulkGenerateSaltsResponse{
				Count: 1,
				Data: []api.BulkGenerateSaltResponseItem{
					{Username: "alice", Salt: "test-salt"},
				},
			}, nil
		},
		BulkCreateIAMUsersFunc: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			_ *api.BulkCreateIAMUsersRequestBody,
		) (*api.BulkCreateIAMUsersResponse, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}

			return &api.BulkCreateIAMUsersResponse{
				Count: 2,
				Data: []api.BulkCreateIAMUserResponseItem{
					{
						ID:       "user-001",
						Username: "alice",
						Email:    &email,
						Created:  true,
						Status:   "pending",
				},
					{
						ID:       "user-002",
						Username: "bob",
						Created:  false,
						Status:   "active",
				},
				},
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
		"user", "create",
		"--username", "alice",
		"--password", "test-password",
		"--email", email,
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `data:
    - id: user-001
      username: alice
      email: alice@example.com
      created: true
      status: pending
    - id: user-002
      username: bob
      email: null
      created: false
      status: active
count: 2
`
	if commandOutput.String() != expected {
		t.Errorf("YAML snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}
