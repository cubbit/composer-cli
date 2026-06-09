package api

import "github.com/cubbit/composer-cli/src/configuration"

// MockProcessAPI implements ProcessAPIInterface for testing.
type MockProcessAPI struct {
	ListProcessesFunc func(urlConfig configuration.URLs, apiKey string, organizationID string, processType ProcessType) ([]Process, error)
	GetProcessFunc    func(urlConfig configuration.URLs, apiKey string, organizationID string, processID string) (*Process, error)
}

func (m *MockProcessAPI) ListProcesses(urlConfig configuration.URLs, apiKey string, organizationID string, processType ProcessType) ([]Process, error) {
	if m.ListProcessesFunc != nil {
		return m.ListProcessesFunc(urlConfig, apiKey, organizationID, processType)
	}
	return []Process{}, nil
}

func (m *MockProcessAPI) GetProcess(urlConfig configuration.URLs, apiKey string, organizationID string, processID string) (*Process, error) {
	if m.GetProcessFunc != nil {
		return m.GetProcessFunc(urlConfig, apiKey, organizationID, processID)
	}
	return &Process{}, nil
}
