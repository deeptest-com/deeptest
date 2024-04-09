package model

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"time"
)

type ProjectCron struct {
	BaseModel
	Name          string                `json:"name"`
	Switch        consts.SwitchStatus   `json:"switch"`
	Cron          string                `json:"cron"`
	Source        consts.CronSource     `json:"source"`
	ConfigId      uint                  `json:"configId"`
	ProjectId     uint                  `json:"projectId"`
	ExecTime      *time.Time            `json:"execTime"`
	ExecStatus    consts.CronExecStatus `json:"execStatus"`
	ExecErr       string                `gorm:"type:text" json:"execErr"`
	CreateUserId  uint                  `json:"createUserId"`
	UpdateUserId  uint                  `json:"updateUserId"`
	SwaggerConfig SwaggerSync           `gorm:"-" json:"swaggerReq"`
	LecangConfig  CronConfigLecang      `gorm:"-" json:"lecangReq"`
}

type ProjectCronList struct {
	ProjectCron
	CategoryId     int    `json:"categoryId"`
	CategoryName   string `json:"categoryName"`
	CreateUserName string `json:"createUserName"`
}

func (ProjectCron) TableName() string {
	return "biz_project_cron"
}
