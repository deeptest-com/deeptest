package model

import (
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/kataras/iris/v12"
)

type DiagnoseInterface struct {
	BaseModel
	CreatedBy uint                               `json:"createdBy"`
	Name      string                             `json:"name"`
	Title     string                             `json:"title"`
	Desc      string                             `json:"desc"`
	IsLeaf    bool                               `json:"isLeaf"`
	Type      serverConsts.DiagnoseInterfaceType `json:"type"`

	ParentId  uint `json:"parentId"`
	ServeId   uint `json:"serveId"`
	ProjectId uint `json:"projectId"`
	UseID     uint `json:"useId"`

	Ordr     int                  `json:"ordr"`
	Children []*DiagnoseInterface `gorm:"-" json:"children"`
	Slots    iris.Map             `gorm:"-" json:"slots"`

	DebugInterfaceId uint            `gorm:"default:0" json:"debugInterfaceId"`
	DebugData        *DebugInterface `gorm:"-" json:"debugData"`
}

func (DiagnoseInterface) TableName() string {
	return "biz_diagnose_interface"
}
