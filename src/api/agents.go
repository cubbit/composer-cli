package api

import (
	"encoding/json"
	"net/http"

	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/src/request_utils"
)

func CreateAgentV4(urls configuration.Url, accessToken string, swarmID string, nexusID string, nodeID string, nodesBody BulkInsertNewAgentRequestBody) (*NewAgentsResponse, error) {
	var err error
	var response NewAgentsResponse
	url := urls.ChUrl + "/v4/swarms/" + swarmID + "/nexuses/" + nexusID + "/nodes/" + nodeID + "/agents"

	bodyRequest, err := json.Marshal(nodesBody)
	if err != nil {
		return nil, err
	}

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPost),
		request_utils.WithAccessToken(accessToken),
		request_utils.WithRequestBodyByte(bodyRequest),
		request_utils.WithExpectedStatusCode(http.StatusCreated),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, err
	}

	return &response, nil
}
