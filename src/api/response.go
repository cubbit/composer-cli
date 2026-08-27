package api

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	UnknownGatewayStatus  GatewayStatus = "unknown"
	InactiveGatewayStatus GatewayStatus = "inactive"
	ActiveGatewayStatus   GatewayStatus = "active"
	OfflineGatewayStatus  GatewayStatus = "offline"
	OnlineGatewayStatus   GatewayStatus = "online"
)

const (
	UnknownAgentModel AgentModel = "unknown"
	CellAgentModel    AgentModel = "cell"
	VirtualAgentModel AgentModel = "virtual"
)

const (
	ExpansionStatusSuccess  ExpansionStatus = "success"
	ExpansionStatusPossible ExpansionStatus = "possible"
	ExpansionStatusError    ExpansionStatus = "error"

	RecoveryStatusCompleted RecoveryStatus = "completed"
	RecoveryStatusOngoing   RecoveryStatus = "ongoing"
	RecoveryStatusErrored   RecoveryStatus = "errored"
)

const (
	RCSessionStatusRunning   RCSessionStatus = "running"
	RCSessionStatusCompleted RCSessionStatus = "completed"
	RCSessionStatusFailed    RCSessionStatus = "failed"
)

type GatewayStatus string
type AgentModel string
type ExpansionStatus string
type RecoveryStatus string
type RCSessionStatus string

type ErrorResponseModel struct {
	Message string `json:"message" yaml:"message"`
	Param   string `json:"param" yaml:"param"`
}

type ChallengeResponseModel struct {
	Challenge string `json:"challenge" yaml:"challenge"`
	Salt      string `json:"salt" yaml:"salt" example:"SGVsbG8gd29ybGQ="`
}

type OperatorAPIKey struct {
	ID         string     `json:"id" yaml:"id" example:"695ed3dd-e77d-42b9-88ed-70bd3a1704ee"`
	Name       string     `json:"name" yaml:"name" example:"API Key for Mario Rossi"`
	Key        string     `json:"key" yaml:"key" example:"4f2a1b3c5d6e7f8a9b0c1d2e3f4a5b6c"`
	OperatorID string     `json:"operator_id" yaml:"operator_id" example:"695ed3dd-e77d-42b9-88ed-70bd3a1704ee"`
	CreatedAt  time.Time  `json:"created_at" yaml:"created_at" example:"2023-01-18T12:42:59.089247Z"`
	ExpiresAt  *time.Time `json:"expires_at" yaml:"expires_at" example:"2023-01-18T12:42:59.089247Z"`
	DeletedAt  *time.Time `json:"deleted_at" yaml:"deleted_at" example:"2023-01-18T12:42:59.089247Z"`
	BannedAt   *time.Time `json:"banned_at" yaml:"banned_at" example:"2023-01-18T12:42:59.089247Z"`
	Enabled    bool       `json:"enabled" yaml:"enabled" example:"true"`
}

type TokenAndExpirationResponseModel struct {
	Token   string    `json:"token" yaml:"token"`
	Exp     int       `json:"exp" yaml:"exp"`
	ExpDate time.Time `json:"expDate" yaml:"expDate"`
}

type GenericIDResponseModel struct {
	ID string `json:"id" yaml:"id"`
}

type Tenant struct {
	ID          string          `json:"id" yaml:"id"`
	Name        string          `json:"name" yaml:"name"`
	Description *string         `json:"description" yaml:"description"`
	OwnerID     string          `json:"owner_id" yaml:"owner_id"`
	CreatedAt   time.Time       `json:"created_at" yaml:"created_at"`
	DeletedAt   *time.Time      `json:"deleted_at" yaml:"deleted_at"`
	Settings    *TenantSettings `json:"settings" yaml:"settings"`
	CouponID    *string         `json:"coupon_id" yaml:"coupon_id"`
}

type TenantSettings struct {
	ConsoleUrl        *string                                    `json:"console_url" yaml:"console_url"`
	GatewayUrl        *string                                    `json:"gateway_url" yaml:"gateway_url"`
	Project           *TenantSettingsProject                     `json:"project" yaml:"project"`
	Account           *TenantSettingsAccount                     `json:"account" yaml:"account"`
	AllowedDomains    *[]string                                  `json:"allowed_domains" yaml:"allowed_domains"`
	BlockedDomains    *[]string                                  `json:"blocked_domains" yaml:"blocked_domains"`
	Notifications     *MonitoredResourceForSeverityNotifications `json:"notifications" yaml:"notifications"`
	SignupDisabled    *bool                                      `json:"signup_disabled" yaml:"signup_disabled"`
	SupportLink       *string                                    `json:"support_link" yaml:"support_link"`
	DisplayName       *string                                    `json:"display_name" yaml:"display_name"`
	WhitelabelEnabled *bool                                      `json:"whitelabel_enabled" yaml:"whitelabel_enabled"`
	WhitelabelAssets  *WhitelabelAssets                          `json:"whitelabel_assets" yaml:"whitelabel_assets"`
	WhiteLabel        *WhiteLabel                                `json:"white_label" yaml:"white_label"`
}

type WhiteLabel struct {
	DNS         *WhiteLabelDNS         `json:"dns" yaml:"dns"`
	EmailDomain *WhiteLabelEmailDomain `json:"email_domain" yaml:"email_domain"`
}

type WhiteLabelDNS struct {
	Value     string `json:"value" yaml:"value"`
	Challenge string `json:"challenge" yaml:"challenge"`
	Verified  bool   `json:"verified" yaml:"verified"`
}

type WhiteLabelEmailDomain struct {
	Value     string               `json:"value" yaml:"value"`
	User      *string              `json:"user" yaml:"user"`
	Challenge []DNSChallengeRecord `json:"challenge" yaml:"challenge"`
	Verified  bool                 `json:"verified" yaml:"verified"`
}

type DNSChallengeRecord struct {
	Type  string `json:"type" yaml:"type"`
	Name  string `json:"name" yaml:"name"`
	Value string `json:"value" yaml:"value"`
}

type TenantSettingsProject struct {
	DefaultMaxProjectEgressBandwidthGB *int64 `json:"default_max_egress_bandwidth" yaml:"default_max_egress_bandwidth"`
	DefaultMaxProjectStorageGB         *int64 `json:"default_max_storage" yaml:"default_max_storage"`
}

type WhitelabelAssets struct {
	Icons      *IconAssets  `json:"icon_assets" yaml:"icon_assets"`
	LocalFiles *LocalAssets `json:"local_assets" yaml:"local_assets"`
	LegalFiles *LegalAssets `json:"legal_assets" yaml:"legal_assets"`
}

type IconAssets struct {
	Favicon      *string `json:"favicon" yaml:"favicon"`
	Logo         *string `json:"logo" yaml:"logo"`
	ExtendedLogo *string `json:"extended_logo" yaml:"extended_logo"`
	ComposerLogo *string `json:"composer_logo" yaml:"composer_logo"`
}

type LocalAssets struct {
	SignUp *string `json:"sign_up" yaml:"sign_up"`
	SignIn *string `json:"sign_in" yaml:"sign_in"`
}

type LegalAssets struct {
	TermsOfService *string `json:"terms_of_service" yaml:"terms_of_service"`
	DataPrivacy    *string `json:"data_privacy" yaml:"data_privacy"`
}

type TenantSettingsAccount struct {
	DefaultMaxProject          *int                        `json:"default_max_project" yaml:"default_max_project"`
	EnabledAuthProviders       *[]AccountAuthProvider      `json:"enabled_auth_providers" yaml:"enabled_auth_providers"`
	EnabledAuthProvidersConfig *EnabledAuthProvidersConfig `json:"enabled_auth_providers_config" yaml:"enabled_auth_providers_config"`
	OAuthProviders             *[]OAuthProvider            `json:"oauth_providers" yaml:"oauth_providers"`
}

type OAuthProvider struct {
	Name         string `json:"name" yaml:"name"`
	ClientID     string `json:"client_id" yaml:"client_id"`
	ClientSecret string `json:"client_secret" yaml:"client_secret"`
	IssuerURL    string `json:"issuer_url" yaml:"issuer_url"`
	RedirectURL  string `json:"redirect_url" yaml:"redirect_url"`
}

type AccountAuthProvider string

type EnabledAuthProvidersConfig struct {
	GoogleClientID         *string `json:"google_client_id" yaml:"google_client_id"`
	MicrosoftApplicationID *string `json:"microsoft_application_id" yaml:"microsoft_application_id"`
	MicrosoftDirectoryID   *string `json:"microsoft_directory_id" yaml:"microsoft_directory_id"`
}

type Severity string

type SeverityForNotificationType map[Severity]SeverityNotificationType

type MonitoredResourceForSeverityNotifications map[string]SeverityForNotificationType

type SeverityNotificationType struct {
	Threshold int          `json:"threshold" yaml:"threshold"`
	Notify    Notification `json:"notify" yaml:"notify"`
}

type Notification struct {
	NotifyOwner bool     `json:"notify_owner" yaml:"notify_owner"`
	Emails      []string `json:"emails" yaml:"emails"`
}

type TenantList struct {
	Tenants []*Tenant `json:"tenants" yaml:"tenants"`
}

type Swarm struct {
	ID             string                 `json:"id" yaml:"id"`
	Name           string                 `json:"name" yaml:"name"`
	Configuration  map[string]interface{} `json:"configuration" yaml:"configuration"`
	Size           int64                  `json:"size" yaml:"size"`
	OwnerID        string                 `json:"owner_id" yaml:"owner_id"`
	Used           int64                  `json:"used" yaml:"used"`
	CreationDate   time.Time              `json:"creation_date" yaml:"creation_date"`
	Description    *string                `json:"description" yaml:"description"`
	OrganizationID *string                `json:"organization_id" yaml:"organization_id"`
}

type SwarmList struct {
	Swarms []Swarm `json:"swarms" yaml:"swarms"`
}

type TenantSwarmList struct {
	Swarms []*TenantSwarm `json:"swarms" yaml:"swarms"`
}

type TenantSwarm struct {
	ID                  string `json:"id" yaml:"id"`
	SwarmID             string `json:"swarm_id" yaml:"swarm_id"`
	TenantID            string `json:"tenant_id" yaml:"tenant_id"`
	RedundancyClassID   string `json:"redundancy_class_id" yaml:"redundancy_class_id"`
	Default             bool   `json:"default" yaml:"default"`
	SwarmName           string `json:"swarm_name" yaml:"swarm_name"`
	RedundancyClassName string `json:"redundancy_class_name" yaml:"redundancy_class_name"`
}

type IAMUserPolicy struct {
	ID   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
}

type IAMUser struct {
	ID                 string          `json:"id" yaml:"id"`
	Username           string          `json:"username" yaml:"username"`
	FirstName          string          `json:"first_name" yaml:"first_name"`
	LastName           string          `json:"last_name" yaml:"last_name"`
	Email              string          `json:"email" yaml:"email"`
	Emails             []IAMUserEmail  `json:"emails" yaml:"emails"`
	Enabled            bool            `json:"enabled" yaml:"enabled"`
	Internal           bool            `json:"internal" yaml:"internal"`
	Banned             bool            `json:"banned" yaml:"banned"`
	IsRoot             bool            `json:"is_root" yaml:"is_root"`
	TwoFactorEnabled   bool            `json:"two_factor_enabled" yaml:"two_factor_enabled"`
	Status             string          `json:"status" yaml:"status"`
	CreatedAt          time.Time       `json:"created_at" yaml:"created_at"`
	DeletedAt          *time.Time      `json:"deleted_at" yaml:"deleted_at"`
	LastActivityAt     *time.Time      `json:"last_activity_at" yaml:"last_activity_at"`
	MaxAllowedProjects int             `json:"max_allowed_projects" yaml:"max_allowed_projects"`
	PoliciesCount      int             `json:"policies_count" yaml:"policies_count"`
	Policies           []IAMUserPolicy `json:"policies" yaml:"policies"`
	OrganizationID     *string         `json:"organization_id" yaml:"organization_id"`
	OrganizationName   *string         `json:"organization_name" yaml:"organization_name"`
}

type IAMUserList struct {
	Operators []*IAMUser `json:"operators" yaml:"operators"`
}

type IAMUserEmail struct {
	ID             string    `json:"id" yaml:"id"`
	Email          string    `json:"email" yaml:"email"`
	Verified       bool      `json:"verified" yaml:"verified"`
	Default        bool      `json:"default" yaml:"default"`
	CreatedAt      time.Time `json:"created_at" yaml:"created_at"`
	OrganizationID *string   `json:"organization_id" yaml:"organization_id"`
}

type Policy struct {
	ID        string `json:"id" yaml:"id"`
	AuthorID  string `json:"author_id" yaml:"author_id"`
	Name      string `json:"name" yaml:"name"`
	UserCount int    `json:"user_count" yaml:"user_count"`
	Version   string `json:"version" yaml:"version"`
	CreatedAt string `json:"created_at" yaml:"created_at"`
	UpdatedAt string `json:"updated_at" yaml:"updated_at"`
}

type PolicyList struct {
	Policies []Policy `json:"policies" yaml:"policies"`
}

type Distributor struct {
	ID          string    `json:"id" yaml:"id"`
	Name        string    `json:"name" yaml:"name"`
	Description string    `json:"description" yaml:"description"`
	OwnerID     string    `json:"owner_id" yaml:"owner_id"`
	CreatedAt   time.Time `json:"created_at" yaml:"created_at"`
	DeletedAt   time.Time `json:"deleted_at" yaml:"deleted_at"`
	ImageURL    string    `json:"image_url" yaml:"image_url"`
}

type DistributorList struct {
	Distributors []*Distributor `json:"distributors" yaml:"distributors"`
}

type DistributorCoupon struct {
	ID             string    `json:"id" yaml:"id"`
	Name           string    `json:"name" yaml:"name"`
	Description    string    `json:"description" yaml:"description"`
	Redemptions    int       `json:"redemptions" yaml:"redemptions"`
	MaxRedemptions int64     `json:"max_redemptions" yaml:"max_redemptions"`
	Code           string    `json:"code" yaml:"code"`
	Zone           string    `json:"zone" yaml:"zone"`
	CreatedAt      time.Time `json:"created_at" yaml:"created_at"`
	ExternalID     string    `json:"external_id" yaml:"external_id"`
}

type DistributorCouponList struct {
	Coupons []*DistributorCoupon `json:"coupons" yaml:"coupons"`
}

type DistributorCouponCodeResponseModel struct {
	CouponCode string `json:"coupon_code" yaml:"coupon_code"`
}

type TenantReport struct {
	ID         string        `json:"id" yaml:"id"`
	ExternalID string        `json:"external_id" yaml:"external_id"`
	TenantName string        `json:"tenant_name" yaml:"tenant_name"`
	Code       string        `json:"code" yaml:"code"`
	StorageBH  string        `json:"storage_bh" yaml:"storage_bh"`
	StorageGB  string        `json:"storage_gb" yaml:"storage_gb"`
	EngressGB  string        `json:"engress_gb" yaml:"engress_gb"`
	FromTime   string        `json:"from_time" yaml:"from_time"`
	ToTime     string        `json:"to_time" yaml:"to_time"`
	Status     string        `json:"status" yaml:"status"`
	Timestamp  time.Duration `json:"timestamp" yaml:"timestamp"`
}

type DistributorReportResponseModel struct {
	Report []TenantReport `json:"report" yaml:"report"`
}

type ProjectReport struct {
	ID                  string        `json:"id" yaml:"id"`
	ExternalID          string        `json:"external_id" yaml:"external_id"`
	ProjectName         string        `json:"project_name" yaml:"project_name"`
	StorageBH           string        `json:"storage_bh" yaml:"storage_bh"`
	StorageAVGTB        string        `json:"storage_avg_tb" yaml:"storage_avg_tb"`
	StorageMaxTB        string        `json:"storage_max_tb" yaml:"storage_max_tb"`
	StorageTB           string        `json:"storage_tb" yaml:"storage_tb"`
	StorageReservedTB   string        `json:"storage_reserved_tb" yaml:"storage_reserved_tb"`
	BandwidthReservedTB string        `json:"bandwidth_reserved_tb" yaml:"bandwidth_reserved_tb"`
	EgressTB            string        `json:"egress_tb" yaml:"egress_tb"`
	IngressTB           string        `json:"ingress_tb" yaml:"ingress_tb"`
	FromTime            string        `json:"from_time" yaml:"from_time"`
	ToTime              string        `json:"to_time" yaml:"to_time"`
	Status              string        `json:"status" yaml:"status"`
	Timestamp           time.Duration `json:"timestamp" yaml:"timestamp"`
}

type TenantReportResponseModel struct {
	Report []ProjectReport `json:"report" yaml:"report"`
}

type ZoneResponse struct {
	Name        string `json:"name" yaml:"name"`
	Key         string `json:"key" yaml:"key"`
	Description string `json:"description" yaml:"description"`
}

type ZoneMap struct {
	Zones map[string]ZoneResponse `json:"zones" yaml:"zones"`
}

type Account struct {
	ID                 string              `json:"id" yaml:"id"`
	FirstName          string              `json:"first_name" yaml:"first_name"`
	LastName           string              `json:"last_name" yaml:"last_name"`
	Internal           bool                `json:"internal" yaml:"internal"`
	Banned             bool                `json:"banned" yaml:"banned"`
	CreatedAt          time.Time           `json:"created_at" yaml:"created_at"`
	DeletedAt          *time.Time          `json:"deleted_at" yaml:"deleted_at"`
	MaxAllowedProjects int                 `json:"max_allowed_projects" yaml:"max_allowed_projects"`
	Emails             []AccountEmail      `json:"emails" yaml:"emails"`
	TwoFactorEnabled   bool                `json:"two_factor_enabled" yaml:"two_factor_enabled"`
	EndpointGateway    string              `json:"endpoint_gateway" yaml:"endpoint_gateway"`
	TenantID           string              `json:"tenant_id" yaml:"tenant_id"`
	AuthProvider       AccountAuthProvider `json:"auth_provider" yaml:"auth_provider"`
}

type AccountEmail struct {
	ID        string    `json:"id" yaml:"id"`
	Email     string    `json:"email" yaml:"email"`
	Verified  bool      `json:"verified" yaml:"verified"`
	Default   bool      `json:"default" yaml:"default"`
	CreatedAt time.Time `json:"created_at" yaml:"created_at"`
}

type GenericPaginatedResponse[T interface{}] struct {
	Data     []T  `json:"data" yaml:"data"`
	NextPage *int `json:"next_page" yaml:"next_page"`
	Count    int  `json:"count" yaml:"count"`
}

type ProjectItem struct {
	ProjectID          string     `json:"project_id" yaml:"project_id"`
	ProjectName        string     `json:"project_name" yaml:"project_name"`
	ProjectDescription string     `json:"project_description" yaml:"project_description"`
	ProjectEmail       string     `json:"project_email" yaml:"project_email"`
	ProjectCreatedAt   time.Time  `json:"project_created_at" yaml:"project_created_at"`
	ProjectDeletedAt   *time.Time `json:"project_deleted_at" yaml:"project_deleted_at"`
	ProjectBannedAt    *time.Time `json:"project_banned_at" yaml:"project_banned_at"`
	ProjectImageURL    string     `json:"project_image_url" yaml:"project_image_url"`
	ProjectTenantID    string     `json:"project_tenant_id" yaml:"project_tenant_id"`
	RootAccountEmail   string     `json:"root_account_email" yaml:"root_account_email"`
	RootAccountID      string     `json:"root_account_id" yaml:"root_account_id"`
}

type Nexus struct {
	ID           string    `json:"id" yaml:"id"`
	Name         string    `json:"name" yaml:"name"`
	Description  string    `json:"description" yaml:"description"`
	Location     string    `json:"location" yaml:"location"`
	CreationDate time.Time `json:"creation_date" yaml:"creation_date"`
	LastModified time.Time `json:"last_modified" yaml:"last_modified"`
	SwarmID      string    `json:"swarm_id" yaml:"swarm_id"`
	Capacity     int       `json:"capacity" yaml:"capacity"`
	Used         int       `json:"used" yaml:"used"`
}

type NexusList struct {
	Nexuses  []*Nexus `json:"nexuses" yaml:"nexuses"`
	Page     int      `json:"page" yaml:"page"`
	Count    int      `json:"count" yaml:"count"`
	NextPage *int     `json:"next_page" yaml:"next_page"`
}

type ProviderList struct {
	Providers []Provider `json:"providers" yaml:"providers"`
	Page      int        `json:"page" yaml:"page"`
	Count     int        `json:"count" yaml:"count"`
}

type Provider struct {
	ID      string `json:"id" yaml:"id"`
	Name    string `json:"name" yaml:"name"`
	SwarmID string `json:"swarm_id" yaml:"swarm_id"`
	Email   string `json:"email" yaml:"email"`
}

type Node struct {
	ID           string                 `json:"id" yaml:"id"`
	Name         string                 `json:"name" yaml:"name"`
	Description  string                 `json:"description" yaml:"description"`
	Status       string                 `json:"status" yaml:"status"`
	CreationDate time.Time              `json:"creation_date" yaml:"creation_date"`
	ProviderID   string                 `json:"provider_id" yaml:"provider_id"`
	SecretID     string                 `json:"secret_id" yaml:"secret_id"`
	NexusID      string                 `json:"nexus_id" yaml:"nexus_id"`
	Config       map[string]interface{} `json:"config" yaml:"config"`
}

type NodeList struct {
	Nodes []*Node `json:"nodes" yaml:"nodes"`
	Page  int     `json:"page" yaml:"page"`
	Count int     `json:"count" yaml:"count"`
}

type RedundancyClass struct {
	ID                string    `json:"id" yaml:"id"`
	Name              string    `json:"name" yaml:"name"`
	Description       string    `json:"description" yaml:"description"`
	InnerK            int       `json:"inner_k" yaml:"inner_k"`
	InnerN            int       `json:"inner_n" yaml:"inner_n"`
	OuterK            int       `json:"outer_k" yaml:"outer_k"`
	OuterN            int       `json:"outer_n" yaml:"outer_n"`
	AntiAffinityGroup int       `json:"anti_affinity_group" yaml:"anti_affinity_group"`
	Capacity          int       `json:"capacity" yaml:"capacity"`
	CreationDate      time.Time `json:"creation_date" yaml:"creation_date"`
	SwarmID           string    `json:"swarm_id" yaml:"swarm_id"`
}

type RedundancyClassList struct {
	Data []*RedundancyClass `json:"data" yaml:"data"`
	Page int                `json:"page" yaml:"page"`
}

type RingBulk struct {
	RedundancyClassID string   `json:"redundancy_class_id" yaml:"redundancy_class_id"`
	Nexuses           []string `json:"nexuses" yaml:"nexuses"`
	RingsNumber       *int     `json:"number_of_rings" yaml:"number_of_rings"`
	AntiAffinityGroup *int     `json:"anti_affinity_group" yaml:"anti_affinity_group"`
}

type Ring struct {
	ID           string      `json:"id" yaml:"id"`
	N            int         `json:"n" yaml:"n"`
	K            int         `json:"k" yaml:"k"`
	SwarmID      string      `json:"swarm_id" yaml:"swarm_id"`
	Capacity     int         `json:"capacity" yaml:"capacity"`
	Used         int         `json:"used" yaml:"used"`
	Status       string      `json:"status" yaml:"status"`
	CreationDate string      `json:"creation_date" yaml:"creation_date"`
	Nexuses      []RingNexus `json:"nexuses" yaml:"nexuses"`
	Nodes        []RingNode  `json:"nodes" yaml:"nodes"`
}

type RingNode struct {
	NodeID    string `json:"node_id" yaml:"node_id"`
	NexusID   string `json:"nexus_id" yaml:"nexus_id"`
	Sequence1 int    `json:"sequence1" yaml:"sequence1"`
	Sequence2 int    `json:"sequence2" yaml:"sequence2"`
}

type RingNexus struct {
	NexusID  string `json:"nexus_id" yaml:"nexus_id"`
	N        int    `json:"n" yaml:"n"`
	K        int    `json:"k" yaml:"k"`
	Sequence int    `json:"sequence" yaml:"sequence"`
}

type RingList struct {
	Data  []*Ring `json:"data" yaml:"data"`
	Page  int     `json:"page" yaml:"page"`
	Count int     `json:"count" yaml:"count"`
}

type HumanReadableDistributorCoupon struct {
	ID             string    `json:"id" yaml:"id"`
	Name           string    `json:"name" yaml:"name"`
	Description    string    `json:"description" yaml:"description"`
	Redemptions    int       `json:"redemptions" yaml:"redemptions"`
	MaxRedemptions string    `json:"max_redemptions" yaml:"max_redemptions"`
	Code           string    `json:"code" yaml:"code"`
	Zone           string    `json:"zone" yaml:"zone"`
	CreatedAt      time.Time `json:"created_at" yaml:"created_at"`
	ExternalID     string    `json:"external_id" yaml:"external_id"`
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

type AgentVolume struct {
	MountPoint string `json:"mount_point" yaml:"mount_point"`
	Disk       string `json:"disk" yaml:"disk"`
}

type NewNodesResponse struct {
	Nodes []*NewNodeResponseItem `json:"nodes" yaml:"nodes"`
}

type NewAgentsResponse struct {
	Agents []*NewAgentResponse `json:"agents" yaml:"agents"`
}

type NewNodeResponseItem struct {
	ID            string                 `json:"id" yaml:"id"`
	Name          string                 `json:"name" yaml:"name"`
	Label         *string                `json:"label" yaml:"label"`
	Configuration map[string]interface{} `json:"config" yaml:"config"`
	PublicIP      string                 `json:"public_ip" yaml:"public_ip"`
	PrivateIP     string                 `json:"private_ip" yaml:"private_ip"`
	Agents        []NewAgentResponse     `json:"agents" yaml:"agents"`
	CRN           map[string]interface{} `json:"crn" yaml:"crn"`
}

type NewAgentResponse struct {
	ID       string                 `json:"id" yaml:"id"`
	Secret   string                 `json:"secret" yaml:"secret"`
	Port     int                    `json:"port" yaml:"port"`
	Features map[string]interface{} `json:"features" yaml:"features"`
	Volume   AgentVolume            `json:"volume" yaml:"volume"`
	CRN      map[string]interface{} `json:"crn" yaml:"crn"`
}

type NewNode struct {
	ID            string                 `json:"id" yaml:"id"`
	Name          string                 `json:"name" yaml:"name"`
	Label         *string                `json:"label" yaml:"label"`
	Configuration map[string]interface{} `json:"config" yaml:"config"`
	CreatedAt     time.Time              `json:"created_at" yaml:"created_at"`
	DeletedAt     *time.Time             `json:"deleted_at" yaml:"deleted_at"`
	NexusID       string                 `json:"nexus_id" yaml:"nexus_id"`
	PrivateIP     string                 `json:"private_ip" yaml:"private_ip"`
	PublicIP      string                 `json:"public_ip" yaml:"public_ip"`
	CRN           map[string]interface{} `json:"crn" yaml:"crn"`
}

type NewAgent struct {
	ID                 string                 `json:"id" yaml:"id"`
	NodeID             string                 `json:"node_id" yaml:"node_id"`
	PublicKey          *string                `json:"public_key" yaml:"public_key"`
	Features           map[string]interface{} `json:"features" yaml:"features"`
	CreatedAt          time.Time              `json:"created" yaml:"created"`
	DeletedAt          *time.Time             `json:"deleted" yaml:"deleted"`
	LastStatusUpdateAt time.Time              `json:"last_status_update_at" yaml:"last_status_update_at"`
	Secret             string                 `json:"secret" yaml:"secret"`
	Online             bool                   `json:"online" yaml:"online"`
	ConnectedOn        *time.Time             `json:"connected_on" yaml:"connected_on"`
	Model              AgentModel             `json:"model" yaml:"model"`
	Serial             *string                `json:"serial" yaml:"serial"`
	Version            *string                `json:"version" yaml:"version"`
	Port               int                    `json:"cccp_port" yaml:"cccp_port"`
	TotalStorage       int64                  `json:"total_storage" yaml:"total_storage"`
	UsedStorage        int64                  `json:"used_storage" yaml:"used_storage"`
	Volume             AgentVolume            `json:"volume" yaml:"volume"`
	CRN                map[string]interface{} `json:"crn" yaml:"crn"`
	AnnouncedPrivateIP *string                `json:"announced_private_ip" yaml:"announced_private_ip"`
	AnnouncedPublicIP  *string                `json:"announced_public_ip" yaml:"announced_public_ip"`
}

type RedundancyClassExpanded struct {
	Status         ExpansionStatus `json:"status" yaml:"status"`
	Message        string          `json:"message" yaml:"message"`
	ExpandedSize   int64           `json:"expanded_size" yaml:"expanded_size"`
	AgentsInvolved []string        `json:"agents_involved" yaml:"agents_involved"`
}

type RCSession struct {
	Status RCSessionStatus `json:"status" yaml:"status"`
}
type RCProgress struct {
	Session           RCSession          `json:"session" yaml:"session"`
	ActiveRecoveries  []RecoveryProgress `json:"active_recoveries" yaml:"active_recoveries"`
	FailedRecoveries  []RecoveryProgress `json:"failed_recoveries" yaml:"failed_recoveries"`
	CompletionPercent float64            `json:"completion_percent" yaml:"completion_percent"`
}

type RecoveryProgress struct {
	RecoveryID      string  `json:"recovery_id" yaml:"recovery_id"`
	ProgressPercent float64 `json:"progress_percent" yaml:"progress_percent"`
}

type RedundancyClassErrorData struct {
	IssueFound []interface{} `json:"issue_found" yaml:"issue_found"`
}

type RedundancyClassRecovery struct {
	Message string `json:"message" yaml:"message"`
}

type SummaryDetailsWithStatusNullable struct {
	Details *SummaryDetails `json:"details" yaml:"details"`
	SummaryStatusNullable `yaml:",inline"`
}

type EvaluatedStatusType string

type SummaryStatusNullable struct {
	EvaluatedStatus              *EvaluatedStatusType `json:"evaluated_status" yaml:"evaluated_status"`
	EvaluatedStatusLastUpdatedAt *time.Time           `json:"evaluated_status_last_updated_at" yaml:"evaluated_status_last_updated_at"`
}

type SummaryDetails struct {
	Pending int `json:"pending" yaml:"pending"`
	Online  int `json:"online" yaml:"online"`
	Offline int `json:"offline" yaml:"offline"`
	Warning int `json:"warning" yaml:"warning"`
	Error   int `json:"error" yaml:"error"`
}

type HumanReadableStatus struct {
	Status            string `json:"status" yaml:"status"`
	ChildType         string `json:"child_type" yaml:"child_type"`
	ChildCount        int    `json:"child_count" yaml:"child_count"`
	ChildOnlineCount  int    `json:"child_online_count" yaml:"child_online_count"`
	ChildOfflineCount int    `json:"child_offline_count" yaml:"child_offline_count"`
	ChildPendingCount int    `json:"child_pending_count" yaml:"child_pending_count"`
	ChildWarningCount int    `json:"child_warning_count" yaml:"child_warning_count"`
	ChildErrorCount   int    `json:"child_error_count" yaml:"child_error_count"`
}

func (s *SummaryDetailsWithStatusNullable) ToHumanReadableStatus(childEntityType string) HumanReadableStatus {
	result := HumanReadableStatus{
		ChildType: childEntityType,
	}

	if s.EvaluatedStatus != nil {
		result.Status = string(*s.EvaluatedStatus)
	} else {
		result.Status = "unknown"
	}

	if s.Details != nil {
		result.ChildOnlineCount = s.Details.Online
		result.ChildOfflineCount = s.Details.Offline
		result.ChildPendingCount = s.Details.Pending
		result.ChildWarningCount = s.Details.Warning
		result.ChildErrorCount = s.Details.Error
		result.ChildCount = s.Details.Online + s.Details.Offline + s.Details.Pending + s.Details.Warning + s.Details.Error
	}

	return result
}

type NodeConfig struct {
	ID        string        `json:"id,omitempty" yaml:"id"`
	Name      string        `json:"name" yaml:"name"`
	PublicIP  string        `json:"public_ip" yaml:"public_ip"`
	PrivateIP string        `json:"private_ip" yaml:"private_ip"`
	Label     string        `json:"label,omitempty" yaml:"label"`
	Agents    []AgentConfig `json:"agents" yaml:"agents"`
}

type AgentConfig struct {
	ID         string `json:"id" yaml:"id"`
	MountPoint string `json:"mount_point" yaml:"mount_point"`
	Disk       string `json:"disk" yaml:"disk"`
	Port       int    `json:"port" yaml:"port"`
	Secret     string `json:"secret" yaml:"secret"`
}

type AnsibleConfig struct {
	Nodes []NodeConfig `json:"nodes" yaml:"nodes"`
}

type Gateway struct {
	ID             string     `json:"id" yaml:"id"`
	Name           string     `json:"name" yaml:"name"`
	Location       string     `json:"location" yaml:"location"`
	CreatedAt      time.Time  `json:"created_at" yaml:"created_at"`
	DeletedAt      *time.Time `json:"deleted_at" yaml:"deleted_at"`
	Secret         string     `json:"secret" yaml:"secret"`
	OrganizationID *string    `json:"organization_id" yaml:"organization_id"`
}

type GatewayWithGatewayTenant struct {
	Gateway       *Gateway       `json:"gateway" yaml:"gateway"`
	GatewayTenant *GatewayTenant `json:"gateway_tenant" yaml:"gateway_tenant"`
}

type GatewayTenant struct {
	ID            string                 `json:"id" yaml:"id"`
	GatewayID     string                 `json:"gateway_id" yaml:"gateway_id"`
	TenantID      string                 `json:"tenant_id" yaml:"tenant_id"`
	CreatedAt     time.Time              `json:"created_at" yaml:"created_at"`
	DeletedAt     *time.Time             `json:"deleted_at" yaml:"deleted_at"`
	HardDeleteAt  *time.Time             `json:"hard_delete_at" yaml:"hard_delete_at"`
	Configuration map[string]interface{} `json:"configuration" yaml:"configuration"`
	Hidden        bool                   `json:"hidden" yaml:"hidden"`
}

type GatewayInstanceListResponse struct {
	Data []*GatewayInstance `json:"data" yaml:"data"`
}

type GatewayInstance struct {
	ID              string                 `json:"id" yaml:"id"`
	GatewayID       string                 `json:"gateway_id" yaml:"gateway_id"`
	IP              string                 `json:"ip" yaml:"ip"`
	Status          map[string]interface{} `json:"status" yaml:"status"`
	StatusUpdatedAt *time.Time             `json:"status_updated_at" yaml:"status_updated_at"`
	CreatedAt       time.Time              `json:"created_at" yaml:"created_at"`
	Features        map[string]interface{} `json:"features" yaml:"features"`
	EvaluatedStatus GatewayStatus          `json:"evaluated_status" yaml:"evaluated_status"`
}

type DeviceRegistrationResponse struct {
	DeviceCode string `json:"device_code" yaml:"device_code"`
	ClientURL  string `json:"client_url" yaml:"client_url"`
}

type APIKeyResponse struct {
	APIKey string `json:"api_key" yaml:"api_key"`
	Status string `json:"status,omitempty" yaml:"status"`
}

type GetAgentEvaluatedStatusResponse struct {
	Status EvaluatedStatusType `json:"status" yaml:"status"`
}

type SignInToken struct {
	AccessToken  string `json:"access_token" yaml:"access_token"`
	RefreshToken string `json:"refresh_token" yaml:"refresh_token"`
}

type InfraClusterConnectCmdResponse struct {
	// mandatory, the generated command to use to execute a cluster-connect. It will contains the orgranization-id and its api-key
	Command string `json:"command" yaml:"command" example:"curl -fsSL https://download.operator.cubbit.eu/install-kubectl-plugin.sh | bash && kubectl cubbit cluster-connect --api-key=b9716863-5a0d-4e40-93e9-53f21f442f57 --org-id=00000000-0000-0000-0000-000000000000"`
}

// #region infrastructure

type InfrastructureCluster struct {
	ClusterID string `json:"cluster_id" yaml:"cluster_id"`
	Name      string `json:"name" yaml:"name"`
	Type      string `json:"type" yaml:"type"`
}

type InfraAggregateStatusCode string

const (
	StatusCodeOk      InfraAggregateStatusCode = "status_ok"
	StatusCodeWarning InfraAggregateStatusCode = "status_warning"
	StatusCodeError   InfraAggregateStatusCode = "status_error"
	StatusCodeInUse   InfraAggregateStatusCode = "status_in_use"
)

type InfraClusterType string

const (
	ClusterTypePhysical InfraClusterType = "physical"
	ClusterTypeVirtual  InfraClusterType = "virtual"
)

type InfraVirtualStorageType string

const (
	VirtualStorageTypeS3 InfraVirtualStorageType = "s3"
)

type InfraAggregateStatus struct {
	Code    string `json:"code" yaml:"code"`
	Details string `json:"details,omitempty" yaml:"details"`
}

type InfraNodeCPUInfo struct {
	Cores int `json:"cores" yaml:"cores"`
}

type InfraNodeRAMInfo struct {
	Available float64 `json:"available" yaml:"available"`
}

type InfraAggregateDiskDetail struct {
	DiskUUID              string               `json:"disk_uuid" yaml:"disk_uuid"`
	Path                  string               `json:"path" yaml:"path"`
	Used                  bool                 `json:"used" yaml:"used"`
	PVRef                 string               `json:"pv_ref,omitempty" yaml:"pv_ref"`
	TotalStorageSizeBytes int64                `json:"total_storage_size_bytes" yaml:"total_storage_size_bytes"`
	UsedStorageBytes      int64                `json:"used_storage_bytes" yaml:"used_storage_bytes"`
	Status                InfraAggregateStatus `json:"status" yaml:"status"`
}

type InfraAggregateNodeDetail struct {
	NodeID     string                     `json:"node_id" yaml:"node_id"`
	NodeName   string                     `json:"node_name" yaml:"node_name"`
	Status     InfraAggregateStatus       `json:"status" yaml:"status"`
	OSName     *string                    `json:"os_name,omitempty" yaml:"os_name"`
	CPU        *InfraNodeCPUInfo          `json:"cpu,omitempty" yaml:"cpu"`
	RAM        *InfraNodeRAMInfo          `json:"ram,omitempty" yaml:"ram"`
	ExternalIP *string                    `json:"external_ip,omitempty" yaml:"external_ip"`
	InternalIP *string                    `json:"internal_ip" yaml:"internal_ip"`
	Disks      []InfraAggregateDiskDetail `json:"disks" yaml:"disks"`
}

type InfraAggregateVirtualNodeDetail struct {
	NodeID               string                  `json:"node_id" yaml:"node_id"`
	NodeName             string                  `json:"node_name" yaml:"node_name"`
	Status               InfraAggregateStatus    `json:"status" yaml:"status"`
	StorageType          InfraVirtualStorageType `json:"storage_type" yaml:"storage_type"`
	StorageConfiguration map[string]any          `json:"storage_configuration" yaml:"storage_configuration"`
}

type InfraAggregateClusterDetail struct {
	LastUpdate   time.Time                         `json:"last_update" yaml:"last_update"`
	NextUpdate   time.Time                         `json:"next_update" yaml:"next_update"`
	IsUpdateOk   bool                              `json:"is_update_ok" yaml:"is_update_ok"`
	Nodes        []InfraAggregateNodeDetail        `json:"nodes" yaml:"nodes"`
	VirtualNodes []InfraAggregateVirtualNodeDetail `json:"virtual_nodes" yaml:"virtual_nodes"`
}

type InfraAggregateCluster struct {
	ClusterID string                      `json:"cluster_id" yaml:"cluster_id"`
	Name      string                      `json:"name" yaml:"name"`
	Details   InfraAggregateClusterDetail `json:"details" yaml:"details"`
	Type      InfraClusterType            `json:"type" yaml:"type"`
}

// #endregion

// #region swarm v5

type ListSwarmV5Item struct {
	ID                   string    `json:"id" yaml:"id"`
	Name                 string    `json:"name" yaml:"name"`
	TotalStorageBytes    int64     `json:"total_storage_bytes" yaml:"total_storage_bytes"`
	UsedStorageBytes     int64     `json:"used_storage_bytes" yaml:"used_storage_bytes"`
	CreatedAt            time.Time `json:"created_at" yaml:"created_at"`
	NexusCount           int       `json:"nexus_count" yaml:"nexus_count"`
	RedundancyClassCount int       `json:"redundancy_class_count" yaml:"redundancy_class_count"`
}

type ListSwarmV5ItemPresentation struct {
	ListSwarmV5Item        `yaml:",inline"`
	SummaryStatusNullable  `yaml:",inline"`
}

type SwarmV5 struct {
	ListSwarmV5Item        `yaml:",inline"`
	OrganizationID string                 `json:"organization_id" yaml:"organization_id"`
	OwnerID        string                 `json:"owner_id" yaml:"owner_id"`
	Description    *string                `json:"description,omitempty" yaml:"description"`
	Configuration  map[string]interface{} `json:"configuration" yaml:"configuration"`
	CreationStatus string                 `json:"creation_status,omitempty" yaml:"creation_status"`
}

type SwarmV5Presentation struct {
	SwarmV5                `yaml:",inline"`
	SummaryStatusNullable  `yaml:",inline"`
}

// #endregion

// #region swarm creation v5

type CreateSwarmV5Request struct {
	Name              string                   `json:"name" yaml:"name"`
	Description       *string                  `json:"description,omitempty" yaml:"description"`
	OwnerID           *string                  `json:"owner_id,omitempty" yaml:"owner_id"`
	Configuration     map[string]interface{}   `json:"configuration" yaml:"configuration"`
	Nexuses           []NexusV5Request         `json:"nexuses" yaml:"nexuses"`
	RedundancyClasses []RedundancyClassRequest `json:"redundancy_classes" yaml:"redundancy_classes"`
}

type NexusV5Request struct {
	ClusterID    string               `json:"cluster_id" yaml:"cluster_id"`
	ClusterType  InfraClusterType     `json:"cluster_type" yaml:"cluster_type"`
	Nodes        []NodeRequest        `json:"nodes,omitempty" yaml:"nodes"`
	VirtualNodes []VirtualNodeRequest `json:"virtual_nodes,omitempty" yaml:"virtual_nodes"`
}

type NodeRequest struct {
	ServerID string          `json:"server_id" yaml:"server_id"`
	Volumes  []VolumeRequest `json:"volumes" yaml:"volumes"`
}

type VolumeRequest struct {
	VolumeID string `json:"volume_id" yaml:"volume_id"`
}

type VirtualNodeRequest struct {
	ServerID string `json:"server_id" yaml:"server_id"`
}

type RedundancyClassRequest struct {
	Name              string   `json:"name" yaml:"name"`
	Description       *string  `json:"description,omitempty" yaml:"description"`
	OuterN            int      `json:"outer_n" yaml:"outer_n"`
	OuterK            int      `json:"outer_k" yaml:"outer_k"`
	InnerN            int      `json:"inner_n" yaml:"inner_n"`
	InnerK            int      `json:"inner_k" yaml:"inner_k"`
	AntiAffinityGroup int      `json:"anti_affinity_group,omitempty" yaml:"anti_affinity_group"`
	ClusterIDs        []string `json:"cluster_ids" yaml:"cluster_ids"`
}

type CreateSwarmV5Response struct {
	ID string `json:"id" yaml:"id"`
}

// #endregion

// #region gateway v5

// GatewayV5Status represents the lifecycle status of a gateway
type GatewayV5Status string

const (
	// GatewayV5StatusReady indicates the gateway is fully operational
	GatewayV5StatusReady GatewayV5Status = "ready"
	// GatewayV5StatusNotReady indicates the gateway is not yet ready
	GatewayV5StatusNotReady GatewayV5Status = "not-ready"
)

type GatewayV5GetRedundancyClass struct {
	ID   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
}

type GatewayV5GetResponse struct {
	ID                string                        `json:"id" yaml:"id"`
	Name              string                        `json:"name" yaml:"name"`
	Slug              string                        `json:"slug" yaml:"slug"`
	Type              CubbitIngressType             `json:"type" yaml:"type"`
	RedundancyClasses []GatewayV5GetRedundancyClass `json:"redundancy_classes" yaml:"redundancy_classes"`
	Status            GatewayV5Status               `json:"status" yaml:"status"`
}

// GatewayV5ListItemResponse represents a gateway summary item returned in paginated list responses
type GatewayV5ListItemResponse struct {
	// The unique identifier of the gateway
	ID string `json:"id" yaml:"id"`
	// The name of the gateway
	Name string `json:"name" yaml:"name"`
	// The slug of the gateway, a URL-friendly identifier
	Slug string `json:"slug" yaml:"slug"`
	// The ingress type configuration (manual, singlecluster, multicluster_controller, multicluster_worker)
	Type string `json:"type" yaml:"type"`
	// The number of swarms associated with this gateway
	NumberOfSwarms int `json:"number_of_swarms" yaml:"number_of_swarms"`
	// The number of tenants configured on this gateway
	NumberOfTenants int `json:"number_of_tenants" yaml:"number_of_tenants"`
	// The current lifecycle status of the gateway
	Status GatewayV5Status `json:"status" yaml:"status"`
}

// #endregion

// #region process

type ProcessType string
type ProcessStatus string
type ProcessStep string

const (
	ProcessTypeSwarmCreation   ProcessType = "swarm_creation"
	ProcessTypeGatewayCreation ProcessType = "gateway_creation"
	ProcessTypeTenantCreation  ProcessType = "tenant_creation"

	ProcessStatusRunning ProcessStatus = "running"
	ProcessStatusSuccess ProcessStatus = "success"
	ProcessStatusFailed  ProcessStatus = "failed"

	ProcessStepInitializing ProcessStep = "initializing"
	ProcessStepCompleted    ProcessStep = "completed"

	ProcessStepCreatingNexuses       ProcessStep = "creating_nexuses"
	ProcessStepCreatingAgents        ProcessStep = "creating_agents"
	ProcessStepCreatingRedundancyCls ProcessStep = "creating_redundancy_classes"

	ProcessStepGatewayProfileDeployment ProcessStep = "gateway_profile_deployment"
	ProcessStepGatewayInstallation      ProcessStep = "gateway_installation"

	ProcessStepCreatingTenantGateway            ProcessStep = "creating_tenant_gateway"
	ProcessStepWaitingForTenantGatewayToBeReady ProcessStep = "waiting_for_tenant_gateway_to_be_ready"
)

type Process struct {
	ID        string          `json:"id" yaml:"id"`
	Type      ProcessType     `json:"type" yaml:"type"`
	CreatedAt time.Time       `json:"created_at" yaml:"created_at"`
	OwnerID   string          `json:"owner_id" yaml:"owner_id"`
	Step      ProcessStep     `json:"step" yaml:"step"`
	Status    ProcessStatus   `json:"status" yaml:"status"`
	Data      json.RawMessage `json:"data" yaml:"data"`
}

type ProcessError struct {
	Code    string `json:"code" yaml:"code"`
	Message string `json:"message" yaml:"message"`
}

type GatewayCreationProcessData struct {
	ID    string        `json:"id" yaml:"id"`
	Error *ProcessError `json:"error,omitempty" yaml:"error"`
}

type GatewayCreationProcess struct {
	ID        string                     `json:"id" yaml:"id"`
	Type      ProcessType                `json:"type" yaml:"type"`
	CreatedAt time.Time                  `json:"created_at" yaml:"created_at"`
	OwnerID   string                     `json:"owner_id" yaml:"owner_id"`
	Step      ProcessStep                `json:"step" yaml:"step"`
	Status    ProcessStatus              `json:"status" yaml:"status"`
	Data      GatewayCreationProcessData `json:"data" yaml:"data"`
}

func (p *Process) CastToGatewayCreationProcess() (*GatewayCreationProcess, bool) {
	var data GatewayCreationProcessData
	if err := json.Unmarshal(p.Data, &data); err != nil {
		return nil, false
	}

	return &GatewayCreationProcess{
		ID:        p.ID,
		Type:      p.Type,
		CreatedAt: p.CreatedAt,
		OwnerID:   p.OwnerID,
		Step:      p.Step,
		Status:    p.Status,
		Data:      data,
	}, true
}

type TenantCreationProcessData struct {
	TenantID string                             `json:"tenant_id" yaml:"tenant_id"`
	Gateways []TenantGatewayCreationProcessData `json:"gateways,omitempty" yaml:"gateways"`
	Error    *ProcessError                      `json:"error,omitempty" yaml:"error"`
}

type TenantGatewayCreationProcessData struct {
	ID     string        `json:"id" yaml:"id"`
	Step   string        `json:"step" yaml:"step"`
	Status ProcessStatus `json:"status" yaml:"status"`
	Error  *ProcessError `json:"error,omitempty" yaml:"error"`
}

type TenantCreationProcess struct {
	ID        string                    `json:"id" yaml:"id"`
	Type      ProcessType               `json:"type" yaml:"type"`
	CreatedAt time.Time                 `json:"created_at" yaml:"created_at"`
	OwnerID   string                    `json:"owner_id" yaml:"owner_id"`
	Step      ProcessStep               `json:"step" yaml:"step"`
	Status    ProcessStatus             `json:"status" yaml:"status"`
	Data      TenantCreationProcessData `json:"data" yaml:"data"`
}

func (p *Process) CastToTenantCreationProcess() (*TenantCreationProcess, bool) {
	var data TenantCreationProcessData
	if err := json.Unmarshal(p.Data, &data); err != nil {
		return nil, false
	}

	return &TenantCreationProcess{
		ID:        p.ID,
		Type:      p.Type,
		CreatedAt: p.CreatedAt,
		OwnerID:   p.OwnerID,
		Step:      p.Step,
		Status:    p.Status,
		Data:      data,
	}, true
}

// #region domain

type DomainDTO struct {
	ID             string     `json:"id" yaml:"id"`
	DomainName     string     `json:"domain_name" yaml:"domain_name"`
	CreatedAt      time.Time  `json:"created_at" yaml:"created_at"`
	DeletedAt      *time.Time `json:"deleted_at,omitempty" yaml:"deleted_at"`
	Challenge      string     `json:"challenge" yaml:"challenge"`
	VerifiedAt     *time.Time `json:"verified_at,omitempty" yaml:"verified_at"`
	OrganizationID string     `json:"organization_id" yaml:"organization_id"`
	IsShared       bool       `json:"is_shared" yaml:"is_shared"`
}

type DomainVerifyResult struct {
	Verified bool    `json:"verified" yaml:"verified"`
	Reasons  *string `json:"reasons,omitempty" yaml:"reasons"`
}

// #region users v3

type BulkCreateIAMUserResponseItem struct {
	ID       string  `json:"id" yaml:"id"`
	Username string  `json:"username" yaml:"username"`
	Email    *string `json:"email" yaml:"email"`
	Created  bool    `json:"created" yaml:"created"`
	Status   string  `json:"status" yaml:"status"`
}

type BulkCreateIAMUsersResponse struct {
	Data  []BulkCreateIAMUserResponseItem `json:"data" yaml:"data"`
	Count int                             `json:"count" yaml:"count"`
}

type BulkGenerateSaltResponseItem struct {
	Username string `json:"username" yaml:"username"`
	Salt     string `json:"salt" yaml:"salt"`
}

type BulkGenerateSaltsResponse struct {
	Data  []BulkGenerateSaltResponseItem `json:"data" yaml:"data"`
	Count int                            `json:"count" yaml:"count"`
}

type IAMUserListItem struct {
	ID               string        `json:"id" yaml:"id"`
	Username         string        `json:"username" yaml:"username"`
	FirstName        *string       `json:"first_name" yaml:"first_name"`
	LastName         *string       `json:"last_name" yaml:"last_name"`
	Enabled          bool          `json:"enabled" yaml:"enabled"`
	CreatedAt        time.Time     `json:"created_at" yaml:"created_at"`
	DeletedAt        *time.Time    `json:"deleted_at" yaml:"deleted_at"`
	Emails           []IAMUserEmail `json:"emails" yaml:"emails"`
	PoliciesCount    int           `json:"policies_count" yaml:"policies_count"`
	Status           string        `json:"status" yaml:"status"`
	IsRoot           bool          `json:"is_root" yaml:"is_root"`
	TwoFactorEnabled bool          `json:"two_factor_enabled" yaml:"two_factor_enabled"`
	OrganizationID   string        `json:"organization_id" yaml:"organization_id"`
	LastActivityAt   *time.Time    `json:"last_activity_at" yaml:"last_activity_at"`
}

// #endregion

// #region tenant v5

type TenantV5DTO struct {
	ID          string          `json:"id" yaml:"id"`
	Name        string          `json:"name" yaml:"name"`
	Slug        string          `json:"slug" yaml:"slug"`
	Description *string         `json:"description" yaml:"description"`
	CreatedAt   time.Time       `json:"created_at" yaml:"created_at"`
	Storage     *UsageDTO       `json:"storage" yaml:"storage"`
	Bandwidth   *UsageDTO       `json:"bandwidth" yaml:"bandwidth"`
	Settings    *TenantSettings `json:"settings" yaml:"settings"`
	ZKEnabled   bool            `json:"zk_enabled" yaml:"zk_enabled"`
}

type UsageDTO struct {
	Consumed   int64   `json:"consumed" yaml:"consumed"`
	Reserved   int64   `json:"reserved" yaml:"reserved"`
	Percentage float64 `json:"percentage" yaml:"percentage"`
}

// #endregion
