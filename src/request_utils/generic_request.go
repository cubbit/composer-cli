package request_utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/charmbracelet/lipgloss"
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

type Error struct {
	Message         string   `json:"message"`
	ActionsRequired []string `json:"actions_required"`
	Reason          string   `json:"reason"`
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
		var err Error
		if err := json.Unmarshal(body, &err); err != nil {
			return fmt.Errorf("error while unmarshaling the request response : %w", err)
		}

		keyStyle := lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("12"))

		valueStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("15"))
		redStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("9"))
		yellowStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("11"))

		var formattedError string

		formattedError += keyStyle.Render("INF ") + "code status expected " + yellowStyle.Render(fmt.Sprint(opt.status)) + ", but recieved " + redStyle.Render(fmt.Sprint(res.StatusCode)) + " instead\n"

		if err.Message != "" {
			formattedError += keyStyle.Render("INF ") + valueStyle.Render(err.Message) + "\n"
		}

		if len(err.ActionsRequired) > 0 {
			formattedError += keyStyle.Render("INF ") + "actions required [" + valueStyle.Render(strings.Join(err.ActionsRequired, ", ")) + "]\n"
		}

		if err.Reason != "" {
			formattedError += keyStyle.Render("INF ") + "reason" + valueStyle.Render(err.Reason) + "\n"
		}
		return fmt.Errorf(fmt.Sprintf("\n%s", formattedError))
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

		opt.headers[constants.Authorization] = constants.Bearer + " " + accessToken

		return nil
	}
}

func WithRefreshToken(refreshToken string) RequestModifier {
	return func(opt *RequestOptions, res *http.Response) error {
		if opt == nil {
			return nil
		}

		opt.headers["Cookie"] = constants.RefreshTokenName + "=" + refreshToken

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
