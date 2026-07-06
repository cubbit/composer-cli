package api

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/src/request_utils"
	"github.com/cubbit/composer-cli/utils"
)

type AuthAPIInterface interface {
	Activate(
		endpoints configuration_models.EndpointsV2,
		token string,
	) error

	SignUp(
		endpoints configuration_models.EndpointsV2,
		email string,
		username string,
		firstName *string,
		lastName *string,
		authenticationPublicKey *string,
		organizationName string,
		organizationBasePolicy map[string]interface{},
		organizationSettings map[string]interface{},
	) error
	SignIn(
		endpoints configuration_models.EndpointsV2,
		username string,
		organization string,
		password string,
		tfaCode string,
	) (*SignInToken, error)
	GenerateChallenge(
		endpoints configuration_models.EndpointsV2,
		email *string,
		username *string,
		organizationName *string,
	) (*ChallengeResponseModel, error)
	ForgeToken(
		endpoints configuration_models.EndpointsV2,
		operatorID string,
		username string,
		organizationName string,
		password string,
		tfaCode string,
		tokenType string,
		token string,
		refreshToken string,
	) (string, error)
	CreateApiKey(
		endpoints configuration_models.EndpointsV2,
		operatorID string,
		name string,
		token string,
		forgeApiKeyToken string,
	) (string, error)
}

type AuthAPI struct{}

func NewAuthAPI() *AuthAPI {
	return &AuthAPI{}
}

func (api *AuthAPI) Activate(
	endpoints configuration_models.EndpointsV2,
	token string,
) error {
	url := NewURLBuilder(endpoints.IAM).
		Path("v1", "operators", "activate").
		QueryParam("token", token).
		Build()

	if err := request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodGet),
		request_utils.WithExpectedStatusCode(http.StatusNoContent),
	); err != nil {
		return fmt.Errorf("failed to perform activation request: %w", err)
	}

	return nil
}

func (api *AuthAPI) SignUp(
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
	url := NewURLBuilder(endpoints.IAM).
		Path("v2", "operators", "signup").
		Build()

	requestBody := CreateSignUpRequestBody(
		email,
		username,
		firstName,
		lastName,
		authenticationPublicKey,
		organizationName,
		organizationBasePolicy,
		organizationSettings,
	)

	if err := request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPost),
		request_utils.WithRequestBodyObject(requestBody),
		request_utils.WithExpectedStatusCode(http.StatusCreated),
	); err != nil {
		return fmt.Errorf("failed to perform sign up request: %w", err)
	}

	return nil
}

func (api *AuthAPI) SignIn(
	endpoints configuration_models.EndpointsV2,
	username string,
	organization string,
	password string,
	tfaCode string,
) (*SignInToken, error) {
	var err error
	url := NewURLBuilder(endpoints.IAM).
		Path("v3", "auth", "operators", "signin").
		Build()

	challenge, err := api.GenerateChallenge(
		endpoints,
		nil,
		&username,
		&organization,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to generate challenge: %w", err)
	}

	signedChallenge, err := utils.SignChallenge(
		challenge.Salt,
		challenge.Challenge,
		password,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to sign challenge: %w", err)
	}

	requestBody := map[string]interface{}{
		"username":          username,
		"organization_name": organization,
		"signed_challenge":  base64.StdEncoding.EncodeToString(signedChallenge),
		"tfa_code":          tfaCode,
	}

	var response TokenAndExpirationResponseModel
	var refreshToken string

	if err := request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPost),
		request_utils.WithRequestBodyObject(requestBody),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		ExtractGenericModel(&response),
		ExtractRefreshCookie(&refreshToken),
	); err != nil {
		return nil, fmt.Errorf("failed to perform sign in request: %w", err)
	}

	return &SignInToken{
		AccessToken:  response.Token,
		RefreshToken: refreshToken,
	}, nil
}

func (api *AuthAPI) GenerateChallenge(
	endpoints configuration_models.EndpointsV2,
	email *string,
	username *string,
	organizationName *string,
) (*ChallengeResponseModel, error) {
	url := NewURLBuilder(endpoints.IAM).
		Path("v3", "auth", "operators", "signin", "challenge").
		Build()

	body := CreateChallengeRequestBodyV3(
		email,
		username,
		organizationName,
	)

	var response ChallengeResponseModel

	if err := request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPost),
		request_utils.WithRequestBodyObject(body),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, fmt.Errorf("failed to generate challenge: %w", err)
	}

	return &response, nil
}

func (api *AuthAPI) ForgeToken(
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
	url := NewURLBuilder(endpoints.IAM).
		Path("v3", "auth", "operators", "forge", "token").
		QueryParam("capabilities", tokenType).
		QueryParam("operator_id", operatorID).
		Build()

	challenge, err := api.GenerateChallenge(
		endpoints,
		nil,
		&username,
		&organizationName,
	)

	if err != nil {
		return "", fmt.Errorf("failed to generate challenge: %w", err)
	}

	signedChallenge, err := utils.SignChallenge(
		challenge.Salt,
		challenge.Challenge,
		password,
	)

	if err != nil {
		return "", fmt.Errorf("failed to sign challenge: %w", err)
	}

	requestBody := map[string]interface{}{
		"username":          username,
		"organization_name": organizationName,
		"signed_challenge":  base64.StdEncoding.EncodeToString(signedChallenge),
		"tfa_code":          tfaCode,
	}

	var response TokenAndExpirationResponseModel

	if err := request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPost),
		request_utils.WithRequestBody(requestBody),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		request_utils.WithAccessToken(token),
		request_utils.WithRefreshToken(refreshToken),
		ExtractGenericModel(&response),
	); err != nil {
		return "", fmt.Errorf("failed to forge custom token key: %w", err)
	}

	return response.Token, nil
}

func (api *AuthAPI) CreateApiKey(
	endpoints configuration_models.EndpointsV2,
	operatorID string,
	name string,
	token string,
	forgeApiKeyToken string,
) (string, error) {
	url := NewURLBuilder(endpoints.IAM).
		Path("v1", "operators", operatorID, "api-keys").
		QueryParam("token", forgeApiKeyToken).
		Build()

	requestBody := map[string]string{
		"name": name,
	}

	var response OperatorAPIKey

	if err := request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPost),
		request_utils.WithRequestBodyObject(requestBody),
		request_utils.WithExpectedStatusCode(http.StatusCreated),
		request_utils.WithAccessToken(token),
		ExtractGenericModel(&response),
	); err != nil {
		return "", fmt.Errorf("failed to create API key: %w", err)
	}

	return response.Key, nil
}
