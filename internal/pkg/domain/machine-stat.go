package domain

type MachineStat struct {
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
