package ptdomain

import (
	"github.com/aaronchen2k/deeptest/internal/performance/pkg/consts"
	ptProto "github.com/aaronchen2k/deeptest/internal/performance/proto"
)

type PerformanceTestReq struct {
	BaseExecReqOfRunner

	Runners []*ptProto.Runner `json:"runners"`

	GoalAvgResponseTime float64 `json:"goalAvgResponseTime"`
	GoalAvgQps          float64 `json:"goalAvgQps"`
	GoalFailed          string  `json:"goalFailed"`
}

type PerformanceLogReq struct {
	Room string `json:"room"`
}

type PerformanceTestReqOfRunner struct {
	BaseExecReqOfRunner

	Weight int32 `json:"weight"`
}

type BaseExecReqOfRunner struct {
	Room   string `json:"room"`
	PlanId int    `json:"planId"`
	Title  string `json:"title"`

	GenerateType ptconsts.GenerateType `json:"generateType"`
	Stages       []*ptProto.Stage      `json:"stages"`

	Mode      ptconsts.ExecMode   `json:"mode"`
	Scenarios []*ptProto.Scenario `json:"scenarios"`

	ServerAddress   string `json:"serverAddress,omitempty"`
	InfluxdbAddress string `json:"influxdbAddress,omitempty"`
	InfluxdbOrg     string `json:"influxdbOrg,omitempty"`
	InfluxdbToken   string `json:"influxdbToken,omitempty"`
}

type Scenario struct {
	Name string `json:"name"`

	GenerateType string           `protobuf:"bytes,2,opt,name=generateType,proto3" json:"generateType,omitempty"`
	Stages       []*ptProto.Stage `protobuf:"bytes,3,rep,name=stages,proto3" json:"stages,omitempty"`

	Uuid string `json:"uuid,omitempty"`
	Dur  int    `json:"dur,omitempty"`

	TargetQps      int32  `json:"targetQps"`
	TargetDuration int32  `json:"targetDuration"`
	TargetFailRate string `json:"targetFailRate"`

	VuNo       int                  `json:"vuNo,omitempty"`
	Processors []*ptProto.Processor `json:"processors,omitempty"`

	NsqServerAddress string `json:"nsqServerAddress,omitempty"`
	NsqLookupAddress string `json:"nsqLookupAddress,omitempty"`
}

type Metrics struct {
	Name      string `json:"name"`
	Value     string `gorm:"type:text" json:"value"`
	Timestamp string `json:"timestamp"`
}
