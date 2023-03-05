package model

type Plan struct {
	BaseModel

	Version float64 `json:"version" yaml:"version"`
	Name    string  `json:"name" yaml:"name"`
	Desc    string  `json:"desc" yaml:"desc"`

	CategoryId uint `json:"categoryId"`
	ProjectId  uint `json:"projectId"`
}

func (Plan) TableName() string {
	return "biz_plan"
}

type RelaPlanScenario struct {
	BaseModel

	PlanId     uint `json:"plan_id"`
	ScenarioId uint `json:"scenario_id"`

	ServiceId uint `json:"service_id"`
	ProjectId uint `json:"project_id"`
}

func (RelaPlanScenario) TableName() string {
	return "biz_plan_scenario_r"
}
