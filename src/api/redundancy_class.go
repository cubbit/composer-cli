package api

import (
	"encoding/json"
	"net/http"
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

func ListRedundancyClasses(urls configuration.Url, accessToken, swamID string) (*RedundancyClassList, error) {
	var err error
	var finalResponse RedundancyClassList
	url := urls.ChUrl + constants.RedundancyClasses + "?swarm_id=" + swamID

	page := 0
	resultsPerPage := 1000

	for {
		var response RedundancyClassList
		if err = request_utils.DoRequest(
			url+"&start_page="+strconv.Itoa(page)+"&results_per_page="+strconv.Itoa(resultsPerPage),
			request_utils.WithAccessToken(accessToken),
			request_utils.WithExpectedStatusCode(http.StatusOK),
			ExtractGenericModel(&response),
		); err != nil {
			return nil, err
		}

		finalResponse.Data = append(finalResponse.Data, response.Data...)

		if len(response.Data) == 0 {
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
