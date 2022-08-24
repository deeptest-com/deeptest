package service

import (
	"encoding/json"
	"fmt"
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

	// TODO: save to
	s.ExtractorService.ExtractByInterface(interf.ID, resp, interf.ProjectId)
	s.CheckpointService.CheckByInterface(interf.ID, resp, interf.ProjectId)

	logPo, err := s.ExecLogService.CreateInterfaceLog(req, resp, parentLog)
	if err != nil {
		return
	}

	// TODO: set checkpoint results to interface log

	// send msg to client
	reqContent, _ := json.Marshal(req)
	respContent, _ := json.Marshal(resp)

	interfaceLog := &domain.Log{
		Id:                logPo.ID,
		Name:              fmt.Sprintf("%s - %s %s", interfaceProcessor.Name, req.Method, req.Url),
		ProcessorCategory: consts.ProcessorInterface,
		ProcessorType:     consts.ProcessorInterfaceDefault,
		ParentId:          parentLog.PersistentId,

		InterfaceId:  interf.ID,
		ReqContent:   string(reqContent),
		RespContent:  string(respContent),
		ResultStatus: consts.Pass,
	}

	*parentLog.Logs = append(*parentLog.Logs, interfaceLog)
	execHelper.SendExecMsg(*interfaceLog, wsMsg)

	return
}
