package model

import (
	"github.com/kataras/iris/v12"
)

type ScenarioCategory struct {
	BaseModel

	Name   string `json:"name"`
	Desc   string `json:"desc"`
	IsLeaf bool   `json:"isLeaf"`

	ParentId  uint `json:"parentId"`
	ProjectId uint `json:"projectId"`
	UseID     uint `json:"useId"`

	Ordr     int          `json:"ordr"`
	Children []*Processor `gorm:"-" json:"children"`
	Slots    iris.Map     `gorm:"-" json:"slots"`
}

func (ScenarioCategory) TableName() string {
	return "biz_scenario_category"
}
