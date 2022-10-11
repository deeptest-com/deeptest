package repo

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	model "github.com/aaronchen2k/deeptest/internal/server/modules/model"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ExtractorRepo struct {
	DB            *gorm.DB       `inject:""`
	InterfaceRepo *InterfaceRepo `inject:""`
}

func (r *ExtractorRepo) List(interfaceId uint) (pos []model.InterfaceExtractor, err error) {
	err = r.DB.
		Where("interface_id=?", interfaceId).
		Where("NOT deleted").
		Order("created_at ASC").
		Find(&pos).Error
	return
}

func (r *ExtractorRepo) ListTo(interfaceId uint) (ret []domain.Extractor, err error) {
	pos := make([]model.InterfaceExtractor, 0)

	err = r.DB.
		Where("interface_id=?", interfaceId).
		Where("NOT deleted").
		Find(&pos).Error

	for _, po := range pos {
		extractor := domain.Extractor{}
		copier.CopyWithOption(&extractor, po, copier.Option{DeepCopy: true})

		ret = append(ret, extractor)
	}

	return
}

func (r *ExtractorRepo) Get(id uint) (extractor model.InterfaceExtractor, err error) {
	err = r.DB.
		Where("id=?", id).
		Where("NOT deleted").
		First(&extractor).Error
	return
}

func (r *ExtractorRepo) GetByVariable(variable string, id uint, interfaceId uint) (extractor model.InterfaceExtractor, err error) {
	db := r.DB.Model(&extractor).
		Where("variable = ? AND interface_id =? AND not deleted",
			variable, interfaceId)

	if id > 0 {
		db.Where("id != ?", id)
	}

	db.First(&extractor)

	return
}

func (r *ExtractorRepo) Save(extractor *model.InterfaceExtractor) (id uint, bizErr _domain.BizErr) {
	po, _ := r.GetByVariable(extractor.Variable, extractor.ID, extractor.InterfaceId)
	if po.ID > 0 {
		bizErr.Code = _domain.ErrNameExist.Code
		return
	}

	err := r.DB.Save(extractor).Error
	if err != nil {
		bizErr.Code = _domain.SystemErr.Code
		return
	}

	id = extractor.ID

	return
}

func (r *ExtractorRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&model.InterfaceExtractor{}).
		Where("id=?", id).
		Update("deleted", true).
		Error

	return
}

func (r *ExtractorRepo) UpdateResult(extractor model.InterfaceExtractor) (err error) {
	values := map[string]interface{}{
		"result": extractor.Result,
	}
	err = r.DB.Model(&extractor).Where("id = ?", extractor.ID).
		Updates(values).Error
	if err != nil {
		logUtils.Errorf("update scenario error", zap.String("error:", err.Error()))
		return err
	}

	return
}
func (r *ExtractorRepo) UpdateResultToExecLog(extractor model.InterfaceExtractor, log *model.ExecLogProcessor) (
	logExtractor model.ExecLogExtractor, err error) {

	copier.CopyWithOption(&logExtractor, extractor, copier.Option{DeepCopy: true})
	logExtractor.ID = 0
	logExtractor.LogId = log.ID
	logExtractor.CreatedAt = nil
	logExtractor.UpdatedAt = nil

	err = r.DB.Save(&logExtractor).Error

	return
}

func (r *ExtractorRepo) ListValidExtractorVariable(interfaceId, projectId uint) (variables []v1.Variable, err error) {
	err = r.DB.Model(&model.InterfaceExtractor{}).
		Select("id, variable AS name, result AS value, "+
			"interface_id AS interfaceId, scope AS scope").
		//Where("(is_share OR interface_id = ?) AND enable_share", interfaceId).
		Where("project_id=?", projectId).
		Where("NOT deleted AND NOT disabled").
		Order("created_at ASC").
		Find(&variables).Error

	return
}

func (r *ExtractorRepo) ListExtractorVariableByInterface(interfaceId uint) (variables []v1.Variable, err error) {
	err = r.DB.Model(&model.InterfaceExtractor{}).
		Select("id, variable AS name, result AS value").
		Where("interface_id=?", interfaceId).
		Where("NOT deleted AND NOT disabled").
		Order("created_at ASC").
		Find(&variables).Error

	return
}
