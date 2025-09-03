// Package action provides CLI actions for managing gateways.
package action

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/api"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/src/tui"
	"github.com/cubbit/cubbit/client/cli/utils"
	"github.com/spf13/cobra"
)

func ConfigureAndVerifyDNSInteractive(cmd *cobra.Command) error {
	var err error
	var id, domain string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs
	var choice string
	var choices []string
	var tenants *api.GenericPaginatedResponse[*api.Tenant]
	var tenant *api.Tenant
	var force bool

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if tenants, err = api.ListTenants(*urls, resolvedProfile.APIKey, "", ""); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingTenantsRequest, err)
	}

	for _, t := range tenants.Data {
		choices = append(choices, fmt.Sprintf("• %s, %s, %s", t.ID, t.Name, utils.StringOrEmpty(t.Description)))
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

	if tenant, err = api.GetTenant(*urls, resolvedProfile.APIKey, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
	}

	var dnsAlreadyConfigured bool
	var currentDomain string
	if tenant.Settings != nil && tenant.Settings.WhiteLabel != nil && tenant.Settings.WhiteLabel.DNS != nil && tenant.Settings.WhiteLabel.DNS.Challenge != "" {
		dnsAlreadyConfigured = true
		currentDomain = tenant.Settings.WhiteLabel.DNS.Value
	}

	var shouldConfigure bool

	if dnsAlreadyConfigured && !force {

		configChoices := []string{"Skip configuration and proceed to verification", "Reconfigure DNS (override existing)", "Cancel"}
		if choice, err = tui.ChooseOne("DNS is already configured for this tenant with the domain: "+currentDomain+"\nWhat would you like to do?", false, false, configChoices); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
		}

		switch choice {
		case "Skip configuration and proceed to verification":
			shouldConfigure = false
			if err = api.VerifyDNS(*urls, resolvedProfile.APIKey, id); err != nil {
				return fmt.Errorf("%s: %w", constants.ErrorWhileVerifyingTenantDNS, err)
			}
			utils.PrintSuccess(fmt.Sprintf("Tenant %s DNS verified successfully", tenant.Name))
			return nil
		case "Reconfigure DNS (override existing)":
			shouldConfigure = true
			force = true
		case "Cancel":
			utils.PrintInfo("Operation cancelled")
			return nil
		}
	} else {
		shouldConfigure = true
	}

	if shouldConfigure {
		if domain == "" {
			if _, err = tui.TextInputs(
				"Fill in the form bellow",
				false,
				tui.Input{Placeholder: "Domain for DNS configuration*", IsPassword: false, Value: &domain},
			); err != nil {
				return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)

			}
		}

		tenantSettings := api.TenantSettings{
			WhiteLabel: &api.WhiteLabel{
				DNS: &api.WhiteLabelDNS{
					Value: domain,
				},
			},
		}

		if err = api.EditTenantSettings(*urls, resolvedProfile.APIKey, id, tenantSettings); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorWhileConfiguringTenantDNS, err)
		}

		if tenant, err = api.GetTenant(*urls, resolvedProfile.APIKey, id); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}

		utils.PrintSuccess(fmt.Sprintf("Tenant %s DNS configured successfully", tenant.Name))
		utils.PrintInfo(fmt.Sprintf("Please add a TXT record named _acme-challenge with the following value: %s", tenant.Settings.WhiteLabel.DNS.Challenge))
		utils.PrintHint("Once the DNS record is added, wait a couple of minutes for the DNS to propagate, you can then verify it using the 'Verify DNS now' option.")
		utils.PrintEmptyLine()

		verifyChoices := []string{"Verify DNS now", "Skip verification"}
		if choice, err = tui.ChooseOne("What would you like to do next?", false, false, verifyChoices); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
		}

		switch choice {
		case "Verify DNS now":
			if err = api.VerifyDNS(*urls, resolvedProfile.APIKey, id); err != nil {
				return fmt.Errorf("%s: %w", constants.ErrorWhileVerifyingTenantDNS, err)
			}
			utils.PrintSuccess(fmt.Sprintf("Tenant %s DNS verified successfully", tenant.Name))
		case "Skip verification":
			utils.PrintInfo("DNS verification skipped")
		}
	}

	return nil
}

func InstallTenantGatewayInteractive(cmd *cobra.Command) error {
	var err error
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs
	var choice string
	var choices []string
	var tenants *api.GenericPaginatedResponse[*api.Tenant]
	var tenant *api.Tenant
	var gateways *api.GenericPaginatedResponse[*api.Gateway]
	var gateway *api.Gateway

	// Installation configuration
	var cachePath string
	var enableTLS bool
	var certRootPath string = "./cert" // default value
	var initNode bool = true
	var setupInfra bool = true
	var setupApp bool = true
	var setupConsole bool = true
	var setupOffloader bool = true
	var setupS3 bool = true
	var installOnlyIngress bool = false

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	utils.PrintInfo("🚀 Starting Gateway Installation")
	utils.PrintEmptyLine()

	utils.PrintInfo("Step 1: Choose Tenant")
	if tenants, err = api.ListTenants(*urls, resolvedProfile.APIKey, "", ""); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingTenantsRequest, err)
	}

	choices = []string{}
	for _, t := range tenants.Data {
		choices = append(choices, fmt.Sprintf("• %s, %s, %s", t.ID, t.Name, utils.StringOrEmpty(t.Description)))
	}

	if len(choices) == 0 {
		utils.PrintNotFound("No tenants found")
		return nil
	}

	if choice, err = tui.ChooseOne("Choose your tenant", false, false, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
	}

	_, withoutPrefix, _ := strings.Cut(choice, " ")
	tenantID, _, _ := strings.Cut(withoutPrefix, ",")

	if tenant, err = api.GetTenant(*urls, resolvedProfile.APIKey, tenantID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
	}

	utils.PrintSuccess(fmt.Sprintf("Selected tenant: %s", tenant.Name))
	utils.PrintEmptyLine()

	utils.PrintInfo("Step 2: Choose Gateway")
	if gateways, err = api.ListGateways(*urls, resolvedProfile.APIKey, tenantID, "", ""); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingGatewaysRequest, err)
	}

	choices = []string{}
	for _, g := range gateways.Data {
		choices = append(choices, fmt.Sprintf("• %s, %s", g.ID, g.Name))
	}

	if len(choices) == 0 {
		utils.PrintNotFound("No gateways found for this tenant")
		return nil
	}

	if choice, err = tui.ChooseOne("Choose your gateway", false, false, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingGatewayRequest, err)
	}

	_, withoutPrefix, _ = strings.Cut(choice, " ")
	gatewayID, _, _ := strings.Cut(withoutPrefix, ",")

	if gateway, err = api.GetGateway(*urls, resolvedProfile.APIKey, tenantID, gatewayID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingGatewayRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("Selected gateway: %s", gateway.Name))
	utils.PrintEmptyLine()

	utils.PrintInfo("Step 3: Configure Cache Path")
	if _, err = tui.TextInputs(
		"Configure cache settings",
		false,
		tui.Input{Placeholder: "Cache Path", IsPassword: false, Value: &cachePath},
	); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if cachePath != "" {
		utils.PrintSuccess(fmt.Sprintf("Cache path set to: %s", cachePath))
		utils.PrintEmptyLine()
	} else {
		utils.PrintInfo("No cache path provided, skipping cache configuration.")
		utils.PrintEmptyLine()
	}

	utils.PrintInfo("Step 4: Select Installation Options")

	// Enable TLS
	tlsChoices := []string{"Yes", "No"}
	if choice, err = tui.ChooseOne("Enable TLS?", false, false, tlsChoices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}
	enableTLS = (choice == "Yes")

	if enableTLS {
		if _, err = tui.TextInputs(
			"TLS Configuration",
			false,
			tui.Input{Placeholder: fmt.Sprintf("Certificate Root Path (default: %s)", certRootPath), IsPassword: false, Value: &certRootPath},
		); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
		}
		if certRootPath == "" {
			certRootPath = "./cert"
		}
	}

	// More installation options
	yesNoChoices := []string{"Yes", "No"}

	if choice, err = tui.ChooseOne("Initialize Node?", false, false, yesNoChoices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}
	initNode = (choice == "Yes")

	if choice, err = tui.ChooseOne("Setup Infrastructure?", false, false, yesNoChoices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}
	setupInfra = (choice == "Yes")

	if choice, err = tui.ChooseOne("Setup Application?", false, false, yesNoChoices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}
	setupApp = (choice == "Yes")

	if choice, err = tui.ChooseOne("Setup Console?", false, false, yesNoChoices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}
	setupConsole = (choice == "Yes")

	if choice, err = tui.ChooseOne("Setup Offloader?", false, false, yesNoChoices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}
	setupOffloader = (choice == "Yes")

	if choice, err = tui.ChooseOne("Setup S3?", false, false, yesNoChoices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}
	setupS3 = (choice == "Yes")

	if choice, err = tui.ChooseOne("Install Only Ingress?", false, false, yesNoChoices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}
	installOnlyIngress = (choice == "Yes")

	utils.PrintEmptyLine()

	utils.PrintInfo("Step 5: Installation Summary")
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

	confirmChoices := []string{"Proceed with installation", "Cancel"}
	if choice, err = tui.ChooseOne("Do you want to proceed with the installation?", false, false, confirmChoices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if choice == "Cancel" {
		utils.PrintInfo("Installation cancelled")
		return nil
	}

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
