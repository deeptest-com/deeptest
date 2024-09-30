package model

import "github.com/deeptest-com/deeptest/internal/pkg/domain"

type AiMeasurement struct {
	BaseModel

	domain.AiMeasurement

	MetricsIds string `json:"metricsIds"`
}

func (AiMeasurement) TableName() string {
	return "ai_measurement"
}
