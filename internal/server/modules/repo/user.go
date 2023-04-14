package repo

import (
	"errors"
	"fmt"
	"github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/aaronchen2k/deeptest/internal/server/core/casbin"
	"github.com/aaronchen2k/deeptest/internal/server/core/dao"
	model "github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	_stringUtils "github.com/aaronchen2k/deeptest/pkg/lib/string"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"strings"

	"github.com/snowlyg/helper/arr"
	"github.com/snowlyg/multi"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB              *gorm.DB         `inject:""`
	ProfileRepo     *ProfileRepo     `inject:""`
	RoleRepo        *RoleRepo        `inject:""`
	ProjectRepo     *ProjectRepo     `inject:""`
	EnvironmentRepo *EnvironmentRepo `inject:""`
	ProjectRoleRepo *ProjectRoleRepo `inject:""`
}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (r *UserRepo) Paginate(req domain.UserReqPaginate) (data _domain.PageData, err error) {
	var count int64

	db := r.DB.Model(&model.SysUser{})
	if len(req.Name) > 0 {
		db = db.Where("name LIKE ?", fmt.Sprintf("%s%%", req.Name))
	}
	if len(req.UserName) > 0 {
		db = db.Where("username LIKE ?", fmt.Sprintf("%s%%", req.UserName))
	}

	err = db.Count(&count).Error
	if err != nil {
		logUtils.Errorf("获取用户总数错误", zap.String("错误:", err.Error()))
		return
	}

	users := make([]*domain.UserResp, 0)
	err = db.Scopes(dao.PaginateScope(req.Page, req.PageSize, req.Order, req.Field)).
		Find(&users).Error
	if err != nil {
		logUtils.Errorf("获取用户分页数据错误", zap.String("错误:", err.Error()))
		return
	}

	// 查询用户角色
	r.GetSysRoles(users...)

	data.Result = users
	data.Populate(users, count, req.Page, req.PageSize)

	return
}

// GetSysRoles
func (r *UserRepo) GetSysRoles(users ...*domain.UserResp) {
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

	roles, err := r.RoleRepo.FindInId(roleIds)
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

// getRoles
func (r *UserRepo) GetProjectRoles(users ...*domain.UserResp) {
	if len(users) == 0 {
		return
	}

	for _, user := range users {
		projectRoles, err := r.ProjectRepo.FindRolesByUser(user.Id)
		if err != nil {
			break
		}

		user.ProjectRoles = map[uint]consts.RoleType{}
		for _, projectRole := range projectRoles {
			user.ProjectRoles[projectRole.ProjectId] = projectRole.ProjectRoleName
		}
	}
}

func (r *UserRepo) FindByUserName(username string, ids ...uint) (domain.UserResp, error) {
	user := domain.UserResp{}
	db := r.DB.Model(&model.SysUser{}).Where("username = ?", username)

	if len(ids) == 1 {
		db.Where("id != ?", ids[0])
	}

	err := db.First(&user).Error
	if err != nil {
		return user, err
	}

	r.GetSysRoles(&user)
	return user, nil
}

func (r *UserRepo) FindByEmail(email string, ids ...uint) (domain.UserResp, error) {
	user := domain.UserResp{}
	db := r.DB.Model(&model.SysUser{}).Where("email = ?", email)

	if len(ids) == 1 {
		db.Where("id != ?", ids[0])
	}

	err := db.First(&user).Error
	if err != nil {
		return user, err
	}

	r.GetSysRoles(&user)
	return user, nil
}

func (r *UserRepo) FindPasswordByUserName(username string, ids ...uint) (domain.LoginResp, error) {
	user := domain.LoginResp{}
	db := r.DB.Model(&model.SysUser{}).Select("id,password").Where("username = ?", username)
	if len(ids) == 1 {
		db.Where("id != ?", ids[0])
	}
	err := db.First(&user).Error
	if err != nil {
		logUtils.Errorf("根据用户名查询用户错误", zap.String("用户名:", username), zap.Uints("ids:", ids), zap.String("错误:", err.Error()))
		return user, err
	}

	return user, nil
}

func (r *UserRepo) FindPasswordByEmail(email string) (domain.LoginResp, error) {
	user := domain.LoginResp{}
	db := r.DB.Model(&model.SysUser{}).Select("id,password").Where("email = ?", email)

	err := db.First(&user).Error
	if err != nil {
		logUtils.Errorf("根据邮箱查询用户错误", zap.String("邮箱:", email), zap.String("错误:", err.Error()))
		return user, err
	}

	return user, nil
}

func (r *UserRepo) FindByUserNameAndVcode(username, vcode string) (user model.SysUser, err error) {
	err = r.DB.Model(&model.SysUser{}).Where("username = ? AND vcode = ?", username, vcode).
		First(&user).Error

	if err != nil {
		return user, err
	}

	return
}

func (r *UserRepo) Register(user *model.SysUser) (err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	user.Password = string(hash)
	err = r.DB.Model(&model.SysUser{}).Create(&user).Error
	if err != nil {
		return
	}

	project, err := r.AddProjectForUser(user)
	if err != nil {
		return
	}

	err = r.AddProfileForUser(user, project.ID)
	if err != nil {
		return
	}

	err = r.AddRoleForUser(user)
	if err != nil {
		return
	}

	return
}

func (r *UserRepo) Create(req domain.UserReq) (uint, error) {
	if _, err := r.FindByUserName(req.Username); !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, fmt.Errorf("用户名 %s 已经被使用", req.Username)
	}

	user := model.SysUser{UserBase: req.UserBase, RoleIds: req.RoleIds}
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	user.Password = string(hash)
	err = r.DB.Model(&model.SysUser{}).Create(&user).Error
	if err != nil {
		return 0, err
	}

	project, err := r.AddProjectForUser(&user)
	if err != nil {
		return 0, err
	}

	if err := r.AddProfileForUser(&user, project.ID); err != nil {
		return 0, err
	}

	if len(user.RoleIds) == 0 {
		role, _ := r.RoleRepo.FindByName("user")
		user.RoleIds = append(user.RoleIds, role.Id)
	}

	if err := r.AddRoleForUser(&user); err != nil {
		return 0, err
	}

	return user.ID, nil
}

func (r *UserRepo) Update(id uint, req domain.UserReq) error {
	if b, err := r.IsAdminUser(id); err != nil {
		return err
	} else if b {
		return errors.New("不能编辑超级管理员")
	}
	if _, err := r.FindByUserName(req.Username, id); !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	user := model.SysUser{UserBase: req.UserBase}
	if req.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(hash)
	}

	err := r.DB.Model(&model.SysUser{}).Where("id = ?", id).Updates(&user).Error
	if err != nil {
		logUtils.Errorf("更新用户错误", zap.String("错误:", err.Error()))
		return err
	}

	if err := r.AddRoleForUser(&user); err != nil {
		logUtils.Errorf("添加用户角色错误", zap.String("错误:", err.Error()))
		return err
	}

	return nil
}

func (r *UserRepo) InviteToProject(req domain.InviteUserReq) (user model.SysUser, err error) {
	user, err = r.GetByUserId(req.UserId)
	if err != nil {
		err = errors.New("用户不存在，请先创建用户")
		return
	}

	projectMemberRole, err := r.ProjectRepo.FindRolesByProjectAndUser(uint(req.ProjectId), user.ID)
	if projectMemberRole.ID != 0 {
		err = errors.New("用户已经存在于项目中")
		return
	}
	var roleName consts.RoleType
	if req.RoleName == "" {
		roleName = "user"
	} else {
		role, err := r.ProjectRoleRepo.FindByName(req.RoleName)
		if err != nil || role.ID == 0 {
			err = errors.New("角色不存在")
			return user, err
		}
		roleName = req.RoleName
	}

	err = r.ProjectRepo.AddProjectMember(uint(req.ProjectId), user.ID, roleName)
	if err != nil {
		return
	}

	return
}

func (r *UserRepo) IsAdminUser(id uint) (bool, error) {
	user, err := r.FindDetailById(id)
	if err != nil {
		return false, err
	}

	return arr.InArrayS(user.SysRoles, serverConsts.AdminRoleName), nil
}

func (r *UserRepo) FindById(id uint) (user domain.UserResp, err error) {
	err = r.DB.Model(&model.SysUser{}).Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}

	return
}
func (r *UserRepo) FindDetailById(id uint) (user domain.UserResp, err error) {
	user, err = r.FindById(id)
	if err != nil {
		return user, err
	}

	r.GetSysRoles(&user)
	r.GetProjectRoles(&user)

	return user, nil
}

func (r *UserRepo) GetByUsernameOrPassword(usernameOrPassword string) (user model.SysUser, err error) {
	err = r.DB.Model(&model.SysUser{}).
		Where("NOT deleted").
		Where("username = ? OR email = ?", usernameOrPassword, usernameOrPassword).
		First(&user).Error

	if err != nil {
		return
	}

	return
}

func (r *UserRepo) GetByUserId(id uint) (user model.SysUser, err error) {
	err = r.DB.Model(&model.SysUser{}).
		Where("NOT deleted").
		Where("id = ?", id).
		First(&user).Error

	if err != nil {
		return
	}

	return
}

func (r *UserRepo) DeleteById(id uint) error {
	err := r.DB.Unscoped().Delete(&model.SysUser{}, id).Error
	if err != nil {
		logUtils.Errorf("delete user by id get  err ", zap.String("错误:", err.Error()))
		return err
	}
	return nil
}

func (r *UserRepo) AddProfileForUser(user *model.SysUser, projectId uint) (err error) {
	_, err = r.ProfileRepo.FindByUserId(user.ID)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("用户 %s 信息已经被使用", user.Name)
	}

	profile := model.SysUserProfile{UserId: user.ID, CurrProjectId: projectId}
	err = r.DB.Create(&profile).Error
	if err != nil {
		logUtils.Errorf("添加用户错误", zap.String("错误:", err.Error()))
		return err
	}

	return
}

// AddRoleForUser add roles for user
func (r *UserRepo) AddRoleForUser(user *model.SysUser) error {
	userId := strconv.FormatUint(uint64(user.ID), 10)
	oldRoleIds, err := casbin.Instance().GetRolesForUser(userId)
	if err != nil {
		logUtils.Errorf("获取用户角色错误", zap.String("错误:", err.Error()))
		return err
	}

	if len(oldRoleIds) > 0 {
		if _, err := casbin.Instance().DeleteRolesForUser(userId); err != nil {
			logUtils.Errorf("添加角色到用户错误", zap.String("错误:", err.Error()))
			return err
		}
	}
	if len(user.RoleIds) == 0 {
		role, _ := r.RoleRepo.FindByName("user")
		user.RoleIds = append(user.RoleIds, role.Id)
	}

	var roleIds []string
	for _, userRoleId := range user.RoleIds {
		roleIds = append(roleIds, strconv.FormatUint(uint64(userRoleId), 10))
	}

	if _, err := casbin.Instance().AddRolesForUser(userId, roleIds); err != nil {
		logUtils.Errorf("添加角色到用户错误", zap.String("错误:", err.Error()))
		return err
	}

	return nil
}

func (r *UserRepo) AddProjectForUser(user *model.SysUser) (project model.Project, err error) {
	_, err = r.ProjectRepo.GetCurrProjectByUser(user.ID)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		err = fmt.Errorf("用户%s的默认项目已存在", user.Name)
		return
	}

	project = model.Project{ProjectBase: domain.ProjectBase{Name: "默认项目"}}
	err = r.DB.Create(&project).Error
	if err != nil {
		logUtils.Errorf("添加项目错误", zap.String("错误:", err.Error()))
		return
	}

	err = r.ProjectRepo.AddProjectMember(project.ID, user.ID, "admin")
	if err != nil {
		logUtils.Errorf("添加项目角色错误", zap.String("错误:", err.Error()))
		return
	}

	err = r.EnvironmentRepo.AddDefaultForProject(project.ID)
	if err != nil {
		logUtils.Errorf("添加项目默认环境错误", zap.String("错误:", err.Error()))
		return
	}

	serve, err := r.ProjectRepo.AddProjectDefaultServe(project.ID, user.ID)
	if err != nil {
		logUtils.Errorf("添加默认服务错误", zap.String("错误:", err.Error()))
		return
	}

	err = r.ProjectRepo.AddProjectRootInterface(project.ID, serve.ID)
	if err != nil {
		logUtils.Errorf("添加接口错误", zap.String("错误:", err.Error()))
		return
	}

	return
}

// DelToken 删除token
func (r *UserRepo) DelToken(token string) error {
	err := multi.AuthDriver.DelUserTokenCache(token)
	if err != nil {
		logUtils.Errorf("del token", zap.Any("err", err))
		return fmt.Errorf("del token %w", err)
	}
	return nil
}

// CleanToken 清空 token
func (r *UserRepo) CleanToken(authorityType int, userId string) error {
	err := multi.AuthDriver.CleanUserTokenCache(authorityType, userId)
	if err != nil {
		logUtils.Errorf("clean token", zap.Any("err", err))
		return fmt.Errorf("clean token %w", err)
	}
	return nil
}

func (r *UserRepo) UpdatePasswordByName(name string, password string) (err error) {
	err = r.DB.Model(&model.SysUser{}).Where("username = ?", name).
		Updates(map[string]interface{}{"password": password}).Error
	if err != nil {
		logUtils.Errorf("更新用户错误", zap.String("错误:", err.Error()))
		return err
	}

	return nil
}
func (r *UserRepo) UpdateAvatar(id uint, avatar string) error {
	return nil
}

func (r *UserRepo) UpdateEmail(email string, id uint) (err error) {
	err = r.DB.Model(&model.SysUser{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{"email": email}).Error
	if err != nil {
		logUtils.Errorf("更新用户邮箱错误 %s", err.Error())
		return err
	}

	return
}

func (r *UserRepo) UpdateName(username string, id uint) (err error) {
	err = r.DB.Model(&model.SysUser{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{"username": username}).Error
	if err != nil {
		logUtils.Errorf("更新用户名称错误 %s", err.Error())
		return err
	}

	return
}

func (r *UserRepo) ChangePassword(req domain.UpdateUserReq, id uint) (err error) {
	user, err := r.FindById(id)
	if err != nil {
		if err != nil {
			return
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		err = errors.New("原有密码错误")
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	req.NewPassword = string(hash)

	err = r.DB.Model(&model.SysUser{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{"password": req.NewPassword}).Error
	if err != nil {
		return err
	}

	return
}

func (r *UserRepo) UpdatePassword(password string, id uint) (err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	newPassword := string(hash)

	err = r.DB.Model(&model.SysUser{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{"password": newPassword}).Error

	if err != nil {
		return err
	}

	return
}

func (r *UserRepo) GenAndUpdateVcode(id uint) (vcode string, err error) {
	vcode = strings.ToLower(_stringUtils.RandStr(6))

	err = r.DB.Model(&model.SysUser{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{"vcode": vcode}).Error
	if err != nil {
		return
	}

	return
}

func (r *UserRepo) ClearVcode(id uint) (err error) {
	err = r.DB.Model(&model.SysUser{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{"vcode": ""}).Error

	if err != nil {
		return
	}

	return
}
