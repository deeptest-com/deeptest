package service

import (
	"encoding/json"
	extractCache "github.com/aaronchen2k/deeptest/internal/pkg/cache/extract"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/business"
	serverDomain "github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	execHelper "github.com/aaronchen2k/deeptest/internal/server/modules/v1/helper/exec"
	requestHelper "github.com/aaronchen2k/deeptest/internal/server/modules/v1/helper/request"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
	"github.com/jinzhu/copier"
	"github.com/kataras/iris/v12/websocket"
)

type ExecInterfaceService struct {
	InterfaceService *InterfaceService `inject:""`
	ExecLogService   *ExecLogService   `inject:""`

	ExtractorService  *ExtractorService     `inject:""`
	CheckpointService *CheckpointService    `inject:""`
	ExecContext       *business.ExecContext `inject:""`

	InterfaceRepo   *repo.InterfaceRepo   `inject:""`
	EnvironmentRepo *repo.EnvironmentRepo `inject:""`
	ExtractorRepo   *repo.ExtractorRepo   `inject:""`
}

func (s *ExecInterfaceService) ExecInterfaceProcessor(interfaceProcessor *model.Processor, parentLog *domain.ExecLog, wsMsg *websocket.Message) (err error) {
	interf, err := s.InterfaceRepo.Get(interfaceProcessor.InterfaceId)
	if err != nil {
		return
	}

	req := serverDomain.InvocationRequest{}
	copier.CopyWithOption(&req, interf, copier.Option{DeepCopy: true})

	// replace variables
	environmentVariables, _ := s.EnvironmentRepo.ListByInterface(req.Id)
	interfaceExtractorVariables := extractCache.GetAll()
	processorExtractorVariables := s.ExecContext.ListVariable(interfaceProcessor.ID)
	variableArr := requestHelper.MergeVariables(environmentVariables, interfaceExtractorVariables, processorExtractorVariables)

	requestHelper.ReplaceAll(&req, variableArr)

	resp, err := s.InterfaceService.Test(req)
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

	interfaceLog := &domain.ExecLog{
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
