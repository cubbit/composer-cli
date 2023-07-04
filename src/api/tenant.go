package api

import (
	"fmt"
	"net/http"

	"github.com/cubbit/cubbit/client/cli/src/request_utils"
)

func CreateTenant(apiServerUrl, accessToken, name string, description *string, imageUrl *string, settings map[string]interface{}) (*GenericIDResponseModel, error) {
	var err error
	var response GenericIDResponseModel
	url := apiServerUrl + "/v1/tenants"

	requestBody := map[string]interface{}{
		"name":     name,
		"settings": settings,
	}

	if description != nil {
		requestBody["description"] = *description
	}

	if imageUrl != nil {
		requestBody["image_url"] = *imageUrl
	}

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPost),
		request_utils.WithRequestBody(requestBody),
		request_utils.WithExpectedStatusCode(http.StatusCreated),
		extractGenericIDResponseModel(&response),
		request_utils.WithAccessToken(accessToken),
	); err != nil {
		return nil, fmt.Errorf("failed unable to create get operator request: %w", err)
	}

	return &response, nil
}
