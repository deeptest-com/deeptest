package repo

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

type ScenarioCategoryRepo struct {
	DB          *gorm.DB     `inject:""`
	ProjectRepo *ProjectRepo `inject:""`
}

func (r *ScenarioCategoryRepo) GetTree(projectId uint) (root *v1.ScenarioCategory, err error) {
	project, _ := r.ProjectRepo.Get(projectId)

	pos, err := r.ListByProject(projectId)
	if err != nil {
		return
	}

	tos := r.toTos(pos)

	root = tos[0]
	root.Name = project.Name
	root.Slots = iris.Map{"icon": "icon"}

	r.makeTree(tos[1:], root)

	return
}

func (r *ScenarioCategoryRepo) ListByProject(projectId uint) (pos []*model.ScenarioCategory, err error) {
	err = r.DB.
		Where("project_id=?", projectId).
		Where("NOT deleted").
		Order("parent_id ASC, ordr ASC").
		Find(&pos).Error
	return
}

func (r *ScenarioCategoryRepo) Get(id uint) (processor model.ScenarioCategory, err error) {
	err = r.DB.Where("id = ?", id).First(&processor).Error
	return
}

func (r *ScenarioCategoryRepo) toTos(pos []*model.ScenarioCategory) (tos []*v1.ScenarioCategory) {
	for _, po := range pos {
		to := v1.ScenarioCategory{
			Id:   po.ID,
			Name: po.Name,
			Desc: po.Desc,
		}

		tos = append(tos, &to)
	}

	return
}

func (r *ScenarioCategoryRepo) makeTree(findIn []*v1.ScenarioCategory, parent *v1.ScenarioCategory) { //参数为父节点，添加父节点的子节点指针切片
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

func (r *ScenarioCategoryRepo) hasChild(categories []*v1.ScenarioCategory, parent *v1.ScenarioCategory) (
	ret []*v1.ScenarioCategory, yes bool) {

	for _, item := range categories {
		if item.ParentId == parent.Id {
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

func (r *ScenarioCategoryRepo) CreateDefault(scenarioId uint) (po model.Processor, err error) {
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

func (r *ScenarioCategoryRepo) Save(processor *model.ScenarioCategory) (err error) {
	err = r.DB.Save(processor).Error

	return
}

func (r *ScenarioCategoryRepo) UpdateOrder(pos serverConsts.DropPos, targetId uint) (parentId uint, ordr int) {
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

func (r *ScenarioCategoryRepo) UpdateName(id int, name string) (err error) {
	err = r.DB.Model(&model.Processor{}).
		Where("id = ?", id).
		Update("name", name).Error

	return
}

func (r *ScenarioCategoryRepo) Update(req v1.ScenarioCategoryReq) (err error) {
	po := model.ScenarioCategory{
		Name: req.Name,
		Desc: req.Desc,
	}

	err = r.DB.Save(po).Error

	return
}

func (r *ScenarioCategoryRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&model.Processor{}).
		Where("id=?", id).
		Update("deleted", true).
		Error

	return
}

func (r *ScenarioCategoryRepo) GetChildren(nodeId uint) (children []*model.Processor, err error) {
	err = r.DB.Where("parent_id=?", nodeId).Find(&children).Error
	return
}

func (r *ScenarioCategoryRepo) UpdateOrdAndParent(node model.ScenarioCategory) (err error) {
	err = r.DB.Model(&node).
		Updates(model.Processor{Ordr: node.Ordr, ParentId: node.ParentId}).
		Error

	return
}

func (r *ScenarioCategoryRepo) GetMaxOrder(parentId uint) (order int) {
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

func (r *ScenarioCategoryRepo) GetScopeHierarchy(scenarioId uint, scopeHierarchyMap *map[uint]*[]uint) {
	processors, err := r.ListByProject(scenarioId)
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

func (r *ScenarioCategoryRepo) addSuperParent(id, parentId uint, childToParentIdMap map[uint]uint, scopeHierarchyMap *map[uint]*[]uint) {
	superId, ok := childToParentIdMap[parentId]
	if ok {
		*(*scopeHierarchyMap)[id] = append(*(*scopeHierarchyMap)[id], superId)

		r.addSuperParent(id, superId, childToParentIdMap, scopeHierarchyMap)
	}
}

func (r *ScenarioCategoryRepo) IsDir(po model.Processor) (ret bool) {
	ret = po.EntityCategory == consts.ProcessorRoot ||
		//po.EntityCategory == consts.ProcessorThread ||
		po.EntityCategory == consts.ProcessorGroup ||
		po.EntityCategory == consts.ProcessorLoop ||
		po.EntityCategory == consts.ProcessorLogic ||
		po.EntityCategory == consts.ProcessorData

	return
}
