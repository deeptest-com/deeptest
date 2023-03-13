package repo

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
)

type SummaryProjectUserRankingRepo struct {
	DB *gorm.DB `inject:""`
}

func NewSummaryProjectUserRankingRepo() *SummaryProjectUserRankingRepo {
	return &SummaryProjectUserRankingRepo{}
}

func (r *SummaryProjectUserRankingRepo) FindByProjectId(projectId uint) (summaryProjectUserRanking []model.SummaryProjectUserRanking, err error) {
	err = r.DB.Model(&model.SummaryProjectUserRanking{}).Where("project_id = ?", projectId).Find(&summaryProjectUserRanking).Error
	return
}
