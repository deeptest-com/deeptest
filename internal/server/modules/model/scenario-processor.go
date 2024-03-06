package model

import (
	agentExec "github.com/aaronchen2k/deeptest/internal/agent/exec"
	ptconsts "github.com/aaronchen2k/deeptest/internal/performance/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/kataras/iris/v12"
	"time"
)

type Processor struct {
	BaseModel
	CreatedBy uint `json:"createdBy"`

	Name     string            `json:"name" yaml:"name"`
	Comments string            `json:"comments" yaml:"comments"`
	Method   consts.HttpMethod `json:"method" yaml:"method"`

	ParentId   uint `json:"parentId"`
	ScenarioId uint `json:"scenarioId"`
	ProjectId  uint `json:"projectId"`

	EntityCategory consts.ProcessorCategory `json:"entityCategory"`
	EntityType     consts.ProcessorType     `json:"entityType"`

	EntityId              uint                         `json:"entityId"` // Debug Interface Id for ProcessorInterface type
	EndpointInterfaceId   uint                         `json:"endpointInterfaceId"`
	ProcessorInterfaceSrc consts.ProcessorInterfaceSrc `json:"processorInterfaceSrc"`

	Ordr     int          `json:"ordr"`
	Children []*Processor `gorm:"-" json:"children"`
	Slots    iris.Map     `gorm:"-" json:"slots"`
}

func (Processor) TableName() string {
	return "biz_processor"
}

//type ProcessorThreadGroup struct {
//	BaseModel
//	agentExec.ProcessorEntityBase
//
//	Count int `json:"count" yaml:"count"`
//	Loop  int `json:"loop" yaml:"loop"`
//
//	StartupDelay int `json:"startupDelay" yaml:"startupDelay"`
//	RampUpPeriod int `json:"rampUpPeriod" yaml:"rampUpPeriod"`
//	Duration     int `json:"duration" yaml:"duration"`
//
//	ErrorAction consts.ErrorAction
//}
//
//func (ProcessorThreadGroup) TableName() string {
//	return "biz_processor_thread_group"
//}

type ProcessorGroup struct {
	BaseModel
	agentExec.ProcessorEntityBase
}

func (ProcessorGroup) TableName() string {
	return "biz_processor_group"
}

type ProcessorLogic struct {
	BaseModel
	agentExec.ProcessorEntityBase

	Expression string `json:"expression" yaml:"expression"`
}

func (ProcessorLogic) TableName() string {
	return "biz_processor_logic"
}

type ProcessorLoop struct {
	BaseModel
	agentExec.ProcessorEntityBase

	Times int    `json:"times" yaml:"times"` // time
	Range string `json:"range" yaml:"range"` // range

	InType   string `json:"inType" yaml:"inType"`     // in
	Variable string `json:"variable" yaml:"variable"` // variable
	List     string `json:"list" yaml:"list"`         // list
	Step     string `json:"step" yaml:"step"`

	IsRand       bool   `json:"isRand" yaml:"isRand"`
	VariableName string `json:"variableName" yaml:"variableName"`

	UntilExpression   string `json:"untilExpression" yaml:"untilExpression"` // until
	BreakIfExpression string `json:"breakIfExpression" yaml:"breakIfExpression"`
}

func (ProcessorLoop) TableName() string {
	return "biz_processor_loop"
}

type ProcessorTimer struct {
	BaseModel
	agentExec.ProcessorEntityBase

	SleepTime int `json:"sleepTime" yaml:"sleepTime"`

	Unit string `json:"unit" yaml:"unit"`
}

func (ProcessorTimer) TableName() string {
	return "biz_processor_timer"
}

type ProcessorPrint struct {
	BaseModel
	agentExec.ProcessorEntityBase

	RightValue string `json:"rightValue" yaml:"rightValue"`
}

func (ProcessorPrint) TableName() string {
	return "biz_processor_print"
}

type ProcessorVariable struct {
	BaseModel
	agentExec.ProcessorEntityBase

	VariableName string `json:"variableName" yaml:"variableName"`
	Expression   string `json:"expression" yaml:"expression"`
}

func (ProcessorVariable) TableName() string {
	return "biz_processor_variable"
}

type ProcessorAssertion struct {
	BaseModel
	agentExec.ProcessorEntityBase

	Expression string `json:"expression" yaml:"expression"`
}

func (ProcessorAssertion) TableName() string {
	return "biz_processor_assertion"
}

type ProcessorExtractor struct {
	BaseModel
	agentExec.ProcessorEntityBase

	Src  consts.ExtractorSrc  `json:"src"`
	Type consts.ExtractorType `json:"type"`
	Key  string               `json:"key"` // form header

	Expression string `json:"expression"`
	//NodeProp       string `json:"prop"`

	BoundaryStart    string `json:"boundaryStart"`
	BoundaryEnd      string `json:"boundaryEnd"`
	BoundaryIndex    int    `json:"boundaryIndex"`
	BoundaryIncluded bool   `json:"boundaryIncluded"`

	Variable string `json:"variable"`

	Result      string `json:"result"`
	InterfaceId uint   `json:"interfaceId"`
}

func (ProcessorExtractor) TableName() string {
	return "biz_processor_extractor"
}

type ProcessorData struct {
	BaseModel
	agentExec.ProcessorEntityBase

	Src  consts.DataItSrc  `json:"src" yaml:"src"`
	Type consts.DataItType `json:"type,omitempty" yaml:"type,omitempty"`

	Url        string `json:"url,omitempty" yaml:"url,omitempty"`
	DatapoolId uint   `json:"datapoolId,omitempty" yaml:"datapoolId,omitempty"`

	Separator string `json:"separator,omitempty" yaml:"separator,omitempty"`

	RepeatTimes int  `json:"repeatTimes,omitempty" yaml:"repeatTimes,omitempty"`
	IsRand      bool `json:"isRand,omitempty" yaml:"isRand,omitempty"`

	VariableName string `json:"variableName,omitempty" yaml:"variableName,omitempty"`
}

func (ProcessorData) TableName() string {
	return "biz_processor_data"
}

type ProcessorCookie struct {
	BaseModel
	agentExec.ProcessorEntityBase

	CookieName   string     `json:"cookieName" yaml:"cookieName"`
	VariableName string     `json:"variableName" yaml:"variableName"`
	RightValue   string     `json:"rightValue" yaml:"rightValue"`
	Domain       string     `json:"domain" yaml:"domain"`
	ExpireTime   *time.Time `json:"expireTime" yaml:"expireTime"`
}

func (ProcessorCookie) TableName() string {
	return "biz_processor_cookie"
}

type ProcessorCustomCode struct {
	BaseModel
	agentExec.ProcessorEntityBase

	Content string `gorm:"type:longtext;" json:"content" yaml:"content"`
	Desc    string `json:"desc" yaml:"desc"`

	//Output       string              `gorm:"type:longtext;" json:"output"`
	//ResultStatus consts.ResultStatus `json:"resultStatus"`
	//ResultMsg    string              `json:"resultMsg"`
}

func (ProcessorCustomCode) TableName() string {
	return "biz_processor_custom_code"
}

type ProcessorPerformanceRunner struct {
	BaseModel
	agentExec.ProcessorEntityBase

	Ip       string `json:"ip"`
	WebPort  uint   `json:"webPort"`
	GrpcPort uint   `json:"grpcPort"`
	Weight   uint   `json:"weight"`
}

func (ProcessorPerformanceRunner) TableName() string {
	return "biz_processor_performance_runner"
}

type ProcessorPerformanceScenario struct {
	BaseModel
	agentExec.ProcessorEntityBase

	GeneratorType ptconsts.GeneratorType `json:"generatorType"`

	Target int `json:"target"`

	Goal         ptconsts.GoalType `json:"goal"`
	Duration     int               `json:"duration"`
	Loop         int               `json:"loop"`
	ResponseTime float32           `json:"responseTime"`
	Qps          float32           `json:"qps"`
	FailRate     float32           `json:"failRate"`

	Stages []ProcessorPerformanceStage `gorm:"-" json:"stages"`
}

func (ProcessorPerformanceScenario) TableName() string {
	return "biz_processor_performance_scenario"
}

type ProcessorPerformanceStage struct {
	BaseModel
	agentExec.ProcessorEntityBase

	Duration int `json:"duration"`
	Target   int `json:"target"`
}

func (ProcessorPerformanceStage) TableName() string {
	return "biz_processor_performance_stage"
}

type ProcessorComm struct {
	Id uint `json:"id" yaml:"id"`
	agentExec.ProcessorEntityBase

	EntityId            uint              `json:"entityId"`
	EndpointInterfaceId uint              `json:"endpointInterfaceId"`
	Method              consts.HttpMethod `json:"method"`
	SrcName             string            `json:"srcName"`

	ProcessorInterfaceSrc consts.ProcessorInterfaceSrc `json:"processorInterfaceSrc"`
}
