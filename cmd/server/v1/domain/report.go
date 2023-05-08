package serverDomain

import (
	"github.com/aaronchen2k/deeptest/pkg/domain"
	"time"
)

type ReportReqPaginate struct {
	_domain.PaginateReq
	ExecuteStartTime *time.Time `json:"executeStartTime"`
	ExecuteEndTime   *time.Time `json:"executeEndTime"`
	CreateUserId     uint       `json:"createUserId"`
	Keywords         string     `json:"keywords"`
	ScenarioId       int        `json:"scenarioId"`
}
