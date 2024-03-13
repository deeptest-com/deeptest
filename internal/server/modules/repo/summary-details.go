package repo

import (
	"database/sql"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
	"time"
)

type SummaryDetailsRepo struct {
	*BaseRepo `inject:""`
	DB        *gorm.DB `inject:""`
}

func (r *SummaryDetailsRepo) Create(tenantId consts.TenantId, summaryDetails model.SummaryDetails) (err error) {
	err = r.GetDB(tenantId).Model(&model.SummaryDetails{}).Create(&summaryDetails).Error
	return
}

func (r *SummaryDetailsRepo) UpdateColumnsByDate(tenantId consts.TenantId, id int64, summaryDetails model.SummaryDetails) (err error) {
	now := time.Now()
	summaryDetails.UpdatedAt = &now
	err = r.GetDB(tenantId).Model(&model.SummaryDetails{}).Where("id = ? and not deleted", id).UpdateColumns(&summaryDetails).Error
	return
}

func (r *SummaryDetailsRepo) Existed(tenantId consts.TenantId, startTime string, endTime string, projectId int64) (id int64, err error) {
	err = r.GetDB(tenantId).Model(&model.SummaryDetails{}).Raw("select id from biz_summary_details where created_at >= ? and created_at < ? AND project_id = ? AND NOT deleted;", startTime, endTime, projectId).Last(&id).Error
	return
}

func (r *SummaryDetailsRepo) Count(tenantId consts.TenantId) (count int64, err error) {
	err = r.GetDB(tenantId).Model(&model.Project{}).Select("count(id) ").Where("NOT deleted").Find(&count).Error
	return
}

func (r *SummaryDetailsRepo) CountByUserId(tenantId consts.TenantId, userId int64) (count int64, err error) {
	err = r.GetDB(tenantId).Model(&model.ProjectMember{}).Select("count(distinct project_id)").Where("user_id = ? AND NOT deleted", userId).Find(&count).Error
	return
}

func (r *SummaryDetailsRepo) CountUserTotal(tenantId consts.TenantId) (count int64, err error) {
	err = r.GetDB(tenantId).Model(&model.ProjectMember{}).Select("count(distinct user_id) ").Where("NOT deleted").Find(&count).Error
	return
}

func (r *SummaryDetailsRepo) CountProjectUserTotal(tenantId consts.TenantId, projectId int64) (count int64, err error) {
	err = r.GetDB(tenantId).Model(&model.ProjectMember{}).Select("count(distinct id) ").Where("project_id = ? And NOT deleted", projectId).Find(&count).Error
	return
}

func (r *SummaryDetailsRepo) FindAllProjectInfo(tenantId consts.TenantId) (projectsInfo []model.SummaryProjectInfo, err error) {
	err = r.GetDB(tenantId).Model(&model.Project{}).Select("biz_project.id,biz_project.created_at,biz_project.deleted,biz_project.disabled,biz_project.updated_at,biz_project.name,biz_project.descr,biz_project.logo,biz_project.short_name,biz_project.admin_id,biz_project.include_example ,sys_user.name as admin_name ").Joins("left join sys_user on biz_project.admin_id = sys_user.id").Where("biz_project.deleted !=1").Order("id desc").Find(&projectsInfo).Error
	return
}

func (r *SummaryDetailsRepo) FindAdminNameByAdminId(tenantId consts.TenantId, adminId int64) (adminName string, err error) {
	err = r.GetDB(tenantId).Model(&model.SysUser{}).Select("name").Where("id = ? and not deleted", adminId).Find(&adminName).Error
	return
}

func (r *SummaryDetailsRepo) FindProjectIdsByUserId(tenantId consts.TenantId, userId int64) (projectIds []int64, err error) {
	err = r.GetDB(tenantId).Model(&model.ProjectMember{}).Select("distinct project_id").Where("user_id = ? AND NOT deleted", userId).Find(&projectIds).Order("user_id").Error
	return
}

func (r *SummaryDetailsRepo) FindProjectIds(tenantId consts.TenantId) (ids []int64, err error) {
	err = r.GetDB(tenantId).Model(&model.Project{}).Select("id").Where("NOT deleted").Find(&ids).Error
	return
}

func (r *SummaryDetailsRepo) FindUserIdsByProjectId(tenantId consts.TenantId, projectId int64) (userIds []int64, err error) {
	err = r.GetDB(tenantId).Model(&model.ProjectMember{}).Select("distinct user_id").Where("project_id = ? AND NOT deleted", projectId).Find(&userIds).Error
	return
}

func (r *SummaryDetailsRepo) FindProjectIdsGroupByUserId(tenantId consts.TenantId) (projectIdsGroupByUserId model.ProjectIdsGroupByUserId, err error) {
	err = r.GetDB(tenantId).Model(&model.ProjectMember{}).Select("user_id,group_concat(DISTINCT project_id)").Where("NOT deleted ").Group("user_id").Find(&projectIdsGroupByUserId).Error
	return
}

func (r *SummaryDetailsRepo) FindUserIdsGroupByProjectId(tenantId consts.TenantId) (userIdsGroupByProjectId model.UserIdsGroupByProjectId, err error) {
	err = r.GetDB(tenantId).Model(&model.ProjectMember{}).Select("project_id,group_concat(DISTINCT user_id)").Where("NOT deleted ").Group("project_id").Find(&userIdsGroupByProjectId).Error
	return
}

func (r *SummaryDetailsRepo) FindAllUserIdAndNameOfProject(tenantId consts.TenantId) (users []model.UserIdAndName, err error) {
	err = r.GetDB(tenantId).Model(&model.ProjectMember{}).Raw("select biz_project_member.project_id,sys_user.id as user_id,sys_user.name as user_name from biz_project_member left join sys_user on sys_user.id = biz_project_member.user_id where biz_project_member.deleted !=1 and sys_user.deleted !=1;").Find(&users).Error
	return
}

func (r *SummaryDetailsRepo) FindCreateUserNameByProjectId(tenantId consts.TenantId, projectId int64) (userName string, err error) {
	err = r.GetDB(tenantId).Model(&model.ProjectMember{}).Raw("select sys_user.name as user_name from sys_user inner join biz_project_member on sys_user.id = biz_project_member.user_id where project_id = ? and biz_project_member.deleted !=1 and sys_user.deleted !=1", projectId).First(&userName).Error
	return
}

func (r *SummaryDetailsRepo) FindByProjectId(tenantId consts.TenantId, projectId int64) (summaryDetail model.SummaryDetails, err error) {
	err = r.GetDB(tenantId).Model(&model.SummaryDetails{}).Raw("select * from biz_summary_details where id in (SELECT max(id) FROM biz_summary_details where project_id = ? group by project_id) AND NOT deleted ;", projectId).Find(&summaryDetail).Error
	return
}

func (r *SummaryDetailsRepo) Find(tenantId consts.TenantId) (summaryDetails []model.SummaryDetails, err error) {
	err = r.GetDB(tenantId).Model(&model.SummaryDetails{}).Raw("select * from biz_summary_details where id in (SELECT max(id) FROM biz_summary_details where NOT deleted group by project_id) order by project_id;").Find(&summaryDetails).Error
	return
}

func (r *SummaryDetailsRepo) FindByProjectIds(tenantId consts.TenantId, projectIds []int64) (summaryDetails []model.SummaryDetails, err error) {
	err = r.GetDB(tenantId).Model(&model.SummaryDetails{}).Raw("select * from biz_summary_details where id in (SELECT max(id) FROM biz_summary_details where project_id in ? group by project_id) AND NOT deleted  order by project_id;", projectIds).Find(&summaryDetails).Error
	return
}

func (r *SummaryDetailsRepo) SummaryCard(tenantId consts.TenantId) (summaryCardTotal model.SummaryCardTotal, err error) {
	err = r.GetDB(tenantId).Model(&model.SummaryDetails{}).Raw("select SUM(scenario_total) as scenario_total,sum(interface_total) as interface_total,sum(exec_total) as exec_total,cast(AVG(NULLIF(pass_rate, 0)) as decimal(64,1)) as pass_rate,cast(AVG(NULLIF(coverage, 0)) as decimal(64,1)) as coverage from biz_summary_details where id in (SELECT max(id) FROM biz_summary_details where NOT deleted  group by project_id);").Find(&summaryCardTotal).Error
	return
}

func (r *SummaryDetailsRepo) SummaryCardByDate(tenantId consts.TenantId, startTime string, endTime string) (summaryCardTotal model.SummaryCardTotal, err error) {
	err = r.GetDB(tenantId).Model(&model.SummaryDetails{}).Raw("select SUM(scenario_total) as scenario_total,sum(interface_total) as interface_total,sum(exec_total) as exec_total,cast(AVG(NULLIF(pass_rate, 0)) as decimal(64,1)) as pass_rate,cast(AVG(NULLIF(coverage, 0)) as decimal(64,1)) as coverage from biz_summary_details where id in (SELECT min(id) FROM biz_summary_details where created_at >= ? and created_at < ? AND NOT deleted  group by project_id);", startTime, endTime).Find(&summaryCardTotal).Error
	return
}

func (r *SummaryDetailsRepo) SummaryCardByDateAndProjectId(tenantId consts.TenantId, startTime string, endTime string, projectId int64) (summaryCardTotal model.SummaryCardTotal, err error) {
	err = r.GetDB(tenantId).Model(&model.SummaryDetails{}).Raw("select scenario_total,interface_total,exec_total,pass_rate,coverage from biz_summary_details where created_at >= ? and created_at < ? and project_id = ? AND NOT deleted;", startTime, endTime, projectId).First(&summaryCardTotal).Error
	return
}

func (r *SummaryDetailsRepo) FindByProjectIdAndDate(tenantId consts.TenantId, startTime string, endTime string, projectId int64) (summaryDetails model.SummaryDetails, err error) {
	err = r.GetDB(tenantId).Model(&model.SummaryDetails{}).Raw("select * from biz_summary_details where id in (SELECT min(id) FROM biz_summary_details where project_id = ? and created_at > ? and created_at < ? AND NOT deleted group by project_id);", projectId, startTime, endTime).Find(&summaryDetails).Error
	return
}

func (r *SummaryDetailsRepo) FindAssertionCountByProjectId(tenantId consts.TenantId, projectId int64) (result model.SimplePassRate, err error) {
	err = r.GetDB(tenantId).Model(&model.ScenarioReport{}).Raw("select sum(total_assertion_num)as totalAssertionNum,sum(pass_assertion_num) as passAssertionNum,SUM(JSON_EXTRACT(stat_raw, '$.checkpointPass')) AS checkpointPass,  SUM(JSON_EXTRACT(stat_raw, '$.checkpointFail')) AS checkpointFail  from biz_scenario_report where project_id = ? and not deleted;", projectId).Find(&result).Error
	return
}

func (r *SummaryDetailsRepo) FindAllAssertionCount(tenantId consts.TenantId) (result model.SimplePassRate, err error) {
	err = r.GetDB(tenantId).Model(&model.ScenarioReport{}).Raw("select sum(total_assertion_num) as totalAssertionNum,sum(pass_assertion_num) as passAssertionNum,SUM(JSON_EXTRACT(stat_raw, '$.checkpointPass')) AS checkpointPass,  SUM(JSON_EXTRACT(stat_raw, '$.checkpointFail')) AS checkpointFail  from biz_scenario_report where not deleted;").Find(&result).Error
	return
}

func (r *SummaryDetailsRepo) FindAllAssertionCountGroupByProjectId(tenantId consts.TenantId) (result []model.SimplePassRateByProjectId, err error) {
	err = r.GetDB(tenantId).Model(&model.ScenarioReport{}).Raw("select biz_scenario_report.project_id,SUM(total_assertion_num) as totalAssertionNum ,SUM(pass_assertion_num) as passAssertionNum ,SUM(JSON_EXTRACT(stat_raw, '$.checkpointPass')) AS checkpointPass,  SUM(JSON_EXTRACT(stat_raw, '$.checkpointFail')) AS checkpointFail  from biz_scenario_report where  not deleted group by project_id;").Find(&result).Error
	return
}

func (r *SummaryDetailsRepo) CountBugsGroupByProjectId(tenantId consts.TenantId) (bugsCount []model.ProjectsBugCount, err error) {
	err = r.GetDB(tenantId).Model(&model.SummaryBugs{}).Select("project_id,count(id) as count").Where("NOT deleted ").Group("project_id").Find(&bugsCount).Error
	return
}

func (r *SummaryDetailsRepo) CountScenarioTotalProjectId(tenantId consts.TenantId, projectId int64) (int64, error) {
	var count sql.NullInt64
	err := r.GetDB(tenantId).Model(&model.Scenario{}).Select("count(id)").Where("project_id = ? AND NOT deleted ", projectId).Find(&count).Error
	return count.Int64, err
}

func (r *SummaryDetailsRepo) CountAllScenarioTotal(tenantId consts.TenantId) (count int64, err error) {
	err = r.GetDB(tenantId).Model(&model.Scenario{}).Select("count(id)").Where("NOT deleted ").Find(&count).Error
	return
}

func (r *SummaryDetailsRepo) CountAllScenarioTotalProjectId(tenantId consts.TenantId) (counts []model.ScenarioProjectIdAndId, err error) {
	err = r.GetDB(tenantId).Model(&model.Scenario{}).Select("count(id) as id,project_id").Where("NOT deleted ").Group("project_id").Find(&counts).Error
	return
}

func (r *SummaryDetailsRepo) CountExecTotalProjectId(tenantId consts.TenantId, projectId int64) (int64, error) {
	var count sql.NullInt64
	err := r.GetDB(tenantId).Model(&model.ScenarioReport{}).Select("count(id)").Where("project_id = ? AND NOT deleted ", projectId).Find(&count).Error
	return count.Int64, err
}
func (r *SummaryDetailsRepo) CountAllExecTotal(tenantId consts.TenantId) (int64, error) {
	var count sql.NullInt64
	err := r.GetDB(tenantId).Model(&model.ScenarioReport{}).Select("count(id)").Where("NOT deleted ").Find(&count).Error
	return count.Int64, err
}

func (r *SummaryDetailsRepo) CountAllExecTotalProjectId(tenantId consts.TenantId) (counts []model.ProjectIdAndId, err error) {
	err = r.GetDB(tenantId).Model(&model.ScenarioReport{}).Select("count(id) as id,project_id").Where("NOT deleted ").Group("project_id").Find(&counts).Error
	return
}

func (r *SummaryDetailsRepo) CountAllEndpointTotal(tenantId consts.TenantId) (int64, error) {
	var count sql.NullInt64
	err := r.GetDB(tenantId).Model(&model.EndpointInterface{}).Select("count(id)").Where("NOT deleted ").Find(&count).Error
	return count.Int64, err
}

func (r *SummaryDetailsRepo) CountEndpointInterfaceTotalProjectId(tenantId consts.TenantId, projectId int64) (int64, error) {
	var count sql.NullInt64
	err := r.GetDB(tenantId).Model(&model.EndpointInterface{}).Select("count(id)").Where("project_id = ? AND NOT deleted ", projectId).Find(&count).Error
	return count.Int64, err
}

func (r *SummaryDetailsRepo) CountAllEndpointInterfaceTotalProjectId(tenantId consts.TenantId) (counts []model.ProjectIdAndId, err error) {
	err = r.GetDB(tenantId).Model(&model.EndpointInterface{}).Select("count(id) as id,project_id").Where("NOT deleted ").Group("project_id").Find(&counts).Error
	return
}

func (r *SummaryDetailsRepo) FindEndpointIdsByProjectId(tenantId consts.TenantId, projectId int64) (ids []int64, err error) {
	err = r.GetDB(tenantId).Model(&model.EndpointInterface{}).Select("id").Where("project_id = ? AND NOT deleted ", projectId).Find(&ids).Error
	return
}

func (r *SummaryDetailsRepo) FindAllEndpointIdsGroupByProjectId(tenantId consts.TenantId) (ids []model.ProjectIdAndId, err error) {
	err = r.GetDB(tenantId).Model(&model.EndpointInterface{}).Select("id,project_id").Where("NOT deleted ").Find(&ids).Error
	return
}

func (r *SummaryDetailsRepo) FindExecLogProcessorInterfaceTotalGroupByProjectId(tenantId consts.TenantId, projectId int64) (counts int64, err error) {
	err = r.GetDB(tenantId).Model(&model.ExecLogProcessor{}).Raw("select count(DISTINCT biz_exec_log_processor.endpoint_interface_id)  from biz_scenario_report join biz_exec_log_processor on biz_scenario_report.id = biz_exec_log_processor.report_id where processor_category='processor_interface' AND endpoint_interface_id != 0 and biz_scenario_report.deleted !=1  and biz_scenario_report.deleted !=1 and project_id = ?;", projectId).Find(&counts).Error
	return
}

func (r *SummaryDetailsRepo) FindAllExecLogProcessorInterfaceTotal(tenantId consts.TenantId) (counts int64, err error) {
	err = r.GetDB(tenantId).Model(&model.ExecLogProcessor{}).Raw("select count(DISTINCT biz_exec_log_processor.endpoint_interface_id)  from biz_scenario_report join biz_exec_log_processor on biz_scenario_report.id = biz_exec_log_processor.report_id where processor_category='processor_interface' AND endpoint_interface_id != 0 and biz_scenario_report.deleted !=1  and biz_scenario_report.deleted !=1 ;").Find(&counts).Error
	return
}

func (r *SummaryDetailsRepo) FindAllExecLogProcessorInterfaceTotalGroupByProjectId(tenantId consts.TenantId) (counts []model.ProjectIdAndId, err error) {
	err = r.GetDB(tenantId).Model(&model.ExecLogProcessor{}).Raw("select project_id,count(DISTINCT biz_exec_log_processor.endpoint_interface_id) as id  from biz_scenario_report join biz_exec_log_processor on biz_scenario_report.id = biz_exec_log_processor.report_id where processor_category='processor_interface' AND endpoint_interface_id != 0 and biz_scenario_report.deleted !=1  and biz_scenario_report.deleted !=1 group by project_id;").Find(&counts).Error
	return
}

func (r *SummaryDetailsRepo) CheckCardUpdated(tenantId consts.TenantId, lastUpdateTime *time.Time) (result bool, err error) {
	var newCardUpdatedTime *time.Time
	err = r.GetDB(tenantId).Model(&model.SummaryDetails{}).Select("updated_at").Order("updated_at desc").Limit(1).Find(&newCardUpdatedTime).Error
	result = newCardUpdatedTime.After(*lastUpdateTime)
	return
}

func (r *SummaryDetailsRepo) CheckDetailsUpdated(tenantId consts.TenantId, lastUpdateTime *time.Time) (result bool, err error) {
	result = false
	newTime := time.Now()
	tables := []string{
		"biz_project",
		"biz_summary_bugs",
		"biz_project_member",
		"sys_user",
		"biz_scenario_report",
		"biz_interface"}

	for _, table := range tables {
		sql := "select updated_at from " + table + " order by updated_at limit 1;"
		err = r.GetDB(tenantId).Raw(sql).Find(&newTime).Error

		if newTime.After(*lastUpdateTime) != false {
			result = true
			return
		}
	}
	return
}
