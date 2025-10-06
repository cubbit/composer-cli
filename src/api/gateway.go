// Package api provides functions to interact with the nexus API.
package api

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"github.com/cubbit/composer-cli/src/configuration"
	"github.com/cubbit/composer-cli/src/request_utils"
)

func CreateGateway(urls configuration.URLs, accessToken string, tenantID string, gatewayBody CreateGatewayRequestBody) (*GatewayWithGatewayTenant, error) {
	var err error
	var response GatewayWithGatewayTenant
	var requestBody []byte

	url := NewURLBuilder(urls.ChURL).
		Path("v1", "tenants", tenantID, "gateways").
		Build()

	requestBody, err = json.Marshal(gatewayBody)
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

func UpdateGateway(urls configuration.URLs, accessToken string, tenantID string, gatewayID string, gatewayBody UpdateGatewayRequestBody) error {
	var err error
	var requestBody []byte

	url := NewURLBuilder(urls.ChURL).
		Path("v1", "tenants", tenantID, "gateways", gatewayID).
		Build()

	requestBody, err = json.Marshal(gatewayBody)
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

func GetGateway(urls configuration.URLs, accessToken string, tenantID string, gatewayID string) (*Gateway, error) {
	var err error
	var response Gateway

	url := NewURLBuilder(urls.ChURL).
		Path("v1", "tenants", tenantID, "gateways", gatewayID).
		Build()

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

func ListGateways(urls configuration.URLs, accessToken string, tenantID string, sort string, filter string) (*GenericPaginatedResponse[*Gateway], error) {
	var err error
	var finalResponse GenericPaginatedResponse[*Gateway]
	var nextPage *int

	url := NewURLBuilder(urls.ChURL).
		Path("v1", "tenants", tenantID, "gateways").
		QueryParam("sort_key", sort).
		QueryParam("q", url.QueryEscape(filter)).
		Build()

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

func DeleteGateway(urls configuration.URLs, accessToken string, tenantID string, gatewayID string) error {
	var err error

	url := NewURLBuilder(urls.ChURL).
		Path("v1", "tenants", tenantID, "gateways", gatewayID).
		Build()

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

func VerifyDNS(urls configuration.URLs, accessToken string, tenantID string) error {
	var err error

	url := NewURLBuilder(urls.ChURL).
		Path("v1", "tenants", tenantID, "whitelabel", "verify-dns").
		Build()

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
