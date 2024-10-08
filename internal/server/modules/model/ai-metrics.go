package model

import (
	"github.com/deeptest-com/deeptest/internal/pkg/domain"
)

type AiMetrics struct {
	BaseModel
	domain.AiMetricsBase
	domain.AiMetricsEntityBase
}

func (AiMetrics) TableName() string {
	return "ai_metrics"
}
