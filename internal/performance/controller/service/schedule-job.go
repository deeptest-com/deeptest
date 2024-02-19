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
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/kataras/iris/v12/websocket"
	"time"
)

type ScheduleService struct {
	RunnerIdToNameMap map[int]string

	GrpcService         *GrpcService         `inject:"private"`
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

	influxdbClient := influxdb2.NewClient(req.InfluxdbAddress, req.InfluxdbToken)

	for true {
		time.Sleep(1 * time.Second)

		if time.Now().UnixMilli()-lastTime < 6*1000 {
			continue
		}
		_logUtils.Debug(">>>>>> start server schedule job")

		summary, _ := dao.QueryResponseTimeSummary(influxdbClient, req.InfluxdbOrg)
		vuCount, _ := dao.QueryVuCount(influxdbClient, req.InfluxdbOrg)

		lastAvgResponseTime, _ := dao.QueryLastAvgResponseTime(influxdbClient, req.InfluxdbOrg)
		lastQps, _ := dao.QueryLastQps(influxdbClient, req.InfluxdbOrg)

		responseTimeTable, _ := dao.QueryResponseTimeTableByInterface(influxdbClient, req.InfluxdbOrg)

		metrics, _ := dao.QueryMetrics(influxdbClient, req.InfluxdbOrg)

		data := ptdomain.PerformanceExecResults{
			Timestamp: time.Now().UnixMilli(),

			VuCount: vuCount,
			Summary: summary,

			ReqResponseTime:      lastAvgResponseTime,
			ReqQps:               lastQps,
			ReqResponseTimeTable: responseTimeTable,
			Metrics:              metrics,
		}

		s.SendMetricsByWebsocket(data, req.Room, wsMsg)

		s.SaveReportItems(data, req.Room)

		if IsGoalMet(req, summary.Mean, summary.Qps, int32(summary.Fail+summary.Error), int32(summary.Total)) {
			execCancel()

			s.RemoteRunnerService.CallStop(req)
		}

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
	dao.InsertReportItem(room, ptconsts.ChartSummaryStatusCount.String(), "status_total", float64(summary.Total), timestamp)
	dao.InsertReportItem(room, ptconsts.ChartSummaryStatusCount.String(), "status_pass", float64(summary.Pass), timestamp)
	dao.InsertReportItem(room, ptconsts.ChartSummaryStatusCount.String(), "status_fail", float64(summary.Fail), timestamp)
	dao.InsertReportItem(room, ptconsts.ChartSummaryStatusCount.String(), "status_error", float64(summary.Error), timestamp)

	dao.InsertReportItem(room, ptconsts.ChartSummaryResponseTime.String(), "min", summary.Min, timestamp)
	dao.InsertReportItem(room, ptconsts.ChartSummaryResponseTime.String(), "max", summary.Max, timestamp)
	dao.InsertReportItem(room, ptconsts.ChartSummaryResponseTime.String(), "mean", summary.Mean, timestamp)
	dao.InsertReportItem(room, ptconsts.ChartSummaryResponseTime.String(), "median", summary.Median, timestamp)
	dao.InsertReportItem(room, ptconsts.ChartSummaryResponseTime.String(), "quantile95", summary.Quantile95, timestamp)

	dao.InsertReportItem(room, ptconsts.ChartSummaryQps.String(), "", summary.Qps, timestamp)

	vuCount := data.VuCount
	dao.InsertReportItem(room, ptconsts.ChartSummaryVuCount.String(), "", float64(vuCount), timestamp)

	reqResponseTime := data.ReqResponseTime
	for _, item := range reqResponseTime {
		dao.InsertReportItem(room, ptconsts.ChartRespTime.String(), item.RecordName, float64(item.Value), timestamp)
	}

	reqQps := data.ReqQps
	for _, item := range reqQps {
		dao.InsertReportItem(room, ptconsts.ChartQps.String(), item.RecordName, item.Value, timestamp)
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
