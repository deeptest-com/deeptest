package serverDomain

import (
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/kataras/iris/v12"
)

type TestInterface struct {
	Id int64 `json:"id"`

	Title  string `json:"title"`
	Desc   string `json:"desc"`
	IsLeaf bool   `json:"isLeaf"`

	DebugInterfaceId uint  `json:"debugInterfaceId"`
	ParentId         int64 `json:"parentId"`
	ProjectId        uint  `json:"projectId"`
	ServeId          uint  `json:"serveId"`
	UseID            uint  `json:"useId"`

	Ordr     int              `json:"ordr"`
	Children []*TestInterface `gorm:"-" json:"children"`
	Slots    iris.Map         `gorm:"-" json:"slots"`
}

type TestInterfaceLoadReq struct {
	ServeId   int `json:"serveId"`
	ProjectId int `json:"projectId"`
}

type TestInterfaceSaveReq struct {
	Id        uint   `json:"id"`
	Title     string `json:"title"`
	Mode      string `json:"mode"`
	TargetId  uint   `json:"targetId"`
	ServeId   uint   `json:"serveId"`
	ProjectId uint   `json:"projectId"`

	Type serverConsts.TestInterfaceType `json:"type"`
}

type TestInterfaceReq struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Desc   string `json:"desc"`
	Parent uint
}

type TestInterfaceMoveReq struct {
	DragKey int                  `json:"dragKey"`
	DropKey int                  `json:"dropKey"`
	DropPos serverConsts.DropPos `json:"dropPos"`
}
