package service

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cubbit/composer-cli/constants"
	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	usercreate "github.com/cubbit/composer-cli/src/service/user/create"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

type UserServiceInterface interface {
	ImportUsers(cmd *cobra.Command, args []string) error
	CreateUser(cmd *cobra.Command, args []string) error
	ListUsers(cmd *cobra.Command, args []string) error
	DescribeUser(cmd *cobra.Command, args []string) error
}

type UserService struct {
	configuration configuration_handler.ConfigurationHandlerInterface
	userAPI       api.UserAPIInterface
}

func NewUserService(
	configuration configuration_handler.ConfigurationHandlerInterface,
	userAPI api.UserAPIInterface,
) *UserService {
	return &UserService{
		configuration: configuration,
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

	enabledStr, err := cmd.Flags().GetString("enabled")
	if err != nil {
		return fmt.Errorf("%s enabled: %w", constants.ErrorRetrievingField, err)
	}
	search, err := cmd.Flags().GetString("search")
	if err != nil {
		return fmt.Errorf("%s search: %w", constants.ErrorRetrievingField, err)
	}
	page, err := cmd.Flags().GetInt("page")
	if err != nil {
		return fmt.Errorf("%s page: %w", constants.ErrorRetrievingField, err)
	}
	items, err := cmd.Flags().GetInt("items")
	if err != nil {
		return fmt.Errorf("%s items: %w", constants.ErrorRetrievingField, err)
	}
	sortKey, err := cmd.Flags().GetString("sort-key")
	if err != nil {
		return fmt.Errorf("%s sort-key: %w", constants.ErrorRetrievingField, err)
	}
	sortOrder, err := cmd.Flags().GetString("sort-order")
	if err != nil {
		return fmt.Errorf("%s sort-order: %w", constants.ErrorRetrievingField, err)
	}

	var enabled *bool
	if enabledStr != "" {
		parsed, err := strconv.ParseBool(enabledStr)
		if err != nil {
			return fmt.Errorf("invalid value for --enabled: expected true or false, got %q", enabledStr)
		}
		enabled = &parsed
	}

	if items <= 0 {
		items = 100
	}

	var allUsers []api.IAMUserListItem
	if cmd.Flags().Changed("page") {
		response, err := s.userAPI.ListIAMUsers(
			profile.Endpoints, profile.APIKey, profile.OrganizationID,
			enabled, search, page, items, sortKey, sortOrder,
		)
		if err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingIAMUsersRequest, err)
		}
		allUsers = response.Data
	} else {
		page = 1
		for {
			response, err := s.userAPI.ListIAMUsers(
				profile.Endpoints, profile.APIKey, profile.OrganizationID,
				enabled, search, page, items, sortKey, sortOrder,
			)
			if err != nil {
				return fmt.Errorf("%s: %w", constants.ErrorListingIAMUsersRequest, err)
			}
			allUsers = append(allUsers, response.Data...)
			if response.NextPage == nil {
				break
			}
			page = *response.NextPage
		}
	}

	output, err := resolveCommandOutput(cmd, profile.Output)
	if err != nil {
		return err
	}
	quiet, err := cmd.Flags().GetBool("quiet")
	if err != nil {
		return fmt.Errorf("%s quiet: %w", constants.ErrorRetrievingField, err)
	}
	if output != "human" && !quiet {
		utils.PrintFormattedData(cmd.OutOrStdout(), allUsers, output)
		return nil
	}

	if quiet {
		for _, item := range allUsers {
			firstName := ""
			if item.FirstName != nil {
				firstName = *item.FirstName
			}
			lastName := ""
			if item.LastName != nil {
				lastName = *item.LastName
			}
			utils.PrintQuiet(
				cmd.OutOrStdout(),
				item.Username,
				item.ID,
				defaultIAMUserEmail(item.Emails),
				firstName,
				lastName,
				item.Status,
				strconv.FormatBool(item.Enabled),
			)
		}
		return nil
	}

	return PrintIAMUserList(cmd, allUsers)
}

func (s *UserService) DescribeUser(cmd *cobra.Command, args []string) error {
	profile, err := s.configuration.GetActiveProfile()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	userID, err := s.resolveUserID(cmd, args, profile.Endpoints, profile.APIKey, profile.OrganizationID)
	if err != nil {
		return err
	}

	user, err := s.userAPI.GetIAMUserByID(profile.Endpoints, profile.APIKey, profile.OrganizationID, userID)
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDescribingIAMUserRequest, err)
	}

	user.OrganizationID = &profile.OrganizationID

	organizationName, err := s.getProfileOrganizationName(profile)
	if err == nil {
		user.OrganizationName = &organizationName
	}

	output, err := resolveCommandOutput(cmd, profile.Output)
	if err != nil {
		return err
	}

	if output == string(configuration_models.OutputHuman) {
		return PrintIAMUserDetails(cmd, *user)
	}

	utils.PrintFormattedData(cmd.OutOrStdout(), user, output)
	return nil
}

func (s *UserService) resolveUserID(
	cmd *cobra.Command,
	args []string,
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
) (string, error) {
	identifiers := 0

	userIDFlag, err := cmd.Flags().GetString("user-id")
	if err != nil {
		return "", fmt.Errorf("%s user-id: %w", constants.ErrorRetrievingField, err)
	}
	if userIDFlag != "" {
		identifiers++
	}

	usernameFlag, err := cmd.Flags().GetString("username")
	if err != nil {
		return "", fmt.Errorf("%s username: %w", constants.ErrorRetrievingField, err)
	}
	if usernameFlag != "" {
		identifiers++
	}

	userIDPositional := ""
	if len(args) > 0 {
		userIDPositional = strings.TrimSpace(args[0])
		if userIDPositional != "" {
			identifiers++
		}
	}

	if identifiers != 1 {
		return "", fmt.Errorf("specify exactly one of USER_ID, --user-id or --username")
	}

	if userIDPositional != "" {
		return userIDPositional, nil
	}

	if userIDFlag != "" {
		return userIDFlag, nil
	}

	return s.resolveUserIDByName(endpoints, apiKey, organizationID, usernameFlag)
}

func (s *UserService) resolveUserIDByName(
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
	username string,
) (string, error) {
	page := 1
	const itemsPerPage = 1000

	for {
		response, err := s.userAPI.ListIAMUsers(
			endpoints, apiKey, organizationID,
			nil, username, page, itemsPerPage, "", "",
		)
		if err != nil {
			return "", fmt.Errorf("failed to resolve username '%s': %w", username, err)
		}

		for _, item := range response.Data {
			if item.Username == username {
				return item.ID, nil
			}
		}

		if response.NextPage != nil {
			page = *response.NextPage
			continue
		}

		break
	}

	return "", fmt.Errorf("IAM user with username '%s' not found", username)
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
