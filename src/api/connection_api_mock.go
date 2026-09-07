package api

import (
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
)

type MockConnectionAPI struct {
	ListConnectionsV5Func func(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		tenantID string,
		opts ...ListConnectionsV5Option,
	) (*GenericPaginatedResponse[ConnectionV5DTO], error)
	VerifyConnectionV5Func func(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		tenantID string,
		connectionID string,
	) (*ConnectionV5DTO, error)
}

func NewMockConnectionAPI() *MockConnectionAPI {
	return &MockConnectionAPI{}
}

func (m *MockConnectionAPI) ListConnectionsV5(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	tenantID string,
	opts ...ListConnectionsV5Option,
) (*GenericPaginatedResponse[ConnectionV5DTO], error) {
	if m.ListConnectionsV5Func != nil {
		return m.ListConnectionsV5Func(endpoints, apiKey, organizationID, tenantID, opts...)
	}

	return &GenericPaginatedResponse[ConnectionV5DTO]{Data: []ConnectionV5DTO{}, Count: 0}, nil
}

func (m *MockConnectionAPI) VerifyConnectionV5(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	tenantID string,
	connectionID string,
) (*ConnectionV5DTO, error) {
	if m.VerifyConnectionV5Func != nil {
		return m.VerifyConnectionV5Func(endpoints, apiKey, organizationID, tenantID, connectionID)
	}

	return &ConnectionV5DTO{}, nil
}

var _ ConnectionAPIInterface = (*MockConnectionAPI)(nil)
