package agentExec

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
)

func (s *ExecSession) SetIsRunning(val bool) {
	s.IsRunning = val
}
func (s *ExecSession) GetIsRunning() (ret bool) {
	ret = s.ForceStopExec

	return
}

func (s *ExecSession) SetScopedVariables(val map[uint][]domain.ExecVariable) {
	s.ScopedVariables = val
}
func (s *ExecSession) GetScopedVariables() (ret map[uint][]domain.ExecVariable) {
	ret = s.ScopedVariables

	return
}

func (s *ExecSession) SetScopeHierarchy(val map[uint]*[]uint) {
	s.ScopeHierarchy = val
}
func (s *ExecSession) GetScopeHierarchy() (ret map[uint]*[]uint) {
	ret = s.ScopeHierarchy

	return
}

func (s *ExecSession) SetExecScene(val domain.ExecScene) {
	s.ExecScene = val
}
func (s *ExecSession) GetExecScene() (ret domain.ExecScene) {
	ret = s.ExecScene

	return
}

func (s *ExecSession) SetDatapoolCursor(val map[string]int) {
	s.DatapoolCursor = val
}
func (s *ExecSession) GetDatapoolCursor() (ret map[string]int) {
	ret = s.DatapoolCursor

	return
}

func (s *ExecSession) SetScopedCookies(val map[uint][]domain.ExecCookie) {
	s.ScopedCookies = val
}
func (s *ExecSession) GetScopedCookies() (ret map[uint][]domain.ExecCookie) {
	ret = s.ScopedCookies

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
	s.CurrRequest = val
}
func (s *ExecSession) GetCurrRequest() (ret domain.BaseRequest) {
	ret = s.CurrRequest

	return
}

func (s *ExecSession) SetCurrResponse(val domain.DebugResponse) {
	s.CurrResponse = val
}
func (s *ExecSession) GetCurrResponse() (ret domain.DebugResponse) {
	ret = s.CurrResponse

	return
}

func (s *ExecSession) SetCurrScenarioProcessor(val *Processor) {
	s.CurrScenarioProcessor = val
}
func (s *ExecSession) GetCurrScenarioProcessor() (ret *Processor) {
	ret = s.CurrScenarioProcessor

	return
}

func (s *ExecSession) SetCurrScenarioProcessorId(val uint) {
	s.CurrScenarioProcessorId = val
}
func (s *ExecSession) GetCurrScenarioProcessorId() (ret uint) {
	ret = s.CurrScenarioProcessorId

	return
}

func (s *ExecSession) SetCurrEnvironmentId(id int) {
	s.CurrEnvironmentId = id
}
func (s *ExecSession) GetCurrEnvironmentId() (id int) {
	id = s.CurrEnvironmentId

	return
}

func (s *ExecSession) SetCurrDebugInterfaceId(val uint) {
	s.CurrDebugInterfaceId = val
}
func (s *ExecSession) GetCurrDebugInterfaceId() (ret uint) {
	ret = s.CurrDebugInterfaceId

	return
}
