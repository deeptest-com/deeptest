package indicator

import (
	"context"
	"fmt"
	ptlog "github.com/aaronchen2k/deeptest/internal/performance/pkg/log"
	ptProto "github.com/aaronchen2k/deeptest/internal/performance/proto"
	influxdb2 "github.com/influxdata/influxdb-client-go"
	"github.com/influxdata/influxdb-client-go/api"
	"github.com/influxdata/influxdb-client-go/domain"
	"time"
)

var (
	precision         = "ms"
	tableResponseTime = "response_time"

	//orgName           = "deeptest"
	bucketName        = "performance"
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

	// 删除已有bucket
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	bucketsAPI := influxdbClient.BucketsAPI()

	bucket, err := bucketsAPI.FindBucketByName(ctx, bucketName)
	if err == nil {
		err = influxdbClient.BucketsAPI().DeleteBucket(ctx, bucket)
		if err != nil {
			ptlog.Logf("failed to delete bucketName %s, err %s", bucketName, err.Error())
			return nil
		}
		ptlog.Logf("success to delete bucketName %s", bucketName)
	}

	org, err2 := influxdbClient.OrganizationsAPI().FindOrganizationByName(ctx, orgName)
	if err2 != nil {
		org, err2 = influxdbClient.OrganizationsAPI().CreateOrganizationWithName(ctx, orgName)
		if err2 != nil {
			ptlog.Logf("failed to create org by name %s, err %s", orgName, err.Error())
			return nil
		}
		ptlog.Logf("success to create org %s", bucketName)
	}

	// Create  a bucket with 1 day retention policy
	bucket, err = bucketsAPI.CreateBucketWithName(ctx, org, bucketName, domain.RetentionRule{EverySeconds: 3600 * 24})
	if err != nil {
		ptlog.Logf("failed to create bucketName %s, err %s", bucketName, err.Error())
		return nil
	}

	ptlog.Logf("success to create bucketName %s", bucket.Name)

	writeAPI := influxdbClient.WriteAPIBlocking(orgName, bucketName)

	InfluxdbInstant = &InfluxdbSender{
		Client:    influxdbClient,
		WriteAPI:  writeAPI,
		DbAddress: dbAddress,
	}

	return InfluxdbInstant
}

func (s InfluxdbSender) Send(result ptProto.PerformanceExecResp) (err error) {
	var lines []string

	// 1. send responseTime
	if len(result.Requests) > 0 {
		for _, request := range result.Requests {
			s.addResponseTimePoint(request, result.Room, &lines)
		}
	}

	// OR 2. send metrics
	metrics := result.GetMetrics()
	if metrics != nil {
		s.addCpuUsagePoint(metrics.CpuUsage, result.Room, &lines)
		s.addMemoryUsagePoint(metrics.MemoryUsage, result.Room, &lines)
		s.addDiskUsagePoint(metrics.DiskUsages, result.Room, &lines)
		s.addNetworkUsagePoint(metrics.NetworkUsages, result.Room, &lines)
	}

	if len(lines) > 0 {
		err = s.WriteAPI.WriteRecord(context.Background(), lines...)
		if err != nil {
			ptlog.Logf("failed to write influxdb dta, err: %s", err.Error())
			return
		}
	}

	return
}

func (s InfluxdbSender) addResponseTimePoint(request *ptProto.PerformanceExecRecord, room string, lines *[]string) (
	err error) {
	line := fmt.Sprintf("%s,name=%s value=%d", tableResponseTime, request.RecordName, request.Duration)

	*lines = append(*lines, line)

	return
}

func (s InfluxdbSender) addCpuUsagePoint(value float64, room string, lines *[]string) (err error) {
	line := fmt.Sprintf("%s value=%f", tableCpuUsage, value)

	*lines = append(*lines, line)

	return
}

func (s InfluxdbSender) addMemoryUsagePoint(value float64, room string, lines *[]string) (err error) {
	line := fmt.Sprintf("%s value=%f", tableMemoryUsage, value)

	*lines = append(*lines, line)

	return
}

func (s InfluxdbSender) addDiskUsagePoint(mp map[string]float64, room string, lines *[]string) (err error) {
	for name, value := range mp {
		line := fmt.Sprintf("%s,name=%s value=%f", tableDiskUsage, name, value)
		*lines = append(*lines, line)
	}

	return
}

func (s InfluxdbSender) addNetworkUsagePoint(mp map[string]float64, room string, lines *[]string) (err error) {
	for name, value := range mp {
		line := fmt.Sprintf("%s,name=%s value=%f", tableNetworkUsage, name, value)
		*lines = append(*lines, line)
	}

	return
}
