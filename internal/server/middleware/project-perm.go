package middleware

import (
	"errors"
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/aaronchen2k/deeptest/internal/server/core/casbin"
	"github.com/aaronchen2k/deeptest/internal/server/core/dao"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/snowlyg/helper/arr"
	"github.com/snowlyg/multi"
	"go.uber.org/zap"
	"net/http"
	"strconv"
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

type SysUser struct {
	BaseModel

	v1.UserBase

	Password string              `gorm:"type:varchar(250)" json:"password" validate:"required"`
	Vcode    string              `json:"vcode"`
	Profile  SysUserProfileModel `gorm:"foreignKey:user_id"`

	RoleIds []uint `gorm:"-" json:"role_ids"`
}

func (SysUser) TableName() string {
	return "sys_user"
}

type SysRole struct {
	BaseModel
	v1.RoleBase
}

func (SysRole) TableName() string {
	return "sys_role"
}

// ProjectPerm  项目权限权鉴中间件
func ProjectPerm() iris.Handler {
	return func(ctx *context.Context) {
		userId := multi.GetUserId(ctx)

		isAdminUser, err := IsAdminUser(userId)
		if err != nil {
			ctx.JSON(_domain.Response{Code: _domain.AuthActionErr.Code, Data: nil, Msg: "系统异常，请重新登录或者联系管理员"})
			ctx.StopExecution()
			return
		}
		if !isAdminUser {
			check, err := CheckProjectPerm(ctx.Request(), userId)
			if err != nil || !check {
				ctx.JSON(_domain.Response{Code: _domain.AuthActionErr.Code, Data: nil, Msg: "你未拥有当前项目操作权限，请联系管理员"})
				ctx.StopExecution()
				return
			}
		}

		ctx.Next()
	}
}

func CheckProjectPerm(r *http.Request, userId uint) (bool, error) {
	method := r.Method
	path := r.URL.Path

	projectPerm, err := GetProjectPerm(path, method)
	if err != nil {
		logUtils.Errorf(fmt.Sprintf("项目权限不存在：%d-%s-%s", userId, path, method), zap.Any("project-perm-err", err.Error()))
		return false, err
	}
	if projectPerm.ID == 0 {
		logUtils.Errorf(fmt.Sprintf("项目权限不存在：%d-%s-%s", userId, path, method))
		return false, err
	}

	projectMemberRole, err := GetUserCurrentRole(userId)
	if err != nil {
		logUtils.Errorf(fmt.Sprintf("用户角色不存在：%d-%s-%s", userId, path, method), zap.Any("user-role-in-project-err", err.Error()))
		return false, err
	}

	_, err = GetProjectRolePerm(projectMemberRole.ProjectRoleId, projectPerm.ID)
	if err != nil {
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

func IsAdminUser(id uint) (bool, error) {
	user, err := FindUserDetailById(id)
	if err != nil {
		return false, err
	}

	return arr.InArrayS(user.SysRoles, serverConsts.AdminRoleName), nil
}

func FindUserDetailById(id uint) (user v1.UserResp, err error) {
	user, err = FindUserById(id)
	if err != nil {
		return user, err
	}

	GetSysRoles(&user)

	return user, nil
}

func FindUserById(id uint) (user v1.UserResp, err error) {
	err = dao.GetDB().Model(&SysUser{}).Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}

	return
}

func FindInId(ids []string) (roles []v1.RoleResp, error error) {
	err := dao.GetDB().Model(&SysRole{}).Where("id in ?", ids).Find(&roles).Error
	if err != nil {
		logUtils.Errorf("通过ids查询角色错误", zap.String("错误:", err.Error()))
		return
	}
	return
}

func GetSysRoles(users ...*v1.UserResp) {
	var roleIds []string
	userRoleIds := make(map[uint][]string, 10)

	if len(users) == 0 {
		return
	}

	for _, user := range users {
		user.ToString()
		userRoleId := casbin.GetRolesForUser(user.Id)
		userRoleIds[user.Id] = userRoleId
		roleIds = append(roleIds, userRoleId...)
	}

	roles, err := FindInId(roleIds)
	if err != nil {
		logUtils.Errorf("get role get err ", zap.String("错误:", err.Error()))
	}

	for _, user := range users {
		for _, role := range roles {
			sRoleId := strconv.FormatInt(int64(role.Id), 10)
			if arr.InArrayS(userRoleIds[user.Id], sRoleId) {
				user.SysRoles = append(user.SysRoles, role.Name)
			}
		}
	}
}
