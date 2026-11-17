package service

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/spf13/cobra"
)

func newMutationCommand() *cobra.Command {
	cmd := setupTestCommand()
	cmd.Flags().String("cluster-id", "", "Cluster ID")
	cmd.Flags().String("name", "", "Name")
	cmd.Flags().String("description", "", "Description")
	cmd.Flags().String("position", "", "Position")
	cmd.Flags().String("contact-user", "", "Contact user")
	return cmd
}

func mockLocationConfig() *configuration_handler.MockConfigurationHandler {
	mockCfg := configuration_handler.NewMockConfigurationHandler()
	mockCfg.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{
			APIKey:         "test-api-key",
			OrganizationID: "test-org-id",
			Output:         configuration_models.OutputHuman,
			Endpoints: configuration_models.EndpointsV2{
				CH: "https://ch.example.com",
			},
		}, nil
	}
	return mockCfg
}

func TestLocationService_Update_SendsOnlyProvidedFields(t *testing.T) {
	var captured *api.UpdateInfraClusterRequest

	mockAPI := &api.MockLocationAPI{
		UpdateFunc: func(_ configuration_models.EndpointsV2, _, _, _ string, req *api.UpdateInfraClusterRequest) error {
			captured = req
			return nil
		},
	}

	service := NewLocationService(mockLocationConfig(), mockAPI, &api.UserAPI{}, nil)

	cmd := newMutationCommand()
	cmd.Flags().Set("cluster-id", "cluster-1")
	cmd.Flags().Set("name", "New name")

	if err := service.Update(cmd, []string{}); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if captured == nil {
		t.Fatal("Expected update request to be sent")
	}
	if captured.DisplayName == nil || *captured.DisplayName != "New name" {
		t.Errorf("Expected display name to be sent, got %v", captured.DisplayName)
	}
	if captured.Description != nil {
		t.Errorf("Expected description to be omitted, got %v", *captured.Description)
	}
	if captured.Position != nil {
		t.Errorf("Expected position to be omitted, got %v", *captured.Position)
	}
	if captured.ContactIAMUserID != nil {
		t.Errorf("Expected contact user to be omitted, got %v", *captured.ContactIAMUserID)
	}

	if !strings.Contains(cmd.OutOrStdout().(*bytes.Buffer).String(), "updated successfully") {
		t.Errorf("Expected success message, got %q", cmd.OutOrStdout().(*bytes.Buffer).String())
	}
}

func TestLocationService_Update_NoFields(t *testing.T) {
	mockAPI := &api.MockLocationAPI{}

	service := NewLocationService(mockLocationConfig(), mockAPI, &api.UserAPI{}, nil)

	cmd := newMutationCommand()
	cmd.Flags().Set("cluster-id", "cluster-1")

	err := service.Update(cmd, []string{})
	if err == nil {
		t.Fatal("Expected an error when no field is provided")
	}
	if !strings.Contains(err.Error(), "no fields to update") {
		t.Errorf("Expected 'no fields to update' error, got %q", err.Error())
	}
}

func TestLocationService_Update_Conflict(t *testing.T) {
	mockAPI := &api.MockLocationAPI{
		UpdateFunc: func(_ configuration_models.EndpointsV2, _, _, _ string, _ *api.UpdateInfraClusterRequest) error {
			return &api.InfraClusterConflictError{Reason: api.InfraClusterConflictReasonDeletionInProgress}
		},
	}

	service := NewLocationService(mockLocationConfig(), mockAPI, &api.UserAPI{}, nil)

	cmd := newMutationCommand()
	cmd.Flags().Set("cluster-id", "cluster-1")
	cmd.Flags().Set("description", "desc")

	err := service.Update(cmd, []string{})
	if err == nil {
		t.Fatal("Expected an error on conflict")
	}
	if !strings.Contains(err.Error(), "already in progress") {
		t.Errorf("Expected deletion-in-progress message, got %q", err.Error())
	}
}

func TestLocationService_Update_ContactUserByUsername(t *testing.T) {
	var captured *api.UpdateInfraClusterRequest

	mockAPI := &api.MockLocationAPI{
		UpdateFunc: func(_ configuration_models.EndpointsV2, _, _, _ string, req *api.UpdateInfraClusterRequest) error {
			captured = req
			return nil
		},
	}

	mockUserAPI := &api.MockUserAPI{
		ListIAMUsersFunc: func(_ configuration_models.EndpointsV2, _, _ string, _ *bool, search string, _ int, _ int, _ string, _ string) (*api.GenericPaginatedResponse[api.IAMUserListItem], error) {
			return &api.GenericPaginatedResponse[api.IAMUserListItem]{
				Data: []api.IAMUserListItem{{ID: "user-uuid-1", Username: search}},
			}, nil
		},
	}

	service := NewLocationService(mockLocationConfig(), mockAPI, mockUserAPI, nil)

	cmd := newMutationCommand()
	cmd.Flags().Set("cluster-id", "cluster-1")
	cmd.Flags().Set("contact-user", "jdoe")

	if err := service.Update(cmd, []string{}); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if captured == nil || captured.ContactIAMUserID == nil || *captured.ContactIAMUserID != "user-uuid-1" {
		t.Fatalf("Expected contact user id to be resolved to user-uuid-1, got %v", captured)
	}
}

func TestLocationService_Update_ContactUserByUUID(t *testing.T) {
	var captured *api.UpdateInfraClusterRequest

	mockAPI := &api.MockLocationAPI{
		UpdateFunc: func(_ configuration_models.EndpointsV2, _, _, _ string, req *api.UpdateInfraClusterRequest) error {
			captured = req
			return nil
		},
	}

	mockUserAPI := &api.MockUserAPI{
		ListIAMUsersFunc: func(_ configuration_models.EndpointsV2, _, _ string, _ *bool, _ string, _ int, _ int, _ string, _ string) (*api.GenericPaginatedResponse[api.IAMUserListItem], error) {
			t.Error("Expected no username lookup when a uuid is provided")
			return nil, nil
		},
	}

	service := NewLocationService(mockLocationConfig(), mockAPI, mockUserAPI, nil)

	cmd := newMutationCommand()
	cmd.Flags().Set("cluster-id", "cluster-1")
	cmd.Flags().Set("contact-user", "550e8400-e29b-41d4-a716-446655440000")

	if err := service.Update(cmd, []string{}); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if captured == nil || captured.ContactIAMUserID == nil || *captured.ContactIAMUserID != "550e8400-e29b-41d4-a716-446655440000" {
		t.Fatalf("Expected contact user id to be forwarded as-is, got %v", captured)
	}
}

func TestLocationService_Delete_Success(t *testing.T) {
	mockAPI := &api.MockLocationAPI{
		DeleteFunc: func(_ configuration_models.EndpointsV2, _, _, clusterID string) (*api.InfraClusterDeleteResponse, error) {
			return &api.InfraClusterDeleteResponse{ID: "process-1"}, nil
		},
	}

	mockProcess := &api.MockProcessAPI{
		GetProcessFunc: func(_ configuration_models.EndpointsV2, _, _, processID string) (*api.Process, error) {
			return &api.Process{ID: processID, Type: api.ProcessTypeClusterDeletion, Status: api.ProcessStatusSuccess}, nil
		},
	}

	service := NewLocationService(mockLocationConfig(), mockAPI, &api.UserAPI{}, mockProcess)

	cmd := newMutationCommand()
	cmd.Flags().Set("cluster-id", "cluster-1")

	if err := service.Delete(cmd, []string{}); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	output := cmd.OutOrStdout().(*bytes.Buffer).String()
	if !strings.Contains(output, "deletion started") || !strings.Contains(output, "deleted successfully") {
		t.Errorf("Expected deletion started and success messages, got %q", output)
	}
}

func TestLocationService_Delete_GuardProfileInstalled(t *testing.T) {
	mockAPI := &api.MockLocationAPI{
		DeleteFunc: func(_ configuration_models.EndpointsV2, _, _, _ string) (*api.InfraClusterDeleteResponse, error) {
			return nil, &api.InfraClusterConflictError{
				Reason:        api.InfraClusterConflictReasonProfileInstalled,
				ConflictsWith: "ProfileGateway",
			}
		},
	}

	service := NewLocationService(mockLocationConfig(), mockAPI, &api.UserAPI{}, &api.MockProcessAPI{})

	cmd := newMutationCommand()
	cmd.Flags().Set("cluster-id", "cluster-1")

	err := service.Delete(cmd, []string{})
	if err == nil {
		t.Fatal("Expected a guard error")
	}
	if !strings.Contains(err.Error(), "ProfileGateway") || !strings.Contains(err.Error(), "cannot be deleted") {
		t.Errorf("Expected a guard message mentioning ProfileGateway, got %q", err.Error())
	}
}

func TestLocationService_Delete_GuardDeletionInProgress(t *testing.T) {
	mockAPI := &api.MockLocationAPI{
		DeleteFunc: func(_ configuration_models.EndpointsV2, _, _, _ string) (*api.InfraClusterDeleteResponse, error) {
			return nil, &api.InfraClusterConflictError{Reason: api.InfraClusterConflictReasonDeletionInProgress}
		},
	}

	service := NewLocationService(mockLocationConfig(), mockAPI, &api.UserAPI{}, &api.MockProcessAPI{})

	cmd := newMutationCommand()
	cmd.Flags().Set("cluster-id", "cluster-1")

	err := service.Delete(cmd, []string{})
	if err == nil {
		t.Fatal("Expected a guard error")
	}
	if !strings.Contains(err.Error(), "already in progress") {
		t.Errorf("Expected deletion-in-progress message, got %q", err.Error())
	}
}

func TestLocationService_Delete_ProcessFailed(t *testing.T) {
	mockAPI := &api.MockLocationAPI{
		DeleteFunc: func(_ configuration_models.EndpointsV2, _, _, _ string) (*api.InfraClusterDeleteResponse, error) {
			return &api.InfraClusterDeleteResponse{ID: "process-1"}, nil
		},
	}

	mockProcess := &api.MockProcessAPI{
		GetProcessFunc: func(_ configuration_models.EndpointsV2, _, _, processID string) (*api.Process, error) {
			return &api.Process{
				ID:     processID,
				Type:   api.ProcessTypeClusterDeletion,
				Status: api.ProcessStatusFailed,
				Data:   json.RawMessage(`{"cluster_id":"cluster-1","error":{"code":"cluster_deletion_failed","message":"boom"}}`),
			}, nil
		},
	}

	service := NewLocationService(mockLocationConfig(), mockAPI, &api.UserAPI{}, mockProcess)

	cmd := newMutationCommand()
	cmd.Flags().Set("cluster-id", "cluster-1")

	err := service.Delete(cmd, []string{})
	if err == nil {
		t.Fatal("Expected a failure error")
	}
	if !strings.Contains(err.Error(), "boom") || !strings.Contains(err.Error(), "cluster_deletion_failed") {
		t.Errorf("Expected the process error to be reported, got %q", err.Error())
	}
}

func TestLocationService_Delete_PollsUntilTerminal(t *testing.T) {
	previousInterval := clusterDeletionPollInterval
	clusterDeletionPollInterval = time.Millisecond
	defer func() { clusterDeletionPollInterval = previousInterval }()

	callCount := 0

	mockAPI := &api.MockLocationAPI{
		DeleteFunc: func(_ configuration_models.EndpointsV2, _, _, _ string) (*api.InfraClusterDeleteResponse, error) {
			return &api.InfraClusterDeleteResponse{ID: "process-1"}, nil
		},
	}

	mockProcess := &api.MockProcessAPI{
		GetProcessFunc: func(_ configuration_models.EndpointsV2, _, _, processID string) (*api.Process, error) {
			callCount++
			status := api.ProcessStatusRunning
			if callCount >= 3 {
				status = api.ProcessStatusSuccess
			}
			return &api.Process{ID: processID, Type: api.ProcessTypeClusterDeletion, Status: status}, nil
		},
	}

	service := NewLocationService(mockLocationConfig(), mockAPI, &api.UserAPI{}, mockProcess)

	cmd := newMutationCommand()
	cmd.Flags().Set("cluster-id", "cluster-1")

	if err := service.Delete(cmd, []string{}); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if callCount < 3 {
		t.Errorf("Expected the process to be polled until success, got %d calls", callCount)
	}
}

func TestLocationService_Update_EmptyProvidedValueClearsTheField(t *testing.T) {
	var captured *api.UpdateInfraClusterRequest

	mockAPI := &api.MockLocationAPI{
		UpdateFunc: func(_ configuration_models.EndpointsV2, _, _, _ string, req *api.UpdateInfraClusterRequest) error {
			captured = req
			return nil
		},
	}

	mockUserAPI := &api.MockUserAPI{
		ListIAMUsersFunc: func(_ configuration_models.EndpointsV2, _, _ string, _ *bool, _ string, _ int, _ int, _ string, _ string) (*api.GenericPaginatedResponse[api.IAMUserListItem], error) {
			t.Error("Expected no username lookup when the contact is cleared")
			return nil, nil
		},
	}

	service := NewLocationService(mockLocationConfig(), mockAPI, mockUserAPI, nil)

	cmd := newMutationCommand()
	cmd.Flags().Set("cluster-id", "cluster-1")
	cmd.Flags().Set("description", "")
	cmd.Flags().Set("name", "")
	cmd.Flags().Set("contact-user", "")

	if err := service.Update(cmd, []string{}); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if captured == nil {
		t.Fatal("Expected update request to be sent")
	}
	if captured.Description == nil || *captured.Description != "" {
		t.Errorf("Expected an empty description to be sent, got %v", captured.Description)
	}
	if captured.DisplayName == nil || *captured.DisplayName != "" {
		t.Errorf("Expected an empty display name to be sent, got %v", captured.DisplayName)
	}
	if captured.ContactIAMUserID == nil || *captured.ContactIAMUserID != "" {
		t.Errorf("Expected an empty contact to be sent, got %v", captured.ContactIAMUserID)
	}
	if captured.Position != nil {
		t.Errorf("Expected position to be omitted, got %v", *captured.Position)
	}
}

func TestLocationService_Delete_StructuredOutputIsSingleDocument(t *testing.T) {
	mockAPI := &api.MockLocationAPI{
		DeleteFunc: func(_ configuration_models.EndpointsV2, _, _, _ string) (*api.InfraClusterDeleteResponse, error) {
			return &api.InfraClusterDeleteResponse{ID: "process-1"}, nil
		},
	}

	mockProcess := &api.MockProcessAPI{
		GetProcessFunc: func(_ configuration_models.EndpointsV2, _, _, processID string) (*api.Process, error) {
			return &api.Process{ID: processID, Type: api.ProcessTypeClusterDeletion, Status: api.ProcessStatusSuccess}, nil
		},
	}

	service := NewLocationService(mockLocationConfig(), mockAPI, &api.UserAPI{}, mockProcess)

	cmd := newMutationCommand()
	cmd.Flags().Set("cluster-id", "cluster-1")
	cmd.Flags().Set("output", "json")

	if err := service.Delete(cmd, []string{}); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	output := cmd.OutOrStdout().(*bytes.Buffer).String()

	var parsed map[string]any
	if err := json.Unmarshal([]byte(output), &parsed); err != nil {
		t.Fatalf("Expected a single JSON document, got %q: %v", output, err)
	}
	if message, _ := parsed["message"].(string); !strings.Contains(message, "deleted successfully") {
		t.Errorf("Unexpected message %v", parsed["message"])
	}
}
