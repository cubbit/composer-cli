package api

import (
	"net/http"

	"github.com/cubbit/composer-cli/src/configuration"
	"github.com/cubbit/composer-cli/src/request_utils"
)

type LocationAPIInterface interface {
	List(
		urlConfig configuration.URLs,
		apiKey string,
		organizationID string,
	) ([]InfrastructureCluster, error)
	ListAggregated(
		urlConfig configuration.URLs,
		apiKey string,
		organizationID string,
	) ([]InfraAggregateCluster, error)

	CreateVirtualCluster(
		urlConfig configuration.URLs,
		apiKey string,
		organizationID string,
		name string,
		description *string,
	) (*InfrastructureCluster, error)

	CreateVirtualNode(
		urlConfig configuration.URLs,
		apiKey string,
		organizationID string,
		clusterID string,
		name string,
		storageType string,
		configuration map[string]any,
	) (*InfraAggregateVirtualNodeDetail, error)
}

type LocationAPI struct{}

func NewLocationAPI() *LocationAPI {
	return &LocationAPI{}
}

func (api *LocationAPI) List(
	urlConfig configuration.URLs,
	apiKey string,
	organizationID string,
) ([]InfrastructureCluster, error) {
	url := NewURLBuilder(urlConfig.ChURL).
		Path("v1", "organizations", organizationID, "infra", "clusters").
		Build()

	var response GenericPaginatedResponse[InfrastructureCluster]

	if err := request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodGet),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		request_utils.WithApiKey(apiKey),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (api *LocationAPI) ListAggregated(
	urlConfig configuration.URLs,
	apiKey string,
	organizationID string,
) ([]InfraAggregateCluster, error) {
	url := NewURLBuilder(urlConfig.ChURL).
		Path("v1", "organizations", organizationID, "infra", "aggregate_clusters").
		Build()

	var response GenericPaginatedResponse[InfraAggregateCluster]

	if err := request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodGet),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		request_utils.WithApiKey(apiKey),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (api *LocationAPI) CreateVirtualCluster(
	urlConfig configuration.URLs,
	apiKey string,
	organizationID string,
	name string,
	description *string,
) (*InfrastructureCluster, error) {
	url := NewURLBuilder(urlConfig.ChURL).
		Path("v1", "organizations", organizationID, "infra", "clusters", "virtual").
		Build()

	var response InfrastructureCluster

	if err := request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPost),
		request_utils.WithExpectedStatusCode(http.StatusCreated),
		request_utils.WithApiKey(apiKey),
		request_utils.WithRequestBody(map[string]interface{}{
			"name":        name,
			"description": description,
		}),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, err
	}

	return &response, nil
}

func (api *LocationAPI) CreateVirtualNode(
	urlConfig configuration.URLs,
	apiKey string,
	organizationID string,
	clusterID string,
	name string,
	storageType string,
	configuration map[string]any,
) (*InfraAggregateVirtualNodeDetail, error) {
	url := NewURLBuilder(urlConfig.ChURL).
		Path("v1", "organizations", organizationID, "infra", "clusters", "virtual", clusterID, "nodes").
		Build()

	var response InfraAggregateVirtualNodeDetail

	if err := request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPost),
		request_utils.WithExpectedStatusCode(http.StatusCreated),
		request_utils.WithApiKey(apiKey),
		request_utils.WithRequestBody(map[string]interface{}{
			"name":                  name,
			"storage_type":          storageType,
			"storage_configuration": configuration,
		}),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, err
	}

	return &response, nil
}
