package repo

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
	"time"
)

type SummaryProjectUserRankingRepo struct {
	DB *gorm.DB `inject:""`
}

func NewSummaryProjectUserRankingRepo() *SummaryProjectUserRankingRepo {
	return &SummaryProjectUserRankingRepo{}
}

func (r *SummaryProjectUserRankingRepo) Create(summaryProjectUserRanking model.SummaryProjectUserRanking) (err error) {
	err = r.DB.Model(&model.SummaryProjectUserRanking{}).Create(&summaryProjectUserRanking).Error
	return
}

func (r *SummaryProjectUserRankingRepo) UpdateColumnsByDate(summaryProjectUserRanking model.SummaryProjectUserRanking, startTime string, endTime string) (err error) {
	err = r.DB.Model(&model.SummaryProjectUserRanking{}).Where("project_id = ? and created_at > ? and created_at < ?", summaryProjectUserRanking.ProjectId, startTime, endTime).UpdateColumns(&summaryProjectUserRanking).Error
	return
}

func (r *SummaryProjectUserRankingRepo) LastByDate(startTime string, endTime string) (ret bool, err error) {
	var count int64
	err = r.DB.Model(&model.SummaryProjectUserRanking{}).Raw("select count(id) from (deeptest.biz_summary_project_user_ranking) where created_at > ? and created_at < ? AND NOT deleted);", startTime, endTime).Last(&count).Error
	if count == 0 {
		ret = true
	}
	return
}

func (r *SummaryProjectUserRankingRepo) FindByDateAndProjectId(startTime string, endTime string, projectId int64) (summaryProjectUserRanking []model.SummaryProjectUserRanking, err error) {
	err = r.DB.Model(&model.SummaryProjectUserRanking{}).Raw("select scenario_total,testcases_total,updated_at,user_name,sort,user_id,project_id from deeptest.biz_summary_project_user_ranking where id in (SELECT max(id) FROM deeptest.biz_summary_project_user_ranking where created_at > ? and created_at < ? AND NOT deleted And project_id = ? group by user_id ORDER BY sort asc);", startTime, endTime, projectId).Find(&summaryProjectUserRanking).Error
	return
}

func (r *SummaryProjectUserRankingRepo) FindByProjectId(projectId int64) (summaryProjectUserRanking []model.SummaryProjectUserRanking, err error) {
	err = r.DB.Model(&model.SummaryProjectUserRanking{}).Raw("select scenario_total,testcases_total,updated_at,user_name,sort,user_id,project_id from deeptest.biz_summary_project_user_ranking where id in (SELECT max(id) FROM deeptest.biz_summary_project_user_ranking where NOT deleted And project_id = ? group by user_id ORDER BY sort asc);", projectId).Find(&summaryProjectUserRanking).Error
	return
}

func (r *SummaryProjectUserRankingRepo) FindGroupByProjectId(projectId int64) (summaryProjectUserRanking []model.SummaryProjectUserRanking, err error) {
	err = r.DB.Model(&model.SummaryProjectUserRanking{}).Raw("select scenario_total,testcases_total,updated_at,user_name,sort,user_id,project_id from deeptest.biz_summary_project_user_ranking where id in (SELECT max(id) FROM deeptest.biz_summary_project_user_ranking where NOT deleted group by user_id ORDER BY sort asc);").Find(&summaryProjectUserRanking).Error
	return
}

func (r *SummaryProjectUserRankingRepo) CheckUpdated(oldTime *time.Time) (result bool, err error) {
	var newTime *time.Time
	err = r.DB.Model(&model.SummaryProjectUserRanking{}).Select("updated_at").Order("updated_at desc").Limit(1).Find(&newTime).Error
	result = newTime.After(*oldTime)
	return
}
