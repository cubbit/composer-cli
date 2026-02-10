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

func CreateNexus(urls configuration.URLs, accessToken string, swarmID string, nexusBody CreateNexusRequestBody) (*Nexus, error) {
	var err error
	var response Nexus
	var requestBody []byte

	url := NewURLBuilder(urls.ChURL).
		Path("v2", "swarms", swarmID, "nexuses").
		Build()

	requestBody, err = json.Marshal(nexusBody)
	if err != nil {
		return nil, err
	}

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPost),
		request_utils.WithRequestBodyByte(requestBody),
		request_utils.WithExpectedStatusCode(http.StatusCreated),
		request_utils.WithApiKey(accessToken),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, err
	}

	return &response, nil
}

func GetNexus(urls configuration.URLs, accessToken string, swarmID string, nexusID string) (*Nexus, error) {
	var err error
	var response Nexus

	url := NewURLBuilder(urls.ChURL).
		Path("v2", "swarms", swarmID, "nexuses", nexusID).
		Build()

	if err = request_utils.DoRequest(
		url,
		request_utils.WithExpectedStatusCode(http.StatusOK),
		request_utils.WithApiKey(accessToken),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, err
	}

	return &response, nil
}

func UpdateNexus(urls configuration.URLs, accessToken string, swarmID string, nexusID string, nexusBody UpdateNexusRequestBody) error {
	var err error
	var requestBody []byte

	url := NewURLBuilder(urls.ChURL).
		Path("v2", "swarms", swarmID, "nexuses", nexusID).
		Build()

	requestBody, err = json.Marshal(nexusBody)
	if err != nil {
		return err
	}

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPatch),
		request_utils.WithRequestBodyByte(requestBody),
		request_utils.WithExpectedStatusCode(http.StatusNoContent),
		request_utils.WithApiKey(accessToken),
	); err != nil {
		return err
	}

	return nil
}

func ListNexuses(urls configuration.URLs, accessToken string, swarmID string, sort string, filter string) (*GenericPaginatedResponse[*Nexus], error) {
	var err error
	var finalResponse GenericPaginatedResponse[*Nexus]
	var nextPage *int

	url := NewURLBuilder(urls.ChURL).
		Path("v2", "swarms", swarmID, "nexuses").
		QueryParam("sort_key", sort).
		QueryParam("q", url.QueryEscape(filter)).
		Build()

	page := 1

	for {
		var response GenericPaginatedResponse[*Nexus]
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

func DeleteNexus(urls configuration.URLs, accessToken string, swarmID string, nexusID string) error {
	var err error

	url := NewURLBuilder(urls.ChURL).
		Path("v2", "swarms", swarmID, "nexuses", nexusID).
		Build()

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodDelete),
		request_utils.WithExpectedStatusCode(http.StatusNoContent),
		request_utils.WithApiKey(accessToken),
	); err != nil {
		return err
	}

	return nil
}
