package repo

import (
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/core/dao"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ProjectRolePermRepo struct {
	DB              *gorm.DB         `inject:""`
	ProjectRepo     *ProjectRepo     `inject:""`
	ProjectRoleRepo *ProjectRoleRepo `inject:""`
}

func (r *ProjectRolePermRepo) PaginateRolePerms(req v1.ProjectRolePermPaginateReq) (data _domain.PageData, err error) {
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

func (r *ProjectRolePermRepo) UserPermList(req v1.ProjectUserPermsPaginate, userId uint) (data _domain.PageData, err error) {
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
// TODO: 后续可能会有其他需求，禁止某些角色访问部分接口，需要对这部分代码进行优化，并改成配置化
func (r *ProjectRolePermRepo) GetProjectPermsForRole() (res map[consts.RoleType][]uint, err error) {
	var permIds, excludePermIds []uint

	//查找所有ID，放在permIds中
	err = r.DB.Model(&model.ProjectPerm{}).Select("id").Find(&permIds).Error
	if err != nil {
		return
	}

	//需要特殊处理的接口路径（当某些角色不能访问某些接口时，通过这块把id提出来）
	apiPath := [...]string{"/api/v1/users/invite", "/api/v1/projects/changeUserRole"}
	err = r.DB.Model(&model.ProjectPerm{}).Select("id").Where("name in ?", apiPath).Find(&excludePermIds).Error
	if err != nil {
		return
	}
	res = make(map[consts.RoleType][]uint)
	// 遍历枚举类型，并将每个枚举值添加到映射中
	for role := range map[consts.RoleType]struct{}{
		consts.Admin:          {},
		consts.User:           {},
		consts.Tester:         {},
		consts.Developer:      {},
		consts.ProductManager: {},
	} {
		res[role] = []uint{} // 将枚举值作为键，空结构体作为值存储在映射中
	}

	for role := range res {
		//复制permIds的值给tmp，避免传递索引，导致最终改动一个role，其他都变化
		tmp := make([]uint, len(permIds)) // 创建一次tmp切片，然后对其进行修改
		copy(tmp, permIds)                // 复制permIds的值给tmp

		if role == consts.Developer || role == consts.Tester || role == consts.User {
			// 移除与excludePermIds中的路径ID相匹配的权限ID
			tmp = excludePerms(tmp, excludePermIds)
		}
		res[role] = tmp
	}

	return
}

func excludePerms(perms []uint, excludePerms []uint) []uint {
	for _, perm := range excludePerms {
		for i := 0; i < len(perms); i++ {
			if perms[i] == perm {
				perms = append(perms[:i], perms[i+1:]...)
				i-- // 调整索引，确保不跳过元素
			}
		}
	}
	return perms
}

func (r *ProjectRolePermRepo) AddPermForProjectRole(roleName consts.RoleType, perms []uint) (successCount int, failItems []string) {
	projectRole, err := r.ProjectRoleRepo.FindByName(roleName)
	if err != nil {
		failItems = append(failItems, fmt.Sprintf("为角色%+v添加权限%+v失败，错误%s", roleName, perms, err.Error()))
		return
	}

	projectRoleId := projectRole.ID
	err = r.DB.Delete(&model.ProjectRolePerm{}, "project_role_id = ?", projectRoleId).Error
	if err != nil {
		failItems = append(failItems, fmt.Sprintf("为角色%+v添加权限%+v失败，错误%s", roleName, perms, err.Error()))
		return
	}

	for _, perm := range perms {
		permModel := &model.ProjectRolePerm{ProjectRolePermBase: v1.ProjectRolePermBase{ProjectRoleId: projectRoleId, ProjectPermId: perm}}
		err := r.DB.Model(&model.ProjectRolePerm{}).Create(&permModel).Error
		if err != nil {
			failItems = append(failItems, fmt.Sprintf("为角色%+v添加权限%d失败，错误%s", roleName, perm, err.Error()))
		} else {
			successCount++
		}
	}
	return
}
