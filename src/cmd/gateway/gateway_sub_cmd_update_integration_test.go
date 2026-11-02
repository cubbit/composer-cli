package cmd_gateway

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	servicegateway "github.com/cubbit/composer-cli/src/service/gateway"
)

func setupGatewayUpdateIntegrationCommand(gatewayService servicegateway.GatewayServiceInterface) (*bytes.Buffer, func([]string) error) {
	gatewayCmd := NewGatewayCmd(gatewayService)
	gatewayCmd.PersistentFlags().String("profile", "", "Profile")
	gatewayCmd.PersistentFlags().String("output", "human", "Output format")
	gatewayCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	gatewayCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	commandOutput := new(bytes.Buffer)
	gatewayCmd.SetOut(commandOutput)
	gatewayCmd.SetErr(commandOutput)

	return commandOutput, func(args []string) error {
		gatewayCmd.SetArgs(args)
		return gatewayCmd.Execute()
	}
}

func newGatewayUpdateService(gatewayAPI api.GatewayAPIInterface) servicegateway.GatewayService {
	return servicegateway.NewGatewayService(
		newTestGatewayConfig(),
		gatewayAPI,
		&api.MockSwarmAPI{},
		&api.MockRedundancyClassAPI{},
		&api.MockProcessAPI{},
		&api.MockLocationAPI{},
	)
}

func TestGatewaySubCmd_Update_Integration_HumanOutput(t *testing.T) {
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
			if gatewayID != "gw-001" {
				t.Fatalf("Expected gateway ID gw-001, got %q", gatewayID)
			}
			if request.Name == nil || *request.Name != "renamed-gateway" {
				t.Fatalf("Expected name renamed-gateway, got %+v", request.Name)
			}
			if request.Description != nil {
				t.Fatalf("Expected description to be omitted, got %q", *request.Description)
			}
			if request.Location != nil {
				t.Fatalf("Expected location to be omitted, got %q", *request.Location)
			}

			return &api.GatewayV5GetResponse{
				ID:     "gw-001",
				Name:   "renamed-gateway",
				Slug:   "renamed-gateway",
				Type:   api.IngressTypeManual,
				Status: api.GatewayV5StatusReady,
			}, nil
		},
	}

	commandOutput, execute := setupGatewayUpdateIntegrationCommand(newGatewayUpdateService(mockGatewayAPI))

	if err := execute([]string{"update", "gw-001", "--name", "renamed-gateway"}); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	output := commandOutput.String()
	if !strings.Contains(output, "Gateway renamed-gateway updated successfully.") {
		t.Fatalf("Expected success message, got %q", output)
	}
	if !strings.Contains(output, "Gateway: renamed-gateway") || !strings.Contains(output, "ID: gw-001") {
		t.Fatalf("Expected gateway details, got %q", output)
	}
}

func TestGatewaySubCmd_Update_Integration_JSONOutput(t *testing.T) {
	mockGatewayAPI := &api.MockGatewayAPI{
		UpdateGatewayV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			request *api.UpdateGatewayV5Request,
		) (*api.GatewayV5GetResponse, error) {
			description := "Updated description"
			location := "eu-west-1"

			return &api.GatewayV5GetResponse{
				ID:          "gw-json",
				Name:        "json-gateway",
				Slug:        "json-gateway",
				Description: &description,
				Location:    location,
				Type:        api.IngressTypeManual,
				Status:      api.GatewayV5StatusReady,
			}, nil
		},
	}

	commandOutput, execute := setupGatewayUpdateIntegrationCommand(newGatewayUpdateService(mockGatewayAPI))

	if err := execute([]string{
		"update", "gw-json",
		"--description", "Updated description",
		"--location", "eu-west-1",
		"--output", "json",
	}); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := `{
  "id": "gw-json",
  "name": "json-gateway",
  "slug": "json-gateway",
  "description": "Updated description",
  "location": "eu-west-1",
  "type": "manual",
  "redundancy_classes": null,
  "status": "ready",
  "legacy": false
}
`
	if output := commandOutput.String(); output != expectedOutput {
		t.Fatalf("Expected JSON output %q, got %q", expectedOutput, output)
	}
}

func TestGatewaySubCmd_Update_Integration_YAMLOutput_ClearsFields(t *testing.T) {
	mockGatewayAPI := &api.MockGatewayAPI{
		UpdateGatewayV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			request *api.UpdateGatewayV5Request,
		) (*api.GatewayV5GetResponse, error) {
			if request.Description == nil || *request.Description != "" {
				t.Fatalf("Expected an explicit empty description, got %+v", request.Description)
			}
			if request.Location == nil || *request.Location != "" {
				t.Fatalf("Expected an explicit empty location, got %+v", request.Location)
			}

			description := ""
			return &api.GatewayV5GetResponse{
				ID:          "gw-yaml-clear",
				Name:        "cleared-gateway",
				Slug:        "cleared-gateway",
				Description: &description,
				Location:    "",
				Type:        api.IngressTypeManual,
				Status:      api.GatewayV5StatusReady,
			}, nil
		},
	}

	commandOutput, execute := setupGatewayUpdateIntegrationCommand(newGatewayUpdateService(mockGatewayAPI))

	if err := execute([]string{
		"update", "gw-yaml-clear",
		"--description", "",
		"--location", "",
		"--output", "yaml",
	}); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	output := commandOutput.String()
	for _, expected := range []string{
		`id: gw-yaml-clear`,
		`description: ""`,
		`location: ""`,
		`legacy: false`,
	} {
		if !strings.Contains(output, expected) {
			t.Fatalf("Expected YAML output to contain %s, got %q", expected, output)
		}
	}
}

func TestGatewaySubCmd_Update_Integration_RequiresAtLeastOneField(t *testing.T) {
	updateCalled := false
	mockGatewayAPI := &api.MockGatewayAPI{
		UpdateGatewayV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			request *api.UpdateGatewayV5Request,
		) (*api.GatewayV5GetResponse, error) {
			updateCalled = true
			return &api.GatewayV5GetResponse{}, nil
		},
	}

	_, execute := setupGatewayUpdateIntegrationCommand(newGatewayUpdateService(mockGatewayAPI))

	err := execute([]string{"update", "gw-001"})
	if err == nil {
		t.Fatal("Expected validation error, got nil")
	}
	if !strings.Contains(err.Error(), "specify at least one of --name, --description or --location") {
		t.Fatalf("Expected validation error, got %q", err.Error())
	}
	if updateCalled {
		t.Fatal("UpdateGatewayV5 should not be called without update fields")
	}
}

func TestGatewaySubCmd_Update_Integration_PropagatesAPIError(t *testing.T) {
	mockGatewayAPI := &api.MockGatewayAPI{
		UpdateGatewayV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			request *api.UpdateGatewayV5Request,
		) (*api.GatewayV5GetResponse, error) {
			return nil, errors.New("code status expected 200, but received 404 instead")
		},
	}

	_, execute := setupGatewayUpdateIntegrationCommand(newGatewayUpdateService(mockGatewayAPI))

	err := execute([]string{"update", "gw-001", "--name", "renamed-gateway"})
	if err == nil {
		t.Fatal("Expected API error, got nil")
	}
	if !strings.Contains(err.Error(), "failed to update gateway") || !strings.Contains(err.Error(), "received 404") {
		t.Fatalf("Expected wrapped API error, got %q", err.Error())
	}
}
