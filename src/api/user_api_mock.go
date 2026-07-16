package api

import "github.com/cubbit/composer-cli/src/configuration/configuration_models"

type MockUserAPI struct {
	GetIAMUserFunc         func(endpoints configuration_models.EndpointsV2, accessToken string, apiKey string, meOrID string) (*IAMUser, error)
	GetIAMUserSelfFunc     func(endpoints configuration_models.EndpointsV2, accessToken string, apiKey string) (*IAMUser, error)
	GetIAMUserByIDFunc     func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, userID string) (*IAMUser, error)
	PromoteIAMUserFunc     func(endpoints configuration_models.EndpointsV2, email string, policyName string, secret string) error
	BulkCreateIAMUsersFunc func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, request *BulkCreateIAMUsersRequestBody) (*BulkCreateIAMUsersResponse, error)
	ListIAMUsersFunc       func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, enabled *bool, search string, page int, items int, sortKey string, sortOrder string) (*GenericPaginatedResponse[IAMUserListItem], error)
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

func (m *MockUserAPI) ListIAMUsers(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, enabled *bool, search string, page int, items int, sortKey string, sortOrder string) (*GenericPaginatedResponse[IAMUserListItem], error) {
	if m.ListIAMUsersFunc != nil {
		return m.ListIAMUsersFunc(endpoints, apiKey, organizationID, enabled, search, page, items, sortKey, sortOrder)
	}

	return &GenericPaginatedResponse[IAMUserListItem]{}, nil
}

var _ UserAPIInterface = (*MockUserAPI)(nil)
