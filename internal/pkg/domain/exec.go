package domain

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"time"
)

type Variable struct {
	Id         uint        `json:"id"`
	Name       string      `json:"name"`
	Value      interface{} `json:"value"`
	Expression string      `json:"expression"`
}

type Cookie struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`

	Domain     string    `json:"domain"`
	ExpireTime time.Time `json:"expireTime"`
}

type Log struct {
	Id             uint                  `json:"id"`
	Name           string                `json:"name"`
	Desc           string                `json:"desc"`
	ProgressStatus consts.ProgressStatus `json:"progressStatus"`
	ResultStatus   consts.ResultStatus   `json:"resultStatus"`
	StartTime      *time.Time            `json:"startTime"`
	EndTime        *time.Time            `json:"endTime"`

	ParentId uint    `json:"parentId"`
	Logs     *[]*Log `json:"logs"`

	// type
	ProcessorCategory consts.ProcessorCategory `json:"processorCategory"`

	// for interface
	InterfaceId uint     `json:"interfaceId"`
	ReqContent  string   `json:"reqContent,omitempty"`
	RespContent string   `json:"respContent,omitempty"`
	RespSummary []string `json:"respSummary,omitempty"`
	Output      Output   `json:"output,omitempty"`

	// for processor
	ProcessorType  consts.ProcessorType `json:"processorType"`
	ProcessId      uint                 `json:"processId,omitempty"`
	ProcessContent string               `json:"processContent,omitempty"`
	ProcessResult  string               `json:"processResult,omitempty"`
}

type ExecIterator struct {
	ProcessorCategory consts.ProcessorCategory
	ProcessorType     consts.ProcessorType

	// loop times
	Times []int `json:"times"`

	// loop range
	Items     []interface{}    `json:"items"`
	RangeType consts.RangeType `json:"rangeType"`
}

type Output struct {
	// loop - times
	Times int `json:"times"`
	// loop - range
	Range      string           `json:"range"`
	RangeStart interface{}      `json:"rangeStart"`
	RangeEnd   interface{}      `json:"rangeEnd"`
	RangeType  consts.RangeType `json:"rangeType"`

	Text string `json:"text"`
}
