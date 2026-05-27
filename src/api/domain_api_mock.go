package api

import "github.com/cubbit/composer-cli/src/configuration"

type MockDomainAPI struct {
	CreateFunc func(urlConfig configuration.URLs, apiKey string, organizationID string, request *CreateDomainRequestBody) (*DomainDTO, error)
	GetFunc    func(urlConfig configuration.URLs, apiKey string, organizationID string, domainID string) (*DomainDTO, error)
	ListFunc   func(urlConfig configuration.URLs, apiKey string, organizationID string, page int, itemsPerPage int) (*GenericPaginatedResponse[DomainDTO], error)
	DeleteFunc func(urlConfig configuration.URLs, apiKey string, organizationID string, domainID string) error
	VerifyFunc func(urlConfig configuration.URLs, apiKey string, organizationID string, domainID string) (*DomainVerifyResult, error)
}

func (m *MockDomainAPI) Create(urlConfig configuration.URLs, apiKey string, organizationID string, request *CreateDomainRequestBody) (*DomainDTO, error) {
	if m.CreateFunc != nil {
		return m.CreateFunc(urlConfig, apiKey, organizationID, request)
	}

	return &DomainDTO{}, nil
}

func (m *MockDomainAPI) Get(urlConfig configuration.URLs, apiKey string, organizationID string, domainID string) (*DomainDTO, error) {
	if m.GetFunc != nil {
		return m.GetFunc(urlConfig, apiKey, organizationID, domainID)
	}

	return &DomainDTO{}, nil
}

func (m *MockDomainAPI) List(urlConfig configuration.URLs, apiKey string, organizationID string, page int, itemsPerPage int) (*GenericPaginatedResponse[DomainDTO], error) {
	if m.ListFunc != nil {
		return m.ListFunc(urlConfig, apiKey, organizationID, page, itemsPerPage)
	}

	return &GenericPaginatedResponse[DomainDTO]{}, nil
}

func (m *MockDomainAPI) Delete(urlConfig configuration.URLs, apiKey string, organizationID string, domainID string) error {
	if m.DeleteFunc != nil {
		return m.DeleteFunc(urlConfig, apiKey, organizationID, domainID)
	}

	return nil
}

func (m *MockDomainAPI) Verify(urlConfig configuration.URLs, apiKey string, organizationID string, domainID string) (*DomainVerifyResult, error) {
	if m.VerifyFunc != nil {
		return m.VerifyFunc(urlConfig, apiKey, organizationID, domainID)
	}

	return &DomainVerifyResult{Verified: true}, nil
}

var _ DomainAPIInterface = (*MockDomainAPI)(nil)
