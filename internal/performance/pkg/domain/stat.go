package ptdomain

import (
	ptconsts "github.com/aaronchen2k/deeptest/internal/performance/pkg/consts"
	ptProto "github.com/aaronchen2k/deeptest/internal/performance/proto"
)

type Stat struct {
	StartTime int64
	EndTime   int64
	Duration  int64

	Requests []*ptProto.PerformanceExecRecord `json:"requests"`

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

type StatData struct {
	Name              string             `json:"name,omitempty"`
	Ip                string             `json:"ip,omitempty"`
	CpuUsage          float64            `json:"cpuUsage,omitempty"`
	MemoryUsage       float64            `json:"memoryUsage,omitempty"`
	DiskUsages        map[string]float64 `json:"diskUsages,omitempty"`
	NetworkUsages     map[string]float64 `json:"networks,omitempty"`
	MaxGoroutines     int32              `json:"maxGoroutines,omitempty"`
	CurrentGoroutines int32              `json:"currentGoroutines,omitempty"`
	ServerType        int32              `json:"serverType,omitempty"`
}
