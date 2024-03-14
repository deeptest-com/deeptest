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
	ctx, cancel := context.WithCancel(context.Background())
	val := UserContext{
		ExecCtx:       ctx,
		ExecCancel:    cancel,
		InterfaceStat: &agentDomain.InterfaceStat{},
	}

	ExecContextStore.Store(execUuid, &val)
}

func GetExecCtx(execUuid string) (ctx context.Context, cancel context.CancelFunc) {
	userContext := GetUserExecContext(execUuid)

	if userContext != nil {
		ctx = userContext.ExecCtx
		cancel = userContext.ExecCancel
	}

	return
}
func CloseExecCtx(execUuid string) (ctx context.Context, cancel context.CancelFunc) {
	userContext := GetUserExecContext(execUuid)

	if userContext.ExecCancel != nil {
		userContext.ExecCancel()
	}

	userContext.ExecCtx = nil
	userContext.ExecCancel = nil

	ExecContextStore.Store(execUuid, nil)

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
		return
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
