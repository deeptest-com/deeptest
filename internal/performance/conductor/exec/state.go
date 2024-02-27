package conductorExec

import (
	ptconsts "github.com/aaronchen2k/deeptest/internal/performance/pkg/consts"
	ptdomain "github.com/aaronchen2k/deeptest/internal/performance/pkg/domain"
	ptproto "github.com/aaronchen2k/deeptest/internal/performance/proto"
	"sync"
	"time"
)

const (
	KeyTests = "tests"
)

var (
	suspendWsMsg bool
)

var (
	ItemsStore    sync.Map
	ServicesStore sync.Map
)

func GetCurrItem() (ret *ptdomain.TestItem) {
	obj, ok := ItemsStore.Load(KeyTests)
	if !ok {
		return
	}

	items := obj.(*[]*ptdomain.TestItem)

	if len(*items) > 0 {
		ret = (*items)[len(*items)-1] // get last one
	}

	return
}
func GetTestItems() (ret *[]*ptdomain.TestItem) {
	obj, ok := ItemsStore.Load(KeyTests)
	if ok {
		ret = obj.(*[]*ptdomain.TestItem)
	}

	return
}
func AddTestItem(room string, role ptconsts.TestRole,
	conductorReq *ptdomain.PerformanceTestReq, runnerReq *ptproto.PerformanceExecStartReq) {

	arr := make([]*ptdomain.TestItem, 0)
	tests := &arr

	obj, ok := ItemsStore.Load(KeyTests)
	if ok {
		tests = obj.(*[]*ptdomain.TestItem)
	} else {
		ItemsStore.Store(KeyTests, tests)
	}

	test := &ptdomain.TestItem{
		Room:       room,
		Role:       role,
		CreateTime: time.Now(),
	}

	if test.Role == ptconsts.Conductor {
		test.ConductorReq = conductorReq
	} else {
		test.RunnerReq = runnerReq
	}

	*tests = append(*tests, test)
}
func RemoveTestItem(room string) {
	obj, ok := ItemsStore.Load(KeyTests)
	if !ok {
		return
	}

	tests := obj.(*[]*ptdomain.TestItem)

	index := -1
	for i, item := range *tests {
		if (*item).Room == room {
			index = i
			break
		}
	}

	*tests = append((*tests)[:index], (*tests)[index+1:]...)
}

func GetTestService(room string) (ret *PerformanceTestService) {
	obj, ok := ServicesStore.Load(room)
	if ok {
		ret = obj.(*PerformanceTestService)
	}

	return
}
func SetTestService(room string, service *PerformanceTestService) {
	ServicesStore.Store(room, service)

	return
}
func DeleteTestService(room string) {
	ServicesStore.Range(func(key, value interface{}) bool {
		if room == key {
			ServicesStore.Delete(key)
		}
		return true
	})

	return
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
