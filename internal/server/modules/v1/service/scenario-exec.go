package service

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	execHelper "github.com/aaronchen2k/deeptest/internal/server/modules/v1/helper/exec"
	websocketHelper "github.com/aaronchen2k/deeptest/internal/server/modules/v1/helper/websocket"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
	_i118Utils "github.com/aaronchen2k/deeptest/pkg/lib/i118"
	_logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	_stringUtils "github.com/aaronchen2k/deeptest/pkg/lib/string"
	"github.com/jinzhu/copier"
	"github.com/kataras/iris/v12/websocket"
	"time"
)

type ScenarioExecService struct {
	ScenarioProcessorRepo *repo.ScenarioProcessorRepo `inject:""`
	ScenarioRepo          *repo.ScenarioRepo          `inject:""`
	TestResultRepo        *repo.TestResultRepo        `inject:""`
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

	result := s.CopyResult(resultPo)
	s.SendStartMsg(result, wsMsg)

	rootProcessor, err := s.ScenarioProcessorRepo.GetRootProcessor(scenario.ID)
	if err != nil {
		return
	}

	rootLog := domain.Log{
		ID:                rootProcessor.ID,
		Name:              rootProcessor.Name,
		ProcessorCategory: rootProcessor.EntityCategory,
		ProcessorType:     rootProcessor.EntityType,
		ParentId:          0,
	}

	children, _ := s.ScenarioProcessorRepo.GetChildrenProcessor(rootProcessor.ID, rootProcessor.ScenarioId)
	for _, child := range children {
		s.ExecRecursiveProcessor(child, &rootLog)
	}

	return
}

func (s *ScenarioExecService) ExecRecursiveProcessor(processor model.TestProcessor, parentLog *domain.Log) (err error) {
	if parentLog.Logs == nil {
		logs := make([]*domain.Log, 0)
		parentLog.Logs = &logs
	}

	if s.isContainerProcessor(processor.EntityCategory) {
		var containerLog *domain.Log
		if s.isExecutableContainerProcessor(processor.EntityCategory) {
			containerLog, _ = s.ExecContainerProcessor(processor, parentLog)
		} else {
			containerLog, _ = s.AddContainerProcessor(processor, parentLog)
		}

		children, _ := s.ScenarioProcessorRepo.GetChildrenProcessor(processor.ID, processor.ScenarioId)
		for _, child := range children {
			s.ExecRecursiveProcessor(child, containerLog)
		}
	} else if processor.EntityCategory == consts.ProcessorInterface {
		s.ExecInterface(processor, parentLog)
	} else {
		s.ExecActionProcessor(processor, parentLog)
	}

	return
}

func (s *ScenarioExecService) AddContainerProcessor(processor model.TestProcessor, parentLog *domain.Log) (
	containerLog *domain.Log, err error) {

	containerLog = &domain.Log{
		ID:                processor.ID,
		Name:              processor.Name,
		ProcessorCategory: processor.EntityCategory,
		ProcessorType:     processor.EntityType,
		ParentId:          processor.ParentId,
	}

	*parentLog.Logs = append(*parentLog.Logs, containerLog)

	return
}

func (s *ScenarioExecService) ExecContainerProcessor(processor model.TestProcessor, parentLog *domain.Log) (
	containerLog *domain.Log, err error) {

	// TODO: exec

	containerLog = &domain.Log{
		ID:                processor.ID,
		Name:              processor.Name,
		ProcessorCategory: processor.EntityCategory,
		ProcessorType:     processor.EntityType,
		ParentId:          processor.ParentId,
	}

	*parentLog.Logs = append(*parentLog.Logs, containerLog)

	return
}

func (s *ScenarioExecService) ExecActionProcessor(processor model.TestProcessor, parentLog *domain.Log) (err error) {
	// TODO: exec

	actionLog := &domain.Log{
		ID:                processor.ID,
		Name:              processor.Name,
		ProcessorCategory: processor.EntityCategory,
		ProcessorType:     processor.EntityType,
		ParentId:          processor.ParentId,
	}

	*parentLog.Logs = append(*parentLog.Logs, actionLog)

	return
}

func (s *ScenarioExecService) ExecInterface(interf model.TestProcessor, parentLog *domain.Log) (err error) {
	// TODO: exec

	interfaceLog := &domain.Log{
		ID:                interf.ID,
		Name:              interf.Name,
		ProcessorCategory: interf.EntityCategory,
		ProcessorType:     interf.EntityType,
		ParentId:          interf.ParentId,
	}

	*parentLog.Logs = append(*parentLog.Logs, interfaceLog)

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

func (s *ScenarioExecService) SendStartMsg(result domain.Result, wsMsg websocket.Message) (err error) {
	execHelper.SetRunning(true)
	msg := _i118Utils.Sprintf("start_exec")
	websocketHelper.SendExecMsg(msg, result, &wsMsg)
	_logUtils.Infof(msg)

	return
}

func (s *ScenarioExecService) Complete(scenarioId int, wsMsg websocket.Message) (err error) {
	s.TestResultRepo.UpdateStatus(consts.Complete, "", scenarioId)

	execHelper.SetRunning(false)
	msg := _i118Utils.Sprintf("end_exec")
	websocketHelper.SendExecMsg(msg, domain.Result{ProgressStatus: consts.Complete}, &wsMsg)
	_logUtils.Infof(_i118Utils.Sprintf(msg))

	return
}

func (s *ScenarioExecService) CancelAndSendMsg(scenarioId int, wsMsg websocket.Message) (err error) {
	s.TestResultRepo.UpdateStatus(consts.Cancel, "", scenarioId)

	execHelper.SetRunning(false)
	msg := _i118Utils.Sprintf("end_exec")
	websocketHelper.SendExecMsg(msg, domain.Result{ProgressStatus: consts.Cancel}, &wsMsg)
	_logUtils.Infof(_i118Utils.Sprintf(msg))

	return
}

func (s *ScenarioExecService) SendErrorMsg(scenarioId int, wsMsg websocket.Message) (err error) {
	msg := _i118Utils.Sprintf("wrong_req_params", err.Error())
	websocketHelper.SendExecMsg(msg, domain.Result{ProgressStatus: consts.Error}, &wsMsg)
	_logUtils.Infof(msg)

	return
}

func (s *ScenarioExecService) SendAlreadyRunningMsg(scenarioId int, wsMsg websocket.Message) (err error) {
	msg := _i118Utils.Sprintf("pls_stop_previous")
	websocketHelper.SendExecMsg(msg, domain.Result{ProgressStatus: consts.InProgress}, &wsMsg)
	_logUtils.Infof(msg)

	return
}

func (s *ScenarioExecService) CopyResult(result model.TestResult) (to domain.Result) {
	copier.Copy(&to, result)
	return
}
