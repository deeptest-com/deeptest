package domain

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/pkg/consts"
	"time"
)

type TestReport struct {
	TestSetId uint `json:"testSetId" yaml:"testSetId"`

	Version float64 `json:"version" yaml:"version"`
	Name    string  `json:"name" yaml:"name"`
	Code    int     `json:"code"`
	Msg     string  `json:"msg"`

	StartTime time.Time `json:"startTime" yaml:"startTime"`
	EndTime   time.Time `json:"endTime" yaml:"endTime"`
	Duration  int       `json:"duration" yaml:"duration"` // sec

	TotalNum  int `json:"totalNum" yaml:"totalNum"`
	PassNum   int `json:"passNum" yaml:"passNum"`
	FailNum   int `json:"failNum" yaml:"failNum"`
	MissedNum int `json:"missedNum" yaml:"missedNum"`

	Payload interface{} `json:"payload"`
}

func (result *TestReport) Pass(msg string) {
	result.Code = _consts.ResultCodeSuccess.Int()
	result.Msg = msg
}
func (result *TestReport) Passf(str string, args ...interface{}) {
	result.Code = _consts.ResultCodeSuccess.Int()
	result.Msg = fmt.Sprintf(str+"\n", args...)
}

func (result *TestReport) Fail(msg string) {
	result.Code = _consts.ResultCodeFail.Int()
	result.Msg = msg
}

func (result *TestReport) Failf(str string, args ...interface{}) {
	result.Code = _consts.ResultCodeFail.Int()
	result.Msg = fmt.Sprintf(str+"\n", args...)
}

func (result *TestReport) IsSuccess() bool {
	return result.Code == _consts.ResultCodeSuccess.Int()
}
