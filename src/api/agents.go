// Package api provides functions to interact with the agents API.
package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/src/request_utils"
)

func CreateAgent(urls configuration.URLs, accessToken string, swarmID string, nexusID string, nodeID string, agentsBody BulkInsertNewAgentRequestBody) (*NewAgentsResponse, error) {
	var err error
	var response NewAgentsResponse
	var bodyRequest []byte

	url := NewURLBuilder(urls.ChURL).
		Path("v4", "swarms", swarmID, "nexuses", nexusID, "nodes", nodeID, "agents").
		Build()

	bodyRequest, err = json.Marshal(agentsBody)
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

func ListAgents(urls configuration.URLs, accessToken string, swarmID string, nexusID, nodeID, sort string, filter string) (*GenericPaginatedResponse[*NewAgent], error) {
	var err error
	var finalResponse GenericPaginatedResponse[*NewAgent]
	var nextPage *int

	url := NewURLBuilder(urls.ChURL).
		Path("v4", "swarms", swarmID, "nexuses", nexusID, "nodes", nodeID, "agents").
		QueryParam("sort_key", sort).
		QueryParam("q", url.QueryEscape(filter)).
		Build()

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

func GetAgent(urls configuration.URLs, accessToken string, swarmID, nexusID, nodeID, agentID string) (*GenericPaginatedResponse[*NewAgent], error) {
	var err error
	var response GenericPaginatedResponse[*NewAgent]

	url := NewURLBuilder(urls.ChURL).
		Path("v4", "swarms", swarmID, "nexuses", nexusID, "nodes", nodeID, "agents").
		QueryParam("q", fmt.Sprintf("agent:eq(%s)", agentID)).
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

func UpdateAgent(urls configuration.URLs, accessToken string, swarmID string, nexusID string, nodeID string, agentID string, nodesBody UpdateNewAgentRequestBody) error {
	var err error
	var bodyRequest []byte

	url := NewURLBuilder(urls.ChURL).
		Path("v4", "swarms", swarmID, "nexuses", nexusID, "nodes", nodeID, "agents", agentID).
		Build()

	bodyRequest, err = json.Marshal(nodesBody)
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

func DeleteAgent(urls configuration.URLs, accessToken string, swarmID string, nexusID string, nodeID string, agentID string) error {
	var err error

	url := NewURLBuilder(urls.ChURL).
		Path("v4", "swarms", swarmID, "nexuses", nexusID, "nodes", nodeID, "agents", agentID).
		Build()

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

func ListAgentsForRC(urls configuration.URLs, accessToken string, swarmID string, rcID, sort string, filter string) (*GenericPaginatedResponse[*NewAgent], error) {
	var err error
	var finalResponse GenericPaginatedResponse[*NewAgent]
	var nextPage *int

	url := NewURLBuilder(urls.ChURL).
		Path("v4", "swarms", swarmID, "redundancy_class", rcID, "agents").
		QueryParam("sort_key", sort).
		QueryParam("q", url.QueryEscape(filter)).
		Build()

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

func GetAgentStatus(urls configuration.URLs, accessToken, swarmID, nexusID, nodeID, agentID string) (*GetAgentEvaluatedStatusResponse, error) {

	var err error

	var response GetAgentEvaluatedStatusResponse

	url := NewURLBuilder(urls.ChURL).
		Path("v4", "swarms", swarmID, "nexuses", nexusID, "nodes", nodeID, "agents", agentID, "status").
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
