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
