// Package api provides functions to interact with the swarm API.
package api

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"github.com/cubbit/composer-cli/src/configuration"
	"github.com/cubbit/composer-cli/src/request_utils"
)

func CreateSwarm(urls configuration.URLs, accessToken string, req CreateSwarmRequest) (*GenericIDResponseModel, error) {
	var err error
	var response GenericIDResponseModel
	var bodyRequest []byte

	url := NewURLBuilder(urls.ChURL).
		Path("v1", "swarms").
		Build()

	if bodyRequest, err = json.Marshal(req); err != nil {
		return nil, err
	}

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPost),
		request_utils.WithRequestBodyByte(bodyRequest),
		request_utils.WithExpectedStatusCode(http.StatusCreated),
		ExtractGenericModel(&response),
		request_utils.WithAccessToken(accessToken),
	); err != nil {
		return nil, err
	}

	return &response, nil
}

func GetSwarm(urls configuration.URLs, accessToken, ownerID string, swarmID string) (*Swarm, error) {
	var err error
	var response Swarm

	url := NewURLBuilder(urls.ChURL).
		Path("v1", "swarms", swarmID).
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

func EditSwarm(urls configuration.URLs, accessToken string, swarmID string, req UpdateSwarmRequest) error {
	var err error
	var bodyRequest []byte

	url := NewURLBuilder(urls.ChURL).
		Path("v1", "swarms", swarmID).Build()

	if bodyRequest, err = json.Marshal(req); err != nil {
		return err
	}

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPatch),
		request_utils.WithAccessToken(accessToken),
		request_utils.WithRequestBodyByte(bodyRequest),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		ExtractGenericModel(nil),
	); err != nil {
		return err
	}

	return nil
}

func ListSwarms(urls configuration.URLs, accessToken, ownerID string) ([]*Swarm, error) {
	var err error
	var response []*Swarm

	url := NewURLBuilder(urls.ChURL).
		Path("v2", "swarms").
		Build()

	if err = request_utils.DoRequest(
		url,
		request_utils.WithAccessToken(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, err
	}

	return response, nil
}

func GetSwarmOperator(urls configuration.URLs, accessToken, swarmID, operatorID string) (*Operator, error) {
	var err error
	var response Operator

	url := NewURLBuilder(urls.IamURL).
		Path("v1", "swarms", swarmID, "operators", operatorID).
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

func RemoveSwarm(urls configuration.URLs, accessToken, swarmID string) error {
	var err error

	url := NewURLBuilder(urls.ChURL).
		Path("v2", "swarms", swarmID).
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

func InviteOperatorToSwarm(urls configuration.URLs, accessToken, swarmID string, req InviteOperatorRequestBody) error {
	var err error
	var bodyRequest []byte

	url := NewURLBuilder(urls.IamURL).
		Path("v1", "swarms", swarmID, "invites").
		Build()

	if bodyRequest, err = json.Marshal(req); err != nil {
		return err
	}

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPost),
		request_utils.WithRequestBodyByte(bodyRequest),
		request_utils.WithExpectedStatusCode(http.StatusCreated),
		request_utils.WithAccessToken(accessToken),
	); err != nil {
		return err
	}

	return nil
}

func ListSwarmOperators(urls configuration.URLs, accessToken, swarmID string) (*OperatorList, error) {
	var err error
	var response OperatorList

	url := NewURLBuilder(urls.IamURL).
		Path("v1", "swarms", swarmID, "operators").
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

func RemoveSwarmOperator(urls configuration.URLs, accessToken, swarmID, operatorID string) error {
	var err error

	url := NewURLBuilder(urls.IamURL).
		Path("v1", "swarms", swarmID, "operators", operatorID).
		Build()

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodDelete),
		request_utils.WithAccessToken(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusOK),
	); err != nil {
		return err
	}

	return nil
}

func GetSwarmStatus(urls configuration.URLs, accessToken, swarmID string) (*SummaryDetailsWithStatusNullable, error) {
	var err error
	var response SummaryDetailsWithStatusNullable

	url := NewURLBuilder(urls.ChURL).
		Path("v2", "swarms", swarmID, "status").
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

func ListSwarmsV2(urls configuration.URLs, accessToken string, sort string, filter string) (*GenericPaginatedResponse[*Swarm], error) {
	var err error
	var finalResponse GenericPaginatedResponse[*Swarm]
	var nextPage *int

	url := NewURLBuilder(urls.ChURL).
		Path("v2", "swarms").
		QueryParam("sort_key", sort).
		QueryParam("q", url.QueryEscape(filter)).
		Build()

	page := 1

	for {
		var response GenericPaginatedResponse[*Swarm]
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

func EditOperatorRoleInSwarm(urls configuration.URLs, accessToken, swarmID, operatorID string, req ChangeOperatorPolicyRequestBody) error {
	var err error
	var bodyRequest []byte

	url := NewURLBuilder(urls.IamURL).
		Path("v1", "swarms", swarmID, "operators", operatorID, "roles").
		Build()

	if bodyRequest, err = json.Marshal(req); err != nil {
		return err
	}

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPut),
		request_utils.WithRequestBodyByte(bodyRequest),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		request_utils.WithAccessToken(accessToken),
	); err != nil {
		return err
	}

	return nil
}
