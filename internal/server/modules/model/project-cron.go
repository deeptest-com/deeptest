package model

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"time"
)

type ProjectCron struct {
	BaseModel
	Name       string                `json:"name"`
	Switch     consts.SwitchStatus   `json:"switch"`
	Cron       string                `json:"cron"`
	Source     consts.CronSource     `json:"source"`
	ConfigId   uint                  `json:"configId"`
	ExecTime   *time.Time            `json:"execTime"`
	ExecStatus consts.CronExecStatus `json:"execStatus"`
	ExecErr    string                `json:"execErr"`
}

func (ProjectCron) TableName() string {
	return "biz_project_cron"
}
