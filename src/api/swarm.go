package api

import (
	"fmt"
	"net/http"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/src/request_utils"
)

func CreateSwarm(urls configuration.Url, accessToken, ownerID string, name string, description string, swarmConfig map[string]interface{}) (*GenericIDResponseModel, error) {
	var err error
	var response GenericIDResponseModel
	url := fmt.Sprintf("%s%s", urls.HiveUrl, constants.Swarms)

	requestBody := map[string]interface{}{
		"name":          name,
		"description":   description,
		"owner_id":      ownerID,
		"configuration": swarmConfig,
	}

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPost),
		request_utils.WithRequestBody(requestBody),
		request_utils.WithExpectedStatusCode(http.StatusCreated),
		extractGenericIDResponseModel(&response),
		request_utils.WithAccessToken(accessToken),
	); err != nil {
		return nil, fmt.Errorf("%s: %w", constants.ErrorCreatingSwarmRequest, err)
	}

	return &response, nil
}

func ListSwarms(urls configuration.Url, accessToken, ownerID string) ([]Swarm, error) {
	var err error
	url := fmt.Sprintf("%s%s", urls.HiveUrl, constants.Swarms)
	var response []Swarm

	if err = request_utils.DoRequest(
		url,
		request_utils.WithAccessToken(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		extractSwarmListResponseModel(&response),
	); err != nil {
		return nil, fmt.Errorf("%s: %w", constants.ErrorListingSwarmsRequest, err)
	}
	return response, nil
}

func GetSwarm(urls configuration.Url, accessToken, ownerID string, swarmID string) (*Swarm, error) {
	var err error
	url := fmt.Sprintf("%s%s/%s", urls.HiveUrl, constants.Swarms, swarmID)
	var response Swarm

	if err = request_utils.DoRequest(
		url,
		request_utils.WithAccessToken(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		extractSwarmResponseModel(&response),
	); err != nil {
		return nil, fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarmsRequest, err)
	}
	return &response, nil
}

func RemoveSwarm(urls configuration.Url, accessToken, swarmId, deleteSwarmToken string) error {
	var err error

	url := urls.HiveUrl + constants.Swarms + "/" + swarmId + "?token=" + deleteSwarmToken

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodDelete),
		request_utils.WithAccessToken(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusOK),
	); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDeletingSwarmRequest, err)
	}

	return nil
}

func EditSwarmDescription(urls configuration.Url, accessToken, swarmID, description string) error {
	var err error

	requestBody := map[string]interface{}{
		"description": description,
	}

	url := urls.HiveUrl + constants.Swarms + "/" + swarmID

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPut),
		request_utils.WithAccessToken(accessToken),
		request_utils.WithRequestBody(requestBody),
		request_utils.WithExpectedStatusCode(http.StatusOK),
	); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorEditingSwarmRequest, err)
	}

	return nil
}

func EditSwarmName(urls configuration.Url, accessToken, swarmID, name string) error {
	var err error

	requestBody := map[string]interface{}{
		"name": name,
	}

	url := urls.HiveUrl + constants.Swarms + "/" + swarmID

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPut),
		request_utils.WithAccessToken(accessToken),
		request_utils.WithRequestBody(requestBody),
		request_utils.WithExpectedStatusCode(http.StatusOK),
	); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorEditingSwarmRequest, err)
	}

	return nil
}

func ListSwarmPolicies(urls configuration.Url, accessToken, swarmID string) (*PolicyList, error) {
	var err error
	url := urls.IamUrl + constants.Policies + "?swarm=" + swarmID
	var response PolicyList

	if err = request_utils.DoRequest(
		url,
		request_utils.WithAccessToken(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		extractPolicyListModel(&response),
	); err != nil {
		return nil, fmt.Errorf("%s: %w", constants.ErrorListingPoliciesRequest, err)
	}
	return &response, nil
}

func InviteOperatorToSwarm(urls configuration.Url, accessToken, swarmID, email, role, firstName, lastName string) error {
	var err error
	url := urls.IamUrl + constants.Swarms + "/" + swarmID + constants.Invites

	requestBody := map[string]interface{}{
		"email":      email,
		"policy_id":  role,
		"first_name": firstName,
		"last_name":  lastName,
	}

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPost),
		request_utils.WithRequestBody(requestBody),
		request_utils.WithExpectedStatusCode(http.StatusCreated),
		request_utils.WithAccessToken(accessToken),
	); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorInvitingOperatorRequest, err)
	}

	return nil
}

func ListSwarmOperators(urls configuration.Url, accessToken, swarmID string) (*OperatorList, error) {
	var err error
	url := urls.IamUrl + constants.Swarms + "/" + swarmID + "/operators"
	var response OperatorList

	if err = request_utils.DoRequest(
		url,
		request_utils.WithAccessToken(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		extractOperatorListModel(&response),
	); err != nil {
		return nil, fmt.Errorf("%s: %w", constants.ErrorListingOperatorsRequest, err)
	}
	return &response, nil
}

func RemoveSwarmOperator(urls configuration.Url, accessToken, swarmID, operatorID string) error {
	var err error
	url := urls.IamUrl + constants.Swarms + "/" + swarmID + "/operators/" + operatorID
	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodDelete),
		request_utils.WithAccessToken(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusOK),
	); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRemovingOperatorsRequest, err)
	}

	return nil
}
