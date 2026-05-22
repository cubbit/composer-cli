package api

import (
	"net/http"
	"strings"

	"github.com/cubbit/composer-cli/src/configuration"
	"github.com/cubbit/composer-cli/src/request_utils"
)

type ListProcessesOptions struct {
	ProcessType   ProcessType
	ProcessStatus ProcessStatus
	Step          ProcessStep
	OwnerID       string
}

type ListProcessesOption func(*ListProcessesOptions)

func WithProcessType(t ProcessType) ListProcessesOption {
	return func(o *ListProcessesOptions) { o.ProcessType = t }
}

func WithProcessStatus(s ProcessStatus) ListProcessesOption {
	return func(o *ListProcessesOptions) { o.ProcessStatus = s }
}

func WithProcessStep(s ProcessStep) ListProcessesOption {
	return func(o *ListProcessesOptions) { o.Step = s }
}

func WithOwnerID(id string) ListProcessesOption {
	return func(o *ListProcessesOptions) { o.OwnerID = id }
}

type ProcessAPIInterface interface {
	ListProcesses(
		urlConfig configuration.URLs,
		apiKey string,
		organizationID string,
		opts ...ListProcessesOption,
	) ([]Process, error)
	GetProcess(
		urlConfig configuration.URLs,
		apiKey string,
		organizationID string,
		processID string,
	) (*Process, error)
}

type ProcessAPI struct{}

func NewProcessAPI() *ProcessAPI {
	return &ProcessAPI{}
}

func (api *ProcessAPI) ListProcesses(
	urlConfig configuration.URLs,
	apiKey string,
	organizationID string,
	opts ...ListProcessesOption,
) ([]Process, error) {
	var options ListProcessesOptions
	for _, opt := range opts {
		opt(&options)
	}

	var qParts []string
	if options.ProcessType != "" {
		qParts = append(qParts, "type:eq("+string(options.ProcessType)+")")
	}
	if options.ProcessStatus != "" {
		qParts = append(qParts, "status:eq("+string(options.ProcessStatus)+")")
	}
	if options.Step != "" {
		qParts = append(qParts, "step:eq("+string(options.Step)+")")
	}
	if options.OwnerID != "" {
		qParts = append(qParts, "owner_id:eq("+options.OwnerID+")")
	}

	q := ""
	if len(qParts) > 0 {
		q = strings.Join(qParts, ",")
	}

	url := NewURLBuilder(urlConfig.ChURL).
		Path("v1", "organizations", organizationID, "process").
		QueryParam("q", q).
		Build()

	var response GenericPaginatedResponse[Process]

	if err := request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodGet),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		request_utils.WithApiKey(apiKey),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (api *ProcessAPI) GetProcess(
	urlConfig configuration.URLs,
	apiKey string,
	organizationID string,
	processID string,
) (*Process, error) {
	url := NewURLBuilder(urlConfig.ChURL).
		Path("v1", "organizations", organizationID, "process", processID).
		Build()

	var response Process

	if err := request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodGet),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		request_utils.WithApiKey(apiKey),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, err
	}

	return &response, nil
}
