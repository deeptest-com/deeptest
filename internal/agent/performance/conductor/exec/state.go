package conductorExec

import (
	ptconsts "github.com/aaronchen2k/deeptest/internal/agent/performance/pkg/consts"
	ptdomain "github.com/aaronchen2k/deeptest/internal/agent/performance/pkg/domain"
	ptproto "github.com/aaronchen2k/deeptest/internal/agent/performance/proto"
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
	ItemsStore        sync.Map
	TestServicesStore sync.Map
	LogServicesStore  sync.Map
)

func GetCurrItem() (ret *ptdomain.TestItem) {
	obj, ok := ItemsStore.Load(KeyTests)
	if !ok {
		return
	}

	items := obj.(*[]*ptdomain.TestItem)

	for i := len(*items) - 1; i >= 0; i-- {
		if (*items)[i].ConductorReq != nil { // get last one conductor item
			ret = (*items)[i]
		}
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
func AddTestItem(room string, role ptconsts.TestRole, conductorReq *ptdomain.PerformanceTestReq,
	runners []*ptdomain.Runner, runnerReq *ptproto.PerformanceExecStartReq) (ret *ptdomain.TestItem) {

	arr := make([]*ptdomain.TestItem, 0)
	tests := &arr

	obj, ok := ItemsStore.Load(KeyTests)
	if ok {
		tests = obj.(*[]*ptdomain.TestItem)
	} else {
		ItemsStore.Store(KeyTests, tests)
	}

	ret = &ptdomain.TestItem{
		Room:       room,
		Role:       role,
		Runners:    runners,
		CreateTime: time.Now(),
	}

	if ret.Role == ptconsts.Conductor {
		ret.ConductorReq = conductorReq
	} else {
		ret.RunnerReq = runnerReq
	}

	*tests = append(*tests, ret)

	return
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

	if index > 0 {
		*tests = append((*tests)[:index], (*tests)[index+1:]...)
	}
}

func GetTestService(room string) (ret *PerformanceTestService) {
	obj, ok := TestServicesStore.Load(room)
	if ok {
		ret = obj.(*PerformanceTestService)
	}

	return
}
func SetTestService(room string, service *PerformanceTestService) {
	TestServicesStore.Store(room, service)

	return
}
func DeleteTestService(room string) {
	TestServicesStore.Range(func(key, value interface{}) bool {
		if room == key {
			TestServicesStore.Delete(key)
		}
		return true
	})

	return
}

func GetLogService(room string) (ret *PerformanceLogService) {
	obj, ok := LogServicesStore.Load(room)
	if ok {
		ret = obj.(*PerformanceLogService)
	}

	return
}
func SetLogService(room string, service *PerformanceLogService) {
	LogServicesStore.Store(room, service)

	return
}
func DeleteLogService(room string) {
	LogServicesStore.Range(func(key, value interface{}) bool {
		if room == key {
			LogServicesStore.Delete(key)
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
