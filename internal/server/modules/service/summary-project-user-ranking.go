package service

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type SummaryProjectUserRankingService struct {
	SummaryProjectUserRankingRepo *repo.SummaryProjectUserRankingRepo `inject:""`
}

func NewSummaryProjectUserRankingService() *SummaryProjectUserRankingService {
	return &SummaryProjectUserRankingService{}
}

func (s *SummaryProjectUserRankingService) ProjectUserRanking(projectId int64) (summaryProjectUserRanking model.SummaryProjectUserRanking, err error) {
	return

}

// FindByProjectId
func (s *SummaryProjectUserRankingService) FindByProjectId(projectId uint) (summaryProjectUserRanking []model.SummaryProjectUserRanking, err error) {
	return s.SummaryProjectUserRankingRepo.FindByProjectId(projectId)
}
