package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/src/request_utils"
)

type LocationProfileType string

const (
	LocationProfileAgent       LocationProfileType = "ProfileAgent"
	LocationProfileCoordinator LocationProfileType = "ProfileCoordinator"
	LocationProfileGateway     LocationProfileType = "ProfileGateway"
	LocationProfileNodeScanner LocationProfileType = "ProfileNodeScanner"
)

type LocationListOptions struct {
	ProfileType *LocationProfileFilter
}

type LocationProfileFilter struct {
	Op     string // "in" or "notin"
	Values []LocationProfileType
}

type LocationListOption func(*LocationListOptions)

func WithProfileType(op string, types ...LocationProfileType) LocationListOption {
	return func(opts *LocationListOptions) {
		opts.ProfileType = &LocationProfileFilter{Op: op, Values: types}
	}
}

func buildProfileTypeQuery(filter *LocationProfileFilter) string {
	values := make([]string, len(filter.Values))
	for i, v := range filter.Values {
		values[i] = string(v)
	}
	return fmt.Sprintf("profile-type:%s(%s)", filter.Op, strings.Join(values, ","))
}

type LocationAPIInterface interface {
	List(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		opts ...LocationListOption,
	) ([]InfrastructureCluster, error)
	ListAggregated(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		opts ...LocationListOption,
	) ([]InfraAggregateCluster, error)

	CreateVirtualCluster(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		name string,
		description *string,
	) (*InfrastructureCluster, error)

	CreateVirtualNode(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		clusterID string,
		name string,
		storageType string,
		configuration map[string]any,
	) (*InfraAggregateVirtualNodeDetail, error)

	GetAggregated(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		clusterID string,
	) (*InfraAggregateCluster, error)

	Update(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		clusterID string,
		request *UpdateInfraClusterRequest,
	) error

	Delete(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		clusterID string,
	) (*InfraClusterDeleteResponse, error)
}

type LocationAPI struct{}

func NewLocationAPI() *LocationAPI {
	return &LocationAPI{}
}

func (api *LocationAPI) List(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	opts ...LocationListOption,
) ([]InfrastructureCluster, error) {
	options := &LocationListOptions{}
	for _, opt := range opts {
		opt(options)
	}

	builder := NewURLBuilder(endpoints.CH).
		Path("v1", "organizations", organizationID, "infra", "clusters")

	if options.ProfileType != nil {
		builder = builder.QueryParam("q", buildProfileTypeQuery(options.ProfileType))
	}

	url := builder.Build()

	var response GenericPaginatedResponse[InfrastructureCluster]

	if err := request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodGet),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		request_utils.WithApiKey(apiKey),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (api *LocationAPI) ListAggregated(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	opts ...LocationListOption,
) ([]InfraAggregateCluster, error) {
	options := &LocationListOptions{}
	for _, opt := range opts {
		opt(options)
	}

	builder := NewURLBuilder(endpoints.CH).
		Path("v1", "organizations", organizationID, "infra", "aggregate_clusters")

	if options.ProfileType != nil {
		builder = builder.QueryParam("q", buildProfileTypeQuery(options.ProfileType))
	}

	url := builder.Build()

	var response GenericPaginatedResponse[InfraAggregateCluster]

	if err := request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodGet),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		request_utils.WithApiKey(apiKey),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (api *LocationAPI) CreateVirtualCluster(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	name string,
	description *string,
) (*InfrastructureCluster, error) {
	url := NewURLBuilder(endpoints.CH).
		Path("v1", "organizations", organizationID, "infra", "clusters", "virtual").
		Build()

	var response InfrastructureCluster

	if err := request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPost),
		request_utils.WithExpectedStatusCode(http.StatusCreated),
		request_utils.WithApiKey(apiKey),
		request_utils.WithRequestBody(map[string]interface{}{
			"name":        name,
			"description": description,
		}),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, err
	}

	return &response, nil
}

func (api *LocationAPI) CreateVirtualNode(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	clusterID string,
	name string,
	storageType string,
	configuration map[string]any,
) (*InfraAggregateVirtualNodeDetail, error) {
	url := NewURLBuilder(endpoints.CH).
		Path("v1", "organizations", organizationID, "infra", "clusters", "virtual", clusterID, "nodes").
		Build()

	var response InfraAggregateVirtualNodeDetail

	if err := request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPost),
		request_utils.WithExpectedStatusCode(http.StatusCreated),
		request_utils.WithApiKey(apiKey),
		request_utils.WithRequestBody(map[string]interface{}{
			"name":                  name,
			"storage_type":          storageType,
			"storage_configuration": configuration,
		}),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, err
	}

	return &response, nil
}

func (api *LocationAPI) GetAggregated(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	clusterID string,
) (*InfraAggregateCluster, error) {
	url := NewURLBuilder(endpoints.CH).
		Path("v1", "organizations", organizationID, "infra", "aggregate_clusters", clusterID).
		Build()

	status, body, err := doRawRequest(url, http.MethodGet, apiKey, nil)
	if err != nil {
		return nil, err
	}

	switch status {
	case http.StatusOK:
		var response InfraAggregateCluster
		if err := json.Unmarshal(body, &response); err != nil {
			return nil, fmt.Errorf("failed to parse location response: %w", err)
		}
		return &response, nil
	case http.StatusNotFound:
		return nil, &InfraClusterNotFoundError{ClusterID: clusterID}
	default:
		return nil, unexpectedStatusError(status, body)
	}
}

func (api *LocationAPI) Update(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	clusterID string,
	request *UpdateInfraClusterRequest,
) error {
	url := NewURLBuilder(endpoints.CH).
		Path("v1", "organizations", organizationID, "infra", "clusters", clusterID).
		Build()

	status, body, err := doRawRequest(url, http.MethodPatch, apiKey, request)
	if err != nil {
		return err
	}

	switch status {
	case http.StatusOK, http.StatusNoContent:
		return nil
	case http.StatusConflict:
		return parseInfraClusterConflict(body)
	case http.StatusNotFound:
		return &InfraClusterNotFoundError{ClusterID: clusterID}
	default:
		return unexpectedStatusError(status, body)
	}
}

func (api *LocationAPI) Delete(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	clusterID string,
) (*InfraClusterDeleteResponse, error) {
	url := NewURLBuilder(endpoints.CH).
		Path("v1", "organizations", organizationID, "infra", "clusters", clusterID).
		Build()

	status, body, err := doRawRequest(url, http.MethodDelete, apiKey, nil)
	if err != nil {
		return nil, err
	}

	switch status {
	case http.StatusAccepted:
		var response InfraClusterDeleteResponse
		if err := json.Unmarshal(body, &response); err != nil {
			return nil, fmt.Errorf("failed to parse cluster deletion response: %w", err)
		}
		return &response, nil
	case http.StatusConflict:
		return nil, parseInfraClusterConflict(body)
	case http.StatusNotFound:
		return nil, &InfraClusterNotFoundError{ClusterID: clusterID}
	default:
		return nil, unexpectedStatusError(status, body)
	}
}

// v2ErrorEnvelope is the Composer Hub V2 error envelope. Conflict operations report the reason
// in data.reason (with data.conflicts_with); the top-level reason/conflicts_with are kept as a
// fallback for the legacy envelope.
type v2ErrorEnvelope struct {
	Message       string              `json:"message"`
	Reason        string              `json:"reason"`
	ConflictsWith string              `json:"conflicts_with"`
	Data          v2ErrorConflictData `json:"data"`
}

type v2ErrorConflictData struct {
	Reason        string `json:"reason"`
	ConflictsWith string `json:"conflicts_with"`
	Param         string `json:"param"`
}

// parseInfraClusterConflict maps a 409 response body to a typed conflict error.
func parseInfraClusterConflict(body []byte) error {
	var envelope v2ErrorEnvelope
	if err := json.Unmarshal(body, &envelope); err != nil {
		return &InfraClusterConflictError{Reason: "conflict"}
	}

	reason := envelope.Data.Reason
	if reason == "" {
		reason = envelope.Reason
	}
	if reason == "" {
		reason = "conflict"
	}

	conflictsWith := envelope.Data.ConflictsWith
	if conflictsWith == "" {
		conflictsWith = envelope.ConflictsWith
	}

	return &InfraClusterConflictError{
		Reason:        reason,
		ConflictsWith: conflictsWith,
		Message:       envelope.Message,
	}
}

func unexpectedStatusError(status int, body []byte) error {
	var envelope v2ErrorEnvelope
	if err := json.Unmarshal(body, &envelope); err == nil && envelope.Message != "" {
		return fmt.Errorf("unexpected status %d: %s", status, envelope.Message)
	}

	if message := strings.TrimSpace(string(body)); message != "" {
		return fmt.Errorf("unexpected status %d: %s", status, message)
	}

	return fmt.Errorf("unexpected status %d", status)
}

// doRawRequest performs a request without the generic status/error flattening used by
// request_utils.DoRequest, so the caller can inspect the HTTP status code and the raw body.
// This is required to surface the 409 guard reasons (profile_installed / deletion_in_progress)
// and to distinguish the 202 accepted response of an asynchronous deletion.
func doRawRequest(url string, method string, apiKey string, body any) (int, []byte, error) {
	var status int
	var rawBody []byte

	modifiers := []request_utils.RequestModifier{
		request_utils.WithRequestMethod(method),
		request_utils.WithApiKey(apiKey),
		request_utils.WithExpectedStatusCode(request_utils.AnyStatusCode),
	}
	if body != nil {
		modifiers = append(modifiers, request_utils.WithRequestBodyObject(body))
	}
	modifiers = append(modifiers, func(opt *request_utils.RequestOptions, res *http.Response) error {
		if res == nil {
			return nil
		}

		status = res.StatusCode

		readBody, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		rawBody = readBody

		return nil
	})

	if err := request_utils.DoRequest(url, modifiers...); err != nil {
		return 0, nil, err
	}

	return status, rawBody, nil
}
