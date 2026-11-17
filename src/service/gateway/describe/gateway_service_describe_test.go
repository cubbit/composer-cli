package describe

import (
	"bytes"
	"strings"
	"testing"

	api "github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/spf13/cobra"
)

func setupGatewayDescribeTestCommand() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.SetOut(new(bytes.Buffer))
	cmd.SetErr(new(bytes.Buffer))
	cmd.Flags().String("profile", "", "Profile")
	cmd.Flags().String("output", "human", "Output format")
	cmd.Flags().Bool("no-headers", false, "No headers")
	cmd.Flags().Bool("quiet", false, "Quiet mode")
	return cmd
}

func gatewayDescribeTestProfile() configuration_models.ProfileV2 {
	return configuration_models.ProfileV2{
		Output:         configuration_models.OutputHuman,
		APIKey:         "test-api-key",
		OrganizationID: "test-org-id",
	}
}

func TestDescribe_WithPositionalID_Human(t *testing.T) {
	mockGatewayAPI := &api.MockGatewayAPI{
		GetGatewayV5Func: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, gatewayID string) (*api.GatewayV5GetResponse, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}
			if gatewayID != "gateway-123" {
				t.Fatalf("Expected gateway ID gateway-123, got %q", gatewayID)
			}

			return &api.GatewayV5GetResponse{
				ID:     "gateway-123",
				Name:   "test-gateway",
				Slug:   "test-gateway",
				Type:   "singlecluster",
				Status: "ready",
				RedundancyClasses: []api.GatewayV5GetRedundancyClass{
					{
						ID:   "rc-123",
						Name: "standard",
					},
				},
			}, nil
		},
	}

	cmd := setupGatewayDescribeTestCommand()
	cmd.SetOut(new(bytes.Buffer))
	cmd.Flags().String("gateway-id", "", "Gateway ID")
	cmd.Flags().String("gateway-name", "", "Gateway name")

	profile := gatewayDescribeTestProfile()
	err := Describe(Dependencies{GatewayAPI: mockGatewayAPI}, cmd, nil, profile, []string{"gateway-123"})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	output := cmd.OutOrStdout().(*bytes.Buffer).String()
	if !strings.Contains(output, "Gateway: test-gateway") || !strings.Contains(output, "Metadata:") || !strings.Contains(output, "ID: gateway-123") {
		t.Fatalf("Expected human output to contain gateway details, got %q", output)
	}
}

func TestDescribe_WithGatewayName_JSON(t *testing.T) {
	mockGatewayAPI := &api.MockGatewayAPI{
		ListGatewaysV5Func: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, opts ...api.ListGatewaysV5Option) (*api.GenericPaginatedResponse[api.GatewayV5ListItemResponse], error) {
			options := &api.ListGatewaysV5Options{}
			for _, opt := range opts {
				opt(options)
			}

			if options.Page != 1 || options.Items != 1000 {
				t.Fatalf("Expected default pagination to resolve gateway by name, got page=%d items=%d", options.Page, options.Items)
			}

			return &api.GenericPaginatedResponse[api.GatewayV5ListItemResponse]{
				Data: []api.GatewayV5ListItemResponse{
					{
						ID:   "gateway-456",
						Name: "named-gateway",
					},
				},
			}, nil
		},
		GetGatewayV5Func: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, gatewayID string) (*api.GatewayV5GetResponse, error) {
			if gatewayID != "gateway-456" {
				t.Fatalf("Expected resolved gateway ID gateway-456, got %q", gatewayID)
			}

			return &api.GatewayV5GetResponse{
				ID:     "gateway-456",
				Name:   "named-gateway",
				Slug:   "named-gateway",
				Type:   "manual",
				Status: "ready",
			}, nil
		},
	}

	cmd := setupGatewayDescribeTestCommand()
	cmd.SetOut(new(bytes.Buffer))
	cmd.Flags().String("gateway-id", "", "Gateway ID")
	cmd.Flags().String("gateway-name", "", "Gateway name")
	cmd.Flags().Set("gateway-name", "named-gateway")
	cmd.Flags().Set("output", "json")

	profile := gatewayDescribeTestProfile()
	err := Describe(Dependencies{GatewayAPI: mockGatewayAPI}, cmd, nil, profile, nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	output := cmd.OutOrStdout().(*bytes.Buffer).String()
	if !strings.Contains(output, `"id": "gateway-456"`) {
		t.Fatalf("Expected json output to contain resolved gateway, got %q", output)
	}
}

func TestDescribe_WithGatewayName_PaginatesUntilFound(t *testing.T) {
	callCount := 0
	mockGatewayAPI := &api.MockGatewayAPI{
		ListGatewaysV5Func: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, opts ...api.ListGatewaysV5Option) (*api.GenericPaginatedResponse[api.GatewayV5ListItemResponse], error) {
			callCount++
			options := &api.ListGatewaysV5Options{}
			for _, opt := range opts {
				opt(options)
			}

			if options.Items != 1000 {
				t.Fatalf("Expected items per page 1000, got %d", options.Items)
			}

			switch options.Page {
			case 1:
				nextPage := 2
				return &api.GenericPaginatedResponse[api.GatewayV5ListItemResponse]{
					Data: []api.GatewayV5ListItemResponse{
						{
							ID:   "gateway-001",
							Name: "first-page-gateway",
						},
					},
					NextPage: &nextPage,
				}, nil
			case 2:
				return &api.GenericPaginatedResponse[api.GatewayV5ListItemResponse]{
					Data: []api.GatewayV5ListItemResponse{
						{
							ID:   "gateway-789",
							Name: "paged-gateway",
						},
					},
				}, nil
			default:
				t.Fatalf("Unexpected page requested: %d", options.Page)
				return nil, nil
			}
		},
		GetGatewayV5Func: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, gatewayID string) (*api.GatewayV5GetResponse, error) {
			if gatewayID != "gateway-789" {
				t.Fatalf("Expected resolved gateway ID gateway-789, got %q", gatewayID)
			}

			return &api.GatewayV5GetResponse{
				ID:     "gateway-789",
				Name:   "paged-gateway",
				Slug:   "paged-gateway",
				Type:   "manual",
				Status: "ready",
			}, nil
		},
	}

	cmd := setupGatewayDescribeTestCommand()
	cmd.SetOut(new(bytes.Buffer))
	cmd.Flags().String("gateway-id", "", "Gateway ID")
	cmd.Flags().String("gateway-name", "", "Gateway name")
	cmd.Flags().Set("gateway-name", "paged-gateway")
	cmd.Flags().Set("output", "json")

	profile := gatewayDescribeTestProfile()
	err := Describe(Dependencies{GatewayAPI: mockGatewayAPI}, cmd, nil, profile, nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if callCount != 2 {
		t.Fatalf("Expected two list calls to page through results, got %d", callCount)
	}
}

func TestDescribe_WithUnknownGatewayName(t *testing.T) {
	mockGatewayAPI := &api.MockGatewayAPI{
		ListGatewaysV5Func: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, opts ...api.ListGatewaysV5Option) (*api.GenericPaginatedResponse[api.GatewayV5ListItemResponse], error) {
			return &api.GenericPaginatedResponse[api.GatewayV5ListItemResponse]{
				Data: []api.GatewayV5ListItemResponse{},
			}, nil
		},
	}

	cmd := setupGatewayDescribeTestCommand()
	cmd.Flags().String("gateway-id", "", "Gateway ID")
	cmd.Flags().String("gateway-name", "", "Gateway name")
	cmd.Flags().Set("gateway-name", "missing-gateway")

	profile := gatewayDescribeTestProfile()
	err := Describe(Dependencies{GatewayAPI: mockGatewayAPI}, cmd, nil, profile, nil)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	if !strings.Contains(err.Error(), "gateway with name 'missing-gateway' not found") {
		t.Fatalf("Expected not found error, got %v", err)
	}
}

func TestDescribe_WithMultipleIdentifiers(t *testing.T) {
	cmd := setupGatewayDescribeTestCommand()
	cmd.Flags().String("gateway-id", "", "Gateway ID")
	cmd.Flags().String("gateway-name", "", "Gateway name")
	cmd.Flags().Set("gateway-id", "gateway-123")

	profile := gatewayDescribeTestProfile()
	err := Describe(Dependencies{GatewayAPI: &api.MockGatewayAPI{}}, cmd, nil, profile, []string{"gateway-456"})
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	if !strings.Contains(err.Error(), "specify exactly one of GATEWAY_ID, --gateway-id or --gateway-name") {
		t.Fatalf("Expected multiple identifiers error, got %v", err)
	}
}
