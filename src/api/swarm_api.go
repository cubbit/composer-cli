package api

import (
	"net/http"

	"github.com/cubbit/composer-cli/src/configuration"
	"github.com/cubbit/composer-cli/src/request_utils"
)

type SwarmAPIInterface interface {
	CreateSwarmV5(
		urlConfig configuration.URLs,
		apiKey string,
		organizationID string,
		request *CreateSwarmV5Request,
	) (*CreateSwarmV5Response, error)
}

type SwarmAPI struct{}

func NewSwarmAPI() *SwarmAPI {
	return &SwarmAPI{}
}

func (api *SwarmAPI) CreateSwarmV5(
	urlConfig configuration.URLs,
	apiKey string,
	organizationID string,
	request *CreateSwarmV5Request,
) (*CreateSwarmV5Response, error) {
	url := NewURLBuilder(urlConfig.ChURL).
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
