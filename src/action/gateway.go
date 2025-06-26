package action

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/api"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/utils"
	"github.com/spf13/cobra"
)

func CreateTenantGateway(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, gatewayName, gatewayLocation, configPath string
	var conf *configuration.Config
	var tenant *api.Tenant
	var tenantGateway *api.GatewayWithGatewayTenant

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if gatewayName, err = cmd.Flags().GetString("gateway-name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if gatewayLocation, err = cmd.Flags().GetString("gateway-location"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id == "" {
		if tenant, err = getTenantByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}
		id = tenant.ID
	}

	gatewayBodyRequest := api.CreateGatewayRequestBody{
		Name:          gatewayName,
		Location:      gatewayLocation,
		Configuration: map[string]interface{}{},
	}

	if tenantGateway, err = api.CreateGateway(conf.Urls, *accessToken, id, gatewayBodyRequest); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCreatingGatewayRequest, err)
	}

	utils.PrintCreateSuccess("tenant gateway", tenantGateway.Gateway.ID)

	return nil
}

func ListTenantGateways(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, configPath, sort, filter string
	var conf *configuration.Config
	var gateways *api.GenericPaginatedResponse[*api.Gateway]
	var tenant *api.Tenant
	var verbose, l bool

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

	if gateways, err = api.ListGateways(conf.Urls, *accessToken, id, sort, filter); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingGatewaysRequest, err)
	}

	if verbose, err = cmd.Flags().GetBool("verbose"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if l, err = cmd.Flags().GetBool("line"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if len(gateways.Data) == 0 {
		utils.PrintEmptyList()
		return nil
	}

	utils.PrintList("Your Tenant Gateways List")

	if verbose {
		utils.PrintVerbose(gateways.Data, l)
	} else {
		var IDs []string
		for _, gateway := range gateways.Data {
			IDs = append(IDs, gateway.ID)
		}
		utils.PrintSimpleList(IDs)
	}

	return nil
}

func DescribeTenantGateway(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, gatewayID, configPath, format string
	var conf *configuration.Config
	var tenant *api.Tenant
	var gateway *api.Gateway

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if gatewayID, err = cmd.Flags().GetString("gateway-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id == "" {
		if tenant, err = getTenantByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}

		id = tenant.ID
	}

	if gateway, err = api.GetGateway(conf.Urls, *accessToken, id, gatewayID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingGatewayRequest, err)
	}

	if format, err = cmd.Flags().GetString("format"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	utils.PrintFormattedData(*gateway, format)

	return nil
}

func RemoveTenantGateway(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, gatewayID, email, password, code, configPath, deleteTenantGatewayToken string
	var conf *configuration.Config
	var challenge *api.ChallengeResponseModel
	var tenants *api.GenericPaginatedResponse[*api.Tenant]

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if gatewayID, err = cmd.Flags().GetString("gateway-id"); err != nil {
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

	if deleteTenantGatewayToken, err = api.ForgeOperatorDeleteTenantGatewayToken(conf.Urls, email, password, conf.RefreshToken, challenge, code, gatewayID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorForgingTenantGatewayDeleteTokenRequest, err)
	}

	if err = api.DeleteGateway(conf.Urls, *accessToken, id, gatewayID, deleteTenantGatewayToken); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDeletingGatewayRequest, err)
	}

	utils.PrintDelete(fmt.Sprintf("gateway %s removed successfully", gatewayID))

	return nil
}

func UpdateTenantGateway(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, gatewayID, gatewayName, gatewayLocation, configPath string
	var conf *configuration.Config
	var tenant *api.Tenant

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if gatewayID, err = cmd.Flags().GetString("gateway-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if gatewayName, err = cmd.Flags().GetString("gateway-name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if gatewayLocation, err = cmd.Flags().GetString("gateway-location"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id == "" {
		if tenant, err = getTenantByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}
		id = tenant.ID
	}

	gatewayRequestBody := api.UpdateGatewayRequestBody{}

	if gatewayName != "" {
		gatewayRequestBody.Name = &gatewayName
	}

	if gatewayLocation != "" {
		gatewayRequestBody.Location = &gatewayLocation
	}

	if err = api.UpdateGateway(conf.Urls, *accessToken, id, gatewayID, gatewayRequestBody); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorEditingGatewayRequest, err)
	}

	utils.PrintSuccess("tenant gateway updated successfully")
	return nil
}

func ListTenantGatewayInstances(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, gatewayID, configPath string
	var conf *configuration.Config
	var gatewayInstances *api.GatewayInstanceListResponse
	var tenant *api.Tenant
	var verbose, l bool

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if gatewayID, err = cmd.Flags().GetString("gateway-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id == "" {
		if tenant, err = getTenantByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}
		id = tenant.ID
	}

	if gatewayInstances, err = api.ListGatewayInstances(conf.Urls, *accessToken, id, gatewayID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingGatewayInstancesRequest, err)
	}

	if verbose, err = cmd.Flags().GetBool("verbose"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if l, err = cmd.Flags().GetBool("line"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if len(gatewayInstances.Data) == 0 {
		utils.PrintEmptyList()
		return nil
	}

	utils.PrintList("Your Tenant Gateway Instances List")

	if verbose {
		utils.PrintVerbose(gatewayInstances.Data, l)
	} else {
		var IDs []string
		for _, gateway := range gatewayInstances.Data {
			IDs = append(IDs, gateway.ID)
		}
		utils.PrintSimpleList(IDs)
	}

	return nil
}

func ConfigureTenantDNS(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, domain, configPath string
	var conf *configuration.Config
	var tenant *api.Tenant
	var force bool

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if domain, err = cmd.Flags().GetString("domain"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if force, err = cmd.Flags().GetBool("force"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id == "" {
		if tenant, err = getTenantByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}
		id = tenant.ID
	} else {
		if tenant, err = api.GetTenant(conf.Urls, *accessToken, id); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}
	}

	if tenant.Settings != nil {
		if tenant.Settings.WhiteLabel != nil && tenant.Settings.WhiteLabel.DNS != nil && tenant.Settings.WhiteLabel.DNS.Challenge != "" && !force {
			return fmt.Errorf("%s: %w", constants.ErrorTenantDNSAlreadyConfigured, fmt.Errorf("domain %s is already configured for tenant %s, add --force to override", tenant.Settings.WhiteLabel.DNS.Value, tenant.Name))

		}
	}

	tenantSettings := api.TenantSettings{
		WhiteLabel: &api.WhiteLabel{
			DNS: &api.WhiteLabelDNS{
				Value: domain,
			},
		},
	}

	if err = api.EditTenantSettings(conf.Urls, *accessToken, id, tenantSettings); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorWhileConfiguringTenantDNS, err)
	}

	if tenant, err = api.GetTenant(conf.Urls, *accessToken, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
	}

	utils.PrintSuccess(fmt.Sprintf("tenant %s DNS configured successfully", tenant.Name))
	utils.PrintInfo(fmt.Sprintf("Please add a TXT record named _acme-challenge with the following value: %s", tenant.Settings.WhiteLabel.DNS.Challenge))
	return nil
}

func VerifyTenantDNS(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, configPath string
	var conf *configuration.Config
	var tenant *api.Tenant

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
		if tenant, err = getTenantByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}
		id = tenant.ID
	} else {
		if tenant, err = api.GetTenant(conf.Urls, *accessToken, id); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}
	}

	if err = api.VerifyDNS(conf.Urls, *accessToken, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorWhileVerifyingTenantDNS, err)
	}

	utils.PrintSuccess(fmt.Sprintf("tenant %s DNS verified successfully", tenant.Name))

	return nil
}

func InstallTenantGateway(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, gatewayID, configPath string
	var conf *configuration.Config
	var tenant *api.Tenant
	var gateway *api.Gateway

	// Installation configuration flags
	var cachePath string
	var enableTLS bool
	var certRootPath string
	var initNode bool
	var setupInfra bool
	var setupApp bool
	var setupConsole bool
	var setupOffloader bool
	var setupS3 bool
	var installOnlyIngress bool

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}
	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}
	if gatewayID, err = cmd.Flags().GetString("gateway-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if cachePath, err = cmd.Flags().GetString("cache"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	var noTLS, noInit, noInfra, noApp, noConsole, noOffloader, noS3 bool
	if noTLS, err = cmd.Flags().GetBool("no-tls"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}
	if certRootPath, err = cmd.Flags().GetString("cert-root"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}
	if noInit, err = cmd.Flags().GetBool("no-init"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}
	if noInfra, err = cmd.Flags().GetBool("no-infra"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}
	if noApp, err = cmd.Flags().GetBool("no-app"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}
	if noConsole, err = cmd.Flags().GetBool("no-console"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}
	if noOffloader, err = cmd.Flags().GetBool("no-offloader"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}
	if noS3, err = cmd.Flags().GetBool("no-s3"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}
	if installOnlyIngress, err = cmd.Flags().GetBool("ingress"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id == "" {
		if tenant, err = getTenantByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}
		id = tenant.ID
	} else {
		if tenant, err = api.GetTenant(conf.Urls, *accessToken, id); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}
	}

	if gateway, err = api.GetGateway(conf.Urls, *accessToken, id, gatewayID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingGatewayRequest, err)
	}

	enableTLS = !noTLS
	initNode = !noInit
	setupInfra = !noInfra
	setupApp = !noApp
	setupConsole = !noConsole
	setupOffloader = !noOffloader
	setupS3 = !noS3

	if certRootPath == "" {
		certRootPath = "./cert"
	}

	utils.PrintInfo("🚀 Starting Gateway Installation")
	utils.PrintEmptyLine()

	var domain string
	var coordinatorDomain string = conf.Urls.BaseURL

	if tenant.Settings != nil && tenant.Settings.WhiteLabel != nil && tenant.Settings.WhiteLabel.DNS != nil {
		domain = tenant.Settings.WhiteLabel.DNS.Value
	}

	utils.PrintInfo("📋 Installation Configuration:")
	utils.PrintInfo(fmt.Sprintf("  Tenant: %s (%s)", tenant.Name, tenant.ID))
	utils.PrintInfo(fmt.Sprintf("  Gateway: %s (%s)", gateway.Name, gateway.ID))
	utils.PrintInfo(fmt.Sprintf("  Domain: %s", domain))
	utils.PrintInfo(fmt.Sprintf("  Gateway Secret: %s", gateway.Secret))
	utils.PrintInfo(fmt.Sprintf("  Cache Path: %s", cachePath))
	utils.PrintInfo(fmt.Sprintf("  Coordinator Domain: %s", coordinatorDomain))
	utils.PrintInfo(fmt.Sprintf("  Enable TLS: %t", enableTLS))
	if enableTLS {
		utils.PrintInfo(fmt.Sprintf("  Certificate Root Path: %s", certRootPath))
	}
	utils.PrintInfo(fmt.Sprintf("  Initialize Node: %t", initNode))
	utils.PrintInfo(fmt.Sprintf("  Setup Infrastructure: %t", setupInfra))
	utils.PrintInfo(fmt.Sprintf("  Setup Application: %t", setupApp))
	utils.PrintInfo(fmt.Sprintf("  Setup Console: %t", setupConsole))
	utils.PrintInfo(fmt.Sprintf("  Setup Offloader: %t", setupOffloader))
	utils.PrintInfo(fmt.Sprintf("  Setup S3: %t", setupS3))
	utils.PrintInfo(fmt.Sprintf("  Install Only Ingress: %t", installOnlyIngress))
	utils.PrintEmptyLine()

	var curlArgs []string
	curlArgs = append(curlArgs, "--domain", domain)
	curlArgs = append(curlArgs, "--tenant-name", tenant.Name)
	curlArgs = append(curlArgs, "--tenant-id", tenant.ID)
	curlArgs = append(curlArgs, "--gateway-id", gateway.ID)
	curlArgs = append(curlArgs, "--gateway-secret", gateway.Secret)
	curlArgs = append(curlArgs, "--coordinator-domain", coordinatorDomain)

	if cachePath != "" {
		curlArgs = append(curlArgs, "--cache", cachePath)
	}

	if !enableTLS {
		curlArgs = append(curlArgs, "--no-tls")
	} else {
		curlArgs = append(curlArgs, "--cert-root", certRootPath)
	}

	if !initNode {
		curlArgs = append(curlArgs, "--no-init")
	}

	if !setupInfra {
		curlArgs = append(curlArgs, "--no-infra")
	}

	if !setupApp {
		curlArgs = append(curlArgs, "--no-app")
	}

	if !setupConsole {
		curlArgs = append(curlArgs, "--no-console")
	}

	if !setupOffloader {
		curlArgs = append(curlArgs, "--no-offloader")
	}

	if !setupS3 {
		curlArgs = append(curlArgs, "--no-s3")
	}

	if installOnlyIngress {
		curlArgs = append(curlArgs, "--ingress")
	}

	utils.PrintInfo("🚀 Starting gateway installation...")
	utils.PrintEmptyLine()

	curlCommand := fmt.Sprintf("curl -fSsL https://installer.s3.cubbit.eu/gateway/installer.sh | bash -s -- %s", strings.Join(curlArgs, " "))

	utils.PrintInfo("Executing command:")
	utils.PrintInfo(curlCommand)
	utils.PrintEmptyLine()

	command := exec.Command("bash", "-c", curlCommand)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Stdin = os.Stdin

	if conf.Urls.BaseURL != "" {
		if err = command.Run(); err != nil {
			return fmt.Errorf("installation failed: %w", err)
		}
	} else {
		utils.PrintInfo("Skipping command execution, as we are in a development environment")
		return nil
	}

	utils.PrintSuccess("Gateway installation command executed successfully.")

	return nil
}
