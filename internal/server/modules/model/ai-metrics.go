package model

import (
	"github.com/deeptest-com/deeptest/internal/pkg/domain"
)

type AiMetricsSummarization struct {
	BaseModel
	domain.AiMetricsSummarizationBase
}

func (AiMetricsSummarization) TableName() string {
	return "ai_metrics_summarization"
}

type AiMetricsAnswerRelevancy struct {
	BaseModel
	domain.AiMetricsAnswerRelevancyBase
}

func (AiMetricsAnswerRelevancy) TableName() string {
	return "ai_metrics_answer_relevancy"
}

type AiMetricsFaithfulness struct {
	BaseModel
	domain.AiMetricsFaithfulnessBase
}

func (AiMetricsFaithfulness) TableName() string {
	return "ai_metrics_faithfulness"
}

type AiMetricsContextualPrecision struct {
	BaseModel
	domain.AiMetricsContextualPrecisionBase
}

func (AiMetricsContextualPrecision) TableName() string {
	return "ai_metrics_contextual_precision"
}

type AiMetricsContextualRecall struct {
	BaseModel
	domain.AiMetricsContextualRecallBase
}

func (AiMetricsContextualRecall) TableName() string {
	return "ai_metrics_contextual_recall"
}

type AiMetricsContextualRelevancy struct {
	BaseModel
	domain.AiMetricsContextualRelevancyBase
}

func (AiMetricsContextualRelevancy) TableName() string {
	return "ai_metrics_contextual_relevancy"
}

type AiMetricsHallucination struct {
	BaseModel
	domain.AiMetricsHallucinationBase
}

func (AiMetricsHallucination) TableName() string {
	return "ai_metrics_hallucination"
}

type AiMetricsBias struct {
	BaseModel
	domain.AiMetricsBiasBase
}

func (AiMetricsBias) TableName() string {
	return "ai_metrics_bias"
}

type AiMetricsToxicity struct {
	BaseModel
	domain.AiMetricsToxicityBase
}

func (AiMetricsToxicity) TableName() string {
	return "ai_metrics_toxicity"
}

type AiMetrics struct {
	BaseModel
	domain.AiMetricsBase
}

func (AiMetrics) TableName() string {
	return "ai_metrics"
}
