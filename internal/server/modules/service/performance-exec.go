package service

import (
	"encoding/json"
	agentExec "github.com/aaronchen2k/deeptest/internal/agent/exec"
	ptdomain "github.com/aaronchen2k/deeptest/internal/agent/performance/pkg/domain"
	ptProto "github.com/aaronchen2k/deeptest/internal/agent/performance/proto"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type PerformanceExecService struct {
	EnvironmentRepo         *repo.EnvironmentRepo         `inject:""`
	PerformanceTestPlanRepo *repo.PerformanceTestPlanRepo `inject:""`
	ScenarioRepo            *repo.ScenarioRepo            `inject:""`
	ProjectSettingsRepo     *repo.ProjectSettingsRepo     `inject:""`
	PerformanceRunnerRepo   *repo.PerformanceRunnerRepo   `inject:""`

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

	s.UpdateServerInfo(plan.ProjectId, &ret)

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
	runners, _ := s.PerformanceRunnerRepo.List(int(execObj.ScenarioId))

	for _, runner := range runners {
		runner := ptdomain.Runner{
			Id:         int32(runner.ID),
			Name:       runner.Name,
			WebAddress: runner.WebAddress,
			Weight:     int32(runner.Weight),
			Scenarios:  s.getScenarioIdsForRunner(runner.ID, execObj),
		}

		ret = append(ret, &runner)
	}

	return
}

func (s *PerformanceExecService) getScenarioIdsForRunner(runnerId uint, execObj agentExec.ScenarioExecObj) (
	ret []int32) {

	for _, processor := range execObj.RootProcessor.Children {
		if processor.EntityType == consts.ProcessorPerformanceScenariosDefault {
			for _, scenarioProcessor := range processor.Children {
				scenarioEntity := agentExec.ProcessorPerformanceScenario{}
				json.Unmarshal(scenarioProcessor.EntityRaw, &scenarioEntity)

				for _, id := range scenarioEntity.RunnerIds {
					if uint(id) == runnerId {
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
						ID:             scenarioProcessor.ID,
						Name:           "root",
						EntityCategory: consts.ProcessorRoot,
						EntityType:     consts.ProcessorRootDefault,
						Children:       scenarioProcessor.Children,

						ParentId:  0,
						ProjectId: processor.ProjectId,
					},
				}
				scenario.ProcessorRaw, _ = json.Marshal(rootProcessor)

				ret = append(ret, &scenario)
			}
		}
	}

	return
}

func (s *PerformanceExecService) UpdateServerInfo(projectId uint, data *ptdomain.PerformanceTestData) {
	po, _ := s.ProjectSettingsRepo.GetPerformance(projectId)

	data.ConductorGrpcAddress = po.ConductorGrpcAddress

	data.InfluxdbAddress = po.InfluxdbAddress
	data.InfluxdbOrg = po.InfluxdbOrg
	data.InfluxdbToken = po.InfluxdbToken
}
