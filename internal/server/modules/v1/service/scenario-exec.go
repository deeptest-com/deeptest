package service

import (
	"container/list"
	"encoding/json"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	serverDomain "github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	execHelper "github.com/aaronchen2k/deeptest/internal/server/modules/v1/helper/exec"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
	_stringUtils "github.com/aaronchen2k/deeptest/pkg/lib/string"
	"github.com/jinzhu/copier"
	"github.com/kataras/iris/v12/websocket"
	"time"
)

type ScenarioExecService struct {
	ScenarioProcessorRepo        *repo.ScenarioProcessorRepo   `inject:""`
	ScenarioRepo                 *repo.ScenarioRepo            `inject:""`
	ScenarioNodeRepo             *repo.ScenarioNodeRepo        `inject:""`
	TestResultRepo               *repo.TestResultRepo          `inject:""`
	TestLogRepo                  *repo.TestLogRepo             `inject:""`
	InterfaceRepo                *repo.InterfaceRepo           `inject:""`
	InterfaceService             *InterfaceService             `inject:""`
	ScenarioProcessorExecService *ScenarioProcessorExecService `inject:""`
}

func (s *ScenarioExecService) Load(scenarioId int) (result domain.Result, err error) {
	scenario, err := s.ScenarioRepo.Get(uint(scenarioId))
	if err != nil {
		return
	}

	result.Name = scenario.Name

	return
}

func (s *ScenarioExecService) ExecScenario(scenarioId int, wsMsg websocket.Message) (err error) {
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
	s.ScenarioNodeRepo.GetScopeHierarchy(scenarioId, &execHelper.ScopeHierarchyMap)

	execHelper.IteratorContextStack = list.New()

	rootLog := domain.Log{
		Id:                rootProcessor.ID,
		Name:              rootProcessor.Name,
		ProcessorCategory: rootProcessor.EntityCategory,
		ProcessorType:     rootProcessor.EntityType,
		ParentId:          0,
	}

	execHelper.SendStartMsg(wsMsg)
	execHelper.SendExecMsg(rootLog, wsMsg)

	for _, child := range rootProcessor.Children {
		s.ExecRecursive(child, &rootLog, wsMsg)
	}

	execHelper.SendEndMsg(wsMsg)

	return
}

func (s *ScenarioExecService) ExecRecursive(processor *model.TestProcessor, parentLog *domain.Log,
	wsMsg websocket.Message) (err error) {
	if parentLog.Logs == nil {
		logs := make([]*domain.Log, 0)
		parentLog.Logs = &logs
	}

	if s.isContainerProcessor(processor.EntityCategory) {
		var containerLog *domain.Log
		if s.isExecutableContainerProcessor(processor.EntityCategory) {
			containerLog, _ = s.ExecContainerProcessorWithResp(processor, parentLog, wsMsg)

			if containerLog.Output.Times > 0 { // is loop times
				iterator, _ := execHelper.GenerateLoopTimes(*containerLog)
				execHelper.IteratorContextStack.PushFront(iterator)
				for range iterator.Times {
					s.ExecChildren(processor, containerLog, wsMsg)
				}
				execHelper.IteratorContextStack.Remove(execHelper.IteratorContextStack.Front())
				execHelper.IteratorContextIndex = 0
			}

		} else {
			containerLog, _ = s.AddContainerProcessorWithResp(processor, parentLog, wsMsg)
			s.ExecChildren(processor, containerLog, wsMsg)
		}

	} else if processor.EntityCategory == consts.ProcessorInterface {
		s.ExecInterfaceWithResp(processor, parentLog, wsMsg)
	} else {
		s.ExecActionProcessorWithResp(processor, parentLog, wsMsg)
	}

	return
}

func (s *ScenarioExecService) ExecChildren(processor *model.TestProcessor, parentLog *domain.Log, wsMsg websocket.Message) {
	for _, child := range processor.Children {
		s.ExecRecursive(child, parentLog, wsMsg)
	}
}

func (s *ScenarioExecService) AddContainerProcessorWithResp(processor *model.TestProcessor, parentLog *domain.Log, wsMsg websocket.Message) (
	containerLog *domain.Log, err error) {

	_, desc, _ := execHelper.RetrieveIteratorsVal()

	containerLog = &domain.Log{
		Id:                processor.ID,
		Name:              processor.Name,
		ProcessId:         processor.ID,
		ProcessorCategory: processor.EntityCategory,
		ProcessorType:     processor.EntityType,
		ParentId:          processor.ParentId,
		RespSummary:       []string{desc},
	}

	*parentLog.Logs = append(*parentLog.Logs, containerLog)
	execHelper.SendExecMsg(*containerLog, wsMsg)

	return
}

func (s *ScenarioExecService) ExecContainerProcessorWithResp(processor *model.TestProcessor, parentLog *domain.Log, wsMsg websocket.Message) (
	containerLog *domain.Log, err error) {

	output := domain.Output{}
	//if processor.EntityCategory == consts.ProcessorThreadGroup {
	//	result, _ = s.ScenarioProcessorExecService.ExecThreadGroup(processor, parentLog, wsMsg)
	//} else
	if processor.EntityCategory == consts.ProcessorLogic {
		output, _ = s.ScenarioProcessorExecService.ExecLogic(processor, parentLog, wsMsg)

	} else if processor.EntityCategory == consts.ProcessorLoop {
		output, _ = s.ScenarioProcessorExecService.ExecLoop(processor, parentLog, wsMsg)

	} else if processor.EntityCategory == consts.ProcessorData {
		output, _ = s.ScenarioProcessorExecService.ExecData(processor, parentLog, wsMsg)

	}

	containerLog = &domain.Log{
		Id:                processor.ID,
		Name:              processor.Name,
		ProcessId:         processor.ID,
		ProcessorCategory: processor.EntityCategory,
		ProcessorType:     processor.EntityType,
		ParentId:          processor.ParentId,

		Output:      output,
		RespSummary: []string{output.Text},
	}

	*parentLog.Logs = append(*parentLog.Logs, containerLog)
	execHelper.SendExecMsg(*containerLog, wsMsg)

	return
}

func (s *ScenarioExecService) ExecActionProcessorWithResp(processor *model.TestProcessor, parentLog *domain.Log, wsMsg websocket.Message) (err error) {
	output := domain.Output{}
	if processor.EntityCategory == consts.ProcessorTimer {
		output, _ = s.ScenarioProcessorExecService.ExecTimer(processor, parentLog, wsMsg)

	} else if processor.EntityCategory == consts.ProcessorVariable {
		output, _ = s.ScenarioProcessorExecService.ExecVariable(processor, parentLog, wsMsg)

	} else if processor.EntityCategory == consts.ProcessorAssertion {
		output, _ = s.ScenarioProcessorExecService.ExecAssertion(processor, parentLog, wsMsg)

	} else if processor.EntityCategory == consts.ProcessorExtractor {
		output, _ = s.ScenarioProcessorExecService.ExecExtractor(processor, parentLog, wsMsg)

	} else if processor.EntityCategory == consts.ProcessorCookie {
		output, _ = s.ScenarioProcessorExecService.ExecCookie(processor, parentLog, wsMsg)

	}

	actionLog := &domain.Log{
		Id:                processor.ID,
		Name:              processor.Name,
		ProcessId:         processor.ID,
		ProcessorCategory: processor.EntityCategory,
		ProcessorType:     processor.EntityType,
		ParentId:          processor.ParentId,
		Output:            output,
	}

	*parentLog.Logs = append(*parentLog.Logs, actionLog)
	execHelper.SendExecMsg(*actionLog, wsMsg)

	return
}

func (s *ScenarioExecService) ExecInterfaceWithResp(interfaceProcessor *model.TestProcessor, parentLog *domain.Log, wsMsg websocket.Message) (err error) {
	interf, err := s.InterfaceRepo.Get(interfaceProcessor.InterfaceId)
	if err != nil {
		return
	}

	invocation := serverDomain.InvocationRequest{}
	copier.CopyWithOption(&invocation, interf, copier.Option{DeepCopy: true})
	err = s.InterfaceService.ReplaceVariables(&invocation)
	if err != nil {
		return
	}

	resp, err := s.InterfaceService.Test(invocation)
	if err != nil {
		return
	}

	_, err = s.CreateLog(invocation, resp)
	if err != nil {
		return
	}

	reqContent, _ := json.Marshal(invocation)
	respContent, _ := json.Marshal(resp)

	interfaceLog := &domain.Log{
		Id:                interfaceProcessor.ID,
		Name:              interfaceProcessor.Name,
		ProcessorCategory: consts.ProcessorInterface,
		ProcessorType:     consts.ProcessorInterfaceDefault,
		ParentId:          interfaceProcessor.ParentId,

		InterfaceId: interf.ID,
		ReqContent:  string(reqContent),
		RespContent: string(respContent),
	}

	*parentLog.Logs = append(*parentLog.Logs, interfaceLog)
	execHelper.SendExecMsg(*interfaceLog, wsMsg)

	return
}

func (s *ScenarioExecService) CreateResult(scenario model.TestScenario) (result model.TestResult, err error) {
	startTime := time.Now()
	result = model.TestResult{
		Name:           scenario.Name,
		StartTime:      &startTime,
		ProgressStatus: consts.InProgress,
		ScenarioId:     scenario.ID,
	}

	s.TestResultRepo.Create(&result)

	return
}

func (s *ScenarioExecService) RestartResult(result *model.TestResult, scenario model.TestScenario) (err error) {
	result.Name = scenario.Name

	startTime := time.Now()
	result.StartTime = &startTime

	s.TestResultRepo.ResetResult(*result)
	s.TestResultRepo.ClearLogs(result.ID)

	return
}

func (s *ScenarioExecService) CreateLog(req serverDomain.InvocationRequest, resp serverDomain.InvocationResponse) (
	log model.TestLog, err error) {
	log = model.TestLog{
		Name:        time.Now().Format("01-02 15:04:05"),
		InterfaceId: req.Id,
	}

	bytesReq, _ := json.Marshal(req)
	log.ReqContent = string(bytesReq)

	bytesReps, _ := json.Marshal(resp)
	log.RespContent = string(bytesReps)

	err = s.TestLogRepo.Save(&log)

	return
}

func (s *ScenarioExecService) isContainerProcessor(category consts.ProcessorCategory) bool {
	arr := []string{
		consts.ProcessorRoot.ToString(),
		//consts.ProcessorThreadGroup.ToString(),
		consts.ProcessorGroup.ToString(),
		consts.ProcessorLogic.ToString(),
		consts.ProcessorLoop.ToString(),
		consts.ProcessorData.ToString(),
	}
	return _stringUtils.FindInArr(category.ToString(), arr)
}

func (s *ScenarioExecService) isExecutableContainerProcessor(category consts.ProcessorCategory) bool {
	arr := []string{
		//consts.ProcessorThreadGroup.ToString(),
		consts.ProcessorLogic.ToString(),
		consts.ProcessorLoop.ToString(),
		consts.ProcessorData.ToString(),
	}
	return _stringUtils.FindInArr(category.ToString(), arr)
}

func (s *ScenarioExecService) isActionProcessor(category consts.ProcessorCategory) bool {
	arr := []string{
		consts.ProcessorTimer.ToString(),
		consts.ProcessorVariable.ToString(),
		consts.ProcessorAssertion.ToString(),
		consts.ProcessorExtractor.ToString(),
		consts.ProcessorCookie.ToString(),
	}
	return _stringUtils.FindInArr(category.ToString(), arr)
}

func (s *ScenarioExecService) CancelAndSendMsg(scenarioId int, wsMsg websocket.Message) (err error) {
	s.TestResultRepo.UpdateStatus(consts.Cancel, "", uint(scenarioId))
	execHelper.SendCancelMsg(wsMsg)
	return
}
