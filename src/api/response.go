package api

import (
	"errors"
	"fmt"
	"time"
)

type ErrorResponseModel struct {
	Message string `json:"message"`
	Param   string `json:"param"`
}

type ChallengeResponseModel struct {
	Challenge string `json:"challenge"`
	Salt      string `json:"salt" example:"SGVsbG8gd29ybGQ="`
}

var ErrorBadRequest = errors.New("bad request")
var ErrorInternalServerError = errors.New("internal server error")
var ErrorNotFound = errors.New("not found")
var ErrorUnauthorized = errors.New("unauthorized")
var ErrorUniqueConstraintViolation = errors.New("unique constraint violation")

var ErrorBadFormat = errors.New("unique constraint violation")

var ErrorBadRequestResponseModel ErrorResponseModel = ErrorResponseModel{
	Message: ErrorBadRequest.Error(),
}

var ErrorInternalServerErrorResponseModel ErrorResponseModel = ErrorResponseModel{
	Message: ErrorInternalServerError.Error(),
}

var ErrorNotFoundResponseModel ErrorResponseModel = ErrorResponseModel{
	Message: ErrorNotFound.Error(),
}

var ErrorUnauthorizedResponseModel ErrorResponseModel = ErrorResponseModel{
	Message: ErrorUnauthorized.Error(),
}

var ErrorUniqueConstraintViolationResponseModel ErrorResponseModel = ErrorResponseModel{
	Message: ErrorUniqueConstraintViolation.Error(),
}

type TokenAndExpirationResponseModel struct {
	Token   string    `json:"token" example:"SGVsbG8gd29ybGQ="`
	Exp     int       `json:"exp" example:"3600000"`
	ExpDate time.Time `json:"expDate" example:"2022-12-27T11:21:23.478555Z"`
}

type GenericIDResponseModel struct {
	ID string `json:"id" example:"8da936da-f50d-42ce-bd3a-0764bd8595ee"`
}

type Tenant struct {
	ID          string          `json:"id" example:"784da023-cc6c-4a46-8e31-bc0a521d19e0"`
	Name        string          `json:"name" example:"Cubbit"`
	Description *string         `json:"description" example:"Cloud storage: privacy, powered by p2p collaborations and eco-friendly"`
	OwnerID     string          `json:"owner_id" example:"847390b4-a5b0-4ef7-949d-a15e84875d7e"`
	CreatedAt   time.Time       `json:"created_at" example:"2023-01-18T12:42:59.089247Z"`
	DeletedAt   *time.Time      `json:"deleted_at" example:"2023-01-18T12:42:59.089247Z"`
	Settings    *TenantSettings `json:"settings"`
	CouponID    *string         `json:"coupon_id" example:"9OPADNOEJFNO"`
}

type TenantSettings struct {
	ConsoleUrl        *string                                    `json:"console_url" example:"https://console.cubbit.eu"`
	GatewayUrl        *string                                    `json:"gateway_url" example:"https://s3.cubbit.eu"`
	Project           *TenantSettingsProject                     `json:"project"`
	Account           *TenantSettingsAccount                     `json:"account"`
	AllowedDomains    *[]string                                  `json:"allowed_domains"`
	BlockedDomains    *[]string                                  `json:"blocked_domains"`
	Notifications     *MonitoredResourceForSeverityNotifications `json:"notifications"`
	SignupDisabled    bool                                       `json:"signup_disabled"`
	SupportLink       *string                                    `json:"support_link"`
	DisplayName       *string                                    `json:"display_name"`
	WhitelabelEnabled bool                                       `json:"whitelabel_enabled"`
	WhiteLabel        *WhiteLabel                                `json:"white_label"`
}

type WhiteLabel struct {
	DNS         *WhiteLabelDNS         `json:"dns"`
	EmailDomain *WhiteLabelEmailDomain `json:"email_domain"`
}

type WhiteLabelDNS struct {
	Value     string `json:"value"`
	Challenge string `json:"challenge"`
	Verified  bool   `json:"verified"`
}

type WhiteLabelEmailDomain struct {
	Value     string               `json:"value"`
	User      *string              `json:"user"`
	Challenge []DNSChallengeRecord `json:"challenge"`
	Verified  bool                 `json:"verified"`
}

type DNSChallengeRecord struct {
	Type  string `json:"type" example:"TXT"`
	Name  string `json:"name" example:"_cubbit-challenge.tum-rating6.com"`
	Value string `json:"value" example:"some-value"`
}

type TenantSettingsProject struct {
	DefaultMaxProjectEgressBandwidth *int64 `json:"default_max_egress_bandwidth" example:"42"`
	DefaultMaxProjectStorage         *int64 `json:"default_max_storage" example:"42"`
}

type TenantSettingsAccount struct {
	DefaultMaxProject          *int                        `json:"default_max_project"`
	EnabledAuthProviders       *[]AccountAuthProvider      `json:"enabled_auth_providers"`
	EnabledAuthProvidersConfig *EnabledAuthProvidersConfig `json:"enabled_auth_providers_config"`
}

type AccountAuthProvider string
type EnabledAuthProvidersConfig struct {
	GoogleClientID         *string `json:"google_client_id"`
	MicrosoftApplicationID *string `json:"microsoft_application_id"`
	MicrosoftDirectoryID   *string `json:"microsoft_directory_id"`
}
type Severity string
type SeverityForNotificationType map[Severity]SeverityNotificationType
type MonitoredResourceForSeverityNotifications map[string]SeverityForNotificationType

type SeverityNotificationType struct {
	Threshold int          `json:"threshold" binding:"required,gte=1,lte=100" example:"33"`
	Notify    Notification `json:"notify" binding:"required"`
}

type Notification struct {
	NotifyOwner bool     `json:"notify_owner" binding:"required" example:"true"`
	Emails      []string `json:"emails" binding:"required,dive,email" example:"true"`
}
type TenantList struct {
	Tenants []*Tenant `json:"tenants"`
}

type Swarm struct {
	ID             string                 `json:"id" binding:"required,uuid"`
	Name           string                 `json:"name" binding:"required,min=1"`
	Configuration  map[string]interface{} `json:"configuration"`
	Size           int64                  `json:"size" binding:"required"`
	OwnerID        string                 `json:"owner_id"`
	Used           int64                  `json:"used"`
	CreationDate   time.Time              `json:"creation_date"`
	Description    *string                `json:"description"`
	OrganizationID *string                `json:"organization_id"`
}

type SwarmList struct {
	Swarms []Swarm `json:"swarms"`
}

type TenantSwarmList struct {
	Swarms []*TenantSwarm `json:"swarms"`
}
type TenantSwarm struct {
	ID                  string `json:"id"`
	SwarmID             string `json:"swarm_id"`
	TenantID            string `json:"tenant_id"`
	RedundancyClassID   string `json:"redundancy_class_id"`
	Default             bool   `json:"default"`
	SwarmName           string `json:"swarm_name"`
	RedundancyClassName string `json:"redundancy_class_name"`
}

type Operator struct {
	ID                 string          `json:"id" example:"695ed3dd-e77d-42b9-88ed-70bd3a1704ee"`
	FirstName          string          `json:"first_name" example:"Mario"`
	LastName           string          `json:"last_name" example:"Rossi"`
	Internal           bool            `json:"internal" example:"false"`
	Banned             bool            `json:"banned" example:"false"`
	CreatedAt          time.Time       `json:"created_at" example:"2023-01-18T12:42:59.089247Z"`
	DeletedAt          *time.Time      `json:"deleted_at" example:"2023-01-18T12:42:59.089247Z"`
	MaxAllowedProjects int             `json:"max_allowed_projects" example:"3"`
	Email              string          `json:"email"`
	Emails             []OperatorEmail `json:"emails"`
	TwoFactorEnabled   bool            `json:"two_factor_enabled" example:"true"`
	Status             string          `json:"status"`
	PolicyName         string          `json:"policy_name"`
}

type OperatorList struct {
	Operators []*Operator `json:"operators"`
}
type OperatorEmail struct {
	ID        string    `json:"id" example:"5ff281ee-75e7-4543-a304-ca861521f2a7"`
	Email     string    `json:"email" example:"mario.rossi@cubbit.io"`
	Verified  bool      `json:"verified" example:"true"`
	Default   bool      `json:"default" example:"true"`
	CreatedAt time.Time `json:"created_at" example:"2023-01-18T12:42:59.089247Z"`
}

type Policy struct {
	ID        string `json:"id"`
	AuthorID  string `json:"author_id"`
	Name      string `json:"name"`
	UserCount int    `json:"user_count"`
	Version   string `json:"version"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type PolicyList struct {
	Policies []Policy `json:"policies"`
}

type Distributor struct {
	ID          string    `json:"id" example:"784da023-cc6c-4a46-8e31-bc0a521d19e0"`
	Name        string    `json:"name" example:"Cubbit"`
	Description string    `json:"description" example:"Cloud storage: privacy, powered by p2p collaborations and eco-friendly"`
	OwnerID     string    `json:"owner_id" example:"847390b4-a5b0-4ef7-949d-a15e84875d7e"`
	CreatedAt   time.Time `json:"created_at" example:"2023-01-18T12:42:59.089247Z"`
	DeletedAt   time.Time `json:"deleted_at" example:"2023-01-18T12:42:59.089247Z"`
	ImageUrl    string    `json:"image_url" example:"https://s3.cubbit.io/my-new-test-bucket/Screenshot.png"`
}

type DistributorList struct {
	Distributors []*Distributor `json:"distributors"`
}

type DistributorCoupon struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	Redemptions    int       `json:"redemptions"`
	MaxRedemptions int64     `json:"max_redemptions"`
	Code           string    `json:"code"`
	Zone           string    `json:"zone"`
	CreatedAt      time.Time `json:"created_at"`
	ExternalID     string    `json:"external_id"`
}

type DistributorCouponList struct {
	Coupons []*DistributorCoupon `json:"coupons"`
}

type DistributorCouponCodeResponseModel struct {
	CouponCode string `json:"coupon_code"`
}

type TenantReport struct {
	ID         string        `json:"id"`
	ExternalID string        `json:"external_id"`
	TenantName string        `json:"tenant_name"`
	Code       string        `json:"code"`
	StorageBH  string        `json:"storage_bh"`
	StorageGB  string        `json:"storage_gb"`
	EngressGB  string        `json:"engress_gb"`
	FromTime   string        `json:"from_time"`
	ToTime     string        `json:"to_time"`
	Status     string        `json:"status"`
	Timestamp  time.Duration `json:"timestamp"`
}

type DistributorReportResponseModel struct {
	Report []TenantReport `json:"report"`
}

type ProjectReport struct {
	ID                  string        `json:"id"`
	ExternalID          string        `json:"external_id"`
	ProjectName         string        `json:"project_name"`
	StorageBH           string        `json:"storage_bh"`
	StorageAVGTB        string        `json:"storage_avg_tb"`
	StorageMaxTB        string        `json:"storage_max_tb"`
	StorageTB           string        `json:"storage_tb"`
	StorageReservedTB   string        `json:"storage_reserved_tb"`
	BandwidthReservedTB string        `json:"bandwidth_reserved_tb"`
	EgressTB            string        `json:"egress_tb"`
	IngressTB           string        `json:"ingress_tb"`
	FromTime            string        `json:"from_time"`
	ToTime              string        `json:"to_time"`
	Status              string        `json:"status"`
	Timestamp           time.Duration `json:"timestamp"`
}

type TenantReportResponseModel struct {
	Report []ProjectReport `json:"report"`
}

type ZoneResponse struct {
	Name        string `json:"name" binding:"required" example:"France"`
	Key         string `json:"key" binding:"required" example:"fr"`
	Description string `json:"description"`
}

type ZoneMap struct {
	Zones map[string]ZoneResponse `json:"zones"`
}
type UpdateAccountRequest struct {
	FirstName          *string `json:"first_name" example:"mario" binding:"omitempty,min=1,max=256"`
	LastName           *string `json:"last_name" example:"rossi" binding:"omitempty,min=1,max=256"`
	EndpointGateway    *string `json:"endpoint_gateway" binding:"omitempty,url,excludes= " example:"s3.cubbit.eu"`
	Internal           *bool   `json:"internal" example:"false"`
	MaxAllowedProjects *int    `json:"max_allowed_projects" example:"3"`
}

type Account struct {
	ID                 string              `json:"id" example:"695ed3dd-e77d-42b9-88ed-70bd3a1704ee"`
	FirstName          string              `json:"first_name" example:"Mario"`
	LastName           string              `json:"last_name" example:"Rossi"`
	Internal           bool                `json:"internal" example:"false"`
	Banned             bool                `json:"banned" example:"false"`
	CreatedAt          time.Time           `json:"created_at"`
	DeletedAt          *time.Time          `json:"deleted_at" example:"false"`
	MaxAllowedProjects int                 `json:"max_allowed_projects"  example:"3"`
	Emails             []AccountEmail      `json:"emails"`
	TwoFactorEnabled   bool                `json:"two_factor_enabled"`
	EndpointGateway    string              `json:"endpoint_gateway"`
	TenantID           string              `json:"tenant_id"`
	AuthProvider       AccountAuthProvider `json:"auth_provider" example:"cubbit"`
}

type AccountEmail struct {
	ID        string    `json:"id" example:"5ff281ee-75e7-4543-a304-ca861521f2a7"`
	Email     string    `json:"email" example:"mario.rossi@cubbit.io"`
	Verified  bool      `json:"verified" example:"true"`
	Default   bool      `json:"default" example:"true"`
	CreatedAt time.Time `json:"created_at"`
}

type GenericPaginatedResponse[T interface{}] struct {
	Data     []T  `json:"data"`
	NextPage *int `json:"next_page"`
	Count    int  `json:"count"`
}

type ProjectItem struct {
	ProjectID          string     `json:"project_id" example:"132d7f00-9e10-425e-afca-515cc4240d9f"`
	ProjectName        string     `json:"project_name" example:"Cubbit"`
	ProjectDescription string     `json:"project_description" example:"Cloud storage: best cloud service"`
	ProjectEmail       string     `json:"project_email" example:"321d7f00-9e10-425e-afca-515cc4240d9f@cubbit.io"`
	ProjectCreatedAt   time.Time  `json:"project_created_at" example:"2023-01-18T12:42:59.089247Z"`
	ProjectDeletedAt   *time.Time `json:"project_deleted_at" example:"2023-01-18T12:42:59.089247Z"`
	ProjectBannedAt    *time.Time `json:"project_banned_at" example:"2023-01-18T12:42:59.089247Z"`
	ProjectImageUrl    string     `json:"project_image_url" example:"https://s3.cubbit.io/my-new-test-bucket/Screenshot.png"`
	ProjectTenantID    string     `json:"project_tenant_id" example:"7add5517-9ddf-4037-a2e2-79e8bd664494"`
	RootAccountEmail   string     `json:"root_account_email" example:"piero@cubbit.io"`
	RootAccountID      string     `json:"root_account_id" example:"132d7f00-9e10-425e-afca-515cc4240d9f"`
}

type UpdateTenantProjectRequestBody struct {
	Name        *string `json:"name" example:"CubbitCloud"`
	Description *string `json:"description" example:"Cloud storage made easy"`
	ImageUrl    *string `json:"image_url" example:"https://s3.cubbit.io/my-new-test-bucket/Screenshot.png"`
}

type CreateNexusRequestBody struct {
	Name        string `json:"name" `
	Description string `json:"description,omitempty"`
	Location    string `json:"location"`
}

type UpdateNexusRequestBody struct {
	Name        string `json:"name,omitempty" `
	Description string `json:"description,omitempty"`
}

type Nexus struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Location     string    `json:"location"`
	CreationDate time.Time `json:"creation_date"`
	LastModified time.Time `json:"last_modified"`
	SwarmID      string    `json:"swarm_id"`
	Capacity     int       `json:"capacity"`
	Used         int       `json:"used"`
}

type NexusList struct {
	Nexuses  []*Nexus `json:"nexuses"`
	Page     int      `json:"page"`
	Count    int      `json:"count"`
	NextPage *int     `json:"next_page"`
}

type CreateNodeBodyRequest struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	NexusID     string `json:"nexus_id"`
	SecretID    string `json:"secret_id"`
}

type UpdateNodeBodyRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type ProviderList struct {
	Providers []Provider `json:"providers"`
	Page      int        `json:"page"`
	Count     int        `json:"count"`
}

type Provider struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	SwarmID string `json:"swarm_id"`
	Email   string `json:"email"`
}

type Node struct {
	ID           string                 `json:"id"`
	Name         string                 `json:"name"`
	Description  string                 `json:"description"`
	Status       string                 `json:"status"`
	CreationDate time.Time              `json:"creation_date"`
	ProviderID   string                 `json:"provider_id"`
	SecretID     string                 `json:"secret_id"`
	NexusID      string                 `json:"nexus_id"`
	Config       map[string]interface{} `json:"config"`
}

type NodeList struct {
	Nodes []*Node `json:"nodes"`
	Page  int     `json:"page"`
	Count int     `json:"count"`
}

type CreateRedundancyClassRequestBody struct {
	Name              string   `json:"name"`
	Description       string   `json:"description,omitempty"`
	InnerK            int      `json:"inner_k"`
	InnerN            int      `json:"inner_n"`
	OuterK            int      `json:"outer_k"`
	OuterN            int      `json:"outer_n"`
	AntiAffinityGroup int      `json:"anti_affinity_group"`
	Nexuses           []string `json:"nexuses"`
}

type RedundancyClass struct {
	ID                string    `json:"id"`
	Name              string    `json:"name"`
	Description       string    `json:"description"`
	InnerK            int       `json:"inner_k"`
	InnerN            int       `json:"inner_n"`
	OuterK            int       `json:"outer_k"`
	OuterN            int       `json:"outer_n"`
	AntiAffinityGroup int       `json:"anti_affinity_group"`
	Capacity          int       `json:"capacity"`
	CreationDate      time.Time `json:"creation_date"`
	SwarmID           string    `json:"swarm_id"`
}

type RedundancyClassList struct {
	Data []*RedundancyClass `json:"data"`
	Page int                `json:"page"`
}

type RingBulk struct {
	RedundancyClassID string   `json:"redundancy_class_id"`
	Nexuses           []string `json:"nexuses"`
	RingsNumber       *int     `json:"number_of_rings,omitempty"`
	AntiAffinityGroup *int     `json:"anti_affinity_group,omitempty"`
}

type Ring struct {
	ID           string      `json:"id"`
	N            int         `json:"n"`
	K            int         `json:"k"`
	SwarmID      string      `json:"swarm_id"`
	Capacity     int         `json:"capacity"`
	Used         int         `json:"used"`
	Status       string      `json:"status"`
	CreationDate string      `json:"creation_date"`
	Nexuses      []RingNexus `json:"nexuses,omitempty"`
	Nodes        []RingNode  `json:"nodes,omitempty"`
}

type RingNode struct {
	NodeID    string `json:"node_id"`
	NexusID   string `json:"nexus_id"`
	Sequence1 int    `json:"sequence1"`
	Sequence2 int    `json:"sequence2"`
}

type RingNexus struct {
	NexusID  string `json:"nexus_id"`
	N        int    `json:"n"`
	K        int    `json:"k"`
	Sequence int    `json:"sequence"`
}

type RingList struct {
	Data  []*Ring `json:"data"`
	Page  int     `json:"page"`
	Count int     `json:"count"`
}

type HumanReadableDistributorCoupon struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	Redemptions    int       `json:"redemptions"`
	MaxRedemptions string    `json:"max_redemptions"`
	Code           string    `json:"code"`
	Zone           string    `json:"zone"`
	CreatedAt      time.Time `json:"created_at"`
	ExternalID     string    `json:"external_id"`
}

func (c *DistributorCoupon) ToHumanReadableDistributorCode() *HumanReadableDistributorCoupon {
	maxRedemptions := fmt.Sprintf("%d", c.MaxRedemptions)
	if c.MaxRedemptions == -1 {
		maxRedemptions = "unlimited"
	}

	return &HumanReadableDistributorCoupon{
		ID:             c.ID,
		Name:           c.Name,
		Description:    c.Description,
		Redemptions:    c.Redemptions,
		MaxRedemptions: maxRedemptions,
		Code:           c.Code,
		Zone:           c.Zone,
		CreatedAt:      c.CreatedAt,
		ExternalID:     c.ExternalID,
	}
}

type BulkInsertNewNodeRequestBody struct {
	Nodes []CreateNewNodeRequestBody `json:"nodes" binding:"required,dive"`
}

type BulkInsertNewAgentRequestBody struct {
	Agents []CreateNewAgentRequestBody `json:"agents" binding:"required,dive"`
}

type CreateNewNodeRequestBody struct {
	Name          string                      `json:"name" binding:"required,min=3,max=63" example:"cubbit"`
	Label         *string                     `json:"label" example:"cubbit"`
	Configuration map[string]interface{}      `json:"config" swaggertype:"object,string" example:"{\"key\":\"value\"}"`
	PublicIP      string                      `json:"public_ip" binding:"required"`
	PrivateIP     string                      `json:"private_ip" binding:"required"`
	Agents        []CreateNewAgentRequestBody `json:"agents" binding:"dive"`
}

type CreateNewAgentRequestBody struct {
	Port     int                    `json:"port" binding:"required"`
	Features map[string]interface{} `json:"features" swaggertype:"object,string" example:"{\"key\":\"value\"}"`
	Volume   AgentVolume            `json:"volume" binding:"required,dive"`
}

type AgentVolume struct {
	MountPoint string `json:"mount_point" binding:"required,min=1" example:"/mnt/cubbit"`
	Disk       string `json:"disk" binding:"required,min=1" example:"sda"`
}

type NewNodesResponse struct {
	Nodes []NewNodeResponseItem `json:"nodes"`
}

type NewAgentsResponse struct {
	Agents []NewAgentResponse `json:"agents"`
}

type NewNodeResponseItem struct {
	ID            string                 `json:"id"`
	Name          string                 `json:"name" example:"cubbit"`
	Label         *string                `json:"label" example:"cubbit"`
	Configuration map[string]interface{} `json:"config" swaggertype:"object,string" example:"{\"key\":\"value\"}"`
	PublicIP      string                 `json:"public_ip"`
	PrivateIP     string                 `json:"private_ip"`
	Agents        []NewAgentResponse     `json:"agents"`
	CRN           map[string]interface{} `json:"crn"`
}

type NewAgentResponse struct {
	ID       string                 `json:"id"`
	Secret   string                 `json:"secret"`
	Port     int                    `json:"port"`
	Features map[string]interface{} `json:"features" swaggertype:"object,string" example:"{\"key\":\"value\"}"`
	Volume   AgentVolume            `json:"volume"`
	CRN      map[string]interface{} `json:"crn"`
}

type NewNode struct {
	ID            string                 `json:"id"`
	Name          string                 `json:"name"`
	Label         *string                `json:"label"`
	Configuration map[string]interface{} `json:"config"`
	CreatedAt     time.Time              `json:"created_at"`
	DeletedAt     *time.Time             `json:"deleted_at"`
	NexusID       string                 `json:"nexus_id"`
	PrivateIP     string                 `json:"private_ip"`
	PublicIP      string                 `json:"public_ip"`
	CRN           map[string]interface{} `json:"crn"`
}

type UpdateNewNodeRequestBody struct {
	Name          *string                `json:"name" binding:"omitempty,min=3,max=63" example:"cubbit"`
	Label         *string                `json:"label" example:"cube"`
	Configuration map[string]interface{} `json:"config" binding:"omitempty"`
	PublicIP      *string                `json:"public_ip" binding:"omitempty"`
	PrivateIP     *string                `json:"private_ip" binding:"omitempty"`
}

type NewAgent struct {
	ID                 string                 `json:"id"`
	NodeID             string                 `json:"node_id"`
	PublicKey          *string                `json:"public_key"`
	Features           map[string]interface{} `json:"features"`
	CreatedAt          time.Time              `json:"created"`
	DeletedAt          *time.Time             `json:"deleted"`
	LastStatusUpdateAt time.Time              `json:"last_status_update_at"`
	Secret             string                 `json:"secret"`
	Online             bool                   `json:"online"`
	ConnectedOn        *time.Time             `json:"connected_on,omitempty"`
	Model              AgentModel             `json:"model"`
	Serial             *string                `json:"serial,omitempty"`
	Version            *string                `json:"version"`
	Port               int                    `json:"cccp_port,omitempty"`
	TotalStorage       int64                  `json:"total_storage"`
	UsedStorage        int64                  `json:"used_storage"`
	Volume             AgentVolume            `json:"volume"`
	CRN                map[string]interface{} `json:"crn"`
	AnnouncedPrivateIP *string                `json:"announced_private_ip"`
	AnnouncedPublicIP  *string                `json:"announced_public_ip"`
}

type AgentModel string

const (
	UnknownAgentModel AgentModel = "unknown"
	CellAgentModel    AgentModel = "cell"
	VirtualAgentModel AgentModel = "virtual"
)

type UpdateNewAgentRequestBody struct {
	Port     *int                      `json:"port"`
	Features map[string]interface{}    `json:"features" swaggertype:"object,string" example:"{\"key\":\"value\"}"`
	Volume   *UpdateAgentVolumeRequest `json:"volume" binding:"omitempty,dive"`
}

type UpdateAgentVolumeRequest struct {
	MountPoint *string `json:"mount_point" binding:"omitempty,min=1" example:"/mnt/cubbit"`
	Disk       *string `json:"disk" binding:"omitempty,min=1" example:"sda"`
}

type ExpansionStatus string
type RecoveryStatus string

const (
	ExpansionStatusSuccess  ExpansionStatus = "success"
	ExpansionStatusPossible ExpansionStatus = "possible"
	ExpansionStatusError    ExpansionStatus = "error"

	RecoveryStatusCompleted RecoveryStatus = "completed"
	RecoveryStatusOngoing   RecoveryStatus = "ongoing"
	RecoveryStatusErrored   RecoveryStatus = "errored"
)

type RedundancyClassExpanded struct {
	Status         ExpansionStatus `json:"status"`
	Message        string          `json:"message"`
	ExpandedSize   int64           `json:"expanded_size"`
	AgentsInvolved []string        `json:"agents_involved"`
}

type RedundancyClassRecoveryStatus struct {
	Status      RecoveryStatus      `json:"status"`
	Message     string              `json:"message"`
	Errored     int                 `json:"errored"`
	Total       int                 `json:"total"`
	RecoveryIds []string            `json:"recovery_ids"`
	IssueFound  []map[string]string `json:"issue_found"`
}

type RedundancyClassErrorData struct {
	IssueFound []interface{} `json:"issue_found"`
}

type RedundancyClassRecovery struct {
	Message string `json:"message"`
}

type SummaryDetailsWithStatusNullable struct {
	Details *SummaryDetails `json:"details"`
	SummaryStatusNullable
}

type EvaluatedStatusType string

type SummaryStatusNullable struct {
	EvaluatedStatus              *EvaluatedStatusType `json:"evaluated_status"`
	EvaluatedStatusLastUpdatedAt *time.Time           `json:"evaluated_status_last_updated_at"`
}

type SummaryDetails struct {
	Pending int `json:"pending"`
	Online  int `json:"online"`
	Offline int `json:"offline"`
	Warning int `json:"warning"`
	Error   int `json:"error"`
}

type NodeConfig struct {
	ID        string        `json:"id,omitempty"`
	Name      string        `json:"name"`
	PublicIP  string        `json:"public_ip"`
	PrivateIP string        `json:"private_ip"`
	Label     string        `json:"label,omitempty"`
	Agents    []AgentConfig `json:"agents"`
}

type AgentConfig struct {
	ID         string `json:"id"`
	MountPoint string `json:"mount_point"`
	Disk       string `json:"disk"`
	Port       int    `json:"port"`
	Secret     string `json:"secret"`
}

type AnsibleConfig struct {
	Nodes []NodeConfig `json:"nodes"`
}

type YAMLGenerationEnvs struct {
	HiveURL                  string `yaml:"hive_url"`
	MetricsURL               string `yaml:"metrics_url"`
	MetricsRoutesSend        string `yaml:"metrics_routes_send"`
	CCCPSwarmGatewayEndpoint string `yaml:"cccp_swarm_gateway_endpoint"`
	CCCPSwarmGatewayPort     string `yaml:"cccp_swarm_gateway_port"`
	CCCPSwarmGatewaySecure   string `yaml:"cccp_swarm_gateway_secure"`
}

type SecretYAML struct {
	APIVersion string            `yaml:"apiVersion"`
	Kind       string            `yaml:"kind"`
	Metadata   SecretMetadata    `yaml:"metadata"`
	Type       string            `yaml:"type"`
	Data       map[string]string `yaml:"data"`
}

type SecretMetadata struct {
	Name string `yaml:"name"`
}

type ClusterAgentYAML struct {
	APIVersion string               `yaml:"apiVersion"`
	Kind       string               `yaml:"kind"`
	Metadata   ClusterAgentMetadata `yaml:"metadata"`
	Spec       ClusterAgentSpec     `yaml:"spec"`
}

type ClusterAgentMetadata struct {
	Name string `yaml:"name"`
}

type ClusterAgentSpec struct {
	InstancesCounter  int                    `yaml:"instancesCounter"`
	BaseName          string                 `yaml:"baseName"`
	SecretName        string                 `yaml:"secretName"`
	AgentImage        string                 `yaml:"agentImage"`
	AdditionalEnvVars []EnvVar               `yaml:"additionalEnvVars"`
	Volume            VolumeSpec             `yaml:"volume"`
	AgentsDetail      map[string]AgentDetail `yaml:"agentsDetail"`
}

type EnvVar struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value,omitempty"`
}

type VolumeSpec struct {
	Type    string `yaml:"type"`
	PVCSize string `yaml:"pvcSize"`
}

type AgentDetail struct {
	LocalPath        string `yaml:"localPath"`
	NodeNameSelector string `yaml:"nodeNameSelector"`
}

type AgentSecret struct {
	AgentSecret string `json:"agentSecret"`
	AgentUUID   string `json:"agentUUID"`
}

type CreateGatewayRequestBody struct {
	Name          string                 `json:"name" binding:"required,min=3,max=63" example:"gateway"`
	Location      string                 `json:"location" binding:"required" example:"eu-west-1"`
	Configuration map[string]interface{} `json:"configuration"`
}

type UpdateGatewayRequestBody struct {
	Name     *string `json:"name" binding:"omitempty,min=3,max=63" example:"cubbit"`
	Location *string `json:"location" binding:"omitempty,min=1" example:"eu-west-1"`
}
type Gateway struct {
	ID             string     `json:"id" binding:"required,uuid"`
	Name           string     `json:"name" binding:"required"`
	Location       string     `json:"location" binding:"required"`
	CreatedAt      time.Time  `json:"created_at" example:"2023-01-18T12:42:59.089247Z"`
	DeletedAt      *time.Time `json:"deleted_at" example:"2023-01-18T12:42:59.089247Z"`
	HardDeleteAt   *time.Time `json:"hard_delete_at" example:"2023-01-18T12:42:59.089247Z"`
	Secret         string     `json:"secret" binding:"required"`
	OrganizationID *string    `json:"organization_id" example:"00000000-0000-0000-0000-000000000000"`
}

type GatewayWithGatewayTenant struct {
	Gateway       *Gateway       `json:"gateway"`
	GatewayTenant *GatewayTenant `json:"gateway_tenant"`
}

type GatewayTenant struct {
	ID            string                 `json:"id" binding:"required,uuid"`
	GatewayID     string                 `json:"gateway_id" binding:"required,uuid"`
	TenantID      string                 `json:"tenant_id" binding:"required,uuid"`
	CreatedAt     time.Time              `json:"created_at" example:"2023-01-18T12:42:59.089247Z"`
	DeletedAt     *time.Time             `json:"deleted_at" example:"2023-01-18T12:42:59.089247Z"`
	HardDeleteAt  *time.Time             `json:"hard_delete_at" example:"2023-01-18T12:42:59.089247Z"`
	Configuration map[string]interface{} `json:"configuration"`
	Hidden        bool                   `json:"hidden"`
}

type GatewayInstanceListResponse struct {
	Data []*GatewayInstance `json:"data"`
}

type GatewayInstance struct {
	ID              string                 `json:"id" binding:"required,uuid"`
	GatewayID       string                 `json:"gateway_id" binding:"required,uuid"`
	IP              string                 `json:"ip" binding:"required"`
	Status          map[string]interface{} `json:"status"`
	StatusUpdatedAt *time.Time             `json:"status_updated_at" example:"2023-01-18T12:42:59.089247Z"`
	CreatedAt       time.Time              `json:"created_at" example:"2023-01-18T12:42:59.089247Z"`
	Features        map[string]interface{} `json:"features"`
	EvaluatedStatus GatewayStatus          `json:"evaluated_status"`
}

type GatewayStatus string

const (
	UnknownGatewayStatus  GatewayStatus = "unknown"
	InactiveGatewayStatus GatewayStatus = "inactive"
	ActiveGatewayStatus   GatewayStatus = "active"
	OfflineGatewayStatus  GatewayStatus = "offline"
	OnlineGatewayStatus   GatewayStatus = "online"
)
