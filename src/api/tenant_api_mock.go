package api

import (
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
)

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

type MockTenantAPI struct {
	CreateTenantV5Func func(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		request *CreateTenantV5Request,
	) (*GenericIDResponseModel, error)
	ListTenantsV5Func func(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		opts ...ListTenantsV5Option,
	) (*GenericPaginatedResponse[TenantV5DTO], error)
	GetTenantV5Func func(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		tenantID string,
	) (*TenantV5DTO, error)
}

func (m *MockTenantAPI) ListTenantsV5(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	opts ...ListTenantsV5Option,
) (*GenericPaginatedResponse[TenantV5DTO], error) {
	if m.ListTenantsV5Func != nil {
		return m.ListTenantsV5Func(endpoints, apiKey, organizationID, opts...)
	}

	return &GenericPaginatedResponse[TenantV5DTO]{Data: []TenantV5DTO{}, Count: 0}, nil
}

func (m *MockTenantAPI) GetTenantV5(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	tenantID string,
) (*TenantV5DTO, error) {
	if m.GetTenantV5Func != nil {
		return m.GetTenantV5Func(endpoints, apiKey, organizationID, tenantID)
	}

	return &TenantV5DTO{}, nil
}
