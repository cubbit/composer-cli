package api

import (
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
)

type MockAuthAPI struct {
	ActivateFunc         func(endpoints configuration_models.EndpointsV2, token string) error
	SignUpFunc           func(endpoints configuration_models.EndpointsV2, email string, username string, firstName *string, lastName *string, authenticationPublicKey *string, organizationName string, organizationBasePolicy map[string]interface{}, organizationSettings map[string]interface{}) error
	SignInFunc           func(endpoints configuration_models.EndpointsV2, username string, organization string, password string, tfaCode string) (*SignInToken, error)
	GenerateChallengeFunc func(endpoints configuration_models.EndpointsV2, email *string, username *string, organizationName *string) (*ChallengeResponseModel, error)
	ForgeTokenFunc       func(endpoints configuration_models.EndpointsV2, operatorID string, username string, organizationName string, password string, tfaCode string, tokenType string, token string, refreshToken string) (string, error)
	CreateApiKeyFunc     func(endpoints configuration_models.EndpointsV2, operatorID string, name string, token string, forgeApiKeyToken string) (string, error)
}

func (m *MockAuthAPI) Activate(endpoints configuration_models.EndpointsV2, token string) error {
	if m.ActivateFunc != nil {
		return m.ActivateFunc(endpoints, token)
	}
	return nil
}

func (m *MockAuthAPI) SignUp(
	endpoints configuration_models.EndpointsV2,
	email string,
	username string,
	firstName *string,
	lastName *string,
	authenticationPublicKey *string,
	organizationName string,
	organizationBasePolicy map[string]interface{},
	organizationSettings map[string]interface{},
) error {
	if m.SignUpFunc != nil {
		return m.SignUpFunc(endpoints, email, username, firstName, lastName, authenticationPublicKey, organizationName, organizationBasePolicy, organizationSettings)
	}
	return nil
}

func (m *MockAuthAPI) SignIn(
	endpoints configuration_models.EndpointsV2,
	username string,
	organization string,
	password string,
	tfaCode string,
) (*SignInToken, error) {
	if m.SignInFunc != nil {
		return m.SignInFunc(endpoints, username, organization, password, tfaCode)
	}
	return &SignInToken{}, nil
}

func (m *MockAuthAPI) GenerateChallenge(
	endpoints configuration_models.EndpointsV2,
	email *string,
	username *string,
	organizationName *string,
) (*ChallengeResponseModel, error) {
	if m.GenerateChallengeFunc != nil {
		return m.GenerateChallengeFunc(endpoints, email, username, organizationName)
	}
	return &ChallengeResponseModel{}, nil
}

func (m *MockAuthAPI) ForgeToken(
	endpoints configuration_models.EndpointsV2,
	operatorID string,
	username string,
	organizationName string,
	password string,
	tfaCode string,
	tokenType string,
	token string,
	refreshToken string,
) (string, error) {
	if m.ForgeTokenFunc != nil {
		return m.ForgeTokenFunc(endpoints, operatorID, username, organizationName, password, tfaCode, tokenType, token, refreshToken)
	}
	return "", nil
}

func (m *MockAuthAPI) CreateApiKey(
	endpoints configuration_models.EndpointsV2,
	operatorID string,
	name string,
	token string,
	forgeApiKeyToken string,
) (string, error) {
	if m.CreateApiKeyFunc != nil {
		return m.CreateApiKeyFunc(endpoints, operatorID, name, token, forgeApiKeyToken)
	}
	return "", nil
}
