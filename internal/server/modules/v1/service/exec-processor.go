package service

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/business"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
	"github.com/kataras/iris/v12/websocket"
)

type ExecProcessorService struct {
	ScenarioProcessorRepo *repo.ScenarioProcessorRepo `inject:""`
	ScenarioRepo          *repo.ScenarioRepo          `inject:""`
	TestResultRepo        *repo.ReportRepo            `inject:""`
	TestLogRepo           *repo.LogRepo               `inject:""`
	InterfaceRepo         *repo.InterfaceRepo         `inject:""`
	InterfaceService      *InterfaceService           `inject:""`
	ExecRequestService    *business.ExecRequest       `inject:""`
	ExecHelperService     *ExecHelperService          `inject:""`
}

//func (s *ExecLogService) ExecThreadGroup(processor model.Processor, log *domain.ExecLog, msg websocket.Message) (
//	result string, err error) {
// threadGroup, err := s.ScenarioProcessorRepo.GetThreadGroup(*processor)
// s.ExecComm.ExecThreadGroup(&threadGroup, parentLog, msg)
//	return
//}

func (s *ExecProcessorService) ExecLogic(processor *model.Processor, parentLog *domain.ExecLog, msg *websocket.Message) (
	output domain.ExecOutput, err error) {

	logic, err := s.ScenarioProcessorRepo.GetLogic(*processor)
	output, _ = s.ExecHelperService.HandleLogic(&logic, parentLog, msg)

	return
}

func (s *ExecProcessorService) ExecLoop(processor *model.Processor, parentLog *domain.ExecLog, msg *websocket.Message) (
	output domain.ExecOutput, err error) {

	loop, err := s.ScenarioProcessorRepo.GetLoop(*processor)
	output, _ = s.ExecHelperService.HandleLoop(&loop, parentLog, msg)

	return
}
func (s *ExecProcessorService) ExecLoopBreak(processor *model.Processor, parentLog *domain.ExecLog, msg *websocket.Message) (
	output domain.ExecOutput, err error) {

	loop, err := s.ScenarioProcessorRepo.GetLoop(*processor)
	output, _ = s.ExecHelperService.HandleLoopBreak(&loop, parentLog, msg)

	return
}

func (s *ExecProcessorService) ExecData(processor *model.Processor, parentLog *domain.ExecLog, msg *websocket.Message) (
	output domain.ExecOutput, err error) {

	data, err := s.ScenarioProcessorRepo.GetData(*processor)
	output, _ = s.ExecHelperService.HandleData(&data, parentLog, msg)

	return
}

func (s *ExecProcessorService) ExecTimer(processor *model.Processor, parentLog *domain.ExecLog, msg *websocket.Message) (
	output domain.ExecOutput, err error) {

	timer, err := s.ScenarioProcessorRepo.GetTimer(*processor)
	output, _ = s.ExecHelperService.HandleTimer(&timer, parentLog, msg)

	return
}

func (s *ExecProcessorService) ExecVariable(processor *model.Processor, parentLog *domain.ExecLog, msg *websocket.Message) (
	output domain.ExecOutput, err error) {

	variable, err := s.ScenarioProcessorRepo.GetVariable(*processor)
	output, _ = s.ExecHelperService.HandleVariable(&variable, parentLog, msg)

	return
}

func (s *ExecProcessorService) ExecAssertion(processor *model.Processor, parentLog *domain.ExecLog, msg *websocket.Message) (
	output domain.ExecOutput, err error) {

	assertion, err := s.ScenarioProcessorRepo.GetAssertion(*processor)
	output, _ = s.ExecHelperService.HandleAssertion(&assertion, parentLog, msg)

	return
}

func (s *ExecProcessorService) ExecExtractor(processor *model.Processor, parentLog *domain.ExecLog, msg *websocket.Message) (
	output domain.ExecOutput, err error) {

	extractor, err := s.ScenarioProcessorRepo.GetExtractor(*processor)
	output, _ = s.ExecHelperService.HandleExtractor(&extractor, parentLog, msg)

	return
}

func (s *ExecProcessorService) ExecCookie(processor *model.Processor, parentLog *domain.ExecLog, msg *websocket.Message) (
	output domain.ExecOutput, err error) {

	cookie, err := s.ScenarioProcessorRepo.GetCookie(*processor)
	output, _ = s.ExecHelperService.HandleCookie(&cookie, parentLog, msg)

	return
}
