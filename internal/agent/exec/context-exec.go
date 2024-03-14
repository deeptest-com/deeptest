package agentExec

import (
	"context"
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

func SetExecCtx(execUuid string, ctx context.Context, cancel context.CancelFunc) {
	entity := GetUserExecContext(execUuid)
	entity.ExecCtx = ctx
	entity.ExecCancel = cancel
}
func GetExecCtx(execUuid string) (ctx context.Context, cancel context.CancelFunc) {
	userContext := GetUserExecContext(execUuid)

	ctx = userContext.ExecCtx
	cancel = userContext.ExecCancel

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

func GetUserExecContext(execUuid string) (val *UserContext) {
	inf, ok := ExecContextStore.Load(execUuid)
	if !ok {
		InitUserExecContext(execUuid)
	}

	inf, _ = ExecContextStore.Load(execUuid)
	val = inf.(*UserContext)

	return
}

type UserContext struct {
	ExecCtx       context.Context
	ExecCancel    context.CancelFunc
	InterfaceStat *agentDomain.InterfaceStat // for report data
}
