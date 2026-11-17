package service

import (
	"strings"
	"testing"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/spf13/cobra"
)

func newDescribeCommand() *cobra.Command {
	cmd := setupTestCommand()
	cmd.Flags().String("cluster-name", "", "Cluster name")
	cmd.Flags().String("cluster-id", "", "Cluster ID")
	return cmd
}

// describeService builds a LocationService whose single-cluster read returns the cluster with the given
// id, recording the id it was called with.
func describeService(t *testing.T, clusters []api.InfraAggregateCluster, requestedID *string) LocationService {
	t.Helper()

	mockAPI := &api.MockLocationAPI{
		ListAggregatedFunc: func(_ configuration_models.EndpointsV2, _, _ string, _ ...api.LocationListOption) ([]api.InfraAggregateCluster, error) {
			return clusters, nil
		},
		GetAggregatedFunc: func(_ configuration_models.EndpointsV2, _, _, clusterID string) (*api.InfraAggregateCluster, error) {
			*requestedID = clusterID

			for _, cluster := range clusters {
				if cluster.ClusterID == clusterID {
					return &cluster, nil
				}
			}

			return nil, &api.InfraClusterNotFoundError{ClusterID: clusterID}
		},
	}

	return NewLocationService(mockLocationConfig(), mockAPI, &api.UserAPI{}, nil)
}

func TestLocationService_Describe_ResolvesTheClusterName(t *testing.T) {
	clusters := []api.InfraAggregateCluster{
		{ClusterID: "cluster-1", Name: "Milan DC"},
	}
	requestedID := ""

	service := describeService(t, clusters, &requestedID)

	cmd := newDescribeCommand()
	cmd.Flags().Set("cluster-name", "Milan DC")

	if err := service.Describe(cmd, []string{}); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if requestedID != "cluster-1" {
		t.Fatalf("Expected cluster-1 to be described, got %q", requestedID)
	}
}

func TestLocationService_Describe_AmbiguousName(t *testing.T) {
	clusters := []api.InfraAggregateCluster{
		{ClusterID: "cluster-1", Name: "Milan DC"},
		{ClusterID: "cluster-2", Name: "Milan DC"},
	}
	requestedID := ""

	service := describeService(t, clusters, &requestedID)

	cmd := newDescribeCommand()
	cmd.Flags().Set("cluster-name", "Milan DC")

	err := service.Describe(cmd, []string{})
	if err == nil {
		t.Fatal("Expected an error when several locations share the name")
	}

	if !strings.Contains(err.Error(), "several clusters match") {
		t.Fatalf("Expected an ambiguity error, got %q", err.Error())
	}

	if requestedID != "" {
		t.Fatalf("Expected no cluster to be described, got %q", requestedID)
	}
}

func TestLocationService_Describe_NameNotFound(t *testing.T) {
	clusters := []api.InfraAggregateCluster{
		{ClusterID: "cluster-1", Name: "Milan DC"},
	}
	requestedID := ""

	service := describeService(t, clusters, &requestedID)

	cmd := newDescribeCommand()
	cmd.Flags().Set("cluster-name", "missing")

	err := service.Describe(cmd, []string{})
	if err == nil {
		t.Fatal("Expected a not found error")
	}

	if !strings.Contains(err.Error(), "not found") {
		t.Fatalf("Expected a not found error, got %q", err.Error())
	}
}
