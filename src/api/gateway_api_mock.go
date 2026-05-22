package api

import "github.com/cubbit/composer-cli/src/configuration"

type MockGatewayAPI struct {
	CreateGatewayV5Func func(
		urlConfig configuration.URLs,
		apiKey string,
		organizationID string,
		request *CreateGatewayV5Request,
	) (*CreateGatewayV5Response, error)
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
