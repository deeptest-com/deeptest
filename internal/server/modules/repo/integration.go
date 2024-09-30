package repo

import (
	v1 "github.com/deeptest-com/deeptest/cmd/server/v1/domain"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
)

type IntegrationRepo struct {
	*BaseRepo `inject:""`
	DB        *gorm.DB `inject:""`
}

func (r *IntegrationRepo) GetProjectListWithRoleBySpace(tenantId consts.TenantId, spaceCode string) (res []v1.ProjectListWithRole, err error) {
	err = r.GetDB(tenantId).Model(&model.Project{}).
		Joins("LEFT JOIN biz_integration_project_space_rel rel ON biz_project.id=rel.project_id").
		Select("biz_project.id, biz_project.name, biz_project.short_name").
		Where("rel.space_code = ? AND not biz_project.deleted AND not biz_project.disabled", spaceCode).
		Find(&res).Error

	return
}

func (r *IntegrationRepo) DeleteBySpaceCode(tenantId consts.TenantId, spaceCode string) (err error) {
	err = r.GetDB(tenantId).Model(&model.ProjectSpaceRel{}).
		Where("space_code = ?", spaceCode).
		Delete(&model.ProjectSpaceRel{}).Error

	return
}

func (r *IntegrationRepo) DeleteSpaceByProject(tenantId consts.TenantId, projectId uint) (err error) {
	err = r.GetDB(tenantId).Model(&model.ProjectSpaceRel{}).
		Where("project_id = ?", projectId).
		Delete(&model.ProjectSpaceRel{}).Error

	return
}

func (r *IntegrationRepo) BatchCreateProjectSpaceRel(tenantId consts.TenantId, relations []model.ProjectSpaceRel) (err error) {
	err = r.GetDB(tenantId).Model(&model.ProjectSpaceRel{}).Create(&relations).Error

	return
}

func (r *IntegrationRepo) DeleteProductByProject(tenantId consts.TenantId, projectId uint) (err error) {
	err = r.GetDB(tenantId).Model(&model.ProjectProductRel{}).
		Where("project_id = ?", projectId).
		Delete(&model.ProjectProductRel{}).Error

	return
}

func (r *IntegrationRepo) BatchCreateProjectProductRel(tenantId consts.TenantId, relations []model.ProjectProductRel) (err error) {
	err = r.GetDB(tenantId).Model(&model.ProjectProductRel{}).Create(&relations).Error

	return
}

func (r *IntegrationRepo) GetAllProductIds(tenantId consts.TenantId) (res []uint, err error) {
	err = r.GetDB(tenantId).Model(&model.ProjectProductRel{}).
		Select("distinct product_id").
		Where("NOT deleted AND NOT disabled").
		Find(&res).Error

	return
}

func (r *IntegrationRepo) GetAllProjectProductRels(tenantId consts.TenantId) (res []model.ProjectProductRel, err error) {
	err = r.GetDB(tenantId).Model(&model.ProjectProductRel{}).
		Where("NOT deleted AND NOT disabled").
		Find(&res).Error

	return
}

func (r *IntegrationRepo) GetProductsByProject(tenantId consts.TenantId, projectId uint) (res []uint, err error) {
	err = r.GetDB(tenantId).Model(&model.ProjectProductRel{}).
		Select("product_id").
		Where("project_id = ? AND NOT deleted AND NOT disabled", projectId).
		Find(&res).Error

	return
}

func (r *IntegrationRepo) GetSpacesByProject(tenantId consts.TenantId, projectId uint) (res []string, err error) {
	err = r.GetDB(tenantId).Model(&model.ProjectSpaceRel{}).
		Select("space_code").
		Where("project_id = ? AND NOT deleted AND NOT disabled", projectId).
		Find(&res).Error

	return
}

func (r *IntegrationRepo) DeleteEngineeringByProject(tenantId consts.TenantId, projectId uint) (err error) {
	err = r.GetDB(tenantId).Model(&model.ProjectEngineeringRel{}).
		Where("project_id = ?", projectId).
		Delete(&model.ProjectEngineeringRel{}).Error

	return
}

func (r *IntegrationRepo) BatchCreateProjectEngineeringRel(tenantId consts.TenantId, relations []model.ProjectEngineeringRel) (err error) {
	err = r.GetDB(tenantId).Model(&model.ProjectEngineeringRel{}).Create(&relations).Error

	return
}

func (r *IntegrationRepo) GetEngineeringByProject(tenantId consts.TenantId, projectId uint) (res []string, err error) {
	err = r.GetDB(tenantId).Model(&model.ProjectEngineeringRel{}).
		Select("code").
		Where("project_id = ? AND NOT deleted AND NOT disabled", projectId).
		Find(&res).Error

	return
}

func (r *IntegrationRepo) GetProjectByEngineering(tenantId consts.TenantId, engineering string) (res []uint, err error) {
	err = r.GetDB(tenantId).Model(&model.ProjectEngineeringRel{}).
		Select("project_id").
		Where("code = ? AND NOT deleted AND NOT disabled", engineering).
		Find(&res).Error

	return
}
