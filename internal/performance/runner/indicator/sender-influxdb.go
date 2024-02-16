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
	//orgName           = "deeptest"
	bucketName            = "performance"
	bucketNameDownsampled = "performance-downsampled"

	taskStatus    = domain.TaskStatusTypeActive
	taskEveryNumb = 10
	taskEvery     = fmt.Sprintf("%ds", taskEveryNumb)
	taskOffset    = "0s"

	taskVuNumb         = "task_vu_numb"
	taskFailNumb       = "task_fail_numb"
	taskResponseTime   = "task_response_time"
	taskResponseTime90 = "task_response_time_90"
	taskResponseTime95 = "task_response_time_95"
	taskQps            = "task_qps"
	taskCpuUsage       = "task_cpu_usage"
	taskMemoryUsage    = "task_memory_usage"
	taskDiskUsage      = "task_disk_usage"
	taskNetworkUsage   = "task_network_usage"

	tableVuCount      = "vu_count"
	tableResponseTime = "response_time"
	tableQps          = "qps"
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

func ResetInfluxdb(room, dbAddress, orgName, token string) {
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
			ptlog.Logf("failed to delete bucket %s, err %s", bucketName, err.Error())
			return
		}
		ptlog.Logf("success to delete bucket %s", bucketName)
	}

	org, err2 := influxdbClient.OrganizationsAPI().FindOrganizationByName(ctx, orgName)
	if err2 != nil {
		org, err2 = influxdbClient.OrganizationsAPI().CreateOrganizationWithName(ctx, orgName)
		if err2 != nil {
			ptlog.Logf("failed to create org by name %s, err %s", orgName, err.Error())
			return
		}
		ptlog.Logf("success to create org %s", bucketName)
	}

	bucket, err = bucketsAPI.CreateBucketWithName(ctx, org, bucketName, domain.RetentionRule{EverySeconds: 3600 * 24})
	if err != nil {
		ptlog.Logf("failed to create bucketName %s, err %s", bucketName, err.Error())
		return
	}
	ptlog.Logf("success to create bucket %s", bucket.Name)

	// tasks
	err = createVuNumbTask(ctx, influxdbClient, *org.Id)
	if err != nil {
		return
	}
	err = createFailNumbTask(ctx, influxdbClient, *org.Id)
	if err != nil {
		return
	}
	err = createResponseTimeAllTask(ctx, influxdbClient, *org.Id)
	if err != nil {
		return
	}
	err = createResponseTime90Task(ctx, influxdbClient, *org.Id)
	if err != nil {
		return
	}
	err = createResponseTime95Task(ctx, influxdbClient, *org.Id)
	if err != nil {
		return
	}
	err = createQpsTask(ctx, influxdbClient, *org.Id)
	if err != nil {
		return
	}
	err = createCpuTask(ctx, influxdbClient, *org.Id)
	if err != nil {
		return
	}
	err = createMemoryTask(ctx, influxdbClient, *org.Id)
	if err != nil {
		return
	}
	err = createDiskTask(ctx, influxdbClient, *org.Id)
	if err != nil {
		return
	}
	err = createNetworkTask(ctx, influxdbClient, *org.Id)
	if err != nil {
		return
	}

	return
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

func QueryResponseTimeSummary(ctx context.Context, influxdbClient influxdb2.Client, orgId string) (err error) {
	sql := fmt.Sprintf(`
baseData = from(bucket: "%s")
  |> range(start: -1d)
  |> filter(
      fn: (r) =>
          r._measurement == "%s",
  )

minData = baseData
  |> min()
  |> rename(columns: {"_value":"min"})

maxData = baseData
  |> max()
  |> rename(columns: {"_value":"max"})

meanData = baseData
  |> mean()
  |> rename(columns: {"_value":"mean"})

medianData = baseData
  |> median()
  |> rename(columns: {"_value":"median"})

union(tables: [minData, maxData, meanData, medianData])
  |> keep(columns: ["name", "min", "max", "mean", "median"])
`, bucketName, tableResponseTime)

	err = queryData(influxdbClient, orgId, sql)

	return
}

func createVuNumbTask(ctx context.Context, influxdbClient influxdb2.Client, orgId string) (err error) {
	name := taskVuNumb
	flux := fmt.Sprintf(`
from(bucket: "%s")
    |> filter(
        fn: (r) =>
            r._measurement == "%s",
    )
    |> aggregateWindow(
	   every: task.every, 
	   fn: sum, 
	   createEmpty: false)
    |> to(bucket: "%s")
`, bucketName, tableVuCount, bucketNameDownsampled)

	err = createTask(ctx, influxdbClient, orgId, name, flux)

	return
}

func createFailNumbTask(ctx context.Context, influxdbClient influxdb2.Client, orgId string) (err error) {
	name := taskNetworkUsage
	flux := fmt.Sprintf(`
from(bucket: "%s")
    |> filter(
        fn: (r) =>
            r._measurement == "%s" and (r.status == "fail" or r.status == "err"),
    )
    |> aggregateWindow(
	   every: task.every, 
	   fn: sum, 
	   createEmpty: false)
    |> to(bucket: "%s")
`, bucketName, tableNetworkUsage, bucketNameDownsampled)

	err = createTask(ctx, influxdbClient, orgId, name, flux)

	return
}

func createResponseTimeAllTask(ctx context.Context, influxdbClient influxdb2.Client, orgId string) (err error) {
	flux := fmt.Sprintf(`
from(bucket: "%s")
    |> range(start: -task.every)
    |> filter(
        fn: (r) =>
            r._measurement == "%s",
    )
    |> aggregateWindow(every: task.every, fn: mean)
	|> set(key: "_measurement", value: "response_time_90")
    |> to(bucket: "%s")
`, bucketName, tableResponseTime, bucketNameDownsampled)

	err = createTask(ctx, influxdbClient, orgId, taskResponseTime, flux)

	return
}

func createResponseTime90Task(ctx context.Context, influxdbClient influxdb2.Client, orgId string) (err error) {
	name := taskResponseTime90
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
            |> quantile(q: 0.9, method: "exact_selector"),
    )
	|> set(key: "_measurement", value: "%s")
    |> to(bucket: "%s")
`, bucketName, tableResponseTime, tableResponseTime+"_90", bucketNameDownsampled)

	err = createTask(ctx, influxdbClient, orgId, name, flux)

	return
}

func createResponseTime95Task(ctx context.Context, influxdbClient influxdb2.Client, orgId string) (err error) {
	name := taskResponseTime95
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
            |> quantile(q: 0.95, method: "exact_selector"),
    )
	|> set(key: "_measurement", value: "%s")
    |> to(bucket: "%s")
`, bucketName, tableResponseTime, tableResponseTime+"_95", bucketNameDownsampled)

	err = createTask(ctx, influxdbClient, orgId, name, flux)

	return
}

func createQpsTask(ctx context.Context, influxdbClient influxdb2.Client, orgId string) (err error) {
	name := taskQps
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
		 |> count(column: "_value") / %d, 
	   createEmpty: false)
    |> to(bucket: "%s")
`, bucketName, tableQps, taskEveryNumb, bucketNameDownsampled)

	err = createTask(ctx, influxdbClient, orgId, name, flux)

	return
}

func createCpuTask(ctx context.Context, influxdbClient influxdb2.Client, orgId string) (err error) {
	name := taskCpuUsage
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

	err = createTask(ctx, influxdbClient, orgId, name, flux)

	return
}

func createMemoryTask(ctx context.Context, influxdbClient influxdb2.Client, orgId string) (err error) {
	name := taskMemoryUsage
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

	err = createTask(ctx, influxdbClient, orgId, name, flux)

	return
}

func createDiskTask(ctx context.Context, influxdbClient influxdb2.Client, orgId string) (err error) {
	name := taskDiskUsage
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

	err = createTask(ctx, influxdbClient, orgId, name, flux)

	return
}

func createNetworkTask(ctx context.Context, influxdbClient influxdb2.Client, orgId string) (err error) {
	name := taskNetworkUsage
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
`, bucketName, tableNetworkUsage, bucketNameDownsampled)

	err = createTask(ctx, influxdbClient, orgId, name, flux)

	return
}

func queryData(influxdbClient influxdb2.Client, orgId string, query string) (err error) {
	queryAPI := influxdbClient.QueryAPI(orgId)
	result, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		ptlog.Logf("failed to query, err %s.", err.Error())
		return
	}

	for result.Next() {
		mp := result.Record().Values()
		name := mp["name"]
		min := mp["min"]
		max := mp["max"]
		mean := mp["mean"]
		median := mp["median"]

		ptlog.Logf("%s, %v, %v, %v, %v", name, min, max, mean, median)
	}

	return
}

func createTask(ctx context.Context, influxdbClient influxdb2.Client, orgId string,
	name, flux string) (err error) {

	tasksAPI := influxdbClient.TasksAPI()

	tasks, err := tasksAPI.FindTasks(ctx, &api.TaskFilter{Name: name})
	for _, task := range tasks {
		err = tasksAPI.DeleteTaskWithID(ctx, task.Id)
		if err != nil {
			ptlog.Logf("failed to delete task %s, err %s.", task.Id, err.Error())
			continue
		}
		ptlog.Logf("success to create task %s", task.Id)
	}

	newTask := &domain.Task{
		Name:   name,
		Every:  &taskEvery,
		Offset: &taskOffset,
		Flux:   flux,
		OrgID:  orgId,
		Status: &taskStatus,
	}

	_, err = tasksAPI.CreateTask(ctx, newTask)
	if err != nil {
		ptlog.Logf("failed to create task %s, err %s.", name, err.Error())
		return
	}
	ptlog.Logf("success to create task %s", name)

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

	// OR 3. send metrics
	if result.VuCount > 0 {
		s.addVuCount(result.VuCount, result.Room, &lines)
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
	line := fmt.Sprintf("%s,name=%s,status=%s value=%d",
		tableResponseTime, request.RecordName, request.Status, request.Duration)

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

func (s InfluxdbSender) addVuCount(count int32, room string, lines *[]string) (err error) {
	line := fmt.Sprintf("%s value=%d", tableVuCount, count)
	*lines = append(*lines, line)

	return
}
