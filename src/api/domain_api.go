package api

import (
	"net/http"

	"github.com/cubbit/composer-cli/src/configuration"
	"github.com/cubbit/composer-cli/src/request_utils"
)

type DomainAPIInterface interface {
	Create(
		urlConfig configuration.URLs,
		apiKey string,
		organizationID string,
		request *CreateDomainRequestBody,
	) (*DomainDTO, error)
	Get(
		urlConfig configuration.URLs,
		apiKey string,
		organizationID string,
		domainID string,
	) (*DomainDTO, error)
	List(
		urlConfig configuration.URLs,
		apiKey string,
		organizationID string,
		page int,
		itemPerPage int,
	) (*GenericPaginatedResponse[DomainDTO], error)
	Delete(
		urlConfig configuration.URLs,
		apiKey string,
		organizationID string,
		domainID string,
	) error
	Verify(
		urlConfig configuration.URLs,
		apiKey string,
		organizationID string,
		domainID string,
	) (*DomainVerifyResult, error)
}

type DomainAPI struct{}

func NewDomainAPI() *DomainAPI {
	return &DomainAPI{}
}

func (api *DomainAPI) Create(
	urlConfig configuration.URLs,
	apiKey string,
	organizationID string,
	request *CreateDomainRequestBody,
) (*DomainDTO, error) {
	url := NewURLBuilder(urlConfig.ChURL).
		Path("v1", "organizations", organizationID, "domains").
		Build()

	var response DomainDTO

	if err := request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPost),
		request_utils.WithExpectedStatusCode(http.StatusCreated),
		request_utils.WithApiKey(apiKey),
		request_utils.WithRequestBodyObject(request),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, err
	}

	return &response, nil
}

func (api *DomainAPI) Get(
	urlConfig configuration.URLs,
	apiKey string,
	organizationID string,
	domainID string,
) (*DomainDTO, error) {
	url := NewURLBuilder(urlConfig.ChURL).
		Path("v1", "organizations", organizationID, "domains", domainID).
		Build()

	var response DomainDTO

	if err := request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodGet),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		request_utils.WithApiKey(apiKey),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, err
	}

	return &response, nil
}

func (api *DomainAPI) List(
	urlConfig configuration.URLs,
	apiKey string,
	organizationID string,
	page int,
	itemPerPage int,
) (*GenericPaginatedResponse[DomainDTO], error) {
	url := NewURLBuilder(urlConfig.ChURL).
		Path("v1", "organizations", organizationID, "domains").
		QueryParamInt("page", page).
		QueryParamInt("items", itemPerPage).
		Build()

	var response GenericPaginatedResponse[DomainDTO]

	if err := request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodGet),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		request_utils.WithApiKey(apiKey),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, err
	}

	return &response, nil
}

func (api *DomainAPI) Delete(
	urlConfig configuration.URLs,
	apiKey string,
	organizationID string,
	domainID string,
) error {
	url := NewURLBuilder(urlConfig.ChURL).
		Path("v1", "organizations", organizationID, "domains", domainID).
		Build()

	if err := request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodDelete),
		request_utils.WithExpectedStatusCode(http.StatusNoContent),
		request_utils.WithApiKey(apiKey),
	); err != nil {
		return err
	}

	return nil
}

func (api *DomainAPI) Verify(
	urlConfig configuration.URLs,
	apiKey string,
	organizationID string,
	domainID string,
) (*DomainVerifyResult, error) {
	url := NewURLBuilder(urlConfig.ChURL).
		Path("v1", "organizations", organizationID, "domains", domainID, "verify").
		Build()

	var response DomainVerifyResult

	if err := request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPatch),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		request_utils.WithApiKey(apiKey),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, err
	}

	return &response, nil
}
