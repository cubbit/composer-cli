// constants/constants.go
package constants

const (
	// Paths
	DefaultConfigFileName = "cubbit-session"
	DefaultFilePath       = "./"

	// Default Values
	DefaultProfile = "default"

	// Authentication
	Authorization = "Authorization"
	APIKey        = "ApiKey"
	Bearer        = "Bearer"
	RefreshCookie = "_refresh"

	// Base API URL
	BaseAPIURL = "https://api.eu00wi.cubbit.services"

	// Base DASH URL
	BaseDashURL = "https://dashboard.cubbit.eu"

	// API URIS
	BaseKeyvaultURI = "keyvault"
	BaseIamURI      = "/iam"
	BaseDashURI     = "/api"
	BaseChURI       = "/composer-hub"

	AnsibleTarURL = "https://cubbit-public.s3.cubbit.eu/agent/cubbit-agent-playbook.tar"
	MetricsSender = "/composer-hub/v2/agent/metrics"
)
