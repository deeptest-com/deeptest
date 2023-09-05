package repo

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type EnvironmentRepo struct {
	*BaseRepo   `inject:""`
	DB          *gorm.DB     `inject:""`
	ProjectRepo *ProjectRepo `inject:""`
	ServeRepo   *ServeRepo   `inject:""`
}

func (r *EnvironmentRepo) List(projectId int) (pos []model.Environment, err error) {
	err = r.DB.
		Select("id", "name").
		Where("NOT deleted and project_id=?", projectId).
		Order("sort ASC").
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

	err = r.DB.
		Where("project_id=?", projectId).
		Where("NOT deleted").
		First(&env).Error

	return
}

// GetDefaultByProject 默认/Mock
func (r *EnvironmentRepo) GetDefaultByProject(projectId uint) (envs []model.Environment, err error) {
	err = r.DB.
		Where("project_id=?", projectId).
		Where("name IN (?)", []string{"默认环境", "Mock环境"}).
		Where("NOT deleted").
		Find(&envs).Error

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
		Where("name=? AND right_value=?", vari.Name, vari.RightValue).
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
		ProjectId: projectId,
		Name:      "默认环境",
	}
	if err = r.Save(&env); err != nil {
		return
	}

	mockEnv := model.Environment{
		ProjectId: projectId,
		Name:      "Mock环境",
	}
	err = r.Save(&mockEnv)
	//err = r.ProjectRepo.UpdateDefaultEnvironment(projectId, env.ID)

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
	err = r.DB.Model(&model.DebugConditionExtractor{}).
		Where("id=?", id).
		Update("disable_share", true).
		Error

	return
}

func (r *EnvironmentRepo) DisableAllShareVar(projectId uint) (err error) {
	err = r.DB.Model(&model.DebugConditionExtractor{}).
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

//func (r *EnvironmentRepo) ListVariableByProject(projectId uint) (vars []modelRef.EnvironmentVar, err error) {
//	environment, _ := r.GetByProject(projectId)
//	vars, _ = r.GetVars(environment.ID)
//
//	return
//}

func (r *EnvironmentRepo) SaveEnvironment(environment *model.Environment) (err error) {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		err = r.BaseRepo.Save(environment.ID, environment)
		if err != nil {
			return err
		}
		err = r.ServeRepo.SaveServer(environment.ID, environment.Name, environment.ServeServers)
		if err != nil {
			return err
		}
		err = r.SaveVars(environment.ProjectId, environment.ID, environment.Vars)
		if err != nil {
			return err
		}
		return nil
	})
}

func (r *EnvironmentRepo) DeleteEnvironment(id uint) (err error) {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		err = r.DB.Delete(&model.Environment{}, id).Error
		if err != nil {
			return err
		}
		err = r.DB.Delete(&model.ServeServer{}, "environment_id=?", id).Error
		if err != nil {
			return err
		}
		err = r.DB.Delete(&model.EnvironmentVar{}, "environment_id=?", id).Error
		if err != nil {
			return err
		}
		return nil
	})
}

func (r *EnvironmentRepo) SaveVars(projectId, environmentId uint, environmentVars []model.EnvironmentVar) (err error) {
	if len(environmentVars) == 0 {
		return
	}

	err = r.DB.Delete(&model.EnvironmentVar{}, "environment_id=? and project_id=?", environmentId, projectId).Error
	if err != nil {
		return err
	}

	for key, _ := range environmentVars {
		environmentVars[key].ID = 0
		environmentVars[key].EnvironmentId = environmentId
		environmentVars[key].ProjectId = projectId
	}
	err = r.DB.Create(environmentVars).Error

	if err != nil {
		return err
	}

	return
}

func (r *EnvironmentRepo) GetListByProjectId(projectId uint) (environments []model.Environment, err error) {
	err = r.DB.Order("sort").Find(&environments, "project_id=?", projectId).Error
	if err != nil {
		return
	}
	for key, _ := range environments {
		err = r.GetEnvironmentDetail(&environments[key])
		if err != nil {
			return
		}
	}
	return
}

func (r *EnvironmentRepo) GetEnvironmentById(id uint) (env *model.Environment, err error) {
	err = r.DB.First(&env, id).Error
	return
}

func (r *EnvironmentRepo) GetEnvironmentDetail(env *model.Environment) (err error) {
	var vars []model.EnvironmentVar
	err = r.DB.Find(&vars, "environment_id=?", env.ID).Error
	if err != nil {
		return
	}
	env.Vars = vars

	var servers []model.ServeServer
	err = r.DB.Find(&servers, "environment_id=?", env.ID).Error
	if err != nil {
		return
	}

	for key, server := range servers {
		var serve model.Serve
		r.DB.First(&serve, "id=?", server.ServeId)
		servers[key].ServeName = serve.Name
	}

	env.ServeServers = servers

	return
}

func (r *EnvironmentRepo) ListGlobalVar(projectId uint) (vars []model.EnvironmentVar, err error) {
	err = r.DB.Find(&vars, "project_id=? and environment_id=0", projectId).Error
	return
}

func (r *EnvironmentRepo) SaveParams(projectId uint, params []model.EnvironmentParam) (err error) {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		err = r.DB.Delete(&model.EnvironmentParam{}, " project_id=?", projectId).Error
		if err != nil {
			return err
		}
		err = r.DB.Create(params).Error
		if err != nil {
			return err
		}
		return nil
	})
}

func (r *EnvironmentRepo) ListParamModel(projectId uint) (ret []model.EnvironmentParam, err error) {
	err = r.DB.Find(&ret, "project_id=?", projectId).Error
	if err != nil {
		return
	}

	return
}

func (r *EnvironmentRepo) ListParams(projectId uint) (res map[string]interface{}, err error) {
	res = map[string]interface{}{}

	var params []model.EnvironmentParam
	err = r.DB.Find(&params, "project_id=?", projectId).Error
	if err != nil {
		return
	}

	for _, param := range params {
		in := string(param.In)

		res["projectId"] = param.ProjectId
		if res[in] == nil {
			res[in] = []model.EnvironmentParam{param}
		} else {
			res[in] = append(res[in].([]model.EnvironmentParam), param)
		}

	}
	return
}

func (r *EnvironmentRepo) SaveOrder(ids []uint) (err error) {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		for key, id := range ids {
			err = r.DB.Model(&model.Environment{}).Where("id=?", id).Update("sort", key).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *EnvironmentRepo) GetByIds(ids []uint) (envs map[uint]model.Environment, err error) {
	var res []model.Environment
	err = r.DB.Where("NOT disabled and NOT deleted and id in ?", ids).Find(&res).Error

	envs = make(map[uint]model.Environment)
	for _, item := range res {
		envs[item.ID] = item
	}
	return
}
