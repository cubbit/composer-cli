package cmd_domain

import (
	"bytes"
	"testing"

	api "github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/src/service"
)

func TestDomainSubCmd_Verify_Output_JSON(t *testing.T) {
	mockCfg := configuration_handler.NewMockConfigurationHandler()
	mockCfg.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{
			APIKey:         "test-api-key",
			OrganizationID: "test-org-id",
			Output:         configuration_models.OutputHuman,
			Endpoints: configuration_models.EndpointsV2{
				IAM: "https://iam.example.com",
				Dash: "https://dash.example.com",
				CH:  "https://ch.example.com",
			},
		}, nil
	}

	mockAPI := &api.MockDomainAPI{
		VerifyFunc: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			domainID string,
		) (*api.DomainVerifyResult, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}
			if domainID != "dom-verify-json" {
				t.Fatalf("Expected domain ID dom-verify-json, got %q", domainID)
			}

			return &api.DomainVerifyResult{Verified: true}, nil
		},
	}

	svc := service.NewDomainService(mockCfg, mockAPI, &api.UserAPI{})

	domainCmd := NewDomainCmd(svc)
	domainCmd.PersistentFlags().String("profile", "", "Profile")
	domainCmd.PersistentFlags().String("output", "human", "Output format")
	domainCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	domainCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	commandOutput := new(bytes.Buffer)
	domainCmd.SetOut(commandOutput)
	domainCmd.SetErr(commandOutput)
	domainCmd.SetArgs([]string{
		"verify",
		"--domain-id", "dom-verify-json",
		"--output", "json",
	})

	err := domainCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `{
  "verified": true
}
`
	if commandOutput.String() != expected {
		t.Errorf("JSON snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

func TestDomainSubCmd_Verify_Output_YAML(t *testing.T) {
	mockCfg := configuration_handler.NewMockConfigurationHandler()
	mockCfg.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{
			APIKey:         "test-api-key",
			OrganizationID: "test-org-id",
			Output:         configuration_models.OutputHuman,
			Endpoints: configuration_models.EndpointsV2{
				IAM: "https://iam.example.com",
				Dash: "https://dash.example.com",
				CH:  "https://ch.example.com",
			},
		}, nil
	}

	mockAPI := &api.MockDomainAPI{
		VerifyFunc: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			domainID string,
		) (*api.DomainVerifyResult, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}
			if domainID != "dom-verify-yaml" {
				t.Fatalf("Expected domain ID dom-verify-yaml, got %q", domainID)
			}

			return &api.DomainVerifyResult{Verified: false}, nil
		},
	}

	svc := service.NewDomainService(mockCfg, mockAPI, &api.UserAPI{})

	domainCmd := NewDomainCmd(svc)
	domainCmd.PersistentFlags().String("profile", "", "Profile")
	domainCmd.PersistentFlags().String("output", "human", "Output format")
	domainCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	domainCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	commandOutput := new(bytes.Buffer)
	domainCmd.SetOut(commandOutput)
	domainCmd.SetErr(commandOutput)
	domainCmd.SetArgs([]string{
		"verify",
		"--domain-id", "dom-verify-yaml",
		"--output", "yaml",
	})

	err := domainCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `verified: false
reasons: null
`
	if commandOutput.String() != expected {
		t.Errorf("YAML snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

func TestDomainSubCmd_Verify_Output_JSON_FromProfile(t *testing.T) {
	mockCfg := configuration_handler.NewMockConfigurationHandler()
	mockCfg.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{
			APIKey:         "test-api-key",
			OrganizationID: "test-org-id",
			Output:         "json",
			Endpoints: configuration_models.EndpointsV2{
				IAM: "https://iam.example.com",
				Dash: "https://dash.example.com",
				CH:  "https://ch.example.com",
			},
		}, nil
	}

	mockAPI := &api.MockDomainAPI{
		VerifyFunc: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			domainID string,
		) (*api.DomainVerifyResult, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}
			if domainID != "dom-verify-json" {
				t.Fatalf("Expected domain ID dom-verify-json, got %q", domainID)
			}

			return &api.DomainVerifyResult{Verified: true}, nil
		},
	}

	svc := service.NewDomainService(mockCfg, mockAPI, &api.UserAPI{})

	domainCmd := NewDomainCmd(svc)
	domainCmd.PersistentFlags().String("profile", "", "Profile")
	domainCmd.PersistentFlags().String("output", "human", "Output format")
	domainCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	domainCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	commandOutput := new(bytes.Buffer)
	domainCmd.SetOut(commandOutput)
	domainCmd.SetErr(commandOutput)
	domainCmd.SetArgs([]string{
		"verify",
		"--domain-id", "dom-verify-json",
	})

	err := domainCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `{
  "verified": true
}
`
	if commandOutput.String() != expected {
		t.Errorf("JSON snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

func TestDomainSubCmd_Verify_Output_YAML_FromProfile(t *testing.T) {
	mockCfg := configuration_handler.NewMockConfigurationHandler()
	mockCfg.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{
			APIKey:         "test-api-key",
			OrganizationID: "test-org-id",
			Output:         "yaml",
			Endpoints: configuration_models.EndpointsV2{
				IAM: "https://iam.example.com",
				Dash: "https://dash.example.com",
				CH:  "https://ch.example.com",
			},
		}, nil
	}

	mockAPI := &api.MockDomainAPI{
		VerifyFunc: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			domainID string,
		) (*api.DomainVerifyResult, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}
			if domainID != "dom-verify-yaml" {
				t.Fatalf("Expected domain ID dom-verify-yaml, got %q", domainID)
			}

			return &api.DomainVerifyResult{Verified: false}, nil
		},
	}

	svc := service.NewDomainService(mockCfg, mockAPI, &api.UserAPI{})

	domainCmd := NewDomainCmd(svc)
	domainCmd.PersistentFlags().String("profile", "", "Profile")
	domainCmd.PersistentFlags().String("output", "human", "Output format")
	domainCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	domainCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	commandOutput := new(bytes.Buffer)
	domainCmd.SetOut(commandOutput)
	domainCmd.SetErr(commandOutput)
	domainCmd.SetArgs([]string{
		"verify",
		"--domain-id", "dom-verify-yaml",
	})

	err := domainCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `verified: false
reasons: null
`
	if commandOutput.String() != expected {
		t.Errorf("YAML snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

