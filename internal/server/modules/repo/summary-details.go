package repo

import (
	"database/sql"
	"github.com/aaronchen2k/deeptest/internal/server/core/dao"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
	"time"
)

type SummaryDetailsRepo struct {
	DB *gorm.DB `inject:""`
}

func NewSummaryDetailsRepo() *SummaryDetailsRepo {
	db := dao.GetDB()
	return &SummaryDetailsRepo{db}
}

func (r *SummaryDetailsRepo) Create(summaryDetails model.SummaryDetails) (err error) {
	err = r.DB.Model(&model.SummaryDetails{}).Create(&summaryDetails).Error
	return
}

func (r *SummaryDetailsRepo) UpdateColumnsByDate(id int64, summaryDetails model.SummaryDetails) (err error) {
	err = r.DB.Model(&model.SummaryDetails{}).Where("id = ? ", id).UpdateColumns(&summaryDetails).Error
	return
}

func (r *SummaryDetailsRepo) Existed(startTime string, endTime string, projectId int64) (id int64, err error) {
	err = r.DB.Model(&model.SummaryDetails{}).Raw("select id from (deeptest.biz_summary_details) where created_at >= ? and created_at < ? AND project_id = ? AND NOT deleted;", startTime, endTime, projectId).Last(&id).Error
	return
}

func (r *SummaryDetailsRepo) Count() (count int64, err error) {
	err = r.DB.Model(&model.Project{}).Select("count(id) ").Where("NOT deleted").Find(&count).Error
	return
}

func (r *SummaryDetailsRepo) CountByUserId(userId int64) (count int64, err error) {
	err = r.DB.Model(&model.ProjectMember{}).Select("count(distinct project_id)").Where("user_id = ? AND NOT deleted", userId).Find(&count).Error
	return
}

func (r *SummaryDetailsRepo) CountUserTotal() (count int64, err error) {
	err = r.DB.Model(&model.ProjectMember{}).Select("count(distinct user_id) ").Where("NOT deleted").Find(&count).Error
	return
}

func (r *SummaryDetailsRepo) CountProjectUserTotal(projectId int64) (count int64, err error) {
	err = r.DB.Model(&model.ProjectMember{}).Select("count(distinct id) ").Where("project_id = ? And NOT deleted", projectId).Find(&count).Error
	return
}

func (r *SummaryDetailsRepo) FindAllProjectInfo() (projectsInfo []model.SummaryProjectInfo, err error) {
	err = r.DB.Model(&model.Project{}).Select("biz_project.id,biz_project.created_at,biz_project.deleted,biz_project.disabled,biz_project.updated_at,biz_project.name,biz_project.descr,biz_project.logo,biz_project.short_name,biz_project.admin_id,biz_project.include_example ,sys_user.name as admin_name ").Joins("left join sys_user on biz_project.admin_id = sys_user.id").Where("biz_project.deleted !=1 ").Order("id desc").Find(&projectsInfo).Error
	return
}

func (r *SummaryDetailsRepo) FindAdminNameByAdminId(adminId int64) (adminName string, err error) {
	err = r.DB.Model(&model.SysUser{}).Select("name").Where("id = ?", adminId).Find(&adminName).Error
	return
}

func (r *SummaryDetailsRepo) FindProjectIdsByUserId(userId int64) (projectIds []int64, err error) {
	err = r.DB.Model(&model.ProjectMember{}).Select("distinct project_id").Where("user_id = ? AND NOT deleted", userId).Find(&projectIds).Order("user_id").Error
	return
}

func (r *SummaryDetailsRepo) FindProjectIds() (ids []int64, err error) {
	err = r.DB.Model(&model.Project{}).Select("id").Where("NOT deleted").Find(&ids).Error
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

func (r *SummaryDetailsRepo) FindAllUserIdAndNameOfProject() (users []model.UserIdAndName, err error) {
	err = r.DB.Model(&model.ProjectMember{}).Raw("select biz_project_member.project_id,sys_user.id as user_id,sys_user.name as user_name from biz_project_member left join sys_user on sys_user.id = biz_project_member.user_id;").Find(&users).Error
	return
}

func (r *SummaryDetailsRepo) FindCreateUserNameByProjectId(projectId int64) (userName string, err error) {
	err = r.DB.Model(&model.ProjectMember{}).Raw("select sys_user.name as user_name from sys_user inner join biz_project_member on sys_user.id = biz_project_member.user_id where project_id = ?", projectId).First(&userName).Error
	return
}

func (r *SummaryDetailsRepo) FindByProjectId(projectId int64) (summaryDetail model.SummaryDetails, err error) {
	err = r.DB.Model(&model.SummaryDetails{}).Raw("select * from (deeptest.biz_summary_details) where id in (SELECT max(id) FROM deeptest.biz_summary_details where project_id = ? group by project_id) AND NOT deleted ;", projectId).Find(&summaryDetail).Error
	return
}

func (r *SummaryDetailsRepo) Find() (summaryDetails []model.SummaryDetails, err error) {
	err = r.DB.Model(&model.SummaryDetails{}).Raw("select * from deeptest.biz_summary_details where id in (SELECT max(id) FROM deeptest.biz_summary_details where NOT deleted group by project_id) order by project_id;").Find(&summaryDetails).Error
	return
}

func (r *SummaryDetailsRepo) FindByProjectIds(projectIds []int64) (summaryDetails []model.SummaryDetails, err error) {
	err = r.DB.Model(&model.SummaryDetails{}).Raw("select * from (deeptest.biz_summary_details) where id in (SELECT max(id) FROM deeptest.biz_summary_details where project_id in ? group by project_id) AND NOT deleted  order by project_id;", projectIds).Find(&summaryDetails).Error
	return
}

func (r *SummaryDetailsRepo) SummaryCard() (summaryCardTotal model.SummaryCardTotal, err error) {
	err = r.DB.Model(&model.SummaryDetails{}).Raw("select SUM(scenario_total) as scenario_total,sum(interface_total) as interface_total,sum(exec_total) as exec_total,cast(AVG(NULLIF(pass_rate, 0)) as decimal(64,1)) as pass_rate,cast(AVG(NULLIF(coverage, 0)) as decimal(64,1)) as coverage from (deeptest.biz_summary_details) where id in (SELECT max(id) FROM deeptest.biz_summary_details where NOT deleted  group by project_id);").Find(&summaryCardTotal).Error
	return
}

func (r *SummaryDetailsRepo) SummaryCardByDate(startTime string, endTime string) (summaryCardTotal model.SummaryCardTotal, err error) {
	err = r.DB.Model(&model.SummaryDetails{}).Raw("select SUM(scenario_total) as scenario_total,sum(interface_total) as interface_total,sum(exec_total) as exec_total,cast(AVG(NULLIF(pass_rate, 0)) as decimal(64,1)) as pass_rate,cast(AVG(NULLIF(coverage, 0)) as decimal(64,1)) as coverage from (deeptest.biz_summary_details) where id in (SELECT max(id) FROM deeptest.biz_summary_details where created_at >= ? and created_at < ? AND NOT deleted  group by project_id);", startTime, endTime).Find(&summaryCardTotal).Error
	return
}

func (r *SummaryDetailsRepo) SummaryCardByProjectId(projectId int64) (summaryCardTotal model.SummaryCardTotal, err error) {
	err = r.DB.Model(&model.SummaryDetails{}).Raw("select SUM(scenario_total) as scenario_total,sum(interface_total) as interface_total,sum(exec_total) as exec_total,cast(AVG(NULLIF(pass_rate, 0)) as decimal(64,1)) as pass_rate,cast(AVG(NULLIF(coverage, 0)) as decimal(64,1)) as coverage from (deeptest.biz_summary_details) where id in (SELECT max(id) FROM deeptest.biz_summary_details where project_id = ? And NOT deleted  group by project_id);", projectId).Find(&summaryCardTotal).Error
	return
}

func (r *SummaryDetailsRepo) SummaryCardByDateAndProjectId(startTime string, endTime string, projectId int64) (summaryCardTotal model.SummaryCardTotal, err error) {
	err = r.DB.Model(&model.SummaryDetails{}).Raw("select SUM(scenario_total) as scenario_total,sum(interface_total) as interface_total,sum(exec_total) as exec_total,cast(AVG(NULLIF(pass_rate, 0)) as decimal(64,1)) as pass_rate,cast(AVG(NULLIF(coverage, 0)) as decimal(64,1)) as coverage from (deeptest.biz_summary_details) where id in (SELECT max(id) FROM deeptest.biz_summary_details where created_at >= ? and created_at < ? and project_id = ? AND NOT deleted  group by project_id);", startTime, endTime, projectId).Find(&summaryCardTotal).Error
	return
}

func (r *SummaryDetailsRepo) FindByProjectIdAndDate(startTime string, endTime string, projectId int64) (summaryDetails model.SummaryDetails, err error) {
	err = r.DB.Model(&model.SummaryDetails{}).Raw("select * from (deeptest.biz_summary_details) where id in (SELECT max(id) FROM deeptest.biz_summary_details where project_id = ? and created_at > ? and created_at < ? AND NOT deleted group by project_id);", projectId, startTime, endTime).Find(&summaryDetails).Error
	return
}

func (r *SummaryDetailsRepo) FindPassRateByProjectId(projectId int64) (float64, error) {
	var passRate sql.NullFloat64
	err := r.DB.Model(&model.ScenarioReport{}).Raw("select (SUM(pass_assertion_num)/SUM(total_assertion_num))*100 from (deeptest.biz_scenario_report) where project_id = ?;", projectId).Find(&passRate).Error
	return passRate.Float64, err
}

func (r *SummaryDetailsRepo) FindAllPassRateByProjectId() (passRate []model.ProjectIdAndFloat, err error) {
	err = r.DB.Model(&model.ScenarioReport{}).Raw("select biz_scenario_report.project_id,(SUM(biz_scenario_report.pass_assertion_num)/SUM(biz_scenario_report.total_assertion_num))*100 as coverage from (deeptest.biz_scenario_report) group by project_id;").Find(&passRate).Error
	return
}

func (r *SummaryDetailsRepo) CountBugsGroupByProjectId() (bugsCount []model.ProjectsBugCount, err error) {
	err = r.DB.Model(&model.SummaryBugs{}).Select("project_id,count(id) as count").Where("NOT deleted ").Group("project_id").Find(&bugsCount).Error
	return
}

func (r *SummaryDetailsRepo) CountScenarioTotalProjectId(projectId int64) (int64, error) {
	var count sql.NullInt64
	err := r.DB.Model(&model.Scenario{}).Select("count(id)").Where("project_id = ? AND NOT deleted ", projectId).Find(&count).Error
	return count.Int64, err
}

func (r *SummaryDetailsRepo) CountAllScenarioTotalProjectId() (counts []model.ScenarioProjectIdAndId, err error) {
	err = r.DB.Model(&model.Scenario{}).Select("count(id) as id,project_id").Where("NOT deleted ").Group("project_id").Find(&counts).Error
	return
}

func (r *SummaryDetailsRepo) CountExecTotalProjectId(projectId int64) (int64, error) {
	var count sql.NullInt64
	err := r.DB.Model(&model.ScenarioReport{}).Select("count(id)").Where("project_id = ? AND NOT deleted ", projectId).Find(&count).Error
	return count.Int64, err
}

func (r *SummaryDetailsRepo) CountAllExecTotalProjectId() (counts []model.ProjectIdAndId, err error) {
	err = r.DB.Model(&model.ScenarioReport{}).Select("count(id) as id,project_id").Where("NOT deleted ").Group("project_id").Find(&counts).Error
	return
}

func (r *SummaryDetailsRepo) CountEndpointInterfaceTotalProjectId(projectId int64) (int64, error) {
	var count sql.NullInt64
	err := r.DB.Model(&model.EndpointInterface{}).Select("count(id)").Where("project_id = ? AND NOT deleted ", projectId).Find(&count).Error
	return count.Int64, err
}

func (r *SummaryDetailsRepo) CountAllEndpointInterfaceTotalProjectId() (counts []model.ProjectIdAndId, err error) {
	err = r.DB.Model(&model.EndpointInterface{}).Select("count(id) as id,project_id").Where("NOT deleted ").Group("project_id").Find(&counts).Error
	return
}

func (r *SummaryDetailsRepo) FindEndpointIdsByProjectId(projectId int64) (ids []int64, err error) {
	err = r.DB.Model(&model.EndpointInterface{}).Select("id").Where("project_id = ? AND NOT deleted ", projectId).Find(&ids).Error
	return
}

func (r *SummaryDetailsRepo) FindAllEndpointIdsGroupByProjectId() (ids []model.ProjectIdAndId, err error) {
	err = r.DB.Model(&model.EndpointInterface{}).Select("id,project_id").Where("NOT deleted ").Find(&ids).Error
	return
}

func (r *SummaryDetailsRepo) FindAllExecLogProcessorInterfaceTotalGroupByProjectId() (counts []model.ProjectIdAndId, err error) {
	err = r.DB.Model(&model.ExecLogProcessor{}).Raw("select project_id,count(DISTINCT biz_exec_log_processor.interface_id) as id  from deeptest.biz_scenario_report join biz_exec_log_processor on biz_scenario_report.id = biz_exec_log_processor.report_id where processor_category='processor_interface' AND interface_id != 0 group by project_id;").Find(&counts).Error
	return
}

func (r *SummaryDetailsRepo) CheckCardUpdated(lastUpdateTime *time.Time) (result bool, err error) {
	var newCardUpdatedTime *time.Time
	err = r.DB.Model(&model.SummaryDetails{}).Select("updated_at").Order("updated_at desc").Limit(1).Find(&newCardUpdatedTime).Error
	result = newCardUpdatedTime.After(*lastUpdateTime)
	return
}

func (r *SummaryDetailsRepo) CheckDetailsUpdated(lastUpdateTime *time.Time) (result bool, err error) {
	result = false
	newTime := time.Now()
	tables := []string{
		"deeptest.biz_project",
		"deeptest.biz_summary_bugs",
		"deeptest.biz_project_member",
		"deeptest.sys_user",
		"deeptest.biz_scenario_report",
		"deeptest.biz_interface"}

	for _, table := range tables {
		sql := "select updated_at from " + table + " order by updated_at limit 1;"
		err = r.DB.Raw(sql).Find(&newTime).Error

		if newTime.After(*lastUpdateTime) != false {
			result = true
			return
		}
	}
	return
}
