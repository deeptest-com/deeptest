package model

import (
	"time"
)

type Story struct {
	BaseModel
	BaseItem

	Iteration          uint       `json:"iteration" gorm:"comment:'所属迭代'"`
	Module             uint       `json:"module" gorm:"comment:'所属模块'"`
	ScheduledStartDate *time.Time `json:"scheduledStartDate" gorm:"comment:'预计开始时间'"`
	ScheduledEndDate   *time.Time `json:"scheduledEndDate" gorm:"comment:'预计结束时间'"`
	EvaluatedWorkload  int        `json:"evaluatedWorkload" gorm:"comment:'预估工作量'"`
	ActualWorkload     int        `json:"actualWorkload" gorm:"comment:'实际工作量'"`
	Deadline           *time.Time `json:"deadline" gorm:"comment:'截止日期'"`
}

func (Story) TableName() string {
	return "item_story"
}
