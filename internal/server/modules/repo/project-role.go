package repo

import (
	"errors"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"gorm.io/gorm"
)

type ProjectRoleRepo struct {
	DB          *gorm.DB     `inject:""`
	ProjectRepo *ProjectRepo `inject:""`
}

func NewProjectRoleRepo() *ProjectRoleRepo {
	return &ProjectRoleRepo{}
}

func (r *ProjectRoleRepo) GetAdminRecord() (projectRole model.ProjectRole, err error) {
	db := r.DB.Model(&model.ProjectRole{}).Where("name='admin'").Order("id ASC")
	err = db.First(&projectRole).Error
	return
}
func (r *ProjectRoleRepo) GetUserRecord() (projectRole model.ProjectRole, err error) {
	db := r.DB.Model(&model.ProjectRole{}).Where("name='user'").Order("id ASC")
	err = db.First(&projectRole).Error
	return
}

func (r *ProjectRoleRepo) FindById(id uint) (projectRole model.ProjectRole, err error) {
	db := r.DB.Model(&model.ProjectRole{}).Where("id = ?", id)

	err = db.First(&projectRole).Error
	return
}

func (r *ProjectRoleRepo) FindByName(name consts.RoleType) (projectRole model.ProjectRole, err error) {
	db := r.DB.Model(&model.ProjectRole{}).Where("name = ?", name)

	err = db.First(&projectRole).Error
	return
}

func (r *ProjectRoleRepo) Create(projectRole model.ProjectRole) (err error) {
	_, err = r.FindByName(projectRole.Name)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		logUtils.Errorf("项目角色已经存在")
		return
	}

	err = r.DB.Create(&projectRole).Error
	if err != nil {
		logUtils.Errorf("创建项目角色失败%s", err.Error())
		return
	}

	return
}

func (r *ProjectRoleRepo) ProjectUserRoleList(userId, projectId uint) (projectRole model.ProjectRole, err error) {
	//获取用户在项目中拥有的角色
	projectMemberRole, err := r.ProjectRepo.FindRolesByProjectAndUser(projectId, userId)
	if err != nil {
		return
	}

	err = r.DB.Model(&model.ProjectRole{}).Where("id = ?", projectMemberRole.ProjectRoleId).Scan(&projectRole).Error

	return
}

func (r *ProjectRoleRepo) AllRoleList() (projectRoles []model.ProjectRole, err error) {
	err = r.DB.Model(&model.ProjectRole{}).Scan(&projectRoles).Error

	return
}

func (r *ProjectRoleRepo) FindByIds(ids []uint) (projectRoles []model.ProjectRole, err error) {
	db := r.DB.Model(&model.ProjectRole{}).Where("id IN (?)", ids)

	err = db.Find(&projectRoles).Error
	return
}
