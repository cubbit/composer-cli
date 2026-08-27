package cmd_auth

import (
	"bytes"
	"testing"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/src/service"
)

func TestAuthSubCmd_Login_Output_JSON(t *testing.T) {
	mockCfg := newTestAuthConfig()

	mockAuthAPI := &api.MockAuthAPI{}

	orgID := "test-org-id"
	mockUserAPI := &api.MockUserAPI{
		GetIAMUserSelfFunc: func(
			_ configuration_models.EndpointsV2,
			_ string,
			apiKey string,
		) (*api.IAMUser, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			return &api.IAMUser{OrganizationID: &orgID}, nil
		},
	}

	authService := service.NewAuthService(mockCfg, mockAuthAPI, mockUserAPI)

	authCmd := NewAuthCmd(authService)
	authCmd.PersistentFlags().String("output", "human", "Output format")
	authCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	authCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	commandOutput := new(bytes.Buffer)
	authCmd.SetOut(commandOutput)
	authCmd.SetErr(commandOutput)
	authCmd.SetArgs([]string{
		"login",
		"--profile", "test-profile",
		"--api-key", "test-api-key",
		"--output", "json",
	})

	err := authCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `{
  "message": "✅ Authentication successful!\n"
}
`
	if commandOutput.String() != expected {
		t.Errorf("JSON snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

func TestAuthSubCmd_Login_Output_YAML(t *testing.T) {
	mockCfg := newTestAuthConfig()

	mockAuthAPI := &api.MockAuthAPI{}

	orgID := "test-org-id"
	mockUserAPI := &api.MockUserAPI{
		GetIAMUserSelfFunc: func(
			_ configuration_models.EndpointsV2,
			_ string,
			apiKey string,
		) (*api.IAMUser, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			return &api.IAMUser{OrganizationID: &orgID}, nil
		},
	}

	authService := service.NewAuthService(mockCfg, mockAuthAPI, mockUserAPI)

	authCmd := NewAuthCmd(authService)
	authCmd.PersistentFlags().String("output", "human", "Output format")
	authCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	authCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	commandOutput := new(bytes.Buffer)
	authCmd.SetOut(commandOutput)
	authCmd.SetErr(commandOutput)
	authCmd.SetArgs([]string{
		"login",
		"--profile", "test-profile",
		"--api-key", "test-api-key",
		"--output", "yaml",
	})

	err := authCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `message: |
    ✅ Authentication successful!
`
	if commandOutput.String() != expected {
		t.Errorf("YAML snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

func TestAuthSubCmd_Login_Output_JSON_FromProfile(t *testing.T) {
	mockCfg := newTestAuthConfig()

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

	mockAuthAPI := &api.MockAuthAPI{}

	orgID := "test-org-id"
	mockUserAPI := &api.MockUserAPI{
		GetIAMUserSelfFunc: func(
			_ configuration_models.EndpointsV2,
			_ string,
			apiKey string,
		) (*api.IAMUser, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			return &api.IAMUser{OrganizationID: &orgID}, nil
		},
	}

	authService := service.NewAuthService(mockCfg, mockAuthAPI, mockUserAPI)

	authCmd := NewAuthCmd(authService)
	authCmd.PersistentFlags().String("output", "human", "Output format")
	authCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	authCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	commandOutput := new(bytes.Buffer)
	authCmd.SetOut(commandOutput)
	authCmd.SetErr(commandOutput)
	authCmd.SetArgs([]string{
		"login",
		"--profile", "test-profile",
		"--api-key", "test-api-key",
	})

	err := authCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `{
  "message": "✅ Authentication successful!\n"
}
`
	if commandOutput.String() != expected {
		t.Errorf("JSON snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

func TestAuthSubCmd_Login_Output_YAML_FromProfile(t *testing.T) {
	mockCfg := newTestAuthConfig()

	mockCfg.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{
			APIKey:         "test-api-key",
			OrganizationID: "test-org-id",
			Output:         "yaml",
		Endpoints: configuration_models.EndpointsV2{
			IAM:  "https://iam.example.com",
			Dash: "https://dash.example.com",
			CH:   "https://ch.example.com",
		},
	}, nil
	}

	mockAuthAPI := &api.MockAuthAPI{}

	orgID := "test-org-id"
	mockUserAPI := &api.MockUserAPI{
		GetIAMUserSelfFunc: func(
			_ configuration_models.EndpointsV2,
			_ string,
			apiKey string,
		) (*api.IAMUser, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			return &api.IAMUser{OrganizationID: &orgID}, nil
		},
	}

	authService := service.NewAuthService(mockCfg, mockAuthAPI, mockUserAPI)

	authCmd := NewAuthCmd(authService)
	authCmd.PersistentFlags().String("output", "human", "Output format")
	authCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	authCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	commandOutput := new(bytes.Buffer)
	authCmd.SetOut(commandOutput)
	authCmd.SetErr(commandOutput)
	authCmd.SetArgs([]string{
		"login",
		"--profile", "test-profile",
		"--api-key", "test-api-key",
	})

	err := authCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `message: |
    ✅ Authentication successful!
`
	if commandOutput.String() != expected {
		t.Errorf("YAML snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}
