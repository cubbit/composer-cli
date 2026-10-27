package api

import "github.com/cubbit/composer-cli/src/configuration/configuration_models"

// MockLocationAPI implements LocationAPIInterface for testing
type MockLocationAPI struct {
	ListFunc              func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, opts ...LocationListOption) ([]InfrastructureCluster, error)
	ListAggregatedFunc    func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, opts ...LocationListOption) ([]InfraAggregateCluster, error)
	CreateVirtualFunc     func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, name string, description *string) (*InfrastructureCluster, error)
	CreateVirtualNodeFunc func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, clusterID string, name string, storageType string, configuration map[string]any) (*InfraAggregateVirtualNodeDetail, error)
	GetAggregatedFunc     func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, clusterID string) (*InfraAggregateCluster, error)
	UpdateFunc            func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, clusterID string, request *UpdateInfraClusterRequest) error
	DeleteFunc            func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, clusterID string) (*InfraClusterDeleteResponse, error)
}

func (m *MockLocationAPI) List(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, opts ...LocationListOption) ([]InfrastructureCluster, error) {
	if m.ListFunc != nil {
		return m.ListFunc(endpoints, apiKey, organizationID)
	}
	return []InfrastructureCluster{}, nil
}

func (m *MockLocationAPI) ListAggregated(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, opts ...LocationListOption) ([]InfraAggregateCluster, error) {
	if m.ListAggregatedFunc != nil {
		return m.ListAggregatedFunc(endpoints, apiKey, organizationID, opts...)
	}
	return []InfraAggregateCluster{}, nil
}

func (m *MockLocationAPI) CreateVirtualCluster(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, name string, description *string) (*InfrastructureCluster, error) {
	if m.CreateVirtualFunc != nil {
		return m.CreateVirtualFunc(endpoints, apiKey, organizationID, name, description)
	}
	return &InfrastructureCluster{ClusterID: "test-cluster-id", Name: name}, nil
}

func (m *MockLocationAPI) CreateVirtualNode(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, clusterID string, name string, storageType string, configuration map[string]any) (*InfraAggregateVirtualNodeDetail, error) {
	if m.CreateVirtualNodeFunc != nil {
		return m.CreateVirtualNodeFunc(endpoints, apiKey, organizationID, clusterID, name, storageType, configuration)
	}
	return &InfraAggregateVirtualNodeDetail{
		NodeID:      "test-node-id",
		NodeName:    name,
		StorageType: InfraVirtualStorageType(storageType),
	}, nil
}

func (m *MockLocationAPI) GetAggregated(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, clusterID string) (*InfraAggregateCluster, error) {
	if m.GetAggregatedFunc != nil {
		return m.GetAggregatedFunc(endpoints, apiKey, organizationID, clusterID)
	}

	return nil, &InfraClusterNotFoundError{ClusterID: clusterID}
}

func (m *MockLocationAPI) Update(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, clusterID string, request *UpdateInfraClusterRequest) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(endpoints, apiKey, organizationID, clusterID, request)
	}
	return nil
}

func (m *MockLocationAPI) Delete(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, clusterID string) (*InfraClusterDeleteResponse, error) {
	if m.DeleteFunc != nil {
		return m.DeleteFunc(endpoints, apiKey, organizationID, clusterID)
	}
	return &InfraClusterDeleteResponse{ID: "test-process-id"}, nil
}
