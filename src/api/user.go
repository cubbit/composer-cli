// Package api provides functions to interact with the operator API.
package api

import (
	"net/http"

	"github.com/cubbit/composer-cli/src/configuration"
	"github.com/cubbit/composer-cli/src/request_utils"
)

func GetIAMUser(urls configuration.URLs, accessToken, apiKey string, meOrID string) (*Operator, error) {
	var err error
	var operator Operator

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

func GetIAMUserSelf(urls configuration.URLs, accessToken, apiKey string) (*Operator, error) {
	return GetIAMUser(urls, accessToken, apiKey, "me")
}

func PromoteIAMUser(urls configuration.URLs, email, policyName, secret string) error {
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
