package api

import (
	"crypto/ed25519"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/src/request_utils"
	"github.com/cubbit/cubbit/client/cli/utils"
)

func GenerateOperatorChallenge(urls configuration.Url, email string) (*ChallengeResponseModel, error) {
	var err error
	var response ChallengeResponseModel
	url := urls.IamUrl + constants.GenerateOperatorChallenge

	requestBody := map[string]interface{}{
		"email": email,
	}

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPost),
		request_utils.WithRequestBody(requestBody),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, err
	}

	return &response, nil
}

func GenerateAccountChallenge(urls configuration.Url, email string) (*ChallengeResponseModel, error) {
	var err error
	var response ChallengeResponseModel
	url := urls.IamUrl + constants.GenerateAccountChallenge

	requestBody := map[string]interface{}{
		"email": email,
	}

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPost),
		request_utils.WithRequestBody(requestBody),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, err
	}

	return &response, nil
}

func PerformOperatorSignin(urls configuration.Url, email, password string, challenge *ChallengeResponseModel, code string) (string, error) {
	var err error
	url := urls.IamUrl + constants.OperatorSignIn
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
		ExtractGenericModel(&tokenExpirationResponse),
		extractRefreshCookie(&refreshTokenCookie),
	); err != nil {
		return "", err
	}

	if tokenExpirationResponse.Exp == 0 {
		return "", fmt.Errorf(constants.ErrorTokenExpiration)
	}

	if tokenExpirationResponse.Token == "" {
		return "", fmt.Errorf(constants.ErrorEmptyToken)
	}

	if refreshTokenCookie == "" {
		return "", fmt.Errorf(constants.ErrorRefreshToken)
	}

	return refreshTokenCookie, nil
}

func PerformAccountSignin(urls configuration.Url, tenantID string, email, password string, challenge *ChallengeResponseModel, code string) (string, error) {
	var err error
	url := urls.IamUrl + constants.AccountSignIn
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
		"tenant_id":        tenantID,
	}

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPost),
		request_utils.WithRequestBody(requestBody),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		ExtractGenericModel(&tokenExpirationResponse),
		extractRefreshCookie(&refreshTokenCookie),
	); err != nil {
		return "", err
	}

	if tokenExpirationResponse.Exp == 0 {
		return "", fmt.Errorf(constants.ErrorTokenExpiration)
	}

	if tokenExpirationResponse.Token == "" {
		return "", fmt.Errorf(constants.ErrorEmptyToken)
	}

	if refreshTokenCookie == "" {
		return "", fmt.Errorf(constants.ErrorRefreshToken)
	}

	return refreshTokenCookie, nil
}

func CreateOperator(urls configuration.Url, firstName, lastName, email, password, secret string) error {
	var err error

	var challenge *ChallengeResponseModel
	if challenge, err = GenerateOperatorChallenge(urls, email); err != nil {
		return err
	}

	hash := sha256.New()
	hash.Write([]byte(password + challenge.Salt))
	seed := hash.Sum(nil)

	var publicKey ed25519.PublicKey
	if publicKey, _, err = utils.GenerateKeyPairFromSeed(seed); err != nil {
		return err
	}

	url := urls.IamUrl + constants.CreateOperator + url.QueryEscape(secret)
	requestBody := map[string]interface{}{
		"first_name":                firstName,
		"last_name":                 lastName,
		"email":                     email,
		"authentication_public_key": base64.StdEncoding.EncodeToString(publicKey),
	}

	if err = request_utils.DoRequest(url, request_utils.WithRequestMethod(http.MethodPost), request_utils.WithRequestBody(requestBody), request_utils.WithExpectedStatusCode(http.StatusNoContent)); err != nil {
		return err
	}

	return nil
}

func CreateAccount(urls configuration.Url, firstName, lastName, email, password, tenantID string) error {
	var err error

	var challenge *ChallengeResponseModel
	if challenge, err = GenerateAccountChallenge(urls, email); err != nil {
		return err
	}

	hash := sha256.New()
	hash.Write([]byte(password + challenge.Salt))
	seed := hash.Sum(nil)

	var publicKey ed25519.PublicKey
	if publicKey, _, err = utils.GenerateKeyPairFromSeed(seed); err != nil {
		return err
	}

	url := urls.IamUrl + constants.CreateAccount
	requestBody := map[string]interface{}{
		"first_name":                firstName,
		"last_name":                 lastName,
		"email":                     email,
		"authentication_public_key": base64.StdEncoding.EncodeToString(publicKey),
		"tenant_id":                 tenantID,
	}

	if err = request_utils.DoRequest(url, request_utils.WithRequestMethod(http.MethodPost), request_utils.WithRequestBody(requestBody), request_utils.WithExpectedStatusCode(http.StatusNoContent)); err != nil {
		return err
	}

	return nil
}

func ForgeOperatorAccessToken(urls configuration.Url, refreshToken string) (string, string, error) {
	url := urls.IamUrl + constants.ForgeOperatorAccessToken

	return getAccessToken(refreshToken, url)
}

func RefreshAccountAccessToken(urls configuration.Url, refreshToken string) (string, string, error) {
	url := urls.IamUrl + constants.RefreshAccountAccessToken

	return getAccessToken(refreshToken, url)
}

func RefreshOperatorAccessToken(urls configuration.Url, refreshToken string) (string, string, error) {
	url := urls.IamUrl + constants.RefreshOperatorAccessToken

	return getAccessToken(refreshToken, url)
}

func getAccessToken(refreshToken string, url string) (string, string, error) {
	var err error
	var tokenExpirationResponse TokenAndExpirationResponseModel

	if err = request_utils.DoRequest(
		url,
		request_utils.WithExpectedStatusCode(http.StatusOK),
		request_utils.WithRefreshToken(refreshToken),
		ExtractGenericModel(&tokenExpirationResponse),
		extractRefreshCookie(&refreshToken),
	); err != nil {
		return "", "", err
	}

	if tokenExpirationResponse.Exp == 0 {
		return "", "", fmt.Errorf(constants.ErrorTokenExpiration)
	}

	if tokenExpirationResponse.Token == "" {
		return "", "", fmt.Errorf(constants.ErrorEmptyToken)
	}

	return tokenExpirationResponse.Token, refreshToken, nil
}

func ForgeOperatorDeleteTenantToken(urls configuration.Url, email, password, refreshToken string, challenge *ChallengeResponseModel, code, tenantID string) (string, error) {

	var err error
	var privateKey ed25519.PrivateKey
	var tokenExpirationResponse TokenAndExpirationResponseModel

	h := sha256.New()
	h.Write([]byte(password + challenge.Salt))

	seed := h.Sum(nil)

	if _, privateKey, err = utils.GenerateKeyPairFromSeed(seed); err != nil {
		return "", err
	}

	url := urls.IamUrl + constants.ForgeOperatorDeleteTenantToken + tenantID
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
		ExtractGenericModel(&tokenExpirationResponse),
	); err != nil {
		return "", err
	}

	if tokenExpirationResponse.Exp == 0 {
		return "", fmt.Errorf(constants.ErrorTokenExpiration)
	}

	if tokenExpirationResponse.Token == "" {
		return "", fmt.Errorf(constants.ErrorEmptyToken)
	}

	return tokenExpirationResponse.Token, nil
}

func ForgeOperatorDeleteSwarmToken(urls configuration.Url, email, password, refreshToken string, challenge *ChallengeResponseModel, code, swarmID string) (string, error) {

	var err error
	var privateKey ed25519.PrivateKey
	var tokenExpirationResponse TokenAndExpirationResponseModel

	h := sha256.New()
	h.Write([]byte(password + challenge.Salt))

	seed := h.Sum(nil)

	if _, privateKey, err = utils.GenerateKeyPairFromSeed(seed); err != nil {
		return "", err
	}

	url := urls.IamUrl + constants.ForgeOperatorDeleteSwarmToken + swarmID
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
		ExtractGenericModel(&tokenExpirationResponse),
	); err != nil {
		return "", err
	}

	if tokenExpirationResponse.Exp == 0 {
		return "", fmt.Errorf(constants.ErrorTokenExpiration)
	}

	if tokenExpirationResponse.Token == "" {
		return "", fmt.Errorf(constants.ErrorEmptyToken)
	}

	return tokenExpirationResponse.Token, nil
}

func ForgeDistributorDeleteToken(urls configuration.Url, email, password, refreshToken string, challenge *ChallengeResponseModel, code, distributor string) (string, error) {

	var err error
	var privateKey ed25519.PrivateKey
	var tokenExpirationResponse TokenAndExpirationResponseModel

	h := sha256.New()
	h.Write([]byte(password + challenge.Salt))

	seed := h.Sum(nil)

	if _, privateKey, err = utils.GenerateKeyPairFromSeed(seed); err != nil {
		return "", err
	}

	url := urls.IamUrl + constants.ForgeDistributorDeleteToken + distributor
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
		ExtractGenericModel(&tokenExpirationResponse),
	); err != nil {
		return "", err
	}

	if tokenExpirationResponse.Exp == 0 {
		return "", fmt.Errorf(constants.ErrorTokenExpiration)
	}

	if tokenExpirationResponse.Token == "" {
		return "", fmt.Errorf(constants.ErrorEmptyToken)
	}

	return tokenExpirationResponse.Token, nil
}

func ForgeOperatorDeleteTenantAccountToken(urls configuration.Url, email, password, refreshToken string, challenge *ChallengeResponseModel, code, accountID string) (string, error) {

	var err error
	var privateKey ed25519.PrivateKey
	var tokenExpirationResponse TokenAndExpirationResponseModel

	h := sha256.New()
	h.Write([]byte(password + challenge.Salt))

	seed := h.Sum(nil)

	if _, privateKey, err = utils.GenerateKeyPairFromSeed(seed); err != nil {
		return "", err
	}

	url := urls.IamUrl + constants.DeleteTenantAccountToken + accountID
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
		ExtractGenericModel(&tokenExpirationResponse),
	); err != nil {
		return "", err
	}

	if tokenExpirationResponse.Exp == 0 {
		return "", fmt.Errorf(constants.ErrorTokenExpiration)
	}

	if tokenExpirationResponse.Token == "" {
		return "", fmt.Errorf(constants.ErrorEmptyToken)
	}

	return tokenExpirationResponse.Token, nil
}

func ForgeOperatorDeleteTenantProjectToken(urls configuration.Url, email, password, refreshToken string, challenge *ChallengeResponseModel, code, projectID string) (string, error) {

	var err error
	var privateKey ed25519.PrivateKey
	var tokenExpirationResponse TokenAndExpirationResponseModel

	h := sha256.New()
	h.Write([]byte(password + challenge.Salt))

	seed := h.Sum(nil)

	if _, privateKey, err = utils.GenerateKeyPairFromSeed(seed); err != nil {
		return "", err
	}

	url := urls.IamUrl + constants.DeleteTenantProjectToken + projectID
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
		ExtractGenericModel(&tokenExpirationResponse),
	); err != nil {
		return "", err
	}

	if tokenExpirationResponse.Exp == 0 {
		return "", fmt.Errorf(constants.ErrorTokenExpiration)
	}

	if tokenExpirationResponse.Token == "" {
		return "", fmt.Errorf(constants.ErrorEmptyToken)
	}

	return tokenExpirationResponse.Token, nil
}

func ForgeOperatorSwarmNodeToken(urls configuration.Url, email, password, refreshToken string, challenge *ChallengeResponseModel, code, nodeID string) (string, error) {

	var err error
	var privateKey ed25519.PrivateKey
	var tokenExpirationResponse TokenAndExpirationResponseModel

	h := sha256.New()
	h.Write([]byte(password + challenge.Salt))

	seed := h.Sum(nil)

	if _, privateKey, err = utils.GenerateKeyPairFromSeed(seed); err != nil {
		return "", err
	}

	url := urls.IamUrl + constants.DeleteSwarmNodeToken + nodeID
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
		ExtractGenericModel(&tokenExpirationResponse),
	); err != nil {
		return "", err
	}

	if tokenExpirationResponse.Exp == 0 {
		return "", fmt.Errorf(constants.ErrorTokenExpiration)
	}

	if tokenExpirationResponse.Token == "" {
		return "", fmt.Errorf(constants.ErrorEmptyToken)
	}

	return tokenExpirationResponse.Token, nil
}

func ForgeOperatorDeleteTenantGatewayToken(urls configuration.Url, email, password, refreshToken string, challenge *ChallengeResponseModel, code, gatewayID string) (string, error) {

	var err error
	var privateKey ed25519.PrivateKey
	var tokenExpirationResponse TokenAndExpirationResponseModel

	h := sha256.New()
	h.Write([]byte(password + challenge.Salt))

	seed := h.Sum(nil)

	if _, privateKey, err = utils.GenerateKeyPairFromSeed(seed); err != nil {
		return "", err
	}

	url := urls.IamUrl + constants.DeleteTenantGatewayToken + gatewayID
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
		ExtractGenericModel(&tokenExpirationResponse),
	); err != nil {
		return "", err
	}

	if tokenExpirationResponse.Exp == 0 {
		return "", fmt.Errorf(constants.ErrorTokenExpiration)
	}

	if tokenExpirationResponse.Token == "" {
		return "", fmt.Errorf(constants.ErrorEmptyToken)
	}

	return tokenExpirationResponse.Token, nil
}
