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
	url := fmt.Sprintf("%s%s", urls.HiveUrl, constants.Swarms)

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
	url := fmt.Sprintf("%s%s", urls.HiveUrl, constants.Swarms)
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
		return nil, fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarmsRequest, err)
	}
	return &response, nil
}
func EditSwarmDescription(urls configuration.Url, accessToken, ownerID string, swarmID string, description string) error {
	var err error
	url := fmt.Sprintf("%s%s/%s", urls.HiveUrl, constants.Swarms, swarmID)

	requestBody := map[string]interface{}{
		"description": description,
	}

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPut),
		request_utils.WithRequestBody(requestBody),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		request_utils.WithAccessToken(accessToken),
	); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorEditingSwarmRequest, err)
	}

	return nil
}

func EditSwarmName(urls configuration.Url, accessToken, ownerID string, swarmID string, name string) error {
	var err error
	url := fmt.Sprintf("%s%s/%s", urls.HiveUrl, constants.Swarms, swarmID)

	requestBody := map[string]interface{}{
		"name": name,
	}

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPut),
		request_utils.WithRequestBody(requestBody),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		request_utils.WithAccessToken(accessToken),
	); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorEditingSwarmRequest, err)
	}

	return nil
}

func ListSwarmOperators(urls configuration.Url, accessToken, swarmID string) (*OperatorList, error) {
	var err error
	url := fmt.Sprintf("%s%s/%s/operators", urls.HiveUrl, constants.Swarms, swarmID)
	var response *OperatorList

	if err = request_utils.DoRequest(
		url,
		request_utils.WithAccessToken(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		extractOperatorListResponseModel(response),
	); err != nil {
		return nil, fmt.Errorf("%s: %w", constants.ErrorListingSwarmsRequest, err)
	}
	return response, nil
}
