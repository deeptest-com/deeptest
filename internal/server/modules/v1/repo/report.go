package repo

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/core/dao"
	serverDomain "github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ReportRepo struct {
	DB       *gorm.DB  `inject:""`
	RoleRepo *RoleRepo `inject:""`
}

func NewReportRepo() *ReportRepo {
	return &ReportRepo{}
}

func (r *ReportRepo) Paginate(req serverDomain.ReportReqPaginate) (data _domain.PageData, err error) {
	var count int64

	db := r.DB.Model(&model.Report{}).Where("NOT deleted")

	if req.Keywords != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", req.Keywords))
	}
	if req.ScenarioId != 0 {
		db = db.Where("scenario_id = ?", req.ScenarioId)
	}

	err = db.Count(&count).Error
	if err != nil {
		logUtils.Errorf("count result error", zap.String("error:", err.Error()))
		return
	}

	results := make([]*model.Report, 0)

	err = db.
		Scopes(dao.PaginateScope(req.Page, req.PageSize, req.Order, req.Field)).
		Find(&results).Error
	if err != nil {
		logUtils.Errorf("query scenario error", zap.String("error:", err.Error()))
		return
	}

	data.Populate(results, count, req.Page, req.PageSize)

	return
}

func (r *ReportRepo) Get(id uint) (scenario model.Report, err error) {
	err = r.DB.Model(&model.Report{}).Where("id = ?", id).First(&scenario).Error
	if err != nil {
		logUtils.Errorf("find scenario by id error", zap.String("error:", err.Error()))
		return scenario, err
	}

	return scenario, nil
}

func (r *ReportRepo) Create(result *model.Report) (bizErr *_domain.BizErr) {
	err := r.DB.Model(&model.Report{}).Create(result).Error
	if err != nil {
		logUtils.Errorf("create test result error", zap.String("error:", err.Error()))
		bizErr.Code = _domain.ErrComm.Code

		return
	}

	return
}

func (r *ReportRepo) DeleteById(id uint) (err error) {
	err = r.DB.Model(&model.Report{}).Where("id = ?", id).
		Updates(map[string]interface{}{"deleted": true}).Error
	if err != nil {
		logUtils.Errorf("delete scenario by id error", zap.String("error:", err.Error()))
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
	err = r.DB.Model(&model.Report{}).
		Where("scenario_id = ? AND progress_status = ?", scenarioId, consts.InProgress).
		Updates(values).Error
	if err != nil {
		logUtils.Errorf("update test result error", zap.String("error:", err.Error()))
		return err
	}

	return nil
}

func (r *ReportRepo) ResetResult(result model.Report) (err error) {
	values := map[string]interface{}{
		"name":       result.Name,
		"start_time": result.StartTime,
	}
	err = r.DB.Model(&result).Where("id = ?", result.ID).Updates(values).Error
	if err != nil {
		logUtils.Errorf("update test result error", zap.String("error:", err.Error()))
		return
	}

	return
}

func (r *ReportRepo) ClearLogs(resultId uint) (err error) {
	err = r.DB.Model(&model.Log{}).Where("result_id = ?", resultId).
		Updates(map[string]interface{}{"deleted": true}).Error
	if err != nil {
		logUtils.Errorf("delete logs by result id error", zap.String("error:", err.Error()))
		return
	}

	return
}

func (r *ReportRepo) FindInProgressResult(scenarioId uint) (result model.Report, err error) {
	err = r.DB.Model(&result).
		Where("progress_status =? AND scenario_id = ? AND  not deleted", consts.InProgress, scenarioId).
		First(&result).Error

	return
}
