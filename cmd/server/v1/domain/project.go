package domain

import (
	"github.com/aaronchen2k/deeptest/pkg/domain"
)

type ProjectReq struct {
	_domain.Model
	ProjectBase
}

type ProjectReqPaginate struct {
	_domain.PaginateReq
	Keywords string `json:"keywords"`
	Enabled  string `json:"enabled"`
}

type ProjectResp struct {
	_domain.PaginateReq
	ProjectBase
}

type ProjectBase struct {
	Name string `json:"name"`
	Desc string `json:"desc" gorm:"column:descr"`

	SchemaId uint `json:"schemaId"`
	OrgId    uint `json:"orgId"`
}
