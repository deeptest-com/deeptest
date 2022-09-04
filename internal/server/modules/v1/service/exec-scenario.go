package service

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/business"
	execHelper "github.com/aaronchen2k/deeptest/internal/server/modules/v1/helper/exec"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/jinzhu/copier"
	"github.com/kataras/iris/v12/websocket"
	"time"
)

var (
	breakOnce = false
)

type ExecScenarioService struct {
	ScenarioProcessorRepo *repo.ScenarioProcessorRepo `inject:""`
	ScenarioRepo          *repo.ScenarioRepo          `inject:""`
	ScenarioNodeRepo      *repo.ScenarioNodeRepo      `inject:""`
	TestResultRepo        *repo.ReportRepo            `inject:""`
	TestLogRepo           *repo.LogRepo               `inject:""`
	InterfaceRepo         *repo.InterfaceRepo         `inject:""`
	InterfaceService      *InterfaceService           `inject:""`
	ExecProcessorService  *ExecProcessorService       `inject:""`

	ExecContextService   *business.ExecContext  `inject:""`
	ExecComm             *business.ExecComm     `inject:""`
	ExecHelperService    *ExecHelperService     `inject:""`
	ExecIteratorService  *business.ExecIterator `inject:""`
	ExecRequestService   *business.ExecRequest  `inject:""`
	ExecLogService       *ExecLogService        `inject:""`
	ExecReportService    *ExecReportService     `inject:""`
	ExecInterfaceService *ExecInterfaceService  `inject:""`
}

func (s *ExecScenarioService) Load(scenarioId int) (result domain.Report, err error) {
	scenario, err := s.ScenarioRepo.Get(uint(scenarioId))
	if err != nil {
		return
	}

	result.Name = scenario.Name

	return
}

func (s *ExecScenarioService) ExecScenario(scenarioId int, wsMsg websocket.Message) (err error) {
	scenario, err := s.ScenarioRepo.Get(uint(scenarioId))
	if err != nil {
		return
	}

	resultPo, err := s.TestResultRepo.FindInProgressResult(uint(scenarioId))
	if resultPo.ID > 0 {
		s.RestartResult(&resultPo, scenario)
	} else {
		resultPo, _ = s.CreateResult(scenario)
	}

	rootProcessor, err := s.ScenarioNodeRepo.GetTree(scenarioId)

	s.ExecContextService.InitScopeHierarchy(uint(scenarioId))
	s.ExecIteratorService.InitIteratorContext()

	rootLog := domain.ExecLog{
		Id:                rootProcessor.ID,
		Name:              rootProcessor.Name,
		ProcessorCategory: rootProcessor.EntityCategory,
		ProcessorType:     rootProcessor.EntityType,
		ParentId:          0,
		ReportId:          resultPo.ID,
	}

	execHelper.SendStartMsg(wsMsg)
	execHelper.SendExecMsg(rootLog, wsMsg)

	for _, child := range rootProcessor.Children {
		s.ExecProcessorRecursively(child, &rootLog, wsMsg)
	}

	execHelper.SendEndMsg(wsMsg)

	report := s.ExecReportService.UpdateTestReport(rootLog)
	reportTo := domain.ReportSimple{}
	copier.CopyWithOption(&reportTo, report, copier.Option{DeepCopy: true})
	reportTo.StartTime = report.StartTime
	reportTo.EndTime = report.EndTime

	execHelper.SendResultMsg(reportTo, wsMsg)

	return
}

func (s *ExecScenarioService) ExecProcessorRecursively(processor *model.Processor, parentLog *domain.ExecLog,
	wsMsg websocket.Message) (err error) {
	if parentLog.Logs == nil {
		logs := make([]*domain.ExecLog, 0)
		parentLog.Logs = &logs
	}

	if s.ExecComm.IsWrapperProcessor(processor.EntityCategory) {
		if s.ExecComm.IsExecutableWrapperProcessor(processor.EntityCategory) {
			s.ExecWrapperProcessor(processor, parentLog, wsMsg)

		} else {
			wrapperLog, _ := s.AddWrapperProcessor(processor, parentLog, wsMsg)

			s.ExecChildren(processor, wrapperLog, wsMsg)

		}

	} else if processor.EntityCategory == consts.ProcessorInterface {
		s.ExecInterfaceService.ExecInterfaceProcessor(processor, parentLog, wsMsg)

	} else {
		s.ExecActionProcessorWithResp(processor, parentLog, wsMsg)

	}

	return
}

func (s *ExecScenarioService) ExecWrapperProcessor(processor *model.Processor, parentLog *domain.ExecLog,
	wsMsg websocket.Message) {
	wrapperLog, _ := s.GetWrapperProcessorResp(processor, parentLog, wsMsg)

	if s.ExecComm.IsLoop(wrapperLog) {
		s.ExecWrapperLoopProcessor(processor, wrapperLog, wsMsg)

	} else if s.ExecComm.IsLogicPass(wrapperLog) {
		s.ExecChildren(processor, wrapperLog, wsMsg)
	}
}

func (s *ExecScenarioService) ExecActionProcessorWithResp(processor *model.Processor, parentLog *domain.ExecLog, wsMsg websocket.Message) (
	wrapperLog *domain.ExecLog, err error) {

	output := domain.ExecOutput{}
	if processor.EntityCategory == consts.ProcessorTimer {
		output, _ = s.ExecProcessorService.ExecTimer(processor, parentLog, wsMsg)

	} else if processor.EntityCategory == consts.ProcessorVariable {
		output, _ = s.ExecProcessorService.ExecVariable(processor, parentLog, wsMsg)

	} else if processor.EntityCategory == consts.ProcessorAssertion {
		output, _ = s.ExecProcessorService.ExecAssertion(processor, parentLog, wsMsg)

	} else if processor.EntityCategory == consts.ProcessorExtractor {
		output, _ = s.ExecProcessorService.ExecExtractor(processor, parentLog, wsMsg)

	} else if processor.EntityCategory == consts.ProcessorCookie {
		output, _ = s.ExecProcessorService.ExecCookie(processor, parentLog, wsMsg)

	}

	wrapperLog, _ = s.wrapperLogAndSendMsg(output, processor, parentLog, wsMsg)

	return
}

func (s *ExecScenarioService) ExecWrapperLoopProcessor(processor *model.Processor, wrapperLog *domain.ExecLog,
	wsMsg websocket.Message) {
	if s.ExecComm.IsLoopTimesPass(wrapperLog) {
		iterator, _ := s.ExecIteratorService.GenerateLoopTimes(*wrapperLog)

		s.ExecIteratorService.Push(iterator)

		for range iterator.Times {
			wrapperLogItem, _ := s.AddWrapperProcessor(processor, wrapperLog, wsMsg)

			s.ExecChildren(processor, wrapperLogItem, wsMsg)

			if breakOnce {
				breakOnce = false
				break
			}
		}

		s.ExecIteratorService.Pop()

	} else if s.ExecComm.IsLoopUntilPass(wrapperLog) {
		expression := wrapperLog.Output.Expression
		for {
			result, err := s.ExecHelperService.ComputerExpress(expression, wrapperLog.ProcessId)
			pass, ok := result.(bool)
			if err != nil || !ok || !pass {
				break
			}

			s.ExecChildren(processor, wrapperLog, wsMsg)

			if breakOnce {
				breakOnce = false
				break
			}
		}

	} else if s.ExecComm.IsLoopInPass(wrapperLog) {
		loopListProcessor, _ := s.ScenarioProcessorRepo.GetLoop(*processor)
		iterator, _ := s.ExecIteratorService.GenerateLoopList(*wrapperLog)

		s.ExecIteratorService.Push(iterator)

		for _, item := range iterator.Items {
			wrapperLogItem, _ := s.AddWrapperProcessor(processor, wrapperLog, wsMsg)

			s.ExecContextService.SetVariable(processor.ID, loopListProcessor.VariableName, item)
			vari, _ := s.ExecContextService.GetVariable(processor.ID, loopListProcessor.VariableName)
			logUtils.Infof("%s = %v", vari.Name, vari.Value)

			s.ExecChildren(processor, wrapperLogItem, wsMsg)

			if breakOnce {
				breakOnce = false
				break
			}
		}

		s.ExecIteratorService.Pop()

	} else if s.ExecComm.IsLoopRangePass(wrapperLog) {
		loopRangeProcessor, _ := s.ScenarioProcessorRepo.GetLoop(*processor)
		iterator, _ := s.ExecIteratorService.GenerateLoopRange(*wrapperLog, loopRangeProcessor.Step, loopRangeProcessor.IsRand)

		s.ExecIteratorService.Push(iterator)

		for _, item := range iterator.Items {
			wrapperLogItem, _ := s.AddWrapperProcessor(processor, wrapperLog, wsMsg)

			s.ExecContextService.SetVariable(processor.ID, loopRangeProcessor.VariableName, item)
			vari, _ := s.ExecContextService.GetVariable(processor.ID, loopRangeProcessor.VariableName)
			logUtils.Infof("%s = %v", vari.Name, vari.Value)

			s.ExecChildren(processor, wrapperLogItem, wsMsg)

			if breakOnce {
				breakOnce = false
				break
			}
		}

		s.ExecIteratorService.Pop()
	} else if s.ExecComm.IsLoopLoopBreak(wrapperLog) {
		breakIfExpress := wrapperLog.Output.Expression

		result, err := s.ExecHelperService.ComputerExpress(breakIfExpress, wrapperLog.ProcessId)
		pass, ok := result.(bool)
		if err != nil || !ok || !pass {
			breakOnce = true
		}
	} else if s.ExecComm.IsDataPass(wrapperLog) {
		data, _ := s.ScenarioProcessorRepo.GetData(*processor)

		iterator, _ := s.ExecIteratorService.GenerateData(*wrapperLog, data)

		s.ExecIteratorService.Push(iterator)

		for _, mapItem := range iterator.Items {
			wrapperLogItem, _ := s.AddWrapperProcessor(processor, wrapperLog, wsMsg)

			s.ExecContextService.SetVariable(processor.ID, data.VariableName, mapItem)
			vari, _ := s.ExecContextService.GetVariable(processor.ID, data.VariableName)
			logUtils.Infof("%s = %v", vari.Name, vari.Value)

			s.ExecChildren(processor, wrapperLogItem, wsMsg)

			if breakOnce {
				breakOnce = false
				break
			}
		}

		s.ExecIteratorService.Pop()
	}
}

func (s *ExecScenarioService) ExecChildren(processor *model.Processor, parentLog *domain.ExecLog, wsMsg websocket.Message) {
	for _, child := range processor.Children {
		s.ExecProcessorRecursively(child, parentLog, wsMsg)
	}
}

func (s *ExecScenarioService) AddWrapperProcessor(processor *model.Processor, parentLog *domain.ExecLog, wsMsg websocket.Message) (
	wrapperLog *domain.ExecLog, err error) {

	_, desc, _ := s.ExecIteratorService.RetrieveIteratorsVal(processor)

	wrapperLog = &domain.ExecLog{
		Name:              processor.Name,
		ProcessId:         processor.ID,
		ProcessorCategory: processor.EntityCategory,
		ProcessorType:     processor.EntityType,
		ParentId:          parentLog.PersistentId,
		Summary:           []string{desc},
		ReportId:          parentLog.ReportId,
	}

	s.ExecLogService.CreateProcessorLog(processor, wrapperLog, parentLog.PersistentId)

	*parentLog.Logs = append(*parentLog.Logs, wrapperLog)
	execHelper.SendExecMsg(*wrapperLog, wsMsg)

	return
}

func (s *ExecScenarioService) GetWrapperProcessorResp(processor *model.Processor, parentLog *domain.ExecLog, wsMsg websocket.Message) (
	wrapperLog *domain.ExecLog, err error) {

	output := domain.ExecOutput{}

	//if processor.EntityCategory == consts.ProcessorThreadGroup {
	//	result, _ = s.ExecLogService.ExecThreadGroup(processor, parentLog, wsMsg)
	//} else
	if processor.EntityCategory == consts.ProcessorLogic {
		output, _ = s.ExecProcessorService.ExecLogic(processor, parentLog, wsMsg)

	} else if processor.EntityCategory == consts.ProcessorLoop {
		output, _ = s.ExecProcessorService.ExecLoop(processor, parentLog, wsMsg)

	} else if processor.EntityCategory == consts.ProcessorData {
		output, _ = s.ExecProcessorService.ExecData(processor, parentLog, wsMsg)

	}

	wrapperLog, _ = s.wrapperLogAndSendMsg(output, processor, parentLog, wsMsg)

	return
}

func (s *ExecScenarioService) wrapperLogAndSendMsg(output domain.ExecOutput, processor *model.Processor, parentLog *domain.ExecLog, wsMsg websocket.Message) (
	wrapperLog *domain.ExecLog, err error) {
	wrapperLog = &domain.ExecLog{
		Id:                processor.ID,
		Name:              processor.Name,
		ProcessId:         processor.ID,
		ProcessorCategory: processor.EntityCategory,
		ProcessorType:     processor.EntityType,
		ParentId:          parentLog.PersistentId,
		ReportId:          parentLog.ReportId,

		Output:  output,
		Summary: []string{output.Msg},
	}

	logs := make([]*domain.ExecLog, 0)
	wrapperLog.Logs = &logs

	s.ExecLogService.CreateProcessorLog(processor, wrapperLog, parentLog.PersistentId)

	*parentLog.Logs = append(*parentLog.Logs, wrapperLog)
	execHelper.SendExecMsg(*wrapperLog, wsMsg)

	return
}

func (s *ExecScenarioService) CreateResult(scenario model.Scenario) (result model.Report, err error) {
	startTime := time.Now()
	result = model.Report{
		Name:           scenario.Name,
		StartTime:      &startTime,
		ProgressStatus: consts.InProgress,
		ScenarioId:     scenario.ID,
	}

	s.TestResultRepo.Create(&result)

	return
}

func (s *ExecScenarioService) RestartResult(result *model.Report, scenario model.Scenario) (err error) {
	result.Name = scenario.Name

	startTime := time.Now()
	result.StartTime = &startTime

	s.TestResultRepo.ResetResult(*result)
	s.TestResultRepo.ClearLogs(result.ID)

	return
}

func (s *ExecScenarioService) CancelAndSendMsg(scenarioId int, wsMsg websocket.Message) (err error) {
	s.TestResultRepo.UpdateStatus(consts.Cancel, "", uint(scenarioId))
	execHelper.SendCancelMsg(wsMsg)
	return
}
