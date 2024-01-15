package repo

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
)

type IntegrationRepo struct {
	*BaseRepo `inject:""`
	DB        *gorm.DB `inject:""`
}

func (r *IntegrationRepo) GetProjectListWithRoleBySpace(spaceCode string) (res []v1.ProjectListWithRole, err error) {
	err = r.DB.Model(&model.Project{}).
		Joins("LEFT JOIN biz_integration_project_space_rel rel ON biz_project.id=rel.project_id").
		Joins("LEFT JOIN biz_project_member m ON biz_project.id=m.project_id").
		Joins("LEFT JOIN biz_project_role r ON m.project_role_id=r.id").
		Select("biz_project.id, biz_project.name, biz_project.short_name, r.name as role_name").
		Where("rel.space_code = ? AND not biz_project.deleted AND not biz_project.disabled", spaceCode).
		Find(&res).Error

	return
}

func (r *IntegrationRepo) DeleteBySpaceCode(spaceCode string) (err error) {
	err = r.DB.Model(&model.ProjectSpaceRel{}).
		Where("space_code = ?", spaceCode).
		Delete(&model.ProjectSpaceRel{}).Error

	return
}

func (r *IntegrationRepo) DeleteSpaceByProject(projectId uint) (err error) {
	err = r.DB.Model(&model.ProjectSpaceRel{}).
		Where("project_id = ?", projectId).
		Delete(&model.ProjectSpaceRel{}).Error

	return
}

func (r *IntegrationRepo) BatchCreateProjectSpaceRel(relations []model.ProjectSpaceRel) (err error) {
	err = r.DB.Model(&model.ProjectSpaceRel{}).Create(&relations).Error

	return
}

func (r *IntegrationRepo) DeleteProductByProject(projectId uint) (err error) {
	err = r.DB.Model(&model.ProjectProductRel{}).
		Where("project_id = ?", projectId).
		Delete(&model.ProjectProductRel{}).Error

	return
}

func (r *IntegrationRepo) BatchCreateProjectProductRel(relations []model.ProjectProductRel) (err error) {
	err = r.DB.Model(&model.ProjectProductRel{}).Create(&relations).Error

	return
}

func (r *IntegrationRepo) GetAllProductIds() (res []uint, err error) {
	err = r.DB.Model(&model.ProjectProductRel{}).
		Select("distinct product_id").
		Where("NOT deleted AND NOT disabled").
		Find(&res).Error

	return
}

func (r *IntegrationRepo) GetAllProjectProductRels() (res []model.ProjectProductRel, err error) {
	err = r.DB.Model(&model.ProjectProductRel{}).
		Where("NOT deleted AND NOT disabled").
		Find(&res).Error

	return
}

func (r *IntegrationRepo) GetProductsByProject(projectId uint) (res []uint, err error) {
	err = r.DB.Model(&model.ProjectProductRel{}).
		Select("product_id").
		Where("project_id = ? AND NOT deleted AND NOT disabled", projectId).
		Find(&res).Error

	return
}

func (r *IntegrationRepo) GetSpacesByProject(projectId uint) (res []string, err error) {
	err = r.DB.Model(&model.ProjectSpaceRel{}).
		Select("space_code").
		Where("project_id = ? AND NOT deleted AND NOT disabled", projectId).
		Find(&res).Error

	return
}
