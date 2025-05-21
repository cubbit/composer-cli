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

func CreateRedundancyClass(urls configuration.Url, accessToken, swamID string, redundancyClass CreateRedundancyClassRequestBody) (*RedundancyClass, error) {
	var err error
	var response RedundancyClass
	url := urls.ChUrl + constants.Swarms + "/" + swamID + "/redundancy_class"

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

func ListRedundancyClasses(urls configuration.Url, accessToken, swamID, sort, filter string) (*GenericPaginatedResponse[*RedundancyClass], error) {
	var err error
	var finalResponse GenericPaginatedResponse[*RedundancyClass]

	url := urls.ChUrl + "/v2/swarms/" + swamID + "/redundancy_class" + "?sort_key=" + sort + "&q=" + url.QueryEscape(filter)

	var nextPage *int
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

func GetRedundancyClass(urls configuration.Url, accessToken, redundancyClassID string) (*RedundancyClass, error) {
	var err error
	var response RedundancyClass
	url := urls.ChUrl + constants.RedundancyClasses + "/" + redundancyClassID

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

func CreateRedundancyClassV4(urls configuration.Url, accessToken, swamID string, redundancyClass CreateRedundancyClassRequestBody) (*RedundancyClass, error) {
	var err error
	var response RedundancyClass
	url := urls.ChUrl + constants.Swarms + "/" + swamID + "/redundancy_class"

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

func CheckRedundancyClassStatus(urls configuration.Url, accessToken, swarmID, redundancyClassID string) (*SummaryDetailsWithStatusNullable, error) {
	var err error
	var response SummaryDetailsWithStatusNullable
	url := urls.ChUrl + "/v2/swarms/" + swarmID + "/redundancy_class/" + redundancyClassID + "/status"

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

func CheckRedundancyClassRecoveryStatus(urls configuration.Url, accessToken, swarmID, redundancyClassID string) (*RedundancyClassRecoveryStatus, error) {
	var err error
	var response RedundancyClassRecoveryStatus

	url := urls.ChUrl + "/v4/swarms/" + swarmID + "/redundancy_class/" + redundancyClassID + "/recovery/status"

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

func RecoverRedundancyClass(urls configuration.Url, accessToken, swarmID, redundancyClassID string, dryRun bool) (*RedundancyClassRecovery, error) {
	var err error
	var response RedundancyClassRecovery
	url := urls.ChUrl + "/v4/swarms/" + swarmID + "/redundancy_class/" + redundancyClassID + "/recovery" + "?dry_run=" + strconv.FormatBool(dryRun)

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

func ExpandRedundancyClass(urls configuration.Url, accessToken, swarmID, redundancyClassID string, dryRun bool) (*RedundancyClassExpanded, error) {
	var err error
	var response RedundancyClassExpanded
	url := urls.ChUrl + "/v4/swarms/" + swarmID + "/redundancy_class/" + redundancyClassID + "/expand" + "?dry_run=" + strconv.FormatBool(dryRun)

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
