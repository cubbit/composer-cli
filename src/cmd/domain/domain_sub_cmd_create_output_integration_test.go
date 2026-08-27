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

func TestDomainSubCmd_Create_Output_JSON(t *testing.T) {
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
		CreateFunc: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			request *api.CreateDomainRequestBody,
		) (*api.DomainDTO, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}
			if request.DomainName != "created-test.com" {
				t.Fatalf("Expected domain name 'created-test.com', got %q", request.DomainName)
			}

			return &api.DomainDTO{
				ID:             "dom-created",
				DomainName:     "created-test.com",
				CreatedAt:      frozen,
				Challenge:      "create-challenge",
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
		"create",
		"--domain-name", "created-test.com",
		"--output", "json",
	})

	err := domainCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `{
  "id": "dom-created",
  "domain_name": "created-test.com",
  "created_at": "2026-08-19T10:00:00Z",
  "challenge": "create-challenge",
  "organization_id": "test-org-id",
  "is_shared": false
}
`
	if commandOutput.String() != expected {
		t.Errorf("JSON snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

func TestDomainSubCmd_Create_Output_YAML(t *testing.T) {
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
		CreateFunc: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			request *api.CreateDomainRequestBody,
		) (*api.DomainDTO, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}
			if request.DomainName != "yaml-created.com" {
				t.Fatalf("Expected domain name 'yaml-created.com', got %q", request.DomainName)
			}

			return &api.DomainDTO{
				ID:             "dom-yaml-created",
				DomainName:     "yaml-created.com",
				CreatedAt:      frozen,
				Challenge:      "yaml-create-ch",
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
		"create",
		"--domain-name", "yaml-created.com",
		"--output", "yaml",
	})

	err := domainCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `id: dom-yaml-created
domain_name: yaml-created.com
created_at: 2026-08-19T10:00:00Z
deleted_at: null
challenge: yaml-create-ch
verified_at: null
organization_id: test-org-id
is_shared: false
`
	if commandOutput.String() != expected {
		t.Errorf("YAML snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

func TestDomainSubCmd_Create_Output_JSON_FromProfile(t *testing.T) {
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
		CreateFunc: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			request *api.CreateDomainRequestBody,
		) (*api.DomainDTO, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}
			if request.DomainName != "created-test.com" {
				t.Fatalf("Expected domain name 'created-test.com', got %q", request.DomainName)
			}

			return &api.DomainDTO{
				ID:             "dom-created",
				DomainName:     "created-test.com",
				CreatedAt:      frozen,
				Challenge:      "create-challenge",
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
		"create",
		"--domain-name", "created-test.com",
	})

	err := domainCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `{
  "id": "dom-created",
  "domain_name": "created-test.com",
  "created_at": "2026-08-19T10:00:00Z",
  "challenge": "create-challenge",
  "organization_id": "test-org-id",
  "is_shared": false
}
`
	if commandOutput.String() != expected {
		t.Errorf("JSON snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

func TestDomainSubCmd_Create_Output_YAML_FromProfile(t *testing.T) {
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
		CreateFunc: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			request *api.CreateDomainRequestBody,
		) (*api.DomainDTO, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}
			if request.DomainName != "yaml-created.com" {
				t.Fatalf("Expected domain name 'yaml-created.com', got %q", request.DomainName)
			}

			return &api.DomainDTO{
				ID:             "dom-yaml-created",
				DomainName:     "yaml-created.com",
				CreatedAt:      frozen,
				Challenge:      "yaml-create-ch",
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
		"create",
		"--domain-name", "yaml-created.com",
	})

	err := domainCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `id: dom-yaml-created
domain_name: yaml-created.com
created_at: 2026-08-19T10:00:00Z
deleted_at: null
challenge: yaml-create-ch
verified_at: null
organization_id: test-org-id
is_shared: false
`
	if commandOutput.String() != expected {
		t.Errorf("YAML snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

