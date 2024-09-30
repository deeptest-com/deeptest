package serverDomain

import _domain "github.com/deeptest-com/deeptest/pkg/domain"

type ProjectRolePermBase struct {
	ProjectRoleId uint `gorm:"index:index_project_role,unique;not null" json:"project_role_id"`
	ProjectPermId uint `gorm:"index:index_project_role,unique;not null" json:"project_perm_id"`
}

type ProjectRolePermPaginateReq struct {
	_domain.PaginateReq
	RoleId uint `json:"role_id"`
}
