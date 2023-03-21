package repo

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
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
	err = r.DB.Model(&model.SummaryDetails{}).Select("count(distinct project_id) ").Where("NOT deleted").Find(&count).Error
	return
}

func (r *SummaryDetailsRepo) CountByUserId(userId int64) (count int64, err error) {
	err = r.DB.Model(&model.ProjectMember{}).Select("count(distinct project_id)").Where("user_id = ? AND NOT deleted", userId).Find(&count).Error
	return
}

func (r *SummaryDetailsRepo) FindProjectIdsByUserId(userId int64) (projectIds []int64, err error) {
	err = r.DB.Model(&model.ProjectMember{}).Select("distinct project_id").Where("user_id = ? AND NOT deleted", userId).Find(&projectIds).Error
	return
}

func (r *SummaryDetailsRepo) FindUserIdsByProjectId(projectId int64) (userIds []int64, err error) {
	err = r.DB.Model(&model.ProjectMember{}).Select("distinct user_id").Where("project_id = ? AND NOT deleted", projectId).Find(&userIds).Error
	return
}

func (r *SummaryDetailsRepo) FindProjectIdsGroupByUserId() (projectIdsGroupByUserId model.ProjectIdsGroupByUserId, err error) {
	err = r.DB.Model(&model.ProjectMember{}).Select("user_id,group_concat(DISTINCT project_id)").Group("user_id").Find(&projectIdsGroupByUserId).Error
	return
}

func (r *SummaryDetailsRepo) FindUserIdsGroupByProjectId() (userIdsGroupByProjectId model.UserIdsGroupByProjectId, err error) {
	err = r.DB.Model(&model.ProjectMember{}).Select("project_id,group_concat(DISTINCT user_id)").Group("project_id").Find(&userIdsGroupByProjectId).Error
	return
}

func (r *SummaryDetailsRepo) FindUserIdAndNameByProjectId(projectId int64) (userIdAndName []v1.ResUserIdAndName, err error) {
	err = r.DB.Model(&model.ProjectMember{}).Raw("select sys_user.id as user_id,sys_user.name as user_name from sys_user inner join biz_project_member on sys_user.id = biz_project_member.user_id where project_id = ?", projectId).Find(&userIdAndName).Error
	return
}

func (r *SummaryDetailsRepo) FindByProjectId(projectId int64) (summaryDetail model.SummaryDetails, err error) {
	err = r.DB.Model(&model.SummaryDetails{}).Where("project_id = ? AND NOT deleted", projectId).Last(&summaryDetail).Error
	return
}

func (r *SummaryDetailsRepo) Find() (summaryDetails []model.SummaryDetails, err error) {
	err = r.DB.Model(&model.SummaryDetails{}).Raw("select * from (deeptest.biz_summary_details) where id in (SELECT max(id) FROM deeptest.biz_summary_details group by project_id) AND NOT deleted;").Find(&summaryDetails).Error
	return
}

func (r *SummaryDetailsRepo) FindByProjectIds(projectIds []int64) (summaryDetails []model.SummaryDetails, err error) {
	err = r.DB.Model(&model.SummaryDetails{}).Raw("select * from (deeptest.biz_summary_details) where id in (SELECT max(id) FROM deeptest.biz_summary_details where project_id in ? group by project_id) AND NOT deleted ;", projectIds).Find(&summaryDetails).Error
	return
}

func (r *SummaryDetailsRepo) SummaryCard() (summaryCardTotal model.SummaryCardTotal, err error) {
	err = r.DB.Model(&model.SummaryDetails{}).Raw("select SUM(scenario_total) as scenario_total,sum(interface_total) as interface_total,sum(exec_total) as exec_total,cast(AVG(pass_rate) as decimal(64,1)) as pass_rate,cast(AVG(coverage) as decimal(64,1)) as coverage from (deeptest.biz_summary_details) where id in (SELECT max(id) FROM deeptest.biz_summary_details where NOT deleted  group by project_id);").Find(&summaryCardTotal).Error
	return
}

func (r *SummaryDetailsRepo) SummaryCardByDate(startTime string, endTime string) (summaryCardTotal model.SummaryCardTotal, err error) {
	err = r.DB.Model(&model.SummaryDetails{}).Raw("select SUM(scenario_total) as scenario_total,sum(interface_total) as interface_total,sum(exec_total) as exec_total,cast(AVG(pass_rate) as decimal(64,1)) as pass_rate,cast(AVG(coverage) as decimal(64,1)) as coverage from (deeptest.biz_summary_details) where id in (SELECT max(id) FROM deeptest.biz_summary_details where created_at > ? and created_at < ? AND NOT deleted  group by project_id);", startTime, endTime).Find(&summaryCardTotal).Error
	return
}

func (r *SummaryDetailsRepo) FindByProjectIdAndDate(startTime string, endTime string, projectId int64) (summaryDetails model.SummaryDetails, err error) {
	err = r.DB.Model(&model.SummaryDetails{}).Raw("select * from (deeptest.biz_summary_details) where id in (SELECT max(id) FROM deeptest.biz_summary_details where project_id = ? and created_at > ? and created_at < ? AND NOT deleted group by project_id);", projectId, startTime, endTime).Find(&summaryDetails).Error
	return
}

func (r *SummaryDetailsRepo) CountBugsByProjectId(projectId int64) (count int64, err error) {
	err = r.DB.Model(&model.SummaryBugs{}).Select("count(id)").Where("project_id = ? AND NOT deleted AND not disabled", projectId).Find(&count).Error
	return
}
