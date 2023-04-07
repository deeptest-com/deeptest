package domain

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
)

type EndpointReqPaginate struct {
	_domain.PaginateReq
	ProjectId    int64  `json:"projectId"`
	CategoryId   uint   `json:"categoryId"`
	Status       int64  `json:"status"`
	CreateUser   string `json:"createUser"`
	Title        string `json:"title"`
	ServeId      uint   `json:"serveId"`
	ServeVersion string `json:"serveVersion"`
}

type OpenApiParam struct {
	domain.Param
	Format      string `json:"format"`
	Example     string `json:"example"`
	Pattern     string `json:"pattern"`
	MinLength   string `json:"minLength"`
	MaxLength   int64  `json:"maxLength"`
	Default     string `json:"default"`
	MultipleOf  int64  `json:"multipleOf"`
	MinItems    int64  `json:"minItems"`
	MaxItems    int64  `json:"maxItems"`
	UniqueItems bool   `json:"uniqueItems"`
	Ref         string `json:"ref"`
}

type EndpointReq struct {
	ID         int64           `json:"id"`
	ProjectId  int64           `json:"projectId"`
	ServeId    int64           `json:"serveId"`
	Status     int64           `json:"status"`
	Title      string          `json:"title" validate:"required"`
	Path       string          `json:"path"`
	Version    string          `json:"version"`
	CreateUser string          `json:"CreateUser"`
	CategoryId uint            `json:"categoryId"`
	PathParams []OpenApiParam  `gorm:"-" json:"pathParams"`
	Interfaces []InterfaceResp `gorm:"-" json:"interfaces"`
}

type EndpointRes struct {
	ProjectId  int64           `json:"project_id"`
	Status     int64           `json:"status"`
	Title      string          `json:"title"`
	Version    string          `json:"version"`
	PathParams []OpenApiParam  `json:"pathParams"`
	Interfaces []InterfaceResp `json:"interfaces"`
}

type EndpointVersionReq struct {
	EndpointId int64  `json:"endpointId"`
	Version    string `json:"version"`
}
