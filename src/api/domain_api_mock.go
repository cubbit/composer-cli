package api

import (
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
)

type MockDomainAPI struct {
	CreateFunc func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, request *CreateDomainRequestBody) (*DomainDTO, error)
	GetFunc    func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, domainID string) (*DomainDTO, error)
	ListFunc   func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, page int, itemsPerPage int) (*GenericPaginatedResponse[DomainDTO], error)
	DeleteFunc func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, domainID string) error
	VerifyFunc func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, domainID string) (*DomainVerifyResult, error)
}

func (m *MockDomainAPI) Create(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, request *CreateDomainRequestBody) (*DomainDTO, error) {
	if m.CreateFunc != nil {
		return m.CreateFunc(endpoints, apiKey, organizationID, request)
	}

	return &DomainDTO{}, nil
}

func (m *MockDomainAPI) Get(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, domainID string) (*DomainDTO, error) {
	if m.GetFunc != nil {
		return m.GetFunc(endpoints, apiKey, organizationID, domainID)
	}

	return &DomainDTO{}, nil
}

func (m *MockDomainAPI) List(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, page int, itemsPerPage int) (*GenericPaginatedResponse[DomainDTO], error) {
	if m.ListFunc != nil {
		return m.ListFunc(endpoints, apiKey, organizationID, page, itemsPerPage)
	}

	return &GenericPaginatedResponse[DomainDTO]{}, nil
}

func (m *MockDomainAPI) Delete(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, domainID string) error {
	if m.DeleteFunc != nil {
		return m.DeleteFunc(endpoints, apiKey, organizationID, domainID)
	}

	return nil
}

func (m *MockDomainAPI) Verify(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, domainID string) (*DomainVerifyResult, error) {
	if m.VerifyFunc != nil {
		return m.VerifyFunc(endpoints, apiKey, organizationID, domainID)
	}

	return &DomainVerifyResult{Verified: true}, nil
}

var _ DomainAPIInterface = (*MockDomainAPI)(nil)
