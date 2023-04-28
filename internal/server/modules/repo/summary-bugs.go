package repo

import (
	"github.com/aaronchen2k/deeptest/internal/server/core/dao"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
	"time"
)

type SummaryBugsRepo struct {
	DB *gorm.DB `inject:""`
}

func NewSummaryBugsRepo() *SummaryBugsRepo {
	db := dao.GetDB()
	return &SummaryBugsRepo{db}
}

func (r *SummaryBugsRepo) Create(bugs model.SummaryBugs) (err error) {
	err = r.DB.Model(&model.SummaryBugs{}).Create(&bugs).Error
	return
}

func (r *SummaryBugsRepo) UpdateColumnsByDate(bugs model.SummaryBugs, startTime string, endTime string) (err error) {
	err = r.DB.Model(&model.SummaryBugs{}).Where("project_id = ? and created_at > ? and created_at < ?", bugs.ProjectId, startTime, endTime).UpdateColumns(&bugs).Error
	return
}

func (r *SummaryBugsRepo) HasDataOfDate(startTime string, endTime string) (ret bool, err error) {
	var count int64
	err = r.DB.Model(&model.SummaryBugs{}).Raw("select count(id) from deeptest.biz_summary_bugs where created_at > ? and created_at < ? AND NOT deleted;", startTime, endTime).Last(&count).Error
	if count == 0 {
		ret = true
	}
	return
}

func (r *SummaryBugsRepo) CountByProjectId(projectId int64) (count int64, err error) {
	var bugsCount int64
	err = r.DB.Model(&model.SummaryBugs{}).Select("count(id)").Where("project_id = ? AND NOT deleted ", projectId).Find(&bugsCount).Error
	count = bugsCount
	return
}

func (r *SummaryBugsRepo) Count() (count int64, err error) {
	err = r.DB.Model(&model.SummaryBugs{}).Select("count(id) ").Where("NOT deleted ").Find(&count).Error
	return
}

func (r *SummaryBugsRepo) FindByProjectIdGroupByBugSeverity(projectId int64) (result []model.SummaryBugsSeverity, err error) {
	err = r.DB.Model(&model.SummaryBugs{}).Select("count(id) as count,bug_severity as severity ").Where("project_id = ? AND NOT deleted ", projectId).Group("bug_severity").Find(&result).Error
	return
}

func (r *SummaryBugsRepo) FindGroupByBugSeverity() (result []model.SummaryBugsSeverity, err error) {
	err = r.DB.Model(&model.SummaryBugs{}).Select("count(id) as count,bug_severity as severity").Where(" NOT deleted ").Group("bug_severity").Find(&result).Error
	return
}

func (r *SummaryBugsRepo) FindProjectIds() (projectIds []int64, err error) {
	err = r.DB.Model(&model.Project{}).Raw("select id from deeptest.biz_project;").Find(&projectIds).Error
	return
}

func (r *SummaryBugsRepo) CheckUpdated(lastUpdateTime *time.Time) (result bool, err error) {
	result = false
	newTime := time.Now()
	err = r.DB.Model(&model.SummaryBugs{}).Raw("select updated_at from  deeptest.biz_summary_bugs order by updated_at desc limit 1").Find(&newTime).Error
	result = newTime.After(*lastUpdateTime)
	return
}
