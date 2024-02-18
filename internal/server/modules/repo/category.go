package repo

import (
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

type CategoryRepo struct {
	*BaseRepo   `inject:""`
	DB          *gorm.DB     `inject:""`
	ProjectRepo *ProjectRepo `inject:""`
}

func (r *CategoryRepo) GetTree(tenantId consts.TenantId, typ serverConsts.CategoryDiscriminator, projectId uint) (root *v1.Category, err error) {
	pos, err := r.ListByProject(tenantId, typ, projectId)
	if err != nil {
		return
	}

	tos := r.toTos(pos)
	if len(tos) == 0 {
		return
	}

	root = tos[0]
	root.Slots = iris.Map{"icon": "icon"}

	r.makeTree(tenantId, tos[1:], root)

	return
}

func (r *CategoryRepo) ListByProject(tenantId consts.TenantId, typ serverConsts.CategoryDiscriminator, projectId uint) (pos []*model.Category, err error) {
	db := r.GetDB(tenantId).
		Where("project_id=?", projectId).
		Where("type=?", typ).
		Where("NOT deleted")

	err = db.
		Order("parent_id ASC, ordr ASC").
		Find(&pos).Error

	return
}

func (r *CategoryRepo) Get(tenantId consts.TenantId, id int) (po model.Category, err error) {
	err = r.GetDB(tenantId).Where("id = ?", id).First(&po).Error
	return
}

func (r *CategoryRepo) toTos(pos []*model.Category) (tos []*v1.Category) {
	for _, po := range pos {
		to := v1.Category{
			Id:       int64(po.ID),
			Name:     po.Name,
			Desc:     po.Desc,
			ParentId: int64(po.ParentId),
		}

		tos = append(tos, &to)
	}

	return
}

func (r *CategoryRepo) makeTree(tenantId consts.TenantId, findIn []*v1.Category, parent *v1.Category) { //参数为父节点，添加父节点的子节点指针切片
	children, _ := r.hasChild(tenantId, findIn, parent) // 判断节点是否有子节点并返回

	if children != nil {
		parent.Children = append(parent.Children, children[0:]...) // 添加子节点

		for _, child := range children { // 查询子节点的子节点，并添加到子节点
			_, has := r.hasChild(tenantId, findIn, child)
			if has {
				r.makeTree(tenantId, findIn, child) // 递归添加节点
			}
		}
	}
}

func (r *CategoryRepo) hasChild(tenantId consts.TenantId, categories []*v1.Category, parent *v1.Category) (
	ret []*v1.Category, yes bool) {

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

func (r *CategoryRepo) Save(tenantId consts.TenantId, category *model.Category) (err error) {
	if category != nil && category.ID == 0 && category.ParentId != 0 && category.Ordr == 0 {
		category.Ordr = r.GetMaxOrder(tenantId, uint(category.ParentId), category.Type, category.ProjectId)
	}
	err = r.GetDB(tenantId).Save(category).Error

	return
}

func (r *CategoryRepo) UpdateOrder(tenantId consts.TenantId, pos serverConsts.DropPos, targetId int, typ serverConsts.CategoryDiscriminator, projectId uint) (
	parentId int, ordr int) {
	if pos == serverConsts.Inner {
		parentId = targetId

		var preChild model.Category
		r.GetDB(tenantId).Where("parent_id=? AND type = ? AND project_id = ?", parentId, typ, projectId).
			Order("ordr DESC").Limit(1).
			First(&preChild)

		ordr = preChild.Ordr + 1

	} else if pos == serverConsts.Before {
		brother, _ := r.Get(tenantId, targetId)
		ordr = brother.Ordr
		parentId = brother.ParentId

		r.GetDB(tenantId).Model(&model.Category{}).
			Where("NOT deleted AND parent_id=? AND type = ? AND project_id = ? AND ordr >= ?",
				parentId, typ, projectId, brother.Ordr).
			Update("ordr", gorm.Expr("ordr + 1"))

	} else if pos == serverConsts.After {
		brother, _ := r.Get(tenantId, targetId)
		parentId = brother.ParentId

		r.GetDB(tenantId).Model(&model.Category{}).
			Where("NOT deleted AND parent_id=? AND type = ? AND project_id = ? AND ordr > ?",
				parentId, typ, projectId, brother.Ordr).
			Update("ordr", gorm.Expr("ordr + 1"))

		ordr = brother.Ordr + 1

	}

	return
}

func (r *CategoryRepo) UpdateName(tenantId consts.TenantId, id int, name string) (err error) {
	err = r.GetDB(tenantId).Model(&model.Category{}).
		Where("id = ?", id).
		Update("name", name).Error

	return
}

func (r *CategoryRepo) Update(tenantId consts.TenantId, req v1.CategoryReq) (err error) {
	po := new(model.Category)
	po.ID = uint(req.Id)

	err = r.GetDB(tenantId).First(&po).Error
	if err != nil {
		return err
	}

	po.Name = req.Name
	po.Desc = req.Desc

	err = r.GetDB(tenantId).Save(&po).Error

	return
}

func (r *CategoryRepo) Delete(tenantId consts.TenantId, id uint) (err error) {
	err = r.GetDB(tenantId).Model(&model.Category{}).
		Where("id=?", id).
		Update("deleted", true).
		Error

	return
}

func (r *CategoryRepo) BatchDelete(tenantId consts.TenantId, ids []uint) (err error) {
	err = r.GetDB(tenantId).Model(&model.Category{}).
		Where("id IN (?)", ids).
		Update("deleted", true).
		Error

	return
}
func (r *CategoryRepo) GetChildren(tenantId consts.TenantId, nodeId uint) (children []*model.Category, err error) {
	err = r.GetDB(tenantId).Where("parent_id=?", nodeId).Find(&children).Error
	return
}

func (r *CategoryRepo) UpdateOrdAndParent(tenantId consts.TenantId, node model.Category) (err error) {
	err = r.GetDB(tenantId).Model(&node).
		Updates(model.Category{Ordr: node.Ordr, ParentId: node.ParentId}).
		Error

	return
}

func (r *CategoryRepo) GetMaxOrder(tenantId consts.TenantId, parentId uint, typ serverConsts.CategoryDiscriminator, projectId uint) (order int) {
	node := model.Category{}

	err := r.GetDB(tenantId).Model(&model.Category{}).
		Where("parent_id=? AND type = ? AND project_id = ?", parentId, typ, projectId).
		Order("ordr DESC").
		First(&node).Error

	if err == nil {
		order = node.Ordr + 1
	}

	return
}

func (r *CategoryRepo) GetByItem(tenantId consts.TenantId, parentId uint, typ serverConsts.CategoryDiscriminator, projectId uint, name string) (res model.Category, err error) {

	err = r.GetDB(tenantId).Model(&model.Category{}).
		Where("parent_id=? AND type = ? AND project_id = ? AND name = ? and not deleted", parentId, typ, projectId, name).
		Order("ordr DESC").
		First(&res).Error

	return

}

func (r *CategoryRepo) GetDetail(tenantId consts.TenantId, req model.Category) (res model.Category, err error) {
	coon := r.GetDB(tenantId).Model(&model.Category{}).Where("not deleted")
	if req.Name != "" {
		coon = coon.Where("name = ?", req.Name)
	}
	if req.ProjectId != 0 {
		coon = coon.Where("project_id = ?", req.ProjectId)
	}
	if req.ServeId != 0 {
		coon = coon.Where("serve_id = ?", req.ServeId)
	}
	if req.Type != "" {
		coon = coon.Where("type = ?", req.Type)
	}
	if req.SourceType != 0 {
		coon = coon.Where("source_type = ?", req.SourceType)
	}
	if req.ParentId != 0 {
		coon = coon.Where("parent_id = ?", req.ParentId)
	}

	err = coon.Order("ordr DESC").First(&res).Error

	return

}
func (r *CategoryRepo) GetChild(tenantId consts.TenantId, categories, result []*model.Category, parentId int) []*model.Category {
	child := make([]*model.Category, 0)
	for _, item := range categories {
		if item.ParentId == parentId {
			result = append(result, item)
			child = append(child, item)
		}
	}

	if len(child) != 0 {
		for _, v := range child {
			return r.GetChild(tenantId, categories, result, int(v.ID))
		}
	}
	return result
}

func (r *CategoryRepo) GetAllChild(tenantId consts.TenantId, typ serverConsts.CategoryDiscriminator, projectId uint, parentId int) (child []*model.Category, err error) {
	pos, err := r.ListByProject(tenantId, typ, projectId)
	if err != nil || len(pos) == 0 {
		return
	}

	child = r.GetChild(tenantId, pos, child, parentId)
	return
}

func (r *CategoryRepo) GetRootNode(tenantId consts.TenantId, projectId uint, typ serverConsts.CategoryDiscriminator) (node model.Category, err error) {
	err = r.GetDB(tenantId).Model(&model.Category{}).
		Where("project_id = ?", projectId).
		Where("type = ?", typ).
		Where("name = ? AND NOT deleted", "分类").
		First(&node).Error

	return
}

func (r *CategoryRepo) CopySelf(tenantId consts.TenantId, id, newParentId int) (category model.Category, err error) {
	category, err = r.Get(tenantId, id)
	category.ID = 0
	if newParentId != 0 {
		category.ParentId = newParentId
	} else { // 复制的第一个节点重命名
		category.Name = fmt.Sprintf("%s_copy", category.Name)
	}
	_, category.Ordr = r.UpdateOrder(tenantId, serverConsts.After, id, category.Type, category.ProjectId)
	err = r.Save(tenantId, &category)

	return category, err
}
