package api

import "github.com/cubbit/composer-cli/src/configuration/configuration_models"

type MockTenantAPI struct {
	CreateTenantV5Func func(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		request *CreateTenantV5Request,
	) (*GenericIDResponseModel, error)
}

func (m *MockTenantAPI) CreateTenantV5(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	request *CreateTenantV5Request,
) (*GenericIDResponseModel, error) {
	if m.CreateTenantV5Func != nil {
		return m.CreateTenantV5Func(endpoints, apiKey, organizationID, request)
	}

	return &GenericIDResponseModel{ID: "test-tenant-id"}, nil
}
