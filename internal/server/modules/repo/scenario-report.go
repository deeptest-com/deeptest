package repo

import (
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/core/dao"
	model "github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	_commUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"gorm.io/gorm"
	"strconv"
)

type ScenarioReportRepo struct {
	DB              *gorm.DB         `inject:""`
	LogRepo         *LogRepo         `inject:""`
	ProjectRepo     *ProjectRepo     `inject:""`
	ScenarioRepo    *ScenarioRepo    `inject:""`
	UserRepo        *UserRepo        `inject:""`
	EnvironmentRepo *EnvironmentRepo `inject:""`
}

func (r *ScenarioReportRepo) Paginate(req v1.ReportReqPaginate) (data _domain.PageData, err error) {
	var count int64

	db := r.DB.Model(&model.ScenarioReport{}).Where(" NOT deleted and scenario_id=?", req.ScenarioId)

	if req.Keywords != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", req.Keywords))
	}

	if req.CreateUserId != 0 {
		db = db.Where("create_user_id = ?", req.CreateUserId)
	}

	err = db.Count(&count).Error
	if err != nil {
		logUtils.Errorf("count report error %s", err.Error())
		return
	}

	results := make([]*model.ScenarioReport, 0)

	err = db.
		Scopes(dao.PaginateScope(req.Page, req.PageSize, req.Order, req.Field)).
		Find(&results).Error
	if err != nil {
		logUtils.Errorf("query report error %s", err.Error())
		return
	}
	r.CombineUserName(results)
	data.Populate(results, count, req.Page, req.PageSize)

	return
}

func (r *ScenarioReportRepo) CombineUserName(data []*model.ScenarioReport) {
	userIds := make([]uint, 0)
	for _, v := range data {
		userIds = append(userIds, v.CreateUserId)
	}
	userIds = _commUtils.ArrayRemoveUintDuplication(userIds)

	users, _ := r.UserRepo.FindByIds(userIds)

	userIdNameMap := make(map[uint]string)
	for _, v := range users {
		userIdNameMap[v.ID] = v.Name
	}

	for _, v := range data {
		if name, ok := userIdNameMap[v.CreateUserId]; ok {
			v.ExecUserName = name
		}
	}
}

func (r *ScenarioReportRepo) Get(id uint) (report model.ScenarioReport, err error) {
	err = r.DB.Where("id = ?", id).First(&report).Error
	if err != nil {
		logUtils.Errorf("find report by id error %s", err.Error())
		return
	}

	var env model.Environment
	env, err = r.EnvironmentRepo.Get(uint(report.ExecEnvId))
	if err != nil {
		logUtils.Errorf("find environment by id error %s", err.Error())
	}

	report.ExecEnv = env.Name
	root, err := r.getLogTree(report)
	report.Logs = root.Logs

	return
}

func (r *ScenarioReportRepo) Create(result *model.ScenarioReport) (bizErr *_domain.BizErr) {
	err := r.DB.Model(&model.ScenarioReport{}).Create(result).Error
	if err != nil {
		logUtils.Errorf("create report error %s", err.Error())
		bizErr.Code = _domain.SystemErr.Code

		return
	}

	if err = r.UpdateSerialNumber(result.ID, result.ProjectId); err != nil {
		logUtils.Errorf("update scenario report serial number error %s", err.Error())
		bizErr.Code = _domain.SystemErr.Code

		return
	}

	return
}

func (r *ScenarioReportRepo) DeleteById(id uint) (err error) {
	err = r.DB.Model(&model.ScenarioReport{}).Where("id = ?", id).
		Updates(map[string]interface{}{"deleted": true}).Error
	if err != nil {
		logUtils.Errorf("delete report by id error %s", err.Error())
		return
	}

	err = r.DB.Model(&model.ExecLogProcessor{}).Where("report_id = ?", id).
		Updates(map[string]interface{}{"deleted": true}).Error
	if err != nil {
		logUtils.Errorf("delete report's logs by id error %s", err.Error())
		return
	}

	return
}

func (r *ScenarioReportRepo) UpdateStatus(progressStatus consts.ProgressStatus, resultStatus consts.ResultStatus, scenarioId uint) (
	err error) {

	values := map[string]interface{}{
		"progress_status": progressStatus,
		"result_status":   resultStatus,
	}
	err = r.DB.Model(&model.ScenarioReport{}).
		Where("scenario_id = ? AND progress_status = ?", scenarioId, consts.InProgress).
		Updates(values).Error

	return
}

func (r *ScenarioReportRepo) UpdateResult(report model.ScenarioReport) (err error) {
	values := map[string]interface{}{
		"pass_num":        report.PassRequestNum,
		"fail_num":        report.FailRequestNum,
		"start_time":      report.StartTime,
		"end_time":        report.EndTime,
		"duration":        report.Duration,
		"progress_status": consts.End,
		"result_status":   report.ResultStatus,
	}
	err = r.DB.Model(&model.ScenarioReport{}).
		Where("id = ?", report.ID).
		Updates(values).Error

	return
}

func (r *ScenarioReportRepo) ResetResult(result model.ScenarioReport) (err error) {
	values := map[string]interface{}{
		"name":       result.Name,
		"start_time": result.StartTime,
	}
	err = r.DB.Model(&result).Where("id = ?", result.ID).Updates(values).Error
	if err != nil {
		logUtils.Errorf("update report error %s", err.Error())
		return
	}

	return
}

func (r *ScenarioReportRepo) ClearLogs(resultId uint) (err error) {
	err = r.DB.Model(&model.ExecLogProcessor{}).Where("result_id = ?", resultId).
		Updates(map[string]interface{}{"deleted": true}).Error
	if err != nil {
		logUtils.Errorf("delete logs by result id error %s", err.Error())
		return
	}

	return
}

func (r *ScenarioReportRepo) FindInProgressResult(scenarioId uint) (result model.ScenarioReport, err error) {
	err = r.DB.Model(&result).
		Where("progress_status =? AND scenario_id = ? AND  not deleted", consts.InProgress, scenarioId).
		First(&result).Error

	return
}

func (r *ScenarioReportRepo) getLogTree(report model.ScenarioReport) (root model.ExecLogProcessor, err error) {
	logs, err := r.LogRepo.ListByReport(report.ID)
	if err != nil {
		return
	}

	for _, log := range logs {
		if log.ProcessorType == consts.ProcessorInterfaceDefault {
			log.InterfaceExtractorsResult, _ = r.listLogExtractors(log.InvokeId)
			log.InterfaceCheckpointsResult, _ = r.listLogCheckpoints(log.InvokeId)

		}
	}

	root = model.ExecLogProcessor{
		Name: report.Name,
	}
	r.makeTree(logs, &root)

	return
}

func (r *ScenarioReportRepo) makeTree(Data []*model.ExecLogProcessor, parent *model.ExecLogProcessor) { //参数为父节点，添加父节点的子节点指针切片
	children, _ := r.haveChild(Data, parent) //判断节点是否有子节点并返回

	if children != nil {
		parent.Logs = append(parent.Logs, children[0:]...) //添加子节点

		for _, child := range children { //查询子节点的子节点，并添加到子节点
			_, has := r.haveChild(Data, child)
			if has {
				r.makeTree(Data, child) //递归添加节点
			}
		}
	}
}

func (r *ScenarioReportRepo) haveChild(Data []*model.ExecLogProcessor, node *model.ExecLogProcessor) (children []*model.ExecLogProcessor, yes bool) {
	for _, v := range Data {
		if v.ParentId == node.ID {
			children = append(children, v)
		}
	}

	if children != nil {
		yes = true
	}

	return
}

func (r *ScenarioReportRepo) listLogExtractors(invokeId uint) (extractors []model.ExecLogExtractor, err error) {
	err = r.DB.
		Where("invoke_id =? AND not deleted", invokeId).
		Find(&extractors).Error

	return
}

func (r *ScenarioReportRepo) listLogCheckpoints(invokeId uint) (checkpoints []model.ExecLogCheckpoint, err error) {
	err = r.DB.
		Where("invoke_id =? AND not deleted", invokeId).
		Find(&checkpoints).Error

	return
}

func (r *ScenarioReportRepo) UpdateSerialNumber(id, projectId uint) (err error) {
	var project model.Project
	project, err = r.ProjectRepo.Get(projectId)
	if err != nil {
		return
	}

	err = r.DB.Model(&model.ScenarioReport{}).Where("id=?", id).Update("serial_number", project.ShortName+"-TR-"+strconv.Itoa(int(id))).Error
	return
}

func (r *ScenarioReportRepo) UpdatePlanReportId(id, planReportId uint) (err error) {
	values := map[string]interface{}{
		"plan_report_id": planReportId,
	}
	err = r.DB.Model(&model.ScenarioReport{}).Where("id = ?", id).Updates(values).Error
	if err != nil {
		logUtils.Errorf("update scenario report error %s", err.Error())
		return
	}

	return
}

func (r *ScenarioReportRepo) BatchUpdatePlanReportId(ids []uint, planReportId uint) (err error) {
	values := map[string]interface{}{
		"plan_report_id": planReportId,
	}
	err = r.DB.Model(&model.ScenarioReport{}).Where("id IN (?)", ids).Updates(values).Error
	if err != nil {
		logUtils.Errorf("batch update scenario reports error %s", err.Error())
		return
	}

	return
}

func (r *ScenarioReportRepo) GetReportsByPlanReportId(planReportId uint) (reports []model.ScenarioReportDetail, err error) {
	var scenarioReports []model.ScenarioReport
	err = r.DB.Model(&model.ScenarioReport{}).Where("plan_report_id = ?", planReportId).Find(&scenarioReports).Error
	if err != nil {
		logUtils.Errorf("find report by id error %s", err.Error())
		return
	}

	for _, report := range scenarioReports {
		root, err := r.getLogTree(report)
		if err != nil {
			continue
		}
		report.Logs = root.Logs
		scenarioReport := model.ScenarioReportDetail{
			ScenarioReport: report,
		}
		reports = append(reports, scenarioReport)
	}
	reports, err = r.CombinePriority(reports)

	return
}

func (r *ScenarioReportRepo) CombinePriority(data []model.ScenarioReportDetail) (res []model.ScenarioReportDetail, err error) {
	scenarioIds := make([]uint, 0)
	for _, v := range data {
		scenarioIds = append(scenarioIds, v.ScenarioId)
	}
	scenarioIds = _commUtils.ArrayRemoveUintDuplication(scenarioIds)

	scenarios, err := r.ScenarioRepo.GetByIds(scenarioIds)
	if err != nil {
		return
	}

	scenarioIdPriorityMap := make(map[uint]string)
	for _, v := range scenarios {
		scenarioIdPriorityMap[v.ID] = v.Priority
	}

	for k, v := range data {
		if priority, ok := scenarioIdPriorityMap[v.ScenarioId]; ok {
			data[k].Priority = priority
		}
	}
	res = data
	return
}

func (r *ScenarioReportRepo) GetBaseReportsByPlanReportId(planReportId uint) (reports []model.ScenarioReport, err error) {
	err = r.DB.Model(&model.ScenarioReport{}).Where("plan_report_id = ?", planReportId).Find(&reports).Error
	if err != nil {
		logUtils.Errorf("find report by id error %s", err.Error())
		return
	}

	return
}

func (r *ScenarioReportRepo) BatchDelete(planReportId uint) (err error) {
	scenarioReportIds := make([]uint, 0)

	scenarioReports, err := r.GetBaseReportsByPlanReportId(planReportId)
	if err != nil {
		return
	}
	for _, v := range scenarioReports {
		scenarioReportIds = append(scenarioReportIds, v.ID)
	}

	if len(scenarioReportIds) == 0 {
		return
	}
	err = r.DB.Model(&model.ScenarioReport{}).Where("id IN (?)", scenarioReportIds).
		Updates(map[string]interface{}{"deleted": true}).Error
	if err != nil {
		logUtils.Errorf("delete report by id error %s", err.Error())
		return
	}

	err = r.DB.Model(&model.ExecLogProcessor{}).Where("report_id IN (?)", scenarioReportIds).
		Updates(map[string]interface{}{"deleted": true}).Error
	if err != nil {
		logUtils.Errorf("delete report's logs by id error %s", err.Error())
		return
	}

	return
}
