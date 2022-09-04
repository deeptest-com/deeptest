package model

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/kataras/iris/v12"
	"time"
)

type Processor struct {
	BaseModel

	Name     string `json:"name" yaml:"name"`
	Comments string `json:"comments" yaml:"comments"`

	ParentId   uint `json:"parentId"`
	ScenarioId uint `json:"scenarioId"`
	ProjectId  uint `json:"projectId"`
	UseID      uint `json:"useId"`

	EntityCategory consts.ProcessorCategory `json:"entityCategory"`
	EntityType     consts.ProcessorType     `json:"entityType"`
	EntityId       uint                     `json:"entityId"`
	InterfaceId    uint                     `json:"interfaceId"`

	Ordr     int          `json:"ordr"`
	Children []*Processor `gorm:"-" json:"children"`
	Slots    iris.Map     `gorm:"-" json:"slots"`
}

func (Processor) TableName() string {
	return "biz_processor"
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
	return "biz_processor_thread_group"
}

type ProcessorGroup struct {
	BaseModel
	ProcessorEntity
}

func (ProcessorGroup) TableName() string {
	return "biz_processor_group"
}

type ProcessorLogic struct {
	BaseModel
	ProcessorEntity

	Expression string `json:"expression" yaml:"expression"`
}

func (ProcessorLogic) TableName() string {
	return "biz_processor_logic"
}

type ProcessorLoop struct {
	BaseModel
	ProcessorEntity

	Times        int    `json:"times" yaml:"times"` // time
	Range        string `json:"range" yaml:"range"` // range
	List         string `json:"list" yaml:"list"`   // in
	Step         string `json:"step" yaml:"step"`
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
	ProcessorEntity

	SleepTime int `json:"sleepTime" yaml:"sleepTime"`

	Unit string `json:"unit" yaml:"unit"`
}

func (ProcessorTimer) TableName() string {
	return "biz_processor_timer"
}

type ProcessorVariable struct {
	BaseModel
	ProcessorEntity

	VariableName string `json:"variableName" yaml:"variableName"`
	RightValue   string `json:"rightValue" yaml:"rightValue"`
}

func (ProcessorVariable) TableName() string {
	return "biz_processor_variable"
}

type ProcessorAssertion struct {
	BaseModel
	ProcessorEntity

	Expression string `json:"expression" yaml:"expression"`
}

func (ProcessorAssertion) TableName() string {
	return "biz_processor_assertion"
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
	return "biz_processor_extractor"
}

type ProcessorData struct {
	BaseModel
	ProcessorEntity

	Type      consts.DataSource `json:"type,omitempty" yaml:"type,omitempty"`
	Url       string            `json:"url,omitempty" yaml:"url,omitempty"`
	Separator string            `json:"separator,omitempty" yaml:"separator,omitempty"`

	RepeatTimes int `json:"repeatTimes,omitempty" yaml:"repeatTimes,omitempty"`
	//StartIndex     int    `json:"startIndex,omitempty" yaml:"startIndex,omitempty"`
	//EndIndex       int    `json:"endIndex,omitempty" yaml:"endIndex,omitempty"`

	IsLoop int  `json:"isLoop,omitempty" yaml:"isLoop,omitempty"`
	IsRand bool `json:"isRand,omitempty" yaml:"isRand,omitempty"`
	IsOnce bool `json:"isOnce,omitempty" yaml:"isOnce,omitempty"`

	VariableName string `json:"variableName,omitempty" yaml:"variableName,omitempty"`
}

func (ProcessorData) TableName() string {
	return "biz_processor_data"
}

type ProcessorCookie struct {
	BaseModel
	ProcessorEntity

	CookieName   string     `json:"cookieName" yaml:"cookieName"`
	VariableName string     `json:"variableName" yaml:"variableName"`
	RightValue   string     `json:"rightValue" yaml:"rightValue"`
	Domain       string     `json:"domain" yaml:"domain"`
	ExpireTime   *time.Time `json:"expireTime" yaml:"expireTime"`
}

func (ProcessorCookie) TableName() string {
	return "biz_processor_cookie"
}

type ProcessorComm struct {
	Id uint `json:"id" yaml:"id"`
	ProcessorEntity
	InterfaceId uint `json:"interfaceId"`
}

type ProcessorEntity struct {
	Name     string `gorm:"-" json:"name" yaml:"name"`
	Comments string `json:"comments" yaml:"comments"`
	Default  string `json:"default" yaml:"default"`

	ProcessorId       uint                     `json:"processorId" yaml:"processorId"`
	ProcessorCategory consts.ProcessorCategory `json:"processorCategory" yaml:"processorCategory"`
	ProcessorType     consts.ProcessorType     `json:"processorType" yaml:"processorType"`

	//ParentId uint `json:"parentId" yaml:"parentId"`

	// interface or Processor
	Children []interface{} `json:"children" yaml:"children" gorm:"-"`
}
