// constants/constants.go
package constants

const (
	// Paths
	DefaultFilePath = "./"

	// Default Values
	DefaultProfile = "default"

	// Token
	AccessTokenName  = "_access"
	RefreshTokenName = "_refresh"
	Authorization    = "Authorization"
	Bearer           = "Bearer"

	// Base API URL
	BaseAPIURL = "https://api.cubbit.eu"

	// API URIS
	BaseKeyvaultURI = "keyvault"
	BaseIamURI      = "/iam"
	BaseHiveURI     = "/hive"
	BaseDashURI     = "/api"

	Invites = "/invites"

	Operators                      = "/v1/operators/"
	GenerateOperatorChallenge      = "/v1/auth/operators/signin/challenge"
	OperatorSignIn                 = "/v1/auth/operators/signin"
	CreateOperator                 = "/v1/operators/signup?secret="
	ForgeOperatorAccessToken       = "/v1/auth/operators/forge/access"
	RefreshOperatorAccessToken     = "/v1/auth/operators/refresh/access"
	ForgeOperatorDeleteTenantToken = "/v1/auth/operators/forge/token?capabilities=delete_tenant&tenant_id="

	AccountSignIn            = "/v1/auth/signin"
	GenerateAccountChallenge = "/v1/auth/signin/challenge"
	CreateAccount            = "/v1/accounts/signup"

	Tenants     = "/v1/tenants"
	ListTenants = "/v1/tenants?owner="

	Swarms                        = "/v1/swarms"
	ForgeOperatorDeleteSwarmToken = "/v1/auth/operators/forge/token?capabilities=delete_swarm&swarm_id="

	Policies = "/v1/policies"

	Distributors                = "/v1/distributors"
	ForgeDistributorDeleteToken = "/v1/auth/operators/forge/token?capabilities=delete_distributor&distributor_id="

	Zones = "/v1/gateways/zones"
)
