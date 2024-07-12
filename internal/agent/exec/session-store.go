package agentExec

import "github.com/aaronchen2k/deeptest/internal/pkg/domain"

func (s *ExecSession) SetIsRunning(val bool) {
	s._isRunning = val
}
func (s *ExecSession) GetIsRunning() (ret bool) {
	ret = s._isRunning
	return
}
func (s *ExecSession) SetForceStop(val bool) {
	s._forceStop = val
}
func (s *ExecSession) GetForceStop() (ret bool) {
	ret = s._forceStop
	return
}

func (s *ExecSession) SetCurrRequest(val domain.BaseRequest) {
	s.InterfaceDebug._currRequest = val
}
func (s *ExecSession) GetCurrRequest() (ret domain.BaseRequest) {
	if s.InterfaceDebug != nil {
		ret = s.InterfaceDebug._currRequest
	}
	return
}

func (s *ExecSession) SetCurrResponse(val domain.DebugResponse) {
	s.InterfaceDebug._currResponse = val
}
func (s *ExecSession) GetCurrResponse() (ret domain.DebugResponse) {
	if s.InterfaceDebug != nil {
		ret = s.InterfaceDebug._currResponse
	}

	return
}

func (s *ExecSession) SetCurrScenarioProcessor(val *Processor) {
	s.ScenarioDebug._currProcessor = val
}
func (s *ExecSession) GetCurrScenarioProcessor() (ret *Processor) {
	if s.ScenarioDebug != nil {
		ret = s.ScenarioDebug._currProcessor
	}
	return
}

func (s *ExecSession) SetCurrScenarioProcessorId(val uint) {
	s.ScenarioDebug._currProcessorId = val
}
func (s *ExecSession) GetCurrScenarioProcessorId() (ret uint) {
	if s.ScenarioDebug != nil {
		ret = s.ScenarioDebug._currProcessorId
	}
	return
}

func (s *ExecSession) SetCurrDebugInterfaceId(val uint) {
	s.InterfaceDebug._debugInterfaceId = val
}
func (s *ExecSession) GetCurrDebugInterfaceId() (ret uint) {
	if s.InterfaceDebug != nil {
		ret = s.InterfaceDebug._debugInterfaceId
	}
	return
}
