package domain

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"time"
)

type ExecVariable struct {
	Id         uint        `json:"id"`
	Name       string      `json:"name"`
	Value      interface{} `json:"value"`
	Expression string      `json:"expression"`

	InterfaceId uint                  `json:"interfaceId"`
	Scope       consts.ExtractorScope `json:"isShare"`
}

type ExecCookie struct {
	Id    uint        `json:"id,omitempty"`
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
	Path  string      `json:"path,omitempty"`

	Domain     string     `json:"domain,omitempty"`
	ExpireTime *time.Time `json:"expireTime,omitempty"`
}

type ExecLog struct {
	Id             uint                  `json:"id"`
	Name           string                `json:"name"`
	Desc           string                `json:"desc"`
	ProgressStatus consts.ProgressStatus `json:"progressStatus"`
	ResultStatus   consts.ResultStatus   `json:"resultStatus"`
	StartTime      *time.Time            `json:"startTime"`
	EndTime        *time.Time            `json:"endTime"`

	ParentId     uint `json:"parentId"`
	PersistentId uint `json:"persistentId"`
	ReportId     uint `json:"reportId"`

	Logs *[]*ExecLog `json:"logs"`

	// type
	ProcessorCategory consts.ProcessorCategory `json:"processorCategory"`

	// for interface
	InterfaceId uint   `json:"interfaceId"`
	ReqContent  string `json:"reqContent,omitempty"`
	RespContent string `json:"respContent,omitempty"`

	// for processor
	ProcessorType  consts.ProcessorType `json:"processorType"`
	ProcessId      uint                 `json:"processId,omitempty"`
	ProcessContent string               `json:"processContent,omitempty"`
	ProcessResult  string               `json:"processResult,omitempty"`

	InterfaceExtractorsResult  []ExecInterfaceExtractor  `gorm:"-" json:"interfaceExtractorsResult,omitempty"`
	InterfaceCheckpointsResult []ExecInterfaceCheckpoint `gorm:"-" json:"interfaceCheckpointsResult,omitempty"`

	Summary []string   `json:"summary,omitempty"`
	Output  ExecOutput `json:"output,omitempty"`
}

type ExecInterfaceExtractor struct {
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
type ExecInterfaceCheckpoint struct {
	Type consts.CheckpointType `json:"type"`

	Expression        string `json:"expression"`
	ExtractorVariable string `json:"extractorVariable"`

	Operator consts.ComparisonOperator `json:"operator"`
	Value    string                    `json:"value"`

	ResultStatus consts.ResultStatus `json:"resultStatus"`
	InterfaceId  uint                `json:"interfaceId"`
}

type ExecIterator struct {
	ProcessorCategory consts.ProcessorCategory
	ProcessorType     consts.ProcessorType
	VariableName      string `json:"variableName,omitempty"`

	// loop range
	Items    []interface{}            `json:"items"`
	Data     []map[string]interface{} `json:"data"`
	DataType consts.DataType          `json:"dataType"`

	// loop condition
	UntilExpression string `json:"untilExpression"`
}

type ExecOutput struct {
	// logic if, else
	Pass bool `json:"pass,omitempty"`

	// loop - times
	Times int `json:"times,omitempty"`
	// loop util
	Expression string `json:"times,omitempty"`
	// loop in
	List string `json:"list,omitempty"`
	// loop - range
	Range      string          `json:"range,omitempty"`
	RangeStart interface{}     `json:"rangeStart,omitempty"`
	RangeEnd   interface{}     `json:"rangeEnd,omitempty"`
	RangeType  consts.DataType `json:"rangeType,omitempty"`
	// loop break
	BreakFrom uint `json:"breakFrom,omitempty"`

	// timer
	SleepTime int `json:"sleepTime"`

	// data
	Url          string `json:"url"`
	RepeatTimes  int    `json:"repeatTimes,omitempty"`
	IsLoop       int    `json:"isLoop,omitempty"`
	IsRand       bool   `json:"isRand,omitempty"`
	IsOnce       bool   `json:"isOnce,omitempty"`
	VariableName string `json:"variableName,omitempty"`

	// extractor
	Src  consts.ExtractorSrc  `json:"src"`
	Type consts.ExtractorType `json:"type"`
	//Expression string `json:"expression"`
	BoundaryStart    string `json:"boundaryStart"`
	BoundaryEnd      string `json:"boundaryEnd"`
	BoundaryIndex    int    `json:"boundaryIndex"`
	BoundaryIncluded bool   `json:"boundaryIncluded"`
	Variable         string `json:"variable"`

	// variable
	VariableValue interface{} `json:"variableValue"`

	// common
	Msg string `json:"msg,omitempty"`
}
