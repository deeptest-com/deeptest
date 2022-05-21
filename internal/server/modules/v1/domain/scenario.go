package serverDomain

import (
	"github.com/aaronchen2k/deeptest/pkg/domain"
)

type ScenarioReq struct {
	_domain.Model
	ScenarioBase
}

type ScenarioReqPaginate struct {
	_domain.PaginateReq
	Keywords string `json:"keywords"`
	Enabled  string `json:"enabled"`
}

type ScenarioResp struct {
	_domain.PaginateReq
	ScenarioBase
}

type ScenarioBase struct {
	Name string `json:"name"`
	Desc string `json:"desc" gorm:"column:descr"`

	SchemaId uint `json:"schemaId"`
	OrgId    uint `json:"orgId"`
}
