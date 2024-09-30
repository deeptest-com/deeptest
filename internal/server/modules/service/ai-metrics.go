package service

import (
	agentDomain "github.com/deeptest-com/deeptest/internal/agent/domain"
	"github.com/deeptest-com/deeptest/internal/pkg/domain"
	"github.com/deeptest-com/deeptest/internal/server/modules/repo"
)

type AiMetricsService struct {
	AiMeasurementRepo *repo.AiMeasurementRepo `inject:""`
	AiMetricsRepo     *repo.AiMetricsRepo     `inject:""`

	AiMetricsResultRelevancyService *AiMetricsResultRelevancyService `inject:""`
}

func (s *AiMetricsService) LoadForExec(req agentDomain.AiMeasurementExecReq) (cs domain.AiMeasurement, metricsArr []domain.AiMetricsAnswerRelevancy, err error) {
	cs, metricsArr, err = s.AiMeasurementRepo.LoadForExec(req.AiMeasurement.ID)

	return
}
