package model

import (
	"time"
)

type Feature struct {
	BaseModel
	BaseItem

	// build in
	requirementType      *time.Time `json:"requirementType" gorm:"comment:'需求类型'"`
	ScheduledReleaseTime *time.Time `json:"scheduledReleaseTime" gorm:"comment:'计划上线日期'"`
	Product              uint       `json:"Product" gorm:"comment:'所属产品'"`
	Project              uint       `json:"Project" gorm:"comment:'所属项目'"`

	// custom
	goal        string `json:"goal" gorm:"comment:'量化目标'"`
	achievement string `json:"achievement" gorm:"comment:'量化目标达成情况'"`
	kickoffTime string `json:"kickoffTime" gorm:"comment:'确认启动日期'"`
}

func (Feature) TableName() string {
	return "item_feature"
}
