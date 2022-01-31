package serverDomain

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
)

type ProjectReq struct {
	model.Project
}

type ProjectReqPaginate struct {
	_domain.PaginateReq
	Keywords string `json:"keywords"`
	Enabled  string `json:"enabled"`
}

type ProjectResp struct {
	model.Project
}
