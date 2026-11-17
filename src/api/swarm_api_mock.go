package api

import "github.com/cubbit/composer-cli/src/configuration/configuration_models"

type MockSwarmAPI struct {
	CreateSwarmV5Func func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, request *CreateSwarmV5Request) (*CreateSwarmV5Response, error)
	GetSwarmV5Func    func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, swarmID string) (*SwarmV5Presentation, error)
	ListSwarmsV5Func  func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, page int, items int) (*GenericPaginatedResponse[ListSwarmV5ItemPresentation], error)
}

func (m *MockSwarmAPI) CreateSwarmV5(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, request *CreateSwarmV5Request) (*CreateSwarmV5Response, error) {
	if m.CreateSwarmV5Func != nil {
		return m.CreateSwarmV5Func(endpoints, apiKey, organizationID, request)
	}

	return &CreateSwarmV5Response{ID: "test-process-id"}, nil
}

func (m *MockSwarmAPI) GetSwarmV5(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, swarmID string) (*SwarmV5Presentation, error) {
	if m.GetSwarmV5Func != nil {
		return m.GetSwarmV5Func(endpoints, apiKey, organizationID, swarmID)
	}

	return &SwarmV5Presentation{}, nil
}

func (m *MockSwarmAPI) ListSwarmsV5(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, page int, items int) (*GenericPaginatedResponse[ListSwarmV5ItemPresentation], error) {
	if m.ListSwarmsV5Func != nil {
		return m.ListSwarmsV5Func(endpoints, apiKey, organizationID, page, items)
	}

	return &GenericPaginatedResponse[ListSwarmV5ItemPresentation]{}, nil
}
