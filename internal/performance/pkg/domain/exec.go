package ptdomain

import (
	"github.com/aaronchen2k/deeptest/internal/performance/pkg/consts"
	ptProto "github.com/aaronchen2k/deeptest/internal/performance/proto"
)

type ExecParamsInCtx struct {
	RunnerId   int32
	RunnerName string
	Room       string

	Target int
	Weight int

	Stages []*ptProto.Stage // ramp generator

	Duration int
	Loop     int

	Mode ptconsts.ExecMode

	Scenario            *ptProto.Scenario
	RunnerExecScenarios map[uint]map[uint]bool

	ServerAddress string
}
