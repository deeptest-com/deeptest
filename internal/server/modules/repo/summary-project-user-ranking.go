package repo

import (
	"github.com/aaronchen2k/deeptest/internal/server/core/dao"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
	"time"
)

type SummaryProjectUserRankingRepo struct {
	DB *gorm.DB `inject:""`
}

func NewSummaryProjectUserRankingRepo() *SummaryProjectUserRankingRepo {
	db := dao.GetDB()
	return &SummaryProjectUserRankingRepo{db}
}

func (r *SummaryProjectUserRankingRepo) Create(summaryProjectUserRanking model.SummaryProjectUserRanking) (err error) {
	err = r.DB.Model(&model.SummaryProjectUserRanking{}).Create(&summaryProjectUserRanking).Error
	return
}

func (r *SummaryProjectUserRankingRepo) UpdateColumnsByDate(id int64, summaryProjectUserRanking model.SummaryProjectUserRanking) (err error) {
	err = r.DB.Model(&model.SummaryProjectUserRanking{}).Where("id = ?", id).UpdateColumns(&summaryProjectUserRanking).Error
	return
}

func (r *SummaryProjectUserRankingRepo) Existed(startTime string, endTime string, projectId int64, userId int64) (id int64, err error) {
	err = r.DB.Model(&model.SummaryProjectUserRanking{}).Raw("select id from biz_summary_project_user_ranking where created_at >= ? and created_at < ? AND project_id = ? And user_id = ? And NOT deleted;", startTime, endTime, projectId, userId).Last(&id).Error
	return
}

func (r *SummaryProjectUserRankingRepo) FindProjectIds() (projectIds []int64, err error) {
	err = r.DB.Model(&model.Project{}).Raw("select id from biz_project;").Find(&projectIds).Error
	return
}

func (r *SummaryProjectUserRankingRepo) FindMaxDataByDateAndProjectId(startTime string, endTime string, projectId int64) (summaryProjectUserRanking []model.SummaryProjectUserRanking, err error) {
	err = r.DB.Model(&model.SummaryProjectUserRanking{}).Raw("select scenario_total,test_case_total,updated_at,sort,user_id,project_id from biz_summary_project_user_ranking where id in (SELECT max(id) FROM biz_summary_project_user_ranking where created_at >= ? and created_at < ? AND NOT deleted And project_id = ? group by user_id ORDER BY sort asc);", startTime, endTime, projectId).Find(&summaryProjectUserRanking).Error
	return
}

func (r *SummaryProjectUserRankingRepo) FindMinDataByDateAndProjectId(startTime string, endTime string, projectId int64) (summaryProjectUserRanking []model.SummaryProjectUserRanking, err error) {
	err = r.DB.Model(&model.SummaryProjectUserRanking{}).Raw("select scenario_total,test_case_total,updated_at,sort,user_id,project_id from biz_summary_project_user_ranking where id in (SELECT min(id) FROM biz_summary_project_user_ranking where created_at >= ? and created_at < ? AND NOT deleted And project_id = ? group by user_id ORDER BY sort asc);", startTime, endTime, projectId).Find(&summaryProjectUserRanking).Error
	return
}

func (r *SummaryProjectUserRankingRepo) FindByProjectId(projectId int64) (summaryProjectUserRanking []model.SummaryProjectUserRanking, err error) {
	err = r.DB.Model(&model.SummaryProjectUserRanking{}).Raw("select scenario_total,test_case_total,updated_at,sort,user_id,project_id from biz_summary_project_user_ranking where id in (SELECT max(id) FROM biz_summary_project_user_ranking where NOT deleted And project_id = ? group by user_id ORDER BY sort asc);", projectId).Find(&summaryProjectUserRanking).Error
	return
}

func (r *SummaryProjectUserRankingRepo) FindGroupByProjectId() (summaryProjectUserRanking []model.SummaryProjectUserRanking, err error) {
	err = r.DB.Model(&model.SummaryProjectUserRanking{}).Raw("select scenario_total,test_case_total,updated_at,sort,user_id,project_id from biz_summary_project_user_ranking where id in (SELECT max(id) FROM biz_summary_project_user_ranking where NOT deleted group by user_id ORDER BY sort asc);").Find(&summaryProjectUserRanking).Error
	return
}

func (r *SummaryProjectUserRankingRepo) FindProjectUserScenarioTotal() (projectUserTotal []model.ProjectUserTotal, err error) {
	err = r.DB.Model(&model.Scenario{}).Raw("select project_id,create_user_id,count(id) as count from biz_scenario where NOT deleted group by project_id,create_user_id; ").Find(&projectUserTotal).Error
	return
}

func (r *SummaryProjectUserRankingRepo) FindProjectUserTestCasesTotal() (projectUserTotal []model.ProjectUserTotal, err error) {
	err = r.DB.Model(&model.Processor{}).Raw("select P.project_id,P.created_by as create_user_id,count(P.id) as count from biz_processor P join biz_scenario S on P.scenario_id = S.id where P.entity_category = 'processor_interface' And P.deleted != 1 And S.deleted != 1 group by project_id,create_user_id order by count desc;").Find(&projectUserTotal).Error
	return
}

func (r *SummaryProjectUserRankingRepo) FindCasesTotalByProjectId(projectId int64) (result []model.UserTotal, err error) {
	err = r.DB.Model(&model.Processor{}).Raw("select P.created_by as create_user_id,count(P.id) as count from biz_processor P join biz_scenario S on P.scenario_id = S.id where P.entity_category = 'processor_interface' and P.project_id = ? And P.deleted != 1 And S.deleted != 1 group by created_by order by count desc;", projectId).Find(&result).Error
	return
}

func (r *SummaryProjectUserRankingRepo) FindScenariosTotalByProjectId(projectId int64) (result []model.UserTotal, err error) {
	err = r.DB.Model(&model.Scenario{}).Raw("select create_user_id,count(id) as count from biz_scenario where project_id = ? And NOT deleted group by create_user_id order by count desc; ", projectId).Find(&result).Error
	return
}

func (r *SummaryProjectUserRankingRepo) FindUserLastUpdateTestCasesByProjectId(projectId int64) (result []model.UserUpdateTime, err error) {
	err = r.DB.Model(&model.Processor{}).Raw("select max(P.updated_at) as updated_at,P.created_by from biz_processor P join biz_scenario S on P.scenario_id = S.id where P.entity_category = 'processor_interface' and P.project_id = ? And P.deleted != 1 And S.deleted != 1 group by created_by;", projectId).Find(&result).Error
	return
}

func (r *SummaryProjectUserRankingRepo) FindAllUserName() (user []model.RankingUser, err error) {
	err = r.DB.Model(&model.SysUser{}).Raw("select id,name from sys_user where NOT deleted; ").Find(&user).Error
	return
}

func (r *SummaryProjectUserRankingRepo) FindUserByProjectId(projectId int64) (user []model.RankingUser, err error) {
	err = r.DB.Model(&model.SysUser{}).Raw("select sys_user.id,sys_user.name from sys_user join biz_project_member on biz_project_member.user_id = sys_user.id where biz_project_member.project_id = ? AND NOT deleted; ", projectId).Find(&user).Error
	return
}

func (r *SummaryProjectUserRankingRepo) FindUserIdsByProjectId(projectId int64) (userIds []int64, err error) {
	err = r.DB.Model(&model.ProjectMember{}).Raw("select user_id from biz_project_member where biz_project_member.project_id = ? AND NOT deleted; ", projectId).Find(&userIds).Error
	return
}

func (r *SummaryProjectUserRankingRepo) CheckUpdated(lastUpdateTime *time.Time) (result bool, err error) {
	result = false
	newTime := time.Now()
	err = r.DB.Model(&model.SummaryBugs{}).Raw("select updated_at from biz_summary_project_user_ranking order by updated_at desc limit 1").Find(&newTime).Error
	result = newTime.After(*lastUpdateTime)
	return
}
