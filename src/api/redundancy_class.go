// Package api provides functions to interact with the redundancy class API.
package api

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"github.com/cubbit/composer-cli/src/configuration"
	"github.com/cubbit/composer-cli/src/request_utils"
)

func ListRedundancyClasses(urls configuration.URLs, accessToken, swamID, sort, filter string) (*GenericPaginatedResponse[*RedundancyClass], error) {
	var err error
	var finalResponse GenericPaginatedResponse[*RedundancyClass]
	var nextPage *int

	url := NewURLBuilder(urls.ChURL).
		Path("v2", "swarms", swamID, "redundancy_class").
		QueryParam("sort_key", sort).
		QueryParam("q", url.QueryEscape(filter)).
		Build()

	page := 1

	for {
		var response GenericPaginatedResponse[*RedundancyClass]
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

func GetRedundancyClass(urls configuration.URLs, accessToken, redundancyClassID string) (*RedundancyClass, error) {
	var err error
	var response RedundancyClass

	url := NewURLBuilder(urls.ChURL).
		Path("v1", "redundancy_class", redundancyClassID).
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

func CreateRedundancyClass(urls configuration.URLs, accessToken, swamID string, redundancyClass CreateRedundancyClassRequestBody) (*RedundancyClass, error) {
	var err error
	var response RedundancyClass

	url := NewURLBuilder(urls.ChURL).
		Path("v4", "swarms", swamID, "redundancy_class").
		Build()

	requestBody, err := json.Marshal(redundancyClass)
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

func CheckRedundancyClassStatus(urls configuration.URLs, accessToken, swarmID, redundancyClassID string) (*SummaryDetailsWithStatusNullable, error) {
	var err error
	var response SummaryDetailsWithStatusNullable

	url := NewURLBuilder(urls.ChURL).
		Path("v2", "swarms", swarmID, "redundancy_class", redundancyClassID, "status").
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

func CheckRedundancyClassRecoveryStatus(urls configuration.URLs, accessToken, swarmID, redundancyClassID string) (*RCProgress, error) {
	var err error
	var response RCProgress

	url := NewURLBuilder(urls.ChURL).
		Path("v4", "swarms", swarmID, "redundancy_class", redundancyClassID, "recovery", "status").
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

func RecoverRedundancyClass(urls configuration.URLs, accessToken, swarmID, redundancyClassID string, dryRun bool) (*RedundancyClassRecovery, error) {
	var err error
	var response RedundancyClassRecovery

	url := NewURLBuilder(urls.ChURL).
		Path("v4", "swarms", swarmID, "redundancy_class", redundancyClassID, "recovery").
		QueryParam("dry_run", strconv.FormatBool(dryRun)).
		Build()

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPost),
		request_utils.WithAccessToken(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, err
	}

	return &response, nil
}

func ExpandRedundancyClass(urls configuration.URLs, accessToken, swarmID, redundancyClassID string, dryRun bool) (*RedundancyClassExpanded, error) {
	var err error
	var response RedundancyClassExpanded

	url := NewURLBuilder(urls.ChURL).
		Path("v4", "swarms", swarmID, "redundancy_class", redundancyClassID, "expand").
		QueryParam("dry_run", strconv.FormatBool(dryRun)).
		Build()

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPost),
		request_utils.WithRequestBodyByte([]byte(`{}`)),
		request_utils.WithAccessToken(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusCreated),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, err
	}

	return &response, nil
}
