package describe

import (
	"fmt"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/utils/printer"
	"github.com/cubbit/composer-cli/utils/printer/tree"
	printerutils "github.com/cubbit/composer-cli/utils/printer/utils"
	"github.com/spf13/cobra"
)

func PrintTenantDetails(cmd *cobra.Command, handler configuration_handler.ConfigurationHandlerInterface, tenant api.TenantV5DTO) error {
	return printer.PrintTree(cmd, handler, tenant, func(t api.TenantV5DTO) []tree.TreeNode {
		return buildTenantTree(cmd, t)
	})
}

func buildTenantTree(cmd *cobra.Command, tenant api.TenantV5DTO) []tree.TreeNode {
	nodes := []tree.TreeNode{
		{Value: fmt.Sprintf("ID: %s", tenant.ID)},
		{Value: fmt.Sprintf("Slug: %s", tenant.Slug)},
		{Value: fmt.Sprintf("Description: %s", printerutils.FormatPtrString(tenant.Description))},
		{Value: fmt.Sprintf("Created at: %s", tenant.CreatedAt.Format("2006-01-02 15:04:05 UTC"))},
		{Value: fmt.Sprintf("ZK Enabled: %s", printerutils.FormatBool(tenant.ZKEnabled))},
		{
			Value:    "Resource Usage",
			Children: resourceUsageChildren(tenant.Storage, tenant.Bandwidth),
		},
		{
			Value:    "Settings",
			Children: settingsChildren(cmd, tenant.Settings),
		},
	}

	return []tree.TreeNode{
		{
			Value:    tenant.Name,
			Children: nodes,
		},
	}
}

func resourceUsageChildren(storage, bandwidth *api.UsageDTO) []tree.TreeNode {
	var children []tree.TreeNode

	if storage != nil {
		children = append(children, tree.TreeNode{
			Value: "Storage",
			Children: []tree.TreeNode{
				{Value: fmt.Sprintf("Consumed: %s", printerutils.FormatBytes(storage.Consumed))},
				{Value: fmt.Sprintf("Reserved: %s", printerutils.FormatBytes(storage.Reserved))},
				{Value: fmt.Sprintf("Used: %.1f%%", storage.Percentage*100)},
			},
		})
	} else {
		children = append(children, tree.TreeNode{Value: "Storage: N/A"})
	}

	if bandwidth != nil {
		children = append(children, tree.TreeNode{
			Value: "Bandwidth",
			Children: []tree.TreeNode{
				{Value: fmt.Sprintf("Consumed: %s", printerutils.FormatBytes(bandwidth.Consumed))},
				{Value: fmt.Sprintf("Reserved: %s", printerutils.FormatBytes(bandwidth.Reserved))},
				{Value: fmt.Sprintf("Used: %.1f%%", bandwidth.Percentage*100)},
			},
		})
	} else {
		children = append(children, tree.TreeNode{Value: "Bandwidth: N/A"})
	}

	return children
}

func settingsChildren(cmd *cobra.Command, settings *api.TenantSettings) []tree.TreeNode {
	showSecrets := printer.ShowSecrets(cmd)
	if settings == nil {
		return append(
			settingsFlatChildren(nil, showSecrets),
			tree.TreeNode{Value: "Project", Children: projectChildren(nil)},
			tree.TreeNode{Value: "Account", Children: accountChildren(nil, showSecrets)},
			tree.TreeNode{Value: "Notifications", Children: notificationsChildren(nil)},
			tree.TreeNode{Value: "White Label", Children: whiteLabelChildren(nil)},
		)
	}

	children := settingsFlatChildren(settings, showSecrets)
	children = append(children,
		tree.TreeNode{Value: "Project", Children: projectChildren(settings.Project)},
		tree.TreeNode{Value: "Account", Children: accountChildren(settings.Account, showSecrets)},
		tree.TreeNode{Value: "Notifications", Children: notificationsChildren(settings.Notifications)},
		tree.TreeNode{Value: "White Label", Children: whiteLabelChildren(settings.WhiteLabel)},
	)
	return children
}

func settingsFlatChildren(settings *api.TenantSettings, showSecrets bool) []tree.TreeNode {
	if settings == nil {
		return []tree.TreeNode{
			{Value: "Console URL: none"},
			{Value: "Gateway URL: none"},
			{Value: "Signup: enabled"},
			{Value: "Display name: none"},
			{Value: "Support link: none"},
			{Value: "Allowed domains: none"},
			{Value: "Blocked domains: none"},
			{Value: "Whitelabel: no"},
		}
	}

	return []tree.TreeNode{
		{Value: fmt.Sprintf("Console URL: %s", printerutils.FormatPtrString(settings.ConsoleUrl))},
		{Value: fmt.Sprintf("Gateway URL: %s", printerutils.FormatPtrString(settings.GatewayUrl))},
		{Value: fmt.Sprintf("Signup: %s", printerutils.FormatDisabledStatus(settings.SignupDisabled))},
		{Value: fmt.Sprintf("Display name: %s", printerutils.FormatPtrString(settings.DisplayName))},
		{Value: fmt.Sprintf("Support link: %s", printerutils.FormatPtrString(settings.SupportLink))},
		{Value: fmt.Sprintf("Allowed domains: %s", printerutils.FormatStringSlicePtr(settings.AllowedDomains))},
		{Value: fmt.Sprintf("Blocked domains: %s", printerutils.FormatStringSlicePtr(settings.BlockedDomains))},
		{Value: fmt.Sprintf("Whitelabel: %s", printerutils.FormatPtrBool(settings.WhitelabelEnabled))},
	}
}

func projectChildren(project *api.TenantSettingsProject) []tree.TreeNode {
	if project == nil {
		return []tree.TreeNode{
			{Value: "Max storage: none"},
			{Value: "Max egress bandwidth: none"},
		}
	}
	return []tree.TreeNode{
		{Value: fmt.Sprintf("Max storage: %s", formatBytesDetailPtr(project))},
		{Value: fmt.Sprintf("Max egress bandwidth: %s", formatBandwidthDetailPtr(project))},
	}
}

func accountChildren(account *api.TenantSettingsAccount, showSecrets bool) []tree.TreeNode {
	if account == nil {
		return []tree.TreeNode{
			{Value: "Max projects: none"},
			{Value: "Auth providers: none"},
			{Value: "Auth Providers Config", Children: authProvidersConfigChildren(nil)},
			{Value: "OAuth Providers: none"},
		}
	}

	children := []tree.TreeNode{
		{Value: fmt.Sprintf("Max projects: %s", formatMaxProjectPtr(account))},
		{Value: fmt.Sprintf("Auth providers: %s", formatAuthProvidersPtr(account))},
		{Value: "Auth Providers Config", Children: authProvidersConfigChildren(account.EnabledAuthProvidersConfig)},
	}

	if oauthChildren := oauthProvidersChildren(account.OAuthProviders, showSecrets); oauthChildren != nil {
		children = append(children, tree.TreeNode{Value: "OAuth Providers", Children: oauthChildren})
	} else {
		children = append(children, tree.TreeNode{Value: "OAuth Providers: none"})
	}

	return children
}

func authProvidersConfigChildren(cfg *api.EnabledAuthProvidersConfig) []tree.TreeNode {
	if cfg == nil {
		return []tree.TreeNode{
			{Value: "Google client ID: none"},
			{Value: "Microsoft application ID: none"},
			{Value: "Microsoft directory ID: none"},
		}
	}
	return []tree.TreeNode{
		{Value: fmt.Sprintf("Google client ID: %s", printerutils.FormatPtrString(cfg.GoogleClientID))},
		{Value: fmt.Sprintf("Microsoft application ID: %s", printerutils.FormatPtrString(cfg.MicrosoftApplicationID))},
		{Value: fmt.Sprintf("Microsoft directory ID: %s", printerutils.FormatPtrString(cfg.MicrosoftDirectoryID))},
	}
}

func oauthProvidersChildren(providers *[]api.OAuthProvider, showSecrets bool) []tree.TreeNode {
	if providers == nil || len(*providers) == 0 {
		return nil
	}

	var children []tree.TreeNode
	for _, p := range *providers {
		children = append(children, tree.TreeNode{
			Value: p.Name,
			Children: []tree.TreeNode{
				{Value: fmt.Sprintf("Client ID: %s", p.ClientID)},
				{Value: fmt.Sprintf("Client Secret: %s", printerutils.FormatSecret(&p.ClientSecret, showSecrets))},
				{Value: fmt.Sprintf("Issuer URL: %s", p.IssuerURL)},
				{Value: fmt.Sprintf("Redirect URL: %s", p.RedirectURL)},
			},
		})
	}
	return children
}

func notificationsChildren(n *api.MonitoredResourceForSeverityNotifications) []tree.TreeNode {
	if n == nil || len(*n) == 0 {
		return []tree.TreeNode{{Value: "none"}}
	}
	return []tree.TreeNode{{Value: "configured"}}
}

func whiteLabelChildren(wl *api.WhiteLabel) []tree.TreeNode {
	if wl == nil {
		return []tree.TreeNode{
			{Value: "DNS", Children: whiteLabelDNSChildren(nil)},
			{Value: "Email Domain", Children: whiteLabelEmailDomainChildren(nil)},
		}
	}
	return []tree.TreeNode{
		{Value: "DNS", Children: whiteLabelDNSChildren(wl.DNS)},
		{Value: "Email Domain", Children: whiteLabelEmailDomainChildren(wl.EmailDomain)},
	}
}

func whiteLabelDNSChildren(dns *api.WhiteLabelDNS) []tree.TreeNode {
	if dns == nil {
		return []tree.TreeNode{
			{Value: "Value: none"},
			{Value: "Challenge: none"},
			{Value: "Verified: no"},
		}
	}
	return []tree.TreeNode{
		{Value: fmt.Sprintf("Value: %s", dns.Value)},
		{Value: fmt.Sprintf("Challenge: %s", dns.Challenge)},
		{Value: fmt.Sprintf("Verified: %s", printerutils.FormatBool(dns.Verified))},
	}
}

func whiteLabelEmailDomainChildren(ed *api.WhiteLabelEmailDomain) []tree.TreeNode {
	if ed == nil {
		return []tree.TreeNode{
			{Value: "Value: none"},
			{Value: "Verified: no"},
		}
	}
	return []tree.TreeNode{
		{Value: fmt.Sprintf("Value: %s", ed.Value)},
		{Value: fmt.Sprintf("Verified: %s", printerutils.FormatBool(ed.Verified))},
	}
}
