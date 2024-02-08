package exec

import (
	ptlog "github.com/aaronchen2k/deeptest/internal/performance/pkg/log"
	ptProto "github.com/aaronchen2k/deeptest/internal/performance/proto"
	"github.com/aaronchen2k/deeptest/internal/performance/runner/indicator"
	_httpUtils "github.com/aaronchen2k/deeptest/pkg/lib/http"
	_intUtils "github.com/aaronchen2k/deeptest/pkg/lib/int"
	"time"
)

func ExecInterfaceProcessor(processor *ptProto.Processor, room string, vuNo, index int) {
	startTime := time.Now().UnixMilli()

	_, err := _httpUtils.Get("http://111.231.16.35:9000/get")
	if err != nil {
		ptlog.Logf("http request failed, err %s", err.Error())
	}

	// simulate processor result
	r := _intUtils.GenUniqueRandNum(100, 300, 1)[0]
	duration := int(processor.Id)*1000 + _intUtils.GenUniqueRandNum(100, 3001, 1)[0]
	endTime := startTime + int64(duration)
	time.Sleep(time.Duration(r) * time.Millisecond)

	status := "pass"
	if index%3 == 0 {
		status = "fail"
	}

	record := ptProto.PerformanceExecRecord{
		RecordId:   processor.Id,
		RecordName: processor.Name,

		StartTime: startTime,
		EndTime:   endTime,
		Duration:  int32(duration), // 毫秒
		Status:    status,

		VuId: int32(vuNo),
	}

	// add request record to requestsCache, will be batch sent to server via grpc by scheduled job.
	indicator.AddRequest(&record)

	return
}
