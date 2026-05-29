package api

import "github.com/cubbit/composer-cli/src/configuration"

type MockGatewayAPI struct {
	CreateGatewayV5Func func(
		urlConfig configuration.URLs,
		apiKey string,
		organizationID string,
		request *CreateGatewayV5Request,
	) (*CreateGatewayV5Response, error)
	ListGatewaysV5Func func(
		urlConfig configuration.URLs,
		apiKey string,
		organizationID string,
		opts ...ListGatewaysV5Option,
	) (*GenericPaginatedResponse[GatewayV5ListItemResponse], error)
	GetGatewayV5Func func(urlConfig configuration.URLs, apiKey string, organizationID string, gatewayID string) (*GatewayV5GetResponse, error)
}

func (m *MockGatewayAPI) CreateGatewayV5(
	urlConfig configuration.URLs,
	apiKey string,
	organizationID string,
	request *CreateGatewayV5Request,
) (*CreateGatewayV5Response, error) {
	if m.CreateGatewayV5Func != nil {
		return m.CreateGatewayV5Func(urlConfig, apiKey, organizationID, request)
	}

	return &CreateGatewayV5Response{ID: "test-gateway-id"}, nil
}

func (m *MockGatewayAPI) ListGatewaysV5(
	urlConfig configuration.URLs,
	apiKey string,
	organizationID string,
	opts ...ListGatewaysV5Option,
) (*GenericPaginatedResponse[GatewayV5ListItemResponse], error) {
	if m.ListGatewaysV5Func != nil {
		return m.ListGatewaysV5Func(urlConfig, apiKey, organizationID, opts...)
	}

	return &GenericPaginatedResponse[GatewayV5ListItemResponse]{Data: []GatewayV5ListItemResponse{}, Count: 0}, nil
}

func (m *MockGatewayAPI) GetGatewayV5(urlConfig configuration.URLs, apiKey string, organizationID string, gatewayID string) (*GatewayV5GetResponse, error) {
	if m.GetGatewayV5Func != nil {
		return m.GetGatewayV5Func(urlConfig, apiKey, organizationID, gatewayID)
	}

	return &GatewayV5GetResponse{}, nil
}
