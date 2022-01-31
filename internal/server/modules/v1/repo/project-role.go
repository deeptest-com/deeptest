package repo

import (
	"errors"
	logUtils "github.com/aaronchen2k/deeptest/internal/pkg/lib/log"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"gorm.io/gorm"
)

type ProjectRoleRepo struct {
	DB *gorm.DB `inject:""`
}

func NewProjectRoleRepo() *ProjectRoleRepo {
	return &ProjectRoleRepo{}
}

// FindByName
func (r *ProjectRoleRepo) FindByName(name string) (projectRole model.ProjectRole, err error) {
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
