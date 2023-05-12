package serverDomain

import (
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	"time"
)

type PlanReportReqPaginate struct {
	_domain.PaginateReq
	ExecuteStartTime *time.Time `json:"executeStartTime"`
	ExecuteEndTime   *time.Time `json:"executeEndTime"`
	CreateUserId     uint       `json:"createUserId"`
	Keywords         string     `json:"keywords"`
	PlanId           uint       `json:"planId"`
}
