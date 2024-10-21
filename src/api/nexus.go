package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/src/request_utils"
)

func CreateNexus(urls configuration.Url, accessToken string, swarmID string, nexusBody CreateNexusRequestBody) (*Nexus, error) {
	var err error
	var response Nexus
	url := urls.ChUrl + constants.Swarms + "/" + swarmID + "/nexuses"

	requestBody, err := json.Marshal(nexusBody)
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

func UpdateNexus(urls configuration.Url, accessToken string, nexusID string, nexusBody UpdateNexusRequestBody) error {
	var err error
	url := urls.ChUrl + constants.Nexuses + "/" + nexusID

	requestBody, err := json.Marshal(nexusBody)
	if err != nil {
		return err
	}

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPatch),
		request_utils.WithRequestBodyByte(requestBody),
		request_utils.WithExpectedStatusCode(http.StatusNoContent),
		request_utils.WithAccessToken(accessToken),
	); err != nil {
		return err
	}

	return nil
}

func GetNexus(urls configuration.Url, accessToken string, nexusID string) (*Nexus, error) {
	var err error
	var response Nexus
	url := urls.ChUrl + constants.Nexuses + "/" + nexusID

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

func ListNexuses(urls configuration.Url, accessToken string, swarmID string) (*NexusList, error) {
	var err error
	var finalResponse NexusList
	url := urls.ChUrl + constants.Nexuses + "?swarm_id=" + swarmID

	page := 0
	resultsPerPage := 1000

	for {
		var response NexusList
		if err = request_utils.DoRequest(
			url+"&start_page="+strconv.Itoa(page)+"&results_per_page="+strconv.Itoa(resultsPerPage),
			request_utils.WithAccessToken(accessToken),
			request_utils.WithExpectedStatusCode(http.StatusOK),
			ExtractGenericModel(&response),
		); err != nil {
			return nil, err
		}

		finalResponse.Nexuses = append(finalResponse.Nexuses, response.Nexuses...)
		finalResponse.Count = response.Count
		finalResponse.NextPage = response.NextPage

		if len(response.Nexuses) == 0 {
			break
		}
		page++
	}

	return &finalResponse, nil
}

func DeleteNexus(urls configuration.Url, accessToken string, nexusID string) error {
	var err error
	url := urls.ChUrl + constants.Nexuses + "/" + nexusID

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
