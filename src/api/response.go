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
