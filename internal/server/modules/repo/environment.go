package repo

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type EnvironmentRepo struct {
	DB          *gorm.DB     `inject:""`
	ProjectRepo *ProjectRepo `inject:""`
}

func (r *EnvironmentRepo) List() (pos []model.Environment, err error) {
	err = r.DB.
		Select("id", "name").
		Where("NOT deleted").
		Order("created_at ASC").
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

func (r *EnvironmentRepo) GetByName(name string) (env model.Environment, err error) {
	var envs []model.Environment

	db := r.DB.Model(&env).
		Where("name =? AND not deleted", name)

	err = db.Find(&envs).Error

	if err != nil {
		return
	}

	if len(envs) > 0 {
		env = envs[0]
	}

	return
}

func (r *EnvironmentRepo) GetByProject(projectId uint) (env model.Environment, err error) {
	project := model.Project{}
	err = r.DB.
		Where("id=?", projectId).
		Where("NOT deleted").
		First(&project).Error

	if err != nil {
		return
	}

	env, err = r.Get(project.EnvironmentId)

	return
}

func (r *EnvironmentRepo) GetVars(envId uint) (vars []model.EnvironmentVar, err error) {
	err = r.DB.
		Where("environment_id=?", envId).
		Where("NOT deleted").
		Order("created_at ASC").
		Find(&vars).Error

	return
}

func (r *EnvironmentRepo) GetSameVar(vari model.EnvironmentVar, envId uint) (ret model.EnvironmentVar, err error) {
	err = r.DB.
		Where("name=? AND value=?", vari.Name, vari.RightValue).
		Where("environment_id=?", envId).
		Where("NOT deleted").
		First(&ret).Error

	return
}

func (r *EnvironmentRepo) Save(env *model.Environment) (err error) {
	err = r.DB.Save(env).Error
	return
}

func (r *EnvironmentRepo) Copy(id int) (err error) {
	src, err := r.Get(uint(id))
	if err != nil {
		return
	}

	dist := model.Environment{}
	dist.Name = r.getCopyName(src.Name)

	r.Save(&dist)

	vars, _ := r.GetVars(src.ID)
	for _, item := range vars {
		item.ID = 0
		item.EnvironmentId = dist.ID

		r.SaveVar(&item)
	}

	return
}

func (r *EnvironmentRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&model.Environment{}).
		Where("id=?", id).
		Update("deleted", true).
		Error

	err = r.DB.Model(&model.EnvironmentVar{}).
		Where("environment_id = ?", id).Update("deleted", true).Error

	return
}

func (r *EnvironmentRepo) AddDefaultForProject(projectId uint) (err error) {
	env := model.Environment{
		Name: "默认环境",
	}
	err = r.Save(&env)
	err = r.ProjectRepo.UpdateDefaultEnvironment(projectId, env.ID)

	return
}

func (r *EnvironmentRepo) GetVar(id uint) (po model.EnvironmentVar, err error) {
	err = r.DB.
		Where("id=?", id).
		Where("NOT deleted").
		First(&po).Error
	return
}

func (r *EnvironmentRepo) SaveVar(po *model.EnvironmentVar) (err error) {
	err = r.DB.Save(po).Error

	return
}

func (r *EnvironmentRepo) DeleteVar(id uint) (err error) {
	err = r.DB.Model(&model.EnvironmentVar{}).
		Where("id=?", id).
		Update("deleted", true).
		Error

	return
}

func (r *EnvironmentRepo) ClearAllVar(environmentId uint) (err error) {
	err = r.DB.Model(&model.EnvironmentVar{}).
		Where("environment_id=?", environmentId).
		Update("deleted", true).
		Error

	return
}

func (r *EnvironmentRepo) DisableShareVar(id uint) (err error) {
	err = r.DB.Model(&model.InterfaceExtractor{}).
		Where("id=?", id).
		Update("disable_share", true).
		Error

	return
}

func (r *EnvironmentRepo) DisableAllShareVar(projectId uint) (err error) {
	err = r.DB.Model(&model.InterfaceExtractor{}).
		Where("project_id=?", projectId).
		Update("disable_share", true).
		Error

	return
}

func (r *EnvironmentRepo) GetVarByName(name string, id, environmentId uint) (envVar model.EnvironmentVar, err error) {
	var envVars []model.EnvironmentVar

	db := r.DB.Model(&envVar).
		Where("name = ? AND environment_id =? AND not deleted", name, environmentId)
	if id > 0 {
		db.Where("id != ?", id)
	}

	err = db.Find(&envVars).Error

	if err != nil {
		return
	}

	if len(envVars) > 0 {
		envVar = envVars[0]
	}

	return
}

func (r *EnvironmentRepo) getCopyName(name string) (ret string) {
	idx := strings.LastIndex(name, " ")

	if idx <= 0 {
		ret = name + " 1"
		env, _ := r.GetByName(ret)
		if env.ID > 0 {
			ret = r.getCopyName(ret)
		}
		return
	}

	left := name[:idx]
	right := name[idx+1:]
	rightNum, err := strconv.Atoi(right)
	if err != nil { // not a valid num
		ret = name + " 1"
		env, _ := r.GetByName(ret)
		if env.ID > 0 {
			ret = r.getCopyName(ret)
		}
		return
	}

	nextNum := rightNum + 1
	ret = left + fmt.Sprintf(" %d", nextNum)

	env, _ := r.GetByName(ret)
	if env.ID > 0 {
		ret = r.getCopyName(ret)
	}

	return
}

func (r *EnvironmentRepo) ListVariableByProject(projectId uint) (vars []model.EnvironmentVar, err error) {
	environment, _ := r.GetByProject(projectId)
	vars, _ = r.GetVars(environment.ID)

	return
}
