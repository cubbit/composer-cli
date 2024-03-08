package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/src/request_utils"
)

func ListTenantAccounts(urls configuration.Url, accessToken, tenantID, sort string) (*GenericPaginatedResponse[*Account], error) {
	var err error
	url := urls.IamUrl + constants.Tenants + "/" + tenantID + "/accounts" + "?sort_key=" + sort
	var finalResponse GenericPaginatedResponse[*Account]

	var nextPage *int
	page := 1

	for {
		var response GenericPaginatedResponse[*Account]
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

func RemoveTenantAccount(urls configuration.Url, accessToken, tenantID, accountID, deleteTenantAccountToken string) error {
	var err error
	url := urls.IamUrl + constants.Tenants + "/" + tenantID + "/accounts/" + accountID + "?token=" + deleteTenantAccountToken
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

func ToggleBanAccount(urls configuration.Url, accessToken, tenantID string, accountID string, banned bool) error {
	var err error
	banUrl := "ban"
	if !banned {
		banUrl = "unban"
	}
	url := urls.IamUrl + constants.Tenants + "/" + tenantID + "/accounts/" + accountID + "/" + banUrl
	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPatch),
		request_utils.WithAccessToken(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusNoContent),
	); err != nil {
		return err
	}

	return nil
}

func RestoreTenantAccount(urls configuration.Url, accessToken, tenantID, accountID string) error {
	var err error

	url := urls.IamUrl + constants.Tenants + "/" + tenantID + "/accounts/" + accountID + "/restore"

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPost),
		request_utils.WithAccessToken(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusNoContent)); err != nil {
		return err
	}

	return nil
}

func DeleteTenantAccountSessions(urls configuration.Url, accessToken, tenantID, accountID string) error {
	var err error

	url := urls.IamUrl + constants.Tenants + "/" + tenantID + "/accounts/" + accountID + "/sessions"

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodDelete),
		request_utils.WithAccessToken(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusNoContent)); err != nil {
		return err
	}

	return nil
}

func CreateTenantAccounts(urls configuration.Url, accessToken, tenantID string, emails []string) error {
	var err error

	url := urls.IamUrl + constants.Tenants + "/" + tenantID + "/accounts"

	requestBody := map[string]interface{}{
		"emails": emails,
	}

	postBody, _ := json.Marshal(requestBody)

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPost),
		request_utils.WithRequestBodyByte(postBody),
		request_utils.WithAccessToken(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusCreated),
	); err != nil {
		return err
	}

	return nil
}

func UpdateAccount(urls configuration.Url, accessToken, tenantID, accountID string, accountBody UpdateAccountRequest) error {
	var err error
	url := urls.IamUrl + constants.Tenants + "/" + tenantID + "/accounts/" + accountID

	requestBody, err := json.Marshal(accountBody)
	if err != nil {
		return err
	}

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPatch),
		request_utils.WithRequestBodyByte(requestBody),
		request_utils.WithAccessToken(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusOK),
	); err != nil {
		return err
	}

	return nil
}
