package service

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/business"
	"github.com/aaronchen2k/deeptest/internal/server/modules/helper/exec"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	repo2 "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/kataras/iris/v12/websocket"
	"time"
)

type ExecProcessorService struct {
	ScenarioProcessorRepo *repo2.ScenarioProcessorRepo `inject:""`
	ScenarioRepo          *repo2.ScenarioRepo          `inject:""`
	TestResultRepo        *repo2.ReportRepo            `inject:""`
	TestLogRepo           *repo2.LogRepo               `inject:""`
	InterfaceRepo         *repo2.InterfaceRepo         `inject:""`
	InterfaceService      *InterfaceService            `inject:""`
	ExecHelperService     *ExecHelperService           `inject:""`
	ExecContext           *business.ExecContext        `inject:""`
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
	output, _ = s.ExecHelperService.EvaluateLogic(&logic, parentLog, msg)

	return
}

func (s *ExecProcessorService) ExecLoop(processor *model.Processor, parentLog *domain.ExecLog, msg *websocket.Message) (
	output domain.ExecOutput, err error) {

	loop, err := s.ScenarioProcessorRepo.GetLoop(*processor)
	output, _ = s.ExecHelperService.EvaluateLoop(&loop, parentLog, msg)

	return
}
func (s *ExecProcessorService) ExecLoopBreak(processor *model.Processor, parentLog *domain.ExecLog, msg *websocket.Message) (
	output domain.ExecOutput, err error) {

	loop, err := s.ScenarioProcessorRepo.GetLoop(*processor)
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

func (s *ExecProcessorService) ExecData(processor *model.Processor, parentLog *domain.ExecLog, msg *websocket.Message) (
	output domain.ExecOutput, err error) {

	data, err := s.ScenarioProcessorRepo.GetData(*processor)
	output, _ = s.ExecHelperService.EvaluateData(&data, parentLog, msg)

	return
}

func (s *ExecProcessorService) ExecTimer(processor *model.Processor, parentLog *domain.ExecLog, msg *websocket.Message) (
	output domain.ExecOutput, err error) {

	timer, err := s.ScenarioProcessorRepo.GetTimer(*processor)
	output, _ = s.ExecHelperService.EvaluateTimer(&timer, parentLog, msg)

	<-time.After(time.Duration(output.SleepTime) * time.Second)

	return
}

func (s *ExecProcessorService) ExecPrint(processor *model.Processor, parentLog *domain.ExecLog, msg *websocket.Message) (
	output domain.ExecOutput, err error) {

	print, err := s.ScenarioProcessorRepo.GetPrint(*processor)
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

func (s *ExecProcessorService) ExecVariable(processor *model.Processor, parentLog *domain.ExecLog, msg *websocket.Message) (
	output domain.ExecOutput, err error) {

	variable, err := s.ScenarioProcessorRepo.GetVariable(*processor)
	output, _ = s.ExecHelperService.EvaluateVariable(&variable, parentLog, msg)

	if processor.EntityType == consts.ProcessorVariableSet {
		s.ExecContext.SetVariable(parentLog.ProcessId, output.VariableName, output.VariableValue, false) // set in parent scope

	} else if processor.EntityType == consts.ProcessorVariableClear {
		s.ExecContext.ClearVariable(parentLog.ProcessId, output.VariableName) // set in parent scope
	}

	return
}

func (s *ExecProcessorService) ExecAssertion(processor *model.Processor, parentLog *domain.ExecLog, msg *websocket.Message) (
	output domain.ExecOutput, err error) {

	assertion, err := s.ScenarioProcessorRepo.GetAssertion(*processor)
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

func (s *ExecProcessorService) ExecExtractor(processor *model.Processor, parentLog *domain.ExecLog, msg *websocket.Message) (
	output domain.ExecOutput, err error) {

	extractor, err := s.ScenarioProcessorRepo.GetExtractor(*processor)
	output, _ = s.ExecHelperService.EvaluateExtractor(&extractor, parentLog, msg)

	return
}

func (s *ExecProcessorService) ExecCookie(processor *model.Processor, parentLog *domain.ExecLog, msg *websocket.Message) (
	output domain.ExecOutput, err error) {

	cookie, err := s.ScenarioProcessorRepo.GetCookie(*processor)
	output, _ = s.ExecHelperService.EvaluateCookie(&cookie, parentLog, msg)

	return
}
