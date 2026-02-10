// Package api provides functions to interact with the account API.
package api

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"github.com/cubbit/composer-cli/src/configuration"
	"github.com/cubbit/composer-cli/src/request_utils"
)

func ListTenantAccounts(urls configuration.URLs, accessToken, tenantID, sort, filter string) (*GenericPaginatedResponse[*Account], error) {
	var err error
	var finalResponse GenericPaginatedResponse[*Account]
	var nextPage *int

	url := NewURLBuilder(urls.IamURL).Path("v1", "tenants", tenantID, "accounts").
		QueryParam("sort_key", sort).
		QueryParam("q", url.QueryEscape(filter)).
		Build()

	page := 1

	for {
		var response GenericPaginatedResponse[*Account]
		if err = request_utils.DoRequest(
			url+"&page="+strconv.Itoa(page),
			request_utils.WithApiKey(accessToken),
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

func RemoveTenantAccount(urls configuration.URLs, accessToken, tenantID, accountID string) error {
	var err error

	url := NewURLBuilder(urls.IamURL).Path("v1", "tenants", tenantID, "accounts", accountID).
		Build()

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodDelete),
		request_utils.WithApiKey(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusOK),
	); err != nil {
		return err
	}

	return nil
}

func ToggleBanAccount(urls configuration.URLs, accessToken, tenantID string, accountID string, banned bool) error {
	var err error

	banURL := "ban"
	if !banned {
		banURL = "unban"
	}

	url := NewURLBuilder(urls.IamURL).
		Path("v1", "tenants", tenantID, "accounts", accountID, banURL).
		Build()

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPatch),
		request_utils.WithApiKey(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusNoContent),
	); err != nil {
		return err
	}

	return nil
}

func RestoreTenantAccount(urls configuration.URLs, accessToken, tenantID, accountID string) error {
	var err error

	url := NewURLBuilder(urls.IamURL).
		Path("v1", "tenants", tenantID, "accounts", accountID, "restore").
		Build()

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPost),
		request_utils.WithApiKey(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusNoContent)); err != nil {
		return err
	}

	return nil
}

func DeleteTenantAccountSessions(urls configuration.URLs, accessToken, tenantID, accountID string) error {
	var err error

	url := NewURLBuilder(urls.IamURL).
		Path("v1", "tenants", tenantID, "accounts", accountID, "sessions").
		Build()

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodDelete),
		request_utils.WithApiKey(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusNoContent)); err != nil {
		return err
	}

	return nil
}

func CreateTenantAccounts(urls configuration.URLs, accessToken, tenantID string, emails []string) error {
	var err error
	var postBody []byte

	url := NewURLBuilder(urls.IamURL).
		Path("v1", "tenants", tenantID, "accounts").
		Build()

	requestBody := map[string]interface{}{
		"emails": emails,
	}

	postBody, err = json.Marshal(requestBody)
	if err != nil {
		return err
	}

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPost),
		request_utils.WithRequestBodyByte(postBody),
		request_utils.WithApiKey(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusCreated),
	); err != nil {
		return err
	}

	return nil
}

func GetTenantAccount(urls configuration.URLs, accessToken, tenantID, accountID string) (*Account, error) {
	var err error
	var response Account

	url := NewURLBuilder(urls.IamURL).
		Path("v1", "tenants", tenantID, "accounts", accountID).
		Build()

	if err = request_utils.DoRequest(
		url,
		request_utils.WithApiKey(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, err
	}

	return &response, nil
}

func UpdateAccount(urls configuration.URLs, accessToken, tenantID, accountID string, accountBody UpdateAccountRequest) error {
	var err error
	var requestBody []byte

	url := NewURLBuilder(urls.IamURL).
		Path("v1", "tenants", tenantID, "accounts", accountID).
		Build()

	requestBody, err = json.Marshal(accountBody)
	if err != nil {
		return err
	}

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPatch),
		request_utils.WithRequestBodyByte(requestBody),
		request_utils.WithApiKey(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusOK),
	); err != nil {
		return err
	}

	return nil
}
