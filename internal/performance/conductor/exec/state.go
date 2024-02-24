package conductorExec

import (
	ptdomain "github.com/aaronchen2k/deeptest/internal/performance/pkg/domain"
	"time"
)

var (
	runningTest *ptdomain.PerformanceTestReq

	suspendLog bool
)

func GetRunningRoom() (ret string) {
	test := GetRunningTest()

	if test != nil {
		ret = test.Room
	}

	return
}
func GetRunningTest() *ptdomain.PerformanceTestReq {
	return runningTest
}
func SetRunningTest(val *ptdomain.PerformanceTestReq) {
	runningTest = val
}

func IsLogSuspend() bool {
	return suspendLog
}

func SuspendLog() {
	suspendLog = true
}

func ResumeLog() {
	go func() {
		time.Sleep(3 * time.Second)
		suspendLog = false
	}()
}
