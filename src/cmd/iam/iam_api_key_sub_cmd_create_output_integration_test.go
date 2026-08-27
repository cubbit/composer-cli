package cmd_iam

import (
	"bytes"
	"testing"
	"time"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	apikey "github.com/cubbit/composer-cli/src/service/api_key"
	"github.com/cubbit/composer-cli/src/service/user"
)

func TestIAMAPIKeySubCmd_Create_Output_JSON(t *testing.T) {
	frozen := time.Date(2026, 6, 24, 20, 24, 13, 0, time.UTC)
	organizationName := "test-org"
	mockUserAPI := &api.MockUserAPI{
		GetIAMUserSelfFunc: func(_ configuration_models.EndpointsV2, _ string, apiKey string) (*api.IAMUser, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			return &api.IAMUser{
				ID:               "operator-001",
				Username:         "admin",
				OrganizationName: &organizationName,
			}, nil
		},
		CreateIAMAPIKeyFunc: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			operatorID string,
			request *api.CreateIAMAPIKeyRequestBody,
		) (*api.OperatorAPIKey, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}
			if operatorID != "operator-001" {
				t.Fatalf("Expected operator ID to be propagated, got %q", operatorID)
			}
			if request.Name != "automation" {
				t.Fatalf("Expected API key name %q, got %q", "automation", request.Name)
			}

			return &api.OperatorAPIKey{
				ID:         "key-001",
				Name:       "automation",
				Key:        "sk-1234567890abcdef",
				OperatorID: "operator-001",
				CreatedAt:  frozen,
				Enabled:    true,
			}, nil
		},
	}

	iamCmd, commandOutput := setupIAMAPIKeyIntegrationCommand(nil, mockUserAPI)
	iamCmd.SetArgs([]string{
		"api-key", "create",
		"--name", "automation",
		"--output", "json",
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `{
  "id": "key-001",
  "name": "automation",
  "key": "sk-1234567890abcdef",
  "operator_id": "operator-001",
  "created_at": "2026-06-24T20:24:13Z",
  "expires_at": null,
  "deleted_at": null,
  "banned_at": null,
  "enabled": true
}
`
	if commandOutput.String() != expected {
		t.Errorf("JSON snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

func TestIAMAPIKeySubCmd_Create_Output_YAML(t *testing.T) {
	frozen := time.Date(2026, 6, 24, 20, 24, 13, 0, time.UTC)
	organizationName := "test-org"
	mockUserAPI := &api.MockUserAPI{
		GetIAMUserSelfFunc: func(_ configuration_models.EndpointsV2, _ string, apiKey string) (*api.IAMUser, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			return &api.IAMUser{
				ID:               "operator-001",
				Username:         "admin",
				OrganizationName: &organizationName,
			}, nil
		},
		CreateIAMAPIKeyFunc: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			operatorID string,
			request *api.CreateIAMAPIKeyRequestBody,
		) (*api.OperatorAPIKey, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}
			if operatorID != "operator-001" {
				t.Fatalf("Expected operator ID to be propagated, got %q", operatorID)
			}
			if request.Name != "automation" {
				t.Fatalf("Expected API key name %q, got %q", "automation", request.Name)
			}

			return &api.OperatorAPIKey{
				ID:         "key-002",
				Name:       "automation",
				Key:        "sk-1234567890abcdef",
				OperatorID: "operator-001",
				CreatedAt:  frozen,
				Enabled:    true,
			}, nil
		},
	}

	iamCmd, commandOutput := setupIAMAPIKeyIntegrationCommand(nil, mockUserAPI)
	iamCmd.SetArgs([]string{
		"api-key", "create",
		"--name", "automation",
		"--output", "yaml",
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `id: key-002
name: automation
key: sk-1234567890abcdef
operator_id: operator-001
created_at: 2026-06-24T20:24:13Z
expires_at: null
deleted_at: null
banned_at: null
enabled: true

`
	if commandOutput.String() != expected {
		t.Errorf("YAML snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

func TestIAMAPIKeySubCmd_Create_Output_JSON_FromProfile(t *testing.T) {
	frozen := time.Date(2026, 6, 24, 20, 24, 13, 0, time.UTC)
	organizationName := "test-org"
	mockUserAPI := &api.MockUserAPI{
		GetIAMUserSelfFunc: func(_ configuration_models.EndpointsV2, _ string, apiKey string) (*api.IAMUser, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			return &api.IAMUser{
				ID:               "operator-001",
				Username:         "admin",
				OrganizationName: &organizationName,
			}, nil
		},
		CreateIAMAPIKeyFunc: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			operatorID string,
			request *api.CreateIAMAPIKeyRequestBody,
		) (*api.OperatorAPIKey, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}
			if operatorID != "operator-001" {
				t.Fatalf("Expected operator ID to be propagated, got %q", operatorID)
			}
			if request.Name != "automation" {
				t.Fatalf("Expected API key name %q, got %q", "automation", request.Name)
			}

			return &api.OperatorAPIKey{
				ID:         "key-001",
				Name:       "automation",
				Key:        "sk-1234567890abcdef",
				OperatorID: "operator-001",
				CreatedAt:  frozen,
				Enabled:    true,
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
	apiKeyService := apikey.NewAPIKeyService(mockCfg, mockUserAPI)
	iamCmd := NewIAMCmd(userService, apiKeyService)
	iamCmd.PersistentFlags().String("profile", "", "Profile")
	iamCmd.PersistentFlags().String("output", "human", "Output format")
	iamCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	iamCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	commandOutput := new(bytes.Buffer)
	iamCmd.SetOut(commandOutput)
	iamCmd.SetErr(commandOutput)

	iamCmd.SetArgs([]string{
		"api-key", "create",
		"--name", "automation",
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `{
  "id": "key-001",
  "name": "automation",
  "key": "sk-1234567890abcdef",
  "operator_id": "operator-001",
  "created_at": "2026-06-24T20:24:13Z",
  "expires_at": null,
  "deleted_at": null,
  "banned_at": null,
  "enabled": true
}
`
	if commandOutput.String() != expected {
		t.Errorf("JSON snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

func TestIAMAPIKeySubCmd_Create_Output_YAML_FromProfile(t *testing.T) {
	frozen := time.Date(2026, 6, 24, 20, 24, 13, 0, time.UTC)
	organizationName := "test-org"
	mockUserAPI := &api.MockUserAPI{
		GetIAMUserSelfFunc: func(_ configuration_models.EndpointsV2, _ string, apiKey string) (*api.IAMUser, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			return &api.IAMUser{
				ID:               "operator-001",
				Username:         "admin",
				OrganizationName: &organizationName,
			}, nil
		},
		CreateIAMAPIKeyFunc: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			operatorID string,
			request *api.CreateIAMAPIKeyRequestBody,
		) (*api.OperatorAPIKey, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}
			if operatorID != "operator-001" {
				t.Fatalf("Expected operator ID to be propagated, got %q", operatorID)
			}
			if request.Name != "automation" {
				t.Fatalf("Expected API key name %q, got %q", "automation", request.Name)
			}

			return &api.OperatorAPIKey{
				ID:         "key-002",
				Name:       "automation",
				Key:        "sk-1234567890abcdef",
				OperatorID: "operator-001",
				CreatedAt:  frozen,
				Enabled:    true,
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
	apiKeyService := apikey.NewAPIKeyService(mockCfg, mockUserAPI)
	iamCmd := NewIAMCmd(userService, apiKeyService)
	iamCmd.PersistentFlags().String("profile", "", "Profile")
	iamCmd.PersistentFlags().String("output", "human", "Output format")
	iamCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	iamCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	commandOutput := new(bytes.Buffer)
	iamCmd.SetOut(commandOutput)
	iamCmd.SetErr(commandOutput)

	iamCmd.SetArgs([]string{
		"api-key", "create",
		"--name", "automation",
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `id: key-002
name: automation
key: sk-1234567890abcdef
operator_id: operator-001
created_at: 2026-06-24T20:24:13Z
expires_at: null
deleted_at: null
banned_at: null
enabled: true

`
	if commandOutput.String() != expected {
		t.Errorf("YAML snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}
