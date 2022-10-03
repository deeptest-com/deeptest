package service

import (
	"encoding/json"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	agentDomain "github.com/aaronchen2k/deeptest/internal/agent/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	model2 "github.com/aaronchen2k/deeptest/internal/server/modules/model"
	repo2 "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"strings"
)

type ExecLogService struct {
	ScenarioProcessorRepo *repo2.ScenarioProcessorRepo `inject:""`
	ScenarioRepo          *repo2.ScenarioRepo          `inject:""`
	TestResultRepo        *repo2.ReportRepo            `inject:""`
	TestLogRepo           *repo2.LogRepo               `inject:""`
	InterfaceRepo         *repo2.InterfaceRepo         `inject:""`
	InterfaceService      *InterfaceService            `inject:""`
}

func (s *ExecLogService) CreateProcessorLog(processor *agentDomain.Processor, log *domain.ExecLog, parentPersistentId uint) (po model2.ExecLogProcessor, err error) {
	po = model2.ExecLogProcessor{
		Name:              processor.Name,
		ProcessorCategory: processor.EntityCategory,
		ProcessorType:     processor.EntityType,
		ProcessorId:       processor.ID,

		ParentId: parentPersistentId,
		ReportId: log.ReportId,
	}

	po.Summary = strings.Join(log.Summary, "; ")

	outputBytes, _ := json.Marshal(log.Output)
	po.Output = string(outputBytes)

	err = s.TestLogRepo.Save(&po)
	log.Id = po.ID
	log.PersistentId = po.ID

	return
}

func (s *ExecLogService) CreateInterfaceLog(req v1.InvocationRequest, resp v1.InvocationResponse, parentLog *domain.ExecLog) (
	po model2.ExecLogProcessor, err error) {
	po = model2.ExecLogProcessor{
		Name:              req.Name,
		ProcessorCategory: consts.ProcessorInterface,
		ProcessorType:     consts.ProcessorInterfaceDefault,
		ResultStatus:      consts.Pass, // TODO:
		InterfaceId:       req.Id,

		ProcessorId: parentLog.ProcessId,
		ParentId:    parentLog.PersistentId,
		ReportId:    parentLog.ReportId,
	}

	bytesReq, _ := json.Marshal(req)
	po.ReqContent = string(bytesReq)

	bytesReps, _ := json.Marshal(resp)
	po.RespContent = string(bytesReps)

	err = s.TestLogRepo.Save(&po)

	return
}
