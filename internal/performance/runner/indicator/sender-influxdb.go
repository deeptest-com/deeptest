package indicator

import (
	"context"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/performance/pkg/log"
	ptproto "github.com/aaronchen2k/deeptest/internal/performance/proto"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"time"
)

var (
	//orgName           = "deeptest"
	bucketName = "performance"

	tableVuNumb       = "vu_numb"
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
	Client    influxdb2.Client
	WriteAPI  api.WriteAPIBlocking
	DbAddress string
}

func GetInfluxdbSenderInstant(room, dbAddress, orgName, token string) MessageSender {
	if InfluxdbInstant != nil {
		return InfluxdbInstant
	}

	influxdbClient := influxdb2.NewClient(dbAddress, token)

	// write
	writeAPI := influxdbClient.WriteAPIBlocking(orgName, bucketName)

	InfluxdbInstant = &InfluxdbSender{
		Client:    influxdbClient,
		WriteAPI:  writeAPI,
		DbAddress: dbAddress,
	}

	return InfluxdbInstant
}

func (s InfluxdbSender) Send(result ptproto.PerformanceExecResp) (err error) {
	var lines []string
	var points []*write.Point

	// 1. send responseTime
	if len(result.Requests) > 0 {
		for _, request := range result.Requests {
			s.addResponseTimePoint(request, result.Room, &points)
		}
	}

	// OR 2. send metrics
	metrics := result.Metrics
	if metrics != nil {
		s.addCpuUsagePoint(metrics.CpuUsage, result.Room, result.RunnerName, &lines)
		s.addMemoryUsagePoint(metrics.MemoryUsage, result.Room, result.RunnerName, &lines)
		s.addDiskUsagePoint(metrics.DiskUsages, result.Room, result.RunnerName, &lines)
		s.addNetworkUsagePoint(metrics.NetworkUsages, result.Room, result.RunnerName, &lines)
	}

	// OR 3. send metrics
	if result.VuCount > 0 {
		s.addVuCount(result.VuCount, result.Room, result.RunnerName, &lines)
	}

	if len(lines) > 0 {
		err = s.WriteAPI.WriteRecord(context.Background(), lines...)
		if err != nil {
			ptlog.Logf("failed to write influxdb lines data, err: %s", err.Error())
			return
		}
	}

	if len(points) > 0 {
		err = s.WriteAPI.WritePoint(context.Background(), points...)
		if err != nil {
			ptlog.Logf("failed to write influxdb points data, err: %s", err.Error())
			return
		}
	}

	return
}

func (s InfluxdbSender) addResponseTimePoint(request *ptproto.PerformanceExecRecord, room string, points *[]*write.Point) (
	err error) {

	p := influxdb2.NewPoint(tableResponseTime,
		map[string]string{"name": request.RecordName},
		map[string]interface{}{
			"value":  request.Duration,
			"status": request.Status,
			"start":  request.StartTime,
			"end":    request.EndTime,
		},
		time.Now(),
	)

	*points = append(*points, p)

	return
}

func (s InfluxdbSender) addCpuUsagePoint(value float64, room string, runnerName string, lines *[]string) (err error) {
	line := fmt.Sprintf("%s,runner=%s value=%f", tableCpuUsage, runnerName, value)

	*lines = append(*lines, line)

	return
}

func (s InfluxdbSender) addMemoryUsagePoint(value float64, room string, runnerName string, lines *[]string) (err error) {
	line := fmt.Sprintf("%s,runner=%s value=%f", tableMemoryUsage, runnerName, value)

	*lines = append(*lines, line)

	return
}

func (s InfluxdbSender) addDiskUsagePoint(mp map[string]float64, room string, runnerName string, lines *[]string) (err error) {
	for name, value := range mp {
		line := fmt.Sprintf("%s,runner=%s,name=%s value=%f", tableDiskUsage, runnerName, name, value)
		*lines = append(*lines, line)
	}

	return
}

func (s InfluxdbSender) addNetworkUsagePoint(mp map[string]float64, room string, runnerName string, lines *[]string) (err error) {
	for name, value := range mp {
		line := fmt.Sprintf("%s,runner=%s,name=%s value=%f", tableNetworkUsage, runnerName, name, value)
		*lines = append(*lines, line)
	}

	return
}

func (s InfluxdbSender) addVuCount(count int32, room, runnerName string, lines *[]string) (err error) {
	line := fmt.Sprintf("%s,runner=%s value=%d", tableVuNumb, runnerName, count)
	*lines = append(*lines, line)

	return
}
