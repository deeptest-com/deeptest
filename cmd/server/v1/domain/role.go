package serverDomain

import (
	"github.com/aaronchen2k/deeptest/pkg/domain"
)

type RoleReq struct {
	_domain.Model
	RoleBase
}

type RoleReqPaginate struct {
	_domain.PaginateReq
	Name string `json:"name"`
}

type RoleResp struct {
	_domain.Model
	RoleBase
}

type RoleBase struct {
	Name        string `gorm:"uniqueIndex;not null; type:varchar(256)" json:"name" validate:"required,gte=4,lte=50" comment:"名称"`
	DisplayName string `gorm:"type:varchar(256)" json:"displayName" comment:"显示名称"`
	Description string `gorm:"type:varchar(256)" json:"description" comment:"描述"`

	Perms [][]string `gorm:"-" json:"perms"`
}

type RoleMenuConfig struct {
	RoleName string   `json:"role_name"`
	Menus    []string `json:"menus"`
}
