package service

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	repo "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type ShareVarService struct {
	ShareVariableRepo *repo.ShareVariableRepo `inject:""`
}

func (s *ShareVarService) Save(name, value string, interfaceId, serveId, scenarioId uint,
	scope consts.ExtractorScope, usedBy consts.UsedBy) (err error) {

	po := model.ShareVariable{
		Name:        name,
		Value:       value,
		InterfaceId: interfaceId,
		ServeId:     serveId,
		ScenarioId:  scenarioId,
		Scope:       scope,
	}

	if usedBy == consts.InterfaceDebug {
		po.ID, err = s.ShareVariableRepo.GetExistByInterfaceDebug(name, serveId)
	} else if usedBy == consts.InterfaceDebug {
		po.ID, err = s.ShareVariableRepo.GetExistByScenarioDebug(name, scenarioId)
	}

	err = s.ShareVariableRepo.Save(&po)

	return
}

func (s *ShareVarService) ListForDebug(serveId, scenarioId uint, usedBy consts.UsedBy) (ret []domain.ShareVars, err error) {
	var pos []model.ShareVariable

	if usedBy == consts.InterfaceDebug {
		pos, err = s.ShareVariableRepo.ListByInterfaceDebug(serveId)
	} else if usedBy == consts.ScenarioDebug {
		pos, err = s.ShareVariableRepo.ListByScenarioDebug(scenarioId)
	}

	for _, po := range pos {
		ret = append(ret, domain.ShareVars{
			"id":    po.ID,
			"name":  po.Name,
			"value": po.Value,
		})
	}

	return
}
