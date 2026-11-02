package update

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	api "github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/spf13/cobra"
)

func setupGatewayUpdateTestCommand() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.SetOut(new(bytes.Buffer))
	cmd.SetErr(new(bytes.Buffer))
	cmd.Flags().String("output", "human", "Output format")
	cmd.Flags().Bool("quiet", false, "Quiet mode")
	cmd.Flags().Bool("no-headers", false, "No headers")
	cmd.Flags().String("name", "", "Gateway name")
	cmd.Flags().String("description", "", "Gateway description")
	cmd.Flags().String("location", "", "Gateway location")
	return cmd
}

func gatewayUpdateTestProfile() configuration_models.ProfileV2 {
	return configuration_models.ProfileV2{
		Output:         configuration_models.OutputHuman,
		APIKey:         "test-api-key",
		OrganizationID: "test-org-id",
	}
}

func TestUpdate_UpdatesOnlyChangedFields(t *testing.T) {
	description := ""

	mockGatewayAPI := &api.MockGatewayAPI{
		UpdateGatewayV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			request *api.UpdateGatewayV5Request,
		) (*api.GatewayV5GetResponse, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}
			if gatewayID != "gateway-123" {
				t.Fatalf("Expected gateway ID gateway-123, got %q", gatewayID)
			}
			if request.Name != nil {
				t.Fatalf("Expected name to be omitted, got %q", *request.Name)
			}
			if request.Description == nil || *request.Description != description {
				t.Fatalf("Expected an explicit empty description, got %+v", request.Description)
			}
			if request.Location != nil {
				t.Fatalf("Expected location to be omitted, got %q", *request.Location)
			}

			return &api.GatewayV5GetResponse{
				ID:          "gateway-123",
				Name:        "test-gateway",
				Slug:        "test-gateway",
				Type:        api.IngressTypeManual,
				Status:      api.GatewayV5StatusReady,
				Description: request.Description,
			}, nil
		},
	}

	cmd := setupGatewayUpdateTestCommand()
	cmd.Flags().Set("description", description)

	err := Update(Dependencies{GatewayAPI: mockGatewayAPI}, cmd, nil, gatewayUpdateTestProfile(), []string{"gateway-123"})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	output := cmd.OutOrStdout().(*bytes.Buffer).String()
	if !strings.Contains(output, "Gateway test-gateway updated successfully.") {
		t.Fatalf("Expected success message, got %q", output)
	}
	if !strings.Contains(output, "Gateway: test-gateway") {
		t.Fatalf("Expected gateway details, got %q", output)
	}
}

func TestUpdate_RequiresAtLeastOneField(t *testing.T) {
	mockGatewayAPI := &api.MockGatewayAPI{}

	cmd := setupGatewayUpdateTestCommand()

	err := Update(Dependencies{GatewayAPI: mockGatewayAPI}, cmd, nil, gatewayUpdateTestProfile(), []string{"gateway-123"})
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	expected := "specify at least one of --name, --description or --location"
	if err.Error() != expected {
		t.Fatalf("Expected error %q, got %q", expected, err.Error())
	}
}

func TestUpdate_ValidatesFieldsBeforeCallingAPI(t *testing.T) {
	mockGatewayAPI := &api.MockGatewayAPI{}

	cmd := setupGatewayUpdateTestCommand()

	err := Update(Dependencies{GatewayAPI: mockGatewayAPI}, cmd, nil, gatewayUpdateTestProfile(), []string{"gateway-123"})
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	expected := "specify at least one of --name, --description or --location"
	if err.Error() != expected {
		t.Fatalf("Expected validation error before the API call, got %q", err.Error())
	}
}

func TestUpdate_JSONOutput(t *testing.T) {
	name := "renamed-gateway"
	location := "eu-west-1"

	mockGatewayAPI := &api.MockGatewayAPI{
		UpdateGatewayV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			request *api.UpdateGatewayV5Request,
		) (*api.GatewayV5GetResponse, error) {
			if request.Name == nil || *request.Name != name {
				t.Fatalf("Expected name %q, got %+v", name, request.Name)
			}
			if request.Location == nil || *request.Location != location {
				t.Fatalf("Expected location %q, got %+v", location, request.Location)
			}

			return &api.GatewayV5GetResponse{
				ID:       "gateway-123",
				Name:     name,
				Slug:     "renamed-gateway",
				Type:     api.IngressTypeManual,
				Status:   api.GatewayV5StatusReady,
				Location: location,
			}, nil
		},
	}

	cmd := setupGatewayUpdateTestCommand()
	cmd.Flags().Set("name", name)
	cmd.Flags().Set("location", location)
	cmd.Flags().Set("output", "json")

	err := Update(Dependencies{GatewayAPI: mockGatewayAPI}, cmd, nil, gatewayUpdateTestProfile(), []string{"gateway-123"})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	output := cmd.OutOrStdout().(*bytes.Buffer).String()
	if !strings.Contains(output, `"id": "gateway-123"`) ||
		!strings.Contains(output, `"name": "renamed-gateway"`) ||
		!strings.Contains(output, `"location": "eu-west-1"`) {
		t.Fatalf("Expected JSON gateway payload, got %q", output)
	}
}

func TestUpdate_UsesPositionalGatewayID(t *testing.T) {
	var updatedGatewayID string

	mockGatewayAPI := &api.MockGatewayAPI{
		UpdateGatewayV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			request *api.UpdateGatewayV5Request,
		) (*api.GatewayV5GetResponse, error) {
			updatedGatewayID = gatewayID
			return &api.GatewayV5GetResponse{
				ID:     gatewayID,
				Name:   "renamed-gateway",
				Slug:   "renamed-gateway",
				Type:   api.IngressTypeManual,
				Status: api.GatewayV5StatusReady,
			}, nil
		},
	}

	cmd := setupGatewayUpdateTestCommand()
	cmd.Flags().Set("name", "renamed-gateway")

	err := Update(Dependencies{GatewayAPI: mockGatewayAPI}, cmd, nil, gatewayUpdateTestProfile(), []string{"gateway-positional"})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if updatedGatewayID != "gateway-positional" {
		t.Fatalf("Expected positional gateway ID gateway-positional, got %q", updatedGatewayID)
	}
}

func TestUpdate_RequiresGatewayID(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{
			name: "no args",
		},
		{
			name: "too many args",
			args: []string{"gateway-123", "gateway-456"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := setupGatewayUpdateTestCommand()
			cmd.Flags().Set("name", "renamed-gateway")

			err := Update(Dependencies{GatewayAPI: &api.MockGatewayAPI{}}, cmd, nil, gatewayUpdateTestProfile(), tt.args)
			if err == nil {
				t.Fatal("Expected error, got nil")
			}

			expected := "GATEWAY_ID is required"
			if err.Error() != expected {
				t.Fatalf("Expected error %q, got %q", expected, err.Error())
			}
		})
	}
}

func TestUpdate_PropagatesAPIError(t *testing.T) {
	mockGatewayAPI := &api.MockGatewayAPI{
		UpdateGatewayV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			request *api.UpdateGatewayV5Request,
		) (*api.GatewayV5GetResponse, error) {
			return nil, errors.New("gateway not found")
		},
	}

	cmd := setupGatewayUpdateTestCommand()
	cmd.Flags().Set("name", "renamed-gateway")

	err := Update(Dependencies{GatewayAPI: mockGatewayAPI}, cmd, nil, gatewayUpdateTestProfile(), []string{"gateway-123"})
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	if !strings.Contains(err.Error(), "failed to update gateway: gateway not found") {
		t.Fatalf("Expected wrapped API error, got %q", err.Error())
	}
}
