package agentExecDomain

import (
	ptconsts "github.com/aaronchen2k/deeptest/internal/agent/performance/pkg/consts"
	ptproto "github.com/aaronchen2k/deeptest/internal/agent/performance/proto"
	"github.com/aaronchen2k/deeptest/internal/agent/performance/runner/metrics"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/kataras/iris/v12"
)

type ExecParamsInCtx struct {
	Room string
	//VuNo                 int
	RunnerId             int32
	RunnerName           string
	WebServerUrl         string
	WebServerToken       string
	ConductorGrpcAddress string
	//InfluxdbSender       metrics.MessageSender

	Weight        int
	EnvironmentId int
	Sender        metrics.MessageSender

	// for constant generate type
	Target   int `json:"target"`
	Duration int `json:"duration"`

	// for ramp generate type
	Stages []*ptproto.Stage // ramp generator

	Scenario            *ptproto.Scenario
	RunnerExecScenarios map[uint]map[uint]bool

	Loop int // used in stage

	Mode           ptconsts.ExecMode
	LocalVarsCache iris.Map
	ExecScene      domain.ExecScene
}
