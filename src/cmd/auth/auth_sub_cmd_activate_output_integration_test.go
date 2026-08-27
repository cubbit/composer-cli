package cmd_auth

import (
	"bytes"
	"testing"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/src/service"
)

func newTestAuthConfig() *configuration_handler.MockConfigurationHandler {
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
	return mockCfg
}

func TestAuthSubCmd_Activate_Output_JSON(t *testing.T) {
	mockCfg := newTestAuthConfig()

	mockAuthAPI := &api.MockAuthAPI{
		ActivateFunc: func(
			_ configuration_models.EndpointsV2,
			token string,
		) error {
			return nil
		},
	}

	mockUserAPI := &api.MockUserAPI{}

	authService := service.NewAuthService(mockCfg, mockAuthAPI, mockUserAPI)

	authCmd := NewAuthCmd(authService)
	authCmd.PersistentFlags().String("profile", "", "Profile")
	authCmd.PersistentFlags().String("output", "human", "Output format")
	authCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	authCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	commandOutput := new(bytes.Buffer)
	authCmd.SetOut(commandOutput)
	authCmd.SetErr(commandOutput)
	authCmd.SetArgs([]string{
		"activate",
		"--token", "test-token-123",
		"--output", "json",
	})

	err := authCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `{
  "message": "Activation completed successfully. You can now log in.\n"
}
`
	if commandOutput.String() != expected {
		t.Errorf("JSON snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

func TestAuthSubCmd_Activate_Output_YAML(t *testing.T) {
	mockCfg := newTestAuthConfig()

	mockAuthAPI := &api.MockAuthAPI{
		ActivateFunc: func(
			_ configuration_models.EndpointsV2,
			token string,
		) error {
			return nil
		},
	}

	mockUserAPI := &api.MockUserAPI{}

	authService := service.NewAuthService(mockCfg, mockAuthAPI, mockUserAPI)

	authCmd := NewAuthCmd(authService)
	authCmd.PersistentFlags().String("profile", "", "Profile")
	authCmd.PersistentFlags().String("output", "human", "Output format")
	authCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	authCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	commandOutput := new(bytes.Buffer)
	authCmd.SetOut(commandOutput)
	authCmd.SetErr(commandOutput)
	authCmd.SetArgs([]string{
		"activate",
		"--token", "test-token-123",
		"--output", "yaml",
	})

	err := authCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `message: |
    Activation completed successfully. You can now log in.
`
	if commandOutput.String() != expected {
		t.Errorf("YAML snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

func TestAuthSubCmd_Activate_Output_JSON_FromProfile(t *testing.T) {
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

	mockAuthAPI := &api.MockAuthAPI{
		ActivateFunc: func(
			_ configuration_models.EndpointsV2,
			token string,
		) error {
			return nil
		},
	}

	mockUserAPI := &api.MockUserAPI{}

	authService := service.NewAuthService(mockCfg, mockAuthAPI, mockUserAPI)

	authCmd := NewAuthCmd(authService)
	authCmd.PersistentFlags().String("profile", "", "Profile")
	authCmd.PersistentFlags().String("output", "human", "Output format")
	authCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	authCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	commandOutput := new(bytes.Buffer)
	authCmd.SetOut(commandOutput)
	authCmd.SetErr(commandOutput)
	authCmd.SetArgs([]string{
		"activate",
		"--token", "test-token-123",
	})

	err := authCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `{
  "message": "Activation completed successfully. You can now log in.\n"
}
`
	if commandOutput.String() != expected {
		t.Errorf("JSON snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

func TestAuthSubCmd_Activate_Output_YAML_FromProfile(t *testing.T) {
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

	mockAuthAPI := &api.MockAuthAPI{
		ActivateFunc: func(
			_ configuration_models.EndpointsV2,
			token string,
		) error {
			return nil
		},
	}

	mockUserAPI := &api.MockUserAPI{}

	authService := service.NewAuthService(mockCfg, mockAuthAPI, mockUserAPI)

	authCmd := NewAuthCmd(authService)
	authCmd.PersistentFlags().String("profile", "", "Profile")
	authCmd.PersistentFlags().String("output", "human", "Output format")
	authCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	authCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	commandOutput := new(bytes.Buffer)
	authCmd.SetOut(commandOutput)
	authCmd.SetErr(commandOutput)
	authCmd.SetArgs([]string{
		"activate",
		"--token", "test-token-123",
	})

	err := authCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `message: |
    Activation completed successfully. You can now log in.
`
	if commandOutput.String() != expected {
		t.Errorf("YAML snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}
