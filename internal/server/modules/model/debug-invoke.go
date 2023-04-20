package model

type DebugInvoke struct {
	BaseModel
	InvocationBase

	ServeId uint `json:"serveId"`

	ProcessorId uint `json:"processorId"`
	ScenarioId  uint `json:"scenarioId"`
}

func (DebugInvoke) TableName() string {
	return "biz_debug_invoke"
}
