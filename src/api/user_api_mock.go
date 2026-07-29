package api

import "github.com/cubbit/composer-cli/src/configuration/configuration_models"

type MockUserAPI struct {
	GetIAMUserFunc           func(endpoints configuration_models.EndpointsV2, accessToken string, apiKey string, meOrID string) (*IAMUser, error)
	GetIAMUserSelfFunc       func(endpoints configuration_models.EndpointsV2, accessToken string, apiKey string) (*IAMUser, error)
	GetIAMUserSelfV3Func     func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string) (*IAMUser, error)
	GetIAMUserByIDFunc       func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, userID string) (*IAMUser, error)
	PromoteIAMUserFunc       func(endpoints configuration_models.EndpointsV2, email string, policyName string, secret string) error
	BulkCreateIAMUsersFunc   func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, request *BulkCreateIAMUsersRequestBody) (*BulkCreateIAMUsersResponse, error)
	BulkGenerateSaltsFunc    func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, request *BulkGenerateSaltsRequestBody) (*BulkGenerateSaltsResponse, error)
	ListIAMUsersFunc         func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, enabled *bool, search string, page int, items int, sortKey string, sortOrder string) (*GenericPaginatedResponse[IAMUserListItem], error)
	DeleteIAMUserFunc        func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, userID string, deleteToken string) error
	UpdateIAMUserFunc        func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, userID string, request *UpdateIAMUserRequestBody) (*IAMUser, error)
	ResetIAMUserPasswordFunc func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, userID string, request *ResetIAMUserPasswordRequestBody) error
	CreateIAMAPIKeyFunc      func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, operatorID string, request *CreateIAMAPIKeyRequestBody) (*OperatorAPIKey, error)
	ListIAMAPIKeysFunc       func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, operatorID string, page int, items int, sortKey string, sortOrder string) (*GenericPaginatedResponse[OperatorAPIKey], error)
	GetIAMAPIKeyByIDFunc     func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, operatorID string, apiKeyID string) (*OperatorAPIKey, error)
	UpdateIAMAPIKeyFunc      func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, operatorID string, apiKeyID string, request *UpdateIAMAPIKeyRequestBody) (*OperatorAPIKey, error)
	DeleteIAMAPIKeyFunc      func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, operatorID string, apiKeyID string) error
}

func (m *MockUserAPI) GetIAMUserByID(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, userID string) (*IAMUser, error) {
	if m.GetIAMUserByIDFunc != nil {
		return m.GetIAMUserByIDFunc(endpoints, apiKey, organizationID, userID)
	}

	return &IAMUser{}, nil
}

func (m *MockUserAPI) GetIAMUser(endpoints configuration_models.EndpointsV2, accessToken string, apiKey string, meOrID string) (*IAMUser, error) {
	if m.GetIAMUserFunc != nil {
		return m.GetIAMUserFunc(endpoints, accessToken, apiKey, meOrID)
	}

	return &IAMUser{}, nil
}

func (m *MockUserAPI) GetIAMUserSelf(endpoints configuration_models.EndpointsV2, accessToken string, apiKey string) (*IAMUser, error) {
	if m.GetIAMUserSelfFunc != nil {
		return m.GetIAMUserSelfFunc(endpoints, accessToken, apiKey)
	}

	return &IAMUser{}, nil
}

func (m *MockUserAPI) GetIAMUserSelfV3(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string) (*IAMUser, error) {
	if m.GetIAMUserSelfV3Func != nil {
		return m.GetIAMUserSelfV3Func(endpoints, apiKey, organizationID)
	}

	return &IAMUser{}, nil
}

func (m *MockUserAPI) PromoteIAMUser(endpoints configuration_models.EndpointsV2, email string, policyName string, secret string) error {
	if m.PromoteIAMUserFunc != nil {
		return m.PromoteIAMUserFunc(endpoints, email, policyName, secret)
	}

	return nil
}

func (m *MockUserAPI) BulkCreateIAMUsers(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, request *BulkCreateIAMUsersRequestBody) (*BulkCreateIAMUsersResponse, error) {
	if m.BulkCreateIAMUsersFunc != nil {
		return m.BulkCreateIAMUsersFunc(endpoints, apiKey, organizationID, request)
	}

	return &BulkCreateIAMUsersResponse{}, nil
}

func (m *MockUserAPI) BulkGenerateSalts(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, request *BulkGenerateSaltsRequestBody) (*BulkGenerateSaltsResponse, error) {
	if m.BulkGenerateSaltsFunc != nil {
		return m.BulkGenerateSaltsFunc(endpoints, apiKey, organizationID, request)
	}

	return &BulkGenerateSaltsResponse{}, nil
}

func (m *MockUserAPI) ListIAMUsers(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, enabled *bool, search string, page int, items int, sortKey string, sortOrder string) (*GenericPaginatedResponse[IAMUserListItem], error) {
	if m.ListIAMUsersFunc != nil {
		return m.ListIAMUsersFunc(endpoints, apiKey, organizationID, enabled, search, page, items, sortKey, sortOrder)
	}

	return &GenericPaginatedResponse[IAMUserListItem]{}, nil
}

func (m *MockUserAPI) DeleteIAMUser(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, userID string, deleteToken string) error {
	if m.DeleteIAMUserFunc != nil {
		return m.DeleteIAMUserFunc(endpoints, apiKey, organizationID, userID, deleteToken)
	}

	return nil
}

func (m *MockUserAPI) UpdateIAMUser(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, userID string, request *UpdateIAMUserRequestBody) (*IAMUser, error) {
	if m.UpdateIAMUserFunc != nil {
		return m.UpdateIAMUserFunc(endpoints, apiKey, organizationID, userID, request)
	}

	return &IAMUser{}, nil
}

func (m *MockUserAPI) ResetIAMUserPassword(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, userID string, request *ResetIAMUserPasswordRequestBody) error {
	if m.ResetIAMUserPasswordFunc != nil {
		return m.ResetIAMUserPasswordFunc(endpoints, apiKey, organizationID, userID, request)
	}

	return nil
}

func (m *MockUserAPI) CreateIAMAPIKey(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, operatorID string, request *CreateIAMAPIKeyRequestBody) (*OperatorAPIKey, error) {
	if m.CreateIAMAPIKeyFunc != nil {
		return m.CreateIAMAPIKeyFunc(endpoints, apiKey, organizationID, operatorID, request)
	}

	return &OperatorAPIKey{}, nil
}

func (m *MockUserAPI) ListIAMAPIKeys(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, operatorID string, page int, items int, sortKey string, sortOrder string) (*GenericPaginatedResponse[OperatorAPIKey], error) {
	if m.ListIAMAPIKeysFunc != nil {
		return m.ListIAMAPIKeysFunc(endpoints, apiKey, organizationID, operatorID, page, items, sortKey, sortOrder)
	}

	return &GenericPaginatedResponse[OperatorAPIKey]{}, nil
}

func (m *MockUserAPI) GetIAMAPIKeyByID(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, operatorID string, apiKeyID string) (*OperatorAPIKey, error) {
	if m.GetIAMAPIKeyByIDFunc != nil {
		return m.GetIAMAPIKeyByIDFunc(endpoints, apiKey, organizationID, operatorID, apiKeyID)
	}

	return &OperatorAPIKey{}, nil
}

func (m *MockUserAPI) UpdateIAMAPIKey(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, operatorID string, apiKeyID string, request *UpdateIAMAPIKeyRequestBody) (*OperatorAPIKey, error) {
	if m.UpdateIAMAPIKeyFunc != nil {
		return m.UpdateIAMAPIKeyFunc(endpoints, apiKey, organizationID, operatorID, apiKeyID, request)
	}

	return &OperatorAPIKey{}, nil
}

func (m *MockUserAPI) DeleteIAMAPIKey(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, operatorID string, apiKeyID string) error {
	if m.DeleteIAMAPIKeyFunc != nil {
		return m.DeleteIAMAPIKeyFunc(endpoints, apiKey, organizationID, operatorID, apiKeyID)
	}

	return nil
}

var _ UserAPIInterface = (*MockUserAPI)(nil)
