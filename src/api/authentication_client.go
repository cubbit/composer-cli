package api

import (
	"net/http"

	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/src/request_utils"
)

func RegisterDevice(endpoints configuration_models.EndpointsV2, uuid string) (*DeviceRegistrationResponse, error) {
	var err error
	var response DeviceRegistrationResponse

	url := NewURLBuilder(endpoints.IAM).
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

func GetDeviceAPIKey(endpoints configuration_models.EndpointsV2, deviceID string) (string, error) {
	var err error
	var response string

	url := NewURLBuilder(endpoints.IAM).
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
