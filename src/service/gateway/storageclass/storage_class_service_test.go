package storageclass

import (
	"bytes"
	"errors"
	"testing"

	api "github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/spf13/cobra"
)

func setupStorageClassTestCommand() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.SetOut(new(bytes.Buffer))
	cmd.SetErr(new(bytes.Buffer))
	cmd.Flags().String("output", "human", "Output format")
	cmd.Flags().Bool("quiet", false, "Quiet mode")
	cmd.Flags().Bool("no-headers", false, "No headers")
	cmd.Flags().String("gateway-id", "", "Gateway ID")
	cmd.Flags().String("name", "", "Storage class name")
	cmd.Flags().String("id", "", "Storage class ID")
	return cmd
}

func storageClassTestProfile() configuration_models.ProfileV2 {
	return configuration_models.ProfileV2{
		Output:         configuration_models.OutputHuman,
		APIKey:         "test-api-key",
		OrganizationID: "test-org-id",
	}
}

func TestList_SendsGatewayIDFromFlag(t *testing.T) {
	var requestedGatewayID string

	mockGatewayAPI := &api.MockGatewayAPI{
		ListGatewayStorageClassesV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
		) (*api.GatewayStorageClassesV5ListResponse, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}
			requestedGatewayID = gatewayID
			return &api.GatewayStorageClassesV5ListResponse{Data: []api.GatewayStorageClassV5{}}, nil
		},
	}

	cmd := setupStorageClassTestCommand()
	_ = cmd.Flags().Set("gateway-id", "gateway-123")

	err := List(Dependencies{GatewayAPI: mockGatewayAPI}, cmd, nil, storageClassTestProfile(), nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if requestedGatewayID != "gateway-123" {
		t.Fatalf("Expected gateway ID gateway-123, got %q", requestedGatewayID)
	}
}

func TestList_TrimsGatewayID(t *testing.T) {
	var requestedGatewayID string

	mockGatewayAPI := &api.MockGatewayAPI{
		ListGatewayStorageClassesV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
		) (*api.GatewayStorageClassesV5ListResponse, error) {
			requestedGatewayID = gatewayID
			return &api.GatewayStorageClassesV5ListResponse{Data: []api.GatewayStorageClassV5{}}, nil
		},
	}

	cmd := setupStorageClassTestCommand()
	_ = cmd.Flags().Set("gateway-id", "  gateway-123  ")

	err := List(Dependencies{GatewayAPI: mockGatewayAPI}, cmd, nil, storageClassTestProfile(), nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if requestedGatewayID != "gateway-123" {
		t.Fatalf("Expected gateway ID gateway-123, got %q", requestedGatewayID)
	}
}

func TestList_RequiresGatewayID(t *testing.T) {
	cmd := setupStorageClassTestCommand()

	err := List(Dependencies{GatewayAPI: &api.MockGatewayAPI{}}, cmd, nil, storageClassTestProfile(), nil)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	expected := "--gateway-id must not be empty"
	if err.Error() != expected {
		t.Fatalf("Expected error %q, got %q", expected, err.Error())
	}
}

func TestList_PropagatesAPIError(t *testing.T) {
	mockGatewayAPI := &api.MockGatewayAPI{
		ListGatewayStorageClassesV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
		) (*api.GatewayStorageClassesV5ListResponse, error) {
			return nil, errors.New("gateway not found")
		},
	}

	cmd := setupStorageClassTestCommand()
	_ = cmd.Flags().Set("gateway-id", "gateway-123")

	err := List(Dependencies{GatewayAPI: mockGatewayAPI}, cmd, nil, storageClassTestProfile(), nil)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	if err.Error() != "failed to list storage classes: gateway not found" {
		t.Fatalf("Expected wrapped API error, got %q", err.Error())
	}
}

func TestCreate_SendsGatewayIDAndName(t *testing.T) {
	mockGatewayAPI := &api.MockGatewayAPI{
		CreateGatewayStorageClassV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			request *api.CreateGatewayStorageClassV5Request,
		) (*api.GatewayStorageClassV5, error) {
			if apiKey != "test-api-key" || organizationID != "test-org-id" {
				t.Fatalf("Expected credentials to be propagated")
			}
			if gatewayID != "gateway-123" {
				t.Fatalf("Expected gateway ID gateway-123, got %q", gatewayID)
			}
			if request.Name != "cold-data" {
				t.Fatalf("Expected name cold-data, got %q", request.Name)
			}

			return &api.GatewayStorageClassV5{}, nil
		},
	}

	cmd := setupStorageClassTestCommand()
	_ = cmd.Flags().Set("gateway-id", "gateway-123")
	_ = cmd.Flags().Set("name", "cold-data")

	err := Create(Dependencies{GatewayAPI: mockGatewayAPI}, cmd, nil, storageClassTestProfile(), nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestCreate_RejectsEmptyNameBeforeCallingAPI(t *testing.T) {
	createCalled := false

	mockGatewayAPI := &api.MockGatewayAPI{
		CreateGatewayStorageClassV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			request *api.CreateGatewayStorageClassV5Request,
		) (*api.GatewayStorageClassV5, error) {
			createCalled = true
			return &api.GatewayStorageClassV5{}, nil
		},
	}

	cmd := setupStorageClassTestCommand()
	_ = cmd.Flags().Set("gateway-id", "gateway-123")
	_ = cmd.Flags().Set("name", "   ")

	err := Create(Dependencies{GatewayAPI: mockGatewayAPI}, cmd, nil, storageClassTestProfile(), nil)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	if err.Error() != "--name must not be empty" {
		t.Fatalf("Expected empty name error, got %q", err.Error())
	}
	if createCalled {
		t.Fatal("CreateGatewayStorageClassV5 should not be called when the name is empty")
	}
}

func TestCreate_RequiresGatewayID(t *testing.T) {
	createCalled := false

	mockGatewayAPI := &api.MockGatewayAPI{
		CreateGatewayStorageClassV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			request *api.CreateGatewayStorageClassV5Request,
		) (*api.GatewayStorageClassV5, error) {
			createCalled = true
			return &api.GatewayStorageClassV5{}, nil
		},
	}

	cmd := setupStorageClassTestCommand()
	_ = cmd.Flags().Set("name", "cold-data")

	err := Create(Dependencies{GatewayAPI: mockGatewayAPI}, cmd, nil, storageClassTestProfile(), nil)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	if err.Error() != "--gateway-id must not be empty" {
		t.Fatalf("Expected empty gateway-id error, got %q", err.Error())
	}
	if createCalled {
		t.Fatal("CreateGatewayStorageClassV5 should not be called when the gateway ID is empty")
	}
}

func TestCreate_PropagatesAPIError(t *testing.T) {
	mockGatewayAPI := &api.MockGatewayAPI{
		CreateGatewayStorageClassV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			request *api.CreateGatewayStorageClassV5Request,
		) (*api.GatewayStorageClassV5, error) {
			return nil, errors.New("code status expected 201, but received 409 instead")
		},
	}

	cmd := setupStorageClassTestCommand()
	_ = cmd.Flags().Set("gateway-id", "gateway-123")
	_ = cmd.Flags().Set("name", "cold-data")

	err := Create(Dependencies{GatewayAPI: mockGatewayAPI}, cmd, nil, storageClassTestProfile(), nil)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	if err.Error() != "failed to create storage class: code status expected 201, but received 409 instead" {
		t.Fatalf("Expected wrapped API error, got %q", err.Error())
	}
}

func TestDelete_SendsGatewayIDAndStorageClassID(t *testing.T) {
	var (
		requestedGatewayID    string
		deletedStorageClassID string
	)

	mockGatewayAPI := &api.MockGatewayAPI{
		DeleteGatewayStorageClassV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			storageClassID string,
		) error {
			if apiKey != "test-api-key" || organizationID != "test-org-id" {
				t.Fatalf("Expected credentials to be propagated")
			}
			requestedGatewayID = gatewayID
			deletedStorageClassID = storageClassID
			return nil
		},
	}

	cmd := setupStorageClassTestCommand()
	_ = cmd.Flags().Set("gateway-id", "gateway-123")
	_ = cmd.Flags().Set("id", "sc-001")

	err := Delete(Dependencies{GatewayAPI: mockGatewayAPI}, cmd, nil, storageClassTestProfile(), nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if requestedGatewayID != "gateway-123" {
		t.Fatalf("Expected gateway ID gateway-123, got %q", requestedGatewayID)
	}
	if deletedStorageClassID != "sc-001" {
		t.Fatalf("Expected storage class ID sc-001, got %q", deletedStorageClassID)
	}
}

func TestDelete_RejectsEmptyIDBeforeCallingAPI(t *testing.T) {
	deleteCalled := false

	mockGatewayAPI := &api.MockGatewayAPI{
		DeleteGatewayStorageClassV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			storageClassID string,
		) error {
			deleteCalled = true
			return nil
		},
	}

	cmd := setupStorageClassTestCommand()
	_ = cmd.Flags().Set("gateway-id", "gateway-123")
	_ = cmd.Flags().Set("id", "")

	err := Delete(Dependencies{GatewayAPI: mockGatewayAPI}, cmd, nil, storageClassTestProfile(), nil)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	if err.Error() != "--id must not be empty" {
		t.Fatalf("Expected empty id error, got %q", err.Error())
	}
	if deleteCalled {
		t.Fatal("DeleteGatewayStorageClassV5 should not be called when the storage class ID is empty")
	}
}

func TestDelete_RequiresGatewayID(t *testing.T) {
	deleteCalled := false

	mockGatewayAPI := &api.MockGatewayAPI{
		DeleteGatewayStorageClassV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			storageClassID string,
		) error {
			deleteCalled = true
			return nil
		},
	}

	cmd := setupStorageClassTestCommand()
	_ = cmd.Flags().Set("id", "sc-001")

	err := Delete(Dependencies{GatewayAPI: mockGatewayAPI}, cmd, nil, storageClassTestProfile(), nil)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	if err.Error() != "--gateway-id must not be empty" {
		t.Fatalf("Expected empty gateway-id error, got %q", err.Error())
	}
	if deleteCalled {
		t.Fatal("DeleteGatewayStorageClassV5 should not be called when the gateway ID is empty")
	}
}

func TestDelete_PropagatesAPIError(t *testing.T) {
	mockGatewayAPI := &api.MockGatewayAPI{
		DeleteGatewayStorageClassV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			storageClassID string,
		) error {
			return errors.New("code status expected 204, but received 404 instead")
		},
	}

	cmd := setupStorageClassTestCommand()
	_ = cmd.Flags().Set("gateway-id", "gateway-456")
	_ = cmd.Flags().Set("id", "sc-001")

	err := Delete(Dependencies{GatewayAPI: mockGatewayAPI}, cmd, nil, storageClassTestProfile(), nil)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	if err.Error() != "failed to delete storage class: code status expected 204, but received 404 instead" {
		t.Fatalf("Expected wrapped API error, got %q", err.Error())
	}
}
