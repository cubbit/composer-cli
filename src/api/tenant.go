package api

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/src/request_utils"
)

func CreateTenant(urls configuration.Url, accessToken, name string, description *string, settings TenantSettings, couponCode, zone string) (*GenericIDResponseModel, error) {
	var err error
	var response GenericIDResponseModel
	url := urls.ChUrl + constants.TenantsV2

	requestBody := map[string]interface{}{
		"name":        name,
		"settings":    settings,
		"coupon_code": strings.ToUpper(couponCode),
	}

	if description != nil {
		requestBody["description"] = *description
	}

	if zone != "" {
		requestBody["zone"] = zone
	}

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPost),
		request_utils.WithRequestBody(requestBody),
		request_utils.WithExpectedStatusCode(http.StatusCreated),
		ExtractGenericModel(&response),
		request_utils.WithAccessToken(accessToken),
	); err != nil {
		return nil, err
	}

	return &response, nil
}

func ListTenants(urls configuration.Url, accessToken, sort, filter string) (*GenericPaginatedResponse[*Tenant], error) {
	var err error
	var finalResponse GenericPaginatedResponse[*Tenant]

	url := urls.ChUrl + constants.Tenants + "?sort_key=" + sort + "&q=" + url.QueryEscape(filter)

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

func RemoveTenant(urls configuration.Url, accessToken, tenantId, deleteTenantToken string) error {
	var err error
	url := urls.ChUrl + constants.Tenants + "/" + tenantId + "?token=" + deleteTenantToken
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

func EditTenantDescription(urls configuration.Url, accessToken, tenantID, description string) error {
	var err error

	requestBody := map[string]interface{}{
		"description": description,
	}

	url := urls.ChUrl + constants.Tenants + "/" + tenantID

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

func EditTenantImage(urls configuration.Url, accessToken, tenantID, imageUrl string) error {
	var err error

	requestBody := map[string]interface{}{
		"image_url": imageUrl,
	}

	url := urls.ChUrl + constants.Tenants + "/" + tenantID

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

func ListAvailableTenantSwarms(urls configuration.Url, accessToken, tenantID string) (*TenantSwarmList, error) {
	var err error
	url := urls.ChUrl + constants.Tenants + "/" + tenantID + "/swarms"
	var response TenantSwarmList

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

func ListTenantPolicies(urls configuration.Url, accessToken, tenantID string) (*PolicyList, error) {
	var err error
	url := urls.IamUrl + constants.Policies + "?tenant=" + tenantID
	var response PolicyList

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
		return err
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
		ExtractGenericModel(&response),
	); err != nil {
		return nil, err
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
		return err
	}

	return nil
}

func GetTenant(urls configuration.Url, accessToken, tenantID string) (*Tenant, error) {
	var err error
	url := urls.ChUrl + constants.Tenants + "/" + tenantID
	var response Tenant

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

func ConnectSwarm(urls configuration.Url, accessToken, tenantID, swarmID string) error {
	var err error

	url := urls.ChUrl + constants.Tenants + "/" + tenantID + "/swarms/" + swarmID

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPut),
		request_utils.WithAccessToken(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusCreated),
	); err != nil {
		return err
	}

	return nil
}

func GetTenantCouponSwarms(urls configuration.Url, accessToken, tenantID string) (*SwarmList, error) {
	var err error

	url := urls.ChUrl + constants.Tenants + "/" + tenantID + "/coupons/default/swarms"

	var response SwarmList

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

func GetGatwayZones(urls configuration.Url) (*ZoneMap, error) {
	var err error

	url := urls.IamUrl + constants.Zones

	var response ZoneMap

	if err = request_utils.DoRequest(
		url,
		request_utils.WithExpectedStatusCode(http.StatusOK),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, err
	}
	return &response, nil
}

func AssignTenantToCoupon(urls configuration.Url, accessToken, tenantID, CouponCode string) (*GenericIDResponseModel, error) {
	var err error
	var response GenericIDResponseModel
	url := urls.ChUrl + constants.Tenants + "/" + tenantID + "/coupon"

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

func EditTenantSettings(urls configuration.Url, accessToken string, tenantID string, settings TenantSettings) error {
	var err error

	requestBody := map[string]interface{}{
		"settings": settings,
	}

	url := urls.ChUrl + constants.TenantsV2 + "/" + tenantID

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

func EditOperatorRoleInTenant(urls configuration.Url, accessToken, tenantID, operatorID, role string) error {
	var err error
	url := urls.IamUrl + constants.Tenants + "/" + tenantID + "/operators/" + operatorID + "/roles"

	requestBody := map[string]interface{}{
		"policy_id": role,
	}

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPut),
		request_utils.WithRequestBody(requestBody),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		request_utils.WithAccessToken(accessToken),
	); err != nil {
		return err
	}

	return nil
}

func DownloadTenantReport(urls configuration.Url, accessToken, tenantID, from, to, output string) (*string, error) {
	var err error
	var url string
	var response string
	url = urls.DashUrl + constants.BaseDashURI + constants.Tenants + "/" + tenantID + "/projects/report" + "?from_date=" + from + "&to_date=" + to

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

func GetTenantReport(urls configuration.Url, accessToken, tenantID, from, to string) (*TenantReportResponseModel, error) {
	var err error
	var url string
	var response TenantReportResponseModel

	url = urls.DashUrl + constants.BaseDashURI + constants.Tenants + "/" + tenantID + "/projects/report" + "?from_date=" + from + "&to_date=" + to

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
