package api

import (
	"fmt"
	"net/http"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/src/request_utils"
)

func CreateSwarm(urls configuration.Url, accessToken, ownerID string, name string, description string, swarmConfig map[string]interface{}) (*GenericIDResponseModel, error) {
	var err error
	var response GenericIDResponseModel
	url := fmt.Sprintf("%s%s",urls.HiveUrl,constants.Swarms)

	requestBody := map[string]interface{}{
		"name":          name,
		"description":   description,
		"owner_id":      ownerID,
		"configuration": swarmConfig,
	}

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPost),
		request_utils.WithRequestBody(requestBody),
		request_utils.WithExpectedStatusCode(http.StatusCreated),
		extractGenericIDResponseModel(&response),
		request_utils.WithAccessToken(accessToken),
	); err != nil {
		return nil, fmt.Errorf("%s: %w", constants.ErrorCreatingSwarmRequest, err)
	}

	return &response, nil
}

func ListSwarms(urls configuration.Url, accessToken, ownerID string) ([]Swarm, error) {
	var err error
	url := fmt.Sprintf("%s%s",urls.HiveUrl,constants.Swarms)
	var response []Swarm

	if err = request_utils.DoRequest(
		url,
		request_utils.WithAccessToken(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		extractSwarmListResponseModel(&response),
	); err != nil {
		return nil, fmt.Errorf("%s: %w", constants.ErrorListingSwarmsRequest, err)
	}
	return response, nil
}

func GetSwarm(urls configuration.Url, accessToken, ownerID string, swarmID string) (*Swarm, error) {
	var err error
	url := fmt.Sprintf("%s%s/%s", urls.HiveUrl, constants.Swarms, swarmID)
	var response Swarm

	if err = request_utils.DoRequest(
		url,
		request_utils.WithAccessToken(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		extractSwarmResponseModel(&response),
	); err != nil {
		return nil, fmt.Errorf("%s: %w", constants.ErrorListingSwarmsRequest, err)
	}
	return &response, nil
}
