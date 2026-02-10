// Package api provides functions to interact with the project API.
package api

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"github.com/cubbit/composer-cli/src/configuration"
	"github.com/cubbit/composer-cli/src/request_utils"
)

func CreateProject(urls configuration.URLs, accessToken, name string, description *string, imageUrl *string) (*GenericIDResponseModel, error) {
	var err error
	var response GenericIDResponseModel

	url := NewURLBuilder(urls.IamURL).
		Path("v1", "projects").
		Build()

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
		request_utils.WithApiKey(accessToken),
	); err != nil {
		return nil, err
	}

	return &response, nil
}

func UpdateProject(urls configuration.URLs, accessToken, tenantID, projectID string, projectBody UpdateTenantProjectRequestBody) error {
	var err error

	url := NewURLBuilder(urls.ChURL).
		Path("v1", "tenants", tenantID, "projects", projectID).
		Build()

	requestBody, err := json.Marshal(projectBody)
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

func ListTenantProjects(urls configuration.URLs, accessToken, tenantID, sort, filter string) (*GenericPaginatedResponse[*ProjectItem], error) {
	var err error
	var finalResponse GenericPaginatedResponse[*ProjectItem]
	var nextPage *int

	url := NewURLBuilder(urls.ChURL).
		Path("v1", "tenants", tenantID, "projects").
		QueryParam("sort_key", sort).
		QueryParam("q", url.QueryEscape(filter)).
		Build()

	page := 1

	for {
		var response GenericPaginatedResponse[*ProjectItem]
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

func RemoveTenantProject(urls configuration.URLs, accessToken, tenantID, projectID string) error {
	var err error

	url := NewURLBuilder(urls.ChURL).
		Path("v1", "tenants", tenantID, "projects", projectID).
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

func ToggleBanProject(urls configuration.URLs, accessToken, tenantID string, projectID string, banned bool) error {
	var err error
	banURL := "ban"
	if !banned {
		banURL = "unban"
	}

	url := NewURLBuilder(urls.ChURL).
		Path("v1", "tenants", tenantID, "projects", projectID, banURL).
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

func RestoreTenantProject(urls configuration.URLs, accessToken, tenantID, projectID string) error {
	var err error

	url := NewURLBuilder(urls.ChURL).
		Path("v1", "tenants", tenantID, "projects", projectID, "restore").
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

func GetTenantProject(urls configuration.URLs, accessToken, tenantID, projectID string) (*ProjectItem, error) {
	var err error
	var response ProjectItem

	url := NewURLBuilder(urls.ChURL).
		Path("v1", "tenants", tenantID, "projects", projectID).
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
