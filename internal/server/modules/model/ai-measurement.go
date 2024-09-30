package model

import "github.com/aaronchen2k/deeptest/internal/pkg/domain"

type AiMeasurement struct {
	BaseModel

	domain.AiMeasurement

	MetricsIds string `json:"metricsIds"`
}

func (AiMeasurement) TableName() string {
	return "ai_measurement"
}
