// Package api provides functions to interact with the node API.
package api

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/src/request_utils"
)

func CreateNodes(urls configuration.URLs, accessToken string, swarmID string, nexusID string, nodesBody BulkInsertNewNodeRequestBody) (*NewNodesResponse, error) {
	var err error
	var response NewNodesResponse

	url := NewURLBuilder(urls.ChURL).
		Path("v4", "swarms", swarmID, "nexuses", nexusID, "nodes").
		Build()

	bodyRequest, err := json.Marshal(nodesBody)
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

func GetNode(urls configuration.URLs, accessToken string, swarmID, nexusID, nodeID string) (*NewNode, error) {
	var err error
	var response NewNode

	url := NewURLBuilder(urls.ChURL).
		Path("v4", "swarms", swarmID, "nexuses", nexusID, "nodes", nodeID).
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

func UpdateNode(urls configuration.URLs, accessToken string, swarmID, nexusID, nodeID string, nodeBody UpdateNewNodeRequestBody) error {
	var err error

	url := NewURLBuilder(urls.ChURL).
		Path("v4", "swarms", swarmID, "nexuses", nexusID, "nodes", nodeID).
		Build()

	bodyRequest, err := json.Marshal(nodeBody)
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

func DeleteNode(urls configuration.URLs, accessToken string, swarmID, nexusID, nodeID string) error {
	var err error

	url := NewURLBuilder(urls.ChURL).
		Path("v4", "swarms", swarmID, "nexuses", nexusID, "nodes", nodeID).
		Build()

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodDelete),
		request_utils.WithAccessToken(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusNoContent),
	); err != nil {
		return err
	}

	return nil
}

func ListNodes(urls configuration.URLs, accessToken string, swarmID string, nexusID, sort string, filter string) (*GenericPaginatedResponse[*NewNode], error) {
	var err error
	var finalResponse GenericPaginatedResponse[*NewNode]
	var nextPage *int

	url := NewURLBuilder(urls.ChURL).
		Path("v4", "swarms", swarmID, "nexuses", nexusID, "nodes").
		QueryParam("sort_key", sort).
		QueryParam("q", url.QueryEscape(filter)).
		Build()

	page := 1

	for {
		var response GenericPaginatedResponse[*NewNode]
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
