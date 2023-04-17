package model

type Debug struct {
	BaseModel
	InvocationBase

	ServeId uint `json:"serveId"`

	ProcessorId uint `json:"processorId"`
	ScenarioId  uint `json:"scenarioId"`
}

func (Debug) TableName() string {
	return "biz_debug"
}
