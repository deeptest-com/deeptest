package service

import (
	"encoding/json"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/business"
	serverDomain "github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
	"strings"
)

type ExecLogService struct {
	ScenarioProcessorRepo *repo.ScenarioProcessorRepo  `inject:""`
	ScenarioRepo          *repo.ScenarioRepo           `inject:""`
	TestResultRepo        *repo.ReportRepo             `inject:""`
	TestLogRepo           *repo.LogRepo                `inject:""`
	InterfaceRepo         *repo.InterfaceRepo          `inject:""`
	InterfaceService      *InterfaceService            `inject:""`
	ExecRequestService    *business.ExecRequestService `inject:""`
}

func (s *ExecLogService) CreateProcessorLog(processor *model.Processor, log *domain.Log, parentPersistentId uint) (po model.Log, err error) {
	po = model.Log{
		Name:              processor.Name,
		ProcessorCategory: processor.EntityCategory,
		ProcessorType:     processor.EntityType,
		ProcessorId:       processor.ID,

		ParentId: parentPersistentId,
		ReportId: log.ResultId,
	}

	po.Summary = strings.Join(log.Summary, "; ")

	outputBytes, _ := json.Marshal(log.Output)
	po.Output = string(outputBytes)

	err = s.TestLogRepo.Save(&po)
	log.Id = po.ID
	log.PersistentId = po.ID

	return
}

func (s *ExecLogService) CreateInterfaceLog(req serverDomain.InvocationRequest, resp serverDomain.InvocationResponse, parentLog *domain.Log) (
	po model.Log, err error) {
	po = model.Log{
		Name:              req.Name,
		ProcessorCategory: consts.ProcessorInterface,
		ProcessorType:     consts.ProcessorInterfaceDefault,
		InterfaceId:       req.Id,

		ParentId: parentLog.PersistentId,
		ReportId: parentLog.ResultId,
	}

	bytesReq, _ := json.Marshal(req)
	po.ReqContent = string(bytesReq)

	bytesReps, _ := json.Marshal(resp)
	po.RespContent = string(bytesReps)

	err = s.TestLogRepo.Save(&po)

	return
}
