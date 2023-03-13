package repo

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

type PlanCategoryRepo struct {
	DB          *gorm.DB     `inject:""`
	ProjectRepo *ProjectRepo `inject:""`
}

func (r *PlanCategoryRepo) GetTree(projectId uint) (root *v1.PlanCategory, err error) {
	pos, err := r.ListByProject(projectId)
	if err != nil {
		return
	}

	tos := r.toTos(pos)
	if len(tos) == 0 {
		return
	}
	root = tos[0]
	root.Slots = iris.Map{"icon": "icon"}

	r.makeTree(tos[1:], root)

	return
}

func (r *PlanCategoryRepo) ListByProject(projectId uint) (pos []*model.PlanCategory, err error) {
	err = r.DB.
		Where("project_id=?", projectId).
		Where("NOT deleted").
		Order("parent_id ASC, ordr ASC").
		Find(&pos).Error
	return
}

func (r *PlanCategoryRepo) Get(id uint) (po model.PlanCategory, err error) {
	err = r.DB.Where("id = ?", id).First(&po).Error
	return
}

func (r *PlanCategoryRepo) toTos(pos []*model.PlanCategory) (tos []*v1.PlanCategory) {
	for _, po := range pos {
		to := v1.PlanCategory{
			Id:       po.ID,
			Name:     po.Name,
			Desc:     po.Desc,
			ParentId: po.ParentId,
		}

		tos = append(tos, &to)
	}

	return
}

func (r *PlanCategoryRepo) makeTree(findIn []*v1.PlanCategory, parent *v1.PlanCategory) { //参数为父节点，添加父节点的子节点指针切片
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

func (r *PlanCategoryRepo) hasChild(categories []*v1.PlanCategory, parent *v1.PlanCategory) (
	ret []*v1.PlanCategory, yes bool) {

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

func (r *PlanCategoryRepo) Save(processor *model.PlanCategory) (err error) {
	err = r.DB.Save(processor).Error

	return
}

func (r *PlanCategoryRepo) UpdateOrder(pos serverConsts.DropPos, targetId uint) (parentId uint, ordr int) {
	if pos == serverConsts.Inner {
		parentId = targetId

		var preChild model.PlanCategory
		r.DB.Where("parent_id=?", parentId).
			Order("ordr DESC").Limit(1).
			First(&preChild)

		ordr = preChild.Ordr + 1

	} else if pos == serverConsts.Before {
		brother, _ := r.Get(targetId)
		parentId = brother.ParentId

		r.DB.Model(&model.PlanCategory{}).
			Where("NOT deleted AND parent_id=? AND ordr >= ?", parentId, brother.Ordr).
			Update("ordr", gorm.Expr("ordr + 1"))

		ordr = brother.Ordr

	} else if pos == serverConsts.After {
		brother, _ := r.Get(targetId)
		parentId = brother.ParentId

		r.DB.Model(&model.PlanCategory{}).
			Where("NOT deleted AND parent_id=? AND ordr > ?", parentId, brother.Ordr).
			Update("ordr", gorm.Expr("ordr + 1"))

		ordr = brother.Ordr + 1

	}

	return
}

func (r *PlanCategoryRepo) UpdateName(id int, name string) (err error) {
	err = r.DB.Model(&model.PlanCategory{}).
		Where("id = ?", id).
		Update("name", name).Error

	return
}

func (r *PlanCategoryRepo) Update(req v1.PlanCategoryReq) (err error) {
	po := model.PlanCategory{
		Name: req.Name,
		Desc: req.Desc,
	}

	err = r.DB.Save(po).Error

	return
}

func (r *PlanCategoryRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&model.PlanCategory{}).
		Where("id=?", id).
		Update("deleted", true).
		Error

	return
}

func (r *PlanCategoryRepo) GetChildren(nodeId uint) (children []*model.PlanCategory, err error) {
	err = r.DB.Where("parent_id=?", nodeId).Find(&children).Error
	return
}

func (r *PlanCategoryRepo) UpdateOrdAndParent(node model.PlanCategory) (err error) {
	err = r.DB.Model(&node).
		Updates(model.PlanCategory{Ordr: node.Ordr, ParentId: node.ParentId}).
		Error

	return
}

func (r *PlanCategoryRepo) GetMaxOrder(parentId uint) (order int) {
	node := model.PlanCategory{}

	err := r.DB.Model(&model.PlanCategory{}).
		Where("parent_id=?", parentId).
		Order("ordr DESC").
		First(&node).Error

	if err == nil {
		order = node.Ordr + 1
	}

	return
}
