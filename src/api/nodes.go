package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/src/request_utils"
)

func CreateNode(urls configuration.Url, accessToken string, nodeBody CreateNodeBodyRequest) (*Node, error) {
	var err error
	var response Node
	url := urls.ChUrl + constants.NodesV2

	bodyRequest, err := json.Marshal(nodeBody)
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

func GetNode(urls configuration.Url, accessToken string, nodeID string) (*Node, error) {
	var err error
	var response Node
	url := urls.ChUrl + constants.Nodes + "/" + nodeID

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

func UpdateNode(urls configuration.Url, accessToken string, nodeID string, nodeBody UpdateNodeBodyRequest) error {
	var err error
	url := urls.ChUrl + constants.Nodes + "/" + nodeID

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

func DeleteNode(urls configuration.Url, accessToken string, nodeID string) error {
	var err error
	url := urls.ChUrl + constants.Nodes + "/" + nodeID

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

func ListNodes(urls configuration.Url, accessToken string, swarmID string, nexusID string) (*NodeList, error) {

	var err error
	var finalResponse NodeList
	url := urls.ChUrl + constants.Nodes + "?swarm_id=" + swarmID

	if nexusID != "" {
		url += "&nexus_id=" + nexusID
	}

	page := 0
	resultsPerPage := 1000

	for {
		var response NodeList
		if err = request_utils.DoRequest(
			url+"&start_page="+strconv.Itoa(page)+"&results_per_page="+strconv.Itoa(resultsPerPage),
			request_utils.WithAccessToken(accessToken),
			request_utils.WithExpectedStatusCode(http.StatusOK),
			ExtractGenericModel(&response),
		); err != nil {
			return nil, err
		}

		finalResponse.Nodes = append(finalResponse.Nodes, response.Nodes...)
		finalResponse.Count = response.Count

		if len(response.Nodes) == 0 {
			break
		}
		page++
	}

	return &finalResponse, nil
}

func CreateNodeV4(urls configuration.Url, accessToken string, swarmID string, nexusID string, nodesBody BulkInsertNewNodeRequestBody) (*NewNodesResponse, error) {
	var err error
	var response NewNodesResponse
	url := urls.ChUrl + "/v4/swarms/" + swarmID + "/nexuses/" + nexusID + "/nodes"

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
