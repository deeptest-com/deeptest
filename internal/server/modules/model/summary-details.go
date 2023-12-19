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
	ScenarioTotal  int64   `gorm:"type:bigint" json:"scenarioTotal"`
	InterfaceTotal int64   `gorm:"type:bigint" json:"interfaceTotal"`
	ExecTotal      int64   `gorm:"type:bigint" json:"execTotal"`
	PassRate       float64 `gorm:"type:double" json:"passRate"`
	Coverage       float64 `gorm:"type:double" json:"coverage"`
}

type ProjectIdsGroupByUserId struct {
	UserId     int64  `gorm:"-" json:"userId"`
	ProjectIds string `gorm:"-" json:"projectIds"`
}

type UserIdsGroupByProjectId struct {
	ProjectIds int64  `gorm:"-" json:"projectIds"`
	UserId     string `gorm:"-" json:"userId"`
}

type ProjectsBugCount struct {
	ProjectId int64 `json:"projectId"`
	Count     int64 `json:"count"`
}

type ProjectIdAndId struct {
	ProjectId int64 `json:"projectId"`
	Id        int64 `json:"id"`
}

type ScenarioProjectIdAndId struct {
	ProjectId uint  `json:"projectId"`
	Id        int64 `json:"id"`
}

type ProjectIdAndFloat struct {
	ProjectId int64   `json:"projectId"`
	Coverage  float64 `json:"coverage"`
}

type UserIdAndName struct {
	ProjectId int64  `json:"projectId"`
	UserId    int64  `json:"userId"`
	UserName  string `json:"userName"`
}

type SummaryProjectInfo struct {
	BaseModel
	Name           string `json:"name"`
	Descr          string `json:"descr"`
	Logo           string `json:"logo"`
	ShortName      string `json:"shortName"`
	IncludeExample bool   `json:"includeExample"`
	AdminId        uint   `json:"adminId"`
	AdminName      string `json:"adminName"`
}

type SimplePassRate struct {
	TotalAssertionNum int64 `gorm:"column:totalAssertionNum"`
	PassAssertionNum  int64 `gorm:"column:passAssertionNum"`
	CheckpointPass    int64 `gorm:"column:checkpointPass"`
	CheckpointFail    int64 `gorm:"column:checkpointFail"`
}

type SimplePassRateByProjectId struct {
	ProjectId int64 `json:"projectId"`
	SimplePassRate
}

func (SummaryDetails) TableName() string {
	return "biz_summary_details"
}
