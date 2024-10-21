package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/src/request_utils"
)

func ListRings(urls configuration.Url, accessToken string, swarmID string, redundancyClassID string) (*RingList, error) {

	var err error
	var finalResponse RingList

	url := urls.ChUrl + constants.Rings + "?swarm_id=" + swarmID

	if redundancyClassID != "" {
		url += "&redundancy_class_id=" + redundancyClassID
	}

	page := 0
	resultsPerPage := 1000

	for {
		var response RingList
		if err = request_utils.DoRequest(
			url+"&start_page="+strconv.Itoa(page)+"&results_per_page="+strconv.Itoa(resultsPerPage),
			request_utils.WithAccessToken(accessToken),
			request_utils.WithExpectedStatusCode(http.StatusOK),
			ExtractGenericModel(&response),
		); err != nil {
			return nil, err
		}

		finalResponse.Data = append(finalResponse.Data, response.Data...)
		finalResponse.Count = response.Count

		if len(response.Data) == 0 {
			break
		}
		page++
	}

	return &finalResponse, nil
}

func CreateRing(urls configuration.Url, accessToken, swarmID, redundancyClassID string, ringBulk RingBulk, dryRun bool) (*RingList, error) {
	var err error
	var response RingList
	url := urls.ChUrl + constants.Swarms + "/" + swarmID + "/rings" + "/bulks"

	if dryRun {
		url += "?dry_run=true"
	}

	bodyRequest, err := json.Marshal(ringBulk)
	if err != nil {
		return nil, err
	}

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPost),
		request_utils.WithRequestBodyByte(bodyRequest),
		request_utils.WithExpectedStatusCode(http.StatusCreated),
		request_utils.WithAccessToken(accessToken),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, err
	}

	return &response, nil
}
