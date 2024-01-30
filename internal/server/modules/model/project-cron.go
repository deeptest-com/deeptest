package model

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"time"
)

type ProjectCron struct {
	BaseModel
	Switch   consts.SwitchStatus `json:"switch"`
	Cron     string              `json:"cron"`
	ExecTime *time.Time          `json:"execTime"`
	ConfigId uint                `json:"configId"`
}

func (ProjectCron) TableName() string {
	return "biz_project_cron"
}
