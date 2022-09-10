package action

import (
	commandConsts "github.com/aaronchen2k/deeptest/internal/command/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/service"
	"strconv"
)

type RunAction struct {
	ExecScenarioService *service.ExecScenarioService `inject:""`
}

func (s *RunAction) Run(scenarioIdOrName string) {
	scenarioId, _ := strconv.Atoi(scenarioIdOrName)

	if scenarioId == 0 {
		scenario := getScenarioByName(scenarioIdOrName)
		scenarioId = int(scenario.ID)
	}

	s.ExecScenarioService.ExecScenario(scenarioId, nil)
}

func getScenarioById(id int) (scenario model.Scenario) {
	commandConsts.DB.
		Where("id = ? && NOT deleted && NOT disabled", id).
		First(&scenario)

	return
}

func getScenarioByName(name string) (scenario model.Scenario) {
	commandConsts.DB.
		Where("name = ? && NOT deleted && NOT disabled", name).
		First(&scenario)

	return
}
