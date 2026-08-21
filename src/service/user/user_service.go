package user

import (
	"fmt"

	"github.com/cubbit/composer-cli/constants"
	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	usercreate "github.com/cubbit/composer-cli/src/service/user/create"
	userdelete "github.com/cubbit/composer-cli/src/service/user/delete"
	userdescribe "github.com/cubbit/composer-cli/src/service/user/describe"
	userlist "github.com/cubbit/composer-cli/src/service/user/list"
	userupdate "github.com/cubbit/composer-cli/src/service/user/update"
	"github.com/spf13/cobra"
)

type UserServiceInterface interface {
	ImportUsers(cmd *cobra.Command, args []string) error
	CreateUser(cmd *cobra.Command, args []string) error
	ListUsers(cmd *cobra.Command, args []string) error
	DescribeUser(cmd *cobra.Command, args []string) error
	EditUser(cmd *cobra.Command, args []string) error
	EnableUser(cmd *cobra.Command, args []string) error
	DisableUser(cmd *cobra.Command, args []string) error
	DeleteUser(cmd *cobra.Command, args []string) error
	ResetUserPassword(cmd *cobra.Command, args []string) error
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

func (s *UserService) ListUsers(cmd *cobra.Command, args []string) error {
	profile, err := s.configuration.GetActiveProfile()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	return userlist.ListUsers(userlist.Dependencies{UserAPI: s.userAPI}, cmd, profile)
}

func (s *UserService) DescribeUser(cmd *cobra.Command, args []string) error {
	profile, err := s.configuration.GetActiveProfile()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	return userdescribe.DescribeUser(userdescribe.Dependencies{UserAPI: s.userAPI}, cmd, profile, args)
}

func (s *UserService) EditUser(cmd *cobra.Command, args []string) error {
	profile, err := s.configuration.GetActiveProfile()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	return userupdate.EditUser(userupdate.Dependencies{UserAPI: s.userAPI}, cmd, profile, args)
}

func (s *UserService) EnableUser(cmd *cobra.Command, args []string) error {
	profile, err := s.configuration.GetActiveProfile()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	return userupdate.EnableUser(userupdate.Dependencies{UserAPI: s.userAPI}, cmd, profile, args)
}

func (s *UserService) DisableUser(cmd *cobra.Command, args []string) error {
	profile, err := s.configuration.GetActiveProfile()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	return userupdate.DisableUser(userupdate.Dependencies{UserAPI: s.userAPI}, cmd, profile, args)
}

func (s *UserService) DeleteUser(cmd *cobra.Command, args []string) error {
	profile, err := s.configuration.GetActiveProfile()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	return userdelete.DeleteUser(
		userdelete.Dependencies{
			UserAPI: s.userAPI,
		},
		cmd,
		profile,
		args,
	)
}

func (s *UserService) ResetUserPassword(cmd *cobra.Command, args []string) error {
	profile, err := s.configuration.GetActiveProfile()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	return userupdate.ResetUserPassword(
		userupdate.Dependencies{
			AuthAPI: s.authAPI,
			UserAPI: s.userAPI,
		},
		cmd,
		profile,
		args,
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
