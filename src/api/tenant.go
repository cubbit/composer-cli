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

func RemoveTenant(apiServerUrl, accessToken, tenantId, deleteTenantToken string) error {
	var err error
	url := apiServerUrl + "/v1/tenants/" + tenantId + "?token=" + deleteTenantToken

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodDelete),
		request_utils.WithAccessToken(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusNoContent),
	); err != nil {
		return fmt.Errorf("failed unable to delete tenant request: %w", err)
	}

	return nil
}

func EditTenantDescription(apiServerUrl, accessToken, tenantID string) (*TenantList, error) {
	var err error
	var response TenantList

	url := apiServerUrl + "/v1/tenants/" + tenantID

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPatch),
		request_utils.WithAccessToken(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		extractTenantListModel(&response),
	); err != nil {
		return nil, fmt.Errorf("failed unable to delete tenant request: %w", err)
	}

	return &response, nil
}

func EditTenantImage(apiServerUrl, accessToken, tenantID string) (*TenantList, error) {
	var err error
	var response TenantList

	url := apiServerUrl + "/v1/tenants/" + tenantID

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPatch),
		request_utils.WithAccessToken(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		extractTenantListModel(&response),
	); err != nil {
		return nil, fmt.Errorf("failed unable to delete tenant request: %w", err)
	}

	return &response, nil
}
