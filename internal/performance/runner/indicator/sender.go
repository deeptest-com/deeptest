package indicator

import (
	ptlog "github.com/aaronchen2k/deeptest/internal/performance/pkg/log"
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
	// response time
	requests := GetRequests()
	requestCountSent += len(*requests)
	ClearRequests()

	ptlog.Logf("****** RUNNER DEBUG: totally %d requests sent to server", requestCountSent)

	// machine metrics
	machineMetrics := ptutils.GetMachineMetrics(prevDiskInfoMap, prevDiskTsMap, prevNetworkInfoMap, prevNetworkTsMap)

	// send results
	result := ptProto.PerformanceExecResp{
		Timestamp: time.Now().UnixMilli(),
		RunnerId:  runnerId,
		Room:      room,
		Requests:  *requests,

		Metrics: &ptProto.PerformanceExecMetrics{
			CpuUsage:      machineMetrics.CpuUsage,
			MemoryUsage:   machineMetrics.MemoryUsage,
			DiskUsages:    machineMetrics.DiskUsages,
			NetworkUsages: machineMetrics.NetworkUsages,
		},
	}

	sender.Send(result)
}
