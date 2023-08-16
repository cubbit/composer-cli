package action

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/url"
	"os"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/api"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/spf13/cobra"
)

func CreateTenant(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var name, description, imageUrl, settingsString, configPath string
	var response *api.GenericIDResponseModel
	var conf *configuration.Config

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}
	if description, err = cmd.Flags().GetString("description"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if len(description) > 200 {
		return fmt.Errorf("t%s: %w", constants.ErrorTenantDescriptionSize, err)
	}

	if imageUrl, err = cmd.Flags().GetString("image-url"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if imageUrl != "" {
		if _, err := url.ParseRequestURI(imageUrl); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorImageURL, err)
		}
	}

	if settingsString == "" {
		settingsString = "{}"
	}

	if conf, configPath, err = configuration.ReadConfig(cmd); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}
	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	var settings map[string]interface{}

	if err = json.Unmarshal([]byte(settingsString), &settings); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorParsingJsonSettings, err)
	}

	if response, err = api.CreateTenant(conf.Urls, *accessToken, name, &description, &imageUrl, settings); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCreatingTenant, err)
	}

	fmt.Printf("Successfully created tenant: %s\n", response.ID)
	return nil
}

func ListTenant(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var configPath string
	var conf *configuration.Config
	var operator *api.Operator
	var tenants *api.TenantList

	fmt.Println("these are your tenants")

	if conf, configPath, err = configuration.ReadConfig(cmd); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}
	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if operator, err = api.GetOperatorSelf(conf.Urls, *accessToken); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingOperator, err)
	}
	if tenants, err = api.ListTenants(conf.Urls, *accessToken, operator.ID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenantList, err)
	}

	var verbose, l bool
	if verbose, err = cmd.Flags().GetBool("verbose"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}
	if l, err = cmd.Flags().GetBool("line"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	for _, tenant := range tenants.Tenants {
		if verbose {
			fmt.Printf("%s %s %s ", tenant.ID, tenant.Name, *tenant.Description)
		} else {
			fmt.Printf("%s ", tenant.Name)
		}
		if l {
			fmt.Println()
		}
	}

	return nil
}

func RemoveTenant(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, name, email, password, code, configPath, deleteTenantToken string
	var conf *configuration.Config
	var challenge *api.ChallengeResponseModel

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}
	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if id == "" && name == "" {
		return fmt.Errorf("invalid tenant id or name: %w", err)
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
	if conf, configPath, err = configuration.ReadConfig(cmd); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id == "" {
		var operator *api.Operator
		var tenants *api.TenantList

		if operator, err = api.GetOperatorSelf(conf.Urls, *accessToken); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingOperator, err)
		}
		if tenants, err = api.ListTenants(conf.Urls, *accessToken, operator.ID); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenantList, err)
		}
		for _, tenant := range tenants.Tenants {
			if name == tenant.Name {
				id = tenant.ID
			}
		}
		if id == "" {
			fmt.Printf("Tenant %s not found\n", name)
			return nil
		}
	}

	if challenge, err = api.GenerateOperatorChallenge(conf.Urls, email); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingOperatorChallenge, err)
	}

	if deleteTenantToken, err = api.ForgeOperatorDeleteTenantToken(conf.Urls, email, password, conf.RefreshToken, challenge, code, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorForgingOperatorDeleteToken, err)
	}

	if err = api.RemoveTenant(conf.Urls, *accessToken, id, deleteTenantToken); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDeletingTenant, err)
	}

	fmt.Printf("tenant %s removed successfully\n", id)
	return nil
}

func DescribeTenant(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, name, format, configPath string
	var conf *configuration.Config
	var tenants *api.TenantList
	var operator *api.Operator

	if conf, configPath, err = configuration.ReadConfig(cmd); err != nil {
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
	if format, err = cmd.Flags().GetString("format"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}
	if operator, err = api.GetOperatorSelf(conf.Urls, *accessToken); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingOperator, err)
	}
	if tenants, err = api.ListTenants(conf.Urls, *accessToken, operator.ID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenantList, err)
	}

	switch {
	case id == "" && name == "":
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenantDescription, err)
	case name == "":
		for _, tenant := range tenants.Tenants {
			if id == tenant.ID {
				FormatTenant(format, tenant)
			}
		}
	case id == "":
		for _, tenant := range tenants.Tenants {
			if name == tenant.Name {
				FormatTenant(format, tenant)
			}
		}
	default:
		return fmt.Errorf("%s: %w", constants.ErrorTenantNameOrID, err)
	}

	return nil
}

func FormatTenant(format string, tenant *api.Tenant) error {
	switch {
	case format == "default":
		fmt.Printf("ID: %s\n", tenant.ID)
		fmt.Printf("Name: %s\n", tenant.Name)

		if tenant.Description != nil {
			fmt.Printf("Description: %s\n", *tenant.Description)
		}

		fmt.Printf("OwnerID: %s\n", tenant.OwnerID)
		fmt.Printf("CreatedAt: %s\n", tenant.CreatedAt)

		if tenant.DeletedAt != nil {
			fmt.Printf("DeletedAt: %s\n", tenant.DeletedAt)
		}

		if tenant.ImageUrl != nil && *tenant.ImageUrl != "" {
			fmt.Printf("ImageUrl: %s\n", *tenant.ImageUrl)
		}

		fmt.Printf("Settings:\n")
		for key, value := range tenant.Settings {
			fmt.Printf(" - %s: %s\n", key, value)
		}

	case format == "json":
		formatJson, err := json.Marshal(api.Tenant{ID: tenant.ID, Name: tenant.Name, Description: tenant.Description, OwnerID: tenant.OwnerID, CreatedAt: tenant.CreatedAt, DeletedAt: tenant.DeletedAt, ImageUrl: tenant.ImageUrl, Settings: tenant.Settings})

		if err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorOpeningJson, err)
		}

		fmt.Println(string(formatJson))

	case format == "csv":
		records := [][]string{
			{"ID", "Name", "Description", "OwnerID", "CreatedAt", "DeletedAt", "ImageUrl"},
		}
		for key := range tenant.Settings {
			fmt.Printf(",%s", key)
		}
		var values []string
		values = append(values, tenant.ID, tenant.Name, *tenant.Description, tenant.OwnerID, tenant.CreatedAt.String(), *tenant.ImageUrl)

		if tenant.DeletedAt != nil {
			values = append(values, tenant.DeletedAt.String())
		} else {
			values = append(values, "")
		}

		records = append(records, values)

		w := csv.NewWriter(os.Stdout)

		for _, record := range records {
			if err := w.Write(record); err != nil {
				return fmt.Errorf("%s: %w", constants.ErrorWritingCvsRecord, err)
			}
		}

		w.Flush()

		if err := w.Error(); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorFlush, err)
		}
	}

	fmt.Println()
	return nil
}

func EditTenantDescription(cmd *cobra.Command, args []string) error {
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

	if id == "" && name == "" {
		return fmt.Errorf("invalid tenant id or name: %w", err)
	}

	if conf, configPath, err = configuration.ReadConfig(cmd); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}
	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("error while generating access and refresh tokens: %w", err)
	}

	if len(args) != 1 {
		return fmt.Errorf("invalid image url: %w", err)
	}

	description := args[0]
	if len(description) > 200 {
		return fmt.Errorf("%s: %w", constants.ErrorTenantDescriptionSize, err)
	}

	if err = api.EditTenantDescription(conf.Urls, *accessToken, id, description); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenantDescription, err)
	}

	fmt.Printf("tenant %s description updated successfully\n", id)
	return nil
}

func EditTenantImage(cmd *cobra.Command, args []string) error {
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
	if id == "" && name == "" {
		return fmt.Errorf("%s: %w", constants.ErrorTenantNameOrID, err)
	}

	if conf, configPath, err = configuration.ReadConfig(cmd); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}
	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if len(args) != 1 {
		return fmt.Errorf("%s: %w", constants.ErrorImageURL, err)
	}

	image := args[0]
	if image != "" {
		if _, err := url.ParseRequestURI(image); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorImageURL, err)
		}
	}

	if err = api.EditTenantImage(conf.Urls, *accessToken, id, image); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorEditingTenant, err)
	}

	fmt.Printf("tenant %s image updated successfully\n", id)
	return nil
}

func ListAvailableSwarmsTenant(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, name, configPath string
	var conf *configuration.Config
	var swarms *api.SwarmList

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}
	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if id == "" && name == "" {
		return fmt.Errorf("%s: %w", constants.ErrorTenantNameOrID, err)
	}

	if conf, configPath, err = configuration.ReadConfig(cmd); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id == "" {
		if id, err = getTenantByName(conf, accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}
	}

	fmt.Printf("those are the swarms connect to the tenant %s\n", id)

	if conf, configPath, err = configuration.ReadConfig(cmd); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}
	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if swarms, err = api.ListAvailableTenantSwarms(conf.Urls, *accessToken, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingAvailableTenantSwarms, err)
	}

	for _, swarm := range swarms.Swarms {
		cross := " "
		if swarm.Default {
			cross = "x"
		}

		fmt.Printf("[%s] %s\n", cross, swarm.SwarmID)

	}
	return nil
}

func getTenantByName(conf *configuration.Config, accessToken *string, name string) (string, error) {
	var err error
	var operator *api.Operator
	var tenants *api.TenantList
	var id string

	if tenants, err = api.ListTenants(conf.Urls, *accessToken, operator.ID); err != nil {
		return "", fmt.Errorf("%s: %w", constants.ErrorRetrievingTenantList, err)
	}
	for _, tenant := range tenants.Tenants {
		if name == tenant.Name {
			id = tenant.ID
		}
	}
	if id == "" {
		return "", fmt.Errorf("tenant %s not found", name)
	}
	return id, nil
}
