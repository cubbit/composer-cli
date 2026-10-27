package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/src/request_utils"
)

type LocationProfileType string

const (
	LocationProfileAgent       LocationProfileType = "ProfileAgent"
	LocationProfileCoordinator LocationProfileType = "ProfileCoordinator"
	LocationProfileGateway     LocationProfileType = "ProfileGateway"
	LocationProfileNodeScanner LocationProfileType = "ProfileNodeScanner"
)

type LocationListOptions struct {
	ProfileType *LocationProfileFilter
}

type LocationProfileFilter struct {
	Op     string // "in" or "notin"
	Values []LocationProfileType
}

type LocationListOption func(*LocationListOptions)

func WithProfileType(op string, types ...LocationProfileType) LocationListOption {
	return func(opts *LocationListOptions) {
		opts.ProfileType = &LocationProfileFilter{Op: op, Values: types}
	}
}

func buildProfileTypeQuery(filter *LocationProfileFilter) string {
	values := make([]string, len(filter.Values))
	for i, v := range filter.Values {
		values[i] = string(v)
	}
	return fmt.Sprintf("profile-type:%s(%s)", filter.Op, strings.Join(values, ","))
}

type LocationAPIInterface interface {
	List(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		opts ...LocationListOption,
	) ([]InfrastructureCluster, error)
	ListAggregated(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		opts ...LocationListOption,
	) ([]InfraAggregateCluster, error)

	CreateVirtualCluster(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		name string,
		description *string,
	) (*InfrastructureCluster, error)

	CreateVirtualNode(
		endpoints configuration_models.EndpointsV2,
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
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	opts ...LocationListOption,
) ([]InfrastructureCluster, error) {
	options := &LocationListOptions{}
	for _, opt := range opts {
		opt(options)
	}

	builder := NewURLBuilder(endpoints.CH).
		Path("v1", "organizations", organizationID, "infra", "clusters")

	if options.ProfileType != nil {
		builder = builder.QueryParam("q", buildProfileTypeQuery(options.ProfileType))
	}

	url := builder.Build()

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
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	opts ...LocationListOption,
) ([]InfraAggregateCluster, error) {
	options := &LocationListOptions{}
	for _, opt := range opts {
		opt(options)
	}

	builder := NewURLBuilder(endpoints.CH).
		Path("v1", "organizations", organizationID, "infra", "aggregate_clusters")

	if options.ProfileType != nil {
		builder = builder.QueryParam("q", buildProfileTypeQuery(options.ProfileType))
	}

	url := builder.Build()

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
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	name string,
	description *string,
) (*InfrastructureCluster, error) {
	url := NewURLBuilder(endpoints.CH).
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
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	clusterID string,
	name string,
	storageType string,
	configuration map[string]any,
) (*InfraAggregateVirtualNodeDetail, error) {
	url := NewURLBuilder(endpoints.CH).
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
