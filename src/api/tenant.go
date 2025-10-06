// Package api provides functions to interact with the tenant API.
package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cubbit/composer-cli/constants"
	"github.com/cubbit/composer-cli/src/configuration"
	"github.com/cubbit/composer-cli/src/request_utils"
)

func CreateTenant(urls configuration.URLs, accessToken string, req CreateTenantRequestBody) (*GenericIDResponseModel, error) {
	var err error
	var response GenericIDResponseModel
	var bodyRequest []byte

	url := NewURLBuilder(urls.ChURL).
		Path("v2", "tenants").
		Build()

	if bodyRequest, err = json.Marshal(req); err != nil {
		return nil, err
	}

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPost),
		request_utils.WithRequestBodyByte(bodyRequest),
		request_utils.WithExpectedStatusCode(http.StatusCreated),
		ExtractGenericModel(&response),
		request_utils.WithAccessToken(accessToken),
	); err != nil {
		return nil, err
	}

	return &response, nil
}

func ListTenants(urls configuration.URLs, accessToken, sort, filter string) (*GenericPaginatedResponse[*Tenant], error) {
	var err error
	var finalResponse GenericPaginatedResponse[*Tenant]

	url := NewURLBuilder(urls.ChURL).
		Path("v1", "tenants").
		QueryParam("sort_key", sort).
		QueryParam("q", filter).
		Build()

	var nextPage *int
	page := 1

	for {
		var response GenericPaginatedResponse[*Tenant]
		if err = request_utils.DoRequest(
			url+"&page="+strconv.Itoa(page),
			request_utils.WithAccessToken(accessToken),
			request_utils.WithExpectedStatusCode(http.StatusOK),
			ExtractGenericModel(&response),
		); err != nil {
			return nil, err
		}

		finalResponse.Data = append(finalResponse.Data, response.Data...)
		finalResponse.Count = response.Count
		finalResponse.NextPage = response.NextPage

		if nextPage = response.NextPage; nextPage == nil {
			break
		}
		page++
	}

	return &finalResponse, nil
}

func RemoveTenant(urls configuration.URLs, accessToken, tenantID string) error {
	var err error

	url := NewURLBuilder(urls.ChURL).
		Path("v1", "tenants", tenantID).
		Build()

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodDelete),
		request_utils.WithAccessToken(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusNoContent),
	); err != nil {
		return err
	}

	return nil
}

func EditTenant(urls configuration.URLs, accessToken string, tenantID string, req UpdateTenantRequestBody) error {
	var err error
	var bodyRequest []byte

	url := NewURLBuilder(urls.ChURL).
		Path("v2", "tenants", tenantID).
		Build()

	if bodyRequest, err = json.Marshal(req); err != nil {
		return err
	}

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPatch),
		request_utils.WithAccessToken(accessToken),
		request_utils.WithRequestBodyByte(bodyRequest),
		request_utils.WithExpectedStatusCode(http.StatusCreated),
	); err != nil {
		return err
	}

	return nil
}

func EditTenantImage(urls configuration.URLs, accessToken, tenantID, imageURL string) error {
	var err error

	url := NewURLBuilder(urls.ChURL).
		Path("v1", "tenants", tenantID).
		Build()

	requestBody := map[string]interface{}{
		"image_url": imageURL,
	}

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPatch),
		request_utils.WithAccessToken(accessToken),
		request_utils.WithRequestBody(requestBody),
		request_utils.WithExpectedStatusCode(http.StatusCreated),
	); err != nil {
		return err
	}

	return nil
}

func ListAvailableTenantSwarms(urls configuration.URLs, accessToken, tenantID string) (*TenantSwarmList, error) {
	var err error
	var response TenantSwarmList

	url := NewURLBuilder(urls.ChURL).
		Path("v1", "tenants", tenantID, "swarms").
		Build()

	if err = request_utils.DoRequest(
		url,
		request_utils.WithAccessToken(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, err
	}
	return &response, nil
}

func ListTenantPolicies(urls configuration.URLs, accessToken, tenantID string) (*PolicyList, error) {
	var err error
	var response PolicyList

	url := NewURLBuilder(urls.IamURL).
		Path("v1", "policies").
		QueryParam("tenant", tenantID).
		Build()

	if err = request_utils.DoRequest(
		url,
		request_utils.WithAccessToken(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, err
	}
	return &response, nil
}

func InviteOperatorToTenant(urls configuration.URLs, accessToken, tenantID string, req InviteOperatorRequestBody) error {
	var err error
	var bodyRequest []byte

	url := NewURLBuilder(urls.IamURL).
		Path("v1", "tenants", tenantID, "invites").
		Build()

	if bodyRequest, err = json.Marshal(req); err != nil {
		return err
	}

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPost),
		request_utils.WithRequestBodyByte(bodyRequest),
		request_utils.WithExpectedStatusCode(http.StatusCreated),
		request_utils.WithAccessToken(accessToken),
	); err != nil {
		return err
	}

	return nil
}

func ListTenantOperators(urls configuration.URLs, accessToken, tenantID string) (*OperatorList, error) {
	var err error
	var response OperatorList

	url := NewURLBuilder(urls.IamURL).
		Path("v1", "tenants", tenantID, "operators").
		Build()

	if err = request_utils.DoRequest(
		url,
		request_utils.WithAccessToken(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, err
	}
	return &response, nil
}

func GetTenantOperator(urls configuration.URLs, accessToken, tenantID, operatorID string) (*Operator, error) {
	var err error
	var response Operator

	url := NewURLBuilder(urls.IamURL).
		Path("v1", "tenants", tenantID, "operators", operatorID).
		Build()

	if err = request_utils.DoRequest(
		url,
		request_utils.WithAccessToken(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, err
	}
	return &response, nil
}

func RemoveTenantOperator(urls configuration.URLs, accessToken, tenantID, operatorID string) error {
	var err error

	url := NewURLBuilder(urls.IamURL).
		Path("v1", "tenants", tenantID, "operators", operatorID).
		Build()

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodDelete),
		request_utils.WithAccessToken(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusOK),
	); err != nil {
		return err
	}

	return nil
}

func GetTenant(urls configuration.URLs, accessToken, tenantID string) (*Tenant, error) {
	var err error
	var response Tenant

	url := NewURLBuilder(urls.ChURL).
		Path("v1", "tenants", tenantID).
		Build()

	if err = request_utils.DoRequest(
		url,
		request_utils.WithAccessToken(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, err
	}
	return &response, nil
}

func ConnectSwarm(urls configuration.URLs, accessToken, tenantID, swarmID string, redundancyClassID string, isDefault bool) error {
	var err error

	url := NewURLBuilder(urls.ChURL).
		Path("v2", "tenants", tenantID, "swarms", swarmID).
		Build()

	requestBody := map[string]interface{}{
		"redundancy_class_id": redundancyClassID,
		"default":             isDefault,
	}

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPut),
		request_utils.WithAccessToken(accessToken),
		request_utils.WithRequestBody(requestBody),
		request_utils.WithExpectedStatusCode(http.StatusCreated),
	); err != nil {
		return err
	}

	return nil
}

func GetTenantCouponSwarms(urls configuration.URLs, accessToken, tenantID string) (*SwarmList, error) {
	var err error
	var response SwarmList

	url := NewURLBuilder(urls.ChURL).
		Path("v1", "tenants", tenantID, "coupons", "default", "swarms").
		Build()

	if err = request_utils.DoRequest(
		url,
		request_utils.WithAccessToken(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, err
	}
	return &response, nil
}

func GetGatewayZones(urls configuration.URLs) (*ZoneMap, error) {
	var err error
	var response ZoneMap

	url := NewURLBuilder(urls.IamURL).
		Path("v1", "zones").
		Build()

	if err = request_utils.DoRequest(
		url,
		request_utils.WithExpectedStatusCode(http.StatusOK),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, err
	}
	return &response, nil
}

func AssignTenantToCoupon(urls configuration.URLs, accessToken, tenantID, CouponCode string) (*GenericIDResponseModel, error) {
	var err error
	var response GenericIDResponseModel

	url := NewURLBuilder(urls.ChURL).
		Path("v1", "tenants", tenantID, "coupons").
		Build()

	requestBody := map[string]interface{}{
		"coupon_code": CouponCode,
	}

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPatch),
		request_utils.WithAccessToken(accessToken),
		request_utils.WithRequestBody(requestBody),
		request_utils.WithExpectedStatusCode(http.StatusCreated),
	); err != nil {
		return nil, err
	}

	return &response, nil
}

func EditTenantSettings(urls configuration.URLs, accessToken string, tenantID string, settings TenantSettings) error {
	var err error

	url := NewURLBuilder(urls.ChURL).
		Path("v2", "tenants", tenantID).
		Build()

	requestBody := map[string]interface{}{
		"settings": settings,
	}

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPatch),
		request_utils.WithAccessToken(accessToken),
		request_utils.WithRequestBody(requestBody),
		request_utils.WithExpectedStatusCode(http.StatusCreated),
	); err != nil {
		return err
	}

	return nil
}

func EditOperatorRoleInTenant(urls configuration.URLs, accessToken, tenantID, operatorID string, req ChangeOperatorPolicyRequestBody) error {
	var err error
	var bodyRequest []byte

	url := NewURLBuilder(urls.IamURL).
		Path("v1", "tenants", tenantID, "operators", operatorID, "roles").
		Build()

	if bodyRequest, err = json.Marshal(req); err != nil {
		return err
	}

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPut),
		request_utils.WithRequestBodyByte(bodyRequest),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		request_utils.WithAccessToken(accessToken),
	); err != nil {
		return err
	}

	return nil
}

func DownloadTenantReport(urls configuration.URLs, accessToken, tenantID, from, to, output string) (*string, error) {
	var err error
	var response string

	url := NewURLBuilder(urls.DashURL).
		Path(constants.BaseDashURI, "v1", "tenants", tenantID, "projects", "report").
		QueryParam("from_date", from).
		QueryParam("to_date", to).Build()

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodGet),
		request_utils.WithAccessToken(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		request_utils.WithAttachement(),
		DownloadReport(output, &response),
	); err != nil {
		return nil, err
	}
	return &response, nil
}

func GetTenantReport(urls configuration.URLs, accessToken, tenantID, from, to string) (*TenantReportResponseModel, error) {
	var err error
	var response TenantReportResponseModel

	url := NewURLBuilder(urls.DashURL).
		Path(constants.BaseDashURI, "v1", "tenants", tenantID, "projects", "report").
		QueryParam("from_date", from).
		QueryParam("to_date", to).Build()

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodGet),
		request_utils.WithAccessToken(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		extractReport(&response),
	); err != nil {
		return nil, err
	}

	return &response, err
}
