package repo

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
)

type SummaryBugsRepo struct {
	DB *gorm.DB `inject:""`
}

func NewSummaryBugsRepo() *SummaryBugsRepo {
	return &SummaryBugsRepo{}
}

func (r *SummaryBugsRepo) CreateByDate(bugs interface{}) (err error) {
	err = r.DB.Model(&model.SummaryBugs{}).Create(bugs).Error
	return
}

func (r *SummaryBugsRepo) CountByProjectId(projectId int64) (count int64, err error) {
	var bugsCount int64
	err = r.DB.Model(&model.SummaryBugs{}).Select("count(id)").Where("project_id = ? AND NOT deleted AND not disabled", projectId).Find(&bugsCount).Error
	count = bugsCount
	return
}

func (r *SummaryBugsRepo) Count() (count int64, err error) {
	err = r.DB.Model(&model.SummaryBugs{}).Select("count(id) ").Where("NOT deleted AND not disabled").Find(&count).Error
	return
}

func (r *SummaryBugsRepo) FindByProjectIdGroupByBugSeverity(projectId int64) (result []model.SummaryBugsSeverity, err error) {
	err = r.DB.Model(&model.SummaryBugs{}).Select("count(id) as count,bug_severity as severity ").Where("project_id = ? AND NOT deleted AND not disabled", projectId).Group("bug_severity").Find(&result).Error
	return
}

func (r *SummaryBugsRepo) FindGroupByBugSeverity() (result []model.SummaryBugsSeverity, err error) {
	err = r.DB.Model(&model.SummaryBugs{}).Select("count(id) as count,bug_severity as severity").Where(" NOT deleted AND not disabled").Group("bug_severity").Find(&result).Error
	return
}
