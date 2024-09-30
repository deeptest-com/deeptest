package serverDomain

import (
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/pkg/domain"
)

type ProjectMenuBase struct {
	Code     string `gorm:"index:code_index,unique;not null;type:varchar(256)" json:"code"`
	Title    string `gorm:"type:varchar(256)" json:"title"`
	Path     string `gorm:"type:varchar(256)" json:"path"`
	Type     string `gorm:"type:varchar(100)" json:"type"`
	ParentId uint   `json:"parentId"`
}

type ProjectMenuReq struct {
	_domain.Model
	ProjectMenuBase
}

type ProjectMenuConfig struct {
	Code   string `json:"code"`
	Title  string `json:"title"`
	Path   string `json:"path"`
	Type   string `json:"type"`
	Parent string `json:"parent"`
}

type ProjectRoleMenuConfig struct {
	RoleName consts.RoleType `json:"role_name"`
	Menus    []string        `json:"menus"`
}
