package repo

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/core/dao"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ProjectRolePermRepo struct {
	DB          *gorm.DB     `inject:""`
	ProjectRepo *ProjectRepo `inject:""`
}

func NewProjectRolePermRepo() *ProjectRolePermRepo {
	return &ProjectRolePermRepo{}
}

func (r *ProjectRolePermRepo) PaginateRolePerms(req domain.ProjectRolePermPaginateReq) (data _domain.PageData, err error) {
	var count int64
	projectPerms := make([]*model.ProjectPerm, 0)

	db := r.DB.Model(&model.ProjectPerm{}).Joins("JOIN biz_project_role_perm p ON biz_project_perm.id=p.project_perm_id AND p.project_role_id = ?", req.RoleId)

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

func (r *ProjectRolePermRepo) UserPermList(req domain.ProjectUserPermsPaginate, userId uint) (data _domain.PageData, err error) {
	currProject, err := r.ProjectRepo.GetCurrProjectByUser(userId)
	if err != nil {
		logUtils.Errorf("query project profile error", zap.String("error:", err.Error()))
		return
	}

	var roleId uint
	r.DB.Model(&model.ProjectMember{}).
		Select("project_role_id").
		Where("user_id = ?", userId).
		Where("project_id = ? AND NOT deleted", currProject.ID).First(&roleId)

	projectPerms := make([]*model.ProjectPerm, 0)

	db := r.DB.Model(&model.ProjectPerm{}).Joins("JOIN biz_project_role_perm p ON biz_project_perm.id=p.project_perm_id AND p.project_role_id = ?", roleId)

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

func (r *ProjectRolePermRepo) GetByRoleAndPerm(roleId, permId uint) (ret model.ProjectRolePerm, err error) {
	err = r.DB.Model(&model.ProjectRolePerm{}).
		Where("project_role_id = ?", roleId).
		Where("project_perm_id = ?", permId).
		First(&ret).Error
	return
}

// GetProjectPermsForRole TODO: 每个角色需要的权限还未确定，需要改动
func (r *ProjectRolePermRepo) GetProjectPermsForRole() (res map[uint][]uint, err error) {
	var permIds, testPermIds []uint
	err = r.DB.Model(&model.ProjectPerm{}).Select("id").Find(&permIds).Error
	err = r.DB.Model(&model.ProjectPerm{}).Select("id").Where("name like ?", "/api/v1/projects%").Find(&testPermIds).Error
	res = map[uint][]uint{
		1: permIds,
		2: permIds,
		3: testPermIds,
	}

	return
}

func (r *ProjectRolePermRepo) AddPermForProjectRole(id uint, perms []uint) (successCount int, failItems []string) {
	err := r.DB.Delete(&model.ProjectRolePerm{}, "project_role_id = ?", id).Error
	if err != nil {
		failItems = append(failItems, fmt.Sprintf("为角色%d添加权限%+v失败，错误%s", id, perms, err.Error()))
		return
	}
	for _, perm := range perms {
		permModel := &model.ProjectRolePerm{ProjectRolePermBase: domain.ProjectRolePermBase{ProjectRoleId: id, ProjectPermId: perm}}
		err := r.DB.Model(&model.ProjectRolePerm{}).Create(&permModel).Error
		if err != nil {
			failItems = append(failItems, fmt.Sprintf("为角色%d添加权限%d失败，错误%s", id, perm, err.Error()))
		} else {
			successCount++
		}
	}
	return
}
