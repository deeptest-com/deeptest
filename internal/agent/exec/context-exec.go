package agentExec

import (
	"context"
	"sync"
)

const (
	ServerApiPath = "api/v1"
)

var (
	ExecContextStore sync.Map
)

func InitUserExecContext(execUuid string) {
	ctx, cancel := context.WithCancel(context.Background())
	val := UserContext{
		ExecCtx:    ctx,
		ExecCancel: cancel,
	}

	ExecContextStore.Store(execUuid, &val)
}

func ClearUserExecContext(execUuid string) {
	ExecContextStore.Store(execUuid, nil)
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
func GetExecCtx(execUuid string) (ctx context.Context, cancel context.CancelFunc) {
	userContext := GetUserExecContext(execUuid)

	if userContext != nil {
		ctx = userContext.ExecCtx
		cancel = userContext.ExecCancel
	}

	return
}

type UserContext struct {
	ExecCtx    context.Context
	ExecCancel context.CancelFunc
}
