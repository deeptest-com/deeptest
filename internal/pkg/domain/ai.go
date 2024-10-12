package domain

type ToolModel struct {
	BaseObj
	ToolModelBase
}

type AiMeasurement struct {
	BaseObj
	AiMeasurementBase
}

type AiTemplate struct {
	BaseObj
	AiTemplateBase
}

type AiMetrics struct {
	BaseObj
	AiMetricsBase
	AiMetricsEntityBase
}
