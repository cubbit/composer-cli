package action

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/api"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/src/tui"
	"github.com/cubbit/cubbit/client/cli/utils"
	"github.com/spf13/cobra"
)

func CreateProjectInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var name, description, imageUrl, configPath string
	var response *api.GenericIDResponseModel
	var conf *configuration.Config

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeAccount, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if _, err = tui.TextInputs(
		"Fill in the form below",
		false,
		tui.Input{Placeholder: "Name*", IsPassword: false, Value: &name},
		tui.Input{Placeholder: "Description", IsPassword: false, Value: &description},
		tui.Input{Placeholder: "Image URL", IsPassword: false, Value: &imageUrl},
	); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if len(description) > 200 {
		return fmt.Errorf("t%s: %w", constants.ErrorDescriptionSize, err)
	}

	if imageUrl != "" {
		if _, err := url.ParseRequestURI(imageUrl); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorImageURL, err)
		}
	}

	if response, err = api.CreateProject(conf.Urls, *accessToken, name, &description, &imageUrl); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCreatingProjectRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("project  %s created successfully", response.ID))

	return nil
}

func ListTenantProjectsInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, name, configPath string
	var conf *configuration.Config
	var projects *api.GenericPaginatedResponse[*api.ProjectItem]

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if id == "" && name == "" {
		var choice string
		var choices []string
		var tenants *api.GenericPaginatedResponse[*api.Tenant]

		if tenants, err = api.ListTenants(conf.Urls, *accessToken, "", ""); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingTenantsRequest, err)
		}

		for _, tenant := range tenants.Data {
			choices = append(choices, fmt.Sprintf("• %s, %s, %s", tenant.ID, tenant.Name, utils.StringOrEmpty(tenant.Description)))
		}

		if len(choices) == 0 {
			utils.PrintNotFound("No tenants found")
			return nil
		}

		if choice, err = tui.ChooseOne("Choose your tenant", false, false, choices); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}

		_, withoutPrefix, _ := strings.Cut(choice, " ")
		id, _, _ = strings.Cut(withoutPrefix, ",")
	}

	if id == "" {
		var tenant *api.Tenant

		if tenant, err = getTenantByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}

		id = tenant.ID
	}

	var sort string
	if sort, err = tui.ChooseOne("Choose your sort key", false, true, []string{"project_id", "project_name", "project_created_at", "project_deleted_at", "project_banned_at", "project_tenant_id", "project_email", "root_account_email"}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if projects, err = api.ListTenantProjects(conf.Urls, *accessToken, id, sort); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingTenantProjectsRequest, err)
	}

	utils.PrintList("Your Tenant Projects List")

	if len(projects.Data) == 0 {
		utils.PrintEmptyList()
		return nil
	}

	var list []string
	for _, project := range projects.Data {
		list = append(list, fmt.Sprintf(" • %s, %s, %s", project.ProjectID, project.ProjectName, project.ProjectDescription))
	}

	tui.List(list)

	return nil
}

func DescribeTenantProjectInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, name, configPath, choice, projectID, format string
	var conf *configuration.Config
	var projects *api.GenericPaginatedResponse[*api.ProjectItem]

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if id == "" && name == "" {
		var choice string
		var choices []string
		var tenants *api.GenericPaginatedResponse[*api.Tenant]

		if tenants, err = api.ListTenants(conf.Urls, *accessToken, "", ""); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingTenantsRequest, err)
		}

		for _, tenant := range tenants.Data {
			choices = append(choices, fmt.Sprintf("• %s, %s, %s", tenant.ID, tenant.Name, utils.StringOrEmpty(tenant.Description)))
		}

		if len(choices) == 0 {
			utils.PrintNotFound("No tenants found")
			return nil
		}

		if choice, err = tui.ChooseOne("Choose your tenant", false, false, choices); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}

		_, withoutPrefix, _ := strings.Cut(choice, " ")
		id, _, _ = strings.Cut(withoutPrefix, ",")
	}

	if id == "" {
		var tenant *api.Tenant

		if tenant, err = getTenantByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}

		id = tenant.ID
	}

	if projects, err = api.ListTenantProjects(conf.Urls, *accessToken, id, ""); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingTenantProjectsRequest, err)
	}

	if len(projects.Data) == 0 {
		utils.PrintNotFound("No projects found")
		return nil
	}

	var choices []string

	for _, project := range projects.Data {
		choices = append(choices, fmt.Sprintf("• %s, %s, %s", project.ProjectID, project.ProjectName, project.ProjectDescription))
	}

	if len(choices) == 0 {
		utils.PrintNotFound("No projects found")
		return nil
	}

	if choice, err = tui.ChooseOne("Which project would you like to describe?", false, false, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDescribingTenantProjectRequest, err)
	}

	_, withoutPrefix, _ := strings.Cut(choice, " ")
	projectID, _, _ = strings.Cut(withoutPrefix, ",")

	if format, err = tui.ChooseOne("Choose your output format", false, true, []string{"json", "semantic", "csv"}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	for _, p := range projects.Data {
		if p.ProjectID == projectID {
			utils.PrintFormattedData(*p, format)
			break
		}
	}

	return nil
}

func RemoveTenantProjectInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, name, configPath, choice, projectID, email, password, deleteTenantProjectToken, code string
	var conf *configuration.Config
	var projects *api.GenericPaginatedResponse[*api.ProjectItem]
	var challenge *api.ChallengeResponseModel

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if id == "" && name == "" {
		var choice string
		var choices []string
		var tenants *api.GenericPaginatedResponse[*api.Tenant]

		if tenants, err = api.ListTenants(conf.Urls, *accessToken, "", ""); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingTenantsRequest, err)
		}

		for _, tenant := range tenants.Data {
			choices = append(choices, fmt.Sprintf("• %s, %s, %s", tenant.ID, tenant.Name, utils.StringOrEmpty(tenant.Description)))
		}

		if len(choices) == 0 {
			utils.PrintNotFound("No tenants found")
			return nil
		}

		if choice, err = tui.ChooseOne("Choose your tenant", false, false, choices); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}

		_, withoutPrefix, _ := strings.Cut(choice, " ")
		id, _, _ = strings.Cut(withoutPrefix, ",")
	}

	if id == "" {
		var tenant *api.Tenant

		if tenant, err = getTenantByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}

		id = tenant.ID
	}

	if projects, err = api.ListTenantProjects(conf.Urls, *accessToken, id, ""); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingOperatorsRequest, err)
	}

	if len(projects.Data) == 0 {
		utils.PrintNotFound("No projects found")
		return nil
	}

	var choices []string

	for _, project := range projects.Data {
		if project.ProjectDeletedAt == nil {
			choices = append(choices, fmt.Sprintf("• %s, %s, %s", project.ProjectID, project.ProjectName, project.ProjectDescription))
		}
	}

	if len(choices) == 0 {
		utils.PrintNotFound("No projects found")
		return nil
	}

	if choice, err = tui.ChooseOne("Which project would you like to delete?", false, false, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	_, withoutPrefix, _ := strings.Cut(choice, " ")
	projectID, _, _ = strings.Cut(withoutPrefix, ",")

	if _, err = tui.TextInputs(fmt.Sprintf("Confirm your login to delete project %s 🚮", utils.RedBg.Render(projectID)), true, tui.Input{Placeholder: "Email*", IsPassword: false, Value: &email}, tui.Input{Placeholder: "Password*", IsPassword: true, Value: &password}, tui.Input{Placeholder: "Code", IsPassword: false, Value: &code}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if challenge, err = api.GenerateOperatorChallenge(conf.Urls, email); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingOperatorChallengeRequest, err)
	}

	if deleteTenantProjectToken, err = api.ForgeOperatorDeleteTenantProjectToken(conf.Urls, email, password, conf.RefreshToken, challenge, code, projectID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorForgingOperatorDeleteTokenRequest, err)
	}

	if err = api.RemoveTenantProject(conf.Urls, *accessToken, id, projectID, deleteTenantProjectToken); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDeletingTenantProjectRequest, err)
	}

	utils.PrintDelete(fmt.Sprintf("project %s removed successfully", projectID))

	return nil
}

func BanTenantProjectInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, name, configPath, choice, projectID string
	var conf *configuration.Config
	var projects *api.GenericPaginatedResponse[*api.ProjectItem]

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if id == "" && name == "" {
		var choice string
		var choices []string
		var tenants *api.GenericPaginatedResponse[*api.Tenant]

		if tenants, err = api.ListTenants(conf.Urls, *accessToken, "", ""); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingTenantsRequest, err)
		}

		for _, tenant := range tenants.Data {
			choices = append(choices, fmt.Sprintf("• %s, %s, %s", tenant.ID, tenant.Name, utils.StringOrEmpty(tenant.Description)))
		}

		if len(choices) == 0 {
			utils.PrintNotFound("No tenants found")
			return nil
		}

		if choice, err = tui.ChooseOne("Choose your tenant", false, false, choices); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}

		_, withoutPrefix, _ := strings.Cut(choice, " ")
		id, _, _ = strings.Cut(withoutPrefix, ",")
	}

	if id == "" {
		var tenant *api.Tenant

		if tenant, err = getTenantByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}

		id = tenant.ID
	}

	if projects, err = api.ListTenantProjects(conf.Urls, *accessToken, id, ""); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingTenantProjectsRequest, err)
	}

	if len(projects.Data) == 0 {
		utils.PrintNotFound("No projects found")
		return nil
	}

	var choices []string

	for _, project := range projects.Data {
		if project.ProjectBannedAt == nil && project.ProjectDeletedAt == nil {
			choices = append(choices, fmt.Sprintf("• %s, %s , %s", project.ProjectID, project.ProjectName, project.ProjectDescription))
		}
	}

	if len(choices) == 0 {
		utils.PrintNotFound("No projects found")
		return nil
	}

	if choice, err = tui.ChooseOne("Which project would you like to freeze?", false, true, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
	}

	_, withoutPrefix, _ := strings.Cut(choice, " ")
	projectID, _, _ = strings.Cut(withoutPrefix, ",")

	if err = api.ToggleBanProject(conf.Urls, *accessToken, id, projectID, true); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorFreezingTenantProjectRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("project %s freezed successfully", projectID))

	return nil
}

func UnbanTenantProjectInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, name, configPath, choice, projectID string
	var conf *configuration.Config
	var projects *api.GenericPaginatedResponse[*api.ProjectItem]

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if id == "" && name == "" {
		var choice string
		var choices []string
		var tenants *api.GenericPaginatedResponse[*api.Tenant]

		if tenants, err = api.ListTenants(conf.Urls, *accessToken, "", ""); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingTenantsRequest, err)
		}

		for _, tenant := range tenants.Data {
			choices = append(choices, fmt.Sprintf("• %s, %s, %s", tenant.ID, tenant.Name, utils.StringOrEmpty(tenant.Description)))
		}

		if len(choices) == 0 {
			utils.PrintNotFound("No tenants found")
			return nil
		}

		if choice, err = tui.ChooseOne("Choose your tenant", false, false, choices); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}

		_, withoutPrefix, _ := strings.Cut(choice, " ")
		id, _, _ = strings.Cut(withoutPrefix, ",")
	}

	if id == "" {
		var tenant *api.Tenant

		if tenant, err = getTenantByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}

		id = tenant.ID
	}

	if projects, err = api.ListTenantProjects(conf.Urls, *accessToken, id, ""); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingTenantProjectsRequest, err)
	}

	if len(projects.Data) == 0 {
		utils.PrintNotFound("No projects found")
		return nil
	}

	var choices []string
	for _, project := range projects.Data {
		if project.ProjectBannedAt != nil && project.ProjectDeletedAt == nil {
			choices = append(choices, fmt.Sprintf("• %s, %s , %s", project.ProjectID, project.ProjectName, project.ProjectDescription))
		}
	}

	if len(choices) == 0 {
		utils.PrintNotFound("No projects found")
		return nil
	}

	if choice, err = tui.ChooseOne("Which project would you like to unfreeze?", false, true, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
	}

	_, withoutPrefix, _ := strings.Cut(choice, " ")
	projectID, _, _ = strings.Cut(withoutPrefix, ",")

	if err = api.ToggleBanProject(conf.Urls, *accessToken, id, projectID, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDeletingTenantRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("project %s unfreezed successfully", projectID))

	return nil
}

func RestoreTenantProjectInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, name, configPath, choice, projectID string
	var conf *configuration.Config
	var projects *api.GenericPaginatedResponse[*api.ProjectItem]

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if id == "" && name == "" {
		var choice string
		var choices []string
		var tenants *api.GenericPaginatedResponse[*api.Tenant]

		if tenants, err = api.ListTenants(conf.Urls, *accessToken, "", ""); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingOperatorsRequest, err)
		}

		for _, tenant := range tenants.Data {
			choices = append(choices, fmt.Sprintf("• %s, %s, %s", tenant.ID, tenant.Name, utils.StringOrEmpty(tenant.Description)))
		}

		if len(choices) == 0 {
			utils.PrintNotFound("No tenants found")
			return nil
		}

		if choice, err = tui.ChooseOne("Choose your tenant", false, false, choices); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}

		_, withoutPrefix, _ := strings.Cut(choice, " ")
		id, _, _ = strings.Cut(withoutPrefix, ",")
	}

	if id == "" {
		var tenant *api.Tenant

		if tenant, err = getTenantByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}

		id = tenant.ID
	}

	if projects, err = api.ListTenantProjects(conf.Urls, *accessToken, id, ""); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingTenantProjectsRequest, err)
	}

	if len(projects.Data) == 0 {
		utils.PrintNotFound("No projects found")
		return nil
	}

	var choices []string

	for _, project := range projects.Data {
		if project.ProjectDeletedAt != nil {
			choices = append(choices, fmt.Sprintf("• %s, %s , %s", project.ProjectID, project.ProjectName, project.ProjectDescription))
		}
	}

	if len(choices) == 0 {
		utils.PrintNotFound("No projects found")
		return nil
	}

	if choice, err = tui.ChooseOne("Which project would you like to restore?", false, true, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
	}

	_, withoutPrefix, _ := strings.Cut(choice, " ")
	projectID, _, _ = strings.Cut(withoutPrefix, ",")

	if err = api.RestoreTenantProject(conf.Urls, *accessToken, id, projectID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRestoringTenantProjectRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("project %s restored successfully", projectID))

	return nil
}

func UpdateTenantProjectInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, name, configPath, choice, description, imageUrl string
	var conf *configuration.Config
	var projects *api.GenericPaginatedResponse[*api.ProjectItem]

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if id == "" && name == "" {
		var choice string
		var choices []string
		var tenants *api.GenericPaginatedResponse[*api.Tenant]

		if tenants, err = api.ListTenants(conf.Urls, *accessToken, "", ""); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingTenantsRequest, err)
		}

		for _, tenant := range tenants.Data {
			choices = append(choices, fmt.Sprintf("• %s, %s, %s", tenant.ID, tenant.Name, utils.StringOrEmpty(tenant.Description)))
		}

		if len(choices) == 0 {
			utils.PrintNotFound("No tenants found")
			return nil
		}

		if choice, err = tui.ChooseOne("Choose your tenant", false, false, choices); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}

		_, withoutPrefix, _ := strings.Cut(choice, " ")
		id, _, _ = strings.Cut(withoutPrefix, ",")
	}

	if id == "" {
		var tenant *api.Tenant

		if tenant, err = getTenantByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}

		id = tenant.ID
	}

	if projects, err = api.ListTenantProjects(conf.Urls, *accessToken, id, ""); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingTenantProjectsRequest, err)
	}

	if len(projects.Data) == 0 {
		utils.PrintNotFound("No projects found")
		return nil
	}

	var choices []string

	for _, project := range projects.Data {
		if project.ProjectDeletedAt == nil {
			choices = append(choices, fmt.Sprintf("• %s, %s , %s", project.ProjectID, project.ProjectName, project.ProjectDescription))
		}
	}

	if len(choices) == 0 {
		utils.PrintNotFound("No projects found")
		return nil
	}

	if choice, err = tui.ChooseOne("Which project would you like to update?", false, false, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorUpdatingTenantProjectRequest, err)
	}

	_, withoutPrefix, _ := strings.Cut(choice, " ")
	projectID, _, _ := strings.Cut(withoutPrefix, ",")

	if _, err = tui.TextInputs("Fill in the form for the project to update", true, tui.Input{Placeholder: "Description", Value: &description}, tui.Input{Placeholder: "Image", Value: &imageUrl}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

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
