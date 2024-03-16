package dao

import (
	ptconsts "github.com/aaronchen2k/deeptest/internal/agent/performance/pkg/consts"
	ptdomain "github.com/aaronchen2k/deeptest/internal/agent/performance/pkg/domain"
	ptProto "github.com/aaronchen2k/deeptest/internal/agent/performance/proto"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"time"
)

func ClearData(room string) {
	db := GetDB(consts.AppServer)

	condtion := "created_at < ?"
	oneHourAgo := time.Now().Add(-1 * time.Minute)

	db.Where(condtion, oneHourAgo).Delete(&BizRequest{})
	db.Where(condtion, oneHourAgo).Delete(&BizMetrics{})
	db.Where(condtion, oneHourAgo).Delete(&BizDiskUsage{})
	db.Where(condtion, oneHourAgo).Delete(&BizNetworkUsage{})
}

func ListLastMinRequestRecordForResponseTime(room string) (ret []BizRequest) {
	specificTime := time.Now().Add(-1 * time.Minute)

	db := GetDB(consts.AppServer).
		Select("interface_id, name, status, duration, start_time, end_time").
		Where("room = ? AND created_at > ?", room, specificTime).
		Order("duration ASC")

	db.Find(&ret)

	return
}
func ListAllRequestRecordForResponseTime(room string) (ret []BizRequest) {
	db := GetDB(consts.AppServer).
		Select("interface_id, name, status, duration, start_time, end_time").
		Where("room = ?", room).
		Order("duration ASC")

	db.Find(&ret)

	return
}
func ListRequestRecordForQps(lastId uint, room string) (ret []BizRequest) {
	db := GetDB(consts.AppServer).
		Where("id > ? AND room = ?", lastId, room)

	db.Find(&ret)

	return
}

func CountAllByStatus(room string) (total, pass, fail, err int) {
	var list []ptdomain.StatusModel

	GetDB(consts.AppServer).
		Raw("SELECT status, COUNT(id) AS `count` FROM biz_request WHERE ROOM = ? GROUP BY status", room).
		Find(&list)

	mp := map[ptconsts.ResultStatus]int{}

	for _, item := range list {
		mp[item.Status] = item.Count
	}

	pass = mp[ptconsts.Pass]
	fail = mp[ptconsts.Fail]
	err = mp[ptconsts.Error]
	total = pass + fail + err

	return
}

func CountAvgTime(room string, duration int64) (responseTime, qps float64) {
	db := GetDB(consts.AppServer)

	db.Raw("SELECT AVG(duration) AS `duration` FROM biz_request WHERE ROOM = ?", room).
		Scan(&responseTime)

	var total int
	db.Raw("SELECT COUNT(id) AS `count` FROM biz_request WHERE ROOM = ?", room).
		Scan(&total)

	qps = float64(total) / float64(duration) / 1000

	return
}

func InsertRequestRecord(records []*ptProto.PerformanceExecRecord, runnerId int32, room string) {
	db := GetDB(consts.AppServer)

	pos := []BizRequest{}

	for _, record := range records {
		po := BizRequest{
			StartTime: record.StartTime,
			EndTime:   record.EndTime,
			Duration:  int(record.Duration),
			Status:    ptconsts.ResultStatus(record.Status),

			Name:        record.RecordName,
			InterfaceId: int(record.RecordId),
			VuId:        int(record.VuId),
			RunnerId:    int(runnerId),
			Room:        room,
		}

		pos = append(pos, po)
	}

	db.Create(&pos)
}

func InsertMetricsRecord(metrics *ptProto.PerformanceExecMetrics, runnerId int32, timestamp int64) {
	db := GetDB(consts.AppServer)

	po := BizMetrics{
		Timestamp: timestamp,
		RunnerId:  int(runnerId),

		CpuUsage:    metrics.CpuUsage,
		MemoryUsage: metrics.MemoryUsage,
	}

	db.Create(&po)

	for key, val := range metrics.DiskUsages {
		child := BizDiskUsage{
			MetricsId: po.ID,
			Name:      key,
			Usage:     val,
		}
		db.Create(&child)
	}

	for key, val := range metrics.NetworkUsages {
		child := BizNetworkUsage{
			MetricsId: po.ID,
			Name:      key,
			Usage:     val,
		}
		db.Create(&child)
	}
}

func ListMetricsByRunner() (ret []ptdomain.PerformanceExecMetrics) {
	var pos []BizMetrics

	DB := GetDB(consts.AppServer)

	DB.Raw("SELECT MAX(id), * FROM biz_metrics GROUP BY runner_id").
		Find(&pos)

	for _, runner := range pos {
		to := ptdomain.PerformanceExecMetrics{
			RunnerId:    runner.RunnerId,
			CpuUsage:    runner.CpuUsage,
			MemoryUsage: runner.MemoryUsage,

			DiskUsages:    map[string]float64{},
			NetworkUsages: map[string]float64{},
		}

		var disks []BizDiskUsage
		DB.Where("metrics_id = ?", runner.ID).Find(&disks)
		for _, disk := range disks {
			to.DiskUsages[disk.Name] = disk.Usage
		}

		var networks []BizNetworkUsage
		DB.Where("metrics_id = ?", runner.ID).Find(&networks)
		for _, network := range networks {
			to.NetworkUsages[network.Name] = network.Usage
		}

		ret = append(ret, to)
	}

	return
}

func GetPercentNumbsByInterface(room string) (ret map[int]map[string]int) {
	ret = map[int]map[string]int{}

	var raws = []struct {
		InterfaceId uint
		Numb        int
	}{}

	GetDB(consts.AppServer).
		Model(&BizRequest{}).
		Select("count(id) AS numb, interface_id").
		Where("room = ?", room).
		Group("interface_id").
		Find(&raws)

	for _, item := range raws {
		percentNumbsMap := map[string]int{}

		total := item.Numb
		percentNumbsMap[ptconsts.ChartRespTime.String()] = total

		ret[int(item.InterfaceId)] = percentNumbsMap
	}

	return
}

func InsertReportItem(room, chart, series string, value float64, timestamp int64) {
	po := BizSummaryReportItem{
		Room:      room,
		Chart:     chart,
		Series:    series,
		Value:     value,
		Timestamp: timestamp,
	}

	GetDB(consts.AppServer).Create(&po)
}

func SaveReport(room string, startTime, endTime int64) {
	report := BizSummaryReport{
		Room:      room,
		StartTime: time.UnixMilli(startTime),
		EndTime:   time.UnixMilli(endTime),
		Duration:  endTime - startTime,
	}

	report.Total, report.Pass, report.Fail, report.Error = CountAllByStatus(room)

	report.AvgResponseTime, report.AvgQps = CountAvgTime(room, report.Duration)

	GetDB(consts.AppServer).Create(&report)
}
