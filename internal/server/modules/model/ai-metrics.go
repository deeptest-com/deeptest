package model

import (
	"github.com/deeptest-com/deeptest/internal/pkg/domain"
)

type AiMetricsAnswerRelevancy struct {
	BaseModel
	domain.AiMetricsAnswerRelevancyBase
}

func (AiMetricsAnswerRelevancy) TableName() string {
	return "ai_metrics_answer_relevancy"
}

type AiMetrics struct {
	BaseModel
	domain.AiMetricsBase
}

func (AiMetrics) TableName() string {
	return "ai_metrics"
}
