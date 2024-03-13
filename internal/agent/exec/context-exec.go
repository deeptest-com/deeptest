package agentExec

import (
	agentDomain "github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"sync"
)

var (
	ExecContextStore sync.Map
)

func InitUserExecContext(execUuid string) {
	val := UserContext{}

	ExecContextStore.Store(execUuid, &val)
}

func ClearExecContext(execUuid string) {
	ExecContextStore.Store(execUuid, nil)
}

func SetForceStopExec(execUuid string, val bool) {
	entity := GetUserExecContext(execUuid)
	entity.ForceStopExec = val
}
func GetForceStopExec(execUuid string) (ret bool) {
	userContext := GetUserExecContext(execUuid)
	ret = userContext.ForceStopExec

	return
}

func SetIsRunning(execUuid string, val bool) {
	entity := GetUserExecContext(execUuid)
	entity.IsRunning = val
}
func GetIsRunning(execUuid string) (ret bool) {
	userContext := GetUserExecContext(execUuid)
	ret = userContext.ForceStopExec

	return
}

func SetInterfaceStat(execUuid string, val *agentDomain.InterfaceStat) {
	entity := GetUserExecContext(execUuid)
	entity.InterfaceStat = val
}
func GetInterfaceStat(execUuid string) (ret *agentDomain.InterfaceStat) {
	userContext := GetUserExecContext(execUuid)
	ret = userContext.InterfaceStat

	return
}

type UserContext struct {
	IsRunning     bool
	ForceStopExec bool
	InterfaceStat *agentDomain.InterfaceStat // for report data
}

func GetUserExecContext(execUuid string) (val *UserContext) {
	inf, ok := ExecContextStore.Load(execUuid)
	if !ok {
		InitUserExecContext(execUuid)
	}

	inf, _ = ExecContextStore.Load(execUuid)
	val = inf.(*UserContext)

	return
}
