package api

import (
	"net/http"

	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/src/request_utils"
)

type ListConnectionsV5Options struct {
	Page          int
	Items         int
	SortKey       string
	SortOrder     string
	Filter        string
	FetchAllPages bool
}

type ListConnectionsV5Option func(*ListConnectionsV5Options)

func WithConnectionsPage(page int) ListConnectionsV5Option {
	return func(opts *ListConnectionsV5Options) {
		opts.Page = page
	}
}

func WithConnectionsItems(items int) ListConnectionsV5Option {
	return func(opts *ListConnectionsV5Options) {
		opts.Items = items
	}
}

func WithConnectionsSortKey(sortKey string) ListConnectionsV5Option {
	return func(opts *ListConnectionsV5Options) {
		opts.SortKey = sortKey
	}
}

func WithConnectionsSortOrder(sortOrder string) ListConnectionsV5Option {
	return func(opts *ListConnectionsV5Options) {
		opts.SortOrder = sortOrder
	}
}

func WithConnectionsFilter(filter string) ListConnectionsV5Option {
	return func(opts *ListConnectionsV5Options) {
		opts.Filter = filter
	}
}

func WithConnectionsFetchAllPages(fetchAll bool) ListConnectionsV5Option {
	return func(opts *ListConnectionsV5Options) {
		opts.FetchAllPages = fetchAll
	}
}

type ConnectionAPIInterface interface {
	ListConnectionsV5(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		tenantID string,
		opts ...ListConnectionsV5Option,
	) (*GenericPaginatedResponse[ConnectionV5DTO], error)
	VerifyConnectionV5(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		tenantID string,
		connectionID string,
	) (*ConnectionV5DTO, error)
}

type ConnectionAPI struct{}

func NewConnectionAPI() *ConnectionAPI {
	return &ConnectionAPI{}
}

func (api *ConnectionAPI) VerifyConnectionV5(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	tenantID string,
	connectionID string,
) (*ConnectionV5DTO, error) {
	url := NewURLBuilder(endpoints.CH).
		Path("v5", "organizations", organizationID, "tenants", tenantID, "connections", connectionID, "verify").
		Build()

	var response ConnectionV5DTO

	if err := request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPost),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		request_utils.WithApiKey(apiKey),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, err
	}

	return &response, nil
}

func (api *ConnectionAPI) ListConnectionsV5(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	tenantID string,
	opts ...ListConnectionsV5Option,
) (*GenericPaginatedResponse[ConnectionV5DTO], error) {
	options := &ListConnectionsV5Options{
		Page:  1,
		Items: 100,
	}

	for _, opt := range opts {
		opt(options)
	}

	allData := make([]ConnectionV5DTO, 0)
	nextPage := &options.Page

	for nextPage != nil {
		urlBuilder := NewURLBuilder(endpoints.CH).
			Path("v5", "organizations", organizationID, "tenants", tenantID, "connections").
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

		var pageResponse GenericPaginatedResponse[ConnectionV5DTO]

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

	return &GenericPaginatedResponse[ConnectionV5DTO]{
		Data:     allData,
		NextPage: nil,
		Count:    count,
	}, nil
}
