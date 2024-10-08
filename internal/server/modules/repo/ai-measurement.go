package repo

import (
	v1 "github.com/deeptest-com/deeptest/cmd/server/v1/domain"
	"github.com/deeptest-com/deeptest/internal/pkg/domain"
	"github.com/deeptest-com/deeptest/internal/server/modules/model"
	_logUtils "github.com/deeptest-com/deeptest/pkg/lib/log"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type AiMeasurementRepo struct {
	*BaseRepo `inject:""`
	DB        *gorm.DB `inject:""`

	AiMetricsRepo *AiMetricsRepo `inject:""`
}

func (r *AiMeasurementRepo) List() (pos []model.AiMeasurement, err error) {
	err = r.DB.Model(&model.AiMeasurement{}).
		Where("NOT deleted").
		Find(&pos).Error

	return
}

func (r *AiMeasurementRepo) Get(id uint) (po model.AiMeasurement, err error) {
	err = r.DB.Where("id = ?", id).
		First(&po).Error

	return
}

func (r *AiMeasurementRepo) LoadForExec(id uint) (cs domain.AiMeasurement, metricsArr []domain.AiMetrics, err error) {
	casePo, err := r.Get(id)
	if err != nil {
		return
	}

	copier.CopyWithOption(&cs, casePo, copier.Option{DeepCopy: true})

	metricsPos, err := r.AiMetricsRepo.ListByMeasurement(casePo)
	if err != nil {
		return
	}

	for _, po := range metricsPos {
		to := domain.AiMetrics{}
		copier.CopyWithOption(&to, po, copier.Option{DeepCopy: true})

		metricsArr = append(metricsArr, to)
	}

	return
}

func (r *AiMeasurementRepo) Create(req *v1.AiMeasurementCreateReq) (user model.AiMeasurement, err error) {
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

func (r *AiMeasurementRepo) Update(user model.AiMeasurement) (err error) {
	err = r.DB.Save(&user).Error
	if err != nil {
		_logUtils.Error(err.Error())
		return
	}

	return
}

func (r *AiMeasurementRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&model.AiMeasurement{}).
		Where("id = ?", id).
		Update("deleted", true).Error

	return
}
