package service

import "github.com/cubbit/composer-cli/src/api"

// RedundancyClassValidatorMock implements RedundancyClassValidatorInterface for testing.
type RedundancyClassValidatorMock struct {
	ValidateRedundancyClassesFunc func(redundancyClasses []api.RedundancyClassRequest, nexuses []api.NexusV5Request) error
}

func (m *RedundancyClassValidatorMock) ValidateRedundancyClasses(redundancyClasses []api.RedundancyClassRequest, nexuses []api.NexusV5Request) error {
	if m.ValidateRedundancyClassesFunc != nil {
		return m.ValidateRedundancyClassesFunc(redundancyClasses, nexuses)
	}
	return nil
}
