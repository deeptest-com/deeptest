package domain

import (
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type PlanReqPaginate struct {
	_domain.PaginateReq

	CategoryId uint   `json:"categoryId"`
	Keywords   string `json:"keywords"`
	Enabled    string `json:"enabled"`
}

type PlanAddScenariosReq struct {
	SelectedNodes []ScenarioSimple `json:"selectedNodes"`

	TargetId  uint `json:"targetId"`
	ProjectId int  `json:"projectId"`
}

// category
type PlanCategory struct {
	Id       uint            `json:"id"`
	Name     string          `json:"name"`
	Desc     string          `json:"desc"`
	ParentId uint            `json:"parentId"`
	Children []*PlanCategory `json:"children"`
	Slots    iris.Map        `json:"slots"`
}

type PlanCategoryCreateReq struct {
	Name string `json:"name"`
	Mode string `json:"mode"`

	TargetId  uint `json:"targetId"`
	ProjectId uint `json:"projectId"`
}

type PlanCategoryReq struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	Parent uint
}

type PlanCategoryMoveReq struct {
	DragKey int                  `json:"dragKey"`
	DropKey int                  `json:"dropKey"`
	DropPos serverConsts.DropPos `json:"dropPos"`
}
