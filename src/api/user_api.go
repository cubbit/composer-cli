package api

import (
	"fmt"
	"net/http"

	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/src/request_utils"
)

type UserAPIInterface interface {
	GetIAMUser(
		endpoints configuration_models.EndpointsV2,
		accessToken, apiKey string,
		meOrID string,
	) (*IAMUser, error)

	GetIAMUserSelf(
		endpoints configuration_models.EndpointsV2,
		accessToken, apiKey string,
	) (*IAMUser, error)

	PromoteIAMUser(
		endpoints configuration_models.EndpointsV2,
		email,
		policyName,
		secret string,
	) error

	BulkCreateIAMUsers(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		request *BulkCreateIAMUsersRequestBody,
	) (*BulkCreateIAMUsersResponse, error)
}

type UserAPI struct{}

func NewUserAPI() *UserAPI {
	return &UserAPI{}
}

func (a *UserAPI) GetIAMUser(endpoints configuration_models.EndpointsV2, accessToken, apiKey string, meOrID string) (*IAMUser, error) {
	var err error
	var operator IAMUser

	url := NewURLBuilder(endpoints.IAM).
		Path("v1", "operators", meOrID).
		Build()

	options := []request_utils.RequestModifier{
		ExtractGenericModel(&operator),
		request_utils.WithExpectedStatusCode(http.StatusOK),
	}

	if accessToken != "" {
		options = append(options, request_utils.WithAccessToken(accessToken))
	}
	if apiKey != "" {
		options = append(options, request_utils.WithApiKey(apiKey))
	}

	if err = request_utils.DoRequest(
		url,

		options...,
	); err != nil {
		return nil, err
	}

	return &operator, nil
}

func (a *UserAPI) GetIAMUserSelf(endpoints configuration_models.EndpointsV2, accessToken, apiKey string) (*IAMUser, error) {
	return a.GetIAMUser(endpoints, accessToken, apiKey, "me")
}

func (a *UserAPI) PromoteIAMUser(endpoints configuration_models.EndpointsV2, email, policyName, secret string) error {
	var err error

	url := NewURLBuilder(endpoints.IAM).
		Path("v1", "operators", "promote").
		Build()

	requestBody := map[string]interface{}{
		"email":       email,
		"policy_name": policyName,
		"secret":      secret,
	}

	if err = request_utils.DoRequest(url, request_utils.WithRequestMethod(http.MethodPost), request_utils.WithRequestBody(requestBody), request_utils.WithExpectedStatusCode(http.StatusCreated)); err != nil {
		return err
	}

	return nil
}

func (a *UserAPI) BulkCreateIAMUsers(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	request *BulkCreateIAMUsersRequestBody,
) (*BulkCreateIAMUsersResponse, error) {
	url := NewURLBuilder(endpoints.IAM).
		Path("v3", "organizations", organizationID, "operators", "bulk").
		Build()

	requestBody := map[string][]BulkCreateIAMUserRequestBody{
		"operators": request.Users,
	}

	var response BulkCreateIAMUsersResponse
	if err := request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPost),
		request_utils.WithExpectedStatusCode(http.StatusCreated),
		request_utils.WithApiKey(apiKey),
		request_utils.WithRequestBodyObject(requestBody),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, fmt.Errorf("failed to perform bulk create users request: %w", err)
	}

	return &response, nil
}
