package service

import (
	"bytes"
	"strings"
	"testing"
	"time"

	api "github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration"
)

func TestSwarmService_Describe_WithPositionalID_Human(t *testing.T) {
	mockCfg := configuration.NewMockConfig(configuration.ProfileTypeComposer, "test-api-key", "test-org-id")

	createdAt := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)
	description := "Test swarm"
	mockSwarmAPI := &api.MockSwarmAPI{
		GetSwarmV5Func: func(urlConfig configuration.URLs, apiKey string, organizationID string, swarmID string) (*api.SwarmV5Presentation, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}
			if swarmID != "swarm-123" {
				t.Fatalf("Expected swarm ID swarm-123, got %q", swarmID)
			}

			return &api.SwarmV5Presentation{
				SwarmV5: api.SwarmV5{
					ListSwarmV5Item: api.ListSwarmV5Item{
						ID:                   "swarm-123",
						Name:                 "test-swarm",
						TotalStorageBytes:    1024,
						UsedStorageBytes:     512,
						CreatedAt:            createdAt,
						NexusCount:           2,
						RedundancyClassCount: 1,
					},
					OrganizationID: "test-org-id",
					OwnerID:        "owner-123",
					Description:    &description,
					Configuration: map[string]interface{}{
						"tier": "hot",
					},
					CreationStatus: "created",
				},
			}, nil
		},
	}

	service := NewSwarmService(mockCfg, mockSwarmAPI, &api.MockLocationAPI{}, &api.ProcessAPI{}, NewRedundancyClassValidator())
	cmd := setupTestCommand()
	cmd.SetOut(new(bytes.Buffer))
	cmd.Flags().String("swarm-id", "", "Swarm ID")
	cmd.Flags().String("swarm-name", "", "Swarm name")

	err := service.Describe(cmd, []string{"swarm-123"})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	output := cmd.OutOrStdout().(*bytes.Buffer).String()
	if !strings.Contains(output, "Swarm: test-swarm") || !strings.Contains(output, "Metadata:") || !strings.Contains(output, "ID: swarm-123") {
		t.Fatalf("Expected human output to contain swarm details, got %q", output)
	}
}

func TestSwarmService_Describe_WithSwarmName_JSON(t *testing.T) {
	mockCfg := configuration.NewMockConfig(configuration.ProfileTypeComposer, "test-api-key", "test-org-id")

	createdAt := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)
	mockSwarmAPI := &api.MockSwarmAPI{
		ListSwarmsV5Func: func(urlConfig configuration.URLs, apiKey string, organizationID string, page int, items int) (*api.GenericPaginatedResponse[api.ListSwarmV5ItemPresentation], error) {
			if page != 1 || items != 1000 {
				t.Fatalf("Expected default pagination to resolve swarm by name, got page=%d items=%d", page, items)
			}

			return &api.GenericPaginatedResponse[api.ListSwarmV5ItemPresentation]{
				Data: []api.ListSwarmV5ItemPresentation{
					{
						ListSwarmV5Item: api.ListSwarmV5Item{
							ID:   "swarm-456",
							Name: "named-swarm",
						},
					},
				},
			}, nil
		},
		GetSwarmV5Func: func(urlConfig configuration.URLs, apiKey string, organizationID string, swarmID string) (*api.SwarmV5Presentation, error) {
			if swarmID != "swarm-456" {
				t.Fatalf("Expected resolved swarm ID swarm-456, got %q", swarmID)
			}

			return &api.SwarmV5Presentation{
				SwarmV5: api.SwarmV5{
					ListSwarmV5Item: api.ListSwarmV5Item{
						ID:                   "swarm-456",
						Name:                 "named-swarm",
						TotalStorageBytes:    2048,
						UsedStorageBytes:     1024,
						CreatedAt:            createdAt,
						NexusCount:           1,
						RedundancyClassCount: 1,
					},
					OrganizationID: "test-org-id",
					OwnerID:        "owner-456",
					Configuration:  map[string]interface{}{},
				},
			}, nil
		},
	}

	service := NewSwarmService(mockCfg, mockSwarmAPI, &api.MockLocationAPI{}, &api.ProcessAPI{}, NewRedundancyClassValidator())
	cmd := setupTestCommand()
	cmd.SetOut(new(bytes.Buffer))
	cmd.Flags().String("swarm-id", "", "Swarm ID")
	cmd.Flags().String("swarm-name", "", "Swarm name")
	cmd.Flags().Set("swarm-name", "named-swarm")
	cmd.Flags().Set("output", "json")

	err := service.Describe(cmd, nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	output := cmd.OutOrStdout().(*bytes.Buffer).String()
	if !strings.Contains(output, `"id": "swarm-456"`) {
		t.Fatalf("Expected json output to contain resolved swarm, got %q", output)
	}
}

func TestSwarmService_Describe_WithSwarmName_PaginatesUntilFound(t *testing.T) {
	mockCfg := configuration.NewMockConfig(configuration.ProfileTypeComposer, "test-api-key", "test-org-id")

	callCount := 0
	mockSwarmAPI := &api.MockSwarmAPI{
		ListSwarmsV5Func: func(urlConfig configuration.URLs, apiKey string, organizationID string, page int, items int) (*api.GenericPaginatedResponse[api.ListSwarmV5ItemPresentation], error) {
			callCount++
			if items != 1000 {
				t.Fatalf("Expected items per page 1000, got %d", items)
			}

			switch page {
			case 1:
				return &api.GenericPaginatedResponse[api.ListSwarmV5ItemPresentation]{
					Data: []api.ListSwarmV5ItemPresentation{
						{
							ListSwarmV5Item: api.ListSwarmV5Item{
								ID:   "swarm-001",
								Name: "first-page-swarm",
							},
						},
					},
					NextPage: func() *int { next := 2; return &next }(),
				}, nil
			case 2:
				return &api.GenericPaginatedResponse[api.ListSwarmV5ItemPresentation]{
					Data: []api.ListSwarmV5ItemPresentation{
						{
							ListSwarmV5Item: api.ListSwarmV5Item{
								ID:   "swarm-789",
								Name: "paged-swarm",
							},
						},
					},
				}, nil
			default:
				t.Fatalf("Unexpected page requested: %d", page)
				return nil, nil
			}
		},
		GetSwarmV5Func: func(urlConfig configuration.URLs, apiKey string, organizationID string, swarmID string) (*api.SwarmV5Presentation, error) {
			if swarmID != "swarm-789" {
				t.Fatalf("Expected resolved swarm ID swarm-789, got %q", swarmID)
			}

			return &api.SwarmV5Presentation{
				SwarmV5: api.SwarmV5{
					ListSwarmV5Item: api.ListSwarmV5Item{
						ID:   "swarm-789",
						Name: "paged-swarm",
					},
					OrganizationID: "test-org-id",
					OwnerID:        "owner-789",
					Configuration:  map[string]interface{}{},
				},
			}, nil
		},
	}

	service := NewSwarmService(mockCfg, mockSwarmAPI, &api.MockLocationAPI{}, &api.ProcessAPI{}, NewRedundancyClassValidator())
	cmd := setupTestCommand()
	cmd.SetOut(new(bytes.Buffer))
	cmd.Flags().String("swarm-id", "", "Swarm ID")
	cmd.Flags().String("swarm-name", "", "Swarm name")
	cmd.Flags().Set("swarm-name", "paged-swarm")
	cmd.Flags().Set("output", "json")

	err := service.Describe(cmd, nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if callCount != 2 {
		t.Fatalf("Expected two list calls to page through results, got %d", callCount)
	}
}

func TestSwarmService_Describe_WithUnknownSwarmName(t *testing.T) {
	mockCfg := configuration.NewMockConfig(configuration.ProfileTypeComposer, "test-api-key", "test-org-id")

	mockSwarmAPI := &api.MockSwarmAPI{
		ListSwarmsV5Func: func(urlConfig configuration.URLs, apiKey string, organizationID string, page int, items int) (*api.GenericPaginatedResponse[api.ListSwarmV5ItemPresentation], error) {
			return &api.GenericPaginatedResponse[api.ListSwarmV5ItemPresentation]{
				Data: []api.ListSwarmV5ItemPresentation{},
			}, nil
		},
	}

	service := NewSwarmService(mockCfg, mockSwarmAPI, &api.MockLocationAPI{}, &api.ProcessAPI{}, NewRedundancyClassValidator())
	cmd := setupTestCommand()
	cmd.Flags().String("swarm-id", "", "Swarm ID")
	cmd.Flags().String("swarm-name", "", "Swarm name")
	cmd.Flags().Set("swarm-name", "missing-swarm")

	err := service.Describe(cmd, nil)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	if !strings.Contains(err.Error(), "swarm with name 'missing-swarm' not found") {
		t.Fatalf("Expected not found error, got %v", err)
	}
}

func TestSwarmService_InterfaceCompliance(t *testing.T) {
	var _ SwarmServiceInterface = SwarmService{}
}
