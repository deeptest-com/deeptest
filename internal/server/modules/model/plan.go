package model

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"time"
)

type Plan struct {
	BaseModel

	Version float64 `json:"version" yaml:"version"`
	Name    string  `json:"name" yaml:"name"`
	Desc    string  `json:"desc" yaml:"desc"`

	CategoryId     int               `json:"categoryId"`
	ProjectId      uint              `json:"projectId"`
	SerialNumber   string            `json:"serialNumber"`
	AdminId        uint              `json:"adminId"` //负责人ID
	CreateUserId   uint              `json:"createUserId"`
	UpdateUserId   uint              `json:"updateUserId"`
	Status         consts.TestStatus `json:"status"`
	TestStage      consts.TestStage  `json:"testStage"`
	Scenarios      []Scenario        `gorm:"-" json:"scenarios"`
	Reports        []PlanReport      `gorm:"-" json:"reports"`
	TestPassRate   string            `gorm:"-" json:"testPassRate"`
	AdminName      string            `gorm:"-" json:"adminName"`      //负责人姓名
	UpdateUserName string            `gorm:"-" json:"updateUserName"` //最近更新人姓名
	CurrEnvId      uint              `json:"currEnvId"`
	CreateUserName string            `gorm:"-" json:"createUserName"` //创建人姓名
	ExecTimes      int64             `gorm:"-" json:"execTimes"`      //执行次数
	ExecutorName   string            `gorm:"-" json:"executorName"`   //执行人姓名
	ExecTime       *time.Time        `gorm:"-" json:"execTime"`       //执行时间
	ExecEnv        string            `gorm:"-" json:"execEnv"`        //执行环境
	IsLy           bool              `json:"IsLy"`                    //是否是leyan计划

}

func (Plan) TableName() string {
	return "biz_plan"
}

type RelaPlanScenario struct {
	BaseModel

	PlanId     uint `json:"planId"`
	ScenarioId uint `json:"scenarioId"`

	ServiceId uint `json:"serviceId"`
	ProjectId uint `json:"projectId"`
	Ordr      int  `gorm:"default:0" json:"ordr"`
}

func (RelaPlanScenario) TableName() string {
	return "biz_plan_scenario_r"
}
