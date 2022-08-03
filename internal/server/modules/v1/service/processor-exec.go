package service

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	execHelper "github.com/aaronchen2k/deeptest/internal/server/modules/v1/helper/exec"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
	"github.com/kataras/iris/v12/websocket"
)

type ScenarioProcessorExecService struct {
	ScenarioProcessorRepo *repo.ScenarioProcessorRepo `inject:""`
	ScenarioRepo          *repo.ScenarioRepo          `inject:""`
	TestResultRepo        *repo.TestResultRepo        `inject:""`
	TestLogRepo           *repo.TestLogRepo           `inject:""`
}

//func (s *ScenarioProcessorExecService) ExecThreadGroup(processor model.TestProcessor, log *domain.Log, msg websocket.Message) (
//	result string, err error) {
// threadGroup, err := s.ScenarioProcessorRepo.GetThreadGroup(*processor)
// execHelper.ExecThreadGroup(&threadGroup, parentLog, msg)
//	return
//}

func (s *ScenarioProcessorExecService) ExecLogic(processor *model.TestProcessor, parentLog *domain.Log, msg websocket.Message) (
	result domain.Output, err error) {

	logic, err := s.ScenarioProcessorRepo.GetLogic(*processor)
	execHelper.ExecLogic(&logic, parentLog, msg)

	return
}

func (s *ScenarioProcessorExecService) ExecLoop(processor *model.TestProcessor, parentLog *domain.Log, msg websocket.Message) (
	output domain.Output, err error) {

	loop, err := s.ScenarioProcessorRepo.GetLoop(*processor)
	execHelper.ExecLoop(&loop, parentLog, msg)

	return
}

func (s *ScenarioProcessorExecService) ExecData(processor *model.TestProcessor, parentLog *domain.Log, msg websocket.Message) (
	result domain.Output, err error) {

	data, err := s.ScenarioProcessorRepo.GetData(*processor)
	execHelper.ExecData(&data, parentLog, msg)

	return
}

func (s *ScenarioProcessorExecService) ExecTimer(processor *model.TestProcessor, parentLog *domain.Log, msg websocket.Message) (
	result domain.Output, err error) {

	timer, err := s.ScenarioProcessorRepo.GetTimer(*processor)
	execHelper.ExecTimer(&timer, parentLog, msg)

	return
}

func (s *ScenarioProcessorExecService) ExecVariable(processor *model.TestProcessor, parentLog *domain.Log, msg websocket.Message) (
	result domain.Output, err error) {

	variable, err := s.ScenarioProcessorRepo.GetVariable(*processor)
	execHelper.ExecVariable(&variable, parentLog, msg)

	return
}

func (s *ScenarioProcessorExecService) ExecAssertion(processor *model.TestProcessor, parentLog *domain.Log, msg websocket.Message) (
	result domain.Output, err error) {

	assertion, err := s.ScenarioProcessorRepo.GetAssertion(*processor)
	execHelper.ExecAssertion(&assertion, parentLog, msg)

	return
}

func (s *ScenarioProcessorExecService) ExecExtractor(processor *model.TestProcessor, parentLog *domain.Log, msg websocket.Message) (
	result domain.Output, err error) {

	extractor, err := s.ScenarioProcessorRepo.GetExtractor(*processor)
	execHelper.ExecExtractor(&extractor, parentLog, msg)

	return
}

func (s *ScenarioProcessorExecService) ExecCookie(processor *model.TestProcessor, parentLog *domain.Log, msg websocket.Message) (
	result domain.Output, err error) {

	cookie, err := s.ScenarioProcessorRepo.GetCookie(*processor)
	execHelper.ExecCookie(&cookie, parentLog, msg)

	return
}
