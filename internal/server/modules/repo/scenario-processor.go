package repo

import (
	domain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	_intUtils "github.com/aaronchen2k/deeptest/pkg/lib/int"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type ScenarioProcessorRepo struct {
	DB *gorm.DB `inject:""`

	ScenarioNodeRepo      *ScenarioNodeRepo      `inject:""`
	ExtractorRepo         *ExtractorRepo         `inject:""`
	CheckpointRepo        *CheckpointRepo        `inject:""`
	DebugInterfaceRepo    *DebugInterfaceRepo    `inject:""`
	PerformanceRunnerRepo *PerformanceRunnerRepo `inject:""`
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

	case consts.ProcessorPerformanceGoal:
		ret, _ = r.GetPerformanceGoal(processor)
	case consts.ProcessorPerformanceScenario:
		ret, _ = r.GetPerformanceScenario(processor)
	case consts.ProcessorPerformanceRendezvous:
		ret, _ = r.GetPerformanceRendezvous(processor)

	default:
	}

	return
}

func (r *ScenarioProcessorRepo) CopyEntity(srcProcessorId, distProcessorId uint) (err error) {
	distProcessor, err := r.Get(distProcessorId)
	if err != nil {
		return err
	}
	distParentId := distProcessor.ParentId

	entity, err := r.GetEntity(srcProcessorId)
	if err != nil {
		return err
	}

	switch entity.(type) {
	case model.ProcessorGroup:
		group := entity.(model.ProcessorGroup)
		group.ID = 0
		group.ProcessorID = distProcessorId
		group.ParentID = distParentId
		err = r.SaveGroup(&group)
	case model.ProcessorLogic:
		logic := entity.(model.ProcessorLogic)
		logic.ID = 0
		logic.ProcessorID = distProcessorId
		logic.ParentID = distParentId
		err = r.SaveLogic(&logic)
	case model.ProcessorLoop:
		loop := entity.(model.ProcessorLoop)
		loop.ID = 0
		loop.ProcessorID = distProcessorId
		loop.ParentID = distParentId
		err = r.SaveLoop(&loop)
	case model.ProcessorTimer:
		timer := entity.(model.ProcessorTimer)
		timer.ID = 0
		timer.ProcessorID = distProcessorId
		timer.ParentID = distParentId
		err = r.SaveTimer(&timer)
	case model.ProcessorPrint:
		printData := entity.(model.ProcessorPrint)
		printData.ID = 0
		printData.ProcessorID = distProcessorId
		printData.ParentID = distParentId
		err = r.SavePrint(&printData)
	case model.ProcessorVariable:
		variable := entity.(model.ProcessorVariable)
		variable.ID = 0
		variable.ProcessorID = distProcessorId
		variable.ParentID = distParentId
		err = r.SaveVariable(&variable)
	case model.ProcessorAssertion:
		assertion := entity.(model.ProcessorAssertion)
		assertion.ID = 0
		assertion.ProcessorID = distProcessorId
		assertion.ParentID = distParentId
		err = r.SaveAssertion(&assertion)
	case model.ProcessorData:
		data := entity.(model.ProcessorData)
		data.ID = 0
		data.ProcessorID = distProcessorId
		data.ParentID = distParentId
		err = r.SaveData(&data)
	case model.ProcessorCookie:
		cookie := entity.(model.ProcessorCookie)
		cookie.ID = 0
		cookie.ProcessorID = distProcessorId
		cookie.ParentID = distParentId
		err = r.SaveCookie(&cookie)
	case model.ProcessorCustomCode:
		customCode := entity.(model.ProcessorCustomCode)
		customCode.ID = 0
		customCode.ProcessorID = distProcessorId
		customCode.ParentID = distParentId
		err = r.SaveCustomCode(&customCode)
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

func (r *ScenarioProcessorRepo) GetPerformanceGoal(processor model.Processor) (ret model.ProcessorPerformanceGoal, err error) {
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
func (r *ScenarioProcessorRepo) GetPerformanceScenario(processor model.Processor) (ret model.ProcessorPerformanceScenario, err error) {
	err = r.DB.Where("processor_id = ?", processor.ID).First(&ret).Error

	if ret.ID == 0 {
		comm := r.genProcessorComm(processor)
		copier.CopyWithOption(&ret, comm, copier.Option{DeepCopy: true})

		return
	}

	ret.Name = processor.Name
	ret.ParentID = processor.ParentId

	ret.Stages, err = r.loadStages(ret.ID)

	idsExist := r.PerformanceRunnerRepo.ListExistOnes(ret.RunnerIdsRaw)

	for _, str := range strings.Split(ret.RunnerIdsRaw, ",") {
		i, err := strconv.Atoi(str)
		if err != nil {
			continue
		}

		if _intUtils.FindInArr(i, idsExist) {
			ret.RunnerIds = append(ret.RunnerIds, i)
		}
	}

	return
}
func (r *ScenarioProcessorRepo) GetPerformanceRendezvous(processor model.Processor) (ret model.ProcessorPerformanceRendezvous, err error) {
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

func (r *ScenarioProcessorRepo) SavePerformanceGoal(po *model.ProcessorPerformanceGoal) (err error) {
	err = r.DB.Save(po).Error
	r.UpdateEntityId(po.ProcessorID, po.ID)

	return
}

func (r *ScenarioProcessorRepo) SavePerformanceScenario(po *model.ProcessorPerformanceScenario) (err error) {
	strArr := make([]string, len(po.RunnerIds))
	for i, v := range po.RunnerIds {
		strArr[i] = strconv.Itoa(v)
	}
	po.RunnerIdsRaw = strings.Join(strArr, ",")

	err = r.DB.Save(po).Error
	r.UpdateEntityId(po.ProcessorID, po.ID)

	r.saveStages(po)

	po.Stages, err = r.loadStages(po.ID)

	return
}

func (r *ScenarioProcessorRepo) SavePerformanceRendezvous(po *model.ProcessorPerformanceRendezvous) (err error) {
	err = r.DB.Save(po).Error
	r.UpdateEntityId(po.ProcessorID, po.ID)

	return
}

func (r *ScenarioProcessorRepo) saveStages(po *model.ProcessorPerformanceScenario) (err error) {
	err = r.removeStages(po.ID)

	var stages []model.ProcessorPerformanceStage
	for _, item := range po.Stages {
		if item.Duration == 0 || item.Target == 0 {
			continue
		}

		stages = append(stages, model.ProcessorPerformanceStage{
			Duration:   item.Duration,
			Target:     item.Target,
			ScenarioId: po.ID,
		})
	}
	err = r.DB.Create(&stages).Error

	return
}
func (r *ScenarioProcessorRepo) removeStages(scenarioId uint) (err error) {
	err = r.DB.
		Where("scenario_id = ?", scenarioId).
		Delete(&model.ProcessorPerformanceStage{}).Error

	return
}
func (r *ScenarioProcessorRepo) loadStages(scenarioId uint) (ret []model.ProcessorPerformanceStage, err error) {
	err = r.DB.Where("scenario_id = ?", scenarioId).Find(&ret).Error

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
