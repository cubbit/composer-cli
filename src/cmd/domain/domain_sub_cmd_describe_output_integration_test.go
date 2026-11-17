package cmd_domain

import (
	"bytes"
	"testing"
	"time"

	api "github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/src/service"
)

func TestDomainSubCmd_Describe_Output_JSON(t *testing.T) {
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

	frozen := time.Date(2026, 8, 19, 10, 0, 0, 0, time.UTC)

	mockAPI := &api.MockDomainAPI{
		GetFunc: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			domainID string,
		) (*api.DomainDTO, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}
			if domainID != "dom-desc" {
				t.Fatalf("Expected domain ID dom-desc, got %q", domainID)
			}

			return &api.DomainDTO{
				ID:             "dom-desc",
				DomainName:     "describe-test.com",
				CreatedAt:      frozen,
				Challenge:      "challenge-token",
			OrganizationID: "test-org-id",
				IsShared:       false,
			}, nil
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
		"describe",
		"--domain-id", "dom-desc",
		"--output", "json",
	})

	err := domainCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `{
  "id": "dom-desc",
  "domain_name": "describe-test.com",
  "created_at": "2026-08-19T10:00:00Z",
  "challenge": "challenge-token",
  "organization_id": "test-org-id",
  "is_shared": false
}
`
	if commandOutput.String() != expected {
		t.Errorf("JSON snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

func TestDomainSubCmd_Describe_Output_JSON_FromProfile(t *testing.T) {
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

	frozen := time.Date(2026, 8, 19, 10, 0, 0, 0, time.UTC)

	mockAPI := &api.MockDomainAPI{
		GetFunc: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			domainID string,
		) (*api.DomainDTO, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}
			if domainID != "dom-desc" {
				t.Fatalf("Expected domain ID dom-desc, got %q", domainID)
			}

			return &api.DomainDTO{
				ID:             "dom-desc",
				DomainName:     "describe-test.com",
				CreatedAt:      frozen,
				Challenge:      "challenge-token",
			OrganizationID: "test-org-id",
				IsShared:       false,
			}, nil
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
		"describe",
		"--domain-id", "dom-desc",
	})

	err := domainCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `{
  "id": "dom-desc",
  "domain_name": "describe-test.com",
  "created_at": "2026-08-19T10:00:00Z",
  "challenge": "challenge-token",
  "organization_id": "test-org-id",
  "is_shared": false
}
`
	if commandOutput.String() != expected {
		t.Errorf("JSON snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

func TestDomainSubCmd_Describe_Output_YAML(t *testing.T) {
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

	frozen := time.Date(2026, 8, 19, 10, 0, 0, 0, time.UTC)

	mockAPI := &api.MockDomainAPI{
		GetFunc: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			domainID string,
		) (*api.DomainDTO, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}
			if domainID != "dom-yaml" {
				t.Fatalf("Expected domain ID dom-yaml, got %q", domainID)
			}

			return &api.DomainDTO{
				ID:             "dom-yaml",
				DomainName:     "yaml-test.com",
				CreatedAt:      frozen,
				Challenge:      "yaml-challenge",
			OrganizationID: "test-org-id",
				IsShared:       false,
			}, nil
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
		"describe",
		"--domain-id", "dom-yaml",
		"--output", "yaml",
	})

	err := domainCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `id: dom-yaml
domain_name: yaml-test.com
created_at: 2026-08-19T10:00:00Z
deleted_at: null
challenge: yaml-challenge
verified_at: null
organization_id: test-org-id
is_shared: false
`
	if commandOutput.String() != expected {
		t.Errorf("YAML snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

func TestDomainSubCmd_Describe_Output_YAML_FromProfile(t *testing.T) {
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

	frozen := time.Date(2026, 8, 19, 10, 0, 0, 0, time.UTC)

	mockAPI := &api.MockDomainAPI{
		GetFunc: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			domainID string,
		) (*api.DomainDTO, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}
			if domainID != "dom-yaml" {
				t.Fatalf("Expected domain ID dom-yaml, got %q", domainID)
			}

			return &api.DomainDTO{
				ID:             "dom-yaml",
				DomainName:     "yaml-test.com",
				CreatedAt:      frozen,
				Challenge:      "yaml-challenge",
			OrganizationID: "test-org-id",
				IsShared:       false,
			}, nil
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
		"describe",
		"--domain-id", "dom-yaml",
	})

	err := domainCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `id: dom-yaml
domain_name: yaml-test.com
created_at: 2026-08-19T10:00:00Z
deleted_at: null
challenge: yaml-challenge
verified_at: null
organization_id: test-org-id
is_shared: false
`
	if commandOutput.String() != expected {
		t.Errorf("YAML snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}
