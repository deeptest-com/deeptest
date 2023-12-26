package repo

import (
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
)

type ComponentSchemaRepo struct {
	DB           *gorm.DB `inject:""`
	*BaseRepo    `inject:""`
	CategoryRepo *CategoryRepo `inject:""`
}

func (r *ComponentSchemaRepo) DeleteByIds(ids []uint) (err error) {
	err = r.DB.Model(&model.ComponentSchema{}).
		Where("id IN (?)", ids).
		Update("deleted", 1).Error

	return
}

func (r *ComponentSchemaRepo) GetCategoryCount(result interface{}, projectId uint) (err error) {
	err = r.DB.Raw("select count(id) count, parent_id category_id from "+model.Category{}.TableName()+" where not deleted and not disabled and project_id=? and type = ? and entity_id != 0 group by category_id", projectId, serverConsts.SchemaCategory).Scan(result).Error
	return
}

func (r *ComponentSchemaRepo) ListAll() (res []model.ComponentSchema, err error) {
	err = r.DB.Where("NOT deleted AND not disabled ").Find(&res).Error
	return
}

func (r *ComponentSchemaRepo) GetSchemasNotExistedInCategory(projectIds []uint) (res []model.ComponentSchema, err error) {
	err = r.DB.Model(&model.ComponentSchema{}).
		Joins("left join biz_category c on biz_project_serve_component_schema.id=c.entity_id").
		Select("biz_project_serve_component_schema.*").
		Where("biz_project_serve_component_schema.project_id IN (?) and c.id is null", projectIds).
		Find(&res).Error

	return
}

func (r *ComponentSchemaRepo) SaveEntity(category *model.Category) (err error) {
	schema := model.ComponentSchema{
		ProjectId: category.ProjectId,
		Name:      category.Name,
		Content:   "{\"type\":\"object\"}",
	}

	joinedPath, err := r.CategoryRepo.GetJoinedPath(category.ID)
	if err != nil {
		return
	}

	schema.Ref = "#/components/schemas" + joinedPath

	err = r.Save(0, &schema)
	if err != nil {
		return
	}

	category.EntityId = schema.ID
	err = r.Save(category.ID, &category)

	return
}

func (r *ComponentSchemaRepo) UpdateRefById(id uint, ref string) (err error) {
	err = r.DB.Model(&model.Category{}).
		Where("id = ?", id).
		Update("ref", ref).Error

	return
}

func (r *ComponentSchemaRepo) ChangeRef(id, categoryId uint) (err error) {
	path, err := r.CategoryRepo.GetJoinedPath(categoryId)
	if err != nil {
		return
	}

	ref := "#/components/schemas" + path
	err = r.UpdateRefById(id, ref)

	return
}
