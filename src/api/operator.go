// Package api provides functions to interact with the operator API.
package api

import (
	"net/http"

	"github.com/cubbit/composer-cli/src/configuration"
	"github.com/cubbit/composer-cli/src/request_utils"
)

func GetOperator(urls configuration.URLs, accessToken, meOrID string) (*Operator, error) {
	var err error
	var operator Operator

	url := NewURLBuilder(urls.IamURL).
		Path("v1", "operators", meOrID).
		Build()

	if err = request_utils.DoRequest(
		url,
		request_utils.WithExpectedStatusCode(http.StatusOK),
		request_utils.WithAccessToken(accessToken),
		ExtractGenericModel(&operator),
	); err != nil {
		return nil, err
	}

	return &operator, nil
}

func GetOperatorSelf(urls configuration.URLs, accessToken string) (*Operator, error) {
	return GetOperator(urls, accessToken, "me")
}

func PromoteOperator(urls configuration.URLs, email, policyName, secret string) error {
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
