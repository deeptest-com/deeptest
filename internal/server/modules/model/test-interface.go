package model

import (
	"github.com/kataras/iris/v12"
)

type TestInterface struct {
	BaseModel

	Title  string `json:"title"`
	Desc   string `json:"desc"`
	IsLeaf bool   `json:"isLeaf"`

	DebugInterfaceId uint `json:"debugInterfaceId"`
	ParentId         uint `json:"parentId"`
	ProjectId        uint `json:"projectId"`
	ServeId          uint `json:"serveId"`
	UseID            uint `json:"useId"`

	Ordr     int              `json:"ordr"`
	Children []*TestInterface `gorm:"-" json:"children"`
	Slots    iris.Map         `gorm:"-" json:"slots"`
}

func (TestInterface) TableName() string {
	return "biz_test_interface"
}
