package indicator

import (
	"fmt"
	ptdomain "github.com/aaronchen2k/deeptest/internal/performance/pkg/domain"
	ptlog "github.com/aaronchen2k/deeptest/internal/performance/pkg/log"
	ptProto "github.com/aaronchen2k/deeptest/internal/performance/proto"
	_floatUtils "github.com/aaronchen2k/deeptest/pkg/lib/float"
	"sync"
	"time"
)

const (
	keyStartTime = "startTime"
	keyEndTime   = "endTime"
	keyDuration  = "duration"

	keyTotalCount = "total"
	keyPassCount  = "pass"
	keyFailCount  = "fail"
	keyErrorCount = "error"

	keyRequests = "requests"

	keyAvgResponseTime = "avgResponseTime"
	keyAvgQps          = "avgQps"
)

var (
	requestsCache sync.Map

	summaryCache sync.Map

	requestResponseTimeCache sync.Map
	requestQpsCache          sync.Map

	otherCache sync.Map

	requestCountCollected = 0
)

func Init() {
	UpdateStartTime(time.Now().UnixMilli())
	UpdateEndTime(time.Now().UnixMilli())
	UpdateDuration(0)

	UpdateTotal(0)
	UpdatePass(0)
	UpdateFail(0)
	UpdateError(0)

	UpdateAvgResponseTime(0)
	UpdateAvgQps(0)

	clearRequestResponseTime()
}

func GetRequests() (ret *[]*ptProto.PerformanceExecRecord) {
	val, ok := requestsCache.Load(keyRequests)
	if !ok {
		val = &[]*ptProto.PerformanceExecRecord{}
		requestsCache.Store(keyRequests, val)
	}

	ret = val.(*[]*ptProto.PerformanceExecRecord)

	return
}
func AddRequest(val *ptProto.PerformanceExecRecord) {
	arr := GetRequests()
	*arr = append(*arr, val)

	requestCountCollected++
	ptlog.Logf("****** RUNNER DEBUG: totally %d requests collected", requestCountCollected)
}
func ClearRequests() {
	requestsCache.Delete(keyRequests)
}

func GetSummary() (ret ptdomain.Stat) {
	ret.StartTime = GetStartTime()
	ret.EndTime = GetEndTime()
	ret.Duration = GetDuration()

	ret.Pass = GetPass()
	ret.Fail = GetFail()
	ret.Error = GetError()

	return
}

func GetStartTime() (ret int64) {
	val, ok := summaryCache.Load(keyStartTime)

	if ok {
		ret = val.(int64)
	}

	return
}
func UpdateStartTime(val int64) {
	summaryCache.Store(keyStartTime, val)
}

func GetEndTime() (ret int64) {
	val, ok := summaryCache.Load(keyEndTime)
	if ok {
		ret = val.(int64)
	}

	return
}
func UpdateEndTime(val int64) {
	summaryCache.Store(keyEndTime, val)
}

func GetDuration() (ret int64) {
	val, ok := summaryCache.Load(keyDuration)
	if ok {
		ret = val.(int64)
	}

	return
}
func UpdateDuration(val int64) {
	summaryCache.Store(keyDuration, val)
}

func GetResponseTimeData(key string) (ret *ptdomain.ResponseTimeData) {
	val, ok := requestResponseTimeCache.Load(key)
	if ok {
		ret = val.(*ptdomain.ResponseTimeData)
	} else {
		ret = &ptdomain.ResponseTimeData{
			Durations: &[]int{},
		}
		requestResponseTimeCache.Store(key, ret)
	}

	return
}
func AddResponseTime(record ptProto.PerformanceExecRecord) {
	id := record.RecordId
	name := record.RecordName
	dur := record.Duration

	key := fmt.Sprintf("%d-%s", id, name)

	// add to durations arr
	recordData := GetResponseTimeData(key)
	durations := recordData.Durations

	*durations = append(*durations, int(dur))
}
func clearRequestResponseTime() {
	requestResponseTimeCache = sync.Map{}
}

func GetTotal() (ret int) {
	val, ok := summaryCache.Load(keyTotalCount)

	if ok {
		ret = val.(int)
	}

	return
}
func UpdateTotal(val int) {
	summaryCache.Store(keyTotalCount, val)
}
func AddTotal(count int) {
	old := GetTotal()
	UpdateTotal(old + count)
}

func GetPass() (ret int) {
	val, ok := summaryCache.Load(keyPassCount)

	if ok {
		ret = val.(int)
	}

	return
}
func UpdatePass(val int) {
	summaryCache.Store(keyPassCount, val)
}
func AddPass(count int) {
	old := GetPass()
	UpdatePass(old + count)
}

func GetFail() (ret int) {
	val, ok := summaryCache.Load(keyFailCount)
	if ok {
		ret = val.(int)
	}

	return
}
func UpdateFail(val int) {
	summaryCache.Store(keyFailCount, val)
}
func AddFail(count int) {
	old := GetPass()
	UpdateFail(old + count)
}

func GetError() (ret int) {
	val, ok := summaryCache.Load(keyErrorCount)
	if ok {
		ret = val.(int)
	}

	return
}
func UpdateError(val int) {
	summaryCache.Store(keyErrorCount, val)
}
func AddError(count int) {
	old := GetError()
	UpdateError(old + count)
}

func GetAvgResponseTime() (ret int) {
	val, ok := summaryCache.Load(keyAvgResponseTime)
	if ok {
		ret = val.(int)
	}

	return
}
func UpdateAvgResponseTime(val int) {
	summaryCache.Store(keyAvgResponseTime, val)
}

func GetAvgQps() (ret float64) {
	val, ok := summaryCache.Load(keyAvgQps)
	if ok {
		ret = val.(float64)

		ret = _floatUtils.PointNumb(ret, 2)
	}

	return
}
func UpdateAvgQps(val float64) {
	summaryCache.Store(keyAvgQps, val)
}

func UpdateRequestQpsCache(key string, val float64) {
	requestQpsCache.Store(key, val)
}
