package api

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/src/request_utils"
)

func CreateAgentV4(urls configuration.Url, accessToken string, swarmID string, nexusID string, nodeID string, agentsBody BulkInsertNewAgentRequestBody) (*NewAgentsResponse, error) {
	var err error
	var response NewAgentsResponse
	url := urls.ChUrl + "/v4/swarms/" + swarmID + "/nexuses/" + nexusID + "/nodes/" + nodeID + "/agents"

	bodyRequest, err := json.Marshal(agentsBody)
	if err != nil {
		return nil, err
	}

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPost),
		request_utils.WithAccessToken(accessToken),
		request_utils.WithRequestBodyByte(bodyRequest),
		request_utils.WithExpectedStatusCode(http.StatusCreated),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, err
	}

	return &response, nil
}

func ListAgentsV4(urls configuration.Url, accessToken string, swarmID string, nexusID, nodeID, sort string, filter string) (*GenericPaginatedResponse[*NewAgent], error) {
	var err error
	var finalResponse GenericPaginatedResponse[*NewAgent]

	url := urls.ChUrl + "/v4/swarms/" + swarmID + "/nexuses/" + nexusID + "/nodes/" + nodeID + "/agents" + "?sort_key=" + sort + "&q=" + url.QueryEscape(filter)

	var nextPage *int
	page := 1

	for {
		var response GenericPaginatedResponse[*NewAgent]
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

func UpdateAgentV4(urls configuration.Url, accessToken string, swarmID string, nexusID string, nodeID string, agentID string, nodesBody UpdateNewAgentRequestBody) error {
	var err error
	url := urls.ChUrl + "/v4/swarms/" + swarmID + "/nexuses/" + nexusID + "/nodes/" + nodeID + "/agents/" + agentID

	bodyRequest, err := json.Marshal(nodesBody)
	if err != nil {
		return err
	}

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPatch),
		request_utils.WithAccessToken(accessToken),
		request_utils.WithRequestBodyByte(bodyRequest),
		request_utils.WithExpectedStatusCode(http.StatusNoContent),
		ExtractGenericModel(nil),
	); err != nil {
		return err
	}

	return nil
}

func DeleteAgentV4(urls configuration.Url, accessToken string, swarmID string, nexusID string, nodeID string, agentID string) error {
	var err error
	url := urls.ChUrl + "/v4/swarms/" + swarmID + "/nexuses/" + nexusID + "/nodes/" + nodeID + "/agents/" + agentID

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodDelete),
		request_utils.WithAccessToken(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusNoContent),
		ExtractGenericModel(nil),
	); err != nil {
		return err
	}

	return nil
}

func ListAgentsForRCV4(urls configuration.Url, accessToken string, swarmID string, rcID, sort string, filter string) (*GenericPaginatedResponse[*NewAgent], error) {
	var err error
	var finalResponse GenericPaginatedResponse[*NewAgent]

	url := urls.ChUrl + "/v4/swarms/" + swarmID + "/redundancy_class/" + rcID + "/agents" + "?sort_key=" + sort + "&q=" + url.QueryEscape(filter)

	var nextPage *int
	page := 1

	for {
		var response GenericPaginatedResponse[*NewAgent]
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
