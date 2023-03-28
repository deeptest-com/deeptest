package domain

import _domain "github.com/aaronchen2k/deeptest/pkg/domain"

type ProjectRolePermBase struct {
	ProjectRoleId uint `gorm:"index:index_project_role_id" json:"project_role_id"`
	ProjectPermId uint `gorm:"index:index_project_perm_id" json:"project_perm_id"`
}

type ProjectRolePermPaginateReq struct {
	_domain.PaginateReq
	RoleId uint `json:"role_id"`
}
