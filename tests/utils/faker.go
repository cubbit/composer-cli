package utils

import (
	"strings"

	"github.com/bxcodec/faker/v4"
	"github.com/bxcodec/faker/v4/pkg/options"
)

func safeTrim(s string, max int) string {
	if len(s) <= max {
		return s
	}

	runes := []rune(s)
	if len(runes) > max {
		return string(runes[:max])
	}
	return s
}

type FakerInstance struct{}

func (f FakerInstance) Email(opts ...options.OptionFunc) string {
	return strings.ToLower(faker.Email(opts...))
}

func (f FakerInstance) Username(opts ...options.OptionFunc) string {
	return faker.Username(opts...)
}

func (f FakerInstance) Password(opts ...options.OptionFunc) string {
	return faker.Password(opts...)
}

func (f FakerInstance) Name(opts ...options.OptionFunc) string {

	return safeTrim(faker.Name(opts...), 128)
}

func (f FakerInstance) NameUnique(opts ...options.OptionFunc) string {
	return safeTrim(faker.Name(opts...), 128) + faker.UUIDHyphenated()
}

var Faker = FakerInstance{}
