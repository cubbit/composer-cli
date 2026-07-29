package apikey

import (
	"fmt"

	"github.com/cubbit/composer-cli/constants"
	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/spf13/cobra"
)

type APIKeyServiceInterface interface {
	CreateAPIKey(cmd *cobra.Command, args []string) error
	ListAPIKeys(cmd *cobra.Command, args []string) error
	DescribeAPIKey(cmd *cobra.Command, args []string) error
	EditAPIKey(cmd *cobra.Command, args []string) error
	RevokeAPIKey(cmd *cobra.Command, args []string) error
}

type APIKeyService struct {
	configuration configuration_handler.ConfigurationHandlerInterface
	userAPI       api.UserAPIInterface
}

func NewAPIKeyService(
	configuration configuration_handler.ConfigurationHandlerInterface,
	userAPI api.UserAPIInterface,
) *APIKeyService {
	return &APIKeyService{
		configuration: configuration,
		userAPI:       userAPI,
	}
}

func (s *APIKeyService) CreateAPIKey(cmd *cobra.Command, args []string) error {
	profile, err := s.configuration.GetActiveProfile()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	return CreateAPIKey(Dependencies{UserAPI: s.userAPI}, cmd, profile)
}

func (s *APIKeyService) ListAPIKeys(cmd *cobra.Command, args []string) error {
	profile, err := s.configuration.GetActiveProfile()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	return ListAPIKeys(Dependencies{UserAPI: s.userAPI}, cmd, profile)
}

func (s *APIKeyService) DescribeAPIKey(cmd *cobra.Command, args []string) error {
	profile, err := s.configuration.GetActiveProfile()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	return DescribeAPIKey(Dependencies{UserAPI: s.userAPI}, cmd, profile)
}

func (s *APIKeyService) EditAPIKey(cmd *cobra.Command, args []string) error {
	profile, err := s.configuration.GetActiveProfile()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	return EditAPIKey(Dependencies{UserAPI: s.userAPI}, cmd, profile)
}

func (s *APIKeyService) RevokeAPIKey(cmd *cobra.Command, args []string) error {
	profile, err := s.configuration.GetActiveProfile()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	return RevokeAPIKey(Dependencies{UserAPI: s.userAPI}, cmd, profile)
}

var _ APIKeyServiceInterface = (*APIKeyService)(nil)

type Dependencies struct {
	UserAPI api.UserAPIInterface
}

func resolveCurrentOperator(deps Dependencies, profile configuration_models.ProfileV2) (*api.IAMUser, string, error) {
	operator, err := deps.UserAPI.GetIAMUserSelf(profile.Endpoints, "", profile.APIKey)
	if err != nil {
		return nil, "", fmt.Errorf("failed to retrieve current IAM user: %w", err)
	}
	if operator.ID == "" {
		return nil, "", fmt.Errorf("current IAM user does not expose an ID")
	}
	if operator.OrganizationName == nil || *operator.OrganizationName == "" {
		return nil, "", fmt.Errorf("current IAM user does not expose an organization name")
	}

	return operator, *operator.OrganizationName, nil
}
