package repo

import (
	"database/sql"
	"fmt"
	v1 "github.com/deeptest-com/deeptest/cmd/server/v1/domain"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	serverConsts "github.com/deeptest-com/deeptest/internal/server/consts"
	"github.com/deeptest-com/deeptest/internal/server/modules/model"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

type CategoryRepo struct {
	*BaseRepo   `inject:""`
	DB          *gorm.DB     `inject:""`
	ProjectRepo *ProjectRepo `inject:""`
}

func (r *CategoryRepo) GetTree(tenantId consts.TenantId, typ serverConsts.CategoryDiscriminator, projectId uint, nodeType serverConsts.NodeCreateType) (root *v1.Category, err error) {
	pos, err := r.ListByProject(tenantId, typ, projectId, nodeType)
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

func (r *CategoryRepo) ListByProject(tenantId consts.TenantId, typ serverConsts.CategoryDiscriminator, projectId uint, nodeType serverConsts.NodeCreateType) (pos []*model.Category, err error) {
	db := r.GetDB(tenantId).
		Where("project_id=?", projectId).
		Where("type=?", typ).
		Where("NOT deleted")

	if nodeType != "" {
		if nodeType == serverConsts.Dir {
			db = db.Where("entity_id = 0 or entity_id IS NULL")
		} else {
			db = db.Where("entity_id != 0")
		}
	}

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
			EntityId: po.EntityId,
		}

		tos = append(tos, &to)
	}

	return
}

func (r *CategoryRepo) makeTree(tenantId consts.TenantId, findIn []*v1.Category, parent *v1.Category) { //参数为父节点，添加父节点的子节点指针切片
	children, _ := r.hasChild(findIn, parent) // 判断节点是否有子节点并返回

	if children != nil {
		parent.Children = append(parent.Children, children[0:]...) // 添加子节点

		for _, child := range children { // 查询子节点的子节点，并添加到子节点
			_, has := r.hasChild(findIn, child)
			if has {
				r.makeTree(tenantId, findIn, child) // 递归添加节点
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
		r.GetDB(tenantId).Model(&model.Category{}).
			Where("NOT deleted AND parent_id=? AND type = ? AND project_id = ? ",
				parentId, typ, projectId).
			Update("ordr", gorm.Expr("ordr + 1"))
		ordr = 1
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

	//if req.Type == serverConsts.SchemaCategory && int(req.Parent) != po.ParentId {
	//	po.ParentId = int(req.Parent)
	//	po.Ordr = r.GetMaxOrder(req.Parent, req.Type, po.ProjectId)
	//}

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
func (r *CategoryRepo) GetChildren(tenantId consts.TenantId, nodeId uint, nodeType serverConsts.NodeCreateType) (children []*model.Category, err error) {
	db := r.GetDB(tenantId).Where("parent_id=? and not deleted", nodeId)
	if nodeType != "" {
		if nodeType == serverConsts.Dir {
			db = db.Where("entity_id = 0 or entity_id IS NULL")
		} else {
			db = db.Where("entity_id != 0")
		}
	}

	err = db.Order("ordr ASC").Find(&children).Error
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

func (r *CategoryRepo) GetAllChild(tenantId consts.TenantId, typ serverConsts.CategoryDiscriminator, projectId uint, parentId int, nodeType serverConsts.NodeCreateType) (child []*model.Category, err error) {
	pos, err := r.ListByProject(tenantId, typ, projectId, nodeType)
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
		Where("parent_id = ? AND NOT deleted", 0).
		First(&node).Error

	return
}

func (r *CategoryRepo) CopySelf(tenantId consts.TenantId, id, newParentId int) (category model.Category, err error) {
	category, err = r.Get(tenantId, id)

	category.ID = 0
	if newParentId != 0 {
		category.ParentId = newParentId
	} else { // 复制的第一个节点重命名
		category.Name = fmt.Sprintf("CopyOf%s", category.Name)
	}
	_, category.Ordr = r.UpdateOrder(tenantId, serverConsts.After, id, category.Type, category.ProjectId)
	err = r.Save(tenantId, &category)

	return category, err
}

func (r *CategoryRepo) GetEntityIdsByIds(tenantId consts.TenantId, ids []uint) (entityIds []uint, err error) {
	err = r.GetDB(tenantId).Model(&model.Category{}).
		Select("entity_id").
		Where("id IN (?)", ids).
		Find(&entityIds).Error

	return
}

func (r *CategoryRepo) GetByEntityId(tenantId consts.TenantId, entityId uint, _type serverConsts.CategoryDiscriminator) (category model.Category, err error) {
	err = r.GetDB(tenantId).Model(&model.Category{}).
		Where("entity_id = ? AND type = ? AND NOT deleted", entityId, _type).
		First(&category).Error

	return
}

func (r *CategoryRepo) DeleteByEntityId(tenantId consts.TenantId, entityId uint, _type serverConsts.CategoryDiscriminator) (err error) {
	err = r.GetDB(tenantId).Model(&model.Category{}).
		Where("entity_id = ? AND type = ?", entityId, _type).
		Update("deleted", true).Error

	return
}

func (r *CategoryRepo) UpdateEntityId(tenantId consts.TenantId, id, entityId uint) (err error) {
	err = r.GetDB(tenantId).Model(&model.Category{}).
		Where("id = ?", id).
		Update("entity_id", entityId).Error

	return
}

func (r *CategoryRepo) UpdateNameByEntityId(tenantId consts.TenantId, entityId uint, name string, _type serverConsts.CategoryDiscriminator) (err error) {
	err = r.GetDB(tenantId).Model(&model.Category{}).
		Where("entity_id = ? AND type = ?", entityId, _type).
		Update("name", name).Error

	return
}

func (r *CategoryRepo) BatchAddProjectRootSchemaCategory(tenantId consts.TenantId, projectIds []uint) (err error) {
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

	err = r.GetDB(tenantId).Create(&roots).Error

	return
}

func (r *CategoryRepo) BatchGetRootNodeProjectIds(tenantId consts.TenantId, projectIds []uint, typ serverConsts.CategoryDiscriminator) (res []uint, err error) {
	err = r.GetDB(tenantId).Model(&model.Category{}).
		Select("project_id").
		Where("project_id IN (?)", projectIds).
		Where("type = ?", typ).
		Where("parent_id = ?", 0).
		Where("name = ? AND NOT deleted", "分类").
		Find(&res).Error

	return
}

func (r *CategoryRepo) BatchGetRootNodes(tenantId consts.TenantId, projectIds []uint, typ serverConsts.CategoryDiscriminator) (res []model.Category, err error) {
	err = r.GetDB(tenantId).Model(&model.Category{}).
		Where("project_id IN (?)", projectIds).
		Where("type = ?", typ).
		Where("parent_id = ?", 0).
		Where("name = ? AND NOT deleted", "分类").
		Find(&res).Error

	return
}

func (r *CategoryRepo) GetJoinedPath(tenantId consts.TenantId, categoryId uint) (path []string, err error) {

	sql := `
		WITH RECURSIVE temp(id,parent_id,name) AS
		(
			SELECT id,parent_id,name from biz_category where id = %d
		  UNION ALL
		  SELECT b.id,b.parent_id,b.name from biz_category b, temp c where c.parent_id = b.id and  b.parent_id > 0
		) 
		select name from temp,(select @row_number := 0) x ORDER BY (@row_number:=@row_number+1) desc
`
	sql = fmt.Sprintf(sql, categoryId)
	err = r.GetDB(tenantId).Raw(sql).Scan(&path).Error

	return
}

func (r *CategoryRepo) GetRoot(tenantId consts.TenantId, typ serverConsts.CategoryDiscriminator, projectId uint) (res model.Category, err error) {

	err = r.GetDB(tenantId).Model(&model.Category{}).
		Where("parent_id = 0 AND type = ? AND project_id = ? and not deleted", typ, projectId).
		First(&res).Error

	return

}

func (r *CategoryRepo) GetChildrenNodes(tenantId consts.TenantId, categoryId uint, nodeType serverConsts.NodeCreateType) (ret []*model.Category, err error) {
	ret, err = r.GetChildren(tenantId, categoryId, nodeType)
	return
}

func (r *CategoryRepo) GetEntityCountByCategoryId(tenantId consts.TenantId, categoryId uint) int64 {
	var ret []sql.NullInt64
	sql := `WITH RECURSIVE temp(id,count) AS
		(
			SELECT id,if(entity_id=0,0,1 ) count from biz_category where parent_id = %d and not deleted
		  UNION ALL
		  SELECT b.id,if(b.entity_id=0,0,1 ) count from biz_category b, temp c where c.id = b.parent_id and not deleted
		) 
		select sum(count) count from temp
`
	sql = fmt.Sprintf(sql, categoryId)
	r.GetDB(tenantId).Raw(sql).Scan(&ret)

	if len(ret) > 0 {
		return ret[0].Int64
	}

	return 0
}

func (r *CategoryRepo) SaveEntityNode(tenantId consts.TenantId, nodeId uint, typ serverConsts.CategoryDiscriminator, projectId, categoryId, entityId uint, name string) (id uint, err error) {
	var entity model.Category
	if nodeId == 0 {
		entity, _ = r.GetByEntityId(tenantId, entityId, typ)
	} else {
		entity, _ = r.Get(tenantId, int(nodeId))
	}

	entity.ProjectId, entity.Name, entity.ParentId, entity.Type, entity.EntityId = projectId, name, int(categoryId), typ, entityId
	entity.Ordr = r.GetMaxOrder(tenantId, categoryId, typ, projectId)
	entity.IsDir = false

	err = r.GetDB(tenantId).Save(&entity).Error
	return entity.ID, err

}

func (r *CategoryRepo) BatchGetByIds(tenantId consts.TenantId, ids []int) (res []model.Category, err error) {
	err = r.GetDB(tenantId).Model(&model.Category{}).
		Where("id IN (?) AND NOT deleted", ids).
		Find(&res).Error

	return
}

func (r *CategoryRepo) GetAllChildrenMap(tenantId consts.TenantId, categoryId uint) (root *model.Category, ret map[uint][]*model.Category, err error) {
	var res []*model.Category
	ret = map[uint][]*model.Category{}
	sql := `WITH RECURSIVE temp(id,name,parent_id,entity_id,ordr) AS
		(
			SELECT id,name,parent_id,entity_id,ordr from biz_category where id = %d and not deleted
		  UNION ALL
		  SELECT b.id,b.name,b.parent_id,b.entity_id,b.ordr from biz_category b, temp c where c.id = b.parent_id and not deleted
		) 
		select * from temp`
	sql = fmt.Sprintf(sql, categoryId)
	err = r.GetDB(tenantId).Raw(sql).Scan(&res).Error

	if len(res) > 0 {
		root = res[0]
	}

	for _, item := range res {
		ret[uint(item.ParentId)] = append(ret[uint(item.ParentId)], item)
	}

	return
}

func (r *CategoryRepo) UpdateParentIdByEntityIds(tenantId consts.TenantId, entityIds []uint, parentId uint, _type serverConsts.CategoryDiscriminator) (err error) {
	for _, entityId := range entityIds {
		entity, err := r.GetByEntityId(tenantId, entityId, _type)
		if err != nil {
			continue
		}
		entity.ParentId = int(parentId)
		entity.Ordr = r.GetMaxOrder(tenantId, parentId, _type, entity.ProjectId)
		err = r.GetDB(tenantId).Save(&entity).Error
	}
	return err
}
