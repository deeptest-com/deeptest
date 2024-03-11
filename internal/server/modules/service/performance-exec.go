package service

import (
	"encoding/json"
	"fmt"
	agentExec "github.com/aaronchen2k/deeptest/internal/agent/exec"
	ptdomain "github.com/aaronchen2k/deeptest/internal/performance/pkg/domain"
	ptProto "github.com/aaronchen2k/deeptest/internal/performance/proto"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type PerformanceExecService struct {
	EnvironmentRepo         *repo.EnvironmentRepo         `inject:""`
	PerformanceTestPlanRepo *repo.PerformanceTestPlanRepo `inject:""`
	ScenarioRepo            *repo.ScenarioRepo            `inject:""`

	ScenarioExecService *ScenarioExecService `inject:""`
	EnvironmentService  *EnvironmentService  `inject:""`
	SceneService        *SceneService        `inject:""`
}

func (s *PerformanceExecService) LoadExecData(planId, environmentId uint) (
	ret ptdomain.PerformanceTestData, err error) {

	plan, err := s.PerformanceTestPlanRepo.Get(planId)
	if err != nil {
		return
	}

	scenarioExecObj, err := s.ScenarioExecService.LoadExecData(plan.ScenarioId, environmentId)
	if err != nil {
		return
	}

	ret.Goal = s.getGoalFromScenarioExecObj(scenarioExecObj)
	ret.Runners = s.getRunnersFromScenarioExecObj(scenarioExecObj)
	ret.Scenarios = s.getScenariosFromScenarioExecObj(scenarioExecObj)

	ret.ExecScene = scenarioExecObj.ExecScene

	return
}

func (s *PerformanceExecService) getGoalFromScenarioExecObj(execObj agentExec.ScenarioExecObj) (ret ptdomain.Goal) {
	for _, processor := range execObj.RootProcessor.Children {
		if processor.EntityType == consts.ProcessorPerformanceGoalDefault {
			goalEntity := agentExec.ProcessorPerformanceGoal{}
			json.Unmarshal(processor.EntityRaw, &goalEntity)

			ret = ptdomain.Goal{
				Type: goalEntity.Type,

				Duration: goalEntity.Duration,
				Loop:     goalEntity.Loop,

				ResponseTime: goalEntity.ResponseTime,
				Qps:          goalEntity.Qps,
				FailRate:     goalEntity.FailRate,
			}

			return
		}
	}
	return
}

func (s *PerformanceExecService) getRunnersFromScenarioExecObj(execObj agentExec.ScenarioExecObj) (ret []*ptdomain.Runner) {
	for _, processor := range execObj.RootProcessor.Children {
		if processor.EntityType == consts.ProcessorPerformanceRunnersDefault {
			for _, runnerProcessor := range processor.Children {
				runnerEntity := agentExec.ProcessorPerformanceRunner{}
				json.Unmarshal(runnerProcessor.EntityRaw, &runnerEntity)

				runner := ptdomain.Runner{
					Id:          int32(runnerProcessor.EntityId),
					Name:        processor.Name,
					WebAddress:  fmt.Sprintf("%s:%d", runnerEntity.Ip, runnerEntity.WebPort),
					GrpcAddress: fmt.Sprintf("%s:%d", runnerEntity.Ip, runnerEntity.GrpcPort),
					Weight:      int32(runnerEntity.Weight),
					Scenarios:   s.getScenarioIdsForRunner(runnerProcessor.EntityId, execObj),
				}

				ret = append(ret, &runner)
			}

			return
		}
	}

	return
}

func (s *PerformanceExecService) getScenarioIdsForRunner(runnerProcessorId uint, execObj agentExec.ScenarioExecObj) (
	ret []int32) {

	for _, processor := range execObj.RootProcessor.Children {
		if processor.EntityType == consts.ProcessorPerformanceScenariosDefault {
			for _, scenarioProcessor := range processor.Children {
				scenarioEntity := agentExec.ProcessorPerformanceScenario{}
				json.Unmarshal(scenarioProcessor.EntityRaw, &scenarioEntity)

				for _, runnerId := range scenarioEntity.RunnerIds {
					if uint(runnerId) == runnerProcessorId {
						ret = append(ret, int32(scenarioProcessor.EntityId))

						break
					}
				}
			}

			return
		}
	}

	return
}

func (s *PerformanceExecService) getScenariosFromScenarioExecObj(execObj agentExec.ScenarioExecObj) (
	ret []*ptProto.Scenario) {

	for _, processor := range execObj.RootProcessor.Children {
		if processor.EntityType == consts.ProcessorPerformanceScenariosDefault {
			for _, scenarioProcessor := range processor.Children {
				scenarioEntity := agentExec.ProcessorPerformanceScenario{}
				json.Unmarshal(scenarioProcessor.EntityRaw, &scenarioEntity)

				scenario := ptProto.Scenario{
					Id:   int32(scenarioProcessor.EntityId),
					Name: processor.Name,

					GenerateType: scenarioEntity.GenerateType.String(),
					Target:       int32(scenarioEntity.Target),
				}

				rootProcessor := agentExec.Processor{
					ProcessorBase: agentExec.ProcessorBase{
						Name:           "root",
						EntityCategory: consts.ProcessorRoot,
						EntityType:     consts.ProcessorRootDefault,
						Children:       scenarioProcessor.Children,
					},
				}
				scenario.ProcessorRaw, _ = json.Marshal(rootProcessor)

				ret = append(ret, &scenario)
			}
		}
	}

	return
}
