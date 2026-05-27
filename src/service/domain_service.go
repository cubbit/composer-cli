package service

import (
	"fmt"

	"github.com/cubbit/composer-cli/constants"
	api "github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration"
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
	configuration configuration.ConfigInterface
	domainAPI     api.DomainAPIInterface
	userAPI       api.UserAPIInterface
}

func NewDomainService(
	configuration configuration.ConfigInterface,
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
	resolvedProfile, urls, err := s.configuration.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer)
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

	response, err := s.domainAPI.Create(*urls, resolvedProfile.APIKey, resolvedProfile.OrganizationID, request)
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCreatingDomainRequest, err)
	}

	output, err := resolveCommandOutput(cmd, resolvedProfile.Output)
	if err != nil {
		return err
	}

	if output == string(configuration.OutputHuman) {
		return PrintDomainDetails(cmd, *response)
	}

	utils.PrintFormattedData(cmd.OutOrStdout(), response, output)
	return nil
}

func (s DomainService) Describe(cmd *cobra.Command, args []string) error {
	resolvedProfile, urls, err := s.configuration.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer)
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	domainID, err := cmd.Flags().GetString("domain-id")
	if err != nil {
		return fmt.Errorf("%s domain-id: %w", constants.ErrorRetrievingField, err)
	}

	response, err := s.domainAPI.Get(*urls, resolvedProfile.APIKey, resolvedProfile.OrganizationID, domainID)
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDescribingDomainRequest, err)
	}

	output, err := resolveCommandOutput(cmd, resolvedProfile.Output)
	if err != nil {
		return err
	}

	if output == string(configuration.OutputHuman) {
		return PrintDomainDetails(cmd, *response)
	}

	utils.PrintFormattedData(cmd.OutOrStdout(), response, output)
	return nil
}

func (s DomainService) List(cmd *cobra.Command, args []string) error {
	resolvedProfile, urls, err := s.configuration.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer)
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	domains, err := s.fetchAllDomains(*urls, resolvedProfile.APIKey, resolvedProfile.OrganizationID)
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingDomainsRequest, err)
	}

	output, err := resolveCommandOutput(cmd, resolvedProfile.Output)
	if err != nil {
		return err
	}

	if output == string(configuration.OutputHuman) {
		return PrintDomainList(cmd, domains)
	}

	utils.PrintFormattedData(cmd.OutOrStdout(), domains, output)
	return nil
}

func (s DomainService) fetchAllDomains(urls configuration.URLs, apiKey string, organizationID string) ([]api.DomainDTO, error) {
	page := 1
	itemsPerPage := 100
	var all []api.DomainDTO

	for {
		response, err := s.domainAPI.List(urls, apiKey, organizationID, page, itemsPerPage)
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
	resolvedProfile, urls, err := s.configuration.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer)
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	domainID, err := cmd.Flags().GetString("domain-id")
	if err != nil {
		return fmt.Errorf("%s domain-id: %w", constants.ErrorRetrievingField, err)
	}

	err = s.domainAPI.Delete(*urls, resolvedProfile.APIKey, resolvedProfile.OrganizationID, domainID)
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDeletingDomainRequest, err)
	}

	return printer.PrintText(cmd, fmt.Sprintf("Domain %s deleted successfully\n", domainID))
}

func (s DomainService) Verify(cmd *cobra.Command, args []string) error {
	resolvedProfile, urls, err := s.configuration.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer)
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	domainID, err := cmd.Flags().GetString("domain-id")
	if err != nil {
		return fmt.Errorf("%s domain-id: %w", constants.ErrorRetrievingField, err)
	}

	result, err := s.domainAPI.Verify(*urls, resolvedProfile.APIKey, resolvedProfile.OrganizationID, domainID)
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorVerifyingDomainRequest, err)
	}

	output, err := resolveCommandOutput(cmd, resolvedProfile.Output)
	if err != nil {
		return err
	}

	if output == string(configuration.OutputHuman) {
		verified := "is not verified"
		if result.Verified {
			verified = "is verified"
		}
		return printer.PrintText(cmd, fmt.Sprintf("Domain %s %s\n", domainID, verified))
	}

	utils.PrintFormattedData(cmd.OutOrStdout(), result, output)
	return nil
}
