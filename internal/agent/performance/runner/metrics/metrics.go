package metrics

import (
	ptproto "github.com/aaronchen2k/deeptest/internal/agent/performance/proto"
	commUtils "github.com/aaronchen2k/deeptest/internal/agent/pkg/utils"
	"time"
)

var (
	requestCountSent = 0
)

type MessageSender interface {
	Send(result ptproto.PerformanceExecResp) error
}

func SendMetrics(sender MessageSender, runnerId int32, runnerName, room string, prevDiskInfoMap *map[string]*int64, prevDiskTsMap *map[string]*int64, prevNetworkInfoMap *map[string]*int64, prevNetworkTsMap *map[string]*int64) {
	// machine metrics
	machineMetrics := commUtils.GetMachineMetrics(prevDiskInfoMap, prevDiskTsMap, prevNetworkInfoMap, prevNetworkTsMap)

	// send results
	result := ptproto.PerformanceExecResp{
		Timestamp:  time.Now().UnixMilli(),
		RunnerId:   runnerId,
		RunnerName: runnerName,
		Room:       room,

		Metrics: &ptproto.PerformanceExecMetrics{
			CpuUsage:      machineMetrics.CpuUsage,
			MemoryUsage:   machineMetrics.MemoryUsage,
			DiskUsages:    machineMetrics.DiskUsages,
			NetworkUsages: machineMetrics.NetworkUsages,
		},
	}

	sender.Send(result)
}
