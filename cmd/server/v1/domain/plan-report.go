package serverDomain

import (
	_domain "github.com/deeptest-com/deeptest/pkg/domain"
)

type PlanReportReqPaginate struct {
	_domain.PaginateReq
	ExecuteStartTime int64  `json:"executeStartTime"`
	ExecuteEndTime   int64  `json:"executeEndTime"`
	CreateUserId     string `json:"createUserId"`
	Keywords         string `json:"keywords"`
	PlanId           uint   `json:"planId"`
}
