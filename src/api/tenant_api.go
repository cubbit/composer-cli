package api

import (
	"net/http"

	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/src/request_utils"
)

type ListTenantsV5Options struct {
	Page          int
	Items         int
	SortKey       string
	SortOrder     string
	Filter        string
	FetchAllPages bool
}

type ListTenantsV5Option func(*ListTenantsV5Options)

func WithTenantPage(page int) ListTenantsV5Option {
	return func(opts *ListTenantsV5Options) {
		opts.Page = page
	}
}

func WithTenantItems(items int) ListTenantsV5Option {
	return func(opts *ListTenantsV5Options) {
		opts.Items = items
	}
}

func WithTenantSortKey(sortKey string) ListTenantsV5Option {
	return func(opts *ListTenantsV5Options) {
		opts.SortKey = sortKey
	}
}

func WithTenantSortOrder(sortOrder string) ListTenantsV5Option {
	return func(opts *ListTenantsV5Options) {
		opts.SortOrder = sortOrder
	}
}

func WithTenantFilter(filter string) ListTenantsV5Option {
	return func(opts *ListTenantsV5Options) {
		opts.Filter = filter
	}
}

func WithTenantListAllPages(fetchAll bool) ListTenantsV5Option {
	return func(opts *ListTenantsV5Options) {
		opts.FetchAllPages = fetchAll
	}
}

type TenantAPIInterface interface {
	CreateTenantV5(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		request *CreateTenantV5Request,
	) (*GenericIDResponseModel, error)
	ListTenantsV5(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		opts ...ListTenantsV5Option,
	) (*GenericPaginatedResponse[TenantV5DTO], error)
	GetTenantV5(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		tenantID string,
	) (*TenantV5DTO, error)
}

type TenantAPI struct{}

func NewTenantAPI() *TenantAPI {
	return &TenantAPI{}
}

func (api *TenantAPI) CreateTenantV5(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	request *CreateTenantV5Request,
) (*GenericIDResponseModel, error) {
	url := NewURLBuilder(endpoints.CH).
		Path("v5", "organizations", organizationID, "tenants").
		Build()

	var response GenericIDResponseModel

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

func (api *TenantAPI) ListTenantsV5(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	opts ...ListTenantsV5Option,
) (*GenericPaginatedResponse[TenantV5DTO], error) {
	options := &ListTenantsV5Options{
		Page:  1,
		Items: 100,
	}

	for _, opt := range opts {
		opt(options)
	}

	allData := make([]TenantV5DTO, 0)
	nextPage := &options.Page

	for nextPage != nil {
		urlBuilder := NewURLBuilder(endpoints.CH).
			Path("v5", "organizations", organizationID, "tenants").
			QueryParamInt("page", *nextPage).
			QueryParamInt("items", options.Items)

		if options.SortKey != "" {
			urlBuilder = urlBuilder.QueryParam("sort_key", options.SortKey)
		}

		if options.SortOrder != "" {
			urlBuilder = urlBuilder.QueryParam("sort_order", options.SortOrder)
		}

		if options.Filter != "" {
			urlBuilder = urlBuilder.QueryParam("q", options.Filter)
		}

		url := urlBuilder.Build()

		var pageResponse GenericPaginatedResponse[TenantV5DTO]

		if err := request_utils.DoRequest(
			url,
			request_utils.WithRequestMethod(http.MethodGet),
			request_utils.WithExpectedStatusCode(http.StatusOK),
			request_utils.WithApiKey(apiKey),
			ExtractGenericModel(&pageResponse),
		); err != nil {
			return nil, err
		}

		allData = append(allData, pageResponse.Data...)

		if !options.FetchAllPages {
			pageResponse.Data = allData
			return &pageResponse, nil
		}

		nextPage = pageResponse.NextPage
	}

	var count int
	if len(allData) > 0 {
		count = len(allData)
	}

	return &GenericPaginatedResponse[TenantV5DTO]{
		Data:     allData,
		NextPage: nil,
		Count:    count,
	}, nil
}

func (api *TenantAPI) GetTenantV5(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	tenantID string,
) (*TenantV5DTO, error) {
	url := NewURLBuilder(endpoints.CH).
		Path("v5", "organizations", organizationID, "tenants", tenantID).
		Build()

	var response TenantV5DTO

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
