package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
)

func TestLocationAPI_Delete_Accepted(t *testing.T) {
	var method, path string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		method = r.Method
		path = r.URL.Path
		w.WriteHeader(http.StatusAccepted)
		_, _ = w.Write([]byte(`{"id":"process-1"}`))
	}))
	defer server.Close()

	api := NewLocationAPI()
	response, err := api.Delete(configuration_models.EndpointsV2{CH: server.URL}, "api-key", "org-1", "cluster-1")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if method != http.MethodDelete {
		t.Errorf("Expected DELETE, got %s", method)
	}
	if path != "/v1/organizations/org-1/infra/clusters/cluster-1" {
		t.Errorf("Unexpected path %q", path)
	}
	if response.ID != "process-1" {
		t.Errorf("Expected process-1, got %q", response.ID)
	}
}

func TestLocationAPI_Delete_ConflictProfileInstalled(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusConflict)
		_, _ = w.Write([]byte(`{"message":"conflict","type":"conflict","data":{"reason":"profile_installed","conflicts_with":"ProfileGateway"}}`))
	}))
	defer server.Close()

	api := NewLocationAPI()
	_, err := api.Delete(configuration_models.EndpointsV2{CH: server.URL}, "api-key", "org-1", "cluster-1")
	if err == nil {
		t.Fatal("Expected a conflict error")
	}

	var conflict *InfraClusterConflictError
	if !errors.As(err, &conflict) {
		t.Fatalf("Expected an InfraClusterConflictError, got %T", err)
	}
	if conflict.Reason != InfraClusterConflictReasonProfileInstalled {
		t.Errorf("Expected profile_installed reason, got %q", conflict.Reason)
	}
	if conflict.ConflictsWith != "ProfileGateway" {
		t.Errorf("Expected ProfileGateway, got %q", conflict.ConflictsWith)
	}
}

func TestLocationAPI_Delete_NotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	api := NewLocationAPI()
	_, err := api.Delete(configuration_models.EndpointsV2{CH: server.URL}, "api-key", "org-1", "cluster-1")
	if err == nil {
		t.Fatal("Expected a not found error")
	}

	var notFound *InfraClusterNotFoundError
	if !errors.As(err, &notFound) {
		t.Fatalf("Expected an InfraClusterNotFoundError, got %T", err)
	}
}

func TestLocationAPI_Update_PatchesOnlyProvidedFields(t *testing.T) {
	var body []byte
	var method string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		method = r.Method
		body, _ = io.ReadAll(r.Body)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	displayName := "New name"
	api := NewLocationAPI()
	err := api.Update(
		configuration_models.EndpointsV2{CH: server.URL},
		"api-key",
		"org-1",
		"cluster-1",
		&UpdateInfraClusterRequest{DisplayName: &displayName},
	)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if method != http.MethodPatch {
		t.Errorf("Expected PATCH, got %s", method)
	}

	var sent map[string]any
	if err := json.Unmarshal(body, &sent); err != nil {
		t.Fatalf("Failed to parse request body %q: %v", string(body), err)
	}
	if sent["display_name"] != "New name" {
		t.Errorf("Expected display_name to be sent, got %v", sent["display_name"])
	}
	if _, ok := sent["description"]; ok {
		t.Errorf("Expected description to be omitted, got %v", sent["description"])
	}
}

func TestLocationAPI_GetAggregated_NotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	api := NewLocationAPI()
	_, err := api.GetAggregated(configuration_models.EndpointsV2{CH: server.URL}, "api-key", "org-1", "cluster-1")
	if err == nil {
		t.Fatal("Expected a not found error")
	}

	var notFound *InfraClusterNotFoundError
	if !errors.As(err, &notFound) {
		t.Fatalf("Expected an InfraClusterNotFoundError, got %T", err)
	}
}

func TestLocationAPI_Update_Conflict(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusConflict)
		_, _ = w.Write([]byte(`{"message":"conflict","type":"conflict","data":{"reason":"deletion_in_progress"}}`))
	}))
	defer server.Close()

	description := "desc"
	api := NewLocationAPI()
	err := api.Update(
		configuration_models.EndpointsV2{CH: server.URL},
		"api-key",
		"org-1",
		"cluster-1",
		&UpdateInfraClusterRequest{Description: &description},
	)
	if err == nil {
		t.Fatal("Expected a conflict error")
	}

	var conflict *InfraClusterConflictError
	if !errors.As(err, &conflict) {
		t.Fatalf("Expected an InfraClusterConflictError, got %T", err)
	}
	if conflict.Reason != InfraClusterConflictReasonDeletionInProgress {
		t.Errorf("Expected deletion_in_progress reason, got %q", conflict.Reason)
	}
}

func TestLocationAPI_Update_NotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	description := "desc"
	api := NewLocationAPI()
	err := api.Update(
		configuration_models.EndpointsV2{CH: server.URL},
		"api-key",
		"org-1",
		"cluster-1",
		&UpdateInfraClusterRequest{Description: &description},
	)
	if err == nil {
		t.Fatal("Expected a not found error")
	}

	var notFound *InfraClusterNotFoundError
	if !errors.As(err, &notFound) {
		t.Fatalf("Expected an InfraClusterNotFoundError, got %T", err)
	}
}

func TestLocationAPI_GetAggregated(t *testing.T) {
	var path string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path = r.URL.Path
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"cluster_id":"cluster-1","name":"c1","type":"physical","details":{"last_update":"2024-01-15T10:30:00Z","next_update":"2024-01-15T11:30:00Z","is_update_ok":true,"nodes":[],"virtual_nodes":[]}}`))
	}))
	defer server.Close()

	api := NewLocationAPI()
	cluster, err := api.GetAggregated(configuration_models.EndpointsV2{CH: server.URL}, "api-key", "org-1", "cluster-1")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !strings.HasSuffix(path, "/v1/organizations/org-1/infra/aggregate_clusters/cluster-1") {
		t.Errorf("Unexpected path %q", path)
	}
	if cluster.ClusterID != "cluster-1" {
		t.Errorf("Expected cluster-1, got %q", cluster.ClusterID)
	}
}

// TestInfrastructureClusterUnmarshalAppliesDisplayName pins the implicit override: the display name
// replaces the cluster name in the parsed value and is never exposed as a separate field.
func TestInfrastructureClusterUnmarshalAppliesDisplayName(t *testing.T) {
	payload := `{
		"cluster_id": "cluster-1",
		"name": "k8s-name",
		"type": "physical",
		"display_name": "Milan DC",
		"position": "Milan, Italy"
	}`

	var cluster InfrastructureCluster
	if err := json.Unmarshal([]byte(payload), &cluster); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if cluster.Name != "Milan DC" {
		t.Errorf("Expected the display name to override the name, got %q", cluster.Name)
	}

	raw, err := json.Marshal(cluster)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if strings.Contains(string(raw), "display_name") {
		t.Errorf("Expected display_name to be hidden, got %s", string(raw))
	}
}

// TestInfraAggregateClusterUnmarshalAppliesDisplayName pins the same override on the single-location
// representation.
func TestInfraAggregateClusterUnmarshalAppliesDisplayName(t *testing.T) {
	payload := `{
		"cluster_id": "cluster-1",
		"name": "k8s-name",
		"type": "physical",
		"display_name": "Milan DC"
	}`

	var cluster InfraAggregateCluster
	if err := json.Unmarshal([]byte(payload), &cluster); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if cluster.Name != "Milan DC" {
		t.Errorf("Expected the display name to override the name, got %q", cluster.Name)
	}

	raw, err := json.Marshal(cluster)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if strings.Contains(string(raw), "display_name") {
		t.Errorf("Expected display_name to be hidden, got %s", string(raw))
	}
}

// TestInfrastructureClusterUnmarshalKeepsNameWhenDisplayNameIsAbsent pins the fallback to the
// Kubernetes cluster name.
func TestInfrastructureClusterUnmarshalKeepsNameWhenDisplayNameIsAbsent(t *testing.T) {
	var cluster InfrastructureCluster
	if err := json.Unmarshal([]byte(`{"cluster_id":"cluster-1","name":"k8s-name","type":"physical"}`), &cluster); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if cluster.Name != "k8s-name" {
		t.Errorf("Expected the cluster name to be kept, got %q", cluster.Name)
	}
}

// TestInfrastructureClusterUnmarshalKeepsContactName pins that the optional contact display name
// survives the custom unmarshal of the list payload.
func TestInfrastructureClusterUnmarshalKeepsContactName(t *testing.T) {
	payload := `{
		"cluster_id": "cluster-1",
		"name": "k8s-name",
		"type": "physical",
		"contact_iam_user_id": "695ed3dd-e77d-42b9-88ed-70bd3a1704ee",
		"contact_iam_user_name": "Mario Rossi"
	}`

	var cluster InfrastructureCluster
	if err := json.Unmarshal([]byte(payload), &cluster); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if cluster.ContactIAMUserName == nil || *cluster.ContactIAMUserName != "Mario Rossi" {
		t.Errorf("Expected the contact name to be decoded, got %v", cluster.ContactIAMUserName)
	}
}
