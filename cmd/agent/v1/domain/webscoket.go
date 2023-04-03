package domain

import (
	agentExec "github.com/aaronchen2k/deeptest/internal/agent/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
)

type WsReq struct {
	Act             consts.ExecType           `json:"act"`
	ScenarioExecReq agentExec.ScenarioExecReq `json:"scenarioExecReq"`
	PlanExecReq     agentExec.PlanExecReq     `json:"planExecReq"`
}
