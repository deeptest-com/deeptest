package service

import (
	"encoding/json"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	repo "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type ExecConditionService struct {
	PreConditionRepo  *repo.PreConditionRepo  `inject:""`
	PostConditionRepo *repo.PostConditionRepo `inject:""`

	ExtractorRepo  *repo.ExtractorRepo  `inject:""`
	CheckpointRepo *repo.CheckpointRepo `inject:""`
	ScriptRepo     *repo.ScriptRepo     `inject:""`

	ShareVarService *ShareVarService `inject:""`
}

func (s *ExecConditionService) SavePreConditionResult(invokeId uint, preConditions []domain.InterfaceExecCondition,
	usedBy consts.UsedBy) (err error) {

	for _, condition := range preConditions {
		if condition.Type == consts.ConditionTypeScript {
			var scriptBase domain.ScriptBase
			json.Unmarshal(condition.Raw, &scriptBase)
			scriptBase.InvokeId = invokeId

			s.ScriptRepo.UpdateResult(scriptBase)
			s.ScriptRepo.CreateLog(scriptBase)
		}
	}

	return
}

func (s *ExecConditionService) SavePostConditionResult(invokeId,
	debugInterfaceId, caseInterfaceId, endpointInterfaceId, serveId, processorId, scenarioId uint, usedBy consts.UsedBy,
	postConditions []domain.InterfaceExecCondition) (err error) {

	for _, condition := range postConditions {
		if condition.Type == consts.ConditionTypeExtractor {
			var extractorBase domain.ExtractorBase
			json.Unmarshal(condition.Raw, &extractorBase)
			extractorBase.InvokeId = invokeId

			s.ExtractorRepo.UpdateResult(extractorBase)
			s.ExtractorRepo.CreateLog(extractorBase)

			// add all ids for easy to load
			s.ShareVarService.Save(extractorBase.Variable, extractorBase.Result,
				invokeId, debugInterfaceId, caseInterfaceId, endpointInterfaceId, serveId, processorId, scenarioId,
				extractorBase.Scope, usedBy)

		} else if condition.Type == consts.ConditionTypeCheckpoint {
			var checkpointBase domain.CheckpointBase
			json.Unmarshal(condition.Raw, &checkpointBase)
			checkpointBase.InvokeId = invokeId

			s.CheckpointRepo.UpdateResult(checkpointBase)
			s.CheckpointRepo.CreateLog(checkpointBase)

		} else if condition.Type == consts.ConditionTypeScript {
			var scriptBase domain.ScriptBase
			json.Unmarshal(condition.Raw, &scriptBase)
			scriptBase.InvokeId = invokeId

			s.ScriptRepo.UpdateResult(scriptBase)
			s.ScriptRepo.CreateLog(scriptBase)
		}
	}

	return
}
