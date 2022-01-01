package serverDomain

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
)

type TestCaseReq struct {
	model.TestCase
}

type TestCaseReqPaginate struct {
	domain.PaginateReq
	Name     string `json:"name"`
	Category string `json:"name"`
	Status   string `json:"status"`
}

type TestCaseResp struct {
	model.TestCase
}
