package api

import "github.com/cubbit/composer-cli/src/configuration"

type MockSwarmAPI struct {
	CreateSwarmV5Func func(urlConfig configuration.URLs, apiKey string, organizationID string, request *CreateSwarmV5Request) (*CreateSwarmV5Response, error)
	GetSwarmV5Func    func(urlConfig configuration.URLs, apiKey string, organizationID string, swarmID string) (*SwarmV5Presentation, error)
	ListSwarmsV5Func  func(urlConfig configuration.URLs, apiKey string, organizationID string, page int, items int) (*GenericPaginatedResponse[ListSwarmV5ItemPresentation], error)
}

func (m *MockSwarmAPI) CreateSwarmV5(urlConfig configuration.URLs, apiKey string, organizationID string, request *CreateSwarmV5Request) (*CreateSwarmV5Response, error) {
	if m.CreateSwarmV5Func != nil {
		return m.CreateSwarmV5Func(urlConfig, apiKey, organizationID, request)
	}

	return &CreateSwarmV5Response{ID: "test-process-id"}, nil
}

func (m *MockSwarmAPI) GetSwarmV5(urlConfig configuration.URLs, apiKey string, organizationID string, swarmID string) (*SwarmV5Presentation, error) {
	if m.GetSwarmV5Func != nil {
		return m.GetSwarmV5Func(urlConfig, apiKey, organizationID, swarmID)
	}

	return &SwarmV5Presentation{}, nil
}

func (m *MockSwarmAPI) ListSwarmsV5(urlConfig configuration.URLs, apiKey string, organizationID string, page int, items int) (*GenericPaginatedResponse[ListSwarmV5ItemPresentation], error) {
	if m.ListSwarmsV5Func != nil {
		return m.ListSwarmsV5Func(urlConfig, apiKey, organizationID, page, items)
	}

	return &GenericPaginatedResponse[ListSwarmV5ItemPresentation]{}, nil
}
