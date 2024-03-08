package action

import (
	"fmt"
	"net/url"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/api"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/utils"
	"github.com/spf13/cobra"
)

func CreateProject(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var name, description, imageUrl, configPath string
	var response *api.GenericIDResponseModel
	var conf *configuration.Config

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeAccount); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if description, err = cmd.Flags().GetString("description"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if len(description) > 200 {
		return fmt.Errorf("t%s: %w", constants.ErrorDescriptionSize, err)
	}

	if imageUrl, err = cmd.Flags().GetString("image-url"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if imageUrl != "" {
		if _, err := url.ParseRequestURI(imageUrl); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorImageURL, err)
		}
	}

	if response, err = api.CreateProject(conf.Urls, *accessToken, name, &description, &imageUrl); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCreatingProjectRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("project: %s created successfully", response.ID))

	return nil
}

func ListTenantProjects(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, configPath, sort, filter string
	var conf *configuration.Config
	var projects *api.GenericPaginatedResponse[*api.ProjectItem]

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if sort, err = cmd.Flags().GetString("sort"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id == "" {
		var tenant *api.Tenant

		if tenant, err = getTenantByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}
		id = tenant.ID
	}

	if filter, err = cmd.Flags().GetString("filter"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if filter != "" {
		filter = utils.BuildFilterQuery(filter)
	}

	if projects, err = api.ListTenantProjects(conf.Urls, *accessToken, id, sort, filter); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingTenantProjectsRequest, err)
	}

	utils.PrintList("Your Tenant Projects List")

	if len(projects.Data) == 0 {
		utils.PrintEmptyList()
		return nil
	}

	var verbose, l bool

	if verbose, err = cmd.Flags().GetBool("verbose"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if l, err = cmd.Flags().GetBool("line"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	for _, project := range projects.Data {
		if verbose {
			fmt.Printf(" • %s, %s, %s\n", project.ProjectID, project.ProjectName, project.ProjectDescription)
		} else {
			fmt.Printf(" • %s\n", project.ProjectID)
		}
		if l {
			fmt.Println()
		}
	}

	return nil
}

func DescribeTenantProject(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, configPath, format string
	var conf *configuration.Config

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id == "" {
		var tenant *api.Tenant

		if tenant, err = getTenantByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}

		id = tenant.ID
	}

	var project *api.ProjectItem
	projectID := args[0]
	if project, err = api.GetTenantProject(conf.Urls, *accessToken, id, projectID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenantProjectRequest, err)
	}

	if format, err = cmd.Flags().GetString("format"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	utils.PrintFormattedData(*project, format)

	return nil
}

func RemoveTenantProject(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, email, password, code, configPath, deleteTenantProjectToken string
	var conf *configuration.Config
	var challenge *api.ChallengeResponseModel

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if email, err = cmd.Flags().GetString("email"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if password, err = cmd.Flags().GetString("password"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if code, err = cmd.Flags().GetString("code"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id == "" {
		var tenants *api.GenericPaginatedResponse[*api.Tenant]

		if tenants, err = api.ListTenants(conf.Urls, *accessToken, "", ""); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingTenantsRequest, err)
		}

		for _, tenant := range tenants.Data {
			if name == tenant.Name {
				id = tenant.ID
			}
		}

		if id == "" {
			utils.PrintNotFound(fmt.Sprintf("Tenant %s not found", name))
			return nil
		}
	}

	if challenge, err = api.GenerateOperatorChallenge(conf.Urls, email); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingOperatorChallengeRequest, err)
	}

	projectID := args[0]
	if deleteTenantProjectToken, err = api.ForgeOperatorDeleteTenantProjectToken(conf.Urls, email, password, conf.RefreshToken, challenge, code, projectID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorForgingTenantProjectDeleteTokenRequest, err)
	}

	if err = api.RemoveTenantProject(conf.Urls, *accessToken, id, projectID, deleteTenantProjectToken); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDeletingTenantProjectRequest, err)
	}

	utils.PrintDelete(fmt.Sprintf("project %s removed successfully", projectID))

	return nil
}

func BanTenantProject(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, configPath string
	var conf *configuration.Config

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id == "" {
		var tenants *api.GenericPaginatedResponse[*api.Tenant]

		if tenants, err = api.ListTenants(conf.Urls, *accessToken, "", ""); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingTenantsRequest, err)
		}

		for _, tenant := range tenants.Data {
			if name == tenant.Name {
				id = tenant.ID
			}
		}

		if id == "" {
			utils.PrintNotFound(fmt.Sprintf("Tenant %s not found", name))
			return nil
		}
	}

	projectID := args[0]

	if err = api.ToggleBanProject(conf.Urls, *accessToken, id, projectID, true); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorFreezingTenantProjectRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("project %s freezed successfully", projectID))

	return nil
}

func UnbanTenantProject(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, configPath string
	var conf *configuration.Config

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id == "" {
		var tenants *api.GenericPaginatedResponse[*api.Tenant]

		if tenants, err = api.ListTenants(conf.Urls, *accessToken, "", ""); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingTenantsRequest, err)
		}

		for _, tenant := range tenants.Data {
			if name == tenant.Name {
				id = tenant.ID
			}
		}

		if id == "" {
			utils.PrintNotFound(fmt.Sprintf("Tenant %s not found", name))
			return nil
		}
	}

	projectID := args[0]

	if err = api.ToggleBanProject(conf.Urls, *accessToken, id, projectID, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorUnfreezingTenantProjectRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("project %s unfreezed successfully", projectID))

	return nil
}

func RestoreTenantProject(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, configPath string
	var conf *configuration.Config

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id == "" {
		var tenants *api.GenericPaginatedResponse[*api.Tenant]

		if tenants, err = api.ListTenants(conf.Urls, *accessToken, "", ""); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingTenantsRequest, err)
		}

		for _, tenant := range tenants.Data {
			if name == tenant.Name {
				id = tenant.ID
			}
		}

		if id == "" {
			utils.PrintNotFound(fmt.Sprintf("Tenant %s not found", name))
			return nil
		}
	}

	projectID := args[0]

	if err = api.RestoreTenantProject(conf.Urls, *accessToken, id, projectID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRestoringTenantProjectRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("project %s restored successfully", projectID))

	return nil
}

func UpdateTenantProject(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, configPath, description, imageUrl string
	var conf *configuration.Config

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if description, err = cmd.Flags().GetString("description"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if imageUrl, err = cmd.Flags().GetString("image-url"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id == "" {
		var tenant *api.Tenant

		if tenant, err = getTenantByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}

		id = tenant.ID
	}

	projectID := args[0]

	var descriptionPtr *string
	if description != "" {
		descriptionPtr = &description
	}

	var imageUrlPtr *string
	if imageUrl != "" {
		imageUrlPtr = &imageUrl
	}

	requestBody := api.UpdateTenantProjectRequestBody{
		Description: descriptionPtr,
		ImageUrl:    imageUrlPtr,
	}

	if err = api.UpdateProject(conf.Urls, *accessToken, id, projectID, requestBody); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorUpdatingTenantProjectRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("project %s updated successfully", projectID))

	return nil
}
