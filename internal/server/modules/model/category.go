package model

import (
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/kataras/iris/v12"
)

type Category struct {
	BaseModel

	Name   string `json:"name"`
	Desc   string `json:"desc"`
	IsLeaf bool   `json:"isLeaf"`

	ParentId  int  `json:"parentId"`
	ProjectId uint `json:"projectId"`
	ServeId   uint `json:"serveId"`
	UseID     uint `json:"useId"`

	Ordr     int          `json:"ordr"`
	Children []*Processor `gorm:"-" json:"children"`
	Slots    iris.Map     `gorm:"-" json:"slots"`

	Type serverConsts.CategoryDiscriminator `json:"type"`
}

func (Category) TableName() string {
	return "biz_category"
}
