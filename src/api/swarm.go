package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cubbit/cubbit/client/cli/src/request_utils"
)

func CreateSwarm(accessToken string, ownerID string, name string, description *string, configuration map[string]interface{}) (*GenericIDResponseModel, error) {
	var err error
	url := "https://9dc4-62-152-126-198.ngrok-free.app" + "/v1/swarms" //sistemare apiServerUrl
	var response GenericIDResponseModel
	requestBody := map[string]interface{}{
		"owner_id":      ownerID,
		"name":          name,
		"configuration": configuration,
	}

	if description != nil {
		requestBody["description"] = *description
	}

	putBody, _ := json.Marshal(requestBody)

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPost),
		request_utils.WithRequestBodyByte(putBody),
		request_utils.WithAccessToken(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusCreated),
		extractGenericIDResponseModel(&response),
	); err != nil {
		return nil, fmt.Errorf("failed unable to create tenant request: %w", err)
	}
	return &response, nil
}
