package dao

import (
	"context"
	"fmt"
	ptdomain "github.com/aaronchen2k/deeptest/internal/performance/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/performance/pkg/log"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/domain"

	"time"
)

var (
	//orgName           = "deeptest"
	bucketName            = "performance"
	bucketNameDownsampled = "performance-downsampled"

	tableVuNumb       = "vu_numb"
	tableFailNumb     = "fail_numb"
	tableResponseTime = "response_time"
	tableQps          = "qps"
	tableCpuUsage     = "cpu_usage"
	tableMemoryUsage  = "memory_usage"
	tableDiskUsage    = "disk_usage"
	tableNetworkUsage = "network_usage"

	taskStatus    = domain.TaskStatusTypeActive
	taskEveryNumb = 10
	taskEvery     = fmt.Sprintf("%ds", taskEveryNumb)
	taskOffset    = "0s"
)

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

	return
}

func QueryResponseTimeSummary(ctx context.Context, influxdbClient influxdb2.Client, orgId string) (
	ret ptdomain.PerformanceExecSummary, err error) {
	flux1 := fmt.Sprintf(`
baseDataResponse =
    from(bucket: "%s")
        |> range(start: -1m)
        |> filter(fn: (r) => r._measurement == "%s" and r._field == "value")
        |> group()

minVal =
   baseDataResponse
       |> min()
       |> toFloat()
       |> set(key: "_field", value: "minVal")
maxVal =
   baseDataResponse
       |> max()
       |> toFloat()
       |> set(key: "_field", value: "maxVal")
meanVal =
   baseDataResponse
       |> median()
       |> set(key: "_field", value: "meanVal")
medianVal =
   baseDataResponse
       |> median()
       |> set(key: "_field", value: "medianVal")

union(tables: [minVal, maxVal, meanVal, medianVal])

`, bucketName, tableResponseTime)

	flux2 := fmt.Sprintf(`
baseDataOthers =
    from(bucket: "%s")
        |> range(start: -1m)
        |> filter(fn: (r) => r._measurement == "%s" and r._field != "value")

startTime =
    baseDataOthers
        |> filter(fn: (r) => r["_field"] == "start")
        |> lowestMin(n: 1, column: "_value")
        |> set(key: "_field", value: "startTime")
endTime =
    baseDataOthers
        |> filter(fn: (r) => r["_field"] == "end")
        |> highestMax(n: 1, column: "_value")
        |> set(key: "_field", value: "endTime")

union(tables: [startTime, endTime])

`, bucketName, tableResponseTime)

	flux3 := fmt.Sprintf(`
baseDataResponse =
    from(bucket: "%s")
        |> range(start: -1m)
        |> filter(fn: (r) => r._measurement == "%s" and r._field == "status")

passNumb =
   baseDataResponse
       |> filter(fn: (r) => r["_value"] == "pass")
       |> count()
       |> set(key: "_field", value: "passNumb")
failNumb =
   baseDataResponse
       |> filter(fn: (r) => r["_value"] == "fail")
       |> count()
       |> set(key: "_field", value: "failNumb")
errNumb =
   baseDataResponse
       |> filter(fn: (r) => r["_value"] == "error")
       |> count()
       |> set(key: "_field", value: "errNumb")

union(tables: [passNumb, failNumb, errNumb])

`, bucketName, tableResponseTime)

	result1, err1 := queryData(influxdbClient, orgId, flux1)
	if err1 != nil {
		ptlog.Logf("query data failed, err %s", err1.Error())
		return
	}

	for result1.Next() {
		mp := result1.Record().Values()

		typ := mp["_field"].(string)

		if typ == "minVal" {
			val := mp["_value"].(float64)
			ret.MinResponseTime = val
		} else if typ == "maxVal" {
			val := mp["_value"].(float64)
			ret.MaxResponseTime = val
		} else if typ == "meanVal" {
			val := mp["_value"].(float64)
			ret.MeanResponseTime = val
		} else if typ == "medianVal" {
			val := mp["_value"].(float64)
			ret.MedianResponseTime = val
		}
	}

	result2, err2 := queryData(influxdbClient, orgId, flux2)
	if err2 != nil {
		ptlog.Logf("query data failed, err %s", err2.Error())
		return
	}

	for result2.Next() {
		mp := result2.Record().Values()

		typ := mp["_field"].(string)

		if typ == "startTime" {
			startTime := mp["_value"].(int64)
			ret.StartTime = startTime
		} else if typ == "endTime" {
			endTime := mp["_value"].(int64)
			ret.EndTime = endTime
		}
	}

	result3, err3 := queryData(influxdbClient, orgId, flux3)
	if err3 != nil {
		ptlog.Logf("query data failed, err %s", err3.Error())
		return
	}

	for result3.Next() {
		mp := result3.Record().Values()

		typ := mp["_field"].(string)

		if typ == "passNumb" {
			val := mp["_value"].(int64)
			ret.Pass = int(val)
		} else if typ == "failNumb" {
			val := mp["_value"].(int64)
			ret.Fail = int(val)
		} else if typ == "errNumb" {
			val := mp["_value"].(int64)
			ret.Error = int(val)
		}
	}

	ret.Total = ret.Pass + ret.Pass + ret.Error
	ret.Duration = ret.EndTime - ret.StartTime
	ret.AvgQps = float64(ret.Total) * 1000 / float64(ret.Duration)

	return
}

func QueryVuCount(ctx context.Context, influxdbClient influxdb2.Client, orgId string) (
	ret int, err error) {

	flux := fmt.Sprintf(`
  from(bucket: "%s")
	|> range(start: -1d)
	|> filter(fn: (r) => r._measurement == "%s")
       |> sum()
       |> set(key: "_field", value: "vmNumb")

`, bucketName, tableVuNumb)

	result, err := queryData(influxdbClient, orgId, flux)
	if err != nil {
		ptlog.Logf("query data failed, err %s", err.Error())
		return
	}

	if result.Next() {
		mp := result.Record().Values()

		ret = int(mp["_value"].(float64))
	}

	return
}

func queryResponseTimeTableByInterface(ctx context.Context, influxdbClient influxdb2.Client, orgId string) (
	ret []ptdomain.PerformanceRequestTable, err error) {
	flux := fmt.Sprintf(`
baseData =
    from(bucket: "%s")
        |> filter(fn: (r) => r._measurement == "%s" and r["_field"] == "_value")

minData =
    baseData
        |> min()
        |> set(key: "_field", value: "min")

maxData =
    baseData
        |> max()
        |> set(key: "_field", value: "max")

meanData =
    baseData
        |> mean()
        |> set(key: "_field", value: "mean")

medianData =
    baseData
        |> median()
        |> set(key: "_field", value: "median")

union(tables: [minData, maxData, meanData, medianData])
`, bucketName, tableResponseTime)

	result, err := queryData(influxdbClient, orgId, flux)

	for result.Next() {
		mp := result.Record().Values()

		item := ptdomain.PerformanceRequestTable{
			RecordName: mp["name"].(string),
			Type:       mp["_field"].(string),
			Value:      int32(mp["_value"].(float64)),
		}
		ret = append(ret, item)
	}

	return
}

func queryResponseTimeAvgAll(ctx context.Context, influxdbClient influxdb2.Client, orgId string) (
	ret []ptdomain.PerformanceRequestResponseTime, err error) {
	flux := fmt.Sprintf(`
from(bucket: "%s")
    |> range(start: -1m)
    |> filter(
        fn: (r) =>
            r._measurement == "%s" and r["_field"] == "_value",
    )
    |> aggregateWindow(every: 1m, fn: mean)
	|> set(key: "_measurement", value: "%s")
`, bucketName, tableResponseTime, bucketNameDownsampled)

	result, err := queryData(influxdbClient, orgId, flux)

	for result.Next() {
		mp := result.Record().Values()

		item := ptdomain.PerformanceRequestResponseTime{
			RecordName: mp["name"].(string),
			Value:      int32(mp["_value"].(float64)),
		}
		ret = append(ret, item)
	}

	return
}

func queryResponseTimeAvg90(ctx context.Context, influxdbClient influxdb2.Client, orgId string) (
	ret []ptdomain.PerformanceRequestResponseTime, err error) {
	flux := fmt.Sprintf(`
from(bucket: "%s")
    |> range(start: -1m)
    |> filter(
        fn: (r) =>
            r._measurement == "%s" and r["_field"] == "_value",
    )
    |> aggregateWindow(
        every: 1m, 
        fn: (table=<-, column) => table 
            |> quantile(q: 0.9, method: "exact_selector"),
    )
	|> set(key: "_measurement", value: "%s")
`, bucketName, tableResponseTime, bucketNameDownsampled)

	result, err := queryData(influxdbClient, orgId, flux)

	for result.Next() {
		mp := result.Record().Values()

		item := ptdomain.PerformanceRequestResponseTime{
			RecordName: mp["name"].(string),
			Value:      int32(mp["_value"].(float64)),
		}
		ret = append(ret, item)
	}

	return
}

func queryResponseTimeAvg95(ctx context.Context, influxdbClient influxdb2.Client, orgId string) (
	ret []ptdomain.PerformanceRequestResponseTime, err error) {
	flux := fmt.Sprintf(`
from(bucket: "%s")
    |> range(start: -1m)
    |> filter(
        fn: (r) =>
            r._measurement == "%s" and r["_field"] == "_value",
    )
    |> aggregateWindow(
        every: task.every, 
        fn: (table=<-, column) => table 
            |> quantile(q: 0.95, method: "exact_selector"),
    )
	|> set(key: "_measurement", value: "%s")
`, bucketName, tableResponseTime, bucketNameDownsampled)

	result, err := queryData(influxdbClient, orgId, flux)

	for result.Next() {
		mp := result.Record().Values()

		item := ptdomain.PerformanceRequestResponseTime{
			RecordName: mp["name"].(string),
			Value:      int32(mp["_value"].(float64)),
		}
		ret = append(ret, item)
	}

	return
}

func queryQps(ctx context.Context, influxdbClient influxdb2.Client, orgId string) (
	ret []ptdomain.PerformanceRequestQps, err error) {
	flux := fmt.Sprintf(`
from(bucket: "%s")
    |> range(start: -1m)
    |> filter(
        fn: (r) =>
            r._measurement == "%s" and r["_field"] == "_value",
    )
    |> aggregateWindow(
	   every: 1m, 
	   fn: (table=<-, column) => table 
		 |> count(column: "_value") / %d, 
	   createEmpty: false)
`, bucketName, tableQps, bucketNameDownsampled)

	result, err := queryData(influxdbClient, orgId, flux)

	for result.Next() {
		mp := result.Record().Values()

		item := ptdomain.PerformanceRequestQps{
			RecordName: mp["name"].(string),
			Value:      mp["_value"].(float64),
		}
		ret = append(ret, item)
	}

	return
}

func queryCpu(ctx context.Context, influxdbClient influxdb2.Client, orgId string) (
	ret float64, err error) {
	flux := fmt.Sprintf(`
from(bucket: "%s")
    |> range(start: -10s)
    |> filter(
        fn: (r) =>
            r._measurement == "%s",
    )
    |> aggregateWindow(
	   every: 10s, 
	   fn: mean, 
	   createEmpty: false)
`, bucketName, tableCpuUsage)

	result, err := queryData(influxdbClient, orgId, flux)

	result.Next()
	mp := result.Record().Values()
	ret = mp["_value"].(float64)

	return
}

func queryMemory(ctx context.Context, influxdbClient influxdb2.Client, orgId string) (
	ret float64, err error) {
	flux := fmt.Sprintf(`
from(bucket: "%s")
    |> range(start: -10s)
    |> filter(
        fn: (r) =>
            r._measurement == "%s",
    )
    |> aggregateWindow(
	   every: 10s, 
	   fn: mean, 
	   createEmpty: false)
`, bucketName, tableMemoryUsage)

	result, err := queryData(influxdbClient, orgId, flux)

	result.Next()
	mp := result.Record().Values()
	ret = mp["_value"].(float64)

	return
}

func queryDisk(ctx context.Context, influxdbClient influxdb2.Client, orgId string) (
	ret map[string]float64, err error) {
	flux := fmt.Sprintf(`
from(bucket: "%s")
    |> range(start: -10s)
    |> filter(
        fn: (r) =>
            r._measurement == "%s",
    )
    |> aggregateWindow(
	   every: 10s, 
	   fn: mean, 
	   createEmpty: false)
`, bucketName, tableDiskUsage)

	result, err := queryData(influxdbClient, orgId, flux)

	for result.Next() {
		mp := result.Record().Values()

		ret[mp["name"].(string)] = mp["_value"].(float64)
	}

	return
}

func queryNetwork(ctx context.Context, influxdbClient influxdb2.Client, orgId string) (
	ret map[string]float64, err error) {
	flux := fmt.Sprintf(`
from(bucket: "%s")
    |> range(start: -10s)
    |> filter(
        fn: (r) =>
            r._measurement == "%s",
    )
    |> aggregateWindow(
	   every: 10s, 
	   fn: mean, 
	   createEmpty: false)
`, bucketName, tableNetworkUsage)

	result, err := queryData(influxdbClient, orgId, flux)

	for result.Next() {
		mp := result.Record().Values()

		ret[mp["name"].(string)] = mp["_value"].(float64)
	}

	return
}

func queryData(influxdbClient influxdb2.Client, orgId string, query string) (result *api.QueryTableResult, err error) {
	queryAPI := influxdbClient.QueryAPI(orgId)

	result, err = queryAPI.Query(context.Background(), query)
	if err != nil {
		ptlog.Logf("failed to query, err %s.", err.Error())
		return
	}

	return
}
