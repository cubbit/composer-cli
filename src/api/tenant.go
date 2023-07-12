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

func ListTenant(apiServerUrl, accessToken, ownerID string) (*TenantList, error) {
	var err error
	url := apiServerUrl + "/v1/tenants?owner=" + ownerID
	var response TenantList

	if err = request_utils.DoRequest(
		url,
		request_utils.WithAccessToken(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		extractTenantListModel(&response),
	); err != nil {
		return nil, fmt.Errorf("failed unable to list tenants request: %w", err)
	}
	return &response, nil
}

func ListAvailableSwarmsTenant(apiServerUrl, accessToken, tenantID string) (*SwarmList, error) {
	var err error
	url := apiServerUrl + "/v1/tenants/" + tenantID + "/swarms"
	var response SwarmList

	if err = request_utils.DoRequest(
		url,
		request_utils.WithAccessToken(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		extractSwarmListModel(&response),
	); err != nil {
		return nil, fmt.Errorf("failed unable to list available swarms request: %w", err)
	}
	return &response, nil
}
