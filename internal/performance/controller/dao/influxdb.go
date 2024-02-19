package dao

import (
	"context"
	"fmt"
	ptdomain "github.com/aaronchen2k/deeptest/internal/performance/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/performance/pkg/log"
	_floatUtils "github.com/aaronchen2k/deeptest/pkg/lib/float"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/domain"
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

func QueryResponseTimeSummary(influxdbClient influxdb2.Client, orgId string) (
	ret ptdomain.PerformanceExecSummary, err error) {

	flux1 := fmt.Sprintf(`
baseDataResponse =
    from(bucket: "%s")
        |> range(start: -1d)
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

quantile95Val =
   baseDataResponse
       |> quantile(q: 0.95, method: "exact_selector")
       |> toFloat()
       |> set(key: "_field", value: "quantile95Val")

union(tables: [minVal, maxVal, meanVal, medianVal, quantile95Val])

`, bucketName, tableResponseTime)

	flux2 := fmt.Sprintf(`
baseDataOthers =
    from(bucket: "%s")
        |> range(start: -1d)
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
        |> range(start: -1d)
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
			ret.Min = val
		} else if typ == "maxVal" {
			val := mp["_value"].(float64)
			ret.Max = val
		} else if typ == "meanVal" {
			val := mp["_value"].(float64)
			ret.Mean = val
		} else if typ == "medianVal" {
			val := mp["_value"].(float64)
			ret.Median = val
		} else if typ == "quantile95Val" {
			val := mp["_value"].(float64)
			ret.Quantile95 = val
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
	ret.Qps = float64(ret.Total) * 1000 / float64(ret.Duration)

	return
}

func QueryVuCount(influxdbClient influxdb2.Client, orgId string) (
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

func QueryLastAvgResponseTime(influxdbClient influxdb2.Client, orgId string) (
	ret []ptdomain.PerformanceRequestResponseTime, err error) {
	flux := fmt.Sprintf(`
from(bucket: "%s")
    |> range(start: -1m)
    |> filter(
        fn: (r) =>
            r._measurement == "%s" and r["_field"] == "value",
    )
    |> mean()
`, bucketName, tableResponseTime)

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

func QueryLastQps(influxdbClient influxdb2.Client, orgId string) (
	ret []ptdomain.PerformanceRequestQps, err error) {
	flux := fmt.Sprintf(`
from(bucket: "%s")
    |> range(start: -1m)
    |> filter(
        fn: (r) =>
            r._measurement == "%s" and r["_field"] == "value",
    )
    |> count()
`, bucketName, tableResponseTime)

	result, err := queryData(influxdbClient, orgId, flux)

	for result.Next() {
		mp := result.Record().Values()

		item := ptdomain.PerformanceRequestQps{
			RecordName: mp["name"].(string),
			Value:      _floatUtils.PointNumb(float64(mp["_value"].(int64))/60, 2),
		}
		ret = append(ret, item)
	}

	return
}

func QueryResponseTimeTableByInterface(influxdbClient influxdb2.Client, orgId string) (
	ret []ptdomain.PerformanceRequestTable, err error) {
	flux := fmt.Sprintf(`
baseData =
    from(bucket: "%s")
    	|> range(start: -1d)
        |> filter(fn: (r) => r._measurement == "%s" and r["_field"] == "value")

totalData =
    baseData
        |> count()
        |> set(key: "_field", value: "total")

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

quantile95Val =
   baseData
       |> quantile(q: 0.95, method: "exact_selector")
       |> toFloat()
       |> set(key: "_field", value: "quantile95Val")

union(tables: [totalData, minData, maxData, meanData, medianData, quantile95Val])
`, bucketName, tableResponseTime)

	result, err := queryData(influxdbClient, orgId, flux)

	tableMap := map[string]*ptdomain.PerformanceRequestTable{}

	for result.Next() {
		mp := result.Record().Values()

		name := mp["name"].(string)
		typ := mp["_field"].(string)

		val, ok := tableMap[name]
		if !ok {
			val = &ptdomain.PerformanceRequestTable{
				RecordName: name,
			}
			tableMap[name] = val
		}

		if typ == "total" {
			val.Total = int32(mp["_value"].(int64))
		} else if typ == "min" {
			val.Min = int32(mp["_value"].(int64))
		} else if typ == "max" {
			val.Max = int32(mp["_value"].(int64))
		} else if typ == "mean" {
			val.Mean = mp["_value"].(float64)
		} else if typ == "median" {
			val.Median = mp["_value"].(float64)
		} else if typ == "quantile95Val" {
			val.Quantile95 = mp["_value"].(float64)
		}
	}

	for _, val := range tableMap {
		ret = append(ret, *val)
	}

	return
}

func QueryMetrics(influxdbClient influxdb2.Client, orgId string) (ret []ptdomain.PerformanceExecMetrics, err error) {
	cpuData, _ := queryCpu(influxdbClient, orgId)
	memoryData, _ := queryMemory(influxdbClient, orgId)
	diskData, _ := queryDisk(influxdbClient, orgId)
	networkData, _ := queryNetwork(influxdbClient, orgId)

	for key, cpuVal := range cpuData {
		to := ptdomain.PerformanceExecMetrics{
			RunnerName: key,

			CpuUsage:    cpuVal,
			MemoryUsage: 0,

			DiskUsages:    map[string]float64{},
			NetworkUsages: map[string]float64{},
		}

		to.MemoryUsage = memoryData[key]

		for name, val := range diskData[key] {
			to.DiskUsages[name] = val
		}

		for name, val := range networkData[key] {
			to.NetworkUsages[name] = val
		}

		ret = append(ret, to)
	}

	return
}
func queryCpu(influxdbClient influxdb2.Client, orgId string) (
	ret map[string]float64, err error) {

	ret = map[string]float64{}

	flux := fmt.Sprintf(`
from(bucket: "%s")
    |> range(start: -1m)
    |> filter(
        fn: (r) =>
            r._measurement == "%s",
    )
    |> aggregateWindow(
	   every: 1m, 
	   fn: mean, 
	   createEmpty: false)
`, bucketName, tableCpuUsage)

	result, err := queryData(influxdbClient, orgId, flux)

	for result.Next() {
		mp := result.Record().Values()

		ret[mp["runner"].(string)] = _floatUtils.PointNumb(mp["_value"].(float64), 2)
	}

	return
}

func queryMemory(influxdbClient influxdb2.Client, orgId string) (
	ret map[string]float64, err error) {

	ret = map[string]float64{}

	flux := fmt.Sprintf(`
from(bucket: "%s")
    |> range(start: -1m)
    |> filter(
        fn: (r) =>
            r._measurement == "%s",
    )
    |> aggregateWindow(
	   every: 1m, 
	   fn: mean, 
	   createEmpty: false)
`, bucketName, tableMemoryUsage)

	result, err := queryData(influxdbClient, orgId, flux)

	for result.Next() {
		mp := result.Record().Values()

		ret[mp["runner"].(string)] = _floatUtils.PointNumb(mp["_value"].(float64), 2)
	}

	return
}
func queryDisk(influxdbClient influxdb2.Client, orgId string) (
	ret map[string]map[string]float64, err error) {

	ret = map[string]map[string]float64{}

	flux := fmt.Sprintf(`
from(bucket: "%s")
    |> range(start: -1m)
    |> filter(
        fn: (r) =>
            r._measurement == "%s",
    )
    |> aggregateWindow(
	   every: 1m, 
	   fn: mean, 
	   createEmpty: false)
`, bucketName, tableDiskUsage)

	result, err := queryData(influxdbClient, orgId, flux)

	for result.Next() {
		mp := result.Record().Values()

		runner := mp["runner"].(string)
		name := mp["name"].(string)
		val := _floatUtils.PointNumb(mp["_value"].(float64), 2)

		_, ok := ret[runner]
		if !ok {
			ret[runner] = map[string]float64{}
		}

		ret[runner][name] = val
	}

	return
}
func queryNetwork(influxdbClient influxdb2.Client, orgId string) (
	ret map[string]map[string]float64, err error) {

	ret = map[string]map[string]float64{}

	flux := fmt.Sprintf(`
from(bucket: "%s")
    |> range(start: -1m)
    |> filter(
        fn: (r) =>
            r._measurement == "%s",
    )
    |> aggregateWindow(
	   every: 1m, 
	   fn: mean, 
	   createEmpty: false)
`, bucketName, tableNetworkUsage)

	result, err := queryData(influxdbClient, orgId, flux)

	for result.Next() {
		mp := result.Record().Values()

		runner := mp["runner"].(string)
		name := mp["name"].(string)
		val := _floatUtils.PointNumb(mp["_value"].(float64), 2)

		_, ok := ret[runner]
		if !ok {
			ret[runner] = map[string]float64{}
		}

		ret[runner][name] = val
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
