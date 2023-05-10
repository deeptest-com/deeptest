package model

import "github.com/aaronchen2k/deeptest/internal/pkg/consts"

type Plan struct {
	BaseModel

	Version float64 `json:"version" yaml:"version"`
	Name    string  `json:"name" yaml:"name"`
	Desc    string  `json:"desc" yaml:"desc"`

	CategoryId   uint              `json:"categoryId"`
	ProjectId    uint              `json:"projectId"`
	SerialNumber string            `json:"serialNumber"`
	DirectorId   uint              `json:"directorId"` //负责人ID
	UpdaterId    uint              `json:"updaterId"`
	Status       consts.TestStatus `json:"status"`
	TestStage    consts.TestStage  `json:"testStage"`
	Scenarios    []Scenario        `gorm:"-" json:"scenarios"`
	Reports      []PlanReport      `gorm:"-" json:"reports"`
	TestPassRate string            `gorm:"-" json:"testPassRate"`
	DirectorName string            `gorm:"-" json:"directorName"` //负责人姓名
	UpdaterName  string            `gorm:"-" json:"updaterName"`  //最近更新人姓名
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
	SortId    uint `json:"sortId"` //排序ID
}

func (RelaPlanScenario) TableName() string {
	return "biz_plan_scenario_r"
}
