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

	ProjectId      int64   `json:"projectId"`
	ScenarioTotal  int64   `json:"scenarioTotal"`
	InterfaceTotal int64   `json:"interfaceTotal"`
	ExecTotal      int64   `json:"execTotal"`
	PassRate       float64 `json:"passRate"`
	Coverage       float64 `json:"coverage"`
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
	ProjectTotal int64   `json:"projectTotal"`
	InterfaceHB  float64 `json:"interfaceHb"`
	ScenarioHB   float64 `json:"scenarioHb"`
	CoverageHB   float64 `json:"coverageHb"`
}

type SummaryCardTotal struct {
	ScenarioTotal  int64               `gorm:"column:scenario_total" json:"scenarioTotal"`
	InterfaceTotal int64               `gorm:"column:interface_total" json:"interfaceTotal"`
	ExecTotal      int64               `gorm:"column:exec_total" json:"execTotal"`
	PassRate       float64             `gorm:"column:pass_rate" json:"passRate"`
	Coverage       float64             `gorm:"column:coverage" json:"coverage"`
	Logs           []*ExecLogProcessor `gorm:"-" json:"logs"`
}

type ProjectIdsGroupByUserId struct {
	UserId     int64  `gorm:"-" json:"userId"`
	ProjectIds string `gorm:"-" json:"projectIds"`
}

type UserIdsGroupByProjectId struct {
	ProjectIds int64  `gorm:"-" json:"projectIds"`
	UserId     string `gorm:"-" json:"userId"`
}

type UserIdAndName struct {
	UserId   int64  `json:"userId"`
	UserName string `json:"userName"`
}

func (SummaryDetails) TableName() string {
	return "biz_summary_details"
}
