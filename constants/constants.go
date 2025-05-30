// constants/constants.go
package constants

const (
	// Paths
	DefaultConfigFileName = "cubbit-session"
	DefaultFilePath       = "./"

	// Default Values
	DefaultProfile = "default"

	// Token
	AccessTokenName  = "_access"
	RefreshTokenName = "_refresh"
	Authorization    = "Authorization"
	Bearer           = "Bearer"

	// Base API URL
	BaseAPIURL = "https://api.eu00wi.cubbit.services"

	// Base DASH URL
	BaseDashURL = "https://dashboard.cubbit.eu"

	// Base Metrics URL
	BaseMetricsURL = "https://metrics.cubbit.io"

	// Base Swarm Gateway URL
	BaseSwarmGatewayURL = "https://swarm-gateway.cubbit.io"

	// API URIS
	BaseKeyvaultURI = "keyvault"
	BaseIamURI      = "/iam"
	BaseDashURI     = "/api"
	BaseChURI       = "/composer-hub"

	Invites = "/invites"

	Operators                      = "/v1/operators/"
	GenerateOperatorChallenge      = "/v1/auth/operators/signin/challenge"
	OperatorSignIn                 = "/v1/auth/operators/signin"
	CreateOperator                 = "/v1/operators/signup?secret="
	ForgeOperatorAccessToken       = "/v1/auth/operators/forge/access"
	RefreshOperatorAccessToken     = "/v1/auth/operators/refresh/access"
	ForgeOperatorDeleteTenantToken = "/v1/auth/operators/forge/token?capabilities=delete_tenant&tenant_id="

	AccountSignIn             = "/v1/auth/signin"
	GenerateAccountChallenge  = "/v1/auth/signin/challenge"
	CreateAccount             = "/v1/accounts/signup"
	RefreshAccountAccessToken = "/v1/auth/refresh/access"

	Tenants     = "/v1/tenants"
	TenantsV2   = "/v2/tenants"
	ListTenants = "/v1/tenants?owner="

	Swarms                        = "/v1/swarms"
	ForgeOperatorDeleteSwarmToken = "/v1/auth/operators/forge/token?capabilities=delete_swarm&swarm_id="

	Policies = "/v1/policies"

	Distributors                = "/v1/distributors"
	ForgeDistributorDeleteToken = "/v1/auth/operators/forge/token?capabilities=delete_distributor&distributor_id="

	Zones = "/v1/gateways/zones"

	Projects = "/v1/projects"

	DeleteTenantAccountToken = "/v1/auth/operators/forge/token?capabilities=delete_account&account_id="
	DeleteTenantProjectToken = "/v1/auth/operators/forge/token?capabilities=delete_project&project_id="
	DeleteSwarmNodeToken     = "/v1/auth/operators/forge/token?capabilities=delete_node&node_id="

	Nexuses = "/v1/nexuses"
	Nodes   = "/v1/nodes"
	NodesV2 = "/v2/nodes"

	RedundancyClasses = "/v1/redundancy_class"
	Rings             = "/v1/rings"

	AnsibleTarUrl = "https://cubbit-public.s3.cubbit.eu/agent/cubbit-agent-playbook.tar"
	MetricsSender = "/composer-hub/v2/agent/metrics"
)
