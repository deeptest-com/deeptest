package repo

import (
	"fmt"
	v1 "github.com/deeptest-com/deeptest/cmd/server/v1/domain"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/server/core/dao"
	"github.com/deeptest-com/deeptest/internal/server/modules/model"
	_domain "github.com/deeptest-com/deeptest/pkg/domain"
	logUtils "github.com/deeptest-com/deeptest/pkg/lib/log"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ProjectRolePermRepo struct {
	*BaseRepo       `inject:""`
	DB              *gorm.DB         `inject:""`
	ProjectRepo     *ProjectRepo     `inject:""`
	ProjectRoleRepo *ProjectRoleRepo `inject:""`
}

func (r *ProjectRolePermRepo) PaginateRolePerms(tenantId consts.TenantId, req v1.ProjectRolePermPaginateReq) (data _domain.PageData, err error) {
	var count int64
	projectPerms := make([]*model.ProjectPerm, 0)

	db := r.GetDB(tenantId).Model(&model.ProjectPerm{}).Joins("JOIN biz_project_role_perm p ON biz_project_perm.id=p.project_perm_id AND p.project_role_id = ?", req.RoleId)

	err = db.Count(&count).Error
	if err != nil {
		logUtils.Errorf("count project role perms error", zap.String("error:", err.Error()))
		return
	}

	err = db.
		Scopes(dao.PaginateScope(req.Page, req.PageSize, req.Order, req.Field)).
		Find(&projectPerms).Error
	if err != nil {
		logUtils.Errorf("query project role perms error", zap.String("error:", err.Error()))
		return
	}

	data.Populate(projectPerms, count, req.Page, req.PageSize)
	return
}

func (r *ProjectRolePermRepo) UserPermList(tenantId consts.TenantId, req v1.ProjectUserPermsPaginate, userId uint) (data _domain.PageData, err error) {
	currProject, err := r.ProjectRepo.GetCurrProjectByUser(tenantId, userId)
	if err != nil {
		logUtils.Errorf("query project profile error", zap.String("error:", err.Error()))
		return
	}

	var roleId uint
	r.GetDB(tenantId).Model(&model.ProjectMember{}).
		Select("project_role_id").
		Where("user_id = ?", userId).
		Where("project_id = ? AND NOT deleted", currProject.ID).First(&roleId)

	projectPerms := make([]*model.ProjectPerm, 0)

	db := r.GetDB(tenantId).Model(&model.ProjectPerm{}).Joins("JOIN biz_project_role_perm p ON biz_project_perm.id=p.project_perm_id AND p.project_role_id = ?", roleId)

	var count int64
	err = db.Count(&count).Error
	if err != nil {
		logUtils.Errorf("count project role perms error", zap.String("error:", err.Error()))
		return
	}

	err = db.
		Scopes(dao.PaginateScope(req.Page, req.PageSize, req.Order, req.Field)).
		Find(&projectPerms).Error
	if err != nil {
		logUtils.Errorf("query project role perms error", zap.String("error:", err.Error()))
		return
	}

	data.Populate(projectPerms, count, req.Page, req.PageSize)

	return
}

func (r *ProjectRolePermRepo) GetByRoleAndPerm(tenantId consts.TenantId, roleId, permId uint) (ret model.ProjectRolePerm, err error) {
	err = r.GetDB(tenantId).Model(&model.ProjectRolePerm{}).
		Where("project_role_id = ?", roleId).
		Where("project_perm_id = ?", permId).
		First(&ret).Error
	return
}

// GetProjectPermsForRole TODO: 每个角色需要的权限还未确定，需要改动
func (r *ProjectRolePermRepo) GetProjectPermsForRole(tenantId consts.TenantId) (res map[consts.RoleType][]uint, err error) {
	var permIds, testPermIds []uint
	err = r.GetDB(tenantId).Model(&model.ProjectPerm{}).Select("id").Find(&permIds).Error
	err = r.GetDB(tenantId).Model(&model.ProjectPerm{}).Select("id").Where("name like ?", "/api/v1/projects%").Find(&testPermIds).Error
	res = map[consts.RoleType][]uint{
		consts.Admin:          permIds,
		consts.User:           permIds,
		consts.Tester:         permIds,
		consts.Developer:      permIds,
		consts.ProductManager: permIds,
	}

	return
}

func (r *ProjectRolePermRepo) AddPermForProjectRole(tenantId consts.TenantId, roleName consts.RoleType, perms []uint) (successCount int, failItems []string) {
	projectRole, err := r.ProjectRoleRepo.FindByName(tenantId, roleName)
	if err != nil {
		failItems = append(failItems, fmt.Sprintf("为角色%+v添加权限%+v失败，错误%s", roleName, perms, err.Error()))
		return
	}

	projectRoleId := projectRole.ID
	err = r.GetDB(tenantId).Delete(&model.ProjectRolePerm{}, "project_role_id = ?", projectRoleId).Error
	if err != nil {
		failItems = append(failItems, fmt.Sprintf("为角色%+v添加权限%+v失败，错误%s", roleName, perms, err.Error()))
		return
	}

	for _, perm := range perms {
		permModel := &model.ProjectRolePerm{ProjectRolePermBase: v1.ProjectRolePermBase{ProjectRoleId: projectRoleId, ProjectPermId: perm}}
		err := r.GetDB(tenantId).Model(&model.ProjectRolePerm{}).Create(&permModel).Error
		if err != nil {
			failItems = append(failItems, fmt.Sprintf("为角色%+v添加权限%d失败，错误%s", roleName, perm, err.Error()))
		} else {
			successCount++
		}
	}
	return
}
