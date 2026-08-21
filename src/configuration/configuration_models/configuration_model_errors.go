package configuration_models

import "errors"

var (
	ErrConfigurationVersionNotSupported = errors.New("configuration version not supported")
	ErrConfigurationIsAlreadyLoaded     = errors.New("configuration is already loaded")
	ErrConfigurationIsNotLoaded         = errors.New("configuration is not loaded")
	ErrActiveProfileNotFound            = errors.New("active profile not found among configured profiles")
	ErrActiveProfileNotSet              = errors.New("active profile not set in configuration")
	ErrProfileNotFound                  = errors.New("profile not found")
	ErrInvalidProfile                   = errors.New("invalid profile")
)
