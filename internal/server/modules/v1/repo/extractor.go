package repo

import (
	"fmt"
	_domain "github.com/aaronchen2k/deeptest/internal/pkg/domain"
	serverDomain "github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"gorm.io/gorm"
)

type ExtractorRepo struct {
	DB *gorm.DB `inject:""`
}

func (r *ExtractorRepo) List(interfaceId uint) (pos []model.InterfaceExtractor, err error) {
	err = r.DB.
		Where("interface_id=?", interfaceId).
		Where("NOT deleted").
		Order("created_at ASC").
		Find(&pos).Error
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

func (r *ExtractorRepo) Save(extractor *model.InterfaceExtractor) (id uint, err error) {
	po, _ := r.GetByVariable(extractor.Variable, extractor.ID, extractor.InterfaceId)
	if po.ID > 0 {
		err = fmt.Errorf("%d", _domain.BizErrNameExist.Code)
		return
	}

	err = r.DB.Save(extractor).Error
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
	err = r.DB.Model(&extractor).
		Where("id=?", extractor.ID).
		Update("result", extractor.Result).
		Error

	return
}

func (r *ExtractorRepo) ListExtractorVariable(interfaceId uint) (variables []serverDomain.Variable, err error) {
	err = r.DB.Model(&model.InterfaceExtractor{}).
		Select("id, variable AS name, result AS value").
		Where("interface_id=?", interfaceId).
		Where("NOT deleted AND NOT disabled").
		Order("created_at ASC").
		Find(&variables).Error

	return
}
