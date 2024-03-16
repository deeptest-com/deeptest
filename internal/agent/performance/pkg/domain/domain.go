package ptdomain

import (
	"github.com/aaronchen2k/deeptest/internal/agent/performance/pkg/consts"
	ptproto "github.com/aaronchen2k/deeptest/internal/agent/performance/proto"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/kataras/iris/v12"
)

type PerformanceTestReq struct {
	UserId    uint   `json:"userId"`
	ProjectId uint   `json:"projectId"`
	ServerUrl string `json:"serverUrl"`
	Token     string `json:"token"`

	Room          string `json:"room"`
	PlanId        uint   `json:"planId"`
	EnvironmentId int    `json:"environmentId"`
}

type PerformanceTestData struct {
	BaseExecReqOfRunner

	Goal    Goal      `json:"goal"`
	Runners []*Runner `json:"runners"`

	ServerUrl string `json:"serverUrl"`
	Token     string `json:"token"`

	ExecScene      domain.ExecScene `json:"execScene"`
	LocalVarsCache iris.Map         `json:"localVarsCache"`
}

type Goal struct {
	Type ptconsts.GoalType `json:"type,omitempty"`

	Duration int `json:"duration,omitempty"`
	Loop     int `json:"loop,omitempty"`

	ResponseTime float32 `json:"responseTime,omitempty"`
	Qps          float32 `json:"qps,omitempty"`
	FailRate     float32 `json:"failRate,omitempty"`
}

type Runner struct {
	Id          int32   `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	GrpcAddress string  `json:"grpcAddress,omitempty"`
	WebAddress  string  `json:"webAddress,omitempty"`
	Weight      int32   `json:"weight,omitempty"`
	Scenarios   []int32 `json:"scenarios,omitempty"`
}

type PerformanceExecResult struct {
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
	Stages       []*ptproto.Stage      `json:"stages"`

	Mode      ptconsts.ExecMode   `json:"mode"`
	Scenarios []*ptproto.Scenario `json:"scenarios"`

	ConductorGrpcAddress string `json:"conductorGrpcAddress,omitempty"`
	InfluxdbAddress      string `json:"influxdbAddress,omitempty"`
	InfluxdbOrg          string `json:"influxdbOrg,omitempty"`
	InfluxdbToken        string `json:"influxdbToken,omitempty"`
}

type Scenario struct {
	Name string `json:"name"`

	GenerateType string           `protobuf:"bytes,2,opt,name=generateType,proto3" json:"generateType,omitempty"`
	Stages       []*ptproto.Stage `protobuf:"bytes,3,rep,name=stages,proto3" json:"stages,omitempty"`

	Uuid string `json:"uuid,omitempty"`
	Dur  int    `json:"dur,omitempty"`

	VuNo int `json:"vuNo,omitempty"`

	NsqServerAddress string `json:"nsqServerAddress,omitempty"`
	NsqLookupAddress string `json:"nsqLookupAddress,omitempty"`
}

type Metrics struct {
	Name      string `json:"name"`
	Value     string `gorm:"type:text" json:"value"`
	Timestamp string `json:"timestamp"`
}
