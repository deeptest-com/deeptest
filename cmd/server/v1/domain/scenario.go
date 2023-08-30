package serverDomain

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type ScenarioReqPaginate struct {
	_domain.PaginateReq

	CategoryId int64  `json:"categoryId"`
	Keywords   string `json:"keywords"`
	Enabled    string `json:"enabled"`
	Status     string `json:"status"`
	Priority   string `json:"priority"`
	Type       string `json:"type"`
}

type ScenarioAddInterfacesReq struct {
	InterfaceIds []int `json:"interfaceIds"`
	TargetId     uint  `json:"targetId"`
	CreateBy     uint  `json:"createBy"`
}
type ScenarioAddInterfacesFromTreeReq struct {
	SelectedNodes []DiagnoseInterface `json:"selectedNodes"`
	TargetId      uint                `json:"targetId"`
	CreateBy      uint                `json:"createBy"`
}
type ScenarioAddCasesFromTreeReq struct {
	SelectedNodes []EndpointCaseTree `json:"selectedNodes"`
	TargetId      uint               `json:"targetId"`
	CreateBy      uint               `json:"createBy"`
}

type ScenarioAddScenarioReq struct {
	Name                  string                       `json:"name"`
	Mode                  string                       `json:"mode"`
	ProcessorCategory     consts.ProcessorCategory     `json:"processorCategory"`
	ProcessorType         consts.ProcessorType         `json:"processorType"`
	ProcessorInterfaceSrc consts.ProcessorInterfaceSrc `json:"processorInterfaceSrc"`

	TargetProcessorCategory consts.ProcessorCategory `json:"targetProcessorCategory"`
	TargetProcessorType     consts.ProcessorType     `json:"targetProcessorType"`
	TargetProcessorId       int                      `json:"targetProcessorId"`
	ProjectId               uint                     `json:"projectId"`
	CreateBy                uint                     `json:"createBy"`
	Disable                 bool                     `json:"disable"`
	Comments                string                   `json:"comments"`
	Method                  consts.HttpMethod        `json:"method"`
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
	ProcessorCategory consts.ProcessorCategory `json:"processorCategory"`
	ProcessorType     consts.ProcessorType     `json:"processorType"`
	ProcessorId       uint                     `json:"processorId"`
	InterfaceId       uint                     `json:"interfaceId"`

	Id   int    `json:"id"`
	Name string `json:"name"`
}

// category
type ScenarioCategory struct {
	Id       uint                `json:"id"`
	Name     string              `json:"name"`
	Desc     string              `json:"desc"`
	ParentId uint                `json:"parentId"`
	Children []*ScenarioCategory `json:"children"`
	Slots    iris.Map            `json:"slots"`
}

type ScenarioCategoryCreateReq struct {
	Name      string `json:"name"`
	Mode      string `json:"mode"`
	ServeId   uint   `json:"serveId"`
	ModuleId  string `json:"moduleId"`
	TargetId  uint   `json:"targetId"`
	ProjectId uint   `json:"projectId"`
}

type ScenarioCategoryReq struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	Parent uint
}

type ScenarioCategoryMoveReq struct {
	DragKey int                  `json:"dragKey"`
	DropKey int                  `json:"dropKey"`
	DropPos serverConsts.DropPos `json:"dropPos"`
}

type ScenarioSimple struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	ParentId  int    `json:"parentId"`
	ProjectId int    `json:"projectId"`
	IsDir     bool   `json:"isDir"`

	Ordr     int              `json:"ordr"`
	Children []ScenarioSimple `json:"children"`
}

type ScenarioProcessorInfo struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Comments string `json:"comments"`
}

type ScenarioPlanReqPaginate struct {
	_domain.PaginateReq
	ProjectId  uint              `json:"projectId"`
	CategoryId int64             `json:"categoryId"`
	Status     consts.TestStatus `json:"status"`
	AdminId    uint              `json:"adminId"`
	Ref        bool              `json:"ref"`
	Keywords   string            `json:"keywords"`
}

type ScenarioCurlImportReq struct {
	Content  string `json:"content"`
	TargetId uint   `json:"targetId"`
	CreateBy uint   `json:"createBy"`
}
