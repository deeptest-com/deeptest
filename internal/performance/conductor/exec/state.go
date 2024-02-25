package conductorExec

import (
	ptdomain "github.com/aaronchen2k/deeptest/internal/performance/pkg/domain"
	"time"
)

var (
	runningTest *ptdomain.PerformanceTestReq

	suspendWsMsg bool
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

func IsWsMsgSuspend() bool {
	return suspendWsMsg
}

func SuspendWsMsg() {
	suspendWsMsg = true
}
func ResumeWsMsg() {
	go func() {
		time.Sleep(3 * time.Second)
		suspendWsMsg = false
	}()
}
