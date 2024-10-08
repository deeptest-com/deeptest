package domain

type AiMeasurement struct {
	BaseObj
	AiMeasurementBase
}

type AiTemplate struct {
	BaseObj
	AiTemplateBase
}

type AiModel struct {
	BaseObj
	AiModelBase
}

type AiMetrics struct {
	BaseObj
	AiMetricsBase
	AiMetricsEntityBase
}
