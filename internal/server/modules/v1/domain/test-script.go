package serverDomain

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
)

type TestScriptReq struct {
	model.TestScript
}

type TestScriptReqPaginate struct {
	_domain.PaginateReq
	Keywords string `json:"keywords"`
	Enabled  string `json:"enabled"`
}

type TestScriptResp struct {
	model.TestScript
}
