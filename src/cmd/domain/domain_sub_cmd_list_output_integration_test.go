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

func TestDomainSubCmd_List_Output_JSON(t *testing.T) {
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
		ListFunc: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			page int,
			itemsPerPage int,
		) (*api.GenericPaginatedResponse[api.DomainDTO], error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}

			return &api.GenericPaginatedResponse[api.DomainDTO]{
				Data: []api.DomainDTO{
					{
						ID:             "dom-list-1",
						DomainName:     "first-domain.com",
						CreatedAt:      frozen,
						Challenge:      "challenge-001",
					OrganizationID: "test-org-id",
						IsShared:       false,
				},
					{
						ID:             "dom-list-2",
						DomainName:     "second-domain.com",
						CreatedAt:      frozen,
						Challenge:      "challenge-002",
					OrganizationID: "test-org-id",
						IsShared:       true,
				},
				},
				NextPage: nil,
				Count:    2,
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
		"list",
		"--output", "json",
	})

	err := domainCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `[
  {
    "id": "dom-list-1",
    "domain_name": "first-domain.com",
    "created_at": "2026-08-19T10:00:00Z",
    "challenge": "challenge-001",
    "organization_id": "test-org-id",
    "is_shared": false
  },
  {
    "id": "dom-list-2",
    "domain_name": "second-domain.com",
    "created_at": "2026-08-19T10:00:00Z",
    "challenge": "challenge-002",
    "organization_id": "test-org-id",
    "is_shared": true
  }
]
`
	if commandOutput.String() != expected {
		t.Errorf("JSON snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

func TestDomainSubCmd_List_Output_JSON_FromProfile(t *testing.T) {
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
		ListFunc: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			page int,
			itemsPerPage int,
		) (*api.GenericPaginatedResponse[api.DomainDTO], error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}

			return &api.GenericPaginatedResponse[api.DomainDTO]{
				Data: []api.DomainDTO{
					{
						ID:             "dom-list-1",
						DomainName:     "first-domain.com",
						CreatedAt:      frozen,
						Challenge:      "challenge-001",
					OrganizationID: "test-org-id",
						IsShared:       false,
				},
					{
						ID:             "dom-list-2",
						DomainName:     "second-domain.com",
						CreatedAt:      frozen,
						Challenge:      "challenge-002",
					OrganizationID: "test-org-id",
						IsShared:       true,
				},
				},
				NextPage: nil,
				Count:    2,
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
		"list",
	})

	err := domainCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `[
  {
    "id": "dom-list-1",
    "domain_name": "first-domain.com",
    "created_at": "2026-08-19T10:00:00Z",
    "challenge": "challenge-001",
    "organization_id": "test-org-id",
    "is_shared": false
  },
  {
    "id": "dom-list-2",
    "domain_name": "second-domain.com",
    "created_at": "2026-08-19T10:00:00Z",
    "challenge": "challenge-002",
    "organization_id": "test-org-id",
    "is_shared": true
  }
]
`
	if commandOutput.String() != expected {
		t.Errorf("JSON snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

func TestDomainSubCmd_List_Output_YAML(t *testing.T) {
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
		ListFunc: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			page int,
			itemsPerPage int,
		) (*api.GenericPaginatedResponse[api.DomainDTO], error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}

			return &api.GenericPaginatedResponse[api.DomainDTO]{
				Data: []api.DomainDTO{
					{
						ID:             "dom-list-y-1",
						DomainName:     "yaml-first.com",
						CreatedAt:      frozen,
						Challenge:      "yaml-ch-001",
					OrganizationID: "test-org-id",
						IsShared:       false,
				},
					{
						ID:             "dom-list-y-2",
						DomainName:     "yaml-second.com",
						CreatedAt:      frozen,
						Challenge:      "yaml-ch-002",
					OrganizationID: "test-org-id",
						IsShared:       true,
				},
				},
				NextPage: nil,
				Count:    2,
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
		"list",
		"--output", "yaml",
	})

	err := domainCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `- id: dom-list-y-1
  domain_name: yaml-first.com
  created_at: 2026-08-19T10:00:00Z
  deleted_at: null
  challenge: yaml-ch-001
  verified_at: null
  organization_id: test-org-id
  is_shared: false
- id: dom-list-y-2
  domain_name: yaml-second.com
  created_at: 2026-08-19T10:00:00Z
  deleted_at: null
  challenge: yaml-ch-002
  verified_at: null
  organization_id: test-org-id
  is_shared: true
`
	if commandOutput.String() != expected {
		t.Errorf("YAML snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

func TestDomainSubCmd_List_Output_YAML_FromProfile(t *testing.T) {
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
		ListFunc: func(
			_ configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			page int,
			itemsPerPage int,
		) (*api.GenericPaginatedResponse[api.DomainDTO], error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}

			return &api.GenericPaginatedResponse[api.DomainDTO]{
				Data: []api.DomainDTO{
					{
						ID:             "dom-list-y-1",
						DomainName:     "yaml-first.com",
						CreatedAt:      frozen,
						Challenge:      "yaml-ch-001",
					OrganizationID: "test-org-id",
						IsShared:       false,
				},
					{
						ID:             "dom-list-y-2",
						DomainName:     "yaml-second.com",
						CreatedAt:      frozen,
						Challenge:      "yaml-ch-002",
					OrganizationID: "test-org-id",
						IsShared:       true,
				},
				},
				NextPage: nil,
				Count:    2,
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
		"list",
	})

	err := domainCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `- id: dom-list-y-1
  domain_name: yaml-first.com
  created_at: 2026-08-19T10:00:00Z
  deleted_at: null
  challenge: yaml-ch-001
  verified_at: null
  organization_id: test-org-id
  is_shared: false
- id: dom-list-y-2
  domain_name: yaml-second.com
  created_at: 2026-08-19T10:00:00Z
  deleted_at: null
  challenge: yaml-ch-002
  verified_at: null
  organization_id: test-org-id
  is_shared: true
`
	if commandOutput.String() != expected {
		t.Errorf("YAML snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}
