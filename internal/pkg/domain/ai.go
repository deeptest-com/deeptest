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

/** 不同于AiMetrics的Model对象，Domain对象及其延伸的类均携带所有属性 */

type AiMetrics struct {
	BaseObj
	AiMetricsBase
}

type AiMetricsInterface interface {
	Run() (err error)
}
