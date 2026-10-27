package api

import (
	"net/http"

	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/src/request_utils"
)

type SwarmAPIInterface interface {
	CreateSwarmV5(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		request *CreateSwarmV5Request,
	) (*CreateSwarmV5Response, error)
	GetSwarmV5(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		swarmID string,
	) (*SwarmV5Presentation, error)
	ListSwarmsV5(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		page int,
		items int,
	) (*GenericPaginatedResponse[ListSwarmV5ItemPresentation], error)
}

type SwarmAPI struct{}

func NewSwarmAPI() *SwarmAPI {
	return &SwarmAPI{}
}

func (api *SwarmAPI) CreateSwarmV5(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	request *CreateSwarmV5Request,
) (*CreateSwarmV5Response, error) {
	url := NewURLBuilder(endpoints.CH).
		Path("v5", "organizations", organizationID, "swarms").
		Build()

	var response CreateSwarmV5Response

	if err := request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPost),
		request_utils.WithExpectedStatusCode(http.StatusCreated),
		request_utils.WithApiKey(apiKey),
		request_utils.WithRequestBodyObject(request),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, err
	}

	return &response, nil
}

func (api *SwarmAPI) GetSwarmV5(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	swarmID string,
) (*SwarmV5Presentation, error) {
	url := NewURLBuilder(endpoints.CH).
		Path("v5", "organizations", organizationID, "swarms", swarmID).
		Build()

	var response SwarmV5Presentation

	if err := request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodGet),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		request_utils.WithApiKey(apiKey),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, err
	}

	return &response, nil
}

func (api *SwarmAPI) ListSwarmsV5(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	page int,
	items int,
) (*GenericPaginatedResponse[ListSwarmV5ItemPresentation], error) {
	url := NewURLBuilder(endpoints.CH).
		Path("v5", "organizations", organizationID, "swarms").
		QueryParamInt("page", page).
		QueryParamInt("items", items).
		Build()

	var response GenericPaginatedResponse[ListSwarmV5ItemPresentation]

	if err := request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodGet),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		request_utils.WithApiKey(apiKey),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, err
	}

	return &response, nil
}
