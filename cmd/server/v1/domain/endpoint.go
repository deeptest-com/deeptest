package domain

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
)

type EndpointReqPaginate struct {
	_domain.PaginateReq
	ProjectId    int64  `json:"project_id"`
	CategoryPath string `json:"category_path"`
	Status       int64  `json:"status"`
	UserId       int64  `json:"user_id"`
	Title        string `json:"title"`
	ServeId      uint   `json:"serveId"`
	ServeVersion string `json:"serveVersion"`
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
	ParentId   int64           `json:"ParentId"`
	PathParams []domain.Param  `gorm:"-" json:"pathParams"`
	Interfaces []InterfaceResp `gorm:"-" json:"interfaces"`
}

type EndpointRes struct {
	ProjectId  int64           `json:"project_id"`
	Status     int64           `json:"status"`
	Title      string          `json:"title"`
	Version    string          `json:"version"`
	PathParams []domain.Param  `json:"pathParams"`
	Interfaces []InterfaceResp `json:"interfaces"`
}

type EndpointVersionReq struct {
	EndpointId int64  `json:"endpointId"`
	Version    string `json:"version"`
}
