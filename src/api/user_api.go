package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/src/request_utils"
)

type UserAPIInterface interface {
	GetIAMUser(
		endpoints configuration_models.EndpointsV2,
		accessToken, apiKey string,
		meOrID string,
	) (*IAMUser, error)

	GetIAMUserSelf(
		endpoints configuration_models.EndpointsV2,
		accessToken, apiKey string,
	) (*IAMUser, error)

	GetIAMUserByID(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		userID string,
	) (*IAMUser, error)

	PromoteIAMUser(
		endpoints configuration_models.EndpointsV2,
		email,
		policyName,
		secret string,
	) error

	BulkCreateIAMUsers(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		request *BulkCreateIAMUsersRequestBody,
	) (*BulkCreateIAMUsersResponse, error)

	BulkGenerateSalts(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		request *BulkGenerateSaltsRequestBody,
	) (*BulkGenerateSaltsResponse, error)

	ListIAMUsers(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		enabled *bool,
		search string,
		page int,
		items int,
		sortKey string,
		sortOrder string,
	) (*GenericPaginatedResponse[IAMUserListItem], error)

	DeleteIAMUser(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		userID string,
		deleteToken string,
	) error

	UpdateIAMUser(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		userID string,
		request *UpdateIAMUserRequestBody,
	) (*IAMUser, error)

	CreateIAMAPIKey(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		operatorID string,
		request *CreateIAMAPIKeyRequestBody,
	) (*OperatorAPIKey, error)

	ListIAMAPIKeys(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		operatorID string,
		page int,
		items int,
		sortKey string,
		sortOrder string,
	) (*GenericPaginatedResponse[OperatorAPIKey], error)

	GetIAMAPIKeyByID(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		operatorID string,
		apiKeyID string,
	) (*OperatorAPIKey, error)

	UpdateIAMAPIKey(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		operatorID string,
		apiKeyID string,
		request *UpdateIAMAPIKeyRequestBody,
	) (*OperatorAPIKey, error)

	DeleteIAMAPIKey(
		endpoints configuration_models.EndpointsV2,
		apiKey string,
		organizationID string,
		operatorID string,
		apiKeyID string,
	) error
}

type UserAPI struct{}

func NewUserAPI() *UserAPI {
	return &UserAPI{}
}

func (a *UserAPI) GetIAMUserByID(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	userID string,
) (*IAMUser, error) {
	url := NewURLBuilder(endpoints.IAM).
		Path("v3", "organizations", organizationID, "operators", userID).
		Build()

	var operator IAMUser
	if err := request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodGet),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		request_utils.WithApiKey(apiKey),
		ExtractGenericModel(&operator),
	); err != nil {
		return nil, fmt.Errorf("failed to describe IAM user: %w", err)
	}

	return &operator, nil
}

func (a *UserAPI) GetIAMUser(endpoints configuration_models.EndpointsV2, accessToken, apiKey string, meOrID string) (*IAMUser, error) {
	var err error
	var operator IAMUser

	url := NewURLBuilder(endpoints.IAM).
		Path("v1", "operators", meOrID).
		Build()

	options := []request_utils.RequestModifier{
		ExtractGenericModel(&operator),
		request_utils.WithExpectedStatusCode(http.StatusOK),
	}

	if accessToken != "" {
		options = append(options, request_utils.WithAccessToken(accessToken))
	}
	if apiKey != "" {
		options = append(options, request_utils.WithApiKey(apiKey))
	}

	if err = request_utils.DoRequest(
		url,

		options...,
	); err != nil {
		return nil, err
	}

	return &operator, nil
}

func (a *UserAPI) GetIAMUserSelf(endpoints configuration_models.EndpointsV2, accessToken, apiKey string) (*IAMUser, error) {
	return a.GetIAMUser(endpoints, accessToken, apiKey, "me")
}

func (a *UserAPI) PromoteIAMUser(endpoints configuration_models.EndpointsV2, email, policyName, secret string) error {
	var err error

	url := NewURLBuilder(endpoints.IAM).
		Path("v1", "operators", "promote").
		Build()

	requestBody := map[string]interface{}{
		"email":       email,
		"policy_name": policyName,
		"secret":      secret,
	}

	if err = request_utils.DoRequest(url, request_utils.WithRequestMethod(http.MethodPost), request_utils.WithRequestBody(requestBody), request_utils.WithExpectedStatusCode(http.StatusCreated)); err != nil {
		return err
	}

	return nil
}

func (a *UserAPI) BulkCreateIAMUsers(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	request *BulkCreateIAMUsersRequestBody,
) (*BulkCreateIAMUsersResponse, error) {
	url := NewURLBuilder(endpoints.IAM).
		Path("v3", "organizations", organizationID, "operators", "bulk").
		Build()

	requestBody := map[string][]BulkCreateIAMUserRequestBody{
		"operators": request.Users,
	}

	var response BulkCreateIAMUsersResponse
	if err := request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPost),
		request_utils.WithExpectedStatusCode(http.StatusCreated),
		request_utils.WithApiKey(apiKey),
		request_utils.WithRequestBodyObject(requestBody),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, fmt.Errorf("failed to perform bulk create users request: %w", err)
	}

	return &response, nil
}

func (a *UserAPI) BulkGenerateSalts(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	request *BulkGenerateSaltsRequestBody,
) (*BulkGenerateSaltsResponse, error) {
	url := NewURLBuilder(endpoints.IAM).
		Path("v3", "organizations", organizationID, "operators", "salts", "bulk").
		Build()

	var response BulkGenerateSaltsResponse
	if err := request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPost),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		request_utils.WithApiKey(apiKey),
		request_utils.WithRequestBodyObject(request),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, fmt.Errorf("failed to bulk generate salts: %w", err)
	}

	return &response, nil
}

func (a *UserAPI) ListIAMUsers(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	enabled *bool,
	search string,
	page int,
	items int,
	sortKey string,
	sortOrder string,
) (*GenericPaginatedResponse[IAMUserListItem], error) {
	urlBuilder := NewURLBuilder(endpoints.IAM).
		Path("v3", "organizations", organizationID, "operators").
		QueryParamInt("page", page).
		QueryParamInt("items", items)

	if enabled != nil {
		urlBuilder.QueryParam("enabled", strconv.FormatBool(*enabled))
	}
	if search != "" {
		urlBuilder.QueryParam("search", search)
	}
	if sortKey != "" {
		urlBuilder.QueryParam("sort_key", sortKey)
	}
	if sortOrder != "" {
		urlBuilder.QueryParam("sort_order", sortOrder)
	}

	url := urlBuilder.Build()

	var response GenericPaginatedResponse[IAMUserListItem]
	if err := request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodGet),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		request_utils.WithApiKey(apiKey),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, fmt.Errorf("failed to list IAM users: %w", err)
	}

	return &response, nil
}

func (a *UserAPI) DeleteIAMUser(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	userID string,
	deleteToken string,
) error {
	url := NewURLBuilder(endpoints.IAM).
		Path("v3", "organizations", organizationID, "operators", userID).
		QueryParam("token", deleteToken).
		Build()

	if err := request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodDelete),
		request_utils.WithExpectedStatusCode(http.StatusNoContent),
		request_utils.WithApiKey(apiKey),
	); err != nil {
		return fmt.Errorf("failed to delete IAM user: %w", err)
	}

	return nil
}

func (a *UserAPI) UpdateIAMUser(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	userID string,
	request *UpdateIAMUserRequestBody,
) (*IAMUser, error) {
	url := NewURLBuilder(endpoints.IAM).
		Path("v3", "organizations", organizationID, "operators", userID).
		Build()

	var user IAMUser
	if err := request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPatch),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		request_utils.WithApiKey(apiKey),
		request_utils.WithRequestBodyObject(request),
		ExtractGenericModel(&user),
	); err != nil {
		return nil, fmt.Errorf("failed to update IAM user: %w", err)
	}

	return &user, nil
}

func (a *UserAPI) CreateIAMAPIKey(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	operatorID string,
	request *CreateIAMAPIKeyRequestBody,
) (*OperatorAPIKey, error) {
	url := NewURLBuilder(endpoints.IAM).
		Path("v3", "organizations", organizationID, "operators", operatorID, "api-keys").
		Build()

	var response OperatorAPIKey
	if err := request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPost),
		request_utils.WithExpectedStatusCode(http.StatusCreated),
		request_utils.WithApiKey(apiKey),
		request_utils.WithRequestBodyObject(request),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, fmt.Errorf("failed to create IAM API key: %w", err)
	}

	return &response, nil
}

func (a *UserAPI) ListIAMAPIKeys(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	operatorID string,
	page int,
	items int,
	sortKey string,
	sortOrder string,
) (*GenericPaginatedResponse[OperatorAPIKey], error) {
	urlBuilder := NewURLBuilder(endpoints.IAM).
		Path("v3", "organizations", organizationID, "operators", operatorID, "api-keys").
		QueryParamInt("page", page).
		QueryParamInt("items", items)

	if sortKey != "" {
		urlBuilder.QueryParam("sort_key", sortKey)
	}
	if sortOrder != "" {
		urlBuilder.QueryParam("sort_order", sortOrder)
	}

	url := urlBuilder.Build()

	var response GenericPaginatedResponse[OperatorAPIKey]
	if err := request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodGet),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		request_utils.WithApiKey(apiKey),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, fmt.Errorf("failed to list IAM API keys: %w", err)
	}

	return &response, nil
}

func (a *UserAPI) GetIAMAPIKeyByID(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	operatorID string,
	apiKeyID string,
) (*OperatorAPIKey, error) {
	url := NewURLBuilder(endpoints.IAM).
		Path("v3", "organizations", organizationID, "operators", operatorID, "api-keys", apiKeyID).
		Build()

	var response OperatorAPIKey
	if err := request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodGet),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		request_utils.WithApiKey(apiKey),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, fmt.Errorf("failed to describe IAM API key: %w", err)
	}

	return &response, nil
}

func (a *UserAPI) DeleteIAMAPIKey(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	operatorID string,
	apiKeyID string,
) error {
	url := NewURLBuilder(endpoints.IAM).
		Path("v3", "organizations", organizationID, "operators", operatorID, "api-keys", apiKeyID).
		Build()

	if err := request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodDelete),
		request_utils.WithExpectedStatusCode(http.StatusNoContent),
		request_utils.WithApiKey(apiKey),
	); err != nil {
		return fmt.Errorf("failed to delete IAM API key: %w", err)
	}

	return nil
}

func (a *UserAPI) UpdateIAMAPIKey(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	operatorID string,
	apiKeyID string,
	request *UpdateIAMAPIKeyRequestBody,
) (*OperatorAPIKey, error) {
	url := NewURLBuilder(endpoints.IAM).
		Path("v3", "organizations", organizationID, "operators", operatorID, "api-keys", apiKeyID).
		Build()

	var response OperatorAPIKey
	if err := request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPatch),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		request_utils.WithApiKey(apiKey),
		request_utils.WithRequestBodyObject(request),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, fmt.Errorf("failed to edit IAM API key: %w", err)
	}

	return &response, nil
}
