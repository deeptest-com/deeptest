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
func (r *SummaryBugsRepo) Create(bugs interface{}) (err error) {
	err = r.DB.Model(&model.SummaryBugs{}).Create(bugs).Error
	return
}

func (r *SummaryBugsRepo) CountByProjectId(projectId int64) (count int64, err error) {
	var bugsCount int64
	err = r.DB.Model(&model.SummaryBugs{}).Select("count(id)").Where("project_id = ? ", projectId).Find(&bugsCount).Error
	count = bugsCount
	return
}

func (r *SummaryBugsRepo) Count() (count int64, err error) {
	var bugsCount int64
	err = r.DB.Model(&model.SummaryBugs{}).Select("count(id)").Find(&bugsCount).Error
	return
}

func (r *SummaryBugsRepo) FindByProjectIdGroupByBugSeverity(projectId int64) (summaryBugs []model.SummaryBugs, err error) {
	var bugsSeveritySummary []struct {
		BugSeverity string `gorm:column:bug_severity`
	}

	err = r.DB.Model(&model.SummaryBugs{}).Select("count(id)").Where("project_id = ? ", projectId).Group("bug_severity").Find(&bugsSeveritySummary).Error
	return
}

func (r *SummaryBugsRepo) FindGroupByBugSeverity() (summaryBugs []model.SummaryBugs, err error) {
	err = r.DB.Model(&model.SummaryBugs{}).Select("count(id)").Group("bug_severity").Find(&summaryBugs).Error
	return
}
