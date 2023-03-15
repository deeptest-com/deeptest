package repo

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
)

type SummaryDetailsRepo struct {
	DB *gorm.DB `inject:""`
}

func NewSummaryDetailsRepo() *SummaryDetailsRepo {
	return &SummaryDetailsRepo{}
}

func (r *SummaryDetailsRepo) CreateByDate(summaryDetails model.SummaryDetails) (err error) {
	err = r.DB.Model(&model.SummaryDetails{}).Create(&summaryDetails).Error
	return
}

func (r *SummaryDetailsRepo) UpdateColumnsByDate(summaryDetails model.SummaryDetails, startTime string, endTime string) (err error) {
	err = r.DB.Model(&model.SummaryDetails{}).Where("project_id = ? and created_at > ? and created_at < ?", summaryDetails.ProjectId, startTime, endTime).UpdateColumns(&summaryDetails).Error
	return
}

func (r *SummaryDetailsRepo) Count() (count int64, err error) {
	err = r.DB.Model(&model.Project{}).Select("count(id) ").Where("NOT deleted AND not disabled").Find(&count).Error
	return
}

func (r *SummaryDetailsRepo) CountByUserId(userId int64) (count int64, err error) {
	err = r.DB.Model(&model.ProjectMember{}).Select("count(distinct project_id)").Where("user_id = ? AND NOT deleted AND not disabled", userId).Find(&count).Error
	return
}

func (r *SummaryDetailsRepo) FindProjectIdsByUserId(userId int64) (projectIds []int64, err error) {
	err = r.DB.Model(&model.ProjectMember{}).Select("distinct project_id").Where("user_id = ? AND NOT deleted AND not disabled", userId).Find(&projectIds).Error
	return
}

func (r *SummaryDetailsRepo) FindByProjectId(projectId int64) (summaryDetail model.SummaryDetails, err error) {
	err = r.DB.Model(&model.SummaryDetails{}).Where("project_id = ? AND NOT deleted AND not disabled", projectId).Last(&summaryDetail).Error
	return
}

func (r *SummaryDetailsRepo) Find() (summaryDetails []model.SummaryDetails, err error) {
	err = r.DB.Model(&model.SummaryDetails{}).Raw("select * from (deeptest.biz_summary_details) where id in (SELECT max(id) FROM deeptest.biz_summary_details group by project_id) AND NOT deleted AND not disabled;").Find(&summaryDetails).Error
	return
}

func (r *SummaryDetailsRepo) FindByProjectIds(projectIds []int64) (summaryDetails []model.SummaryDetails, err error) {
	err = r.DB.Model(&model.SummaryDetails{}).Raw("select * from (deeptest.biz_summary_details) where id in (SELECT max(id) FROM deeptest.biz_summary_details where project_id in ? group by project_id) AND NOT deleted AND not disabled;", projectIds).Find(&summaryDetails).Error
	return
}

func (r *SummaryDetailsRepo) SummaryCard() (summaryCardTotal []model.SummaryCardTotal, err error) {
	err = r.DB.Model(&model.SummaryDetails{}).Raw("select SUM(scenario_total) as scenario_total,sum(interface_total) as interface_total,sum(exec_total) as exec_total,AVG(pass_rate) as pass_rate,AVG(coverage) as coverage from (deeptest.biz_summary_details) where id in (SELECT max(id) FROM deeptest.biz_summary_details where NOT deleted AND not disabled group by project_id);").Find(&summaryCardTotal).Error
	return
}

func (r *SummaryDetailsRepo) SummaryCardByDate(startTime string, endTime string) (summaryDetails []model.SummaryDetails, err error) {
	err = r.DB.Model(&model.SummaryDetails{}).Raw("select SUM(scenario_total) as scenario_total,sum(interface_total) as interface_total,sum(exec_total) as exec_total,AVG(pass_rate) as pass_rate,AVG(coverage) as coverage from (deeptest.biz_summary_details) where id in (SELECT max(id) FROM deeptest.biz_summary_details where created_at > ? and created_at < ? AND NOT deleted AND not disabled group by project_id);", startTime, endTime).Find(&summaryDetails).Error
	return
}

func (r *SummaryDetailsRepo) FindByProjectIdAndDate(startTime string, endTime string, projectId int64) (summaryDetails model.SummaryDetails, err error) {
	err = r.DB.Model(&model.SummaryDetails{}).Raw("select * from (deeptest.biz_summary_details) where id in (SELECT max(id) FROM deeptest.biz_summary_details where project_id = ? and created_at > ? and created_at < ? AND NOT deleted AND not disabled group by project_id);", projectId, startTime, endTime).Find(&summaryDetails).Error
	return
}
