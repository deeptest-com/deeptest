package model

import (
	"database/sql/driver"
	"github.com/goccy/go-json"
)

type SummaryUserList struct {
	UserName string `json:"user_name"`
	UserId   int64  `json:"user_id"`
}

type SummaryDetails struct {
	BaseModel

	ProjectId          int64               `json:"project_id"`
	ProjectName        string              `gorm:"type:text" json:"project_name"`
	ProjectDes         string              `gorm:"type:text" json:"project_des"`
	ProjectChineseName string              `gorm:"type:text" json:"project_chinese_name"`
	ProjectCreateTime  string              `gorm:"type:text" json:"project_create_time"`
	ScenarioTotal      int64               `json:"scenario_total"`
	InterfaceTotal     int64               `json:"interface_total"`
	ExecTotal          int64               `json:"exec_total"`
	PassRate           float64             `json:"pass_rate"`
	Coverage           float64             `json:"coverage"`
	AdminUser          string              `gorm:"type:text" json:"admin_user"`
	AdminId            int64               `json:"adminId"`
	Logo               string              `json:"logo"`
	IncludeExample     bool                `json:"include_example"`
	Logs               []*ExecLogProcessor `gorm:"-" json:"logs"`
}

// Scan 解码json字符串
func (summaryUserList *SummaryUserList) Scan(val interface{}) error {
	b, _ := val.([]byte)
	return json.Unmarshal(b, summaryUserList)
}

// Value 编码json
func (summaryUserList *SummaryUserList) Value() (value driver.Value, err error) {
	return json.Marshal(summaryUserList)
}

type SummaryCard struct {
	SummaryCardTotal
	ProjectTotal int64   `json:"project_total"`
	InterfaceHB  float64 `json:"interface_hb"`
	ScenarioHB   float64 `json:"scenario_hb"`
	CoverageHB   float64 `json:"coverage_hb"`
}

type SummaryCardTotal struct {
	ScenarioTotal  int64               `gorm:"column:scenario_total" json:"scenario_total"`
	InterfaceTotal int64               `gorm:"column:interface_total" json:"interface_total"`
	ExecTotal      int64               `gorm:"column:exec_total" json:"exec_total"`
	PassRate       float64             `gorm:"column:pass_rate" json:"pass_rate"`
	Coverage       float64             `gorm:"column:coverage" json:"coverage"`
	Logs           []*ExecLogProcessor `gorm:"-" json:"logs"`
}

type ProjectIdsGroupByUserId struct {
	UserId     int64  `gorm:"-" json:"user_id"`
	ProjectIds string `gorm:"-" json:"project_ids"`
}

type UserIdsGroupByProjectId struct {
	ProjectIds int64  `gorm:"-" json:"project_ids"`
	UserId     string `gorm:"-" json:"user_id"`
}

type UserIdAndName struct {
	UserId   int64  `json:"user_id"`
	UserName string `json:"user_name"`
}

func (SummaryDetails) TableName() string {
	return "biz_summary_details"
}
