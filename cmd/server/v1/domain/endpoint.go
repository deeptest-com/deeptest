package domain

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
)

type EndpointReqPaginate struct {
	_domain.PaginateReq
	ProjectId    string `json:"project_id"`
	CategoryPath string `json:"category_path"`
	Status       string `json:"status"`
	UserId       string `json:"user_id"`
	Title        string `json:"title"`
}

type EndpointReq struct {
	ProjectId  string          `json:"project_id"`
	Status     string          `json:"status"`
	Title      string          `json:"title"`
	Path       string          `json:"path"`
	Version    string          `json:"version"`
	ParentId   string          `json:"ParentId"`
	PathParams []domain.Param  `json:"path_params"`
	Interfaces []InterfaceResp `json:"interfaces"`
}

type EndpointRes struct {
	ProjectId  string          `json:"project_id"`
	Status     string          `json:"status"`
	CreateUser string          `json:"create_user"`
	Name       string          `json:"name"`
	Version    string          `json:"version"`
	Interfaces []InterfaceResp `json:"interfaces"`
}
