package service

import (
	"errors"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	repo2 "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/jinzhu/copier"
)

type EnvironmentService struct {
	EnvironmentRepo *repo2.EnvironmentRepo `inject:""`
	ScenarioRepo    *repo2.ScenarioRepo    `inject:""`
	InterfaceRepo   *repo2.InterfaceRepo   `inject:""`
	ProjectRepo     *repo2.ProjectRepo     `inject:""`
}

func (s *EnvironmentService) List() (envs []model.Environment, err error) {
	envs, err = s.EnvironmentRepo.List()

	return
}

func (s *EnvironmentService) ListVariableForExec(scenario model.Scenario) (ret map[string]interface{}, err error) {
	ret = map[string]interface{}{}

	pos, err := s.EnvironmentRepo.ListVariableByProject(scenario.ProjectId)
	if err != nil {
		return
	}

	for _, po := range pos {
		ret[po.Name] = po.RightValue
	}

	return
}

func (s *EnvironmentService) Get(id, projectId uint) (env model.Environment, err error) {
	if id > 0 {
		env, err = s.EnvironmentRepo.Get(id)
	} else {
		env, _ = s.EnvironmentRepo.GetByProject(projectId)
	}

	if env.ID > 0 {
		env.Vars, err = s.EnvironmentRepo.GetVars(env.ID)
	}

	return
}

func (s *EnvironmentService) Copy(envId int) (err error) {
	err = s.EnvironmentRepo.Copy(envId)

	return
}

func (s *EnvironmentService) Create(env *model.Environment, projectId uint) (err error) {
	err = s.EnvironmentRepo.Save(env)
	err = s.ProjectRepo.UpdateDefaultEnvironment(projectId, env.ID)

	return
}

func (s *EnvironmentService) Update(env *model.Environment) (err error) {
	err = s.EnvironmentRepo.Save(env)

	return
}

func (s *EnvironmentService) Delete(reqId uint) (err error) {
	err = s.EnvironmentRepo.Delete(reqId)

	return
}

func (s *EnvironmentService) Change(id, projectId int) (err error) {
	err = s.ProjectRepo.UpdateDefaultEnvironment(uint(projectId), uint(id))

	return
}

func (s *EnvironmentService) GetVar(id uint) (env model.EnvironmentVar, err error) {
	env, err = s.EnvironmentRepo.GetVar(id)

	return
}

func (s *EnvironmentService) CreateVar(po *model.EnvironmentVar) (err error) {
	temp, _ := s.EnvironmentRepo.GetVarByName(po.Name, 0, po.EnvironmentId)

	if temp.ID > 0 {
		err = errors.New("")
		return
	}

	err = s.EnvironmentRepo.SaveVar(po)

	return
}

func (s *EnvironmentService) UpdateVar(po *model.EnvironmentVar) (err error) {
	temp, _ := s.EnvironmentRepo.GetVarByName(po.Name, po.ID, po.EnvironmentId)
	if temp.ID > 0 {
		err = errors.New("")
		return
	}

	err = s.EnvironmentRepo.SaveVar(po)

	return
}

func (s *EnvironmentService) DeleteVar(id uint) (err error) {
	err = s.EnvironmentRepo.DeleteVar(id)

	return
}

func (s *EnvironmentService) ClearAllVar(environmentId uint) (err error) {
	err = s.EnvironmentRepo.ClearAllVar(environmentId)

	return
}

func (s *EnvironmentService) DisableShareVar(id uint) (err error) {
	err = s.EnvironmentRepo.DisableShareVar(id)

	return
}

func (s *EnvironmentService) DisableAllShareVar(interfaceId uint) (err error) {
	interf, _ := s.InterfaceRepo.Get(interfaceId)

	err = s.EnvironmentRepo.DisableAllShareVar(interf.ProjectId)

	return
}

func (s *EnvironmentService) Save(req v1.EnvironmentReq) (err error) {
	var environment model.Environment
	copier.CopyWithOption(&environment, req, copier.Option{DeepCopy: true})
	err = s.EnvironmentRepo.SaveEnvironment(environment)
	return
}

func (s *EnvironmentService) ListAll(projectId uint) (res []model.Environment, err error) {
	res, err = s.EnvironmentRepo.GetListByProjectId(projectId)
	return
}

func (s *EnvironmentService) SaveGlobal(projectId uint, req []v1.EnvironmentVariable) (err error) {
	var vars []model.EnvironmentVar
	copier.CopyWithOption(&vars, req, copier.Option{DeepCopy: true})
	err = s.EnvironmentRepo.SaveVars(projectId, 0, vars)
	return
}

func (s *EnvironmentService) ListGlobal(projectId uint) (res []model.EnvironmentVar, err error) {
	res, err = s.EnvironmentRepo.ListGlobal(projectId)
	return
}
