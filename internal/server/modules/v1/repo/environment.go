package repo

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"gorm.io/gorm"
)

type EnvironmentRepo struct {
	DB *gorm.DB `inject:""`
}

func (r *EnvironmentRepo) List(projectId int) (pos []model.Environment, err error) {
	err = r.DB.
		Select("id", "name").
		Where("project_id=?", projectId).
		Where("NOT deleted").
		Order("created_at DESC").
		Find(&pos).Error
	return
}

func (r *EnvironmentRepo) Get(id uint) (env model.Environment, err error) {
	err = r.DB.
		Where("id=?", id).
		Where("NOT deleted").
		First(&env).Error
	return
}

func (r *EnvironmentRepo) GetByInterface(interfaceId uint) (env model.Environment, err error) {
	interf := model.Interface{}
	err = r.DB.
		Where("id=?", interfaceId).
		Where("NOT deleted").
		First(&interf).Error

	if err != nil {
		return
	}

	env, err = r.Get(interf.EnvironmentId)

	return
}

func (r *EnvironmentRepo) GetVars(envId uint) (vars []model.EnvironmentVar, err error) {
	err = r.DB.
		Where("environment_id=?", envId).
		Where("NOT deleted").
		Order("created_at DESC").
		Find(&vars).Error

	return
}

func (r *EnvironmentRepo) Save(env *model.Environment) (err error) {
	err = r.DB.Save(env).Error
	return
}

func (r *EnvironmentRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&model.Environment{}).
		Where("id=?", id).
		Update("deleted", true).
		Error

	return
}
