package api

import (
	"fmt"
	"net/http"

	"github.com/cubbit/composer-cli/src/configuration"
	"github.com/cubbit/composer-cli/src/request_utils"
)

type AuthAPIInterface interface {
	SignUp(
		urlConfig configuration.URLs,
		email string,
		username string,
		firstName *string,
		lastName *string,
		authenticationPublicKey *string,
		organizationName string,
		organizationBasePolicy map[string]interface{},
		organizationSettings map[string]interface{},
	) error
	GenerateChallenge(
		urlConfig configuration.URLs,
		email *string,
		username *string,
		organizationName *string,
	) (*ChallengeResponseModel, error)
}

type AuthAPI struct {
	config configuration.Config
}

func NewAuthAPI(config *configuration.Config) *AuthAPI {
	return &AuthAPI{
		config: *config,
	}
}

func (api *AuthAPI) SignUp(
	urlConfig configuration.URLs,
	email string,
	username string,
	firstName *string,
	lastName *string,
	authenticationPublicKey *string,
	organizationName string,
	organizationBasePolicy map[string]interface{},
	organizationSettings map[string]interface{},
) error {
	url := NewURLBuilder(urlConfig.IamURL).
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

func (api *AuthAPI) GenerateChallenge(
	urlConfig configuration.URLs,
	email *string,
	username *string,
	organizationName *string,
) (*ChallengeResponseModel, error) {
	url := NewURLBuilder(urlConfig.IamURL).
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
