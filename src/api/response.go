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
	ID          string                 `json:"id" example:"784da023-cc6c-4a46-8e31-bc0a521d19e0"`
	Name        string                 `json:"name" example:"Cubbit"`
	Description *string                `json:"description" example:"Cloud storage: privacy, powered by p2p collaborations and eco-friendly"`
	OwnerID     string                 `json:"owner_id" example:"847390b4-a5b0-4ef7-949d-a15e84875d7e"`
	CreatedAt   time.Time              `json:"created_at" example:"2023-01-18T12:42:59.089247Z"`
	DeletedAt   *time.Time             `json:"deleted_at" example:"2023-01-18T12:42:59.089247Z"`
	ImageUrl    *string                `json:"image_url" example:"https://s3.cubbit.io/my-new-test-bucket/Screenshot.png"`
	Settings    map[string]interface{} `json:"settings" example:"{}"`
}

type TenantList struct {
	Tenants []*Tenant `json:"tenants"`
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
	Emails             []OperatorEmail `json:"emails"`
	TwoFactorEnabled   bool            `json:"two_factor_enabled" example:"true"`
}

type OperatorEmail struct {
	ID        string    `json:"id" example:"5ff281ee-75e7-4543-a304-ca861521f2a7"`
	Email     string    `json:"email" example:"mario.rossi@cubbit.io"`
	Verified  bool      `json:"verified" example:"true"`
	Default   bool      `json:"default" example:"true"`
	CreatedAt time.Time `json:"created_at" example:"2023-01-18T12:42:59.089247Z"`
}
