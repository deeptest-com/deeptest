package serverDomain

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/aaronchen2k/deeptest/pkg/domain"
)

type ScenarioReqPaginate struct {
	_domain.PaginateReq
	Keywords string `json:"keywords"`
	Enabled  string `json:"enabled"`
}

type ScenarioAddInterfacesReq struct {
	SelectedNodes []InterfaceSimple `json:"selectedNodes"`

	TargetId  uint `json:"targetId"`
	ProjectId int  `json:"projectId"`
}

type ScenarioAddScenarioReq struct {
	Name              string                   `json:"name"`
	Mode              string                   `json:"mode"`
	ProcessorCategory consts.ProcessorCategory `json:"processorCategory"`
	ProcessorType     consts.ProcessorType     `json:"processorType"`

	TargetProcessorCategory consts.ProcessorCategory `json:"targetProcessorCategory"`
	TargetProcessorType     consts.ProcessorType     `json:"targetProcessorType"`
	TargetProcessorId       int                      `json:"targetProcessorId"`
	ProjectId               int                      `json:"projectId"`
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

type ScenarioProcessorReq struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
