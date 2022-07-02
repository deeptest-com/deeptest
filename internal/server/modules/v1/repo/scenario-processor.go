package repo

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"gorm.io/gorm"
)

type ScenarioProcessorRepo struct {
	DB *gorm.DB `inject:""`

	ScenarioNodeRepo *ScenarioNodeRepo `inject:""`
}

func (r *ScenarioProcessorRepo) GetInterface(id uint, processor model.TestProcessor) (ret interface{}, err error) {
	// TODO
	ret = r.genProcessorComm(processor)
	return
}

func (r *ScenarioProcessorRepo) GetGroup(processorId uint, processor model.TestProcessor) (ret interface{}, err error) {
	var entity model.ProcessorGroup
	err = r.DB.Where("processor_id = ?", processorId).First(&entity).Error

	if entity.ID == 0 {
		ret = r.genProcessorComm(processor)
	} else {
		entity.Name = processor.Name
		ret = entity
	}

	return
}

func (r *ScenarioProcessorRepo) GetTimer(processorId uint, processor model.TestProcessor) (ret interface{}, err error) {
	var entity model.ProcessorTimer
	err = r.DB.Where("processor_id = ?", processorId).First(&entity).Error

	if entity.ID == 0 {
		ret = r.genProcessorComm(processor)
	} else {
		entity.Name = processor.Name
		ret = entity
	}

	return
}

func (r *ScenarioProcessorRepo) GetLogic(processorId uint, processor model.TestProcessor) (ret interface{}, err error) {
	var entity model.ProcessorLogic
	err = r.DB.Where("processor_id = ?", processorId).First(&entity).Error

	if entity.ID == 0 {
		ret = r.genProcessorComm(processor)
	} else {
		entity.Name = processor.Name
		ret = entity
	}

	return
}

func (r *ScenarioProcessorRepo) Get(id uint) (processor model.TestProcessor, err error) {
	err = r.DB.Where("id = ?", id).First(&processor).Error
	return
}

func (r *ScenarioProcessorRepo) UpdateName(id uint, name string) (err error) {
	err = r.DB.Model(&model.TestProcessor{}).
		Where("id = ?", id).
		Update("name", name).Error

	return
}

func (r *ScenarioProcessorRepo) SaveGroup(po model.ProcessorGroup) (err error) {
	err = r.DB.Save(&po).Error

	r.UpdateEntityId(po.ProcessorId, po.ID)

	return
}

func (r *ScenarioProcessorRepo) SaveTimer(po model.ProcessorTimer) (err error) {
	err = r.DB.Save(&po).Error

	r.UpdateEntityId(po.ProcessorId, po.ID)

	return
}

func (r *ScenarioProcessorRepo) SaveLogic(po model.ProcessorLogic) (err error) {
	err = r.DB.Save(&po).Error

	r.UpdateEntityId(po.ProcessorId, po.ID)

	return
}

func (r *ScenarioProcessorRepo) UpdateEntityId(id, entityId uint) (err error) {
	err = r.DB.Model(&model.TestProcessor{}).
		Where("id = ?", id).
		Update("entity_id", entityId).Error

	return
}

func (r *ScenarioProcessorRepo) genProcessorComm(processor model.TestProcessor) (ret model.ProcessorComm) {
	ret.Id = 0
	ret.Name = processor.Name
	ret.Comments = processor.Comments

	ret.ProcessorCategory = processor.EntityCategory
	ret.ProcessorType = processor.EntityType
	ret.ProcessorId = processor.ID

	return
}
