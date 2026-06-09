package api

import (
	"net/http"

	"github.com/cubbit/composer-cli/src/configuration"
	"github.com/cubbit/composer-cli/src/request_utils"
)

type ProcessAPIInterface interface {
	ListProcesses(
		urlConfig configuration.URLs,
		apiKey string,
		organizationID string,
		processType ProcessType,
	) ([]Process, error)
	GetProcess(
		urlConfig configuration.URLs,
		apiKey string,
		organizationID string,
		processID string,
	) (*Process, error)
}

type ProcessAPI struct{}

func NewProcessAPI() *ProcessAPI {
	return &ProcessAPI{}
}

func (api *ProcessAPI) ListProcesses(
	urlConfig configuration.URLs,
	apiKey string,
	organizationID string,
	processType ProcessType,
) ([]Process, error) {
	url := NewURLBuilder(urlConfig.ChURL).
		Path("v1", "organizations", organizationID, "process").
		QueryParam("type", string(processType)).
		Build()

	var response GenericPaginatedResponse[Process]

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

func (api *ProcessAPI) GetProcess(
	urlConfig configuration.URLs,
	apiKey string,
	organizationID string,
	processID string,
) (*Process, error) {
	url := NewURLBuilder(urlConfig.ChURL).
		Path("v1", "organizations", organizationID, "process", processID).
		Build()

	var response Process

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
