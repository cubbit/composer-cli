package api

import "github.com/cubbit/composer-cli/src/configuration/configuration_models"

type MockGatewayAPI struct {
	CreateGatewayV5Func func(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		request *CreateGatewayV5Request,
	) (*CreateGatewayV5Response, error)
	ListGatewaysV5Func func(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		opts ...ListGatewaysV5Option,
	) (*GenericPaginatedResponse[GatewayV5ListItemResponse], error)
	GetGatewayV5Func func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, gatewayID string) (*GatewayV5GetResponse, error)
}

func (m *MockGatewayAPI) CreateGatewayV5(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	request *CreateGatewayV5Request,
) (*CreateGatewayV5Response, error) {
	if m.CreateGatewayV5Func != nil {
		return m.CreateGatewayV5Func(endpoints, apiKey, organizationID, request)
	}

	return &CreateGatewayV5Response{ID: "test-gateway-id"}, nil
}

func (m *MockGatewayAPI) ListGatewaysV5(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	opts ...ListGatewaysV5Option,
) (*GenericPaginatedResponse[GatewayV5ListItemResponse], error) {
	if m.ListGatewaysV5Func != nil {
		return m.ListGatewaysV5Func(endpoints, apiKey, organizationID, opts...)
	}

	return &GenericPaginatedResponse[GatewayV5ListItemResponse]{Data: []GatewayV5ListItemResponse{}, Count: 0}, nil
}

func (m *MockGatewayAPI) GetGatewayV5(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, gatewayID string) (*GatewayV5GetResponse, error) {
	if m.GetGatewayV5Func != nil {
		return m.GetGatewayV5Func(endpoints, apiKey, organizationID, gatewayID)
	}

	return &GatewayV5GetResponse{}, nil
}
