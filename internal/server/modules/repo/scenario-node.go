package repo

import (
	"github.com/aaronchen2k/deeptest/internal/agent/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

type ScenarioNodeRepo struct {
	*BaseRepo `inject:""`

	ScenarioProcessorRepo *ScenarioProcessorRepo `inject:""`
	ScenarioRepo          *ScenarioRepo          `inject:""`
	DebugInterfaceRepo    *DebugInterfaceRepo    `inject:""`
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

func (r *ScenarioNodeRepo) MakeTree(findIn []*agentExec.Processor, parent *agentExec.Processor) { //参数为父节点，添加父节点的子节点指针切片
	children, _ := r.hasChild(findIn, parent) // 判断节点是否有子节点并返回

	if children != nil {
		parent.IsDir = true
		parent.Children = append(parent.Children, children[0:]...) // 添加子节点

		for _, child := range children { // 查询子节点的子节点，并添加到子节点
			_, has := r.hasChild(findIn, child)
			if has {
				r.MakeTree(findIn, child) // 递归添加节点
			}
		}
	}
}

func (r *ScenarioNodeRepo) hasChild(processors []*agentExec.Processor, parent *agentExec.Processor) (
	ret []*agentExec.Processor, yes bool) {

	for _, item := range processors {
		if item.ParentId == parent.ID {
			item.Slots = iris.Map{"icon": "icon"}
			//item.Parent = parent // loop json
			item.Entity = agentExec.ProcessorGroup{}

			ret = append(ret, item)
		}
	}

	if ret != nil {
		yes = true
	}

	return
}

func (r *ScenarioNodeRepo) CreateDefault(scenarioId, projectId, createUserId uint) (po model.Processor, err error) {
	po = model.Processor{
		ScenarioId:     scenarioId,
		ProjectId:      projectId,
		Name:           "root",
		EntityCategory: consts.ProcessorRoot,
		EntityType:     consts.ProcessorRootDefault,
		EntityId:       0,
		CreatedBy:      createUserId,
	}
	err = r.DB.Create(&po).Error

	return
}
func (r *ScenarioNodeRepo) CreateFoldersForPerformance(parentId, scenarioId, projectId, createUserId uint) (err error) {
	runnersPo := model.Processor{
		ParentId:       parentId,
		ScenarioId:     scenarioId,
		ProjectId:      projectId,
		Name:           "执行节点",
		EntityCategory: consts.ProcessorPerformanceRunners,
		EntityType:     consts.ProcessorPerformanceRunnersDefault,
		EntityId:       0,
		CreatedBy:      createUserId,
	}
	err = r.DB.Create(&runnersPo).Error

	scenariosPo := model.Processor{
		ParentId:       parentId,
		ScenarioId:     scenarioId,
		ProjectId:      projectId,
		Name:           "性能场景",
		EntityCategory: consts.ProcessorPerformanceScenarios,
		EntityType:     consts.ProcessorPerformanceScenariosDefault,
		EntityId:       0,
		CreatedBy:      createUserId,
	}
	err = r.DB.Create(&scenariosPo).Error

	return
}

func (r *ScenarioNodeRepo) Save(processor *model.Processor) (err error) {
	err = r.DB.Save(processor).Error

	return
}

func (r *ScenarioNodeRepo) UpdateOrder(pos serverConsts.DropPos, targetId uint) (parentId uint, ordr int, disabled bool) {
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

	parentNode, _ := r.Get(parentId)
	disabled = parentNode.Disabled

	return
}

func (r *ScenarioNodeRepo) UpdateName(id int, name string) (err error) {
	err = r.DB.Model(&model.Processor{}).
		Where("id = ?", id).
		Update("name", name).Error

	return
}

func (r *ScenarioNodeRepo) DeleteWithChildren(id uint) (err error) {
	node, err := r.Get(id)

	ids := []uint{}
	if !r.IsDir(node) {
		ids = append(ids, id)
	} else {
		ids, _ = r.GetAllChildIdsSimple(id, model.Processor{}.TableName())
	}

	err = r.DB.Model(&model.Processor{}).
		Where("id IN (?)", ids).
		Update("deleted", true).Error

	err = r.DebugInterfaceRepo.DeleteByProcessorIds(ids)

	return
}

func (r *ScenarioNodeRepo) DisableWithChildren(id uint) (err error) {
	node, err := r.Get(id)

	action := "disable"
	if node.Disabled {
		action = "enable"
	}

	ids := []uint{}
	if !r.IsDir(node) {
		ids = append(ids, id)
	} else {
		ids, _ = r.GetAllChildIdsSimple(id, model.Processor{}.TableName())
	}

	err = r.DB.Model(&model.Processor{}).
		Where("id IN (?)", ids).
		Update("disabled", !node.Disabled).Error

	if action == "enable" {
		r.EnableAncestors(id)
	}

	return
}

func (r *ScenarioNodeRepo) EnableAncestors(nodeId uint) (err error) {
	ids, err := r.GetAncestorIds(nodeId, model.Processor{}.TableName())

	err = r.DB.Model(&model.Processor{}).
		Where("id IN (?)", ids).
		Update("disabled", false).Error

	return
}

func (r *ScenarioNodeRepo) GetChildren(nodeId uint) (children []*model.Processor, err error) {
	err = r.DB.Where("parent_id=?", nodeId).Find(&children).Error
	return
}

func (r *ScenarioNodeRepo) UpdateOrdAndParent(node model.Processor) (err error) {
	err = r.DB.Model(&node).
		Updates(model.Processor{Ordr: node.Ordr, ParentId: node.ParentId, BaseModel: model.BaseModel{Disabled: node.Disabled}}).
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

func (r *ScenarioNodeRepo) IsDir(po model.Processor) (ret bool) {
	ret = po.EntityCategory == consts.ProcessorRoot ||
		//po.EntityCategory == consts.ProcessorThread ||
		po.EntityCategory == consts.ProcessorGroup ||
		po.EntityCategory == consts.ProcessorLoop ||
		po.EntityCategory == consts.ProcessorLogic ||
		po.EntityCategory == consts.ProcessorData

	return
}

func (r *ScenarioNodeRepo) GetNumberByScenariosAndEntityCategory(scenarioIds []uint, entityCategory consts.ProcessorCategory) (num int64, err error) {
	db := r.DB.Model(model.Processor{}).Where("not deleted and not disabled and scenario_id IN (?)", scenarioIds)
	if entityCategory != "" {
		db.Where("entity_category=?", entityCategory)
	}
	err = db.Count(&num).Error
	return
}

func (r *ScenarioNodeRepo) UpdateEntityId(id, entityId uint) error {
	return r.DB.Model(model.Processor{}).Where("id=?", id).Update("entity_id", entityId).Error
}

func (r *ScenarioNodeRepo) MoveMaxOrder(parentId, order, step uint) (err error) {
	return r.DB.Model(model.Processor{}).Where("parent_id=? and ordr>?", parentId, order).Update("ordr", gorm.Expr("ordr + ?", step)).Error
}

func (r *ScenarioNodeRepo) GetNextNode(id uint) (processor model.Processor, err error) {
	node, err := r.Get(id)
	if err != nil {
		return
	}

	err = r.DB.Where("parent_id = ? and ordr > ?", node.ParentId, node.Ordr).Order("ordr ASC").First(&processor).Error

	return

}

func (r *ScenarioNodeRepo) GetPreNode(id uint) (processor model.Processor, err error) {
	node, err := r.Get(id)
	if err != nil {
		return
	}

	err = r.DB.Where("parent_id = ? and ordr < ?", node.ParentId, node.Ordr).Order("ordr DESC").First(&processor).Error

	return

}
