package cmd_auth

import (
	"bytes"
	"testing"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/src/service"
)

func TestAuthSubCmd_SignUp_Output_JSON(t *testing.T) {
	mockCfg := newTestAuthConfig()

	mockAuthAPI := &api.MockAuthAPI{
		SignUpFunc: func(
			_ configuration_models.EndpointsV2,
			email string,
			username string,
			firstName *string,
			lastName *string,
			authenticationPublicKey *string,
			organizationName string,
			organizationBasePolicy map[string]interface{},
			organizationSettings map[string]interface{},
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
		"signup",
		"--email", "test@example.com",
		"--username", "testuser",
		"--organization", "TestOrg",
		"--output", "json",
	})

	err := authCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `{
  "message": "Sign up completed successfully. Please check your email to verify your account.\n"
}
`
	if commandOutput.String() != expected {
		t.Errorf("JSON snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

func TestAuthSubCmd_SignUp_Output_YAML(t *testing.T) {
	mockCfg := newTestAuthConfig()

	mockAuthAPI := &api.MockAuthAPI{
		SignUpFunc: func(
			_ configuration_models.EndpointsV2,
			email string,
			username string,
			firstName *string,
			lastName *string,
			authenticationPublicKey *string,
			organizationName string,
			organizationBasePolicy map[string]interface{},
			organizationSettings map[string]interface{},
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
		"signup",
		"--email", "test@example.com",
		"--username", "testuser",
		"--organization", "TestOrg",
		"--output", "yaml",
	})

	err := authCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `message: |
    Sign up completed successfully. Please check your email to verify your account.
`
	if commandOutput.String() != expected {
		t.Errorf("YAML snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

func TestAuthSubCmd_SignUp_Output_JSON_FromProfile(t *testing.T) {
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
		SignUpFunc: func(
			_ configuration_models.EndpointsV2,
			email string,
			username string,
			firstName *string,
			lastName *string,
			authenticationPublicKey *string,
			organizationName string,
			organizationBasePolicy map[string]interface{},
			organizationSettings map[string]interface{},
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
		"signup",
		"--email", "test@example.com",
		"--username", "testuser",
		"--organization", "TestOrg",
	})

	err := authCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `{
  "message": "Sign up completed successfully. Please check your email to verify your account.\n"
}
`
	if commandOutput.String() != expected {
		t.Errorf("JSON snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

func TestAuthSubCmd_SignUp_Output_YAML_FromProfile(t *testing.T) {
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
		SignUpFunc: func(
			_ configuration_models.EndpointsV2,
			email string,
			username string,
			firstName *string,
			lastName *string,
			authenticationPublicKey *string,
			organizationName string,
			organizationBasePolicy map[string]interface{},
			organizationSettings map[string]interface{},
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
		"signup",
		"--email", "test@example.com",
		"--username", "testuser",
		"--organization", "TestOrg",
	})

	err := authCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `message: |
    Sign up completed successfully. Please check your email to verify your account.
`
	if commandOutput.String() != expected {
		t.Errorf("YAML snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}
