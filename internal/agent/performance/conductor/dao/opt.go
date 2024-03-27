package dao

import (
	ptconsts "github.com/aaronchen2k/deeptest/internal/agent/performance/pkg/consts"
	ptdomain "github.com/aaronchen2k/deeptest/internal/agent/performance/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"time"
)

func ClearData(room string) (err error) {
	db := GetDB(consts.AppAgent)

	condtion := true // "created_at < ?"
	oneHourAgo := time.Now().Add(-1 * time.Minute)

	err = db.Where(condtion, oneHourAgo).Delete(&BizRequest{}).Error
	if err != nil {
		return
	}

	err = db.Where(condtion, oneHourAgo).Delete(&BizMetrics{}).Error
	if err != nil {
		return
	}

	err = db.Where(condtion, oneHourAgo).Delete(&BizDiskUsage{}).Error
	if err != nil {
		return
	}

	err = db.Where(condtion, oneHourAgo).Delete(&BizNetworkUsage{}).Error
	if err != nil {
		return
	}

	return
}

func CountAllByStatus(room string) (total, pass, fail, err int) {
	var list []ptdomain.StatusModel

	GetDB(consts.AppAgent).
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
	db := GetDB(consts.AppAgent)

	db.Raw("SELECT AVG(duration) AS `duration` FROM biz_request WHERE ROOM = ?", room).
		Scan(&responseTime)

	var total int
	db.Raw("SELECT COUNT(id) AS `count` FROM biz_request WHERE ROOM = ?", room).
		Scan(&total)

	qps = float64(total) / float64(duration) / 1000

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

	GetDB(consts.AppAgent).Create(&po)
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

	GetDB(consts.AppAgent).Create(&report)
}
