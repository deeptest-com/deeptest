package dao

import (
	ptconsts "github.com/aaronchen2k/deeptest/internal/performance/pkg/consts"
	"time"
)

var (
	Models = []interface{}{
		&BizRequest{},
		&BizMetrics{},
		&BizDiskUsage{},
		&BizNetworkUsage{},
		&BizSummaryReport{},
		&BizSummaryReportItem{},
	}
)

type BizRequest struct {
	BaseModel

	Name   string                `json:"name,omitempty"`
	Status ptconsts.ResultStatus `json:"status,omitempty"`

	StartTime int64 `json:"startTime"`
	EndTime   int64 `json:"endTime"`
	Duration  int   `protobuf:"varint,6,opt,name=duration,proto3" json:"duration,omitempty"`

	InterfaceId int    `json:"interfaceId,omitempty"`
	VuId        int    `json:"vuId,omitempty"`
	RunnerId    int    `json:"runnerId,omitempty"`
	Room        string `json:"room,omitempty"`
}

func (BizRequest) TableName() string {
	return "biz_request"
}

type BizMetrics struct {
	BaseModel

	Timestamp     int64              `json:"timestamp"`
	CpuUsage      float64            `json:"cpuUsage,omitempty"`
	MemoryUsage   float64            `json:"memoryUsage,omitempty"`
	DiskUsages    map[string]float64 `gorm:"-" json:"diskUsages,omitempty"`
	NetworkUsages map[string]float64 `gorm:"-" json:"networkUsages,omitempty"`

	RunnerId int `json:"runnerId,omitempty"`
}

func (BizMetrics) TableName() string {
	return "biz_metrics"
}

type BizDiskUsage struct {
	BaseModel

	Name  string  `json:"name"`
	Usage float64 `json:"usage,omitempty"`

	MetricsId uint `json:"metricsId,omitempty"`
}

func (BizDiskUsage) TableName() string {
	return "biz_disk_usage"
}

type BizNetworkUsage struct {
	BaseModel

	Name  string  `json:"name"`
	Usage float64 `json:"usage,omitempty"`

	MetricsId uint `json:"metricsId,omitempty"`
}

func (BizNetworkUsage) TableName() string {
	return "biz_network_usage"
}

type BizSummaryReport struct {
	BaseModel

	Room string `json:"room"`

	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	Duration  int64     `json:"duration"`

	AvgResponseTime float64 `json:"avgResponseTime"`
	AvgQps          float64 `json:"avgQps"`

	Total   int `json:"total"`
	Pass    int `json:"pass"`
	Fail    int `json:"fail"`
	Error   int `json:"error"`
	Unknown int `json:"unknown"`
}

func (BizSummaryReport) TableName() string {
	return "biz_summary_report"
}

type BizSummaryReportItem struct {
	BaseModel

	Room string `json:"room"`

	Chart  string  `json:"chart"`
	Series string  `json:"series"`
	Value  float64 `json:"value"`

	Timestamp int64 `json:"timestamp"`
}

func (BizSummaryReportItem) TableName() string {
	return "biz_summary_report_item"
}

type BaseModel struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
}
