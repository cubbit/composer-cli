package api

import (
	"fmt"
	"net/http"

	"github.com/cubbit/composer-cli/src/configuration"
	"github.com/cubbit/composer-cli/src/request_utils"
)

type OperatorAPIInterface interface {
	Connect(
		urlConfig configuration.URLs,
		apiKey string,
		organizationID string,
	) (string, error)
}

type OperatorAPI struct {
	config configuration.Config
}

func NewOperatorAPI(config *configuration.Config) *OperatorAPI {
	return &OperatorAPI{
		config: *config,
	}
}

func (api *OperatorAPI) Connect(
	urlConfig configuration.URLs,
	apiKey string,
	organizationID string,
) (string, error) {

	url := NewURLBuilder(urlConfig.ChURL).
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
