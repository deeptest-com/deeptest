package ptdomain

type PerformanceExecResults struct {
	Timestamp int64 `json:"timestamp,omitempty"`

	Summary PerformanceExecSummary `json:"summary,omitempty"`

	Metrics []PerformanceExecMetrics `json:"metrics,omitempty"`

	ReqQps []PerformanceRequestQps `json:"reqQps,omitempty"`

	ReqAllResponseTime []PerformanceRequestResponseTime `json:"reqAllResponseTime,omitempty"`
	Req50ResponseTime  []PerformanceRequestResponseTime `json:"req50ResponseTime,omitempty"`
	Req90ResponseTime  []PerformanceRequestResponseTime `json:"req90ResponseTime,omitempty"`
	Req95ResponseTime  []PerformanceRequestResponseTime `json:"req95ResponseTime,omitempty"`
}

type PerformanceExecSummary struct {
	StartTime int64 `json:"startTime,omitempty"`
	EndTime   int64 `json:"endTime,omitempty"`
	Duration  int64 `json:"duration,omitempty"`

	Total   int `json:"total,omitempty"`
	Pass    int `json:"pass,omitempty"`
	Fail    int `json:"fail,omitempty"`
	Error   int `json:"error,omitempty"`
	Unknown int `json:"unknown,omitempty"`

	AvgResponseTime float64 `json:"avgResponseTime,omitempty"`
	AvgQps          float64 `json:"avgQps,omitempty"`
}

type PerformanceRequestQps struct {
	RecordId   int32   `json:"recordId,omitempty"`
	RecordName string  `json:"recordName,omitempty"`
	Value      float64 `json:"value,omitempty"`
	Total      int32   `json:"total,omitempty"`
}

type PerformanceRequestResponseTime struct {
	RecordId   int32  `json:"recordId,omitempty"`
	RecordName string `json:"recordName,omitempty"`
	Value      int32  `json:"value,omitempty"`
	Total      int32  `json:"total,omitempty"`
}

type PerformanceExecMetrics struct {
	Timestamp int64 `json:"timestamp,omitempty"`

	CpuUsage    float64 `json:"cpuUsage,omitempty"`
	MemoryUsage float64 `json:"memoryUsage,omitempty"`

	DiskUsages    map[string]float64 `json:"diskUsages,omitempty"`
	NetworkUsages map[string]float64 `json:"networkUsages,omitempty"`

	RunnerId int `json:"runnerId,omitempty"`
}
