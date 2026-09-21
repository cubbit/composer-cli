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

func TestAuthSubCmd_Login_InlineWithTFACode(t *testing.T) {
	mockCfg := newTestAuthConfig()

	var signInTFACode string
	var createdProfileAPIKey string

	mockCfg.CreateProfileFunc = func(profileName string, apiKey string, organizationID string) error {
		createdProfileAPIKey = apiKey
		return nil
	}

	mockAuthAPI := &api.MockAuthAPI{
		SignInFunc: func(
			_ configuration_models.EndpointsV2,
			username string,
			organization string,
			password string,
			tfaCode string,
		) (*api.SignInToken, error) {
			signInTFACode = tfaCode
			return &api.SignInToken{
				AccessToken:  "test-access-token",
				RefreshToken: "test-refresh-token",
			}, nil
		},
		ForgeTokenFunc: func(
			_ configuration_models.EndpointsV2,
			operatorID string,
			username string,
			organizationName string,
			password string,
			tfaCode string,
			tokenType string,
			token string,
			refreshToken string,
		) (string, error) {
			return "test-forge-token", nil
		},
		CreateApiKeyFunc: func(
			_ configuration_models.EndpointsV2,
			operatorID string,
			name string,
			token string,
			forgeApiKeyToken string,
		) (string, error) {
			return "test-api-key", nil
		},
	}

	orgID := "test-org-id"
	mockUserAPI := &api.MockUserAPI{
		GetIAMUserSelfFunc: func(
			_ configuration_models.EndpointsV2,
			accessToken string,
			apiKey string,
		) (*api.IAMUser, error) {
			return &api.IAMUser{
				ID:               "test-operator-id",
				OrganizationID:   &orgID,
				TwoFactorEnabled: false,
			}, nil
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
		"--username", "test-username",
		"--organization", "test-organization",
		"--password", "test-password",
		"--tfa-code", "123456",
	})

	err := authCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if signInTFACode != "123456" {
		t.Fatalf("Expected tfa-code to be propagated to sign in, got %q", signInTFACode)
	}
	if createdProfileAPIKey != "test-api-key" {
		t.Fatalf("Expected created profile API key to be test-api-key, got %q", createdProfileAPIKey)
	}
}
