package repo

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"gorm.io/gorm"
)

type ScenarioProcessorRepo struct {
	DB *gorm.DB `inject:""`
}

func (r *ScenarioProcessorRepo) GetInterface(id uint, processor model.TestProcessor) (ret interface{}, err error) {
	// TODO
	ret = r.genBaseModel(processor)
	return
}

func (r *ScenarioProcessorRepo) GetGroup(id uint, processor model.TestProcessor) (ret interface{}, err error) {
	var entity model.ProcessorGroup
	err = r.DB.Where("id = ?", id).First(&entity).Error

	if entity.ID == 0 {
		ret = r.genBaseModel(processor)
	} else {
		ret = entity
	}

	return
}

func (r *ScenarioProcessorRepo) GetLogic(id uint, processor model.TestProcessor) (ret interface{}, err error) {
	var entity model.ProcessorLogic
	err = r.DB.Where("id = ?", id).First(&entity).Error

	if entity.ID == 0 {
		ret = r.genBaseModel(processor)
	} else {
		ret = entity
	}

	return
}

func (r *ScenarioProcessorRepo) Get(id uint) (processor model.TestProcessor, err error) {
	err = r.DB.Where("id = ?", id).First(&processor).Error
	return
}

func (r *ScenarioProcessorRepo) UpdateName(id int, name string) (err error) {
	err = r.DB.Model(&model.TestProcessor{}).
		Where("id = ?", id).
		Update("name", name).Error

	return
}

func (r *ScenarioProcessorRepo) Save(po model.ProcessorLogic) (err error) {
	err = r.DB.Save(po).Error
	r.UpdateEntityId(po.ProcessorId, po.ID)

	return
}

func (r *ScenarioProcessorRepo) UpdateEntityId(id, entityId uint) (err error) {
	err = r.DB.Model(&model.TestProcessor{}).
		Where("id = ?", id).
		Update("entityId", entityId).Error

	return
}

func (r *ScenarioProcessorRepo) genBaseModel(processor model.TestProcessor) (ret model.ProcessorBase) {
	ret.Name = processor.Name
	ret.Comments = processor.Comments

	ret.ProcessorCategory = processor.EntityCategory
	ret.ProcessorType = processor.EntityType
	ret.ProcessorId = processor.ID

	return
}
