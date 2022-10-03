package service

import (
	"fmt"
	 "github.com/aaronchen2k/deeptest/internal/agent/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/business"
	"github.com/aaronchen2k/deeptest/internal/server/modules/helper/exec"
	 "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/kataras/iris/v12/websocket"
	"time"
)

type ExecProcessorService struct {
	ScenarioProcessorRepo *repo.ScenarioProcessorRepo `inject:""`
	ScenarioRepo          *repo.ScenarioRepo          `inject:""`
	TestResultRepo        *repo.ReportRepo            `inject:""`
	TestLogRepo           *repo.LogRepo               `inject:""`
	InterfaceRepo         *repo.InterfaceRepo         `inject:""`
	InterfaceService      *InterfaceService           `inject:""`
	ExecHelperService     *ExecHelperService           `inject:""`
	ExecContext           *business.ExecContext        `inject:""`
}

//func (s *ExecLogService) ExecThreadGroup(processor model.Processor, log *domain.ExecLog, msg websocket.Message) (
//	result string, err error) {
// threadGroup, err := s.ScenarioProcessorRepo.GetThreadGroup(*processor)
// s.ExecComm.ExecThreadGroup(&threadGroup, parentLog, msg)
//	return
//}

func (s *ExecProcessorService) ExecLogic(processor *agentDomain.Processor, parentLog *domain.ExecLog, msg *websocket.Message) (
	output domain.ExecOutput, err error) {

	po, _ := s.ScenarioProcessorRepo.Get(processor.ID)
	logic, err := s.ScenarioProcessorRepo.GetLogic(po)
	output, _ = s.ExecHelperService.EvaluateLogic(&logic, parentLog, msg)

	return
}

func (s *ExecProcessorService) ExecLoop(processor *agentDomain.Processor, parentLog *domain.ExecLog, msg *websocket.Message) (
	output domain.ExecOutput, err error) {

	po, _ := s.ScenarioProcessorRepo.Get(processor.ID)
	loop, err := s.ScenarioProcessorRepo.GetLoop(po)
	output, _ = s.ExecHelperService.EvaluateLoop(&loop, parentLog, msg)

	return
}
func (s *ExecProcessorService) ExecLoopBreak(processor *agentDomain.Processor, parentLog *domain.ExecLog, msg *websocket.Message) (
	output domain.ExecOutput, err error) {

	po, _ := s.ScenarioProcessorRepo.Get(processor.ID)
	loop, err := s.ScenarioProcessorRepo.GetLoop(po)
	output, _ = s.ExecHelperService.EvaluateLoopBreak(&loop, parentLog, msg)

	breakFrom := output.BreakFrom
	breakIfExpress := output.Expression

	result, err := s.ExecHelperService.EvaluateGovaluateExpression(breakIfExpress, processor.ID)
	pass, ok := result.(bool)
	if err == nil && ok && pass {
		breakMap.Store(breakFrom, true)
		output.Msg = "真"
	} else {
		output.Msg = "假"
	}

	return
}

func (s *ExecProcessorService) ExecData(processor *agentDomain.Processor, parentLog *domain.ExecLog, msg *websocket.Message) (
	output domain.ExecOutput, err error) {

	po, _ := s.ScenarioProcessorRepo.Get(processor.ID)
	data, err := s.ScenarioProcessorRepo.GetData(po)
	output, _ = s.ExecHelperService.EvaluateData(&data, parentLog, msg)

	return
}

func (s *ExecProcessorService) ExecTimer(processor *agentDomain.Processor, parentLog *domain.ExecLog, msg *websocket.Message) (
	output domain.ExecOutput, err error) {

	po, _ := s.ScenarioProcessorRepo.Get(processor.ID)
	timer, err := s.ScenarioProcessorRepo.GetTimer(po)
	output, _ = s.ExecHelperService.EvaluateTimer(&timer, parentLog, msg)

	<-time.After(time.Duration(output.SleepTime) * time.Second)

	return
}

func (s *ExecProcessorService) ExecPrint(processor *agentDomain.Processor, parentLog *domain.ExecLog, msg *websocket.Message) (
	output domain.ExecOutput, err error) {

	po, _ := s.ScenarioProcessorRepo.Get(processor.ID)
	print, err := s.ScenarioProcessorRepo.GetPrint(po)
	output, _ = s.ExecHelperService.EvaluatePrint(&print, parentLog, msg)

	expression := s.ExecHelperService.ReplaceVariablesWithVerbs(output.Expression)
	variables := execHelper.GetVariablesInVariablePlaceholder(output.Expression)

	variableValues := make([]interface{}, 0)
	for _, name := range variables {
		val, err1 := s.ExecHelperService.GetVariableValueByName(processor.ID, name)
		if err1 != nil {
			val = "空"
		}
		variableValues = append(variableValues, val)
	}

	output.Msg = fmt.Sprintf(expression, variableValues...)

	return
}

func (s *ExecProcessorService) ExecVariable(processor *agentDomain.Processor, parentLog *domain.ExecLog, msg *websocket.Message) (
	output domain.ExecOutput, err error) {

	po, _ := s.ScenarioProcessorRepo.Get(processor.ID)
	variable, err := s.ScenarioProcessorRepo.GetVariable(po)
	output, _ = s.ExecHelperService.EvaluateVariable(&variable, parentLog, msg)

	if processor.EntityType == consts.ProcessorVariableSet {
		s.ExecContext.SetVariable(parentLog.ProcessId, output.VariableName, output.VariableValue, false) // set in parent scope

	} else if processor.EntityType == consts.ProcessorVariableClear {
		s.ExecContext.ClearVariable(parentLog.ProcessId, output.VariableName) // set in parent scope
	}

	return
}

func (s *ExecProcessorService) ExecAssertion(processor *agentDomain.Processor, parentLog *domain.ExecLog, msg *websocket.Message) (
	output domain.ExecOutput, err error) {

	po, _ := s.ScenarioProcessorRepo.Get(processor.ID)
	assertion, err := s.ScenarioProcessorRepo.GetAssertion(po)
	output, _ = s.ExecHelperService.EvaluateAssertion(&assertion, parentLog, msg)

	result, _ := s.ExecHelperService.EvaluateGovaluateExpression(output.Expression, processor.ID)
	output.Pass, _ = result.(bool)

	status := "失败"
	if output.Pass {
		status = "通过"
	}
	output.Msg = fmt.Sprintf("表达式\"%s\"结果为\"%s\"。", output.Expression, status)

	return
}

func (s *ExecProcessorService) ExecExtractor(processor *agentDomain.Processor, parentLog *domain.ExecLog, msg *websocket.Message) (
	output domain.ExecOutput, err error) {

	po, _ := s.ScenarioProcessorRepo.Get(processor.ID)
	extractor, err := s.ScenarioProcessorRepo.GetExtractor(po)
	output, _ = s.ExecHelperService.EvaluateExtractor(&extractor, parentLog, msg)

	return
}

func (s *ExecProcessorService) ExecCookie(processor *agentDomain.Processor, parentLog *domain.ExecLog, msg *websocket.Message) (
	output domain.ExecOutput, err error) {

	po, _ := s.ScenarioProcessorRepo.Get(processor.ID)
	cookie, err := s.ScenarioProcessorRepo.GetCookie(po)
	output, _ = s.ExecHelperService.EvaluateCookie(&cookie, parentLog, msg)

	return
}
