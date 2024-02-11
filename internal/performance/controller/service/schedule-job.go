package controllerService

import (
	"context"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/performance/controller/dao"
	"github.com/aaronchen2k/deeptest/internal/performance/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/performance/pkg/domain"
	ptlog "github.com/aaronchen2k/deeptest/internal/performance/pkg/log"
	websocketHelper "github.com/aaronchen2k/deeptest/internal/performance/pkg/websocket"
	ptProto "github.com/aaronchen2k/deeptest/internal/performance/proto"
	_logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12/websocket"
	"time"
)

type ScheduleService struct {
	RunnerIdToNameMap map[int]string

	StatService         *StatService         `inject:"private"`
	RemoteRunnerService *RemoteRunnerService `inject:"private"`
}

func (s *ScheduleService) Reset(scenarios []*ptProto.Scenario) {
	s.StatService.Reset(scenarios)
}

func (s *ScheduleService) ScheduleJob(execCtx context.Context, execCancel context.CancelFunc,
	req ptdomain.PerformanceTestReq, wsMsg *websocket.Message) {

	s.genRunnerIdToNameMap(req.Runners)

	startTime := time.Now().UnixMilli()
	lastTime := startTime

	s.StatService.Room = wsMsg.Room

	for true {
		time.Sleep(1 * time.Second)

		if time.Now().UnixMilli()-lastTime < 6*1000 {
			continue
		}

		_logUtils.Debug(">>>>>> start server schedule job")

		//responseTimeData, summary := s.StatService.ComputeResponseTimeByInterface(req.Room)
		//
		//data := ptdomain.PerformanceExecResults{
		//	Timestamp: time.Now().UnixMilli(),
		//
		//	Summary: summary,
		//
		//	ReqAllResponseTime: responseTimeData[ptconsts.ChartRespTimeAll.String()],
		//	Req50ResponseTime:  responseTimeData[ptconsts.ChartRespTime50.String()],
		//	Req90ResponseTime:  responseTimeData[ptconsts.ChartRespTime90.String()],
		//	Req95ResponseTime:  responseTimeData[ptconsts.ChartRespTime95.String()],
		//
		//	ReqQps:  s.StatService.ComputeQpsByInterface(req.Room),
		//	Metrics: s.StatService.LoadMetricsByRunner(),
		//}
		//
		//s.SendMetricsByWebsocket(data, req.Room, wsMsg)
		//
		//s.SaveReportItems(data, req.Room)
		//
		//if IsGoalMet(req, summary.AvgResponseTime, summary.AvgQps, int32(summary.Fail+summary.Error), int32(summary.Total)) {
		//	execCancel()
		//
		//	s.RemoteRunnerService.CallStop(req)
		//}

		lastTime = time.Now().UnixMilli()

		select {
		case <-execCtx.Done():
			_logUtils.Debug("<<<<<<< stop server schedule job")
			return

		default:
		}
	}

	s.SaveReport(wsMsg.Room, startTime, lastTime)
}

func (s *ScheduleService) SendMetricsByWebsocket(result ptdomain.PerformanceExecResults, execUUid string, wsMsg *websocket.Message) {
	if wsMsg != nil {
		ptlog.Logf("888888888888 %d", result.Timestamp)

		websocketHelper.SendExecResultToClient(result, ptconsts.MsgResultRecord, execUUid, wsMsg)
	}
}

func (s *ScheduleService) SaveReport(room string, startTime int64, lastTime int64) {
	dao.SaveReport(room, startTime, lastTime)
}

func (s *ScheduleService) SaveReportItems(data ptdomain.PerformanceExecResults, room string) {
	timestamp := data.Timestamp

	summary := data.Summary
	dao.InsertReportItem(room, ptconsts.ChartStatusCount.String(), "status_pass", float64(summary.Pass), timestamp)
	dao.InsertReportItem(room, ptconsts.ChartStatusCount.String(), "status_fail", float64(summary.Fail), timestamp)
	dao.InsertReportItem(room, ptconsts.ChartStatusCount.String(), "status_error", float64(summary.Error), timestamp)
	dao.InsertReportItem(room, ptconsts.ChartStatusCount.String(), "status_unknown", float64(summary.Unknown), timestamp)

	reqAllResponseTime := data.ReqAllResponseTime
	for _, req := range reqAllResponseTime {
		dao.InsertReportItem(room, ptconsts.ChartRespTimeAll.String(), req.RecordName, float64(req.Value), timestamp)
	}

	req50ResponseTime := data.Req50ResponseTime
	for _, req := range req50ResponseTime {
		dao.InsertReportItem(room, ptconsts.ChartRespTime50.String(), req.RecordName, float64(req.Value), timestamp)
	}

	req90ResponseTime := data.Req90ResponseTime
	for _, req := range req90ResponseTime {
		dao.InsertReportItem(room, ptconsts.ChartRespTime90.String(), req.RecordName, float64(req.Value), timestamp)
	}

	req95ResponseTime := data.Req95ResponseTime
	for _, req := range req95ResponseTime {
		dao.InsertReportItem(room, ptconsts.ChartRespTime95.String(), req.RecordName, float64(req.Value), timestamp)
	}

	reqQps := data.ReqQps
	for _, req := range reqQps {
		dao.InsertReportItem(room, ptconsts.ChartQps.String(), req.RecordName, req.Value, timestamp)
	}

	metrics := data.Metrics
	for _, item := range metrics {
		dao.InsertReportItem(room, ptconsts.ChartCpuUsage.String(), s.RunnerIdToNameMap[item.RunnerId],
			item.CpuUsage, timestamp)
		dao.InsertReportItem(room, ptconsts.ChartMemoryUsage.String(), s.RunnerIdToNameMap[item.RunnerId],
			item.MemoryUsage, timestamp)

		for name, val := range item.DiskUsages {
			series := fmt.Sprintf("%s - %s", s.RunnerIdToNameMap[item.RunnerId], name)
			dao.InsertReportItem(room, ptconsts.ChartDiskUsages.String(), series, val, timestamp)
		}

		for name, val := range item.NetworkUsages {
			series := fmt.Sprintf("%s - %s", s.RunnerIdToNameMap[item.RunnerId], name)
			dao.InsertReportItem(room, ptconsts.ChartNetworkUsages.String(), series, val, timestamp)
		}
	}
}

func (s *ScheduleService) genRunnerIdToNameMap(runners []*ptProto.Runner) {
	s.RunnerIdToNameMap = map[int]string{}

	for _, runner := range runners {
		s.RunnerIdToNameMap[int(runner.Id)] = runner.Name
	}
}
