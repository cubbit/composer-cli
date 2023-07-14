package api

import (
	"fmt"
	"net/http"

	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/src/request_utils"
)

func CreateTenant(urls configuration.Url, accessToken, name string, description *string, imageUrl *string, settings map[string]interface{}) (*GenericIDResponseModel, error) {
	var err error
	var response GenericIDResponseModel
	url := urls.IamUrl + "/v1/tenants"

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

func ListTenant(urls configuration.Url, accessToken, ownerID string) (*TenantList, error) {
	var err error
	url := urls.IamUrl + "/v1/tenants?owner=" + ownerID
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

func RemoveTenant(urls configuration.Url, accessToken, tenantId, deleteTenantToken string) error {
	var err error
	url := urls.IamUrl + "/v1/tenants/" + tenantId + "?token=" + deleteTenantToken

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

func EditTenantDescription(urls configuration.Url, accessToken, tenantID, description string) error {
	var err error

	requestBody := map[string]interface{}{
		"description": description,
	}

	url := urls.IamUrl + "/v1/tenants/" + tenantID

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPatch),
		request_utils.WithAccessToken(accessToken),
		request_utils.WithRequestBody(requestBody),
		request_utils.WithExpectedStatusCode(http.StatusCreated),
	); err != nil {
		return fmt.Errorf("failed unable to edit tenant description: %w", err)
	}

	return nil
}

func EditTenantImage(urls configuration.Url, accessToken, tenantID, imageUrl string) error {
	var err error

	requestBody := map[string]interface{}{
		"image_url": imageUrl,
	}

	url := urls.IamUrl + "/v1/tenants/" + tenantID

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPatch),
		request_utils.WithAccessToken(accessToken),
		request_utils.WithRequestBody(requestBody),
		request_utils.WithExpectedStatusCode(http.StatusCreated),
	); err != nil {
		return fmt.Errorf("failed unable to edit tenant image: %w", err)
	}

	return nil
}

func ListAvailableSwarmsTenant(urls configuration.Url, accessToken, tenantID string) (*SwarmList, error) {
	var err error
	url := urls.IamUrl + "/v1/tenants/" + tenantID + "/swarms"
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
