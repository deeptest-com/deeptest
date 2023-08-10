package repo

import (
	domain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type ScenarioProcessorRepo struct {
	DB *gorm.DB `inject:""`

	ScenarioNodeRepo   *ScenarioNodeRepo   `inject:""`
	ExtractorRepo      *ExtractorRepo      `inject:""`
	CheckpointRepo     *CheckpointRepo     `inject:""`
	DebugInterfaceRepo *DebugInterfaceRepo `inject:""`
}

func (r *ScenarioProcessorRepo) Get(id uint) (processor model.Processor, err error) {
	err = r.DB.Where("id = ?", id).First(&processor).Error
	return
}

func (r *ScenarioProcessorRepo) GetEntity(processorId uint) (ret interface{}, err error) {
	processor, _ := r.Get(processorId)

	switch processor.EntityCategory {
	case consts.ProcessorInterface:
		ret, _ = r.GetInterface(processor)

	case consts.ProcessorGroup:
		ret, _ = r.GetGroup(processor)

	case consts.ProcessorLogic:
		ret, _ = r.GetLogic(processor)

	case consts.ProcessorLoop:
		ret, _ = r.GetLoop(processor)

	case consts.ProcessorVariable:
		ret, _ = r.GetVariable(processor)

	case consts.ProcessorTimer:
		ret, _ = r.GetTimer(processor)

	case consts.ProcessorPrint:
		ret, _ = r.GetPrint(processor)

	case consts.ProcessorCookie:
		ret, _ = r.GetCookie(processor)

	case consts.ProcessorAssertion:
		ret, _ = r.GetAssertion(processor)

	case consts.ProcessorExtractor:
		ret, _ = r.GetExtractor(processor)

	case consts.ProcessorData:
		ret, _ = r.GetData(processor)

	case consts.ProcessorCustomCode:
		ret, _ = r.GetCustomCode(processor)

	default:
	}

	return
}

func (r *ScenarioProcessorRepo) UpdateName(id uint, name string) (err error) {
	err = r.DB.Model(&model.Processor{}).
		Where("id = ?", id).
		Update("name", name).Error

	return
}

func (r *ScenarioProcessorRepo) SaveBasicInfo(req domain.ScenarioProcessorInfo) (err error) {
	err = r.DB.Model(&model.Processor{}).
		Where("id = ?", req.Id).
		Updates(map[string]interface{}{"name": req.Name, "comments": req.Comments}).Error

	return
}

func (r *ScenarioProcessorRepo) GetAll(scenarioId uint) (processors []model.Processor, err error) {
	err = r.DB.Where("scenario_id = ?", scenarioId).
		Find(&processors).Error

	return
}

func (r *ScenarioProcessorRepo) GetRoot(processor model.Processor) (ret model.ProcessorComm, err error) {
	// there is no ProcessorRoot obj, just return a common obj
	ret = r.genProcessorComm(processor)

	return
}

func (r *ScenarioProcessorRepo) GetInterface(processor model.Processor) (ret model.ProcessorComm, err error) {
	// processor refer to an interface using interfaceID,
	// there is no ProcessorInterface obj, just return a common obj
	ret = r.genProcessorComm(processor)
	ret.EntityId = processor.EntityId
	ret.ProcessorInterfaceSrc = processor.ProcessorInterfaceSrc
	ret.Method = processor.Method
	srcName, _ := r.DebugInterfaceRepo.GetSourceNameById(processor.EntityId)
	ret.SrcName = srcName

	return
}

func (r *ScenarioProcessorRepo) GetGroup(processor model.Processor) (ret model.ProcessorGroup, err error) {
	err = r.DB.Where("processor_id = ?", processor.ID).First(&ret).Error

	if ret.ID == 0 {
		comm := r.genProcessorComm(processor)
		copier.CopyWithOption(&ret, comm, copier.Option{DeepCopy: true})
	} else {
		ret.Name = processor.Name
		ret.ParentID = processor.ParentId
	}

	return
}

func (r *ScenarioProcessorRepo) GetLogic(processor model.Processor) (ret model.ProcessorLogic, err error) {
	err = r.DB.Where("processor_id = ?", processor.ID).First(&ret).Error

	if ret.ID == 0 {
		comm := r.genProcessorComm(processor)
		copier.CopyWithOption(&ret, comm, copier.Option{DeepCopy: true})
	} else {
		ret.Name = processor.Name
		ret.ParentID = processor.ParentId
	}

	return
}

func (r *ScenarioProcessorRepo) GetLoop(processor model.Processor) (ret model.ProcessorLoop, err error) {
	err = r.DB.Where("processor_id = ?", processor.ID).First(&ret).Error

	if ret.ID == 0 {
		comm := r.genProcessorComm(processor)
		copier.CopyWithOption(&ret, comm, copier.Option{DeepCopy: true})
	} else {
		ret.Name = processor.Name
		ret.ParentID = processor.ParentId
	}

	return
}

func (r *ScenarioProcessorRepo) GetVariable(processor model.Processor) (ret model.ProcessorVariable, err error) {
	err = r.DB.Where("processor_id = ?", processor.ID).First(&ret).Error

	if ret.ID == 0 {
		comm := r.genProcessorComm(processor)
		copier.CopyWithOption(&ret, comm, copier.Option{DeepCopy: true})
	} else {
		ret.Name = processor.Name
		ret.ParentID = processor.ParentId
	}

	return
}

func (r *ScenarioProcessorRepo) GetTimer(processor model.Processor) (ret model.ProcessorTimer, err error) {
	err = r.DB.Where("processor_id = ?", processor.ID).First(&ret).Error

	if ret.ID == 0 {
		comm := r.genProcessorComm(processor)
		copier.CopyWithOption(&ret, comm, copier.Option{DeepCopy: true})
	} else {
		ret.Name = processor.Name
		ret.ParentID = processor.ParentId
	}

	return
}

func (r *ScenarioProcessorRepo) GetPrint(processor model.Processor) (ret model.ProcessorPrint, err error) {
	err = r.DB.Where("processor_id = ?", processor.ID).First(&ret).Error

	if ret.ID == 0 {
		comm := r.genProcessorComm(processor)
		copier.CopyWithOption(&ret, comm, copier.Option{DeepCopy: true})
	} else {
		ret.Name = processor.Name
		ret.ParentID = processor.ParentId
	}

	return
}

func (r *ScenarioProcessorRepo) GetCookie(processor model.Processor) (ret model.ProcessorCookie, err error) {
	err = r.DB.Where("processor_id = ?", processor.ID).First(&ret).Error

	if ret.ID == 0 {
		comm := r.genProcessorComm(processor)
		copier.CopyWithOption(&ret, comm, copier.Option{DeepCopy: true})
	} else {
		ret.Name = processor.Name
		ret.ParentID = processor.ParentId
	}

	return
}

func (r *ScenarioProcessorRepo) GetAssertion(processor model.Processor) (ret model.ProcessorAssertion, err error) {
	err = r.DB.Where("processor_id = ?", processor.ID).First(&ret).Error

	if ret.ID == 0 {
		comm := r.genProcessorComm(processor)
		copier.CopyWithOption(&ret, comm, copier.Option{DeepCopy: true})
	} else {
		ret.Name = processor.Name
		ret.ParentID = processor.ParentId
	}

	return
}

func (r *ScenarioProcessorRepo) GetExtractor(processor model.Processor) (ret model.ProcessorExtractor, err error) {
	err = r.DB.Where("processor_id = ?", processor.ID).First(&ret).Error

	if ret.ID == 0 {
		comm := r.genProcessorComm(processor)
		copier.CopyWithOption(&ret, comm, copier.Option{DeepCopy: true})
	} else {
		ret.Name = processor.Name
		ret.ParentID = processor.ParentId
	}

	return
}

func (r *ScenarioProcessorRepo) GetData(processor model.Processor) (ret model.ProcessorData, err error) {
	err = r.DB.Where("processor_id = ?", processor.ID).First(&ret).Error

	if ret.ID == 0 {
		comm := r.genProcessorComm(processor)
		copier.CopyWithOption(&ret, comm, copier.Option{DeepCopy: true})

		if ret.RepeatTimes == 0 {
			ret.RepeatTimes = 1
		}
	} else {
		ret.Name = processor.Name
		ret.ParentID = processor.ParentId
	}

	return
}

func (r *ScenarioProcessorRepo) GetCustomCode(processor model.Processor) (ret model.ProcessorCustomCode, err error) {
	err = r.DB.Where("processor_id = ?", processor.ID).First(&ret).Error

	if ret.ID == 0 {
		comm := r.genProcessorComm(processor)
		copier.CopyWithOption(&ret, comm, copier.Option{DeepCopy: true})
	} else {
		ret.Name = processor.Name
		ret.ParentID = processor.ParentId
	}

	return
}

func (r *ScenarioProcessorRepo) SaveGroup(po *model.ProcessorGroup) (err error) {
	err = r.DB.Save(po).Error

	r.UpdateEntityId(po.ProcessorID, po.ID)

	return
}

func (r *ScenarioProcessorRepo) SaveTimer(po *model.ProcessorTimer) (err error) {
	err = r.DB.Save(po).Error

	r.UpdateEntityId(po.ProcessorID, po.ID)

	return
}

func (r *ScenarioProcessorRepo) SavePrint(po *model.ProcessorPrint) (err error) {
	err = r.DB.Save(po).Error

	r.UpdateEntityId(po.ProcessorID, po.ID)

	return
}

func (r *ScenarioProcessorRepo) SaveLogic(po *model.ProcessorLogic) (err error) {
	err = r.DB.Save(po).Error

	r.UpdateEntityId(po.ProcessorID, po.ID)

	return
}

func (r *ScenarioProcessorRepo) SaveLoop(po *model.ProcessorLoop) (err error) {
	err = r.DB.Save(po).Error

	r.UpdateEntityId(po.ProcessorID, po.ID)

	return
}

func (r *ScenarioProcessorRepo) SaveVariable(po *model.ProcessorVariable) (err error) {
	err = r.DB.Save(po).Error

	r.UpdateEntityId(po.ProcessorID, po.ID)

	return
}

func (r *ScenarioProcessorRepo) SaveCookie(po *model.ProcessorCookie) (err error) {
	err = r.DB.Save(po).Error

	r.UpdateEntityId(po.ProcessorID, po.ID)

	return
}

func (r *ScenarioProcessorRepo) SaveAssertion(po *model.ProcessorAssertion) (err error) {
	err = r.DB.Save(po).Error

	r.UpdateEntityId(po.ProcessorID, po.ID)

	return
}

func (r *ScenarioProcessorRepo) SaveExtractor(po *model.ProcessorExtractor) (err error) {
	err = r.DB.Save(po).Error

	r.UpdateEntityId(po.ProcessorID, po.ID)

	return
}

func (r *ScenarioProcessorRepo) SaveData(po *model.ProcessorData) (err error) {
	err = r.DB.Save(po).Error

	r.UpdateEntityId(po.ProcessorID, po.ID)

	return
}

func (r *ScenarioProcessorRepo) SaveCustomCode(po *model.ProcessorCustomCode) (err error) {
	err = r.DB.Save(po).Error
	r.UpdateEntityId(po.ProcessorID, po.ID)

	return
}

func (r *ScenarioProcessorRepo) UpdateEntityId(id, entityId uint) (err error) {
	err = r.DB.Model(&model.Processor{}).
		Where("id = ?", id).
		Update("entity_id", entityId).Error

	return
}

func (r *ScenarioProcessorRepo) genProcessorComm(processor model.Processor) (ret model.ProcessorComm) {
	ret.Id = processor.ID
	ret.Name = processor.Name
	ret.Comments = processor.Comments

	ret.ProcessorCategory = processor.EntityCategory
	ret.ProcessorType = processor.EntityType
	ret.ProcessorID = processor.ID
	ret.ParentID = processor.ParentId

	//ret = modelRef.ProcessorComm{
	//	ProcessorEntityBase: agentExec.ProcessorEntityBase{
	//		Name:              processor.Name,
	//		ProcessorCategory: processor.EntityCategory,
	//		ProcessorType:     processor.EntityType,
	//		ProcessorID:       processor.ID,
	//		ParentID:          processor.ParentId,
	//	},
	//}
	//if processor.EndpointInterfaceId > 0 {
	//	ret.EndpointInterfaceId = processor.EndpointInterfaceId
	//}

	return
}

func (r *ScenarioProcessorRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&model.Processor{}).
		Where("id=?", id).
		Update("deleted", true).
		Error

	return
}

func (r *ScenarioProcessorRepo) UpdateInterfaceId(id, debugInterfacceId uint) (err error) {
	err = r.DB.Model(&model.Processor{}).
		Where("id=?", id).
		Update("entity_id", debugInterfacceId).
		Error

	return
}

func (r *ScenarioProcessorRepo) UpdateMethod(id uint, method consts.HttpMethod) (err error) {
	err = r.DB.Model(&model.Processor{}).
		Where("id=?", id).
		Update("method", method).
		Error

	return
}
