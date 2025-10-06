// Package action provides CLI actions for managing gateways.
package action

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/cubbit/composer-cli/constants"
	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func CreateGateway(cmd *cobra.Command, args []string) error {
	var err error
	var tenantID, name, location string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs
	var gateway *api.GatewayWithGatewayTenant

	if tenantID, err = cmd.Flags().GetString("tenant-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if location, err = cmd.Flags().GetString("location"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	gatewayBodyRequest := api.CreateGatewayRequestBody{
		Name:          name,
		Location:      location,
		Configuration: map[string]interface{}{},
	}

	if gateway, err = api.CreateGateway(*urls, resolvedProfile.APIKey, tenantID, gatewayBodyRequest); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCreatingGatewayRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		[]*api.GatewayWithGatewayTenant{gateway},
		func(g *api.GatewayWithGatewayTenant) []string {
			return []string{
				g.Gateway.ID,
			}
		},
		&utils.SmartOutputConfig[*api.GatewayWithGatewayTenant]{
			SingleResourceCompactOutput: true,
			SingleResource:              true,
			DefaultOutput:               resolvedProfile.Output,
		})
}

func DescribeGateway(cmd *cobra.Command, args []string) error {
	var err error
	var tenantID, gatewayID string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs
	var gateway *api.Gateway

	if tenantID, err = cmd.Flags().GetString("tenant-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if gatewayID, err = cmd.Flags().GetString("gateway-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if gateway, err = api.GetGateway(*urls, resolvedProfile.APIKey, tenantID, gatewayID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingGatewayRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		[]*api.Gateway{gateway},
		func(g *api.Gateway) []string {
			return []string{
				g.ID,
				g.Name,
				g.Location,
			}
		},
		&utils.SmartOutputConfig[*api.Gateway]{
			DefaultOutput: resolvedProfile.Output,
		},
	)
}

func UpdateGateway(cmd *cobra.Command, args []string) error {
	var err error
	var tenantID, gatewayID, name, location string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs

	if tenantID, err = cmd.Flags().GetString("tenant-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if gatewayID, err = cmd.Flags().GetString("gateway-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if location, err = cmd.Flags().GetString("location"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	gatewayRequestBody := api.UpdateGatewayRequestBody{}

	if name != "" {
		gatewayRequestBody.Name = &name
	}

	if location != "" {
		gatewayRequestBody.Location = &location
	}

	if err = api.UpdateGateway(*urls, resolvedProfile.APIKey, tenantID, gatewayID, gatewayRequestBody); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorEditingGatewayRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		[]*api.UpdateGatewayRequestBody{&gatewayRequestBody},
		func(g *api.UpdateGatewayRequestBody) []string {
			return []string{
				gatewayID,
			}
		},
		&utils.SmartOutputConfig[*api.UpdateGatewayRequestBody]{
			SingleResource: true,
			DefaultOutput:  resolvedProfile.Output,
		},
	)
}

func ListGateways(cmd *cobra.Command, args []string) error {
	var err error
	var tenantID, sort, filter string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs
	var gateways *api.GenericPaginatedResponse[*api.Gateway]

	if tenantID, err = cmd.Flags().GetString("tenant-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if filter, err = cmd.Flags().GetString("filter"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if sort, err = cmd.Flags().GetString("sort"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if filter != "" {
		filter = utils.BuildFilterQuery(filter)
	}

	if gateways, err = api.ListGateways(*urls, resolvedProfile.APIKey, tenantID, sort, filter); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingGatewaysRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		gateways.Data,
		func(g *api.Gateway) []string {
			return []string{
				g.ID,
			}
		},
		&utils.SmartOutputConfig[*api.Gateway]{
			DefaultOutput: resolvedProfile.Output,
		},
	)
}

func RemoveGateway(cmd *cobra.Command, args []string) error {
	var err error
	var tenantID, gatewayID string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs

	if tenantID, err = cmd.Flags().GetString("tenant-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if gatewayID, err = cmd.Flags().GetString("gateway-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if err = api.DeleteGateway(*urls, resolvedProfile.APIKey, tenantID, gatewayID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDeletingGatewayRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		[]string{gatewayID},
		func(s string) []string { return []string{s} },
		&utils.SmartOutputConfig[string]{
			SingleResource:              true,
			SingleResourceCompactOutput: true,
			DefaultOutput:               resolvedProfile.Output,
		},
	)
}

func InstallGateway(cmd *cobra.Command, args []string) error {
	var err error
	var tenantID, gatewayID string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs
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

	if tenantID, err = cmd.Flags().GetString("tenant-id"); err != nil {
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

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if tenant, err = api.GetTenant(*urls, resolvedProfile.APIKey, tenantID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
	}

	if gateway, err = api.GetGateway(*urls, resolvedProfile.APIKey, tenantID, gatewayID); err != nil {
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
	var coordinatorDomain string = urls.BaseURL

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

	if urls.BaseURL != "" {
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
