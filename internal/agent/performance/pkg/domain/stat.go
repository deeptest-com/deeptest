package ptdomain

import (
	ptconsts "github.com/aaronchen2k/deeptest/internal/agent/performance/pkg/consts"
)

type Stat struct {
	StartTime int64
	EndTime   int64
	Duration  int64

	Pass  int `json:"pass"`
	Fail  int `json:"fail"`
	Error int `json:"error"`

	AvgQps      float64 `json:"avgQps"`
	AvgDuration int     `json:"avgDuration"`

	AvgDuration50 int `json:"avgDuration50"`
	AvgDuration90 int `json:"avgDuration90"`
	AvgDuration95 int `json:"avgDuration95"`
}

type ResponseTimeData struct {
	// RecordId string `json:"recordId"`

	Status    ptconsts.ResultStatus `json:"status"`
	Timestamp int64                 `json:"timestamp"`

	ValueAll int `json:"valueAll"`
	Value50  int `json:"value50"`
	Value90  int `json:"value90"`
	Value95  int `json:"value95"`

	Durations *[]int `json:"Durations"`
}

type StatusModel struct {
	Status ptconsts.ResultStatus `json:"status"`
	Count  int                   `json:"count"`
}
