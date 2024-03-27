package repo

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/getkin/kin-openapi/openapi3"
	"gorm.io/gorm"
	"strings"
)

type ComponentSchemaRepo struct {
	DB           *gorm.DB `inject:""`
	*BaseRepo    `inject:""`
	CategoryRepo *CategoryRepo `inject:""`
}

func (r *ComponentSchemaRepo) DeleteByIds(tenantId consts.TenantId, ids []uint) (err error) {
	err = r.GetDB(tenantId).Model(&model.ComponentSchema{}).
		Where("id IN (?)", ids).
		Update("deleted", 1).Error

	return
}

func (r *ComponentSchemaRepo) GetCategoryCount(tenantId consts.TenantId, result interface{}, projectId uint) (err error) {
	err = r.GetDB(tenantId).Raw("select count(id) count, parent_id category_id from "+model.Category{}.TableName()+" where not deleted and not disabled and project_id=? and type = ? and entity_id != 0 group by category_id", projectId, serverConsts.SchemaCategory).Scan(result).Error
	return
}

func (r *ComponentSchemaRepo) ListAll(tenantId consts.TenantId) (res []model.ComponentSchema, err error) {
	err = r.GetDB(tenantId).Where("NOT deleted AND not disabled ").Find(&res).Error
	return
}

func (r *ComponentSchemaRepo) GetSchemasNotExistedInCategory(tenantId consts.TenantId, projectIds []uint) (res []model.ComponentSchema, err error) {
	err = r.GetDB(tenantId).Model(&model.ComponentSchema{}).
		Joins("left join biz_category c on biz_project_serve_component_schema.id=c.entity_id").
		Select("biz_project_serve_component_schema.*").
		Where("biz_project_serve_component_schema.project_id IN (?) and c.id is null", projectIds).
		Find(&res).Error

	return
}

func (r *ComponentSchemaRepo) SaveEntity(tenantId consts.TenantId, category *model.Category) (err error) {
	schema := model.ComponentSchema{
		ProjectId: category.ProjectId,
		Name:      category.Name,
		Content:   "{\"type\":\"object\"}",
		Type:      openapi3.TypeObject,
	}

	path, err := r.CategoryRepo.GetJoinedPath(tenantId, category.ID)
	if err != nil {
		return
	}

	schema.Ref = "#/components/schemas/" + strings.Join(path, ".")

	err = r.Save(tenantId, 0, &schema)
	if err != nil {
		return
	}

	category.EntityId = schema.ID
	err = r.Save(tenantId, category.ID, &category)

	return
}

func (r *ComponentSchemaRepo) UpdateRefById(tenantId consts.TenantId, id uint, ref string) (err error) {
	err = r.GetDB(tenantId).Model(&model.ComponentSchema{}).
		Where("id = ?", id).
		Update("ref", ref).Error

	return
}

func (r *ComponentSchemaRepo) ChangeRef(tenantId consts.TenantId, id, categoryId uint) (err error) {
	path, err := r.CategoryRepo.GetJoinedPath(tenantId, categoryId)
	if err != nil {
		return
	}

	ref := "#/components/schemas/" + strings.Join(path, ".")
	err = r.UpdateRefById(tenantId, id, ref)

	return
}

func (r *ComponentSchemaRepo) MoveEntity(tenantId consts.TenantId, category *model.Category) (err error) {
	return r.ChangeRef(tenantId, category.EntityId, category.ID)
}
