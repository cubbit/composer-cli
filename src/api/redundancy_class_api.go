package api

import (
	"net/http"

	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/src/request_utils"
)

type RedundancyClassAPIInterface interface {
	ListRedundancyClassesBySwarm(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		swarmID string,
	) ([]RedundancyClass, error)
}

type RedundancyClassAPI struct{}

func NewRedundancyClassAPI() *RedundancyClassAPI {
	return &RedundancyClassAPI{}
}

func (api *RedundancyClassAPI) ListRedundancyClassesBySwarm(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	swarmID string,
) ([]RedundancyClass, error) {
	url := NewURLBuilder(endpoints.CH).
		Path("v5", "organizations", organizationID, "swarms", swarmID, "redundancy_class").
		Build()

	var response GenericPaginatedResponse[RedundancyClass]

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
