package api

import (
	"net/http"

	"github.com/cubbit/composer-cli/src/configuration"
	"github.com/cubbit/composer-cli/src/request_utils"
)

type UserAPIInterface interface {
	GetIAMUser(
		urls configuration.URLs,
		accessToken, apiKey string,
		meOrID string,
	) (*IAMUser, error)

	GetIAMUserSelf(
		urls configuration.URLs,
		accessToken, apiKey string,
	) (*IAMUser, error)

	PromoteIAMUser(
		urls configuration.URLs,
		email,
		policyName,
		secret string,
	) error
}

type UserAPI struct{}

func NewUserAPI() *UserAPI {
	return &UserAPI{}
}

func (a *UserAPI) GetIAMUser(urls configuration.URLs, accessToken, apiKey string, meOrID string) (*IAMUser, error) {
	var err error
	var operator IAMUser

	url := NewURLBuilder(urls.IamURL).
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

func (a *UserAPI) GetIAMUserSelf(urls configuration.URLs, accessToken, apiKey string) (*IAMUser, error) {
	return a.GetIAMUser(urls, accessToken, apiKey, "me")
}

func (a *UserAPI) PromoteIAMUser(urls configuration.URLs, email, policyName, secret string) error {
	var err error

	url := NewURLBuilder(urls.IamURL).
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
