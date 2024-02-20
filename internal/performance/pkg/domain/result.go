package ptdomain

type PerformanceExecResults struct {
	Timestamp int64 `json:"timestamp,omitempty"`

	VuCount int                    `json:"vuCount,omitempty"`
	Summary PerformanceExecSummary `json:"summary,omitempty"`

	ReqResponseTime []PerformanceRequestResponseTime `json:"reqResponseTime,omitempty"`

	ReqQps []PerformanceRequestQps `json:"reqQps,omitempty"`

	ReqResponseTimeTable []PerformanceRequestTable `json:"reqResponseTimeTable,omitempty"`

	Metrics []PerformanceExecMetrics `json:"metrics,omitempty"`
}

type PerformanceExecSummary struct {
	StartTime int64 `json:"startTime,omitempty"`
	EndTime   int64 `json:"endTime,omitempty"`
	Duration  int64 `json:"duration,omitempty"`

	Total int `json:"total,omitempty"`
	Pass  int `json:"pass,omitempty"`
	Fail  int `json:"fail,omitempty"`
	Error int `json:"error,omitempty"`

	Min    float64 `json:"min,omitempty"`
	Max    float64 `json:"max,omitempty"`
	Mean   float64 `json:"mean,omitempty"`
	Median float64 `json:"median,omitempty"`
	Qps    float64 `json:"qps,omitempty"`

	Quantile95 float64 `json:"quantile95,omitempty"`
}

type PerformanceRequestResponseTime struct {
	RecordId   int32  `json:"recordId,omitempty"`
	RecordName string `json:"recordName,omitempty"`
	Value      int32  `json:"value,omitempty"`
}

type PerformanceRequestQps struct {
	RecordId   int32   `json:"recordId,omitempty"`
	RecordName string  `json:"recordName,omitempty"`
	Value      float64 `json:"value,omitempty"`
	Total      int32   `json:"total,omitempty"`
}

type PerformanceRequestTable struct {
	RecordId   int32  `json:"recordId,omitempty"`
	RecordName string `json:"recordName,omitempty"`

	Count  int32   `json:"count,omitempty"`
	Min    int32   `json:"min,omitempty"`
	Max    int32   `json:"max,omitempty"`
	Mean   float64 `json:"mean,omitempty"`
	Median float64 `json:"median,omitempty"`

	Quantile95 float64 `json:"quantile95,omitempty"`
}

type PerformanceExecMetrics struct {
	Timestamp int64 `json:"timestamp,omitempty"`

	CpuUsage    float64 `json:"cpuUsage,omitempty"`
	MemoryUsage float64 `json:"memoryUsage,omitempty"`

	DiskUsages    map[string]float64 `json:"diskUsages,omitempty"`
	NetworkUsages map[string]float64 `json:"networkUsages,omitempty"`

	RunnerId   int    `json:"runnerId,omitempty"`
	RunnerName string `json:"runnerName,omitempty"`
}
