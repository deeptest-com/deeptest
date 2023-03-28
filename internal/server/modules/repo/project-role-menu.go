package repo

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
)

type ProjectRoleMenuRepo struct {
	DB *gorm.DB `inject:""`
}

func NewProjectRoleMenuRepo() *ProjectRoleMenuRepo {
	return &ProjectRoleMenuRepo{}
}

func (r *ProjectRoleMenuRepo) GetRoleMenuList(roleId uint) (roleMenus []model.ProjectRoleMenu, err error) {
	err = r.DB.Model(&model.ProjectRoleMenu{}).
		Where("project_role_id = ?", roleId).
		Scan(&roleMenus).Error

	return
}
