package action

import (
	agentExec "github.com/aaronchen2k/deeptest/internal/agent/exec"
	"github.com/aaronchen2k/deeptest/internal/agent/service"
)

func Run(scenario, plan, env int, server, token string) {
	//s.ExecScenarioService.ExecScenario(scenarioId, nil)

	if scenario > 0 {
		req := agentExec.ScenarioExecReq{
			ServerUrl:     server,
			Token:         token,
			ScenarioId:    scenario,
			EnvironmentId: env,
		}

		service.RunScenario(&req, nil)

	} else if plan > 0 {
		req := agentExec.PlanExecReq{
			ServerUrl:     server,
			Token:         token,
			PlanId:        plan,
			EnvironmentId: env,
		}

		service.RunPlan(&req, nil)

	}
}
