package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/src/request_utils"
)

func CreateTenant(urls configuration.Url, accessToken, name string, description *string, imageUrl *string, settings map[string]interface{}, couponCode string) (*GenericIDResponseModel, error) {
	var err error
	var response GenericIDResponseModel
	url := urls.IamUrl + constants.Tenants

	requestBody := map[string]interface{}{
		"name":        name,
		"settings":    settings,
		"coupon_code": strings.ToUpper(couponCode),
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
		return nil, fmt.Errorf("%s: %w", constants.ErrorGettingOperatorRequest, err)
	}

	return &response, nil
}

func ListTenants(urls configuration.Url, accessToken string) (*TenantList, error) {
	var err error
	url := urls.IamUrl + constants.Tenants
	var response TenantList

	if err = request_utils.DoRequest(
		url,
		request_utils.WithAccessToken(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		extractTenantListModel(&response),
	); err != nil {
		return nil, fmt.Errorf("%s: %w", constants.ErrorListingTenantsRequest, err)
	}
	return &response, nil
}

func RemoveTenant(urls configuration.Url, accessToken, tenantId, deleteTenantToken string) error {
	var err error
	url := urls.IamUrl + constants.Tenants + "/" + tenantId + "?token=" + deleteTenantToken
	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodDelete),
		request_utils.WithAccessToken(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusNoContent),
	); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDeletingTenantRequest, err)
	}

	return nil
}

func EditTenantDescription(urls configuration.Url, accessToken, tenantID, description string) error {
	var err error

	requestBody := map[string]interface{}{
		"description": description,
	}

	url := urls.IamUrl + constants.Tenants + "/" + tenantID

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPatch),
		request_utils.WithAccessToken(accessToken),
		request_utils.WithRequestBody(requestBody),
		request_utils.WithExpectedStatusCode(http.StatusCreated),
	); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorEditingTenantRequest, err)
	}

	return nil
}

func EditTenantImage(urls configuration.Url, accessToken, tenantID, imageUrl string) error {
	var err error

	requestBody := map[string]interface{}{
		"image_url": imageUrl,
	}

	url := urls.IamUrl + constants.Tenants + "/" + tenantID

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPatch),
		request_utils.WithAccessToken(accessToken),
		request_utils.WithRequestBody(requestBody),
		request_utils.WithExpectedStatusCode(http.StatusCreated),
	); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorEditingTenantRequest, err)
	}

	return nil
}

func ListAvailableTenantSwarms(urls configuration.Url, accessToken, tenantID string) (*SwarmList, error) {
	var err error
	url := urls.IamUrl + constants.Tenants + "/" + tenantID + "/swarms"
	var response SwarmList

	if err = request_utils.DoRequest(
		url,
		request_utils.WithAccessToken(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		extractSwarmListModel(&response),
	); err != nil {
		return nil, fmt.Errorf("%s: %w", constants.ErrorListingTenantSwarmsRequest, err)
	}
	return &response, nil
}

func ListTenantPolicies(urls configuration.Url, accessToken, tenantID string) (*PolicyList, error) {
	var err error
	url := urls.IamUrl + constants.Policies + "?tenant=" + tenantID
	var response PolicyList

	if err = request_utils.DoRequest(
		url,
		request_utils.WithAccessToken(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		extractPolicyListModel(&response),
	); err != nil {
		return nil, fmt.Errorf("%s: %w", constants.ErrorListingTenantSwarmsRequest, err)
	}
	return &response, nil
}

func InviteOperatorToTenant(urls configuration.Url, accessToken, tenantID, email, role, firstName, lastName string) error {
	var err error
	url := urls.IamUrl + constants.Tenants + "/" + tenantID + constants.Invites

	requestBody := map[string]interface{}{
		"email":      email,
		"policy_id":  role,
		"first_name": firstName,
		"last_name":  lastName,
	}

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPost),
		request_utils.WithRequestBody(requestBody),
		request_utils.WithExpectedStatusCode(http.StatusCreated),
		request_utils.WithAccessToken(accessToken),
	); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorInvitingOperatorRequest, err)
	}

	return nil
}

func ListTenantOperators(urls configuration.Url, accessToken, tenantID string) (*OperatorList, error) {
	var err error
	url := urls.IamUrl + constants.Tenants + "/" + tenantID + "/operators"
	var response OperatorList

	if err = request_utils.DoRequest(
		url,
		request_utils.WithAccessToken(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		extractOperatorListModel(&response),
	); err != nil {
		return nil, fmt.Errorf("%s: %w", constants.ErrorListingTenantsRequest, err)
	}
	return &response, nil
}

func RemoveTenantOperator(urls configuration.Url, accessToken, tenantID, operatorID string) error {
	var err error
	url := urls.IamUrl + constants.Tenants + "/" + tenantID + "/operators/" + operatorID
	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodDelete),
		request_utils.WithAccessToken(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusOK),
	); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDeletingTenantRequest, err)
	}

	return nil
}

func GetTenant(urls configuration.Url, accessToken, ownerID string, tenantID string) (*Tenant, error) {
	var err error
	url := urls.IamUrl + constants.Tenants + "/" + tenantID
	var response Tenant

	if err = request_utils.DoRequest(
		url,
		request_utils.WithAccessToken(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		extractTenantResponseModel(&response),
	); err != nil {
		return nil, fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarmsRequest, err)
	}
	return &response, nil
}

func ConnectSwarm(urls configuration.Url, accessToken, tenantID, swarmID string) error {
	var err error

	url := urls.IamUrl + constants.Tenants + "/" + tenantID + "/swarms/" + swarmID

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPut),
		request_utils.WithAccessToken(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusCreated),
	); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorConnectingSwarmRequest, err)
	}

	return nil
}
