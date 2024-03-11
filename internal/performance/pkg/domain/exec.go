package ptdomain

import (
	"github.com/aaronchen2k/deeptest/internal/performance/pkg/consts"
	ptproto "github.com/aaronchen2k/deeptest/internal/performance/proto"
)

type ExecParamsInCtx struct {
	RunnerId   int32
	RunnerName string
	Room       string

	Target int
	Weight int

	Stages []*ptproto.Stage // ramp generator

	Duration int
	Loop     int

	Mode ptconsts.ExecMode

	Scenario            *ptproto.Scenario
	RunnerExecScenarios map[uint]map[uint]bool

	ServerAddress string
}
