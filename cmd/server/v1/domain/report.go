package serverDomain

import (
	"github.com/deeptest-com/deeptest/pkg/domain"
)

type ReportReqPaginate struct {
	_domain.PaginateReq
	ExecuteStartTime string `json:"executeStartTime"`
	ExecuteEndTime   string `json:"executeEndTime"`
	CreateUserId     uint   `json:"createUserId"`
	Keywords         string `json:"keywords"`
	ScenarioId       int    `json:"scenarioId"`
}

type ReferBugReq struct {
	ReportId uint   `json:"reportId"`
	BugId    string `json:"bugId"`
	BugType  uint   `json:"bugType"`
	Severity uint   `json:"severity"`
}
