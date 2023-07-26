package serverDomain

import (
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
)

type PlanReportReqPaginate struct {
	_domain.PaginateReq
	ExecuteStartTime int64  `json:"executeStartTime"`
	ExecuteEndTime   int64  `json:"executeEndTime"`
	CreateUserId     uint   `json:"createUserId"`
	CreateUserIds    []uint `json:"createUserIds"`
	Keywords         string `json:"keywords"`
	PlanId           uint   `json:"planId"`
}
