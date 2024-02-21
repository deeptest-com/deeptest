package test

import (
	ptlog "github.com/aaronchen2k/deeptest/internal/performance/pkg/log"
	"testing"
)

const (
	influxdbAddress = "http://localhost:8087"
	influxdbOrg     = "deeptest"
	influxdbToken   = "CjK5KHeIopceCfRznN7RZxlffNrnCOBJ6Ugi9PCFb-mRu4ZQJ01tqpE4oeWmw5VlaDk-y3JMkKSx8k8Klwh04g=="
)

func TestQuery(t *testing.T) {
	ptlog.Init("")

	//influxdbClient := influxdb2.NewClient(influxdbAddress, influxdbToken)

	//indicator.QueryResponseTimeSummary(context.Background(), influxdbClient, influxdbOrg)

}
