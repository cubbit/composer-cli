package api

import (
	"net/http"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/src/request_utils"
)

func GetOperator(urls configuration.Url, accessToken, meOrID string) (*Operator, error) {
	var err error

	url := urls.IamUrl + constants.Operators + meOrID
	var operator Operator

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

func GetOperatorSelf(urls configuration.Url, accessToken string) (*Operator, error) {
	return GetOperator(urls, accessToken, "me")
}
