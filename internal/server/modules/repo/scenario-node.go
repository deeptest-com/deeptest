package repo

import (
	agentDomain "github.com/aaronchen2k/deeptest/internal/agent/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/jinzhu/copier"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

type ScenarioNodeRepo struct {
	DB                    *gorm.DB               `inject:""`
	ScenarioProcessorRepo *ScenarioProcessorRepo `inject:""`
	ScenarioRepo          *ScenarioRepo          `inject:""`
}

//func (r *ScenarioNodeRepo) GenTestScenario(scenarioId uint) (ret *run.TestScenario, err error) {
//	rootProcessor, _ := r.GetTree(scenarioId)
//
//	rootStage := run.TStage{
//		Id:   rootProcessor.ID,
//		Name: rootProcessor.Name,
//	}
//
//	r.getStageTree(*rootProcessor, &rootStage)
//
//	runDomain := runDomain.ProcessorRootStage{
//		Stage: &rootStage,
//	}
//
//	po, err := r.Get(scenarioId)
//	ret = &run.TestScenario {
//		Id: po.ID,
//		Name: po.Name,
//	}
//	ret.TestStages = append(ret.TestStages, &runDomain)
//
//	return
//}
//func (r *ScenarioNodeRepo) getStageTree(parentProcessor model.Processor, parentStage *run.TStage) {
//	r.GetStage(parentProcessor, parentStage)
//
//	for _, processor := range (parentProcessor).Children {
//		r.GetStage(*processor, parentStage)
//
//		for _, child := range processor.Children {
//			childStage := run.TStage{
//				Id:   child.ID,
//				Name: child.Name,
//			}
//
//			parentStage.Children = parentStage.Children
//
//			r.getStageTree(*child, &childStage)
//		}
//	}
//}

func (r *ScenarioNodeRepo) GetTree(scenarioId uint) (root *agentDomain.Processor, err error) {
	scenario, err := r.ScenarioRepo.Get(scenarioId)

	pos, err := r.ListByScenario(scenarioId)
	if err != nil {
		return
	}

	tos := r.toTos(pos)

	root = tos[0]
	root.Name = scenario.Name
	root.Slots = iris.Map{"icon": "icon"}
	r.makeTree(tos[1:], root)

	return
}

func (r *ScenarioNodeRepo) ListByScenario(scenarioId uint) (pos []*model.Processor, err error) {
	err = r.DB.
		Where("scenario_id=?", scenarioId).
		Where("NOT deleted").
		Order("parent_id ASC, ordr ASC").
		Find(&pos).Error
	return
}

func (r *ScenarioNodeRepo) Get(id uint) (processor model.Processor, err error) {
	err = r.DB.Where("id = ?", id).First(&processor).Error
	return
}

func (r *ScenarioNodeRepo) toTos(pos []*model.Processor) (tos []*agentDomain.Processor) {
	for _, po := range pos {
		to := agentDomain.Processor{}
		copier.CopyWithOption(&to, po, copier.Option{DeepCopy: true})
		tos = append(tos, &to)
	}

	return
}

func (r *ScenarioNodeRepo) makeTree(findIn []*agentDomain.Processor, parent *agentDomain.Processor) { //参数为父节点，添加父节点的子节点指针切片
	children, _ := r.hasChild(findIn, parent) // 判断节点是否有子节点并返回

	if children != nil {
		parent.Children = append(parent.Children, children[0:]...) // 添加子节点

		for _, child := range children {                           // 查询子节点的子节点，并添加到子节点
			_, has := r.hasChild(findIn, child)
			if has {
				r.makeTree(findIn, child) // 递归添加节点
			}
		}
	}
}

func (r *ScenarioNodeRepo) hasChild(processors []*agentDomain.Processor, node *agentDomain.Processor) (
		ret []*agentDomain.Processor, yes bool) {
	for _, item := range processors {
		if item.ParentId == node.ID {
			item.Slots = iris.Map{"icon": "icon"}
			ret = append(ret, item)
		}
	}

	if ret != nil {
		yes = true
	}

	return
}

func (r *ScenarioNodeRepo) CreateDefault(scenarioId uint) (po model.Processor, err error) {
	po = model.Processor{
		ScenarioId:     scenarioId,
		Name:           "root",
		EntityCategory: consts.ProcessorRoot,
		EntityId:       0,
	}
	err = r.DB.Create(&po).Error

	return
}

func (r *ScenarioNodeRepo) Save(processor *model.Processor) (err error) {
	err = r.DB.Save(processor).Error

	return
}

func (r *ScenarioNodeRepo) UpdateOrder(pos serverConsts.DropPos, targetId uint) (parentId uint, ordr int) {
	if pos == serverConsts.Inner {
		parentId = targetId

		var preChild model.Processor
		r.DB.Where("parent_id=?", parentId).
			Order("ordr DESC").Limit(1).
			First(&preChild)

		ordr = preChild.Ordr + 1

	} else if pos == serverConsts.Before {
		brother, _ := r.Get(targetId)
		parentId = brother.ParentId

		r.DB.Model(&model.Processor{}).
			Where("NOT deleted AND parent_id=? AND ordr >= ?", parentId, brother.Ordr).
			Update("ordr", gorm.Expr("ordr + 1"))

		ordr = brother.Ordr

	} else if pos == serverConsts.After {
		brother, _ := r.Get(targetId)
		parentId = brother.ParentId

		r.DB.Model(&model.Processor{}).
			Where("NOT deleted AND parent_id=? AND ordr > ?", parentId, brother.Ordr).
			Update("ordr", gorm.Expr("ordr + 1"))

		ordr = brother.Ordr + 1

	}

	return
}

func (r *ScenarioNodeRepo) UpdateName(id int, name string) (err error) {
	err = r.DB.Model(&model.Processor{}).
		Where("id = ?", id).
		Update("name", name).Error

	return
}

func (r *ScenarioNodeRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&model.Processor{}).
		Where("id=?", id).
		Update("deleted", true).
		Error

	return
}

func (r *ScenarioNodeRepo) GetChildren(nodeId uint) (children []*model.Processor, err error) {
	err = r.DB.Where("parent_id=?", nodeId).Find(&children).Error
	return
}

func (r *ScenarioNodeRepo) UpdateOrdAndParent(node model.Processor) (err error) {
	err = r.DB.Model(&node).
		Updates(model.Processor{Ordr: node.Ordr, ParentId: node.ParentId}).
		Error

	return
}

func (r *ScenarioNodeRepo) GetMaxOrder(parentId uint) (order int) {
	node := model.Processor{}

	err := r.DB.Model(&model.Processor{}).
		Where("parent_id=?", parentId).
		Order("ordr DESC").
		First(&node).Error

	if err == nil {
		order = node.Ordr + 1
	}

	return
}

func (r *ScenarioNodeRepo) GetScopeHierarchy(scenarioId uint, scopeHierarchyMap *map[uint]*[]uint) {
	processors, err := r.ListByScenario(scenarioId)
	if err != nil {
		return
	}

	childToParentIdMap := map[uint]uint{}
	for _, processor := range processors {
		childToParentIdMap[processor.ID] = processor.ParentId
	}

	for childId, parentId := range childToParentIdMap {
		if (*scopeHierarchyMap)[childId] == nil {
			arr := []uint{childId}
			(*scopeHierarchyMap)[childId] = &arr
		}
		*(*scopeHierarchyMap)[childId] = append(*(*scopeHierarchyMap)[childId], parentId)

		r.addSuperParent(childId, parentId, childToParentIdMap, scopeHierarchyMap)
	}
}

func (r *ScenarioNodeRepo) addSuperParent(id, parentId uint, childToParentIdMap map[uint]uint, scopeHierarchyMap *map[uint]*[]uint) {
	superId, ok := childToParentIdMap[parentId]
	if ok {
		*(*scopeHierarchyMap)[id] = append(*(*scopeHierarchyMap)[id], superId)

		r.addSuperParent(id, superId, childToParentIdMap, scopeHierarchyMap)
	}
}

//func (r *ScenarioNodeRepo) GetStage(processor model.Processor, parentStage *run.TStage) (err error) {
//	if processor.EntityCategory == consts.ProcessorInterface {
//		stage, _ := r.getInterfaceStage(processor)
//		parentStage.Children = append(parentStage.Children, &stage)
//
//	} else if processor.EntityCategory == consts.ProcessorGroup {
//		stage, _ := r.getGroupStage(processor)
//		parentStage.Children = append(parentStage.Children, &stage)
//
//	} else if processor.EntityCategory == consts.ProcessorLogic {
//		stage, _ := r.getLogicStage(processor)
//		parentStage.Children = append(parentStage.Children, &stage)
//
//	} else if processor.EntityCategory == consts.ProcessorLoop {
//		stage, _ := r.getLoopStage(processor)
//		parentStage.Children = append(parentStage.Children, &stage)
//
//	} else if processor.EntityCategory == consts.ProcessorVariable {
//		stage, _ := r.getVariableStage(processor)
//		parentStage.Children = append(parentStage.Children, &stage)
//
//	} else if processor.EntityCategory == consts.ProcessorTimer {
//		stage, _ := r.getTimerStage(processor)
//		parentStage.Children = append(parentStage.Children, &stage)
//
//	} else if processor.EntityCategory == consts.ProcessorPrint {
//		stage, _ := r.getPrintStage(processor)
//		parentStage.Children = append(parentStage.Children, &stage)
//
//	} else if processor.EntityCategory == consts.ProcessorCookie {
//		stage, _ := r.getCookieStage(processor)
//		parentStage.Children = append(parentStage.Children, &stage)
//
//	} else if processor.EntityCategory == consts.ProcessorAssertion {
//		stage, _ := r.getAssertionStage(processor)
//		parentStage.Children = append(parentStage.Children, &stage)
//
//	} else if processor.EntityCategory == consts.ProcessorExtractor {
//		stage, _ := r.getExtractorStage(processor)
//		parentStage.Children = append(parentStage.Children, &stage)
//
//	} else if processor.EntityCategory == consts.ProcessorData {
//		stage, _ := r.getDataStage(processor)
//		parentStage.Children = append(parentStage.Children, &stage)
//
//	}
//
//	return
//}
//func (r *ScenarioNodeRepo) getInterfaceStage(processor model.Processor) (ret runDomain.ProcessorInterfaceStage, err error) {
//	po, _ := r.ScenarioProcessorRepo.GetInterface(processor)
//	processorInterface := runDomain.ProcessorInterface{}
//
//	copier.CopyWithOption(&processorInterface, po, copier.Option{DeepCopy: true})
//	ret = runDomain.ProcessorInterfaceStage{
//		Stage: &run.TStage{
//			Processor: processorInterface,
//		},
//	}
//
//	return
//}
//func (r *ScenarioNodeRepo) getGroupStage(processor model.Processor) (ret runDomain.ProcessorGroupStage, err error) {
//	po, _ := r.ScenarioProcessorRepo.GetGroup(processor)
//	processorGroup := runDomain.ProcessorGroup{}
//
//	copier.CopyWithOption(&processorGroup, po, copier.Option{DeepCopy: true})
//	ret = runDomain.ProcessorGroupStage{
//		Stage: &run.TStage{
//			Processor: processorGroup,
//		},
//	}
//
//	return
//}
//func (r *ScenarioNodeRepo) getLogicStage(processor model.Processor) (ret runDomain.ProcessorLogicStage, err error) {
//	po, _ := r.ScenarioProcessorRepo.GetLogic(processor)
//	processorLogic := runDomain.ProcessorLogic{}
//
//	copier.CopyWithOption(&processorLogic, po, copier.Option{DeepCopy: true})
//	ret = runDomain.ProcessorLogicStage{
//		Stage: &run.TStage{
//			Processor: processorLogic,
//		},
//	}
//
//	return
//}
//func (r *ScenarioNodeRepo) getLoopStage(processor model.Processor) (ret runDomain.ProcessorInterfaceStage, err error) {
//	po, _ := r.ScenarioProcessorRepo.GetInterface(processor)
//	processorInterface := runDomain.ProcessorInterface{}
//
//	copier.CopyWithOption(&processorInterface, po, copier.Option{DeepCopy: true})
//	ret = runDomain.ProcessorInterfaceStage{
//		Stage: &run.TStage{
//			Processor: processorInterface,
//		},
//	}
//
//	return
//}
//func (r *ScenarioNodeRepo) getVariableStage(processor model.Processor) (ret runDomain.ProcessorVariableStage, err error) {
//	po, _ := r.ScenarioProcessorRepo.GetVariable(processor)
//	processorVariable := runDomain.ProcessorVariable{}
//
//	copier.CopyWithOption(&processorVariable, po, copier.Option{DeepCopy: true})
//	ret = runDomain.ProcessorVariableStage{
//		Stage: &run.TStage{
//			Processor: processorVariable,
//		},
//	}
//
//	return
//}
//func (r *ScenarioNodeRepo) getTimerStage(processor model.Processor) (ret runDomain.ProcessorTimerStage, err error) {
//	po, _ := r.ScenarioProcessorRepo.GetTimer(processor)
//	processorTimer := runDomain.ProcessorTimer{}
//
//	copier.CopyWithOption(&processorTimer, po, copier.Option{DeepCopy: true})
//	ret = runDomain.ProcessorTimerStage{
//		Stage: &run.TStage{
//			Processor: processorTimer,
//		},
//	}
//
//	return
//}
//func (r *ScenarioNodeRepo) getPrintStage(processor model.Processor) (ret runDomain.ProcessorPrintStage, err error) {
//	po, _ := r.ScenarioProcessorRepo.GetPrint(processor)
//	processorPrint := runDomain.ProcessorPrint{}
//
//	copier.CopyWithOption(&processorPrint, po, copier.Option{DeepCopy: true})
//	ret = runDomain.ProcessorPrintStage{
//		Stage: &run.TStage{
//			Processor: processorPrint,
//		},
//	}
//
//	return
//}
//func (r *ScenarioNodeRepo) getCookieStage(processor model.Processor) (ret runDomain.ProcessorCookieStage, err error) {
//	po, _ := r.ScenarioProcessorRepo.GetCookie(processor)
//	processorCookie := runDomain.ProcessorCookie{}
//
//	copier.CopyWithOption(&processorCookie, po, copier.Option{DeepCopy: true})
//	ret = runDomain.ProcessorCookieStage{
//		Stage: &run.TStage{
//			Processor: processorCookie,
//		},
//	}
//
//	return
//}
//func (r *ScenarioNodeRepo) getAssertionStage(processor model.Processor) (ret runDomain.ProcessorAssertionStage, err error) {
//	po, _ := r.ScenarioProcessorRepo.GetAssertion(processor)
//	processorAssertion := runDomain.ProcessorAssertion{}
//
//	copier.CopyWithOption(&processorAssertion, po, copier.Option{DeepCopy: true})
//	ret = runDomain.ProcessorAssertionStage{
//		Stage: &run.TStage{
//			Processor: processorAssertion,
//		},
//	}
//
//	return
//}
//func (r *ScenarioNodeRepo) getExtractorStage(processor model.Processor) (ret runDomain.ProcessorExtractorStage, err error) {
//	po, _ := r.ScenarioProcessorRepo.GetExtractor(processor)
//	processorExtractor := runDomain.ProcessorExtractor{}
//
//	copier.CopyWithOption(&processorExtractor, po, copier.Option{DeepCopy: true})
//	ret = runDomain.ProcessorExtractorStage{
//		Stage: &run.TStage{
//			Processor: processorExtractor,
//		},
//	}
//
//	return
//}
//func (r *ScenarioNodeRepo) getDataStage(processor model.Processor) (ret runDomain.ProcessorDataStage, err error) {
//	po, _ := r.ScenarioProcessorRepo.GetData(processor)
//	processorData := runDomain.ProcessorData{}
//
//	copier.CopyWithOption(&processorData, po, copier.Option{DeepCopy: true})
//	ret = runDomain.ProcessorDataStage{
//		Stage: &run.TStage{
//			Processor: processorData,
//		},
//	}
//
//	return
//}
