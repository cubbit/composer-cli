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

func CreateProject(urls configuration.Url, accessToken, name string, description *string, imageUrl *string) (*GenericIDResponseModel, error) {
	var err error
	var response GenericIDResponseModel
	url := urls.IamUrl + constants.Projects

	requestBody := map[string]interface{}{
		"name": name,
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
		request_utils.WithExpectedStatusCode(http.StatusOK),
		ExtractGenericModel(&response),
		request_utils.WithAccessToken(accessToken),
	); err != nil {
		return nil, err
	}

	return &response, nil
}

func ListTenantProjects(urls configuration.Url, accessToken, tenantID, sort, filter string) (*GenericPaginatedResponse[*ProjectItem], error) {
	var err error
	url := urls.IamUrl + constants.Tenants + "/" + tenantID + "/projects" + "?sort_key=" + sort + "&q=" + url.QueryEscape(filter)
	var finalResponse GenericPaginatedResponse[*ProjectItem]

	var nextPage *int
	page := 1

	for {
		var response GenericPaginatedResponse[*ProjectItem]
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

func RemoveTenantProject(urls configuration.Url, accessToken, tenantID, projectID, deleteTenantProjectToken string) error {
	var err error

	url := urls.IamUrl + constants.Tenants + "/" + tenantID + "/projects/" + projectID + "?token=" + deleteTenantProjectToken
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

func ToggleBanProject(urls configuration.Url, accessToken, tenantID string, projectID string, banned bool) error {
	var err error
	banUrl := "ban"
	if !banned {
		banUrl = "unban"
	}
	url := urls.IamUrl + constants.Tenants + "/" + tenantID + "/projects/" + projectID + "/" + banUrl
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

func RestoreTenantProject(urls configuration.Url, accessToken, tenantID, projectID string) error {
	var err error

	url := urls.IamUrl + constants.Tenants + "/" + tenantID + "/projects/" + projectID + "/restore"

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPost),
		request_utils.WithAccessToken(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusNoContent)); err != nil {
		return err
	}

	return nil
}

func UpdateProject(urls configuration.Url, accessToken, tenantID, projectID string, projectBody UpdateTenantProjectRequestBody) error {
	var err error
	url := urls.IamUrl + constants.Tenants + "/" + tenantID + "/projects/" + projectID

	requestBody, err := json.Marshal(projectBody)
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

func GetTenantProject(urls configuration.Url, accessToken, tenantID, projectID string) (*ProjectItem, error) {
	var err error
	var response ProjectItem

	url := urls.IamUrl + constants.Tenants + "/" + tenantID + "/projects/" + projectID
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
