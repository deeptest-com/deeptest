package indicator

import (
	"context"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/performance/pkg/log"
	ptproto "github.com/aaronchen2k/deeptest/internal/performance/proto"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/domain"

	"time"
)

var (
	precision = "ms"

	//orgName           = "deeptest"
	bucketName            = "performance"
	bucketNameDownsampled = "performance-downsampled"

	taskResponseTime   = "task_response_time"
	taskResponseTime90 = "task_response_time_90"
	taskResponseTime95 = "task_response_time_95"
	taskQps            = "task_qps"
	taskCpuUsage       = "task_cpu_usage"
	taskMemoryUsage    = "task_memory_usage"
	taskDiskUsage      = "task_disk_usage"
	taskNetworkUsage   = "task_network_usage"

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

	// 删除已有bucket
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// recreate bucket if needed
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

	bucket, err = bucketsAPI.CreateBucketWithName(ctx, org, bucketName, domain.RetentionRule{EverySeconds: 3600 * 24})
	if err != nil {
		ptlog.Logf("failed to create bucketName %s, err %s", bucketName, err.Error())
		return nil
	}
	ptlog.Logf("success to create bucket %s", bucket.Name)

	// tasks
	err = createResponseTimeTask(ctx, influxdbClient, *org.Id)
	if err != nil {
		return nil
	}
	err = createResponseTime90Task(ctx, influxdbClient, *org.Id)
	if err != nil {
		return nil
	}
	err = createResponseTime95Task(ctx, influxdbClient, *org.Id)
	if err != nil {
		return nil
	}
	err = createQpsTask(ctx, influxdbClient, *org.Id)
	if err != nil {
		return nil
	}
	err = createCpuTask(ctx, influxdbClient, *org.Id)
	if err != nil {
		return nil
	}
	err = createMemoryTask(ctx, influxdbClient, *org.Id)
	if err != nil {
		return nil
	}
	err = createDiskTask(ctx, influxdbClient, *org.Id)
	if err != nil {
		return nil
	}
	err = createNetworkTask(ctx, influxdbClient, *org.Id)
	if err != nil {
		return nil
	}

	// queries
	//queryAPI := influxdbClient.QueryAPI("my-org")
	//queryAPI.

	// write
	writeAPI := influxdbClient.WriteAPIBlocking(orgName, bucketName)

	InfluxdbInstant = &InfluxdbSender{
		Client:    influxdbClient,
		WriteAPI:  writeAPI,
		DbAddress: dbAddress,
	}

	return InfluxdbInstant
}

func createResponseTimeTask(ctx context.Context, influxdbClient influxdb2.Client, orgId string) (err error) {
	table := tableResponseTime
	flux := fmt.Sprintf(`
from(bucket: "%s")
    |> range(start: -task.every)
    |> filter(
        fn: (r) =>
            r._measurement == "%s",
    )
    |> aggregateWindow(every: task.every, fn: mean)
    |> to(bucket: "%s")
`, bucketName, table, bucketNameDownsampled)

	createTask(ctx, influxdbClient, orgId, table, taskResponseTime, flux)

	return
}

func createResponseTime90Task(ctx context.Context, influxdbClient influxdb2.Client, orgId string) (err error) {
	name := taskResponseTime90
	table := tableResponseTime
	flux := fmt.Sprintf(`
from(bucket: "%s")
    |> range(start: -task.every)
    |> filter(
        fn: (r) =>
            r._measurement == "%s",
    )
    |> aggregateWindow(
       every: task.every, 
       fn: (t=<-, column) => t 
           |> quantile(q: 0.9, method: "exact_selector"), 
    |> to(bucket: "%s")
`, bucketName, tableResponseTime, bucketNameDownsampled)

	createTask(ctx, influxdbClient, orgId, table, name, flux)

	return
}

func createResponseTime95Task(ctx context.Context, influxdbClient influxdb2.Client, orgId string) (err error) {
	name := taskResponseTime95
	table := tableResponseTime
	flux := fmt.Sprintf(`
from(bucket: "%s")
    |> range(start: -task.every)
    |> filter(
        fn: (r) =>
            r._measurement == "%s",
    )
    |> aggregateWindow(
       every: task.every, 
       fn: (t=<-, column) => t 
           |> quantile(q: 0.9, method: "exact_selector"), 
    |> to(bucket: "%s")
`, bucketName, tableResponseTime, bucketNameDownsampled)

	createTask(ctx, influxdbClient, orgId, table, name, flux)

	return
}

func createQpsTask(ctx context.Context, influxdbClient influxdb2.Client, orgId string) (err error) {
	name := taskQps
	table := tableResponseTime
	flux := fmt.Sprintf(`
from(bucket: "%s")
    |> range(start: -task.every)
    |> filter(
        fn: (r) =>
            r._measurement == "%s",
    )
    |> aggregateWindow(
	   every: task.every, 
	   fn: (table=<-, column) => table 
		 |> count(column: "_value") / task.every, 
	   createEmpty: false)
    |> to(bucket: "%s")
`, bucketName, tableResponseTime, bucketNameDownsampled)

	createTask(ctx, influxdbClient, orgId, table, name, flux)

	return
}

func createCpuTask(ctx context.Context, influxdbClient influxdb2.Client, orgId string) (err error) {
	name := taskCpuUsage
	table := tableCpuUsage
	flux := fmt.Sprintf(`
from(bucket: "%s")
    |> range(start: -task.every)
    |> filter(
        fn: (r) =>
            r._measurement == "%s",
    )
    |> aggregateWindow(
	   every: task.every, 
	   fn: mean, 
	   createEmpty: false)
    |> to(bucket: "%s")
`, bucketName, tableCpuUsage, bucketNameDownsampled)

	createTask(ctx, influxdbClient, orgId, table, name, flux)

	return
}

func createMemoryTask(ctx context.Context, influxdbClient influxdb2.Client, orgId string) (err error) {
	name := taskMemoryUsage
	table := tableMemoryUsage
	flux := fmt.Sprintf(`
from(bucket: "%s")
    |> range(start: -task.every)
    |> filter(
        fn: (r) =>
            r._measurement == "%s",
    )
    |> aggregateWindow(
	   every: task.every, 
	   fn: mean, 
	   createEmpty: false)
    |> to(bucket: "%s")
`, bucketName, tableMemoryUsage, bucketNameDownsampled)

	createTask(ctx, influxdbClient, orgId, table, name, flux)

	return
}

func createDiskTask(ctx context.Context, influxdbClient influxdb2.Client, orgId string) (err error) {
	name := taskDiskUsage
	table := tableDiskUsage
	flux := fmt.Sprintf(`
from(bucket: "%s")
    |> range(start: -task.every)
    |> filter(
        fn: (r) =>
            r._measurement == "%s",
    )
    |> aggregateWindow(
	   every: task.every, 
	   fn: mean, 
	   createEmpty: false)
    |> to(bucket: "%s")
`, bucketName, tableDiskUsage, bucketNameDownsampled)

	createTask(ctx, influxdbClient, orgId, table, name, flux)

	return
}

func createNetworkTask(ctx context.Context, influxdbClient influxdb2.Client, orgId string) (err error) {
	name := taskNetworkUsage
	table := tableNetworkUsage
	flux := fmt.Sprintf(`
from(bucket: "%s")
    |> range(start: -task.every)
    |> filter(
        fn: (r) =>
            r._measurement == "%s",
    )
    |> aggregateWindow(
	   every: task.every, 
	   fn: mean, 
	   createEmpty: false)
    |> to(bucket: "%s")
`, bucketName, table, bucketNameDownsampled)

	createTask(ctx, influxdbClient, orgId, table, name, flux)

	return
}

func createTask(ctx context.Context, influxdbClient influxdb2.Client, orgId string,
	table, name, flux string) {

	tasksAPI := influxdbClient.TasksAPI()
	taskStatus := domain.TaskStatusTypeActive
	taskEvery := "10s"
	taskOffset := "0s"
	taskFlux := fmt.Sprintf(flux, bucketName, table, bucketNameDownsampled)

	newTask := &domain.Task{
		Name:   name,
		Every:  &taskEvery,
		Offset: &taskOffset,
		Flux:   taskFlux,
		OrgID:  orgId,
		Status: &taskStatus,
	}

	_, err := tasksAPI.CreateTask(ctx, newTask)
	if err != nil {
		ptlog.Logf("failed to create task %s, err %s.", name, err.Error())
		return
	}
	ptlog.Logf("success to create bucketName %s", name)

	return
}

func (s InfluxdbSender) Send(result ptproto.PerformanceExecResp) (err error) {
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

func (s InfluxdbSender) addResponseTimePoint(request *ptproto.PerformanceExecRecord, room string, lines *[]string) (
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
