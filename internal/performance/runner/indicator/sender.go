package indicator

import (
	ptutils "github.com/aaronchen2k/deeptest/internal/performance/pkg/utils"
	ptProto "github.com/aaronchen2k/deeptest/internal/performance/proto"
	"time"
)

var (
	requestCountSent = 0
)

type MessageSender interface {
	Send(result ptProto.PerformanceExecResp) error
}

func SendMetrics(sender MessageSender, runnerId int32, room string, prevDiskInfoMap *map[string]*int64, prevDiskTsMap *map[string]*int64, prevNetworkInfoMap *map[string]*int64, prevNetworkTsMap *map[string]*int64) {
	// machine metrics
	machineMetrics := ptutils.GetMachineMetrics(prevDiskInfoMap, prevDiskTsMap, prevNetworkInfoMap, prevNetworkTsMap)

	// send results
	result := ptProto.PerformanceExecResp{
		Timestamp: time.Now().UnixMilli(),
		RunnerId:  runnerId,
		Room:      room,

		Metrics: &ptProto.PerformanceExecMetrics{
			CpuUsage:      machineMetrics.CpuUsage,
			MemoryUsage:   machineMetrics.MemoryUsage,
			DiskUsages:    machineMetrics.DiskUsages,
			NetworkUsages: machineMetrics.NetworkUsages,
		},
	}

	sender.Send(result)
}
