package service

import (
	"fmt"

	"github.com/cubbit/composer-cli/constants"
	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration"
	utils "github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

type OperatorServiceInterface interface {
	Connect(cmd *cobra.Command, args []string) error
}

type OperatorService struct {
	configuration *configuration.Config
	operatorApi   api.OperatorAPIInterface
}

func NewOperatorService(
	configuration *configuration.Config,
	operatorAPI api.OperatorAPIInterface,
) *OperatorService {
	return &OperatorService{
		configuration: configuration,
		operatorApi:   operatorAPI,
	}
}

func (s *OperatorService) Connect(cmd *cobra.Command, args []string) error {
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs
	var err error
	var user *api.Operator
	var command string

	if resolvedProfile, urls, err = s.configuration.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if user, err = api.GetIAMUserSelf(*urls, "", resolvedProfile.APIKey); err != nil {
		return fmt.Errorf("failed to get operator information: %w", err)
	}

	if user.OrganizationID == nil {
		return fmt.Errorf("failed to get operator information: organization ID is nil")
	}

	if command, err = s.operatorApi.Connect(*urls, resolvedProfile.APIKey, *user.OrganizationID); err != nil {
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
			DefaultOutput:               resolvedProfile.Output,
		},
	)
}
