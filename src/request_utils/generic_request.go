package request_utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	ACCESS_TOKEN_NAME  = "_access"
	REFRESH_TOKEN_NAME = "_refresh"
	AUTHORIZATION      = "Authorization"
	BEARER             = "Bearer"
)

const BASE_URL = "https://api.cubbit.eu/"
const BASE_KEYVAULT_URL = "keyvault"
const BASE_IAM_URL = "iam"

type RequestBody = map[string]interface{}

type RequestOptions struct {
	method    string
	headers   map[string]string
	body      RequestBody
	bodyBytes []byte
	status    int
}
type RequestModifier = func(*RequestOptions, *http.Response) error

func DoRequest(url string, opts ...RequestModifier) error {
	var err error
	var reqJSONBody []byte
	opt := &RequestOptions{
		method:    http.MethodGet,
		headers:   make(map[string]string),
		body:      make(RequestBody),
		bodyBytes: nil,
		status:    http.StatusOK,
	}

	for _, modifier := range opts {
		if err = modifier(opt, nil); err != nil {
			return fmt.Errorf("error while applying modifier to request: %w", err)
		}
	}

	if reqJSONBody, err = json.Marshal(opt.body); err != nil {
		return fmt.Errorf("error while marshalling request: %w", err)
	}

	reqBody := bytes.NewBuffer(reqJSONBody)

	if opt.bodyBytes != nil {
		reqBody = bytes.NewBuffer(opt.bodyBytes)
	}

	var req *http.Request
	if req, err = http.NewRequest(opt.method, url, reqBody); err != nil {
		return fmt.Errorf("error while creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.ContentLength = int64(reqBody.Len())

	for key, value := range opt.headers {
		req.Header.Set(key, value)
	}
	opt.headers["Cookie"] = "abuse_interstitial=e5fc-62-152-126-198.ngrok-free.app" + opt.headers["Cookie"]

	var res *http.Response
	if res, err = http.DefaultClient.Do(req); err != nil {
		return fmt.Errorf("error while performing the request: %w", err)
	}

	defer res.Body.Close()
	if opt.status != -1 && res.StatusCode != opt.status {
		return fmt.Errorf("error while performing the request status code expected %d but received %d instead", opt.status, res.StatusCode)
	}

	for _, modifier := range opts {
		if err = modifier(nil, res); err != nil {
			return fmt.Errorf("error while applying modifier after request: %w", err)
		}
	}

	return nil
}

func WithRequestMethod(method string) RequestModifier {
	return func(opt *RequestOptions, res *http.Response) error {
		if opt == nil {
			return nil
		}

		opt.method = method

		return nil
	}
}

func WithAccessToken(accessToken string) RequestModifier {
	return func(opt *RequestOptions, res *http.Response) error {
		if opt == nil {
			return nil
		}

		opt.headers[AUTHORIZATION] = BEARER + " " + accessToken

		return nil
	}
}

func WithRefreshToken(refreshToken string) RequestModifier {
	return func(opt *RequestOptions, res *http.Response) error {
		if opt == nil {
			return nil
		}

		opt.headers["Cookie"] = REFRESH_TOKEN_NAME + "=" + refreshToken

		return nil
	}
}

func WithRequestBody(body RequestBody) RequestModifier {
	return func(opt *RequestOptions, res *http.Response) error {
		if opt == nil {
			return nil
		}

		opt.body = body
		return nil
	}
}

func WithRequestBodyByte(body []byte) RequestModifier {
	return func(opt *RequestOptions, res *http.Response) error {
		if opt == nil {
			return nil
		}

		opt.bodyBytes = body
		return nil
	}
}

func WithRequestBodyObject(body any) RequestModifier {
	return func(opt *RequestOptions, res *http.Response) error {
		var reqJSONBody []byte
		var err error

		if opt == nil {
			return nil
		}

		if reqJSONBody, err = json.Marshal(body); err != nil {
			return err
		}

		opt.bodyBytes = reqJSONBody
		return nil
	}
}

func WithExpectedStatusCode(status int) RequestModifier {
	return func(opt *RequestOptions, res *http.Response) error {
		if opt == nil {
			return nil
		}

		opt.status = status
		return nil
	}
}
