package api

import (
	"net/url"
	"strconv"
	"strings"
)

type URLBuilder struct {
	baseURL string
	path    []string
	query   url.Values
}

func NewURLBuilder(baseURL string) *URLBuilder {
	return &URLBuilder{
		baseURL: strings.TrimSuffix(baseURL, "/"),
		path:    make([]string, 0),
		query:   make(url.Values),
	}
}

func (ub *URLBuilder) Path(segments ...string) *URLBuilder {
	for _, segment := range segments {
		ub.path = append(ub.path, url.PathEscape(segment))
	}

	return ub
}

func (ub *URLBuilder) QueryParam(key, value string) *URLBuilder {
	ub.query.Add(key, value)
	return ub
}

func (ub *URLBuilder) QueryParamInt(key string, value int) *URLBuilder {
	ub.query.Add(key, strconv.Itoa(value))
	return ub
}

func (ub *URLBuilder) Build() string {
	urlStr := ub.baseURL

	if len(ub.path) > 0 {
		urlStr += "/" + strings.Join(ub.path, "/")
	}

	if len(ub.query) > 0 {
		urlStr += "?" + ub.query.Encode()
	}

	return urlStr
}
