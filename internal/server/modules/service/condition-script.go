package service

import (
	model "github.com/aaronchen2k/deeptest/internal/server/modules/model"
	repo "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type ScriptService struct {
	ScriptRepo      *repo.ScriptRepo      `inject:""`
	EnvironmentRepo *repo.EnvironmentRepo `inject:""`
	ProjectRepo     *repo.ProjectRepo     `inject:""`
	ExtractorRepo   *repo.ExtractorRepo   `inject:""`
	VariableService *VariableService      `inject:""`
}

func (s *ScriptService) Get(id uint) (script model.DebugConditionScript, err error) {
	script, err = s.ScriptRepo.Get(id)

	return
}

func (s *ScriptService) Create(script *model.DebugConditionScript) (err error) {
	err = s.ScriptRepo.Save(script)

	return
}

func (s *ScriptService) Update(script *model.DebugConditionScript) (err error) {
	err = s.ScriptRepo.Save(script)

	return
}
