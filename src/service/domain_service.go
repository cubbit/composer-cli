package service

import (
	"fmt"

	"github.com/cubbit/composer-cli/constants"
	api "github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/utils"
	"github.com/cubbit/composer-cli/utils/printer"
	"github.com/spf13/cobra"
)

type DomainServiceInterface interface {
	Create(cmd *cobra.Command, args []string) error
	Describe(cmd *cobra.Command, args []string) error
	List(cmd *cobra.Command, args []string) error
	Delete(cmd *cobra.Command, args []string) error
	Verify(cmd *cobra.Command, args []string) error
}

type DomainService struct {
	configuration configuration_handler.ConfigurationHandlerInterface
	domainAPI     api.DomainAPIInterface
	userAPI       api.UserAPIInterface
}

func NewDomainService(
	configuration configuration_handler.ConfigurationHandlerInterface,
	domainAPI api.DomainAPIInterface,
	userAPI api.UserAPIInterface,
) DomainService {
	return DomainService{
		configuration: configuration,
		domainAPI:     domainAPI,
		userAPI:       userAPI,
	}
}

func (s DomainService) Create(cmd *cobra.Command, args []string) error {
	profile, err := s.configuration.GetActiveProfile()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}
	domainName, err := cmd.Flags().GetString("domain-name")
	if err != nil {
		return fmt.Errorf("%s domain-name: %w", constants.ErrorRetrievingField, err)
	}

	request := &api.CreateDomainRequestBody{
		DomainName: domainName,
	}

	response, err := s.domainAPI.Create(
		profile.Endpoints,
		profile.APIKey,
		profile.OrganizationID,
		request,
	)
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCreatingDomainRequest, err)
	}

	if profile.Output == configuration_models.OutputHuman {
		return PrintDomainDetails(cmd, *response)
	}

	utils.PrintFormattedData(cmd.OutOrStdout(), response, string(profile.Output))
	return nil
}

func (s DomainService) Describe(cmd *cobra.Command, args []string) error {
	profile, err := s.configuration.GetActiveProfile()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	domainID, err := cmd.Flags().GetString("domain-id")
	if err != nil {
		return fmt.Errorf("%s domain-id: %w", constants.ErrorRetrievingField, err)
	}

	response, err := s.domainAPI.Get(
		profile.Endpoints, profile.APIKey, profile.OrganizationID, domainID)
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDescribingDomainRequest, err)
	}

	if profile.Output == configuration_models.OutputHuman {
		return PrintDomainDetails(cmd, *response)
	}

	utils.PrintFormattedData(cmd.OutOrStdout(), response, string(profile.Output))
	return nil
}

func (s DomainService) List(cmd *cobra.Command, args []string) error {
	profile, err := s.configuration.GetActiveProfile()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	domains, err := s.fetchAllDomains(profile.Endpoints, profile.APIKey, profile.OrganizationID)
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingDomainsRequest, err)
	}

	if profile.Output == configuration_models.OutputHuman {
		return PrintDomainList(cmd, domains)
	}

	utils.PrintFormattedData(cmd.OutOrStdout(), domains, string(profile.Output))
	return nil
}

func (s DomainService) fetchAllDomains(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string) ([]api.DomainDTO, error) {
	page := 1
	itemsPerPage := 100
	var all []api.DomainDTO

	for {
		response, err := s.domainAPI.List(endpoints, apiKey, organizationID, page, itemsPerPage)
		if err != nil {
			return nil, err
		}

		all = append(all, response.Data...)

		if response.NextPage == nil {
			break
		}

		page = *response.NextPage
	}

	return all, nil
}

func (s DomainService) Delete(cmd *cobra.Command, args []string) error {
	profile, err := s.configuration.GetActiveProfile()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	domainID, err := cmd.Flags().GetString("domain-id")
	if err != nil {
		return fmt.Errorf("%s domain-id: %w", constants.ErrorRetrievingField, err)
	}

	err = s.domainAPI.Delete(profile.Endpoints, profile.APIKey, profile.OrganizationID, domainID)
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDeletingDomainRequest, err)
	}

	return printer.PrintText(cmd, fmt.Sprintf("Domain %s deleted successfully\n", domainID))
}

func (s DomainService) Verify(cmd *cobra.Command, args []string) error {
	profile, err := s.configuration.GetActiveProfile()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	domainID, err := cmd.Flags().GetString("domain-id")
	if err != nil {
		return fmt.Errorf("%s domain-id: %w", constants.ErrorRetrievingField, err)
	}

	result, err := s.domainAPI.Verify(profile.Endpoints, profile.APIKey, profile.OrganizationID, domainID)
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorVerifyingDomainRequest, err)
	}

	if profile.Output == configuration_models.OutputHuman {
		verified := "is not verified"
		if result.Verified {
			verified = "is verified"
		}
		return printer.PrintText(cmd, fmt.Sprintf("Domain %s %s\n", domainID, verified))
	}

	utils.PrintFormattedData(cmd.OutOrStdout(), result, string(profile.Output))
	return nil
}
