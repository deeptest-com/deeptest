package ptdomain

import (
	ptconsts "github.com/aaronchen2k/deeptest/internal/agent/performance/pkg/consts"
)

type WsReq struct {
	Act ptconsts.ExecType `json:"act"`

	PlanExecReq PerformanceTestReq `json:"planExecReq"`
}
