package model

import (
	"time"
)

type Task struct {
	BaseModel
	BaseItem

	ScheduledStartTime *time.Time `json:"scheduledStartTime" gorm:"comment:'预计开始时间'"`
	ScheduledEndTime   *time.Time `json:"scheduledEndTime" gorm:"comment:'预计结束时间'"`
	EvaluatedPoints    int        `json:"evaluatedPoints" gorm:"comment:'预估工作量'"`
	ActualPoints       int        `json:"actualPoints" gorm:"comment:'实际工作量'"`
	Deadline           *time.Time `json:"deadline" gorm:"comment:'截止日期'"`
}

func (Task) TableName() string {
	return "item_task"
}
