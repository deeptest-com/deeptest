package serverDomain

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/pkg/domain"
)

type PlanReqPaginate struct {
	_domain.PaginateReq

	CategoryId int64             `json:"categoryId"`
	Status     consts.TestStatus `json:"status"`
	DirectorId uint              `json:"directorId"`
	Keywords   string            `json:"keywords"`
	Enabled    string            `json:"enabled"`
}

type PlanAddScenariosReq struct {
	SelectedNodes []ScenarioSimple `json:"selectedNodes"`

	TargetId  uint `json:"targetId"`
	ProjectId int  `json:"projectId"`
}
