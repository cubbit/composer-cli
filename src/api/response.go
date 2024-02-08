package api

import (
	"errors"
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
	ImageUrl    *string         `json:"image_url" example:"https://s3.cubbit.io/my-new-test-bucket/Screenshot.png"`
	Settings    *TenantSettings `json:"settings"`
	Metadata    map[string]interface{}
	CouponID    *string `json:"coupon_id" example:"9OPADNOEJFNO"`
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
	ID            string                 `json:"id"`
	SwarmID       string                 `json:"swarm_id"`
	SwarmName     string                 `json:"swarm_name"`
	TenantID      string                 `json:"tenant_id"`
	Name          string                 `json:"name" example:"Cubbit"`
	Description   string                 `json:"description" example:"Cloud storage: privacy, powered by p2p collaborations and eco-friendly"`
	Default       bool                   `json:"default"`
	OwnerID       string                 `json:"owner_id" example:"847390b4-a5b0-4ef7-949d-a15e84875d7e"`
	Configuration map[string]interface{} `json:"configuration" example:"{}"`
	Metrics       map[string]interface{} `json:"metrics" example:"{}"`
	CreatedAt     time.Time              `json:"created_at" example:"2023-01-18T12:42:59.089247Z"`
	DeletedAt     time.Time              `json:"deleted_at" example:"2023-01-18T12:42:59.089247Z"`
}

type SwarmList struct {
	Swarms []Swarm `json:"swarms"`
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
}

type OperatorList struct {
	Operators []Operator `json:"operators"`
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
	Code       int           `json:"string"`
	StorageBH  int64         `json:"storage_bh"`
	StorageGB  int64         `json:"storage_gb"`
	EngressGB  int64         `json:"engress_gb"`
	FromTime   string        `json:"from_time"`
	ToTime     string        `json:"to_time"`
	Status     string        `json:"status"`
	Timestamp  time.Duration `json:"timestamp"`
}

type DistributorReportResponseModel struct {
	Report []TenantReport `json:"report"`
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
