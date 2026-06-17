package api

import (
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
)

type MockProcessAPI struct {
	Processes         []Process
	ListProcessesFunc func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, opts ...ListProcessesOption) ([]Process, error)
	GetProcessFunc    func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, processID string) (*Process, error)
}

func (m *MockProcessAPI) ListProcesses(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, opts ...ListProcessesOption) ([]Process, error) {
	if m.ListProcessesFunc != nil {
		return m.ListProcessesFunc(endpoints, apiKey, organizationID, opts...)
	}

	var options ListProcessesOptions
	for _, opt := range opts {
		opt(&options)
	}

	var result []Process
	for _, p := range m.Processes {
		if options.ProcessType != "" && p.Type != options.ProcessType {
			continue
		}
		if options.ProcessStatus != "" && p.Status != options.ProcessStatus {
			continue
		}
		if options.Step != "" && p.Step != options.Step {
			continue
		}
		if options.OwnerID != "" && p.OwnerID != options.OwnerID {
			continue
		}
		result = append(result, p)
	}
	return result, nil
}

func (m *MockProcessAPI) GetProcess(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, processID string) (*Process, error) {
	if m.GetProcessFunc != nil {
		return m.GetProcessFunc(endpoints, apiKey, organizationID, processID)
	}
	return &Process{ID: processID, Status: ProcessStatusSuccess}, nil
}
