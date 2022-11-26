package repo

import (
	"errors"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"gorm.io/gorm"
)

type ProjectRoleRepo struct {
	DB *gorm.DB `inject:""`
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
