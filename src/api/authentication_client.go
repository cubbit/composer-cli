package api

import (
	"crypto/ed25519"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"

	"github.com/cubbit/cubbit/client/cli/src/request_utils"
	"github.com/cubbit/cubbit/client/cli/utils"
)

func GenerateOperatorChallenge(apiServerUrl, email string) (*ChallengeResponseModel, error) {
	var err error
	var response ChallengeResponseModel
	url := apiServerUrl + "/v1/auth/operators/signin/challenge"

	requestBody := map[string]interface{}{
		"email": email,
	}

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPost),
		request_utils.WithRequestBody(requestBody),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		extractChallengeResponseModel(&response),
	); err != nil {
		return nil, fmt.Errorf("failed unable to create get operator request: %w", err)
	}

	return &response, nil
}

func PerformOperatorSignin(apiServerUrl, email, password string, challenge *ChallengeResponseModel, code string) (string, error) {
	var err error
	url := apiServerUrl + "/v1/auth/operators/signin"
	var privateKey ed25519.PrivateKey
	var tokenExpirationResponse TokenAndExpirationResponseModel
	var refreshTokenCookie string

	h := sha256.New()
	h.Write([]byte(password + challenge.Salt))
	seed := h.Sum(nil)

	if _, privateKey, err = utils.GenerateKeyPairFromSeed(seed); err != nil {
		return "", err
	}

	signedChallenge := ed25519.Sign(privateKey, []byte(challenge.Challenge))

	requestBody := map[string]interface{}{
		"email":            email,
		"signed_challenge": base64.StdEncoding.EncodeToString(signedChallenge),
		"tfa_code":         code,
	}

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPost),
		request_utils.WithRequestBody(requestBody),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		extractTokenExpirationModel(&tokenExpirationResponse),
		extractRefreshCookie(&refreshTokenCookie),
	); err != nil {
		return "", fmt.Errorf("failed unable to sign in request: %w", err)
	}

	if tokenExpirationResponse.Exp == 0 {
		return "", fmt.Errorf("wrong token expiration returned")
	}

	if tokenExpirationResponse.Token == "" {
		return "", errors.New("access token cannot be empty")
	}

	if refreshTokenCookie == "" {
		return "", errors.New("refresh token cannot be empty")
	}

	return refreshTokenCookie, nil
}

func CreateOperator(apiServerUrl, firstName, lastName, email, password string) error {
	var err error

	var challenge *ChallengeResponseModel
	if challenge, err = GenerateOperatorChallenge(apiServerUrl, email); err != nil {
		return err
	}

	hash := sha256.New()
	hash.Write([]byte(password + challenge.Salt))
	seed := hash.Sum(nil)

	var publicKey ed25519.PublicKey
	if publicKey, _, err = utils.GenerateKeyPairFromSeed(seed); err != nil {
		return err
	}

	url := apiServerUrl + "/v1/operators/signup"
	requestBody := map[string]interface{}{
		"first_name":                firstName,
		"last_name":                 lastName,
		"email":                     email,
		"authentication_public_key": base64.StdEncoding.EncodeToString(publicKey),
	}

	if err = request_utils.DoRequest(url, request_utils.WithRequestMethod(http.MethodPost), request_utils.WithRequestBody(requestBody), request_utils.WithExpectedStatusCode(http.StatusNoContent)); err != nil {
		return fmt.Errorf("failed unable to create operator request: %w", err)
	}

	return nil
}

func ForgeOperatorAccessToken(apiServerUrl, refreshToken string) (string, string, error) {
	url := apiServerUrl + "/v1/auth/operators/forge/access"

	return getOperatorAccessToken(refreshToken, url)
}

func RefreshAccessToken(apiServerUrl, refreshToken string) (string, string, error) {
	url := apiServerUrl + "/v1/auth/operators/refresh/access"

	return getOperatorAccessToken(refreshToken, url)
}

func getOperatorAccessToken(refreshToken string, url string) (string, string, error) {
	var err error
	var tokenExpirationResponse TokenAndExpirationResponseModel

	if err = request_utils.DoRequest(
		url,
		request_utils.WithExpectedStatusCode(http.StatusOK),
		request_utils.WithRefreshToken(refreshToken),
		extractTokenExpirationModel(&tokenExpirationResponse),
		extractRefreshCookie(&refreshToken),
	); err != nil {
		return "", "", fmt.Errorf("failed unable to forge token request: %w", err)
	}

	if tokenExpirationResponse.Exp == 0 {
		return "", "", fmt.Errorf("wrong token expiration returned xxx")
	}

	if tokenExpirationResponse.Token == "" {
		return "", "", errors.New("access token cannot be empty")
	}

	return tokenExpirationResponse.Token, refreshToken, nil
}

func ForgeOperatorDeleteTenantToken(apiServerUrl, email, password, refreshToken string, challenge *ChallengeResponseModel, code, tenantID string) (string, error) {

	var err error
	var privateKey ed25519.PrivateKey
	var tokenExpirationResponse TokenAndExpirationResponseModel

	h := sha256.New()
	h.Write([]byte(password + challenge.Salt))

	seed := h.Sum(nil)

	if _, privateKey, err = utils.GenerateKeyPairFromSeed(seed); err != nil {
		return "", err
	}

	url := apiServerUrl + "/v1/auth/operators/forge/token?capabilities=delete_tenant&tenant_id=" + tenantID
	signedChallenge := ed25519.Sign(privateKey, []byte(challenge.Challenge))

	body := map[string]interface{}{
		"email":            email,
		"signed_challenge": base64.StdEncoding.EncodeToString(signedChallenge),
		"tfa_code":         code,
	}

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPost),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		request_utils.WithRefreshToken(refreshToken),
		request_utils.WithRequestBody(body),
		extractTokenExpirationModel(&tokenExpirationResponse),
	); err != nil {
		return "", fmt.Errorf("failed unable to forge token request: %w", err)
	}

	if tokenExpirationResponse.Exp == 0 {
		return "", fmt.Errorf("wrong token expiration returned")
	}

	if tokenExpirationResponse.Token == "" {
		return "", errors.New("access token cannot be empty")
	}

	return tokenExpirationResponse.Token, nil
}
