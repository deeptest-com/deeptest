package agentExec

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
	"sync"
)

const (
	ServerApiPath = "api/v1"
)

var (
	ContextStore sync.Map
)

func InitUserExecContext(execUuid string) {
	val := UserContext{
		ScopedVariables: map[uint][]domain.ExecVariable{},
		ScopedCookies:   map[uint][]domain.ExecCookie{},
		ScopeHierarchy:  map[uint]*[]uint{},

		ExecScene:      domain.ExecScene{},
		DatapoolCursor: map[string]int{},
	}

	ContextStore.Store(execUuid, &val)
}

func ClearExec(execUuid string) {
	ContextStore.Store(execUuid, nil)
}

func GetUserExecContext(execUuid string) (val *UserContext) {
	_, ok := ContextStore.Load(execUuid)
	if !ok {
		InitUserExecContext(execUuid)
	}

	inf, _ := ContextStore.Load(execUuid)
	val = inf.(*UserContext)

	return
}

func SetIsRunning(execUuid string, val bool) {
	entity := GetUserExecContext(execUuid)
	entity.IsRunning = val

	//ContextStore.Store(execUuid, entity)
}
func GetIsRunning(execUuid string) (ret bool) {
	userContext := GetUserExecContext(execUuid)
	ret = userContext.ForceStopExec

	return
}

func SetForceStopExec(execUuid string, val bool) {
	entity := GetUserExecContext(execUuid)
	entity.ForceStopExec = val

	//ContextStore.Store(execUuid, entity)
}
func GetForceStopExec(execUuid string) (ret bool) {
	userContext := GetUserExecContext(execUuid)
	ret = userContext.ForceStopExec

	return
}

func SetScopedVariables(execUuid string, val map[uint][]domain.ExecVariable) {
	entity := GetUserExecContext(execUuid)
	entity.ScopedVariables = val

	//ContextStore.Store(execUuid, entity)
}
func GetScopedVariables(execUuid string) (ret map[uint][]domain.ExecVariable) {
	userContext := GetUserExecContext(execUuid)
	ret = userContext.ScopedVariables

	return
}

func SetScopeHierarchy(execUuid string, val map[uint]*[]uint) {
	entity := GetUserExecContext(execUuid)
	entity.ScopeHierarchy = val

	//ContextStore.Store(execUuid, entity)
}
func GetScopeHierarchy(execUuid string) (ret map[uint]*[]uint) {
	userContext := GetUserExecContext(execUuid)
	ret = userContext.ScopeHierarchy

	return
}

func SetExecScene(execUuid string, val domain.ExecScene) {
	entity := GetUserExecContext(execUuid)
	entity.ExecScene = val

	//ContextStore.Store(execUuid, entity)
}
func GetExecScene(execUuid string) (ret domain.ExecScene) {
	userContext := GetUserExecContext(execUuid)
	ret = userContext.ExecScene

	return
}

func SetDatapoolCursor(execUuid string, val map[string]int) {
	entity := GetUserExecContext(execUuid)
	entity.DatapoolCursor = val

	//ContextStore.Store(execUuid, entity)
}
func GetDatapoolCursor(execUuid string) (ret map[string]int) {
	userContext := GetUserExecContext(execUuid)
	ret = userContext.DatapoolCursor

	return
}

func SetScopedCookies(execUuid string, val map[uint][]domain.ExecCookie) {
	entity := GetUserExecContext(execUuid)
	entity.ScopedCookies = val

	//ContextStore.Store(execUuid, entity)
}
func GetScopedCookies(execUuid string) (ret map[uint][]domain.ExecCookie) {
	userContext := GetUserExecContext(execUuid)
	ret = userContext.ScopedCookies

	return
}

func GetServerUrl(execUuid string) (ret string) {
	userContext := GetUserExecContext(execUuid)
	ret = userContext.ServerUrl

	return
}
func SetServerUrl(execUuid string, val string) {
	entity := GetUserExecContext(execUuid)
	entity.ServerUrl = val

	//ContextStore.Store(execUuid, entity)
}

func SetServerToken(execUuid string, val string) {
	entity := GetUserExecContext(execUuid)
	entity.ServerToken = val

	//ContextStore.Store(execUuid, entity)
}
func GetServerToken(execUuid string) (ret string) {
	userContext := GetUserExecContext(execUuid)
	ret = userContext.ServerToken

	return
}

func SetCurrRequest(execUuid string, val domain.BaseRequest) {
	entity := GetUserExecContext(execUuid)
	entity.CurrRequest = val

	//ContextStore.Store(execUuid, entity)
}
func GetCurrRequest(execUuid string) (ret domain.BaseRequest) {
	userContext := GetUserExecContext(execUuid)
	ret = userContext.CurrRequest

	return
}

func SetCurrResponse(execUuid string, val domain.DebugResponse) {
	entity := GetUserExecContext(execUuid)
	entity.CurrResponse = val

	//ContextStore.Store(execUuid, entity)
}
func GetCurrResponse(execUuid string) (ret domain.DebugResponse) {
	userContext := GetUserExecContext(execUuid)
	ret = userContext.CurrResponse

	return
}

func SetCurrScenarioProcessor(execUuid string, val *Processor) {
	entity := GetUserExecContext(execUuid)
	entity.CurrScenarioProcessor = val

	//ContextStore.Store(execUuid, entity)
}
func GetCurrScenarioProcessor(execUuid string) (ret *Processor) {
	userContext := GetUserExecContext(execUuid)
	ret = userContext.CurrScenarioProcessor

	return
}

func SetCurrScenarioProcessorId(execUuid string, val uint) {
	entity := GetUserExecContext(execUuid)
	entity.CurrScenarioProcessorId = val

	//ContextStore.Store(execUuid, entity)
}
func GetCurrScenarioProcessorId(execUuid string) (ret uint) {
	userContext := GetUserExecContext(execUuid)
	ret = userContext.CurrScenarioProcessorId

	return
}

func SetCurrDebugInterfaceId(execUuid string, val uint) {
	entity := GetUserExecContext(execUuid)
	entity.CurrDebugInterfaceId = val

	//ContextStore.Store(execUuid, entity)
}
func GetCurrDebugInterfaceId(execUuid string) (ret uint) {
	userContext := GetUserExecContext(execUuid)
	ret = userContext.CurrDebugInterfaceId

	return
}

func InitGojaRuntime(execUuid string) (execRuntime *goja.Runtime, execRequire *require.RequireModule) {
	userContext := GetUserExecContext(execUuid)
	execRuntime = userContext.GojaRuntime

	execRuntime = goja.New()
	execRuntime.SetFieldNameMapper(goja.TagFieldNameMapper("json", true))
	registry := new(require.Registry) // registry 能夠被多个goja.Runtime共用
	execRequire = registry.Enable(execRuntime)

	userContext.GojaRuntime = execRuntime
	userContext.GojaRequire = execRequire

	return
}

func GetGojaRuntime(execUuid string) (execRuntime *goja.Runtime, execRequire *require.RequireModule) {
	userContext := GetUserExecContext(execUuid)
	execRuntime = userContext.GojaRuntime
	execRequire = userContext.GojaRequire

	return
}

func GetGojaVariables(execUuid string) (ret []domain.ExecVariable) {
	userContext := GetUserExecContext(execUuid)
	ret = userContext.GojaVariables

	return
}
func SetGojaVariables(execUuid string, val []domain.ExecVariable) {
	entity := GetUserExecContext(execUuid)
	entity.GojaVariables = val

	//ContextStore.Store(execUuid, entity)

	return
}
func AppendGojaVariables(execUuid string, val domain.ExecVariable) {
	userContext := GetUserExecContext(execUuid)
	varis := userContext.GojaVariables

	varis = append(varis, val)

	SetGojaVariables(execUuid, varis)

	return
}

func GetGojaLogs(execUuid string) (ret []string) {
	userContext := GetUserExecContext(execUuid)
	ret = userContext.GojaLogs

	return
}
func SetGojaLogs(execUuid string, val []string) {
	entity := GetUserExecContext(execUuid)
	entity.GojaLogs = val

	//ContextStore.Store(execUuid, entity)

	return
}
func AppendGojaLogs(execUuid string, val string) {
	userContext := GetUserExecContext(execUuid)
	logs := userContext.GojaLogs

	logs = append(logs, val)

	SetGojaLogs(execUuid, logs)

	return
}

type UserContext struct {
	IsRunning     bool
	ForceStopExec bool

	ServerUrl   string
	ServerToken string

	// used to exchange request and response data between goja and go
	CurrRequest  domain.BaseRequest
	CurrResponse domain.DebugResponse

	CurrScenarioProcessor   *Processor
	CurrScenarioProcessorId uint
	CurrDebugInterfaceId    uint

	ScopedVariables map[uint][]domain.ExecVariable // for scenario and debug
	ScopedCookies   map[uint][]domain.ExecCookie   // only for scenario
	ScopeHierarchy  map[uint]*[]uint               // only for scenario (processId -> ancestorProcessIds)

	ExecScene      domain.ExecScene // for scenario and debug, from server
	DatapoolCursor map[string]int   // only for scenario

	// for goja js engine
	GojaRuntime   *goja.Runtime
	GojaRequire   *require.RequireModule
	GojaVariables []domain.ExecVariable
	GojaLogs      []string
}
