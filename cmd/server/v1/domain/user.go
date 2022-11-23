package domain

import (
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
	Name string `json:"name"`
}

type UserResp struct {
	_domain.Model
	UserBase
	Roles []string `gorm:"-" json:"roles"`
}

type UpdateUserReq struct {
	Username    string
	Email       string
	Password    string
	NewPassword string
}

type UserBase struct {
	Username string `gorm:"uniqueIndex;not null;type:varchar(60)" json:"username" validate:"required"`
	Name     string `gorm:"index;not null; type:varchar(60)" json:"name"`
	Email    string `gorm:"index;not null; type:varchar(60)" json:"email"`
	Intro    string `gorm:"not null; type:varchar(512)" json:"intro"`
	Avatar   string `gorm:"type:varchar(1024)" json:"avatar"`

	Password string `json:"password"`
	RoleIds  []uint `gorm:"-" json:"role_ids"`
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
