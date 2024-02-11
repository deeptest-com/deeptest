package indicator

import (
	"context"
	_logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"time"
)

func ScheduleJob(ctx context.Context, runnerId int32, room string, sender MessageSender) {
	prevDiskInfoMap := map[string]*int64{}
	prevDiskTsMap := map[string]*int64{}
	prevNetworkInfoMap := map[string]*int64{}
	prevNetworkTsMap := map[string]*int64{}

	for true {
		time.Sleep(2 * time.Second)

		_logUtils.Debug(">>>>>> start runner schedule job")

		SendMetrics(sender, runnerId, room, &prevDiskInfoMap, &prevDiskTsMap, &prevNetworkInfoMap, &prevNetworkTsMap)

		select {
		case <-ctx.Done():
			_logUtils.Debug("<<<<<<< stop runner schedule job")
			return

		default:
		}
	}
}
