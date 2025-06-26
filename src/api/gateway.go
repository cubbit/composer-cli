package api

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/src/request_utils"
)

func CreateGateway(urls configuration.Url, accessToken string, tenantID string, gatewayBody CreateGatewayRequestBody) (*GatewayWithGatewayTenant, error) {
	var err error
	var response GatewayWithGatewayTenant
	url := urls.ChUrl + constants.Tenants + "/" + tenantID + "/gateways"

	requestBody, err := json.Marshal(gatewayBody)
	if err != nil {
		return nil, err
	}

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPost),
		request_utils.WithRequestBodyByte(requestBody),
		request_utils.WithExpectedStatusCode(http.StatusCreated),
		request_utils.WithAccessToken(accessToken),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, err
	}

	return &response, nil
}

func UpdateGateway(urls configuration.Url, accessToken string, tenantID string, gatewayID string, gatewayBody UpdateGatewayRequestBody) error {
	var err error
	url := urls.ChUrl + constants.Tenants + "/" + tenantID + "/gateways/" + gatewayID

	requestBody, err := json.Marshal(gatewayBody)
	if err != nil {
		return err
	}

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPatch),
		request_utils.WithRequestBodyByte(requestBody),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		request_utils.WithAccessToken(accessToken),
	); err != nil {
		return err
	}

	return nil
}

func GetGateway(urls configuration.Url, accessToken string, tenantID string, gatewayID string) (*Gateway, error) {
	var err error
	var response Gateway
	url := urls.ChUrl + constants.Tenants + "/" + tenantID + "/gateways/" + gatewayID

	if err = request_utils.DoRequest(
		url,
		request_utils.WithExpectedStatusCode(http.StatusOK),
		request_utils.WithAccessToken(accessToken),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, err
	}

	return &response, nil
}

func ListGateways(urls configuration.Url, accessToken string, tenantID string, sort string, filter string) (*GenericPaginatedResponse[*Gateway], error) {
	var err error
	var finalResponse GenericPaginatedResponse[*Gateway]
	url := urls.ChUrl + constants.Tenants + "/" + tenantID + "/gateways" + "?sort_key=" + sort + "&q=" + url.QueryEscape(filter)

	var nextPage *int
	page := 1

	for {
		var response GenericPaginatedResponse[*Gateway]
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

func DeleteGateway(urls configuration.Url, accessToken string, tenantID string, gatewayID string, token string) error {
	var err error
	url := urls.ChUrl + constants.Tenants + "/" + tenantID + "/gateways/" + gatewayID + "?token=" + token

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodDelete),
		request_utils.WithExpectedStatusCode(http.StatusNoContent),
		request_utils.WithAccessToken(accessToken),
	); err != nil {
		return err
	}

	return nil
}

func ListGatewayInstances(urls configuration.Url, accessToken string, tenantID string, gatewayID string) (*GatewayInstanceListResponse, error) {
	var err error
	var finalResponse GatewayInstanceListResponse
	url := urls.ChUrl + constants.Tenants + "/" + tenantID + "/gateways/" + gatewayID + "/instances"

	if err = request_utils.DoRequest(
		url,
		request_utils.WithAccessToken(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		ExtractGenericModel(&finalResponse),
	); err != nil {
		return nil, err
	}

	return &finalResponse, nil
}

func VerifyDNS(urls configuration.Url, accessToken string, tenantID string) error {
	var err error
	url := urls.ChUrl + constants.Tenants + "/" + tenantID + "/whitelabel/verify-dns"

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPut),
		request_utils.WithExpectedStatusCode(http.StatusNoContent),
		request_utils.WithAccessToken(accessToken),
	); err != nil {
		return err
	}

	return nil
}
