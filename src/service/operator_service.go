package service

import (
	"fmt"

	"github.com/cubbit/composer-cli/constants"
	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	utils "github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

type OperatorServiceInterface interface {
	Connect(cmd *cobra.Command, args []string) error
}

type OperatorService struct {
	configuration configuration_handler.ConfigurationHandlerInterface
	operatorAPI   api.OperatorAPIInterface
	userAPI       api.UserAPIInterface
}

func NewOperatorService(
	configuration configuration_handler.ConfigurationHandlerInterface,
	operatorAPI api.OperatorAPIInterface,
	userAPI api.UserAPIInterface,
) *OperatorService {
	return &OperatorService{
		configuration: configuration,
		operatorAPI:   operatorAPI,
		userAPI:       userAPI,
	}
}

func (s *OperatorService) Connect(cmd *cobra.Command, args []string) error {
	var err error
	var command string

	profile, err := s.configuration.GetActiveProfile()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if command, err = s.operatorAPI.Connect(profile.Endpoints, profile.APIKey, profile.OrganizationID); err != nil {
		return fmt.Errorf("failed to connect to the operator: %w", err)
	}

	return utils.PrintSmartOutput(
		cmd,
		[]string{command},
		func(s string) []string {
			return []string{s}
		},
		&utils.SmartOutputConfig[string]{
			SingleResource:              true,
			SingleResourceCompactOutput: true,
			DefaultOutput:               profile.Output,
		},
	)
}
