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

	case consts.ProcessorData:
		ret, _ = r.GetData(processor)

	case consts.ProcessorCustomCode:
		ret, _ = r.GetCustomCode(processor)

	default:
	}

	return
}

func (r *ScenarioProcessorRepo) CopyEntity(srcProcessorId, distProcessorId uint) (err error) {
	//srcProcessor, err := r.Get(srcProcessorId)
	//if err != nil {
	//	return err
	//}

	distProcessor, err := r.Get(distProcessorId)
	if err != nil {
		return err
	}

	if distProcessor.EntityCategory == consts.ProcessorInterface {
		//entityId, err := r.DebugInterfaceRepo.CopyById(srcEntityId)
		//if err != nil {
		//	return err
		//}
	} else {
		entity, err := r.GetEntity(srcProcessorId)
		if err != nil {
			return err
		}

		//var entityBase interface{}
		//entityBase = agentExec.ProcessorEntityBase{
		//	ProcessorID: distProcessorId,
		//	ParentID:    distProcessor.ParentId,
		//}
		//baseModel := model.BaseModel{ID: 0}
		//err = copier.CopyWithOption(&entity, &entityBase, copier.Option{DeepCopy: true})
		//err = copier.CopyWithOption(&entity, &baseModel, copier.Option{IgnoreEmpty: true, DeepCopy: true})
		err = r.DB.Save(&entity).Error
		if err != nil {
			return err
		}

		//_ = copier.CopyWithOption(&baseModel, &entity, copier.Option{IgnoreEmpty: true, DeepCopy: true})

		//err = r.UpdateInterfaceId(distProcessorId, baseModel.ID)
		//if err != nil {
		//	return err
		//}
	}

	return

	//processor, err := r.Get(processorId)
	//if err != nil {
	//	return
	//}
	//
	//srcEntityId := processor.EntityId
	//var entityId uint
	//if entityCategory == consts.ProcessorInterface {
	//	entityId, err = r.DebugInterfaceRepo.CopyById(srcEntityId)
	//	if err != nil {
	//		return err
	//	}
	//} else {
	//	switch entityCategory {
	//	case consts.ProcessorGroup:
	//		group, err := r.GetGroupById(srcEntityId)
	//		if err != nil {
	//			return err
	//		}
	//
	//		group.ID = 0
	//		group.ProcessorID = processor.ID
	//		group.ProcessorCategory = processor.EntityCategory
	//		group.ProcessorType = processor.EntityType
	//		group.ParentID = processor.ParentId
	//
	//		err = r.SaveGroup(&group)
	//		if err != nil {
	//			return err
	//		}
	//
	//		entityId = group.ID
	//	case consts.ProcessorLogic:
	//		logic, err := r.GetLogicById(srcEntityId)
	//		if err != nil {
	//			return err
	//		}
	//
	//		logic.ID = 0
	//		logic.ProcessorID = processor.ID
	//		logic.ProcessorCategory = processor.EntityCategory
	//		logic.ProcessorType = processor.EntityType
	//		logic.ParentID = processor.ParentId
	//
	//		err = r.SaveLogic(&logic)
	//		if err != nil {
	//			return err
	//		}
	//
	//		entityId = logic.ID
	//	case consts.ProcessorLoop:
	//		loop, err := r.GetLoopById(srcEntityId)
	//		if err != nil {
	//			return err
	//		}
	//
	//		loop.ID = 0
	//		loop.ProcessorID = processor.ID
	//		loop.ProcessorCategory = processor.EntityCategory
	//		loop.ProcessorType = processor.EntityType
	//		loop.ParentID = processor.ParentId
	//
	//		err = r.SaveLoop(&loop)
	//		if err != nil {
	//			return err
	//		}
	//
	//		entityId = loop.ID
	//	case consts.ProcessorVariable:
	//		variable, err := r.GetVariableById(srcEntityId)
	//		if err != nil {
	//			return err
	//		}
	//
	//		variable.ID = 0
	//		variable.ProcessorID = processor.ID
	//		variable.ProcessorCategory = processor.EntityCategory
	//		variable.ProcessorType = processor.EntityType
	//		variable.ParentID = processor.ParentId
	//
	//		err = r.SaveVariable(&variable)
	//		if err != nil {
	//			return err
	//		}
	//
	//		entityId = variable.ID
	//	case consts.ProcessorTimer:
	//		timer, err := r.GetTimerById(srcEntityId)
	//		if err != nil {
	//			return err
	//		}
	//
	//		timer.ID = 0
	//		timer.ProcessorID = processor.ID
	//		timer.ProcessorCategory = processor.EntityCategory
	//		timer.ProcessorType = processor.EntityType
	//		timer.ParentID = processor.ParentId
	//
	//		err = r.SaveTimer(&timer)
	//		if err != nil {
	//			return err
	//		}
	//
	//		entityId = timer.ID
	//	case consts.ProcessorPrint:
	//		srcPrint, err := r.GetPrintById(srcEntityId)
	//		if err != nil {
	//			return err
	//		}
	//
	//		srcPrint.ID = 0
	//		srcPrint.ProcessorID = processor.ID
	//		srcPrint.ProcessorCategory = processor.EntityCategory
	//		srcPrint.ProcessorType = processor.EntityType
	//		srcPrint.ParentID = processor.ParentId
	//
	//		err = r.SavePrint(&srcPrint)
	//		if err != nil {
	//			return err
	//		}
	//
	//		entityId = srcPrint.ID
	//	case consts.ProcessorCookie:
	//		cookie, err := r.GetCookieById(srcEntityId)
	//		if err != nil {
	//			return err
	//		}
	//
	//		cookie.ID = 0
	//		cookie.ProcessorID = processor.ID
	//		cookie.ProcessorCategory = processor.EntityCategory
	//		cookie.ProcessorType = processor.EntityType
	//		cookie.ParentID = processor.ParentId
	//
	//		err = r.SaveCookie(&cookie)
	//		if err != nil {
	//			return err
	//		}
	//
	//		entityId = cookie.ID
	//	case consts.ProcessorAssertion:
	//		srcAssertion, err := r.GetAssertionById(srcId)
	//		if err != nil {
	//			return err
	//		}
	//
	//		cookie.ID = 0
	//		cookie.ProcessorID = processor.ID
	//		cookie.ProcessorCategory = processor.EntityCategory
	//		cookie.ProcessorType = processor.EntityType
	//		cookie.ParentID = processor.ParentId
	//
	//		err = r.SaveCookie(&cookie)
	//		if err != nil {
	//			return err
	//		}
	//
	//		entityId = cookie.ID
	//	case consts.ProcessorData:
	//		srcAssertion, err := r.GetDataById(srcId)
	//	case consts.ProcessorCustomCode:
	//		ret, _ = r.GetCustomCode(processor)
	//
	//	default:
	//	}
	//}

	//copier.CopyWithOption(&version, &req, copier.Option{IgnoreEmpty: true, DeepCopy: true})

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

func (r *ScenarioProcessorRepo) GetGroupById(id uint) (ret model.ProcessorGroup, err error) {
	err = r.DB.Where("id = ?", id).First(&ret).Error
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

func (r *ScenarioProcessorRepo) GetLogicById(id uint) (ret model.ProcessorLogic, err error) {
	err = r.DB.Where("id = ?", id).First(&ret).Error

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

func (r *ScenarioProcessorRepo) GetLoopById(id uint) (ret model.ProcessorLoop, err error) {
	err = r.DB.Where("id = ?", id).First(&ret).Error
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

func (r *ScenarioProcessorRepo) GetVariableById(id uint) (ret model.ProcessorVariable, err error) {
	err = r.DB.Where("id = ?", id).First(&ret).Error
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

func (r *ScenarioProcessorRepo) GetTimerById(id uint) (ret model.ProcessorTimer, err error) {
	err = r.DB.Where("id = ?", id).First(&ret).Error
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

func (r *ScenarioProcessorRepo) GetPrintById(id uint) (ret model.ProcessorPrint, err error) {
	err = r.DB.Where("id = ?", id).First(&ret).Error
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

func (r *ScenarioProcessorRepo) GetCookieById(id uint) (ret model.ProcessorCookie, err error) {
	err = r.DB.Where("id = ?", id).First(&ret).Error
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

func (r *ScenarioProcessorRepo) GetAssertionById(id uint) (ret model.ProcessorAssertion, err error) {
	err = r.DB.Where("id = ?", id).First(&ret).Error
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

func (r *ScenarioProcessorRepo) GetDataById(id uint) (ret model.ProcessorData, err error) {
	err = r.DB.Where("id = ?", id).First(&ret).Error
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

func (r *ScenarioProcessorRepo) GetCustomCodeById(id uint) (ret model.ProcessorCustomCode, err error) {
	err = r.DB.Where("id = ?", id).First(&ret).Error
	return
}

func (r *ScenarioProcessorRepo) SaveGroup(po *model.ProcessorGroup) (err error) {
	err = r.DB.Save(po).Error

	_ = r.UpdateEntityId(po.ProcessorID, po.ID)
	_ = r.UpdateName(po.ProcessorID, po.Name)

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
	//ret.Id = processor.ID
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

func (r *ScenarioProcessorRepo) CopyLogic(srcId uint) (id uint, err error) {
	logic, err := r.GetLogicById(srcId)
	if err != nil {
		return
	}

	logic.ID = 0
	if err = r.SaveLogic(&logic); err != nil {
		return
	}

	return logic.ID, nil

}

//func (r *ScenarioProcessorRepo) SwitchEntityInterface(id, debugInterFaceId uint) (err error) {
//	processor, _ := r.Get(id)
//	oldDebugInterFaceId := processor.EntityId
//
//	r.DB.Transaction(func(tx *gorm.DB) error {
//		err = r.DebugInterfaceRepo.UpdateProcessorId(debugInterFaceId, id)
//		if err != nil {
//			return err
//		}
//
//		err = r.UpdateEntityId(id, debugInterFaceId)
//		if err != nil {
//			return err
//		}
//
//		err = r.DebugInterfaceRepo.Delete(oldDebugInterFaceId)
//		return err
//	})
//
//	return
//}
