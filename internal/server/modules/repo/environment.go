package repo

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
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

func (r *EnvironmentRepo) List(tenantId consts.TenantId, projectId int) (pos []model.Environment, err error) {
	err = r.GetDB(tenantId).
		Select("id", "name").
		Where("NOT deleted and project_id=?", projectId).
		Order("sort ASC").
		Find(&pos).Error
	return
}

func (r *EnvironmentRepo) Get(tenantId consts.TenantId, id uint) (env model.Environment, err error) {
	err = r.GetDB(tenantId).
		Where("id=?", id).
		Where("NOT deleted").
		First(&env).Error
	return
}

func (r *EnvironmentRepo) GetByName(tenantId consts.TenantId, name string) (env model.Environment, err error) {
	var envs []model.Environment

	db := r.GetDB(tenantId).Model(&env).
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

func (r *EnvironmentRepo) GetByProject(tenantId consts.TenantId, projectId uint) (env model.Environment, err error) {

	err = r.GetDB(tenantId).
		Where("project_id=?", projectId).
		Where("NOT deleted").
		First(&env).Error

	return
}

func (r *EnvironmentRepo) GetByUserAndProject(tenantId consts.TenantId, userId, projectId uint) (env model.Environment, err error) {
	relaPo, err := r.GetProjectUserServer(tenantId, projectId, userId)
	if err != nil {
		return
	}

	env, err = r.Get(tenantId, relaPo.ServerId)

	return
}

func (r *EnvironmentRepo) SetProjectUserServer(tenantId consts.TenantId, projectId, userId, serverId uint) (err error) {
	data, err := r.GetProjectUserServer(tenantId, projectId, userId)
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}

	data.ProjectId = projectId
	data.UserId = userId
	data.ServerId = serverId
	err = r.GetDB(tenantId).Save(&data).Error

	return
}

func (r *EnvironmentRepo) GetProjectUserServer(tenantId consts.TenantId, projectId, userId uint) (res model.ProjectUserServer, err error) {
	err = r.GetDB(tenantId).
		Where("user_id = ? AND project_id=?", userId, projectId).
		Where("NOT deleted").
		First(&res).Error

	return
}

// GetDefaultByProject 默认/Mock
func (r *EnvironmentRepo) GetDefaultByProject(tenantId consts.TenantId, projectId uint) (envs []model.Environment, err error) {
	err = r.GetDB(tenantId).
		Where("project_id=?", projectId).
		Where("name IN (?)", []string{"默认环境", "Mock环境"}).
		Where("NOT deleted").
		Find(&envs).Error

	return
}

func (r *EnvironmentRepo) GetVars(tenantId consts.TenantId, envId uint) (vars []model.EnvironmentVar, err error) {
	err = r.GetDB(tenantId).
		Where("environment_id=?", envId).
		Where("NOT deleted").
		Order("created_at ASC").
		Find(&vars).Error

	return
}

func (r *EnvironmentRepo) GetSameVar(tenantId consts.TenantId, vari model.EnvironmentVar, envId uint) (ret model.EnvironmentVar, err error) {
	err = r.GetDB(tenantId).
		Where("name=? AND right_value=?", vari.Name, vari.RightValue).
		Where("environment_id=?", envId).
		Where("NOT deleted").
		First(&ret).Error

	return
}

func (r *EnvironmentRepo) Save(tenantId consts.TenantId, env *model.Environment) (err error) {
	err = r.GetDB(tenantId).Save(env).Error
	return
}

func (r *EnvironmentRepo) Copy(tenantId consts.TenantId, id int) (err error) {
	src, err := r.Get(tenantId, uint(id))
	if err != nil {
		return
	}

	dist := model.Environment{}
	dist.Name = r.getCopyName(tenantId, src.Name)

	r.Save(tenantId, &dist)

	vars, _ := r.GetVars(tenantId, src.ID)
	for _, item := range vars {
		item.ID = 0
		item.EnvironmentId = dist.ID

		r.SaveVar(tenantId, &item)
	}

	return
}

func (r *EnvironmentRepo) Delete(tenantId consts.TenantId, id uint) (err error) {
	err = r.GetDB(tenantId).Model(&model.Environment{}).
		Where("id=?", id).
		Update("deleted", true).
		Error

	err = r.GetDB(tenantId).Model(&model.EnvironmentVar{}).
		Where("environment_id = ?", id).Update("deleted", true).Error

	return
}

func (r *EnvironmentRepo) AddDefaultForProject(tenantId consts.TenantId, projectId uint) (err error) {
	env := model.Environment{
		ProjectId: projectId,
		Name:      "默认环境",
	}
	if err = r.Save(tenantId, &env); err != nil {
		return
	}

	mockEnv := model.Environment{
		ProjectId: projectId,
		Name:      "Mock环境",
		Sort:      1,
	}
	err = r.Save(tenantId, &mockEnv)
	//err = r.ProjectRepo.UpdateDefaultEnvironment(projectId, env.ID)

	return
}

func (r *EnvironmentRepo) GetVar(tenantId consts.TenantId, id uint) (po model.EnvironmentVar, err error) {
	err = r.GetDB(tenantId).
		Where("id=?", id).
		Where("NOT deleted").
		First(&po).Error
	return
}

func (r *EnvironmentRepo) SaveVar(tenantId consts.TenantId, po *model.EnvironmentVar) (err error) {
	err = r.GetDB(tenantId).Save(po).Error

	return
}

func (r *EnvironmentRepo) DeleteVar(tenantId consts.TenantId, id uint) (err error) {
	err = r.GetDB(tenantId).Model(&model.EnvironmentVar{}).
		Where("id=?", id).
		Update("deleted", true).
		Error

	return
}

func (r *EnvironmentRepo) ClearAllVar(tenantId consts.TenantId, environmentId uint) (err error) {
	err = r.GetDB(tenantId).Model(&model.EnvironmentVar{}).
		Where("environment_id=?", environmentId).
		Update("deleted", true).
		Error

	return
}

func (r *EnvironmentRepo) DisableShareVar(tenantId consts.TenantId, id uint) (err error) {
	err = r.GetDB(tenantId).Model(&model.DebugConditionExtractor{}).
		Where("id=?", id).
		Update("disable_share", true).
		Error

	return
}

func (r *EnvironmentRepo) DisableAllShareVar(tenantId consts.TenantId, projectId uint) (err error) {
	err = r.GetDB(tenantId).Model(&model.DebugConditionExtractor{}).
		Where("project_id=?", projectId).
		Update("disable_share", true).
		Error

	return
}

func (r *EnvironmentRepo) GetVarByName(tenantId consts.TenantId, name string, id, environmentId uint) (envVar model.EnvironmentVar, err error) {
	var envVars []model.EnvironmentVar

	db := r.GetDB(tenantId).Model(&envVar).
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

func (r *EnvironmentRepo) getCopyName(tenantId consts.TenantId, name string) (ret string) {
	idx := strings.LastIndex(name, " ")

	if idx <= 0 {
		ret = name + " 1"
		env, _ := r.GetByName(tenantId, ret)
		if env.ID > 0 {
			ret = r.getCopyName(tenantId, ret)
		}
		return
	}

	left := name[:idx]
	right := name[idx+1:]
	rightNum, err := strconv.Atoi(right)
	if err != nil { // not a valid num
		ret = name + " 1"
		env, _ := r.GetByName(tenantId, ret)
		if env.ID > 0 {
			ret = r.getCopyName(tenantId, ret)
		}
		return
	}

	nextNum := rightNum + 1
	ret = left + fmt.Sprintf(" %d", nextNum)

	env, _ := r.GetByName(tenantId, ret)
	if env.ID > 0 {
		ret = r.getCopyName(tenantId, ret)
	}

	return
}

//func (r *EnvironmentRepo) ListVariableByProject(projectId uint) (vars []modelRef.EnvironmentVar, err error) {
//	environment, _ := r.GetByProject(projectId)
//	vars, _ = r.GetVars(environment.ID)
//
//	return
//}

func (r *EnvironmentRepo) SaveEnvironment(tenantId consts.TenantId, environment *model.Environment) (err error) {
	return r.GetDB(tenantId).Transaction(func(tx *gorm.DB) error {
		err = r.BaseRepo.Save(tenantId, environment.ID, environment)
		if err != nil {
			return err
		}
		err = r.ServeRepo.SaveServer(tenantId, environment.ID, environment.Name, environment.ServeServers)
		if err != nil {
			return err
		}
		err = r.SaveVars(tenantId, environment.ProjectId, environment.ID, environment.Vars)
		if err != nil {
			return err
		}
		return nil
	})
}

func (r *EnvironmentRepo) DeleteEnvironment(tenantId consts.TenantId, id uint) (err error) {
	return r.GetDB(tenantId).Transaction(func(tx *gorm.DB) error {
		err = r.GetDB(tenantId).Delete(&model.Environment{}, id).Error
		if err != nil {
			return err
		}
		err = r.GetDB(tenantId).Delete(&model.ServeServer{}, "environment_id=?", id).Error
		if err != nil {
			return err
		}
		err = r.GetDB(tenantId).Delete(&model.EnvironmentVar{}, "environment_id=?", id).Error
		if err != nil {
			return err
		}
		return nil
	})
}

func (r *EnvironmentRepo) SaveVars(tenantId consts.TenantId, projectId, environmentId uint, environmentVars []model.EnvironmentVar) (err error) {

	err = r.GetDB(tenantId).Delete(&model.EnvironmentVar{}, "environment_id=? and project_id=?", environmentId, projectId).Error
	if err != nil {
		return err
	}

	if len(environmentVars) == 0 {
		return
	}

	for key, _ := range environmentVars {
		environmentVars[key].ID = 0
		environmentVars[key].EnvironmentId = environmentId
		environmentVars[key].ProjectId = projectId
	}
	err = r.GetDB(tenantId).Create(environmentVars).Error

	if err != nil {
		return err
	}

	return
}

func (r *EnvironmentRepo) GetListByProjectId(tenantId consts.TenantId, projectId uint) (environments []model.Environment, err error) {
	err = r.GetDB(tenantId).Order("sort").Find(&environments, "project_id=?", projectId).Error
	if err != nil {
		return
	}
	for key, _ := range environments {
		err = r.GetEnvironmentDetail(tenantId, &environments[key])
		if err != nil {
			return
		}
	}
	return
}

func (r *EnvironmentRepo) GetEnvironmentById(tenantId consts.TenantId, id uint) (env *model.Environment, err error) {
	err = r.GetDB(tenantId).First(&env, id).Error
	return
}

func (r *EnvironmentRepo) GetEnvironmentDetail(tenantId consts.TenantId, env *model.Environment) (err error) {
	var vars []model.EnvironmentVar
	err = r.GetDB(tenantId).Find(&vars, "environment_id=? AND NOT deleted", env.ID).Error
	if err != nil {
		return
	}
	env.Vars = vars

	var servers []model.ServeServer
	err = r.GetDB(tenantId).Model(&model.ServeServer{}).
		Joins("LEFT JOIN biz_project_serve serve ON biz_project_serve_server.serve_id=serve.id ").
		Select("biz_project_serve_server.*").
		Where("biz_project_serve_server.environment_id=? AND NOT serve.deleted", env.ID).
		Find(&servers).Error
	if err != nil {
		return
	}

	for key, server := range servers {
		var serve model.Serve
		r.GetDB(tenantId).First(&serve, "id=?", server.ServeId)
		servers[key].ServeName = serve.Name
	}

	env.ServeServers = servers

	return
}

func (r *EnvironmentRepo) ListGlobalVar(tenantId consts.TenantId, projectId uint) (vars []model.EnvironmentVar, err error) {
	err = r.GetDB(tenantId).Find(&vars, "project_id=? and environment_id=0", projectId).Error
	return
}

func (r *EnvironmentRepo) SaveParams(tenantId consts.TenantId, projectId uint, params []model.EnvironmentParam) (err error) {
	return r.GetDB(tenantId).Transaction(func(tx *gorm.DB) error {
		err = r.GetDB(tenantId).Delete(&model.EnvironmentParam{}, " project_id=?", projectId).Error
		if err != nil {
			return err
		}
		if len(params) > 0 {
			err = r.GetDB(tenantId).Create(params).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *EnvironmentRepo) ListParamModel(tenantId consts.TenantId, projectId uint) (ret []model.EnvironmentParam, err error) {
	err = r.GetDB(tenantId).Find(&ret, "project_id=?", projectId).Error
	if err != nil {
		return
	}

	return
}

func (r *EnvironmentRepo) ListParams(tenantId consts.TenantId, projectId uint) (res map[string]interface{}, err error) {
	res = map[string]interface{}{}

	var params []model.EnvironmentParam
	err = r.GetDB(tenantId).Find(&params, "project_id=?", projectId).Error
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

func (r *EnvironmentRepo) SaveOrder(tenantId consts.TenantId, ids []uint) (err error) {
	return r.GetDB(tenantId).Transaction(func(tx *gorm.DB) error {
		for key, id := range ids {
			err = r.GetDB(tenantId).Model(&model.Environment{}).Where("id=?", id).Update("sort", key).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *EnvironmentRepo) GetByIds(tenantId consts.TenantId, ids []uint) (envs map[uint]model.Environment, err error) {
	var res []model.Environment
	err = r.GetDB(tenantId).Where("NOT disabled and NOT deleted and id in ?", ids).Find(&res).Error

	envs = make(map[uint]model.Environment)
	for _, item := range res {
		envs[item.ID] = item
	}
	return
}

func (r *EnvironmentRepo) GetByProjectAndName(tenantId consts.TenantId, projectId uint, name string) (env model.Environment, err error) {

	err = r.GetDB(tenantId).
		Where("project_id=?", projectId).
		Where("name=?", name).
		Where("NOT deleted").
		First(&env).Error

	return
}

func (r *EnvironmentRepo) GetMaxOrder(tenantId consts.TenantId, projectId uint) (order uint) {
	environment := model.Environment{}

	err := r.GetDB(tenantId).Model(&model.Environment{}).
		Where("project_id = ?", projectId).
		Order("sort DESC").
		First(&environment).Error

	if err == nil {
		order = environment.Sort + 1
	}

	return
}
