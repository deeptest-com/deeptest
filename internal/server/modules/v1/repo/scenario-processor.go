package repo

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"gorm.io/gorm"
)

type ScenarioProcessorRepo struct {
	DB *gorm.DB `inject:""`

	ScenarioNodeRepo *ScenarioNodeRepo `inject:""`
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

func (r *ScenarioProcessorRepo) GetInterface(id uint, processor model.TestProcessor) (ret interface{}, err error) {
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

func (r *ScenarioProcessorRepo) GetLoop(processorId uint, processor model.TestProcessor) (ret interface{}, err error) {
	var entity model.ProcessorLoop
	err = r.DB.Where("processor_id = ?", processorId).First(&entity).Error

	if entity.ID == 0 {
		ret = r.genProcessorComm(processor)
	} else {
		entity.Name = processor.Name
		ret = entity
	}

	return
}

func (r *ScenarioProcessorRepo) GetVariable(processorId uint, processor model.TestProcessor) (ret interface{}, err error) {
	var entity model.ProcessorVariable
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

func (r *ScenarioProcessorRepo) GetCookie(processorId uint, processor model.TestProcessor) (ret interface{}, err error) {
	var entity model.ProcessorCookie
	err = r.DB.Where("processor_id = ?", processorId).First(&entity).Error

	if entity.ID == 0 {
		ret = r.genProcessorComm(processor)
	} else {
		entity.Name = processor.Name
		ret = entity
	}

	return
}

func (r *ScenarioProcessorRepo) GetAssertion(processorId uint, processor model.TestProcessor) (ret interface{}, err error) {
	var entity model.ProcessorAssertion
	err = r.DB.Where("processor_id = ?", processorId).First(&entity).Error

	if entity.ID == 0 {
		ret = r.genProcessorComm(processor)
	} else {
		entity.Name = processor.Name
		ret = entity
	}

	return
}

func (r *ScenarioProcessorRepo) GetExtractor(processorId uint, processor model.TestProcessor) (ret interface{}, err error) {
	var entity model.ProcessorExtractor
	err = r.DB.Where("processor_id = ?", processorId).First(&entity).Error

	if entity.ID == 0 {
		ret = r.genProcessorComm(processor)
	} else {
		entity.Name = processor.Name
		ret = entity
	}

	return
}

func (r *ScenarioProcessorRepo) GetData(processorId uint, processor model.TestProcessor) (ret interface{}, err error) {
	var entity model.ProcessorData
	err = r.DB.Where("processor_id = ?", processorId).First(&entity).Error

	if entity.ID == 0 {
		ret = r.genProcessorComm(processor)
	} else {
		entity.Name = processor.Name
		ret = entity
	}

	return
}

func (r *ScenarioProcessorRepo) SaveGroup(po *model.ProcessorGroup) (err error) {
	err = r.DB.Save(po).Error

	r.UpdateEntityId(po.ProcessorId, po.ID)

	return
}

func (r *ScenarioProcessorRepo) SaveTimer(po *model.ProcessorTimer) (err error) {
	err = r.DB.Save(po).Error

	r.UpdateEntityId(po.ProcessorId, po.ID)

	return
}

func (r *ScenarioProcessorRepo) SaveLogic(po *model.ProcessorLogic) (err error) {
	err = r.DB.Save(po).Error

	r.UpdateEntityId(po.ProcessorId, po.ID)

	return
}

func (r *ScenarioProcessorRepo) SaveLoop(po *model.ProcessorLoop) (err error) {
	err = r.DB.Save(po).Error

	r.UpdateEntityId(po.ProcessorId, po.ID)

	return
}

func (r *ScenarioProcessorRepo) SaveVariable(po *model.ProcessorVariable) (err error) {
	err = r.DB.Save(po).Error

	r.UpdateEntityId(po.ProcessorId, po.ID)

	return
}

func (r *ScenarioProcessorRepo) SaveCookie(po *model.ProcessorCookie) (err error) {
	err = r.DB.Save(po).Error

	r.UpdateEntityId(po.ProcessorId, po.ID)

	return
}

func (r *ScenarioProcessorRepo) SaveAssertion(po *model.ProcessorAssertion) (err error) {
	err = r.DB.Save(po).Error

	r.UpdateEntityId(po.ProcessorId, po.ID)

	return
}

func (r *ScenarioProcessorRepo) SaveExtractor(po *model.ProcessorExtractor) (err error) {
	err = r.DB.Save(po).Error

	r.UpdateEntityId(po.ProcessorId, po.ID)

	return
}

func (r *ScenarioProcessorRepo) SaveData(po *model.ProcessorData) (err error) {
	err = r.DB.Save(po).Error

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

	if processor.InterfaceId > 0 {
		ret.InterfaceId = processor.InterfaceId
	}

	return
}
