package serverDomain

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
)

type RoleReq struct {
	model.BaseRole
	Perms [][]string `json:"perms"`
}

type RoleReqPaginate struct {
	_domain.PaginateReq
	Name string `json:"name"`
}

type RoleResp struct {
	_domain.Model
	model.BaseRole
}
