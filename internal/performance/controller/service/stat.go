package controllerService

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/performance/controller/dao"
	serverUtils "github.com/aaronchen2k/deeptest/internal/performance/controller/utils"
	"github.com/aaronchen2k/deeptest/internal/performance/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/performance/pkg/domain"
	ptlog "github.com/aaronchen2k/deeptest/internal/performance/pkg/log"
	ptProto "github.com/aaronchen2k/deeptest/internal/performance/proto"
	_dateUtils "github.com/aaronchen2k/deeptest/pkg/lib/date"
	_floatUtils "github.com/aaronchen2k/deeptest/pkg/lib/float"
	_logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"strconv"
	"strings"
	"time"
)

type StatService struct {
	Room string

	lastRequestIdToComputerQps uint
	interfaceIdToNameMap       map[int32]string
}

func (s *StatService) ComputeResponseTimeByInterface(room string) (ret map[string][]ptdomain.PerformanceRequestResponseTime,
	summary ptdomain.PerformanceExecSummary) {

	computeStartTime := time.Now()
	_logUtils.Infof("===== start compute response time @ %d.", computeStartTime)

	ret = map[string][]ptdomain.PerformanceRequestResponseTime{}
	countByInterfaceMap := map[int]*int{}
	sumByInterfaceMap := map[int]*int{}
	sumOfAll := 0

	percentNumbsOfInterface := dao.GetPercentNumbsByInterface(room)

	requests := dao.ListLastMinRequestRecordForResponseTime(room)
	total := len(requests)
	ptlog.Logf("****** SERVER DEBUG: totally load %d requests in sqlite", total)

	for i := 0; i < total; i++ {
		request := requests[i]

		duration := request.Duration
		interfaceId := request.InterfaceId
		interfaceName := s.interfaceIdToNameMap[int32(interfaceId)]
		startTime := request.StartTime
		endTime := request.EndTime

		if summary.StartTime == 0 {
			summary.StartTime = startTime
		}
		summary.EndTime = endTime

		countOfInterface, ok := countByInterfaceMap[interfaceId]
		if !ok {
			v := 0
			countOfInterface = &v
			countByInterfaceMap[interfaceId] = countOfInterface
		}

		sumOfInterface, ok := sumByInterfaceMap[interfaceId]
		if !ok {
			v := 0
			sumOfInterface = &v
			sumByInterfaceMap[interfaceId] = sumOfInterface
		}

		*sumOfInterface += duration
		sumOfAll += duration
		*countOfInterface++

		percentNumbsMap := percentNumbsOfInterface[interfaceId]

		if *countOfInterface == percentNumbsMap[ptconsts.ChartRespTime50.String()] {
			ret[ptconsts.ChartRespTime50.String()] = append(ret[ptconsts.ChartRespTime50.String()], ptdomain.PerformanceRequestResponseTime{
				RecordId:   int32(interfaceId),
				RecordName: interfaceName,
				Value:      int32(*sumOfInterface / *countOfInterface),
			})

		} else if *countOfInterface == percentNumbsMap[ptconsts.ChartRespTime90.String()] {
			ret[ptconsts.ChartRespTime90.String()] = append(ret[ptconsts.ChartRespTime90.String()], ptdomain.PerformanceRequestResponseTime{
				RecordId:   int32(interfaceId),
				RecordName: interfaceName,
				Value:      int32(*sumOfInterface / *countOfInterface),
			})

		} else if *countOfInterface == percentNumbsMap[ptconsts.ChartRespTime95.String()] {
			ret[ptconsts.ChartRespTime95.String()] = append(ret[ptconsts.ChartRespTime95.String()], ptdomain.PerformanceRequestResponseTime{
				RecordId:   int32(interfaceId),
				RecordName: interfaceName,
				Value:      int32(*sumOfInterface / *countOfInterface),
			})

		} else if *countOfInterface == percentNumbsMap[ptconsts.ChartRespTimeAll.String()] {
			ret[ptconsts.ChartRespTimeAll.String()] = append(ret[ptconsts.ChartRespTimeAll.String()], ptdomain.PerformanceRequestResponseTime{
				RecordId:   int32(interfaceId),
				RecordName: interfaceName,
				Value:      int32(*sumOfInterface / *countOfInterface),
			})

		}
	}

	summary.Duration = summary.EndTime - summary.StartTime
	summary.AvgResponseTime = float64(sumOfAll) / float64(summary.Total)
	summary.AvgQps = float64(summary.Total) * 1000 / float64(summary.Duration)

	// for all records
	summary.Total, summary.Pass, summary.Fail, summary.Error, summary.Unknown = dao.CountAllByStatus(room)

	computeEndTime := time.Now()
	ptlog.Logf("****** SERVER DEBUG: end compute %d records' response time, duration: %d micro secs (%s - %s).",
		len(requests),
		computeEndTime.UnixMicro()-computeStartTime.UnixMicro(),
		_dateUtils.DateTimeStrFmt(computeStartTime, "15:04:05.000000"),
		_dateUtils.DateTimeStrFmt(computeEndTime, "15:04:05.000000"))

	return
}

func (s *StatService) ComputeQpsByInterface(room string) (ret []ptdomain.PerformanceRequestQps) {
	countOfInterfaceMap := map[string]*int{}
	totalCount := 0
	startTime := int64(0)
	endTime := int64(0)

	requests := dao.ListRequestRecordForQps(s.lastRequestIdToComputerQps, room)

	for _, request := range requests {
		interfaceId := request.InterfaceId
		interfaceName := s.interfaceIdToNameMap[int32(interfaceId)]

		if startTime == 0 {
			startTime = request.StartTime
		} else {
			if startTime > request.StartTime {
				startTime = request.StartTime
			}
		}
		if endTime < request.EndTime {
			endTime = request.EndTime
		}

		key := fmt.Sprintf("%d-%s", interfaceId, interfaceName)
		if countOfInterfaceMap[key] == nil {
			v := 0
			countOfInterfaceMap[key] = &v
		}
		*countOfInterfaceMap[key]++

		totalCount++

		s.lastRequestIdToComputerQps = request.ID
	}

	duration := endTime - startTime

	if duration == 0 {
		return
	}

	ret = append(ret, ptdomain.PerformanceRequestQps{
		RecordId:   0,
		RecordName: "all",
		Value:      _floatUtils.PointNumb(float64(totalCount*1000)/float64(duration), 2),
	})

	for key, count := range countOfInterfaceMap {
		arr := strings.Split(key, "-")
		interfaceId, _ := strconv.Atoi(arr[0])
		interfaceName := arr[1]

		ret = append(ret, ptdomain.PerformanceRequestQps{
			RecordId:   int32(interfaceId),
			RecordName: interfaceName,
			Value:      _floatUtils.PointNumb(float64(*count*1000)/float64(duration), 2),
		})
	}

	return
}

func (s *StatService) Reset(scenarios []*ptProto.Scenario) {
	s.lastRequestIdToComputerQps = uint(0)
	s.interfaceIdToNameMap = serverUtils.GetInterfaceMap(scenarios)
}

func (s *StatService) LoadMetricsByRunner() (ret []ptdomain.PerformanceExecMetrics) {
	ret = dao.ListMetricsByRunner()

	return
}
