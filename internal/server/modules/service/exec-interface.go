package service

import (
	"encoding/json"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/business"
	"github.com/aaronchen2k/deeptest/internal/server/modules/helper/exec"
	"github.com/aaronchen2k/deeptest/internal/server/modules/helper/request"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	repo2 "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/jinzhu/copier"
	"github.com/kataras/iris/v12/websocket"
)

type ExecInterfaceService struct {
	InterfaceService *InterfaceService `inject:""`
	ExecLogService   *ExecLogService   `inject:""`

	ExtractorService  *ExtractorService     `inject:""`
	CheckpointService *CheckpointService    `inject:""`
	ExecContext       *business.ExecContext `inject:""`
	VariableService   *VariableService      `inject:""`

	ScenarioRepo    *repo2.ScenarioRepo    `inject:""`
	InterfaceRepo   *repo2.InterfaceRepo   `inject:""`
	EnvironmentRepo *repo2.EnvironmentRepo `inject:""`
	ExtractorRepo   *repo2.ExtractorRepo   `inject:""`
}

func (s *ExecInterfaceService) ExecInterfaceProcessor(interfaceProcessor *model.Processor, parentLog *domain.ExecLog, wsMsg *websocket.Message) (err error) {
	interf, err := s.InterfaceRepo.Get(interfaceProcessor.InterfaceId)
	if err != nil {
		return
	}

	req := v1.InvocationRequest{}
	copier.CopyWithOption(&req, interf, copier.Option{DeepCopy: true})

	scenario, _ := s.ScenarioRepo.Get(interfaceProcessor.ScenarioId)

	variableMap, _ := s.VariableService.GetVariablesByInterfaceAndProcessor(req.Id, interfaceProcessor.ID, scenario.ProjectId)
	requestHelper.ReplaceAll(&req, variableMap)

	resp, err := s.InterfaceService.Test(req)
	if err != nil {
		return
	}

	logPo, err := s.ExecLogService.CreateInterfaceLog(req, resp, parentLog)
	if err != nil {
		return
	}

	logExtractors, err := s.ExtractorService.ExtractInterface(interf, resp, &logPo)
	logCheckpoints, status, err := s.CheckpointService.CheckInterface(interf, resp, &logPo)

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
		ResultStatus: status,

		InterfaceExtractorsResult:  logExtractors,
		InterfaceCheckpointsResult: logCheckpoints,
	}

	*parentLog.Logs = append(*parentLog.Logs, interfaceLog)
	execHelper.SendExecMsg(*interfaceLog, wsMsg)

	return
}
