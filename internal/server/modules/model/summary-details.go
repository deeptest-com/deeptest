package model

import (
	"time"
)

type SummaryDetails struct {
	BaseModel

	ProjectId            int64      `json:"project_id"`
	ProjectName          string     `gorm:"type:text" json:"project_name"`
	ProjectChineseName   string     `gorm:"type:text" json:"project_chinese_name"`
	ScenarioTotal        int64      `json:"scenario_total"`
	InterfaceTotal       int64      `json:"interface_total"`
	ExecTotal            int64      `json:"exec_total"`
	PassRate             int        `json:"pass_rate"`
	Coverage             int        `json:"coverage"`
	AdminUser            string     `gorm:"type:text" json:"admin_user"`
	ProjectCreateTime    *time.Time `json:"project_create_time"`
	ProjectUserRankingId int        `json:"project_user_ranking_id"`

	Logs []*ExecLogProcessor `gorm:"-" json:"logs"`
}

func (SummaryDetails) TableName() string {
	return "biz_summary_details"
}
