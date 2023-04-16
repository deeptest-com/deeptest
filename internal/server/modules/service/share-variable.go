package service

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	repo "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type ShareVarService struct {
	ShareVariableRepo *repo.ShareVariableRepo `inject:""`
}

func (s *ShareVarService) Save(name, value string, interfaceId, serveId, scenarioId uint, scope consts.ExtractorScope) (err error) {
	po := model.ShareVariable{
		Name:        name,
		Value:       value,
		InterfaceId: interfaceId,
		ServeId:     serveId,
		ScenarioId:  scenarioId,
		Scope:       scope,
	}

	err = s.ShareVariableRepo.Save(&po)

	return
}

func (s *ShareVarService) ListByInterfaceDebug(serveId uint) (pos []model.ShareVariable, err error) {
	pos, err = s.ShareVariableRepo.ListByInterfaceDebug(serveId)
	return
}

func (s *ShareVarService) ListByScenarioDebug(scenarioId uint) (pos []model.ShareVariable, err error) {
	pos, err = s.ShareVariableRepo.ListByScenarioDebug(scenarioId)
	return
}
