package api

import "github.com/cubbit/composer-cli/src/configuration"

type MockRedundancyClassAPI struct {
	ListRedundancyClassesBySwarmFunc func(urlConfig configuration.URLs, apiKey string, swarmID string) ([]RedundancyClass, error)
}

func (m *MockRedundancyClassAPI) ListRedundancyClassesBySwarm(urlConfig configuration.URLs, apiKey string, swarmID string) ([]RedundancyClass, error) {
	if m.ListRedundancyClassesBySwarmFunc != nil {
		return m.ListRedundancyClassesBySwarmFunc(urlConfig, apiKey, swarmID)
	}
	return []RedundancyClass{}, nil
}
