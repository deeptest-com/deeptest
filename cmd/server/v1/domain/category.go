package serverDomain

import (
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/kataras/iris/v12"
)

// category
type Category struct {
	Id       int64       `json:"id"`
	Name     string      `json:"name"`
	Desc     string      `json:"desc"`
	ParentId int64       `json:"parentId"`
	Children []*Category `json:"children"`
	Slots    iris.Map    `json:"slots"`
	Count    int64       `json:"count"`
	EntityId uint        `json:"entityId"`
}

type CategoryCreateReq struct {
	Name      string                             `json:"name"`
	Mode      string                             `json:"mode"`
	Type      serverConsts.CategoryDiscriminator `json:"type"`
	ServeId   uint                               `json:"serveId"`
	ModuleId  string                             `json:"moduleId"`
	TargetId  int                                `json:"targetId"`
	ProjectId uint                               `json:"projectId"`
	EntityId  uint                               `json:"entityId"`
	IsEntity  bool                               `json:"isEntity"`
}

type CategoryReq struct {
	Id     int                                `json:"id"`
	Name   string                             `json:"name"`
	Desc   string                             `json:"desc"`
	Parent uint                               `json:"parent"`
	Type   serverConsts.CategoryDiscriminator `json:"type"`
}

type CategoryMoveReq struct {
	Type    serverConsts.CategoryDiscriminator `json:"type"`
	DragKey int                                `json:"dragKey"`
	DropKey int                                `json:"dropKey"`
	DropPos serverConsts.DropPos               `json:"dropPos"`
}

type CategoryCount struct {
	Count      int64
	CategoryId int64
}

type BatchAddSchemaRootReq struct {
	ProjectIds []uint `json:"projectIds"`
}
