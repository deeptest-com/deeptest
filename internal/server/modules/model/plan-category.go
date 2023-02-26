package model

import (
	"github.com/kataras/iris/v12"
)

type PlanCategory struct {
	BaseModel

	Name   string `json:"name"`
	Desc   string `json:"desc"`
	IsLeaf bool   `json:"isLeaf"`

	ParentId  uint `json:"parentId"`
	ProjectId uint `json:"projectId"`
	UseID     uint `json:"useId"`

	Ordr     int             `json:"ordr"`
	Children []*PlanCategory `gorm:"-" json:"children"`
	Slots    iris.Map        `gorm:"-" json:"slots"`
}

func (PlanCategory) TableName() string {
	return "biz_plan_category"
}
