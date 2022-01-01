package serverDomain

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
)

type PermReq struct {
	model.BasePerm
}

type PermReqPaginate struct {
	domain.PaginateReq
	Name string `json:"name"`
}

type PermResp struct {
	domain.Model
	model.BasePerm
}
