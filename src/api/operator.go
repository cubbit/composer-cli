package api

import (
	"fmt"
	"net/http"

	"github.com/cubbit/cubbit/client/cli/src/request_utils"
)

func GetOperator(apiServerUrl, accessToken, meOrID string) (*Operator, error) {
	var err error

	url := apiServerUrl + "/v1/operators/" + meOrID
	var operator Operator

	if err = request_utils.DoRequest(
		url,
		request_utils.WithExpectedStatusCode(http.StatusOK),
		request_utils.WithAccessToken(accessToken),
		extractOperatorResponseModel(&operator),
	); err != nil {
		return nil, fmt.Errorf("failed: unable to get an operator: %w", err)
	}

	return &operator, nil
}

func GetOperatorSelf(apiServerUrl, accessToken string) (*Operator, error) {
	return GetOperator(apiServerUrl, accessToken, "me")
}
