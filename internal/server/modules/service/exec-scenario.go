package service

import (
	"github.com/aaronchen2k/deeptest/internal/agent/run"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	business2 "github.com/aaronchen2k/deeptest/internal/server/modules/business"
	"github.com/aaronchen2k/deeptest/internal/server/modules/helper/exec"
	model "github.com/aaronchen2k/deeptest/internal/server/modules/model"
	repo2 "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/jinzhu/copier"
	"github.com/kataras/iris/v12/websocket"
	"sync"
	"time"
)

var (
	breakMap sync.Map
)

type ExecScenarioService struct {
	ScenarioProcessorRepo *repo2.ScenarioProcessorRepo `inject:""`
	ScenarioRepo          *repo2.ScenarioRepo          `inject:""`
	ScenarioNodeRepo      *repo2.ScenarioNodeRepo      `inject:""`
	TestReportRepo        *repo2.ReportRepo            `inject:""`
	TestLogRepo           *repo2.LogRepo               `inject:""`
	InterfaceRepo         *repo2.InterfaceRepo         `inject:""`
	InterfaceService      *InterfaceService            `inject:""`
	ExecProcessorService  *ExecProcessorService        `inject:""`

	ExecContextService   *business2.ExecContext  `inject:""`
	ExecComm             *business2.ExecComm     `inject:""`
	ExecHelperService    *ExecHelperService      `inject:""`
	ExecIteratorService  *business2.ExecIterator `inject:""`
	ExecLogService       *ExecLogService         `inject:""`
	ExecReportService    *ExecReportService      `inject:""`
	ExecInterfaceService *ExecInterfaceService   `inject:""`
}

func (s *ExecScenarioService) Load(scenarioId int) (result domain.Report, err error) {
	scenario, err := s.ScenarioRepo.Get(uint(scenarioId))
	if err != nil {
		return
	}

	result.Name = scenario.Name

	return
}

func (s *ExecScenarioService) RunScenario(scenarioId int, wsMsg *websocket.Message) (err error) {
	po, err := s.ScenarioRepo.Get(uint(scenarioId))
	if err != nil {
		return
	}

	testScenario, err := s.genTestScenario(po)
	if err != nil {
		return
	}

	mainRunner := run.NewMainRunner(nil)
	mainRunner.Run(testScenario)

	return
}
func (s *ExecScenarioService) genTestScenario(po model.Scenario) (ret *run.TestScenario, err error) {
	ret, _ = s.ScenarioNodeRepo.GenTestScenario(po.ID)

	return
}

func (s *ExecScenarioService) ExecScenario(scenarioId int, wsMsg *websocket.Message) (err error) {
	scenario, err := s.ScenarioRepo.Get(uint(scenarioId))
	if err != nil {
		return
	}

	s.genTestScenario(scenario)

	resultPo, err := s.TestReportRepo.FindInProgressResult(uint(scenarioId))
	if resultPo.ID > 0 {
		s.RestartResult(&resultPo, scenario)
	} else {
		resultPo, _ = s.CreateResult(scenario)
	}

	rootProcessor, err := s.ScenarioNodeRepo.GetTree(uint(scenarioId))

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
	wsMsg *websocket.Message) (err error) {
	if parentLog.Logs == nil {
		logs := make([]*domain.ExecLog, 0)
		parentLog.Logs = &logs
	}

	if s.ExecComm.IsExecutableContainerProcessor(processor) {
		s.ExecContainerProcessor(processor, parentLog, wsMsg)

	} else if s.ExecComm.IsNoExecutableContainerProcessor(processor) {
		containerLog, _ := s.AddContainerProcessor(processor, parentLog, wsMsg)
		s.ExecChildren(processor, containerLog, wsMsg)

	} else if s.ExecComm.IsActionProcessor(processor) {
		s.ExecActionProcessorAndDisplay(processor, parentLog, wsMsg)

	} else if s.ExecComm.IsInterfaceProcessor(processor) {
		s.ExecInterfaceService.ExecInterfaceProcessor(processor, parentLog, wsMsg)

	}

	return
}

func (s *ExecScenarioService) ExecContainerProcessor(processor *model.Processor, parentLog *domain.ExecLog,
	wsMsg *websocket.Message) {
	containerLog, _ := s.GenerateContainerProcessorLogAndDisplay(processor, parentLog, wsMsg)

	if s.ExecComm.IsLoopPass(containerLog) {
		s.ExecContainerProcessorChildrenForLoop(processor, containerLog, wsMsg)

	} else if s.ExecComm.IsDataPass(containerLog) {
		s.ExecContainerProcessorChildrenForData(processor, containerLog, wsMsg)

	} else if s.ExecComm.IsLogicPass(containerLog) {
		s.ExecChildren(processor, containerLog, wsMsg)

	}
}

func (s *ExecScenarioService) ExecContainerProcessorChildrenForLoop(processor *model.Processor, containerLog *domain.ExecLog,
	wsMsg *websocket.Message) {
	if s.ExecComm.IsLoopTimesPass(containerLog) {
		iterator, _ := s.ExecIteratorService.GenerateLoopTimes(*containerLog)

		s.ExecIteratorService.Push(iterator)

		for range iterator.Times {
			containerLogItem, _ := s.AddContainerProcessor(processor, containerLog, wsMsg)

			s.ExecChildren(processor, containerLogItem, wsMsg)

			toBreak, ok := breakMap.Load(processor.ID)
			if ok && toBreak.(bool) {
				breakMap.Delete(processor.ID)
				break
			}
		}

		s.ExecIteratorService.Pop()

	} else if s.ExecComm.IsLoopUntilPass(containerLog) {
		expression := containerLog.Output.Expression
		for {
			result, err := s.ExecHelperService.EvaluateGovaluateExpression(expression, containerLog.ProcessId)
			pass, ok := result.(bool)
			if err != nil || !ok || pass {
				break
			}

			s.ExecChildren(processor, containerLog, wsMsg)

			toBreak, ok := breakMap.Load(processor.ID)
			if ok && toBreak.(bool) {
				breakMap.Delete(processor.ID)
				break
			}
		}

	} else if s.ExecComm.IsLoopInPass(containerLog) {
		loopListProcessor, _ := s.ScenarioProcessorRepo.GetLoop(*processor)
		iterator, _ := s.ExecIteratorService.GenerateLoopList(*containerLog)

		s.ExecIteratorService.Push(iterator)

		for _, item := range iterator.Items {
			containerLogItem, _ := s.AddContainerProcessor(processor, containerLog, wsMsg)

			s.ExecContextService.SetVariable(processor.ID, loopListProcessor.VariableName, item, false)
			vari, _ := s.ExecContextService.GetVariable(processor.ID, loopListProcessor.VariableName)
			logUtils.Infof("%s = %v", vari.Name, vari.Value)

			s.ExecChildren(processor, containerLogItem, wsMsg)

			toBreak, ok := breakMap.Load(processor.ID)
			if ok && toBreak.(bool) {
				breakMap.Delete(processor.ID)
				break
			}
		}

		s.ExecIteratorService.Pop()

	} else if s.ExecComm.IsLoopRangePass(containerLog) {
		loopRangeProcessor, _ := s.ScenarioProcessorRepo.GetLoop(*processor)
		iterator, _ := s.ExecIteratorService.GenerateLoopRange(*containerLog, loopRangeProcessor.Step, loopRangeProcessor.IsRand)

		s.ExecIteratorService.Push(iterator)

		for _, item := range iterator.Items {
			containerLogItem, _ := s.AddContainerProcessor(processor, containerLog, wsMsg)

			s.ExecContextService.SetVariable(processor.ID, loopRangeProcessor.VariableName, item, false)
			vari, _ := s.ExecContextService.GetVariable(processor.ID, loopRangeProcessor.VariableName)
			logUtils.Infof("%s = %v", vari.Name, vari.Value)

			s.ExecChildren(processor, containerLogItem, wsMsg)

			toBreak, ok := breakMap.Load(processor.ID)
			if ok && toBreak.(bool) {
				breakMap.Delete(processor.ID)
				break
			}
		}

		s.ExecIteratorService.Pop()
	}
}

func (s *ExecScenarioService) ExecActionProcessorAndDisplay(processor *model.Processor, parentLog *domain.ExecLog, wsMsg *websocket.Message) (
	containerLog *domain.ExecLog, err error) {

	output := domain.ExecOutput{}
	if processor.EntityCategory == consts.ProcessorVariable {
		output, _ = s.ExecProcessorService.ExecVariable(processor, parentLog, wsMsg)

	} else if processor.EntityCategory == consts.ProcessorAssertion {
		output, _ = s.ExecProcessorService.ExecAssertion(processor, parentLog, wsMsg)

	} else if processor.EntityCategory == consts.ProcessorExtractor {
		output, _ = s.ExecProcessorService.ExecExtractor(processor, parentLog, wsMsg)

	} else if processor.EntityCategory == consts.ProcessorCookie {
		output, _ = s.ExecProcessorService.ExecCookie(processor, parentLog, wsMsg)

	} else if processor.EntityCategory == consts.ProcessorTimer {
		output, _ = s.ExecProcessorService.ExecTimer(processor, parentLog, wsMsg)

	} else if processor.EntityCategory == consts.ProcessorPrint {
		output, _ = s.ExecProcessorService.ExecPrint(processor, parentLog, wsMsg)

	} else if processor.EntityType == consts.ProcessorLoopBreak {
		output, _ = s.ExecProcessorService.ExecLoopBreak(processor, parentLog, wsMsg)
	}

	containerLog, _ = s.generateContainerLogAndSendMsg(output, processor, parentLog, wsMsg)

	return
}

func (s *ExecScenarioService) ExecContainerProcessorChildrenForData(processor *model.Processor, containerLog *domain.ExecLog,
	wsMsg *websocket.Message) {
	if s.ExecComm.IsDataPass(containerLog) {
		data, _ := s.ScenarioProcessorRepo.GetData(*processor)

		iterator, _ := s.ExecIteratorService.GenerateData(*containerLog, data)

		s.ExecIteratorService.Push(iterator)

		for _, mapItem := range iterator.Data {
			containerLogItem, _ := s.AddContainerProcessor(processor, containerLog, wsMsg)

			s.ExecContextService.SetVariable(processor.ID, data.VariableName, mapItem, false)
			//vari, _ := s.ExecContextService.GetVariable(processor.ID, data.VariableName)
			//logUtils.Infof("%s = %v", vari.Name, vari.Value)

			s.ExecChildren(processor, containerLogItem, wsMsg)

			toBreak, ok := breakMap.Load(processor.ID)
			if ok && toBreak.(bool) {
				breakMap.Delete(processor.ID)
				break
			}
		}

		s.ExecIteratorService.Pop()
	}
}

func (s *ExecScenarioService) ExecChildren(processor *model.Processor, parentLog *domain.ExecLog, wsMsg *websocket.Message) {
	for _, child := range processor.Children {
		s.ExecProcessorRecursively(child, parentLog, wsMsg)
	}
}

func (s *ExecScenarioService) AddContainerProcessor(processor *model.Processor, parentLog *domain.ExecLog, wsMsg *websocket.Message) (
	containerLog *domain.ExecLog, err error) {

	_, desc, _ := s.ExecIteratorService.RetrieveIteratorsVal(processor)

	containerLog = &domain.ExecLog{
		Name:              processor.Name,
		ProcessId:         processor.ID,
		ProcessorCategory: processor.EntityCategory,
		ProcessorType:     processor.EntityType,
		ParentId:          parentLog.PersistentId,
		Summary:           []string{desc},
		ReportId:          parentLog.ReportId,
	}

	s.ExecLogService.CreateProcessorLog(processor, containerLog, parentLog.PersistentId)

	*parentLog.Logs = append(*parentLog.Logs, containerLog)
	execHelper.SendExecMsg(*containerLog, wsMsg)

	return
}

func (s *ExecScenarioService) GenerateContainerProcessorLogAndDisplay(processor *model.Processor, parentLog *domain.ExecLog, wsMsg *websocket.Message) (
	containerLog *domain.ExecLog, err error) {

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

	containerLog, _ = s.generateContainerLogAndSendMsg(output, processor, parentLog, wsMsg)

	return
}

func (s *ExecScenarioService) generateContainerLogAndSendMsg(output domain.ExecOutput, processor *model.Processor, parentLog *domain.ExecLog, wsMsg *websocket.Message) (
	containerLog *domain.ExecLog, err error) {
	containerLog = &domain.ExecLog{
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
	containerLog.Logs = &logs

	s.ExecLogService.CreateProcessorLog(processor, containerLog, parentLog.PersistentId)

	*parentLog.Logs = append(*parentLog.Logs, containerLog)

	execHelper.SendExecMsg(*containerLog, wsMsg)

	return
}

func (s *ExecScenarioService) CreateResult(scenario model.Scenario) (result model.Report, err error) {
	startTime := time.Now()
	result = model.Report{
		Name:           scenario.Name,
		StartTime:      &startTime,
		ProgressStatus: consts.InProgress,
		ScenarioId:     scenario.ID,
		ProjectId:      scenario.ProjectId,
	}

	s.TestReportRepo.Create(&result)

	return
}

func (s *ExecScenarioService) RestartResult(report *model.Report, scenario model.Scenario) (err error) {
	report.Name = scenario.Name

	startTime := time.Now()
	report.StartTime = &startTime

	s.TestReportRepo.ResetResult(*report)
	s.TestReportRepo.ClearLogs(report.ID)

	return
}

func (s *ExecScenarioService) CancelAndSendMsg(scenarioId int, wsMsg websocket.Message) (err error) {
	s.TestReportRepo.UpdateStatus(consts.Cancel, "", uint(scenarioId))
	execHelper.SendCancelMsg(wsMsg)
	return
}
