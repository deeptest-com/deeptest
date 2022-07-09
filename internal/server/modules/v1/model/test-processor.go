package model

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/kataras/iris/v12"
)

type TestProcessor struct {
	BaseModel

	Name     string `json:"name" yaml:"name"`
	Comments string `json:"comments" yaml:"comments"`

	ParentId   uint `json:"parentId"`
	ScenarioId uint `json:"scenarioId"`
	UseID      uint `json:"useId"`

	EntityCategory consts.ProcessorCategory `json:"entityCategory"`
	EntityType     consts.ProcessorType     `json:"entityType"`
	EntityId       uint                     `json:"entityId"`
	InterfaceId    uint                     `json:"interfaceId"`

	Ordr     int              `json:"ordr"`
	Children []*TestProcessor `gorm:"-" json:"children"`
	Slots    iris.Map         `gorm:"-" json:"slots"`
}

func (TestProcessor) TableName() string {
	return "biz_test_processor"
}

type ProcessorInterface struct {
	BaseModel
	ProcessorEntity

	Name string `json:"name"`
	Desc string `json:"desc"`

	IsDir     bool `json:"isDir"`
	ParentId  uint `json:"parentId"`
	ProjectId uint `json:"projectId"`
	UseID     uint `json:"useId"`

	Ordr     int          `json:"ordr"`
	Children []*Interface `gorm:"-" json:"children"`

	Slots iris.Map `gorm:"-" json:"slots"`

	Url               string                 `json:"url"`
	Method            string                 `gorm:"default:GET" json:"method"`
	Params            []InterfaceParam       `gorm:"-" json:"params"`
	Headers           []InterfaceHeader      `gorm:"-" json:"headers"`
	Body              string                 `gorm:"default:{}" json:"body"`
	BodyType          consts.HttpContentType `gorm:"default:''" json:"bodyType"`
	AuthorizationType string                 `gorm:"default:''" json:"authorizationType"`
	PreRequestScript  string                 `gorm:"default:''" json:"preRequestScript"`
	ValidationScript  string                 `gorm:"default:''" json:"validationScript"`

	BasicAuth   InterfaceBasicAuth   `gorm:"-" json:"basicAuth"`
	BearerToken InterfaceBearerToken `gorm:"-" json:"bearerToken"`
	OAuth20     InterfaceOAuth20     `gorm:"-" json:"oauth20"`
	ApiKey      InterfaceApiKey      `gorm:"-" json:"apiKey"`

	EnvironmentId uint `json:"environmentId"`

	InterfaceExtractors  []InterfaceExtractor  `gorm:"-" json:"interfaceExtractors"`
	InterfaceCheckpoints []InterfaceCheckpoint `gorm:"-" json:"interfaceCheckpoints"`
}

func (ProcessorInterface) TableName() string {
	return "biz_test_processor_interface"
}

type ProcessorThreadGroup struct {
	BaseModel
	ProcessorEntity

	Count int `json:"count" yaml:"count"`
	Loop  int `json:"loop" yaml:"loop"`

	StartupDelay int `json:"startupDelay" yaml:"startupDelay"`
	RampUpPeriod int `json:"rampUpPeriod" yaml:"rampUpPeriod"`
	Duration     int `json:"duration" yaml:"duration"`

	ErrorAction consts.ErrorAction
}

func (ProcessorThreadGroup) TableName() string {
	return "biz_test_processor_thread_group"
}

type ProcessorGroup struct {
	BaseModel
	ProcessorEntity
}

func (ProcessorGroup) TableName() string {
	return "biz_test_processor_group"
}

type ProcessorLogic struct {
	BaseModel
	ProcessorEntity

	LeftValue  string                    `json:"leftValue" yaml:"leftValue"`
	RightValue string                    `json:"rightValue" yaml:"rightValue"`
	Operator   consts.ComparisonOperator `json:"operator" yaml:"operator"`
}

func (ProcessorLogic) TableName() string {
	return "biz_test_processor_logic"
}

type ProcessorIterator struct {
	BaseModel
	ProcessorEntity

	Times int `json:"times" yaml:"times"` // how many
	Count int `json:"count" yaml:"count"` // left

	BreakIfExpression string `json:"breakExpr" yaml:"breakIfExpression"`
}

func (ProcessorIterator) TableName() string {
	return "biz_test_processor_iterator"
}

type ProcessorTimer struct {
	BaseModel
	ProcessorEntity

	SleepBefore int `json:"sleepBefore" yaml:"sleepBefore"`
	SleepAfter  int `json:"sleepAfter" yaml:"sleepAfter"`

	Unit string `json:"unit" yaml:"unit"`
}

func (ProcessorTimer) TableName() string {
	return "biz_test_processor_timer"
}

type ProcessorAssertion struct {
	BaseModel
	ProcessorEntity

	Expression string `json:"expression" yaml:"expression"`
	Expect     string `json:"expect" yaml:"expect"`
}

func (ProcessorAssertion) TableName() string {
	return "biz_test_processor_assertion"
}

type ProcessorExtractor struct {
	BaseModel
	ProcessorEntity

	Src  consts.ExtractorSrc  `json:"src"`
	Type consts.ExtractorType `json:"type"`
	Key  string               `json:"key"`

	Expression string `json:"expression"`
	Prop       string `json:"prop"`

	BoundaryStart    string `json:"boundaryStart"`
	BoundaryEnd      string `json:"boundaryEnd"`
	BoundaryIndex    int    `json:"boundaryIndex"`
	BoundaryIncluded bool   `json:"boundaryIncluded"`

	Variable string `json:"variable"`

	Result      string `json:"result"`
	InterfaceId uint   `json:"interfaceId"`
}

func (ProcessorExtractor) TableName() string {
	return "biz_test_processor_extractor"
}

type ProcessorData struct {
	BaseModel
	ProcessorEntity

	Type consts.DataSource `json:"type,omitempty" yaml:"type,omitempty"`
	Path string            `json:"path,omitempty" yaml:"path,omitempty"`

	Loop           int    `json:"loop,omitempty" yaml:"loop,omitempty"`
	StartIndex     int    `json:"startIndex,omitempty" yaml:"startIndex,omitempty"`
	EndIndex       int    `json:"endIndex,omitempty" yaml:"endIndex,omitempty"`
	IsRand         bool   `json:"isRand,omitempty" yaml:"isRand,omitempty"`
	IsOnce         bool   `json:"isOnce,omitempty" yaml:"isOnce,omitempty"`
	VarNamePostfix string `json:"varNamePostfix,omitempty" yaml:"varNamePostfix,omitempty"`
}

func (ProcessorData) TableName() string {
	return "biz_test_processor_data"
}

type ProcessorCookie struct {
	BaseModel
	ProcessorEntity

	Action   consts.ValueAction `json:"action"`
	Name     string             `json:"name"`
	Variable string             `json:"variable"`
}

func (ProcessorCookie) TableName() string {
	return "biz_test_processor_cookie"
}

type ProcessorComm struct {
	Id uint `json:"id" yaml:"id"`
	ProcessorEntity
}

type ProcessorEntity struct {
	Name     string `gorm:"-" json:"name" yaml:"name"`
	Comments string `json:"comments" yaml:"comments"`

	ProcessorId       uint                     `json:"processorId" yaml:"processorId"`
	ProcessorCategory consts.ProcessorCategory `json:"processorCategory" yaml:"processorCategory"`
	ProcessorType     consts.ProcessorType     `json:"processorType" yaml:"processorType"`

	//ParentId uint `json:"parentId" yaml:"parentId"`

	// interface or Processor
	Children []interface{} `json:"children" yaml:"children" gorm:"-"`
	// results
	Results []string `json:"results" yaml:"results" gorm:"-"`
}
