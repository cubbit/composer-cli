package service

import (
	"fmt"

	"github.com/cubbit/composer-cli/constants"
	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	usercreate "github.com/cubbit/composer-cli/src/service/user/create"
	"github.com/spf13/cobra"
)

type UserServiceInterface interface {
	ImportUsers(cmd *cobra.Command, args []string) error
	CreateUser(cmd *cobra.Command, args []string) error
}

type UserService struct {
	configuration configuration_handler.ConfigurationHandlerInterface
	authAPI       api.AuthAPIInterface
	userAPI       api.UserAPIInterface
}

func NewUserService(
	configuration configuration_handler.ConfigurationHandlerInterface,
	authAPI api.AuthAPIInterface,
	userAPI api.UserAPIInterface,
) *UserService {
	return &UserService{
		configuration: configuration,
		authAPI:       authAPI,
		userAPI:       userAPI,
	}
}

func (s *UserService) ImportUsers(cmd *cobra.Command, args []string) error {
	sampleFormat, err := cmd.Flags().GetString("sample")
	if err != nil {
		return fmt.Errorf("%s sample: %w", constants.ErrorRetrievingField, err)
	}
	if sampleFormat != "" {
		return usercreate.ImportUsers(usercreate.Dependencies{}, cmd, configuration_models.ProfileV2{}, "")
	}

	profile, err := s.configuration.GetActiveProfile()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	organizationName, err := s.getProfileOrganizationName(profile)
	if err != nil {
		return err
	}

	return usercreate.ImportUsers(
		usercreate.Dependencies{
			AuthAPI: s.authAPI,
			UserAPI: s.userAPI,
		},
		cmd,
		profile,
		organizationName,
	)
}

func (s *UserService) CreateUser(cmd *cobra.Command, args []string) error {
	profile, err := s.configuration.GetActiveProfile()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	organizationName, err := s.getProfileOrganizationName(profile)
	if err != nil {
		return err
	}

	return usercreate.CreateUser(
		usercreate.Dependencies{
			AuthAPI: s.authAPI,
			UserAPI: s.userAPI,
		},
		cmd,
		profile,
		organizationName,
	)
}

func (s *UserService) getProfileOrganizationName(profile configuration_models.ProfileV2) (string, error) {
	user, err := s.userAPI.GetIAMUserSelf(profile.Endpoints, "", profile.APIKey)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve current IAM user: %w", err)
	}
	if user.OrganizationName == nil || *user.OrganizationName == "" {
		return "", fmt.Errorf("current IAM user does not expose an organization name")
	}

	return *user.OrganizationName, nil
}
