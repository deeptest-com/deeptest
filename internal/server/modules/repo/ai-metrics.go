package repo

import (
	"fmt"
	v1 "github.com/deeptest-com/deeptest/cmd/server/v1/domain"
	"github.com/deeptest-com/deeptest/internal/server/modules/model"
	_logUtils "github.com/deeptest-com/deeptest/pkg/lib/log"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type AiMetricsRepo struct {
	*BaseRepo `inject:""`
	DB        *gorm.DB `inject:""`
}

func (r *AiMetricsRepo) ListByMeasurement(cs model.AiMeasurement) (pos []model.AiMetrics, err error) {
	err = r.DB.Where(fmt.Sprintf("id IN (%s)", cs.MetricsIds)).
		Find(&pos).Error

	return
}

func (r *AiMetricsRepo) Get(id uint) (po model.AiMeasurement, err error) {
	err = r.DB.Where("id = ?", id).
		First(&po).Error

	return
}
func (r *AiMetricsRepo) GetSummarization(id uint) (po model.AiMetricsSummarization, err error) {
	err = r.DB.Where("id = ?", id).
		First(&po).Error

	return
}
func (r *AiMetricsRepo) GetAnswerRelevancy(id uint) (po model.AiMetricsAnswerRelevancy, err error) {
	err = r.DB.Where("id = ?", id).
		First(&po).Error

	return
}
func (r *AiMetricsRepo) GetFaithfulness(id uint) (po model.AiMetricsFaithfulness, err error) {
	err = r.DB.Where("id = ?", id).
		First(&po).Error

	return
}
func (r *AiMetricsRepo) GetContextualPrecision(id uint) (po model.AiMetricsContextualPrecision, err error) {
	err = r.DB.Where("id = ?", id).
		First(&po).Error

	return
}
func (r *AiMetricsRepo) GetContextualRecall(id uint) (po model.AiMetricsContextualRecall, err error) {
	err = r.DB.Where("id = ?", id).
		First(&po).Error

	return
}
func (r *AiMetricsRepo) GetContextualRelevancy(id uint) (po model.AiMetricsContextualRelevancy, err error) {
	err = r.DB.Where("id = ?", id).
		First(&po).Error

	return
}
func (r *AiMetricsRepo) GetHallucination(id uint) (po model.AiMetricsHallucination, err error) {
	err = r.DB.Where("id = ?", id).
		First(&po).Error

	return
}
func (r *AiMetricsRepo) GetBias(id uint) (po model.AiMetricsBias, err error) {
	err = r.DB.Where("id = ?", id).
		First(&po).Error

	return
}
func (r *AiMetricsRepo) GetToxicity(id uint) (po model.AiMetricsToxicity, err error) {
	err = r.DB.Where("id = ?", id).
		First(&po).Error

	return
}

func (r *AiMetricsRepo) Create(req *v1.AiMeasurementCreateReq) (user model.AiMeasurement, err error) {
	user = model.AiMeasurement{}

	err = copier.CopyWithOption(&user, req, copier.Option{DeepCopy: true})
	if err != nil {
		_logUtils.Error(err.Error())
		return
	}

	err = r.DB.Create(&user).Error
	if err != nil {
		_logUtils.Error(err.Error())
		return
	}

	return
}

func (r *AiMetricsRepo) Update(user model.AiMeasurement) (err error) {
	err = r.DB.Save(&user).Error
	if err != nil {
		_logUtils.Error(err.Error())
		return
	}

	return
}

func (r *AiMetricsRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&model.AiMeasurement{}).
		Where("id = ?", id).
		Update("deleted", true).Error

	return
}
