package middleware

import (
	"errors"
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/core/dao"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/snowlyg/multi"
	"go.uber.org/zap"
	"net/http"
	"strings"
	"time"
)

type BaseModel struct {
	ID        uint       `gorm:"primary_key" sql:"type:INT(10) UNSIGNED NOT NULL" json:"id"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	Deleted  bool `json:"-" gorm:"default:false"`
	Disabled bool `json:"disabled,omitempty" gorm:"default:false"`
}

type ProjectMemberModel struct {
	BaseModel

	ProjectId     uint `json:"projectId"`
	ProjectRoleId uint `json:"projectRoleId"`
	UserId        uint `json:"userId"`
}

func (ProjectMemberModel) TableName() string {
	return "biz_project_member"
}

type ProjectPermModel struct {
	BaseModel
	Name        string `gorm:"index:perm_index,unique;not null ;type:varchar(200)" json:"name" validate:"required,gte=4,lte=50"`
	DisplayName string `gorm:"type:varchar(256)" json:"displayName"`
	Description string `gorm:"type:varchar(256)" json:"description"`
	Act         string `gorm:"index:perm_index;type:varchar(50)" json:"act" validate:"required"`
}

func (ProjectPermModel) TableName() string {
	return "biz_project_perm"
}

type SysUserProfileModel struct {
	BaseModel

	Phone         string `json:"phone"`
	CurrProjectId uint   `json:"currProjectId"`

	UserId uint `json:"userId"`
}

func (SysUserProfileModel) TableName() string {
	return "sys_user_profile"
}

type ProjectRolePermModel struct {
	BaseModel
	v1.ProjectRolePermBase
}

func (ProjectRolePermModel) TableName() string {
	return "biz_project_role_perm"
}

// ProjectPerm  项目权限权鉴中间件
func ProjectPerm() iris.Handler {
	return func(ctx *context.Context) {
		userId := multi.GetUserId(ctx)
		check, err := CheckProjectPerm(ctx.Request(), userId)
		if err != nil || !check {
			ctx.JSON(_domain.Response{Code: _domain.AuthActionErr.Code, Data: nil, Msg: "你未拥有当前操作权限，请联系管理员"})
			ctx.StopExecution()
			return
		}

		ctx.Next()
	}
}

func CheckProjectPerm(r *http.Request, userId uint) (bool, error) {
	method := r.Method
	path := r.URL.Path

	projectPerm, err := GetProjectPerm(path, method)
	if err != nil || projectPerm.ID == 0 {
		logUtils.Errorf(fmt.Sprintf("项目权限不存在：%d-%s-%s", userId, path, method), zap.Any("project-perm-err", err.Error()))
		return false, err
	}

	projectMemberRole, err := GetUserCurrentRole(userId)
	if err != nil || projectMemberRole.ID == 0 {
		logUtils.Errorf(fmt.Sprintf("用户角色不存在：%d-%s-%s", userId, path, method), zap.Any("user-role-in-project-err", err.Error()))
		return false, err
	}

	projectRolePerm, err := GetProjectRolePerm(projectMemberRole.ProjectRoleId, projectPerm.ID)
	if err != nil || projectRolePerm.ID == 0 {
		logUtils.Errorf(fmt.Sprintf("用户没有该权限：%d-%s-%s", userId, path, method), zap.Any("project-role-perm-err", err.Error()))
		return false, err
	}
	return true, nil
}

func GetProjectPerm(path, method string) (res ProjectPermModel, err error) {
	pathArr := strings.Split(path, "/")
	if len(pathArr) < 4 {
		err = errors.New("path is invalid")
		return
	}
	pathArrTmp := pathArr[:4]
	modulePath := strings.Join(pathArrTmp, "/")

	var projectPerms []ProjectPermModel
	err = dao.GetDB().Model(&ProjectPermModel{}).
		Where("name like ?", fmt.Sprintf("%s%%", modulePath)).
		Where("act = ?", method).
		Find(&projectPerms).Error
	if err != nil {
		return
	}

OuterLoop:
	for _, v := range projectPerms {
		tablePathArr := strings.Split(v.Name, "/")
		if len(tablePathArr) != len(pathArr) {
			continue
		}
		for k1, s1 := range tablePathArr {
			if !strings.ContainsAny(s1, ":") && pathArr[k1] != s1 {
				continue OuterLoop
			}
		}
		res = v
	}

	return
}

func GetUserCurrentRole(userId uint) (data ProjectMemberModel, err error) {
	db := dao.GetDB()
	var userProfile SysUserProfileModel
	err = db.Model(&SysUserProfileModel{}).
		Where("user_id = ?", userId).
		First(&userProfile).Error
	if err != nil {
		return
	}

	err = db.Model(&ProjectMemberModel{}).
		Where("user_id = ?", userId).
		Where("project_id = ?", userProfile.CurrProjectId).
		First(&data).Error
	return
}

func GetProjectRolePerm(roleId, permId uint) (data ProjectRolePermModel, err error) {
	err = dao.GetDB().Model(ProjectRolePermModel{}).
		Where("project_role_id = ?", roleId).
		Where("project_perm_id = ?", permId).
		First(&data).Error
	return
}
