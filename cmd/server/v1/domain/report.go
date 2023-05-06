package serverDomain

import (
	"github.com/aaronchen2k/deeptest/pkg/domain"
)

type ReportReqPaginate struct {
	_domain.PaginateReq
	Keywords   string `json:"keywords"`
	ScenarioId int    `json:"scenarioId"`
}
