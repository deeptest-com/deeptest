package repo

import (
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/core/dao"
	model2 "github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"gorm.io/gorm"
)

type ReportRepo struct {
	DB      *gorm.DB `inject:""`
	LogRepo *LogRepo `inject:""`
}

func NewReportRepo() *ReportRepo {
	return &ReportRepo{}
}

func (r *ReportRepo) Paginate(req v1.ReportReqPaginate, projectId int) (data _domain.PageData, err error) {
	var count int64

	db := r.DB.Model(&model2.Report{}).
		Where("project_id = ? AND NOT deleted", projectId)

	if req.Keywords != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", req.Keywords))
	}
	if req.ScenarioId != 0 {
		db = db.Where("scenario_id = ?", req.ScenarioId)
	}

	err = db.Count(&count).Error
	if err != nil {
		logUtils.Errorf("count report error %s", err.Error())
		return
	}

	results := make([]*model2.Report, 0)

	err = db.
		Scopes(dao.PaginateScope(req.Page, req.PageSize, req.Order, req.Field)).
		Find(&results).Error
	if err != nil {
		logUtils.Errorf("query report error %s", err.Error())
		return
	}

	data.Populate(results, count, req.Page, req.PageSize)

	return
}

func (r *ReportRepo) Get(id uint) (report model2.Report, err error) {
	err = r.DB.Where("id = ?", id).First(&report).Error
	if err != nil {
		logUtils.Errorf("find report by id error %s", err.Error())
		return
	}

	root, err := r.getLogTree(report)
	report.Logs = root.Logs

	return
}

func (r *ReportRepo) Create(result *model2.Report) (bizErr *_domain.BizErr) {
	err := r.DB.Model(&model2.Report{}).Create(result).Error
	if err != nil {
		logUtils.Errorf("create report error %s", err.Error())
		bizErr.Code = _domain.SystemErr.Code

		return
	}

	return
}

func (r *ReportRepo) DeleteById(id uint) (err error) {
	err = r.DB.Model(&model2.Report{}).Where("id = ?", id).
		Updates(map[string]interface{}{"deleted": true}).Error
	if err != nil {
		logUtils.Errorf("delete report by id error %s", err.Error())
		return
	}

	err = r.DB.Model(&model2.ExecLogProcessor{}).Where("report_id = ?", id).
		Updates(map[string]interface{}{"deleted": true}).Error
	if err != nil {
		logUtils.Errorf("delete report's logs by id error %s", err.Error())
		return
	}

	return
}

func (r *ReportRepo) UpdateStatus(progressStatus consts.ProgressStatus, resultStatus consts.ResultStatus, scenarioId uint) (
	err error) {

	values := map[string]interface{}{
		"progress_status": progressStatus,
		"result_status":   resultStatus,
	}
	err = r.DB.Model(&model2.Report{}).
		Where("report_id = ? AND progress_status = ?", scenarioId, consts.InProgress).
		Updates(values).Error

	return
}

func (r *ReportRepo) UpdateResult(report model2.Report) (err error) {
	values := map[string]interface{}{
		"pass_num":        report.PassRequestNum,
		"fail_num":        report.FailRequestNum,
		"start_time":      report.StartTime,
		"end_time":        report.EndTime,
		"duration":        report.Duration,
		"progress_status": consts.End,
		"result_status":   report.ResultStatus,
	}
	err = r.DB.Model(&model2.Report{}).
		Where("id = ?", report.ID).
		Updates(values).Error

	return
}

func (r *ReportRepo) ResetResult(result model2.Report) (err error) {
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

func (r *ReportRepo) ClearLogs(resultId uint) (err error) {
	err = r.DB.Model(&model2.ExecLogProcessor{}).Where("result_id = ?", resultId).
		Updates(map[string]interface{}{"deleted": true}).Error
	if err != nil {
		logUtils.Errorf("delete logs by result id error %s", err.Error())
		return
	}

	return
}

func (r *ReportRepo) FindInProgressResult(scenarioId uint) (result model2.Report, err error) {
	err = r.DB.Model(&result).
		Where("progress_status =? AND scenario_id = ? AND  not deleted", consts.InProgress, scenarioId).
		First(&result).Error

	return
}

func (r *ReportRepo) getLogTree(report model2.Report) (root model2.ExecLogProcessor, err error) {
	logs, err := r.LogRepo.ListByReport(report.ID)
	if err != nil {
		return
	}

	for _, log := range logs {
		if log.ProcessorType == consts.ProcessorInterfaceDefault {
			log.InterfaceExtractorsResult, _ = r.listLogExtractors(log.ID)
			log.InterfaceCheckpointsResult, _ = r.listLogCheckpoints(log.ID)
		}
	}

	root = model2.ExecLogProcessor{
		Name: report.Name,
	}
	r.makeTree(logs, &root)

	return
}

func (r *ReportRepo) makeTree(Data []*model2.ExecLogProcessor, parent *model2.ExecLogProcessor) { //参数为父节点，添加父节点的子节点指针切片
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

func (r *ReportRepo) haveChild(Data []*model2.ExecLogProcessor, node *model2.ExecLogProcessor) (children []*model2.ExecLogProcessor, yes bool) {
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

func (r *ReportRepo) listLogExtractors(logId uint) (extractors []model2.ExecLogExtractor, err error) {
	err = r.DB.
		Where("log_id =? AND not deleted", logId).
		Find(&extractors).Error

	return
}

func (r *ReportRepo) listLogCheckpoints(logId uint) (checkpoints []model2.ExecLogCheckpoint, err error) {
	err = r.DB.
		Where("log_id =? AND not deleted", logId).
		Find(&checkpoints).Error

	return
}
