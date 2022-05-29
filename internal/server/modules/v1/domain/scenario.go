package serverDomain

import (
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/aaronchen2k/deeptest/pkg/domain"
)

type ScenarioReqPaginate struct {
	_domain.PaginateReq
	Keywords string `json:"keywords"`
	Enabled  string `json:"enabled"`
}

type ScenarioNodeReq struct {
	Mode      serverConsts.NodeCreateMode `json:"mode"`
	Type      serverConsts.NodeCreateType `json:"type"`
	Target    int                         `json:"target"`
	Name      string                      `json:"name"`
	Id        int                         `json:"id"`
	ProjectId int                         `json:"projectId"`
}

type ScenarioNodeMoveReq struct {
	DragKey int                  `json:"dragKey"`
	DropKey int                  `json:"dropKey"`
	DropPos serverConsts.DropPos `json:"dropPos"`
}
