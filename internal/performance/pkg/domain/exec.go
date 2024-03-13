package ptdomain

import (
	ptconsts "github.com/aaronchen2k/deeptest/internal/performance/pkg/consts"
	ptproto "github.com/aaronchen2k/deeptest/internal/performance/proto"
)

type ExecParamsInCtx struct {
	Room                 string
	VuNo                 int
	RunnerId             int32
	RunnerName           string
	ConductorGrpcAddress string
	//InfluxdbSender       metrics.MessageSender

	Weight int

	// for constant generate type
	Target   int `json:"target"`
	Duration int `json:"duration"`

	// for ramp generate type
	Stages []*ptproto.Stage // ramp generator

	Scenario            *ptproto.Scenario
	RunnerExecScenarios map[uint]map[uint]bool

	Loop int

	Mode ptconsts.ExecMode
}
