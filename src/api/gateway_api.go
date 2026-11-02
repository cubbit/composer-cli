package api

import (
	"encoding/json"
	"net/http"

	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
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

type UpdateGatewayV5Request struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Location    *string `json:"location,omitempty"`
}

// GatewayStorageClassStatusV5 is the status of a gateway storage class.
type GatewayStorageClassStatusV5 string

const (
	GatewayStorageClassStatusV5Active       GatewayStorageClassStatusV5 = "active"
	GatewayStorageClassStatusV5FallbackOnly GatewayStorageClassStatusV5 = "fallback_only"
)

// GatewayStorageClassV5ReferencedRule is a routing rule referencing a storage class.
type GatewayStorageClassV5ReferencedRule struct {
	ID   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
}

// GatewayStorageClassV5 is a default or custom storage class of a gateway.
type GatewayStorageClassV5 struct {
	ID   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
	// Pool reports whether the storage class is a default storage class.
	Pool            bool                                  `json:"pool" yaml:"pool"`
	Status          GatewayStorageClassStatusV5           `json:"status" yaml:"status"`
	ReferencedRules []GatewayStorageClassV5ReferencedRule `json:"referenced_rules" yaml:"referenced_rules"`
}

// GatewayStorageClassesV5ListResponse wraps the storage classes of a gateway.
type GatewayStorageClassesV5ListResponse struct {
	Data []GatewayStorageClassV5 `json:"data" yaml:"data"`
}

// CreateGatewayStorageClassV5Request is the body used to create a custom storage class.
type CreateGatewayStorageClassV5Request struct {
	Name string `json:"name" yaml:"name"`
}

// GatewayConnectionStatusV5 is the status of a gateway connection.
type GatewayConnectionStatusV5 string

const (
	GatewayConnectionStatusV5Active    GatewayConnectionStatusV5 = "active"
	GatewayConnectionStatusV5NotActive GatewayConnectionStatusV5 = "not_active"
)

// GatewayConnectionV5Swarm is the swarm a connection offloads to.
type GatewayConnectionV5Swarm struct {
	ID   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
}

// GatewayConnectionV5RedundancyClass is the redundancy class used by a connection.
type GatewayConnectionV5RedundancyClass struct {
	ID   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
}

// GatewayConnectionV5ImpactedRule is a routing rule referencing a connection.
type GatewayConnectionV5ImpactedRule struct {
	ID      string `json:"id" yaml:"id"`
	Name    string `json:"name" yaml:"name"`
	Enabled bool   `json:"enabled" yaml:"enabled"`
}

// GatewayConnectionV5 is a swarm and redundancy class pair a gateway can offload to.
type GatewayConnectionV5 struct {
	ID              string                             `json:"id" yaml:"id"`
	Swarm           GatewayConnectionV5Swarm           `json:"swarm" yaml:"swarm"`
	RedundancyClass GatewayConnectionV5RedundancyClass `json:"redundancy_class" yaml:"redundancy_class"`
	IsDefault       bool                               `json:"is_default" yaml:"is_default"`
	Status          GatewayConnectionStatusV5          `json:"status" yaml:"status"`
	ImpactedRules   []GatewayConnectionV5ImpactedRule  `json:"impacted_rules" yaml:"impacted_rules"`
}

// GatewayConnectionsV5ListResponse wraps the connections of a gateway.
type GatewayConnectionsV5ListResponse struct {
	Data []GatewayConnectionV5 `json:"data" yaml:"data"`
}

// CreateGatewayConnectionV5RequestItem describes one swarm and its redundancy classes.
type CreateGatewayConnectionV5RequestItem struct {
	SwarmID            string   `json:"swarm_id" yaml:"swarm_id"`
	RedundancyClassIDs []string `json:"redundancy_class_ids" yaml:"redundancy_class_ids"`
}

// CreateGatewayConnectionsV5Request is the body used to create gateway connections.
type CreateGatewayConnectionsV5Request struct {
	Connections []CreateGatewayConnectionV5RequestItem `json:"connections" yaml:"connections"`
}

// GatewayRoutingRuleV5Condition is one node of the SDP condition tree of a
// routing rule. The tree is recursive: group nodes carry an operator and
// children, leaf nodes carry a field, an operator and a value.
type GatewayRoutingRuleV5Condition struct {
	Field         string                           `json:"field" yaml:"field"`
	Operator      string                           `json:"operator" yaml:"operator"`
	Value         any                              `json:"value" yaml:"value"`
	CaseSensitive bool                             `json:"caseSensitive" yaml:"caseSensitive"`
	Conditions    []*GatewayRoutingRuleV5Condition `json:"conditions" yaml:"conditions"`
}

// GatewayRoutingRuleV5FieldMapping routes objects by the value of a field.
type GatewayRoutingRuleV5FieldMapping struct {
	Field                    string            `json:"field" yaml:"field"`
	Mapping                  map[string]string `json:"mapping" yaml:"mapping"`
	DefaultRedundancyClassID string            `json:"defaultRedundancyClassId" yaml:"defaultRedundancyClassId"`
}

// GatewayRoutingRuleV5LoadBalancing spreads objects across redundancy classes.
type GatewayRoutingRuleV5LoadBalancing struct {
	RedundancyClassIDs []string           `json:"redundancyClassIds" yaml:"redundancyClassIds"`
	Algorithm          string             `json:"algorithm" yaml:"algorithm"`
	Weights            map[string]float32 `json:"weights" yaml:"weights"`
}

// GatewayRoutingRuleV5Target is the placement target of a routing rule action.
type GatewayRoutingRuleV5Target struct {
	Type              string                             `json:"type" yaml:"type"`
	RedundancyClassID string                             `json:"redundancyClassId" yaml:"redundancyClassId"`
	FieldMapping      *GatewayRoutingRuleV5FieldMapping  `json:"fieldMapping" yaml:"fieldMapping"`
	LoadBalancing     *GatewayRoutingRuleV5LoadBalancing `json:"loadBalancing" yaml:"loadBalancing"`
}

// GatewayRoutingRuleV5OffloaderTarget offloads matching objects to an external
// endpoint.
type GatewayRoutingRuleV5OffloaderTarget struct {
	URL       string   `json:"url,omitempty" yaml:"url,omitempty"`
	URLs      []string `json:"urls,omitempty" yaml:"urls,omitempty"`
	Algorithm string   `json:"algorithm,omitempty" yaml:"algorithm,omitempty"`
}

// GatewayRoutingRuleV5Action is applied to the objects matching a routing rule.
type GatewayRoutingRuleV5Action struct {
	Type            string                               `json:"type" yaml:"type"`
	Target          *GatewayRoutingRuleV5Target          `json:"target" yaml:"target"`
	OffloaderTarget *GatewayRoutingRuleV5OffloaderTarget `json:"offloaderTarget,omitempty" yaml:"offloaderTarget,omitempty"`
}

// GatewayRoutingRuleV5 is a single smart data placement routing rule: one SDP
// policy attached to a gateway. A gateway can have multiple rules, evaluated in
// order.
type GatewayRoutingRuleV5 struct {
	ID                string                         `json:"id" yaml:"id"`
	Name              string                         `json:"name" yaml:"name"`
	Description       string                         `json:"description" yaml:"description"`
	Enabled           bool                           `json:"enabled" yaml:"enabled"`
	InputStorageClass string                         `json:"inputStorageClass" yaml:"inputStorageClass"`
	Conditions        *GatewayRoutingRuleV5Condition `json:"conditions" yaml:"conditions"`
	Action            GatewayRoutingRuleV5Action     `json:"action" yaml:"action"`
}

// GatewayRoutingRulesV5ListResponse is the bare array returned by the list
// endpoint.
type GatewayRoutingRulesV5ListResponse []GatewayRoutingRuleV5

type ListGatewaysV5Options struct {
	Page          int
	Items         int
	SortKey       string
	SortOrder     string
	Filter        string
	FetchAllPages bool
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

func WithListAllPages(fetchAll bool) ListGatewaysV5Option {
	return func(opts *ListGatewaysV5Options) {
		opts.FetchAllPages = fetchAll
	}
}

type GatewayAPIInterface interface {
	CreateGatewayV5(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		request *CreateGatewayV5Request,
	) (*CreateGatewayV5Response, error)
	GetGatewayV5(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		gatewayID string,
	) (*GatewayV5GetResponse, error)
	ListGatewaysV5(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		opts ...ListGatewaysV5Option,
	) (*GenericPaginatedResponse[GatewayV5ListItemResponse], error)
	UpdateGatewayV5(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		gatewayID string,
		request *UpdateGatewayV5Request,
	) (*GatewayV5GetResponse, error)
	ListGatewayStorageClassesV5(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		gatewayID string,
	) (*GatewayStorageClassesV5ListResponse, error)
	CreateGatewayStorageClassV5(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		gatewayID string,
		request *CreateGatewayStorageClassV5Request,
	) (*GatewayStorageClassV5, error)
	DeleteGatewayStorageClassV5(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		gatewayID string,
		storageClassID string,
	) error
	ListGatewayConnectionsV5(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		gatewayID string,
	) (*GatewayConnectionsV5ListResponse, error)
	CreateGatewayConnectionsV5(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		gatewayID string,
		request *CreateGatewayConnectionsV5Request,
	) (*GatewayConnectionsV5ListResponse, error)
	DeleteGatewayConnectionV5(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		gatewayID string,
		connectionID string,
	) error
	ListGatewayRoutingRulesV5(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		gatewayID string,
	) (*GatewayRoutingRulesV5ListResponse, error)
	GetGatewayRoutingRuleV5Raw(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		gatewayID string,
		ruleID string,
	) (json.RawMessage, error)
	CreateGatewayRoutingRuleV5(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		gatewayID string,
		body json.RawMessage,
	) (json.RawMessage, error)
	UpdateGatewayRoutingRuleV5(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		gatewayID string,
		ruleID string,
		body json.RawMessage,
	) (json.RawMessage, error)
	DeleteGatewayRoutingRuleV5(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		gatewayID string,
		ruleID string,
	) error
}

type GatewayAPI struct{}

func NewGatewayAPI() *GatewayAPI {
	return &GatewayAPI{}
}

func (api *GatewayAPI) CreateGatewayV5(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	request *CreateGatewayV5Request,
) (*CreateGatewayV5Response, error) {
	url := NewURLBuilder(endpoints.CH).
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
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	gatewayID string,
) (*GatewayV5GetResponse, error) {
	url := NewURLBuilder(endpoints.CH).
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

func (api *GatewayAPI) UpdateGatewayV5(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	gatewayID string,
	request *UpdateGatewayV5Request,
) (*GatewayV5GetResponse, error) {
	url := NewURLBuilder(endpoints.CH).
		Path("v5", "organizations", organizationID, "gateways", gatewayID).
		Build()

	var response GatewayV5GetResponse

	if err := request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPatch),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		request_utils.WithApiKey(apiKey),
		request_utils.WithRequestBodyObject(request),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, err
	}

	return &response, nil
}

func (api *GatewayAPI) ListGatewaysV5(
	endpoints configuration_models.EndpointsV2,
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

	allData := make([]GatewayV5ListItemResponse, 0)
	nextPage := &options.Page

	for nextPage != nil {
		urlBuilder := NewURLBuilder(endpoints.CH).
			Path("v5", "organizations", organizationID, "gateways").
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

		var pageResponse GenericPaginatedResponse[GatewayV5ListItemResponse]

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

	return &GenericPaginatedResponse[GatewayV5ListItemResponse]{
		Data:     allData,
		NextPage: nil,
		Count:    count,
	}, nil
}

func (api *GatewayAPI) ListGatewayStorageClassesV5(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	gatewayID string,
) (*GatewayStorageClassesV5ListResponse, error) {
	url := NewURLBuilder(endpoints.CH).
		Path("v5", "organizations", organizationID, "gateways", gatewayID, "storage-classes").
		Build()

	var response GatewayStorageClassesV5ListResponse

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

func (api *GatewayAPI) CreateGatewayStorageClassV5(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	gatewayID string,
	request *CreateGatewayStorageClassV5Request,
) (*GatewayStorageClassV5, error) {
	url := NewURLBuilder(endpoints.CH).
		Path("v5", "organizations", organizationID, "gateways", gatewayID, "storage-classes").
		Build()

	var response GatewayStorageClassV5

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

func (api *GatewayAPI) DeleteGatewayStorageClassV5(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	gatewayID string,
	storageClassID string,
) error {
	url := NewURLBuilder(endpoints.CH).
		Path("v5", "organizations", organizationID, "gateways", gatewayID, "storage-classes", storageClassID).
		Build()

	return request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodDelete),
		request_utils.WithExpectedStatusCode(http.StatusNoContent),
		request_utils.WithApiKey(apiKey),
	)
}

func (api *GatewayAPI) ListGatewayConnectionsV5(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	gatewayID string,
) (*GatewayConnectionsV5ListResponse, error) {
	url := NewURLBuilder(endpoints.CH).
		Path("v5", "organizations", organizationID, "gateways", gatewayID, "connections").
		Build()

	var response GatewayConnectionsV5ListResponse

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

func (api *GatewayAPI) CreateGatewayConnectionsV5(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	gatewayID string,
	request *CreateGatewayConnectionsV5Request,
) (*GatewayConnectionsV5ListResponse, error) {
	url := NewURLBuilder(endpoints.CH).
		Path("v5", "organizations", organizationID, "gateways", gatewayID, "connections").
		Build()

	var response GatewayConnectionsV5ListResponse

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

func (api *GatewayAPI) DeleteGatewayConnectionV5(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	gatewayID string,
	connectionID string,
) error {
	url := NewURLBuilder(endpoints.CH).
		Path("v5", "organizations", organizationID, "gateways", gatewayID, "connections", connectionID).
		Build()

	return request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodDelete),
		request_utils.WithExpectedStatusCode(http.StatusNoContent),
		request_utils.WithApiKey(apiKey),
	)
}

func (api *GatewayAPI) ListGatewayRoutingRulesV5(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	gatewayID string,
) (*GatewayRoutingRulesV5ListResponse, error) {
	url := NewURLBuilder(endpoints.CH).
		Path("v5", "organizations", organizationID, "gateways", gatewayID, "routing-rules").
		Build()

	var response GatewayRoutingRulesV5ListResponse

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

// GetGatewayRoutingRuleV5Raw returns a routing rule as the raw JSON document
// sent by the server, preserving fields this CLI does not model.
func (api *GatewayAPI) GetGatewayRoutingRuleV5Raw(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	gatewayID string,
	ruleID string,
) (json.RawMessage, error) {
	url := NewURLBuilder(endpoints.CH).
		Path("v5", "organizations", organizationID, "gateways", gatewayID, "routing-rules", ruleID).
		Build()

	var response json.RawMessage

	if err := request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodGet),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		request_utils.WithApiKey(apiKey),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, err
	}

	return response, nil
}

// CreateGatewayRoutingRuleV5 creates a routing rule from the given raw JSON
// document. The raw body is forwarded verbatim so fields this CLI does not
// model are preserved. The response is returned as raw JSON as well.
func (api *GatewayAPI) CreateGatewayRoutingRuleV5(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	gatewayID string,
	body json.RawMessage,
) (json.RawMessage, error) {
	url := NewURLBuilder(endpoints.CH).
		Path("v5", "organizations", organizationID, "gateways", gatewayID, "routing-rules").
		Build()

	var response json.RawMessage

	if err := request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPost),
		request_utils.WithExpectedStatusCode(http.StatusCreated),
		request_utils.WithApiKey(apiKey),
		request_utils.WithRequestBodyObject(body),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, err
	}

	return response, nil
}

// UpdateGatewayRoutingRuleV5 replaces a routing rule with the given raw JSON
// document. The raw body is forwarded verbatim so unmodeled fields obtained
// from get survive the round-trip. The response is returned as raw JSON.
func (api *GatewayAPI) UpdateGatewayRoutingRuleV5(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	gatewayID string,
	ruleID string,
	body json.RawMessage,
) (json.RawMessage, error) {
	url := NewURLBuilder(endpoints.CH).
		Path("v5", "organizations", organizationID, "gateways", gatewayID, "routing-rules", ruleID).
		Build()

	var response json.RawMessage

	if err := request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPut),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		request_utils.WithApiKey(apiKey),
		request_utils.WithRequestBodyObject(body),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, err
	}

	return response, nil
}

// DeleteGatewayRoutingRuleV5 deletes a routing rule from a gateway.
func (api *GatewayAPI) DeleteGatewayRoutingRuleV5(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	gatewayID string,
	ruleID string,
) error {
	url := NewURLBuilder(endpoints.CH).
		Path("v5", "organizations", organizationID, "gateways", gatewayID, "routing-rules", ruleID).
		Build()

	return request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodDelete),
		request_utils.WithExpectedStatusCode(http.StatusNoContent),
		request_utils.WithApiKey(apiKey),
	)
}
