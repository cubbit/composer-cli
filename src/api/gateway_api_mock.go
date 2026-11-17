package api

import (
	"encoding/json"

	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
)

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
	GetGatewayV5Func    func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, gatewayID string) (*GatewayV5GetResponse, error)
	UpdateGatewayV5Func func(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		gatewayID string,
		request *UpdateGatewayV5Request,
	) (*GatewayV5GetResponse, error)
	ListGatewayStorageClassesV5Func func(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		gatewayID string,
	) (*GatewayStorageClassesV5ListResponse, error)
	CreateGatewayStorageClassV5Func func(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		gatewayID string,
		request *CreateGatewayStorageClassV5Request,
	) (*GatewayStorageClassV5, error)
	DeleteGatewayStorageClassV5Func func(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		gatewayID string,
		storageClassID string,
	) error
	ListGatewayConnectionsV5Func func(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		gatewayID string,
	) (*GatewayConnectionsV5ListResponse, error)
	CreateGatewayConnectionsV5Func func(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		gatewayID string,
		request *CreateGatewayConnectionsV5Request,
	) (*GatewayConnectionsV5ListResponse, error)
	DeleteGatewayConnectionV5Func func(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		gatewayID string,
		connectionID string,
	) error
	ListGatewayRoutingRulesV5Func func(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		gatewayID string,
	) (*GatewayRoutingRulesV5ListResponse, error)
	GetGatewayRoutingRuleV5RawFunc func(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		gatewayID string,
		ruleID string,
	) (json.RawMessage, error)
	CreateGatewayRoutingRuleV5Func func(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		gatewayID string,
		body json.RawMessage,
	) (json.RawMessage, error)
	UpdateGatewayRoutingRuleV5Func func(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		gatewayID string,
		ruleID string,
		body json.RawMessage,
	) (json.RawMessage, error)
	DeleteGatewayRoutingRuleV5Func func(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		gatewayID string,
		ruleID string,
	) error
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

func (m *MockGatewayAPI) UpdateGatewayV5(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	gatewayID string,
	request *UpdateGatewayV5Request,
) (*GatewayV5GetResponse, error) {
	if m.UpdateGatewayV5Func != nil {
		return m.UpdateGatewayV5Func(endpoints, apiKey, organizationID, gatewayID, request)
	}

	return &GatewayV5GetResponse{}, nil
}

func (m *MockGatewayAPI) ListGatewayStorageClassesV5(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	gatewayID string,
) (*GatewayStorageClassesV5ListResponse, error) {
	if m.ListGatewayStorageClassesV5Func != nil {
		return m.ListGatewayStorageClassesV5Func(endpoints, apiKey, organizationID, gatewayID)
	}

	return &GatewayStorageClassesV5ListResponse{Data: []GatewayStorageClassV5{}}, nil
}

func (m *MockGatewayAPI) CreateGatewayStorageClassV5(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	gatewayID string,
	request *CreateGatewayStorageClassV5Request,
) (*GatewayStorageClassV5, error) {
	if m.CreateGatewayStorageClassV5Func != nil {
		return m.CreateGatewayStorageClassV5Func(endpoints, apiKey, organizationID, gatewayID, request)
	}

	return &GatewayStorageClassV5{}, nil
}

func (m *MockGatewayAPI) DeleteGatewayStorageClassV5(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	gatewayID string,
	storageClassID string,
) error {
	if m.DeleteGatewayStorageClassV5Func != nil {
		return m.DeleteGatewayStorageClassV5Func(endpoints, apiKey, organizationID, gatewayID, storageClassID)
	}

	return nil
}

func (m *MockGatewayAPI) ListGatewayConnectionsV5(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	gatewayID string,
) (*GatewayConnectionsV5ListResponse, error) {
	if m.ListGatewayConnectionsV5Func != nil {
		return m.ListGatewayConnectionsV5Func(endpoints, apiKey, organizationID, gatewayID)
	}

	return &GatewayConnectionsV5ListResponse{Data: []GatewayConnectionV5{}}, nil
}

func (m *MockGatewayAPI) CreateGatewayConnectionsV5(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	gatewayID string,
	request *CreateGatewayConnectionsV5Request,
) (*GatewayConnectionsV5ListResponse, error) {
	if m.CreateGatewayConnectionsV5Func != nil {
		return m.CreateGatewayConnectionsV5Func(endpoints, apiKey, organizationID, gatewayID, request)
	}

	return &GatewayConnectionsV5ListResponse{Data: []GatewayConnectionV5{}}, nil
}

func (m *MockGatewayAPI) DeleteGatewayConnectionV5(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	gatewayID string,
	connectionID string,
) error {
	if m.DeleteGatewayConnectionV5Func != nil {
		return m.DeleteGatewayConnectionV5Func(endpoints, apiKey, organizationID, gatewayID, connectionID)
	}

	return nil
}

func (m *MockGatewayAPI) ListGatewayRoutingRulesV5(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	gatewayID string,
) (*GatewayRoutingRulesV5ListResponse, error) {
	if m.ListGatewayRoutingRulesV5Func != nil {
		return m.ListGatewayRoutingRulesV5Func(endpoints, apiKey, organizationID, gatewayID)
	}

	return &GatewayRoutingRulesV5ListResponse{}, nil
}

func (m *MockGatewayAPI) CreateGatewayRoutingRuleV5(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	gatewayID string,
	body json.RawMessage,
) (json.RawMessage, error) {
	if m.CreateGatewayRoutingRuleV5Func != nil {
		return m.CreateGatewayRoutingRuleV5Func(endpoints, apiKey, organizationID, gatewayID, body)
	}

	return nil, nil
}

func (m *MockGatewayAPI) GetGatewayRoutingRuleV5Raw(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	gatewayID string,
	ruleID string,
) (json.RawMessage, error) {
	if m.GetGatewayRoutingRuleV5RawFunc != nil {
		return m.GetGatewayRoutingRuleV5RawFunc(endpoints, apiKey, organizationID, gatewayID, ruleID)
	}

	return nil, nil
}

func (m *MockGatewayAPI) UpdateGatewayRoutingRuleV5(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	gatewayID string,
	ruleID string,
	body json.RawMessage,
) (json.RawMessage, error) {
	if m.UpdateGatewayRoutingRuleV5Func != nil {
		return m.UpdateGatewayRoutingRuleV5Func(endpoints, apiKey, organizationID, gatewayID, ruleID, body)
	}

	return nil, nil
}

func (m *MockGatewayAPI) DeleteGatewayRoutingRuleV5(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	gatewayID string,
	ruleID string,
) error {
	if m.DeleteGatewayRoutingRuleV5Func != nil {
		return m.DeleteGatewayRoutingRuleV5Func(endpoints, apiKey, organizationID, gatewayID, ruleID)
	}

	return nil
}
