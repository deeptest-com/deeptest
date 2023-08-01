package service

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
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

func (s *ScriptService) List(debugInterfaceId, endpointInterfaceId uint) (scripts []model.DebugConditionScript, err error) {
	scripts, err = s.ScriptRepo.List(debugInterfaceId, endpointInterfaceId)

	return
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

func (s *ScriptService) Delete(reqId uint) (err error) {
	err = s.ScriptRepo.Delete(reqId)

	return
}

func (s *ScriptService) Check(script model.DebugConditionScript, caseInterfaceId, scenarioProcessorId uint, resp domain.DebugResponse,
	usedBy consts.UsedBy) (logScript model.ExecLogScript, err error) {

	if script.Disabled {
		script.ResultStatus = ""

		s.ScriptRepo.UpdateResult(script, usedBy)

		return
	}

	script.ResultStatus = consts.Pass

	return
}
