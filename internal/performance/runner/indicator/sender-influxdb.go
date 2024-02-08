package indicator

import (
	ptlog "github.com/aaronchen2k/deeptest/internal/performance/pkg/log"
	ptProto "github.com/aaronchen2k/deeptest/internal/performance/proto"
	client "github.com/influxdata/influxdb1-client/v2"
	"time"
)

const (
	// token             = "CjK5KHeIopceCfRznN7RZxlffNrnCOBJ6Ugi9PCFb-mRu4ZQJ01tqpE4oeWmw5VlaDk-y3JMkKSx8k8Klwh04g=="

	precision         = "ms"
	tableResponseTime = "response_time"

	tableCpuUsage     = "cpu_usage"
	tableMemoryUsage  = "memory_usage"
	tableDiskUsage    = "disk_usage"
	tableNetworkUsage = "network_usage"
)

var (
	InfluxdbInstant *InfluxdbSender
)

type InfluxdbSender struct {
	Client      client.Client
	BatchPoints client.BatchPoints
	DbAddress   string
}

func GetInfluxdbSenderInstant(room string, dbAddress, username, password string) MessageSender {
	if InfluxdbInstant != nil {
		return InfluxdbInstant
	}

	var err error
	influxdbClient, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://localhost:8086",
		Username: username,
		Password: password,
	})
	if err != nil {
		ptlog.Logf("failed to creat influxdb client, err: %s", err.Error())
		return nil
	}

	batchPoints, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  room,
		Precision: "s",
	})
	if err != nil {
		ptlog.Logf("failed to creat influxdb client, err: %s", err.Error())
		return nil
	}

	InfluxdbInstant = &InfluxdbSender{
		Client:      influxdbClient,
		BatchPoints: batchPoints,
		DbAddress:   dbAddress,
	}

	return InfluxdbInstant
}

func (s InfluxdbSender) Send(result ptProto.PerformanceExecResp) (err error) {
	// 1. send responseTime
	for _, request := range result.Requests {
		s.addResponseTimePoint(request, result.Room)
	}

	// OR 2. send metrics
	metrics := result.GetMetrics()
	if metrics != nil {
		s.addCpuUsagePoint(metrics.CpuUsage, result.Room)
		s.addMemoryUsagePoint(metrics.MemoryUsage, result.Room)
		s.addDiskUsagePoint(metrics.DiskUsages, result.Room)
		s.addNetworkUsagePoint(metrics.NetworkUsages, result.Room)
	}

	return
}

func (s InfluxdbSender) addResponseTimePoint(request *ptProto.PerformanceExecRecord, room string) (err error) {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  room,
		Precision: precision,
	})
	if err != nil {
		ptlog.Logf("failed to creat influxdb batchPoints, err: %s", err.Error())
		return
	}

	tags := map[string]string{
		"name": request.RecordName,
	}
	fields := map[string]interface{}{
		"value": request.Duration,
	}

	point, err := client.NewPoint(tableResponseTime, tags, fields, time.Now())
	if err != nil {
		ptlog.Logf("failed to creat influxdb point, err: %s", err.Error())
		return
	}

	bp.AddPoint(point)
	err = s.Client.Write(bp)
	if err != nil {
		ptlog.Logf("failed to write influxdb dta, err: %s", err.Error())
		return
	}

	return
}

func (s InfluxdbSender) addCpuUsagePoint(value float64, room string) (err error) {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  room,
		Precision: precision,
	})
	if err != nil {
		ptlog.Logf("failed to creat influxdb batchPoints, err: %s", err.Error())
		return
	}

	tags := map[string]string{}
	fields := map[string]interface{}{
		"value": value,
	}

	point, err := client.NewPoint(tableCpuUsage, tags, fields, time.Now())
	if err != nil {
		ptlog.Logf("failed to creat influxdb point, err: %s", err.Error())
		return
	}

	bp.AddPoint(point)
	err = s.Client.Write(bp)
	if err != nil {
		ptlog.Logf("failed to write influxdb dta, err: %s", err.Error())
		return
	}

	return
}

func (s InfluxdbSender) addMemoryUsagePoint(value float64, room string) (err error) {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  room,
		Precision: precision,
	})
	if err != nil {
		ptlog.Logf("failed to creat influxdb batchPoints, err: %s", err.Error())
		return
	}

	tags := map[string]string{}
	fields := map[string]interface{}{
		"value": value,
	}

	point, err := client.NewPoint(tableMemoryUsage, tags, fields, time.Now())
	if err != nil {
		ptlog.Logf("failed to creat influxdb point, err: %s", err.Error())
		return
	}

	bp.AddPoint(point)
	err = s.Client.Write(bp)
	if err != nil {
		ptlog.Logf("failed to write influxdb dta, err: %s", err.Error())
		return
	}

	return
}

func (s InfluxdbSender) addDiskUsagePoint(mp map[string]float64, room string) (err error) {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  room,
		Precision: precision,
	})
	if err != nil {
		ptlog.Logf("failed to creat influxdb batchPoints, err: %s", err.Error())
		return
	}

	for name, value := range mp {
		tags := map[string]string{
			"name": name,
		}
		fields := map[string]interface{}{
			"value": value,
		}

		point, err := client.NewPoint(tableDiskUsage, tags, fields, time.Now())
		if err != nil {
			ptlog.Logf("failed to creat influxdb point, err: %s", err.Error())
			continue
		}

		bp.AddPoint(point)
	}

	err = s.Client.Write(bp)
	if err != nil {
		ptlog.Logf("failed to write influxdb dta, err: %s", err.Error())
		return
	}

	return
}

func (s InfluxdbSender) addNetworkUsagePoint(mp map[string]float64, room string) (err error) {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  room,
		Precision: precision,
	})
	if err != nil {
		ptlog.Logf("failed to creat influxdb batchPoints, err: %s", err.Error())
		return
	}

	for name, value := range mp {
		tags := map[string]string{
			"name": name,
		}
		fields := map[string]interface{}{
			"value": value,
		}

		point, err := client.NewPoint(tableNetworkUsage, tags, fields, time.Now())
		if err != nil {
			ptlog.Logf("failed to creat influxdb point, err: %s", err.Error())
			continue
		}

		bp.AddPoint(point)
	}

	err = s.Client.Write(bp)
	if err != nil {
		ptlog.Logf("failed to write influxdb dta, err: %s", err.Error())
		return
	}

	return
}
