package repo

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type TestResultRepo struct {
	DB       *gorm.DB  `inject:""`
	RoleRepo *RoleRepo `inject:""`
}

func NewTestResultRepo() *TestResultRepo {
	return &TestResultRepo{}
}

func (r *TestResultRepo) Get(id uint) (scenario model.TestResult, err error) {
	err = r.DB.Model(&model.TestResult{}).Where("id = ?", id).First(&scenario).Error
	if err != nil {
		logUtils.Errorf("find scenario by id error", zap.String("error:", err.Error()))
		return scenario, err
	}

	return scenario, nil
}

func (r *TestResultRepo) Create(result *model.TestResult) (bizErr *_domain.BizErr) {
	err := r.DB.Model(&model.TestResult{}).Create(result).Error
	if err != nil {
		logUtils.Errorf("create test result error", zap.String("error:", err.Error()))
		bizErr.Code = _domain.ErrComm.Code

		return
	}

	return
}

func (r *TestResultRepo) DeleteById(id uint) (err error) {
	err = r.DB.Model(&model.TestResult{}).Where("id = ?", id).
		Updates(map[string]interface{}{"deleted": true}).Error
	if err != nil {
		logUtils.Errorf("delete scenario by id error", zap.String("error:", err.Error()))
		return
	}

	return
}

func (r *TestResultRepo) UpdateStatus(progressStatus consts.ProgressStatus, resultStatus consts.ResultStatus, scenarioId int) (
	err error) {

	values := map[string]interface{}{
		"progress_status": progressStatus,
		"result_status":   resultStatus,
	}
	err = r.DB.Model(&model.TestResult{}).
		Where("scenario_id = ? AND progress_status = ?", scenarioId, consts.InProgress).
		Updates(values).Error
	if err != nil {
		logUtils.Errorf("update test result error", zap.String("error:", err.Error()))
		return err
	}

	return nil
}

func (r *TestResultRepo) ResetResult(result model.TestResult) (err error) {
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

func (r *TestResultRepo) ClearLogs(resultId uint) (err error) {
	err = r.DB.Model(&model.TestLog{}).Where("result_id = ?", resultId).
		Updates(map[string]interface{}{"deleted": true}).Error
	if err != nil {
		logUtils.Errorf("delete logs by result id error", zap.String("error:", err.Error()))
		return
	}

	return
}

func (r *TestResultRepo) FindInProgressResult(scenarioId uint) (result model.TestResult, err error) {
	err = r.DB.Model(&result).
		Where("progress_status =? AND scenario_id = ? AND  not deleted", consts.InProgress, scenarioId).
		First(&result).Error

	return
}
