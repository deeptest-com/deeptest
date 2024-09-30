package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type AiMetricsService struct {
	AiMeasurementRepo *repo.AiMeasurementRepo `inject:""`
	AiMetricsRepo     *repo.AiMetricsRepo     `inject:""`

	AiMetricsResultRelevancyService *AiMetricsResultRelevancyService `inject:""`
}

func (s *AiMetricsService) LoadForExec(req v1.AiMeasurementExecReq) (cs domain.AiMeasurement, metricsArr []domain.AiMetricsAnswerRelevancy, err error) {
	cs, metricsArr, err = s.AiMeasurementRepo.LoadForExec(req.ID)

	return
}
