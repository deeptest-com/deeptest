package service

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/business"
	execHelper "github.com/aaronchen2k/deeptest/internal/server/modules/v1/helper/exec"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
	"github.com/kataras/iris/v12/websocket"
)

type ExecProcessorService struct {
	ScenarioProcessorRepo *repo.ScenarioProcessorRepo  `inject:""`
	ScenarioRepo          *repo.ScenarioRepo           `inject:""`
	TestResultRepo        *repo.TestResultRepo         `inject:""`
	TestLogRepo           *repo.TestLogRepo            `inject:""`
	InterfaceRepo         *repo.InterfaceRepo          `inject:""`
	InterfaceService      *InterfaceService            `inject:""`
	ExecRequestService    *business.ExecRequestService `inject:""`
}

//func (s *ExecLogService) ExecThreadGroup(processor model.Processor, log *domain.Log, msg websocket.Message) (
//	result string, err error) {
// threadGroup, err := s.ScenarioProcessorRepo.GetThreadGroup(*processor)
// execHelper.ExecThreadGroup(&threadGroup, parentLog, msg)
//	return
//}

func (s *ExecProcessorService) ExecLogic(processor *model.Processor, parentLog *domain.Log, msg websocket.Message) (
	output domain.Output, err error) {

	logic, err := s.ScenarioProcessorRepo.GetLogic(*processor)
	output, _ = execHelper.ExecLogic(&logic, parentLog, msg)

	return
}

func (s *ExecProcessorService) ExecLoop(processor *model.Processor, parentLog *domain.Log, msg websocket.Message) (
	output domain.Output, err error) {

	loop, err := s.ScenarioProcessorRepo.GetLoop(*processor)
	output, _ = execHelper.ExecLoop(&loop, parentLog, msg)

	return
}

func (s *ExecProcessorService) ExecData(processor *model.Processor, parentLog *domain.Log, msg websocket.Message) (
	output domain.Output, err error) {

	data, err := s.ScenarioProcessorRepo.GetData(*processor)
	output, _ = execHelper.ExecData(&data, parentLog, msg)

	return
}

func (s *ExecProcessorService) ExecTimer(processor *model.Processor, parentLog *domain.Log, msg websocket.Message) (
	output domain.Output, err error) {

	timer, err := s.ScenarioProcessorRepo.GetTimer(*processor)
	output, _ = execHelper.ExecTimer(&timer, parentLog, msg)

	return
}

func (s *ExecProcessorService) ExecVariable(processor *model.Processor, parentLog *domain.Log, msg websocket.Message) (
	output domain.Output, err error) {

	variable, err := s.ScenarioProcessorRepo.GetVariable(*processor)
	output, _ = execHelper.ExecVariable(&variable, parentLog, msg)

	return
}

func (s *ExecProcessorService) ExecAssertion(processor *model.Processor, parentLog *domain.Log, msg websocket.Message) (
	output domain.Output, err error) {

	assertion, err := s.ScenarioProcessorRepo.GetAssertion(*processor)
	output, _ = execHelper.ExecAssertion(&assertion, parentLog, msg)

	return
}

func (s *ExecProcessorService) ExecExtractor(processor *model.Processor, parentLog *domain.Log, msg websocket.Message) (
	output domain.Output, err error) {

	extractor, err := s.ScenarioProcessorRepo.GetExtractor(*processor)
	output, _ = execHelper.ExecExtractor(&extractor, parentLog, msg)

	return
}

func (s *ExecProcessorService) ExecCookie(processor *model.Processor, parentLog *domain.Log, msg websocket.Message) (
	output domain.Output, err error) {

	cookie, err := s.ScenarioProcessorRepo.GetCookie(*processor)
	output, _ = execHelper.ExecCookie(&cookie, parentLog, msg)

	return
}
