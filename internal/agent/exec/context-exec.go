package agentExec

import (
	agentDomain "github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"sync"
)

const (
	ServerApiPath = "api/v1"
)

var (
	ExecContextStore sync.Map
)

func InitUserExecContext(execUuid string) {
	val := UserContext{
		ScopedVariables: map[uint][]domain.ExecVariable{},
		ScopedCookies:   map[uint][]domain.ExecCookie{},
		ScopeHierarchy:  map[uint]*[]uint{},

		ExecScene:      domain.ExecScene{},
		DatapoolCursor: map[string]int{},
	}

	ExecContextStore.Store(execUuid, &val)
}

func ClearExecContext(execUuid string) {
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

func SetIsRunning(execUuid string, val bool) {
	entity := GetUserExecContext(execUuid)
	entity.IsRunning = val

	//ExecContextStore.Store(execUuid, entity)
}
func GetIsRunning(execUuid string) (ret bool) {
	userContext := GetUserExecContext(execUuid)
	ret = userContext.ForceStopExec

	return
}

func SetForceStopExec(execUuid string, val bool) {
	entity := GetUserExecContext(execUuid)
	entity.ForceStopExec = val

	//ExecContextStore.Store(execUuid, entity)
}
func GetForceStopExec(execUuid string) (ret bool) {
	userContext := GetUserExecContext(execUuid)
	ret = userContext.ForceStopExec

	return
}

func SetScopedVariables(execUuid string, val map[uint][]domain.ExecVariable) {
	entity := GetUserExecContext(execUuid)
	entity.ScopedVariables = val

	//ExecContextStore.Store(execUuid, entity)
}
func GetScopedVariables(execUuid string) (ret map[uint][]domain.ExecVariable) {
	userContext := GetUserExecContext(execUuid)
	ret = userContext.ScopedVariables

	return
}

func SetScopeHierarchy(execUuid string, val map[uint]*[]uint) {
	entity := GetUserExecContext(execUuid)
	entity.ScopeHierarchy = val

	//ExecContextStore.Store(execUuid, entity)
}
func GetScopeHierarchy(execUuid string) (ret map[uint]*[]uint) {
	userContext := GetUserExecContext(execUuid)
	ret = userContext.ScopeHierarchy

	return
}

func SetExecScene(execUuid string, val domain.ExecScene) {
	entity := GetUserExecContext(execUuid)
	entity.ExecScene = val

	//ExecContextStore.Store(execUuid, entity)
}
func GetExecScene(execUuid string) (ret domain.ExecScene) {
	userContext := GetUserExecContext(execUuid)
	ret = userContext.ExecScene

	return
}

func SetDatapoolCursor(execUuid string, val map[string]int) {
	entity := GetUserExecContext(execUuid)
	entity.DatapoolCursor = val

	//ExecContextStore.Store(execUuid, entity)
}
func GetDatapoolCursor(execUuid string) (ret map[string]int) {
	userContext := GetUserExecContext(execUuid)
	ret = userContext.DatapoolCursor

	return
}

func SetScopedCookies(execUuid string, val map[uint][]domain.ExecCookie) {
	entity := GetUserExecContext(execUuid)
	entity.ScopedCookies = val

	//ExecContextStore.Store(execUuid, entity)
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

	//ExecContextStore.Store(execUuid, entity)
}

func SetServerToken(execUuid string, val string) {
	entity := GetUserExecContext(execUuid)
	entity.ServerToken = val

	//ExecContextStore.Store(execUuid, entity)
}
func GetServerToken(execUuid string) (ret string) {
	userContext := GetUserExecContext(execUuid)
	ret = userContext.ServerToken

	return
}

func SetCurrRequest(execUuid string, val domain.BaseRequest) {
	entity := GetUserExecContext(execUuid)
	entity.CurrRequest = val

	//ExecContextStore.Store(execUuid, entity)
}
func GetCurrRequest(execUuid string) (ret domain.BaseRequest) {
	userContext := GetUserExecContext(execUuid)
	ret = userContext.CurrRequest

	return
}

func SetCurrResponse(execUuid string, val domain.DebugResponse) {
	entity := GetUserExecContext(execUuid)
	entity.CurrResponse = val

	//ExecContextStore.Store(execUuid, entity)
}
func GetCurrResponse(execUuid string) (ret domain.DebugResponse) {
	userContext := GetUserExecContext(execUuid)
	ret = userContext.CurrResponse

	return
}

func SetCurrScenarioProcessor(execUuid string, val *Processor) {
	entity := GetUserExecContext(execUuid)
	entity.CurrScenarioProcessor = val

	//ExecContextStore.Store(execUuid, entity)
}
func GetCurrScenarioProcessor(execUuid string) (ret *Processor) {
	userContext := GetUserExecContext(execUuid)
	ret = userContext.CurrScenarioProcessor

	return
}

func SetCurrScenarioProcessorId(execUuid string, val uint) {
	entity := GetUserExecContext(execUuid)
	entity.CurrScenarioProcessorId = val

	//ExecContextStore.Store(execUuid, entity)
}
func GetCurrScenarioProcessorId(execUuid string) (ret uint) {
	userContext := GetUserExecContext(execUuid)
	ret = userContext.CurrScenarioProcessorId

	return
}

func SetCurrDebugInterfaceId(execUuid string, val uint) {
	entity := GetUserExecContext(execUuid)
	entity.CurrDebugInterfaceId = val

	//ExecContextStore.Store(execUuid, entity)
}
func GetCurrDebugInterfaceId(execUuid string) (ret uint) {
	userContext := GetUserExecContext(execUuid)
	ret = userContext.CurrDebugInterfaceId

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

func GetGojaVariables(execUuid string) (ret *[]domain.ExecVariable) {
	userContext := GetUserExecContext(execUuid)
	ret = userContext.GojaVariables

	if ret == nil {
		ret = &[]domain.ExecVariable{}
	}

	return
}
func SetGojaVariables(execUuid string, val *[]domain.ExecVariable) {
	entity := GetUserExecContext(execUuid)
	entity.GojaVariables = val

	return
}
func ResetGojaVariables(execUuid string) {
	entity := GetUserExecContext(execUuid)
	entity.GojaVariables = nil
	return
}
func AppendGojaVariables(execUuid string, val domain.ExecVariable) {
	varis := GetGojaVariables(execUuid)

	*varis = append(*varis, val)

	return
}

func GetGojaLogs(execUuid string) (ret *[]string) {
	userContext := GetUserExecContext(execUuid)
	ret = userContext.GojaLogs

	if ret == nil {
		ret = &[]string{}
		SetGojaLogs(execUuid, ret)
	}

	return
}
func SetGojaLogs(execUuid string, val *[]string) {
	entity := GetUserExecContext(execUuid)
	entity.GojaLogs = val

	return
}
func ResetGojaLogs(execUuid string) {
	entity := GetUserExecContext(execUuid)
	entity.GojaLogs = nil

	return
}
func AppendGojaLogs(execUuid string, val string) {
	logs := GetGojaLogs(execUuid)

	*logs = append(*logs, val)

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

	// for report data
	InterfaceStat *agentDomain.InterfaceStat

	// for goja js engine
	GojaVariables *[]domain.ExecVariable
	GojaLogs      *[]string
}
