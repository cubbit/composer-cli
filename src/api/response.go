package api

import (
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
	Message string `json:"message"`
	Param   string `json:"param"`
}

type ChallengeResponseModel struct {
	Challenge string `json:"challenge"`
	Salt      string `json:"salt" example:"SGVsbG8gd29ybGQ="`
}

type TokenAndExpirationResponseModel struct {
	Token   string    `json:"token"`
	Exp     int       `json:"exp"`
	ExpDate time.Time `json:"expDate"`
}

type GenericIDResponseModel struct {
	ID string `json:"id"`
}

type Tenant struct {
	ID          string          `json:"id"`
	Name        string          `json:"name"`
	Description *string         `json:"description"`
	OwnerID     string          `json:"owner_id"`
	CreatedAt   time.Time       `json:"created_at"`
	DeletedAt   *time.Time      `json:"deleted_at"`
	Settings    *TenantSettings `json:"settings"`
	CouponID    *string         `json:"coupon_id"`
}

type TenantSettings struct {
	ConsoleURL        *string                                    `json:"console_url"`
	GatewayURL        *string                                    `json:"gateway_url"`
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
	Type  string `json:"type"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type TenantSettingsProject struct {
	DefaultMaxProjectEgressBandwidth *int64 `json:"default_max_egress_bandwidth"`
	DefaultMaxProjectStorage         *int64 `json:"default_max_storage"`
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
	Threshold int          `json:"threshold"`
	Notify    Notification `json:"notify"`
}

type Notification struct {
	NotifyOwner bool     `json:"notify_owner"`
	Emails      []string `json:"emails"`
}

type TenantList struct {
	Tenants []*Tenant `json:"tenants"`
}

type Swarm struct {
	ID             string                 `json:"id"`
	Name           string                 `json:"name"`
	Configuration  map[string]interface{} `json:"configuration"`
	Size           int64                  `json:"size"`
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
	ID                 string          `json:"id"`
	FirstName          string          `json:"first_name"`
	LastName           string          `json:"last_name"`
	Internal           bool            `json:"internal"`
	Banned             bool            `json:"banned"`
	CreatedAt          time.Time       `json:"created_at"`
	DeletedAt          *time.Time      `json:"deleted_at"`
	MaxAllowedProjects int             `json:"max_allowed_projects"`
	Email              string          `json:"email"`
	Emails             []OperatorEmail `json:"emails"`
	TwoFactorEnabled   bool            `json:"two_factor_enabled"`
	Status             string          `json:"status"`
	PolicyName         string          `json:"policy_name"`
}

type OperatorList struct {
	Operators []*Operator `json:"operators"`
}

type OperatorEmail struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Verified  bool      `json:"verified"`
	Default   bool      `json:"default"`
	CreatedAt time.Time `json:"created_at"`
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
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	OwnerID     string    `json:"owner_id"`
	CreatedAt   time.Time `json:"created_at"`
	DeletedAt   time.Time `json:"deleted_at"`
	ImageURL    string    `json:"image_url"`
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
	Name        string `json:"name"`
	Key         string `json:"key"`
	Description string `json:"description"`
}

type ZoneMap struct {
	Zones map[string]ZoneResponse `json:"zones"`
}

type Account struct {
	ID                 string              `json:"id"`
	FirstName          string              `json:"first_name"`
	LastName           string              `json:"last_name"`
	Internal           bool                `json:"internal"`
	Banned             bool                `json:"banned"`
	CreatedAt          time.Time           `json:"created_at"`
	DeletedAt          *time.Time          `json:"deleted_at"`
	MaxAllowedProjects int                 `json:"max_allowed_projects"`
	Emails             []AccountEmail      `json:"emails"`
	TwoFactorEnabled   bool                `json:"two_factor_enabled"`
	EndpointGateway    string              `json:"endpoint_gateway"`
	TenantID           string              `json:"tenant_id"`
	AuthProvider       AccountAuthProvider `json:"auth_provider"`
}

type AccountEmail struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Verified  bool      `json:"verified"`
	Default   bool      `json:"default"`
	CreatedAt time.Time `json:"created_at"`
}

type GenericPaginatedResponse[T interface{}] struct {
	Data     []T  `json:"data"`
	NextPage *int `json:"next_page"`
	Count    int  `json:"count"`
}

type ProjectItem struct {
	ProjectID          string     `json:"project_id"`
	ProjectName        string     `json:"project_name"`
	ProjectDescription string     `json:"project_description"`
	ProjectEmail       string     `json:"project_email"`
	ProjectCreatedAt   time.Time  `json:"project_created_at"`
	ProjectDeletedAt   *time.Time `json:"project_deleted_at"`
	ProjectBannedAt    *time.Time `json:"project_banned_at"`
	ProjectImageURL    string     `json:"project_image_url"`
	ProjectTenantID    string     `json:"project_tenant_id"`
	RootAccountEmail   string     `json:"root_account_email"`
	RootAccountID      string     `json:"root_account_id"`
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
	RingsNumber       *int     `json:"number_of_rings"`
	AntiAffinityGroup *int     `json:"anti_affinity_group"`
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
	Nexuses      []RingNexus `json:"nexuses"`
	Nodes        []RingNode  `json:"nodes"`
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

type AgentVolume struct {
	MountPoint string `json:"mount_point"`
	Disk       string `json:"disk"`
}

type NewNodesResponse struct {
	Nodes []*NewNodeResponseItem `json:"nodes"`
}

type NewAgentsResponse struct {
	Agents []*NewAgentResponse `json:"agents"`
}

type NewNodeResponseItem struct {
	ID            string                 `json:"id"`
	Name          string                 `json:"name"`
	Label         *string                `json:"label"`
	Configuration map[string]interface{} `json:"config"`
	PublicIP      string                 `json:"public_ip"`
	PrivateIP     string                 `json:"private_ip"`
	Agents        []NewAgentResponse     `json:"agents"`
	CRN           map[string]interface{} `json:"crn"`
}

type NewAgentResponse struct {
	ID       string                 `json:"id"`
	Secret   string                 `json:"secret"`
	Port     int                    `json:"port"`
	Features map[string]interface{} `json:"features"`
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
	ConnectedOn        *time.Time             `json:"connected_on"`
	Model              AgentModel             `json:"model"`
	Serial             *string                `json:"serial"`
	Version            *string                `json:"version"`
	Port               int                    `json:"cccp_port"`
	TotalStorage       int64                  `json:"total_storage"`
	UsedStorage        int64                  `json:"used_storage"`
	Volume             AgentVolume            `json:"volume"`
	CRN                map[string]interface{} `json:"crn"`
	AnnouncedPrivateIP *string                `json:"announced_private_ip"`
	AnnouncedPublicIP  *string                `json:"announced_public_ip"`
}

type RedundancyClassExpanded struct {
	Status         ExpansionStatus `json:"status"`
	Message        string          `json:"message"`
	ExpandedSize   int64           `json:"expanded_size"`
	AgentsInvolved []string        `json:"agents_involved"`
}

type RCSession struct {
	Status RCSessionStatus `json:"status"`
}
type RCProgress struct {
	Session           RCSession          `json:"session"`
	ActiveRecoveries  []RecoveryProgress `json:"active_recoveries"`
	FailedRecoveries  []RecoveryProgress `json:"failed_recoveries"`
	CompletionPercent float64            `json:"completion_percent"`
}

type RecoveryProgress struct {
	RecoveryID      string  `json:"recovery_id"`
	ProgressPercent float64 `json:"progress_percent"`
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

type HumanReadableStatus struct {
	Status            string `json:"status"`
	ChildType         string `json:"child_type"`
	ChildCount        int    `json:"child_count"`
	ChildOnlineCount  int    `json:"child_online_count"`
	ChildOfflineCount int    `json:"child_offline_count"`
	ChildPendingCount int    `json:"child_pending_count"`
	ChildWarningCount int    `json:"child_warning_count"`
	ChildErrorCount   int    `json:"child_error_count"`
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

type Gateway struct {
	ID             string     `json:"id"`
	Name           string     `json:"name"`
	Location       string     `json:"location"`
	CreatedAt      time.Time  `json:"created_at"`
	DeletedAt      *time.Time `json:"deleted_at"`
	Secret         string     `json:"secret"`
	OrganizationID *string    `json:"organization_id"`
}

type GatewayWithGatewayTenant struct {
	Gateway       *Gateway       `json:"gateway"`
	GatewayTenant *GatewayTenant `json:"gateway_tenant"`
}

type GatewayTenant struct {
	ID            string                 `json:"id"`
	GatewayID     string                 `json:"gateway_id"`
	TenantID      string                 `json:"tenant_id"`
	CreatedAt     time.Time              `json:"created_at"`
	DeletedAt     *time.Time             `json:"deleted_at"`
	HardDeleteAt  *time.Time             `json:"hard_delete_at"`
	Configuration map[string]interface{} `json:"configuration"`
	Hidden        bool                   `json:"hidden"`
}

type GatewayInstanceListResponse struct {
	Data []*GatewayInstance `json:"data"`
}

type GatewayInstance struct {
	ID              string                 `json:"id"`
	GatewayID       string                 `json:"gateway_id"`
	IP              string                 `json:"ip"`
	Status          map[string]interface{} `json:"status"`
	StatusUpdatedAt *time.Time             `json:"status_updated_at"`
	CreatedAt       time.Time              `json:"created_at"`
	Features        map[string]interface{} `json:"features"`
	EvaluatedStatus GatewayStatus          `json:"evaluated_status"`
}

type DeviceRegistrationResponse struct {
	DeviceCode string `json:"device_code"`
	ClientURL  string `json:"client_url"`
}

type APIKeyResponse struct {
	APIKey string `json:"api_key"`
	Status string `json:"status,omitempty"`
}

type GetAgentEvaluatedStatusResponse struct {
	Status EvaluatedStatusType `json:"status"`
}
