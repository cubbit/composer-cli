// Package api provides functions to interact with the auth API.
package api

import (

	"net/http"

	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/src/request_utils"
)

func RegisterDevice(urls configuration.URLs, uuid string) (*DeviceRegistrationResponse, error) {
	var err error
	var response DeviceRegistrationResponse

	url := NewURLBuilder(urls.IamURL).
		Path("v1", "devices", "register").
		Build()

	requestBody := map[string]interface{}{
		"device_id": uuid,
	}

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPost),
		request_utils.WithRequestBody(requestBody),
		request_utils.WithExpectedStatusCode(http.StatusCreated),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, err
	}

	return &response, nil
}

func GetDeviceAPIKey(urls configuration.URLs, deviceID string) (string, error) {
	var err error
	var response string

	url := NewURLBuilder(urls.IamURL).
		Path("v1", "devices", deviceID, "api-keys").
		Build()

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodGet),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		ExtractGenericModel(&response),
	); err != nil {
		return "", err
	}

	return response, nil
}
