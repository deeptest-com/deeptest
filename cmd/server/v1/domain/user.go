package serverDomain

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/snowlyg/helper/str"
	"regexp"
)

type UserReq struct {
	_domain.Model
	UserBase
}

type UserReqPaginate struct {
	_domain.PaginateReq
	Name     string `json:"name"`
	UserName string `json:"userName"`
}

type MemberResp struct {
	Id            uint            `json:"id"`
	Username      string          `json:"username"`
	Name          string          `json:"name"`
	Email         string          `json:"email"`
	RoleName      consts.RoleType `json:"roleName"`
	ProjectRoleId uint            `json:"roleId"`
}

type UserResp struct {
	_domain.Model
	UserBase

	SysRoles     []string                 `gorm:"-" json:"sysRoles"`
	ProjectRoles map[uint]consts.RoleType `gorm:"-" json:"projectRoles"`
}

type UpdateUserReq struct {
	Username    string
	Email       string
	Password    string
	NewPassword string
}

type InviteUserReq struct {
	UserId    uint
	Email     string
	ProjectId int
	RoleName  consts.RoleType
}

type UserBase struct {
	Username  string `gorm:"uniqueIndex;not null;type:varchar(60)" json:"username" validate:"required"`
	Name      string `gorm:"index;not null; type:varchar(60)" json:"name"`
	Email     string `gorm:"index;not null; type:varchar(60)" json:"email"`
	Intro     string `gorm:"not null; type:varchar(512)" json:"intro"`
	Avatar    string `gorm:"type:varchar(1024)" json:"avatar"`
	ImAccount string `json:"imAccount"`
	Password  string `json:"password"`
	RoleIds   []uint `gorm:"-" json:"role_ids"`
	Type      bool   `json:"type"`
}

func (res *UserResp) ToString() {
	if res.Avatar == "" {
		return
	}
	re := regexp.MustCompile("^http")
	if !re.MatchString(res.Avatar) {
		res.Avatar = str.Join("http://127.0.0.1:8085/upload/", res.Avatar)
	}
}

type UpdateUserRoleReq struct {
	UserId  uint   `json:"userId" validate:"required"`
	RoleIds []uint `json:"roleIds"`
}
