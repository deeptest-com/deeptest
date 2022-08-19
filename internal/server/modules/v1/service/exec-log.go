package service

import (
	"encoding/json"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/business"
	serverDomain "github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
	"time"
)

type ExecLogService struct {
	ScenarioProcessorRepo *repo.ScenarioProcessorRepo  `inject:""`
	ScenarioRepo          *repo.ScenarioRepo           `inject:""`
	TestResultRepo        *repo.TestResultRepo         `inject:""`
	TestLogRepo           *repo.TestLogRepo            `inject:""`
	InterfaceRepo         *repo.InterfaceRepo          `inject:""`
	InterfaceService      *InterfaceService            `inject:""`
	ExecRequestService    *business.ExecRequestService `inject:""`
}

func (s *ExecLogService) CreateInterfaceProcessorLog(req serverDomain.InvocationRequest, resp serverDomain.InvocationResponse) (
	log model.Log, err error) {
	log = model.Log{
		Name:              time.Now().Format("01-02 15:04:05"),
		ProcessorCategory: consts.ProcessorInterface,
		ProcessorType:     consts.ProcessorInterfaceDefault,
		InterfaceId:       req.Id,
	}

	bytesReq, _ := json.Marshal(req)
	log.ReqContent = string(bytesReq)

	bytesReps, _ := json.Marshal(resp)
	log.RespContent = string(bytesReps)

	err = s.TestLogRepo.Save(&log)

	return
}
