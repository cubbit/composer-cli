package api

import "github.com/cubbit/composer-cli/src/configuration/configuration_models"

type MockRedundancyClassAPI struct {
	ListRedundancyClassesBySwarmFunc func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, swarmID string) ([]RedundancyClass, error)
}

func (m *MockRedundancyClassAPI) ListRedundancyClassesBySwarm(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, swarmID string) ([]RedundancyClass, error) {
	if m.ListRedundancyClassesBySwarmFunc != nil {
		return m.ListRedundancyClassesBySwarmFunc(endpoints, apiKey, organizationID, swarmID)
	}
	return []RedundancyClass{}, nil
}
