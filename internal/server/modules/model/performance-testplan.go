package model

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
)

type PerformanceTestPlan struct {
	BaseModel

	Version float64 `json:"version" yaml:"version"`
	Name    string  `json:"name" yaml:"name"`
	Desc    string  `json:"desc" yaml:"desc"`

	ScenarioId uint `json:"scenarioId"`

	Processor      Processor         `json:"processor" yaml:"processor" gorm:"-"`
	CategoryId     int64             `json:"categoryId"`
	ServeId        uint              `json:"serveId"`
	ProjectId      uint              `json:"projectId"`
	SerialNumber   string            `json:"serialNumber"`
	Status         consts.TestStatus `json:"status"`
	CreateUserId   uint              `json:"createUserId"`
	CreatorName    string            `gorm:"-" json:"creatorName"`
	CreateUserName string            `json:"createUserName"` //用户登录名
	UpdateUserId   uint              `json:"updateUserId"`
	UpdateUserName string            `json:"updateUserName"`
	Priority       string            `json:"priority"`
	Type           consts.TestType   `json:"type"`
	CurrEnvId      uint              `json:"currEnvId"`
}

func (PerformanceTestPlan) TableName() string {
	return "biz_performance_test_plan"
}

type PerformanceRunner struct {
	BaseModel

	AgentId uint `json:"agentId"`
	Weight  uint `json:"weight"`

	Name       string `gorm:"-" json:"name"`
	WebAddress string `gorm:"-" json:"webAddress"`

	SerialNumber string `json:"serialNumber"`
	ScenarioId   uint   `json:"scenarioId"`
	ProjectId    uint   `json:"projectId"`
}

func (PerformanceRunner) TableName() string {
	return "biz_performance_runner"
}

type PerformanceTestPlanDetail struct {
	PerformanceTestPlan
	CategoryName string `json:"categoryName"`
	RefId        uint   `json:"refId"`
}
