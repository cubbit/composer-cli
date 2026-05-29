package api

import (
	"net/http"

	"github.com/cubbit/composer-cli/src/configuration"
	"github.com/cubbit/composer-cli/src/request_utils"
)

type CubbitIngressType string

const (
	IngressTypeManual                 CubbitIngressType = "manual"
	IngressTypeSingleCluster          CubbitIngressType = "singlecluster"
	IngressTypeMulticlusterController CubbitIngressType = "multicluster_controller"
	IngressTypeMulticlusterWorker     CubbitIngressType = "multicluster_worker"
)

type SwarmAndRedundancyClassV5 struct {
	SwarmID           string `json:"swarm_id"`
	RedundancyClassID string `json:"redundancy_class_id"`
	IsDefault         bool   `json:"is_default"`
}

type CubbitIngress struct {
	Type            CubbitIngressType `json:"type"`
	CertSecretName  *string           `json:"cert_secret_name,omitempty"`
	StdHostname     *string           `json:"std_hostname,omitempty"`
	StarHostname    *string           `json:"star_hostname,omitempty"`
	ConsoleHostname *string           `json:"console_hostname,omitempty"`
	ExternalIPs     *[]string         `json:"external_ips,omitempty"`
}

type CreateGatewayV5Request struct {
	ClusterID                string                      `json:"cluster_id"`
	Name                     string                      `json:"name"`
	Slug                     string                      `json:"slug"`
	Description              *string                     `json:"description,omitempty"`
	SwarmsAndRedundancyClass []SwarmAndRedundancyClassV5 `json:"swarms_and_redundancy_class"`
	CubbitIngress            CubbitIngress               `json:"cubbit_ingress"`
}

type CreateGatewayV5Response struct {
	ID string `json:"id"`
}

type ListGatewaysV5Options struct {
	Page      int
	Items     int
	SortKey   string
	SortOrder string
	Filter    string
}

type ListGatewaysV5Option func(*ListGatewaysV5Options)

func WithPage(page int) ListGatewaysV5Option {
	return func(opts *ListGatewaysV5Options) {
		opts.Page = page
	}
}

func WithItems(items int) ListGatewaysV5Option {
	return func(opts *ListGatewaysV5Options) {
		opts.Items = items
	}
}

func WithSortKey(sortKey string) ListGatewaysV5Option {
	return func(opts *ListGatewaysV5Options) {
		opts.SortKey = sortKey
	}
}

func WithSortOrder(sortOrder string) ListGatewaysV5Option {
	return func(opts *ListGatewaysV5Options) {
		opts.SortOrder = sortOrder
	}
}

func WithFilter(filter string) ListGatewaysV5Option {
	return func(opts *ListGatewaysV5Options) {
		opts.Filter = filter
	}
}

type GatewayAPIInterface interface {
	CreateGatewayV5(
		urlConfig configuration.URLs,
		apiKey string,
		organizationID string,
		request *CreateGatewayV5Request,
	) (*CreateGatewayV5Response, error)
	GetGatewayV5(
		urlConfig configuration.URLs,
		apiKey string,
		organizationID string,
		gatewayID string,
	) (*GatewayV5GetResponse, error)
	ListGatewaysV5(
		urlConfig configuration.URLs,
		apiKey string,
		organizationID string,
		opts ...ListGatewaysV5Option,
	) (*GenericPaginatedResponse[GatewayV5ListItemResponse], error)
}

type GatewayAPI struct{}

func NewGatewayAPI() *GatewayAPI {
	return &GatewayAPI{}
}

func (api *GatewayAPI) CreateGatewayV5(
	urlConfig configuration.URLs,
	apiKey string,
	organizationID string,
	request *CreateGatewayV5Request,
) (*CreateGatewayV5Response, error) {
	url := NewURLBuilder(urlConfig.ChURL).
		Path("v5", "organizations", organizationID, "gateways").
		Build()

	var response CreateGatewayV5Response

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

func (api *GatewayAPI) GetGatewayV5(
	urlConfig configuration.URLs,
	apiKey string,
	organizationID string,
	gatewayID string,
) (*GatewayV5GetResponse, error) {
	url := NewURLBuilder(urlConfig.ChURL).
		Path("v5", "organizations", organizationID, "gateways", gatewayID).
		Build()

	var response GatewayV5GetResponse

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

func (api *GatewayAPI) ListGatewaysV5(
	urlConfig configuration.URLs,
	apiKey string,
	organizationID string,
	opts ...ListGatewaysV5Option,
) (*GenericPaginatedResponse[GatewayV5ListItemResponse], error) {
	options := &ListGatewaysV5Options{
		Page:  1,
		Items: 100,
	}

	for _, opt := range opts {
		opt(options)
	}

	urlBuilder := NewURLBuilder(urlConfig.ChURL).
		Path("v5", "organizations", organizationID, "gateways").
		QueryParamInt("page", options.Page).
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

	var response GenericPaginatedResponse[GatewayV5ListItemResponse]

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
