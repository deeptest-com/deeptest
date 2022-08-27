package service

import (
	"encoding/json"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/business"
	serverDomain "github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	execHelper "github.com/aaronchen2k/deeptest/internal/server/modules/v1/helper/exec"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
	"github.com/jinzhu/copier"
	"github.com/kataras/iris/v12/websocket"
)

type ExecInterfaceService struct {
	InterfaceRepo      *repo.InterfaceRepo          `inject:""`
	InterfaceService   *InterfaceService            `inject:""`
	ExecRequestService *business.ExecRequestService `inject:""`
	ExecLogService     *ExecLogService              `inject:""`

	ExtractorService  *ExtractorService  `inject:""`
	CheckpointService *CheckpointService `inject:""`
}

func (s *ExecInterfaceService) ExecInterfaceProcessor(interfaceProcessor *model.Processor, parentLog *domain.Log, wsMsg websocket.Message) (err error) {
	interf, err := s.InterfaceRepo.Get(interfaceProcessor.InterfaceId)
	if err != nil {
		return
	}

	req := serverDomain.InvocationRequest{}
	copier.CopyWithOption(&req, interf, copier.Option{DeepCopy: true})

	// replace variables
	newReq, err := s.InterfaceService.ReplaceEnvironmentVariables(req)
	if err != nil {
		return
	}
	err = s.ExecRequestService.ReplaceProcessorVariables(&newReq, interfaceProcessor)
	if err != nil {
		return
	}

	resp, err := s.InterfaceService.Test(newReq)
	if err != nil {
		return
	}

	logPo, err := s.ExecLogService.CreateInterfaceLog(req, resp, parentLog)
	if err != nil {
		return
	}

	logExtractors, err := s.ExtractorService.ExtractInterface(interf, resp, &logPo)
	logCheckpoints, err := s.CheckpointService.CheckInterface(interf, resp, &logPo)

	// send msg to client
	reqContent, _ := json.Marshal(req)
	respContent, _ := json.Marshal(resp)

	interfaceLog := &domain.Log{
		Id:                logPo.ID,
		Name:              interfaceProcessor.Name,
		ProcessorCategory: consts.ProcessorInterface,
		ProcessorType:     consts.ProcessorInterfaceDefault,
		ParentId:          parentLog.PersistentId,

		InterfaceId:  interf.ID,
		ReqContent:   string(reqContent),
		RespContent:  string(respContent),
		ResultStatus: consts.Pass,

		InterfaceExtractorsResult:  logExtractors,
		InterfaceCheckpointsResult: logCheckpoints,
	}

	*parentLog.Logs = append(*parentLog.Logs, interfaceLog)
	execHelper.SendExecMsg(*interfaceLog, wsMsg)

	return
}
