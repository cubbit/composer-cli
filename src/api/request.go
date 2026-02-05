package api

type UpdateAccountRequest struct {
	FirstName          *string `json:"first_name"`
	LastName           *string `json:"last_name"`
	EndpointGateway    *string `json:"endpoint_gateway"`
	Internal           *bool   `json:"internal"`
	MaxAllowedProjects *int    `json:"max_allowed_projects"`
}

type UpdateTenantProjectRequestBody struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	ImageURL    *string `json:"image_url"`
}

type CreateNexusRequestBody struct {
	Name        string `json:"name" `
	Description string `json:"description,omitempty"`
	Location    string `json:"location"`
	ProviderID  string `json:"provider_id,omitempty"`
}

type UpdateNexusRequestBody struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Location    string `json:"location"`
}

type CreateNodeBodyRequest struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	NexusID     string `json:"nexus_id"`
	SecretID    string `json:"secret_id"`
}

type UpdateNodeBodyRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateRedundancyClassRequestBody struct {
	Name              string   `json:"name"`
	Description       string   `json:"description"`
	InnerK            int      `json:"inner_k"`
	InnerN            int      `json:"inner_n"`
	OuterK            int      `json:"outer_k"`
	OuterN            int      `json:"outer_n"`
	AntiAffinityGroup int      `json:"anti_affinity_group"`
	Nexuses           []string `json:"nexuses"`
}

type BulkInsertNewNodeRequestBody struct {
	Nodes []CreateNewNodeRequestBody `json:"nodes"`
}

type BulkInsertNewAgentRequestBody struct {
	Agents []CreateNewAgentRequestBody `json:"agents"`
}

type CreateNewNodeRequestBody struct {
	Name          string                      `json:"name"`
	Label         *string                     `json:"label"`
	Configuration map[string]interface{}      `json:"config" swaggertype:"object,string"`
	PublicIP      string                      `json:"public_ip"`
	PrivateIP     string                      `json:"private_ip"`
	Agents        []CreateNewAgentRequestBody `json:"agents"`
}

type CreateNewAgentRequestBody struct {
	Port     int                    `json:"port"`
	Features map[string]interface{} `json:"features"`
	Volume   AgentVolume            `json:"volume"`
}

type UpdateNewNodeRequestBody struct {
	Name      *string `json:"name"`
	Label     *string `json:"label"`
	PublicIP  *string `json:"public_ip"`
	PrivateIP *string `json:"private_ip"`
}

type UpdateNewAgentRequestBody struct {
	Port     *int                      `json:"port"`
	Features map[string]interface{}    `json:"features"`
	Volume   *UpdateAgentVolumeRequest `json:"volume"`
}

type UpdateAgentVolumeRequest struct {
	MountPoint *string `json:"mount_point"`
	Disk       *string `json:"disk"`
}

type CreateGatewayRequestBody struct {
	Name          string                 `json:"name"`
	Location      string                 `json:"location"`
	Configuration map[string]interface{} `json:"configuration"`
}

type UpdateGatewayRequestBody struct {
	Name     *string `json:"name"`
	Location *string `json:"location"`
}

type UpdateSwarmRequest struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

type CreateSwarmRequest struct {
	Name          string                 `json:"name"`
	Configuration map[string]interface{} `json:"configuration"`
	Description   *string                `json:"description"`
}

type UpdateTenantRequestBody struct {
	Description *string         `json:"description"`
	Settings    *TenantSettings `json:"settings"`
}

type CreateTenantRequestBody struct {
	Name        string          `json:"name"`
	Description *string         `json:"description"`
	Settings    *TenantSettings `json:"settings"`
	CouponCode  *string         `json:"coupon_code"`
	Zone        *string         `json:"zone"`
	ZKEnabled   *bool           `json:"zk_enabled"`
}

type ChangeOperatorPolicyRequestBody struct {
	PolicyID string `json:"policy_id"`
}

type InviteOperatorRequestBody struct {
	Email     string  `json:"email"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	PolicyID  string  `json:"policy_id"`
}

// #region auth

type SignUpRequestBody struct {
	Operator     CreateOperatorRequestBodyV2     `json:"operator"`
	Organization CreateOrganizationRequestBodyV2 `json:"organization"`
}

type CreateOrganizationRequestBodyV2 struct {
	Name       string                 `json:"name"`
	BasePolicy map[string]interface{} `json:"base_policy"`
	Settings   map[string]interface{} `json:"settings"`
}

type CreateOperatorRequestBodyV2 struct {
	Email                   string  `json:"email"`
	FirstName               *string `json:"first_name"`
	LastName                *string `json:"last_name"`
	Username                string  `json:"username"`
	AuthenticationPublicKey *string `json:"authentication_public_key"`
}

func CreateSignUpRequestBody(
	email string,
	username string,
	firstName *string,
	lastName *string,
	authenticationPublicKey *string,
	organizationName string,
	organizationBasePolicy map[string]interface{},
	organizationSettings map[string]interface{},
) SignUpRequestBody {
	return SignUpRequestBody{
		Operator: CreateOperatorRequestBodyV2{
			Email:                   email,
			Username:                username,
			FirstName:               firstName,
			LastName:                lastName,
			AuthenticationPublicKey: authenticationPublicKey,
		},
		Organization: CreateOrganizationRequestBodyV2{
			Name:       organizationName,
			BasePolicy: organizationBasePolicy,
			Settings:   organizationSettings,
		},
	}
}

type ChallengeRequestBodyV3 struct {
	Email            *string `json:"email,omitempty"`
	Username         *string `json:"username,omitempty"`
	OrganizationName *string `json:"organization_name,omitempty"`
}

func CreateChallengeRequestBodyV3(
	email *string,
	username *string,
	organizationName *string,
) ChallengeRequestBodyV3 {
	return ChallengeRequestBodyV3{
		Email:            email,
		Username:         username,
		OrganizationName: organizationName,
	}
}

// #endregion
