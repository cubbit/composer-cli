package api

import (
	"net/http"

	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/src/request_utils"
)

type TenantAPIInterface interface {
	CreateTenantV5(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		request *CreateTenantV5Request,
	) (*GenericIDResponseModel, error)
}

type TenantAPI struct{}

func NewTenantAPI() *TenantAPI {
	return &TenantAPI{}
}

func (api *TenantAPI) CreateTenantV5(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	request *CreateTenantV5Request,
) (*GenericIDResponseModel, error) {
	url := NewURLBuilder(endpoints.CH).
		Path("v5", "organizations", organizationID, "tenants").
		Build()

	var response GenericIDResponseModel

	if err := request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPost),
		request_utils.WithExpectedStatusCode(http.StatusCreated),
		request_utils.WithApiKey(apiKey),
		request_utils.WithRequestBodyObject(request),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, err
	}

	return &response, nil
}
