package repo

import (
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
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

func (r *CategoryRepo) GetTree(typ serverConsts.CategoryDiscriminator, projectId uint) (root *v1.Category, err error) {
	pos, err := r.ListByProject(typ, projectId)
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

func (r *CategoryRepo) ListByProject(typ serverConsts.CategoryDiscriminator, projectId uint) (pos []*model.Category, err error) {
	db := r.DB.
		Where("project_id=?", projectId).
		Where("type=?", typ).
		Where("NOT deleted")

	err = db.
		Order("parent_id ASC, ordr ASC").
		Find(&pos).Error

	return
}

func (r *CategoryRepo) Get(id int) (po model.Category, err error) {
	err = r.DB.Where("id = ?", id).First(&po).Error
	return
}

func (r *CategoryRepo) toTos(pos []*model.Category) (tos []*v1.Category) {
	for _, po := range pos {
		to := v1.Category{
			Id:       int64(po.ID),
			Name:     po.Name,
			Desc:     po.Desc,
			ParentId: int64(po.ParentId),
			EntityId: po.EntityId,
		}

		tos = append(tos, &to)
	}

	return
}

func (r *CategoryRepo) makeTree(findIn []*v1.Category, parent *v1.Category) { //参数为父节点，添加父节点的子节点指针切片
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

func (r *CategoryRepo) hasChild(categories []*v1.Category, parent *v1.Category) (
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

func (r *CategoryRepo) Save(category *model.Category) (err error) {
	if category != nil && category.ID == 0 && category.ParentId != 0 && category.Ordr == 0 {
		category.Ordr = r.GetMaxOrder(uint(category.ParentId), category.Type, category.ProjectId)
	}
	err = r.DB.Save(category).Error

	return
}

func (r *CategoryRepo) UpdateOrder(pos serverConsts.DropPos, targetId int, typ serverConsts.CategoryDiscriminator, projectId uint) (
	parentId int, ordr int) {
	if pos == serverConsts.Inner {
		parentId = targetId

		//var preChild model.Category
		//r.DB.Where("parent_id=? AND type = ? AND project_id = ?", parentId, typ, projectId).
		//	Order("ordr DESC").Limit(1).
		//	First(&preChild)
		//
		//ordr = preChild.Ordr + 1

		r.DB.Model(&model.Category{}).
			Where("NOT deleted AND parent_id=? AND type = ? AND project_id = ? ",
				parentId, typ, projectId).
			Update("ordr", gorm.Expr("ordr + 1"))
		ordr = 1
	} else if pos == serverConsts.Before {
		brother, _ := r.Get(targetId)
		ordr = brother.Ordr
		parentId = brother.ParentId

		r.DB.Model(&model.Category{}).
			Where("NOT deleted AND parent_id=? AND type = ? AND project_id = ? AND ordr >= ?",
				parentId, typ, projectId, brother.Ordr).
			Update("ordr", gorm.Expr("ordr + 1"))

	} else if pos == serverConsts.After {
		brother, _ := r.Get(targetId)
		parentId = brother.ParentId

		r.DB.Model(&model.Category{}).
			Where("NOT deleted AND parent_id=? AND type = ? AND project_id = ? AND ordr > ?",
				parentId, typ, projectId, brother.Ordr).
			Update("ordr", gorm.Expr("ordr + 1"))

		ordr = brother.Ordr + 1

	}

	return
}

func (r *CategoryRepo) UpdateName(id int, name string) (err error) {
	err = r.DB.Model(&model.Category{}).
		Where("id = ?", id).
		Update("name", name).Error

	return
}

func (r *CategoryRepo) Update(req v1.CategoryReq) (err error) {
	po := new(model.Category)
	po.ID = uint(req.Id)

	err = r.DB.First(&po).Error
	if err != nil {
		return err
	}

	po.Name = req.Name
	po.Desc = req.Desc

	//if req.Type == serverConsts.SchemaCategory && int(req.Parent) != po.ParentId {
	//	po.ParentId = int(req.Parent)
	//	po.Ordr = r.GetMaxOrder(req.Parent, req.Type, po.ProjectId)
	//}

	err = r.DB.Save(&po).Error

	return
}

func (r *CategoryRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&model.Category{}).
		Where("id=?", id).
		Update("deleted", true).
		Error

	return
}

func (r *CategoryRepo) BatchDelete(ids []uint) (err error) {
	err = r.DB.Model(&model.Category{}).
		Where("id IN (?)", ids).
		Update("deleted", true).
		Error

	return
}
func (r *CategoryRepo) GetChildren(nodeId uint) (children []*model.Category, err error) {
	err = r.DB.Where("parent_id=?", nodeId).Find(&children).Error
	return
}

func (r *CategoryRepo) UpdateOrdAndParent(node model.Category) (err error) {
	err = r.DB.Model(&node).
		Updates(model.Category{Ordr: node.Ordr, ParentId: node.ParentId}).
		Error

	return
}

func (r *CategoryRepo) GetMaxOrder(parentId uint, typ serverConsts.CategoryDiscriminator, projectId uint) (order int) {
	node := model.Category{}

	err := r.DB.Model(&model.Category{}).
		Where("parent_id=? AND type = ? AND project_id = ?", parentId, typ, projectId).
		Order("ordr DESC").
		First(&node).Error

	if err == nil {
		order = node.Ordr + 1
	}

	return
}

func (r *CategoryRepo) GetByItem(parentId uint, typ serverConsts.CategoryDiscriminator, projectId uint, name string) (res model.Category, err error) {

	err = r.DB.Model(&model.Category{}).
		Where("parent_id=? AND type = ? AND project_id = ? AND name = ? and not deleted", parentId, typ, projectId, name).
		Order("ordr DESC").
		First(&res).Error

	return

}

func (r *CategoryRepo) GetDetail(req model.Category) (res model.Category, err error) {
	coon := r.DB.Model(&model.Category{}).Where("not deleted")
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
func (r *CategoryRepo) GetChild(categories, result []*model.Category, parentId int) []*model.Category {
	child := make([]*model.Category, 0)
	for _, item := range categories {
		if item.ParentId == parentId {
			result = append(result, item)
			child = append(child, item)
		}
	}

	if len(child) != 0 {
		for _, v := range child {
			return r.GetChild(categories, result, int(v.ID))
		}
	}
	return result
}

func (r *CategoryRepo) GetAllChild(typ serverConsts.CategoryDiscriminator, projectId uint, parentId int) (child []*model.Category, err error) {
	pos, err := r.ListByProject(typ, projectId)
	if err != nil || len(pos) == 0 {
		return
	}

	child = r.GetChild(pos, child, parentId)
	return
}

func (r *CategoryRepo) GetRootNode(projectId uint, typ serverConsts.CategoryDiscriminator) (node model.Category, err error) {
	err = r.DB.Model(&model.Category{}).
		Where("project_id = ?", projectId).
		Where("type = ?", typ).
		Where("name = ? AND NOT deleted", "分类").
		First(&node).Error

	return
}

func (r *CategoryRepo) GetEntityIdsByIds(ids []uint) (entityIds []uint, err error) {
	err = r.DB.Model(&model.Category{}).
		Select("entity_id").
		Where("id IN (?)", ids).
		Find(&entityIds).Error

	return
}

func (r *CategoryRepo) GetByEntityId(entityId uint, _type serverConsts.CategoryDiscriminator) (category model.Category, err error) {
	err = r.DB.Model(&model.Category{}).
		Where("entity_id = ? AND type = ? AND NOT deleted", entityId, _type).
		First(&category).Error

	return
}

func (r *CategoryRepo) DeleteByEntityId(entityId uint) (err error) {
	err = r.DB.Model(&model.Category{}).
		Where("entity_id = ?", entityId).
		Update("deleted", true).Error

	return
}

func (r *CategoryRepo) UpdateEntityId(id, entityId uint) (err error) {
	err = r.DB.Model(&model.Category{}).
		Where("id = ?", id).
		Update("entity_id", entityId).Error

	return
}

func (r *CategoryRepo) UpdateNameByEntityId(entityId uint, name string, _type serverConsts.CategoryDiscriminator) (err error) {
	err = r.DB.Model(&model.Category{}).
		Where("entity_id = ? AND type = ?", entityId, _type).
		Update("name", name).Error

	return
}

func (r *CategoryRepo) BatchAddProjectRootSchemaCategory(projectIds []uint) (err error) {
	roots := make([]model.Category, 0)
	for _, projectId := range projectIds {
		root := model.Category{
			Name:      "分类",
			Type:      serverConsts.SchemaCategory,
			ProjectId: projectId,
			IsDir:     true,
		}

		roots = append(roots, root)
	}

	err = r.DB.Create(&roots).Error

	return
}

func (r *CategoryRepo) BatchGetRootNodeProjectIds(projectIds []uint, typ serverConsts.CategoryDiscriminator) (res []uint, err error) {
	err = r.DB.Model(&model.Category{}).
		Select("project_id").
		Where("project_id IN (?)", projectIds).
		Where("type = ?", typ).
		Where("parent_id = ?", 0).
		Where("name = ? AND NOT deleted", "分类").
		Find(&res).Error

	return
}

func (r *CategoryRepo) BatchGetRootNodes(projectIds []uint, typ serverConsts.CategoryDiscriminator) (res []model.Category, err error) {
	err = r.DB.Model(&model.Category{}).
		Where("project_id IN (?)", projectIds).
		Where("type = ?", typ).
		Where("parent_id = ?", 0).
		Where("name = ? AND NOT deleted", "分类").
		Find(&res).Error

	return
}

func (r *CategoryRepo) GetJoinedPath(categoryId uint) (path []string, err error) {

	sql := `
		WITH RECURSIVE temp(id,parent_id,name) AS
		(
			SELECT id,parent_id,name from biz_category where id = %d
		  UNION ALL
		  SELECT b.id,b.parent_id,b.name from biz_category b, temp c where c.parent_id = b.id and  b.parent_id > 0
		) 
		select name from temp ORDER BY id asc
`
	sql = fmt.Sprintf(sql, categoryId)
	err = r.DB.Raw(sql).Scan(&path).Error

	return
}

func (r *CategoryRepo) GetRoot(typ serverConsts.CategoryDiscriminator, projectId uint) (res model.Category, err error) {

	err = r.DB.Model(&model.Category{}).
		Where("parent_id = 0 AND type = ? AND project_id = ? and not deleted", typ, projectId).
		First(&res).Error

	return

}

func (r *CategoryRepo) CopySelf(id, newParentId int) (category model.Category, err error) {
	category, err = r.Get(id)
	category.ID = 0
	if newParentId != 0 {
		category.ParentId = newParentId
	} else { // 复制的第一个节点重命名
		category.Name = fmt.Sprintf("CopyOf%s", category.Name)
	}
	_, category.Ordr = r.UpdateOrder(serverConsts.After, id, category.Type, category.ProjectId)
	err = r.Save(&category)

	return category, err
}
