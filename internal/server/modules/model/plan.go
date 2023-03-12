package model

type Plan struct {
	BaseModel

	Version float64 `json:"version" yaml:"version"`
	Name    string  `json:"name" yaml:"name"`
	Desc    string  `json:"desc" yaml:"desc"`

	CategoryId uint `json:"categoryId"`
	ProjectId  uint `json:"projectId"`

	Scenarios []Scenario `gorm:"-" json:"scenarios"`
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
}

func (RelaPlanScenario) TableName() string {
	return "biz_plan_scenario_r"
}
