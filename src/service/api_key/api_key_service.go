package apikey

import (
	"fmt"

	"github.com/cubbit/composer-cli/constants"
	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	apikeycreate "github.com/cubbit/composer-cli/src/service/api_key/create"
	apikeydescribe "github.com/cubbit/composer-cli/src/service/api_key/describe"
	apikeyedit "github.com/cubbit/composer-cli/src/service/api_key/edit"
	apikeylist "github.com/cubbit/composer-cli/src/service/api_key/list"
	apikeyrevoke "github.com/cubbit/composer-cli/src/service/api_key/revoke"
	"github.com/cubbit/composer-cli/src/service/api_key/shared"
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

	return apikeycreate.CreateAPIKey(shared.Dependencies{UserAPI: s.userAPI}, cmd, s.configuration, profile)
}

func (s *APIKeyService) ListAPIKeys(cmd *cobra.Command, args []string) error {
	profile, err := s.configuration.GetActiveProfile()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	return apikeylist.ListAPIKeys(shared.Dependencies{UserAPI: s.userAPI}, cmd, s.configuration, profile)
}

func (s *APIKeyService) DescribeAPIKey(cmd *cobra.Command, args []string) error {
	profile, err := s.configuration.GetActiveProfile()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	return apikeydescribe.DescribeAPIKey(shared.Dependencies{UserAPI: s.userAPI}, cmd, s.configuration, profile)
}

func (s *APIKeyService) EditAPIKey(cmd *cobra.Command, args []string) error {
	profile, err := s.configuration.GetActiveProfile()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	return apikeyedit.EditAPIKey(shared.Dependencies{UserAPI: s.userAPI}, cmd, s.configuration, profile)
}

func (s *APIKeyService) RevokeAPIKey(cmd *cobra.Command, args []string) error {
	profile, err := s.configuration.GetActiveProfile()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	return apikeyrevoke.RevokeAPIKey(shared.Dependencies{UserAPI: s.userAPI}, cmd, profile)
}

var _ APIKeyServiceInterface = (*APIKeyService)(nil)
