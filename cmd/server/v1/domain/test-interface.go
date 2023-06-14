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

type TestInterfaceCreateReq struct {
	Name      string `json:"name"`
	Mode      string `json:"mode"`
	ServeId   uint   `json:"serveId"`
	ModuleId  string `json:"moduleId"`
	TargetId  uint   `json:"targetId"`
	ProjectId uint   `json:"projectId"`
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
