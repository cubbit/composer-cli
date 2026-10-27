package api

import (
	"fmt"
	"net/http"

	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/src/request_utils"
)

type OperatorAPIInterface interface {
	Connect(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
	) (string, error)
}

type OperatorAPI struct{}

func NewOperatorAPI() *OperatorAPI {
	return &OperatorAPI{}

}

func (api *OperatorAPI) Connect(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
) (string, error) {

	url := NewURLBuilder(endpoints.CH).
		Path("v1", "organizations", organizationID, "infra", "generate-connect-cmd").
		Build()

	var command InfraClusterConnectCmdResponse
	if err := request_utils.DoRequest(
		url,
		request_utils.WithExpectedStatusCode(http.StatusOK),
		request_utils.WithApiKey(apiKey),
		ExtractGenericModel(&command),
	); err != nil {
		return "", fmt.Errorf("failed to perform connect request: %w", err)
	}

	return command.Command, nil
}
