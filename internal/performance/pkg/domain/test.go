package ptdomain

import (
	ptconsts "github.com/aaronchen2k/deeptest/internal/performance/pkg/consts"
	ptproto "github.com/aaronchen2k/deeptest/internal/performance/proto"
	"time"
)

type TestItem struct {
	Room string            `json:"room"`
	Role ptconsts.TestRole `json:"role"`

	ConductorReq *PerformanceTestReq              `json:"conductorReq"`
	RunnerReq    *ptproto.PerformanceExecStartReq `json:"runnerReq"`

	CreateTime time.Time `json:"createTime"`
}
