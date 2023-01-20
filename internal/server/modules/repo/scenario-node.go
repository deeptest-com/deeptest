package repo

import (
	"encoding/json"
	"github.com/aaronchen2k/deeptest/internal/agent/exec"
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

func (r *ScenarioNodeRepo) GetTree(scenarioId uint, withEntity bool) (root *agentExec.Processor, err error) {
	scenario, err := r.ScenarioRepo.Get(scenarioId)

	pos, err := r.ListByScenario(scenarioId)
	if err != nil {
		return
	}

	tos := r.toTos(pos, withEntity)

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

func (r *ScenarioNodeRepo) toTos(pos []*model.Processor, withDetail bool) (tos []*agentExec.Processor) {
	for _, po := range pos {
		to := agentExec.Processor{
			IsLeaf: r.IsLeaf(*po),
		}
		copier.CopyWithOption(&to, po, copier.Option{DeepCopy: true})

		if withDetail {
			entity, _ := r.ScenarioProcessorRepo.GetEntityTo(&to)

			// avoid json serialization error
			to.EntityRaw, _ = json.Marshal(entity)
			to.Entity = &agentExec.ProcessorGroup{}

		} else {
			to.Entity = &agentExec.ProcessorGroup{} // just to avoid json marshal error for IProcessorEntity
		}

		tos = append(tos, &to)
	}

	return
}

func (r *ScenarioNodeRepo) makeTree(findIn []*agentExec.Processor, parent *agentExec.Processor) { //参数为父节点，添加父节点的子节点指针切片
	children, _ := r.hasChild(findIn, parent) // 判断节点是否有子节点并返回

	if children != nil {
		parent.Children = append(parent.Children, children[0:]...) // 添加子节点

		for _, child := range children { // 查询子节点的子节点，并添加到子节点
			_, has := r.hasChild(findIn, child)
			if has {
				r.makeTree(findIn, child) // 递归添加节点
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
		EntityType:     consts.ProcessorRootDefault,
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

func (r *ScenarioNodeRepo) IsLeaf(po model.Processor) (ret bool) {
	isDir := po.EntityCategory == consts.ProcessorRoot ||
		//po.EntityCategory == consts.ProcessorThread ||
		po.EntityCategory == consts.ProcessorGroup ||
		po.EntityCategory == consts.ProcessorLoop ||
		po.EntityCategory == consts.ProcessorLogic ||
		po.EntityCategory == consts.ProcessorData

	ret = !isDir

	return
}
