package agentExec

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
)

func (s *ExecSession) SetIsRunning(val bool) {
	s.IsRunning = val
}
func (s *ExecSession) GetIsRunning() (ret bool) {
	ret = s.IsRunning
	return
}
func (s *ExecSession) SetForceStop(val bool) {
	s.ForceStop = val
}
func (s *ExecSession) GetForceStop() (ret bool) {
	ret = s.ForceStop
	return
}

//func (s *ExecSession) SetScopedVariables(val map[uint][]domain.ExecVariable) {
//	s.ScenarioDebug.ScopedVariables = val
//}
//func (s *ExecSession) GetScopedVariables() (ret map[uint][]domain.ExecVariable) {
//	ret = s.ScenarioDebug.ScopedVariables
//	return
//}
//
//func (s *ExecSession) SetScopeHierarchy(val map[uint]*[]uint) {
//	s.ScenarioDebug.ScopeHierarchy = val
//}
//func (s *ExecSession) GetScopeHierarchy() (ret map[uint]*[]uint) {
//	ret = s.ScenarioDebug.ScopeHierarchy
//	return
//}

func (s *ExecSession) SetExecScene(val domain.ExecScene) {
	s.ExecScene = val
}
func (s *ExecSession) GetExecScene() (ret domain.ExecScene) {
	ret = s.ExecScene
	return
}

func (s *ExecSession) SetDatapoolCursor(val map[string]int) {
	s.ScenarioDebug.DatapoolCursor = val
}
func (s *ExecSession) GetDatapoolCursor() (ret map[string]int) {
	ret = s.ScenarioDebug.DatapoolCursor
	return
}

func (s *ExecSession) SetScopedCookies(val map[uint][]domain.ExecCookie) {
	s.ScenarioDebug.ScopedCookies = val
}
func (s *ExecSession) GetScopedCookies() (ret map[uint][]domain.ExecCookie) {
	ret = s.ScenarioDebug.ScopedCookies
	return
}

func (s *ExecSession) GetServerUrl() (ret string) {
	ret = s.ServerUrl
	return
}
func (s *ExecSession) SetServerUrl(val string) {
	s.ServerUrl = val
}

func (s *ExecSession) SetServerToken(val string) {
	s.ServerToken = val
}
func (s *ExecSession) GetServerToken() (ret string) {
	ret = s.ServerToken
	return
}

func (s *ExecSession) SetCurrRequest(val domain.BaseRequest) {
	s.InterfaceDebug.CurrRequest = val
}
func (s *ExecSession) GetCurrRequest() (ret domain.BaseRequest) {
	ret = s.InterfaceDebug.CurrRequest
	return
}

func (s *ExecSession) SetCurrResponse(val domain.DebugResponse) {
	s.InterfaceDebug.CurrResponse = val
}
func (s *ExecSession) GetCurrResponse() (ret domain.DebugResponse) {
	ret = s.InterfaceDebug.CurrResponse

	return
}

func (s *ExecSession) SetCurrScenarioProcessor(val *Processor) {
	s.ScenarioDebug.CurrProcessor = val
}
func (s *ExecSession) GetCurrScenarioProcessor() (ret *Processor) {
	ret = s.ScenarioDebug.CurrProcessor
	return
}

func (s *ExecSession) SetCurrScenarioProcessorId(val uint) {
	s.ScenarioDebug.CurrProcessorId = val
}
func (s *ExecSession) GetCurrScenarioProcessorId() (ret uint) {
	ret = s.ScenarioDebug.CurrProcessorId
	return
}

func (s *ExecSession) SetCurrEnvironmentId(id uint) {
	s.EnvironmentId = id
}
func (s *ExecSession) GetCurrEnvironmentId() (id uint) {
	id = s.EnvironmentId
	return
}

func (s *ExecSession) SetCurrDebugInterfaceId(val uint) {
	s.InterfaceDebug.DebugInterfaceId = val
}
func (s *ExecSession) GetCurrDebugInterfaceId() (ret uint) {
	ret = s.InterfaceDebug.DebugInterfaceId
	return
}
