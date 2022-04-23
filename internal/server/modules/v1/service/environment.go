package service

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
)

type EnvironmentService struct {
	EnvironmentRepo *repo.EnvironmentRepo `inject:""`
	InterfaceRepo   *repo.InterfaceRepo   `inject:""`
	ProjectRepo     *repo.ProjectRepo     `inject:""`
}

func (s *EnvironmentService) List(projectId int) (envs []model.Environment, err error) {
	envs, err = s.EnvironmentRepo.List(projectId)

	return
}

func (s *EnvironmentService) Get(id, interfaceId int) (env model.Environment, err error) {
	if id > 0 {
		env, err = s.EnvironmentRepo.Get(uint(id))
	} else {
		env, _ = s.EnvironmentRepo.GetByInterface(uint(interfaceId))
	}

	if env.ID > 0 {
		env.Vars, err = s.EnvironmentRepo.GetVars(env.ID)
	}

	return
}

func (s *EnvironmentService) Create(env *model.Environment) (err error) {
	err = s.EnvironmentRepo.Save(env)

	err = s.ProjectRepo.UpdateDefaultEnvironment(env.ProjectId, env.ID)

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

func (s *EnvironmentService) Change(id, interfaceId, projectId int) (err error) {
	err = s.InterfaceRepo.UpdateDefaultEnvironment(uint(interfaceId), uint(id))
	err = s.ProjectRepo.UpdateDefaultEnvironment(uint(projectId), uint(id))

	return
}
