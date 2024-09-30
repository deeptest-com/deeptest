package repo

import (
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
	"time"
)

type SummaryBugsRepo struct {
	*BaseRepo `inject:""`
	DB        *gorm.DB `inject:""`
}

func (r *SummaryBugsRepo) Create(tenantId consts.TenantId, bugs model.SummaryBugs) (err error) {
	err = r.GetDB(tenantId).Model(&model.SummaryBugs{}).Create(&bugs).Error
	return
}

func (r *SummaryBugsRepo) UpdateColumnsByDate(tenantId consts.TenantId, bugs model.SummaryBugs, id int64) (err error) {
	err = r.GetDB(tenantId).Model(&model.SummaryBugs{}).Where("id = ?", id).UpdateColumns(&bugs).Error
	return
}

func (r *SummaryBugsRepo) Existed(tenantId consts.TenantId, bugId int64, projectId int64) (id int64, err error) {

	err = r.GetDB(tenantId).Model(&model.SummaryBugs{}).Raw("select id from biz_summary_bugs where bugId = ? and project_id = ? AND NOT deleted;", bugId, projectId).Last(&id).Error

	return
}

func (r *SummaryBugsRepo) CountByProjectId(tenantId consts.TenantId, projectId int64) (count int64, err error) {
	var bugsCount int64
	err = r.GetDB(tenantId).Model(&model.SummaryBugs{}).Select("count(id)").Where("project_id = ? AND NOT deleted ", projectId).Find(&bugsCount).Error
	count = bugsCount
	return
}

func (r *SummaryBugsRepo) Count(tenantId consts.TenantId) (count int64, err error) {
	err = r.GetDB(tenantId).Model(&model.SummaryBugs{}).Select("count(id) ").Where("NOT deleted ").Find(&count).Error
	return
}

func (r *SummaryBugsRepo) FindByProjectIdGroupByBugSeverity(tenantId consts.TenantId, projectId int64) (result []model.SummaryBugsSeverity, err error) {
	err = r.GetDB(tenantId).Model(&model.SummaryBugs{}).Select("count(id) as count,bug_severity as severity ").Where("project_id = ? AND NOT deleted ", projectId).Group("bug_severity").Find(&result).Error
	return
}

func (r *SummaryBugsRepo) FindGroupByBugSeverity(tenantId consts.TenantId) (result []model.SummaryBugsSeverity, err error) {
	err = r.GetDB(tenantId).Model(&model.SummaryBugs{}).Select("count(id) as count,bug_severity as severity").Where(" NOT deleted ").Group("bug_severity").Find(&result).Error
	return
}

func (r *SummaryBugsRepo) FindProjectIds(tenantId consts.TenantId) (projectIds []int64, err error) {
	err = r.GetDB(tenantId).Model(&model.Project{}).Raw("select id from biz_project;").Find(&projectIds).Error
	return
}

func (r *SummaryBugsRepo) CheckUpdated(tenantId consts.TenantId, lastUpdateTime *time.Time) (result bool, err error) {
	result = false
	newTime := time.Now()
	err = r.GetDB(tenantId).Model(&model.SummaryBugs{}).Raw("select updated_at from biz_summary_bugs order by updated_at desc limit 1").Find(&newTime).Error
	result = newTime.After(*lastUpdateTime)
	return
}

func (r *SummaryBugsRepo) GetNewBugs(tenantId consts.TenantId) (reports []model.ScenarioReport, err error) {
	err = r.GetDB(tenantId).Model(model.ScenarioReport{}).Where("bug_id != '' and id not in (?)", r.GetDB(tenantId).Model(&model.SummaryBugs{}).Select("bug_id")).Find(&reports).Error
	return
}

func (r *SummaryBugsRepo) Creates(tenantId consts.TenantId, bugs []model.SummaryBugs) (err error) {
	err = r.GetDB(tenantId).Create(&bugs).Error
	return
}
