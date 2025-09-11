package request_utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/cubbit/cubbit/client/cli/constants"
)

type RequestBody = map[string]interface{}

type RequestOptions struct {
	method    string
	headers   map[string]string
	body      RequestBody
	bodyBytes []byte
	status    int
}
type RequestModifier = func(*RequestOptions, *http.Response) error

type Data struct {
	Params          []string      `json:"params" example:"param1,param2"`
	ActionsRequired []string      `json:"actions_required" example:"iam:AttachUserPolicy,iam:DetachUserPolicy"`
	Reason          string        `json:"reason" example:"policy_id"`
	IssueFound      []interface{} `json:"issue_found"`
}

type Error struct {
	Message         string   `json:"message"`
	ActionsRequired []string `json:"actions_required"`
	Reason          string   `json:"reason"`
	Data            Data     `json:"data"`
	Params          []string `json:"params"`
	Param           string   `json:"param"`
}

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

	var res *http.Response
	if res, err = http.DefaultClient.Do(req); err != nil {
		return fmt.Errorf("error while performing the request: %w", err)
	}
	defer res.Body.Close()
	if opt.status != -1 && res.StatusCode != opt.status {
		body, _ := ioutil.ReadAll(res.Body)
		var apiErr Error
		if err := json.Unmarshal(body, &apiErr); err != nil {
			return fmt.Errorf("error while unmarshaling '%s' the request response: %w", string(body), err)
		}

		var errorLines []string
		errorLines = append(errorLines, fmt.Sprintf("code status expected %d, but received %d instead", opt.status, res.StatusCode))

		if apiErr.Message != "" {
			errorLines = append(errorLines, apiErr.Message)
		}

		if len(apiErr.ActionsRequired) > 0 {
			errorLines = append(errorLines, fmt.Sprintf("actions required [%s]", strings.Join(apiErr.ActionsRequired, ", ")))
		}

		if len(apiErr.Data.ActionsRequired) > 0 {
			errorLines = append(errorLines, fmt.Sprintf("actions required [%s]", strings.Join(apiErr.Data.ActionsRequired, ", ")))
		}

		if apiErr.Reason != "" {
			errorLines = append(errorLines, fmt.Sprintf("reason %s", apiErr.Reason))
		}

		if len(apiErr.Params) > 0 {
			errorLines = append(errorLines, fmt.Sprintf("params [%s]", strings.Join(apiErr.Params, ", ")))
		}

		if len(apiErr.Data.IssueFound) > 0 {
			var issues []string
			for _, issue := range apiErr.Data.IssueFound {
				switch v := issue.(type) {
				case map[string]interface{}:
					var parts []string
					for k, val := range v {
						parts = append(parts, fmt.Sprintf("%s: %v", k, val))
					}
					issues = append(issues, strings.Join(parts, ", "))
				default:
					issues = append(issues, fmt.Sprintf("%v", v))
				}
			}
			errorLines = append(errorLines, fmt.Sprintf("issue found [%s]", strings.Join(issues, ", ")))
		}

		return fmt.Errorf(strings.Join(errorLines, "\n"))
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

		opt.headers[constants.Authorization] = constants.APIKey + " " + accessToken

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

func WithAttachement() RequestModifier {
	return func(opt *RequestOptions, res *http.Response) error {
		if opt == nil {
			return nil
		}

		opt.headers["accept"] = "text/csv"
		return nil
	}
}
