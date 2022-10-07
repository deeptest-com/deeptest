package agentExec

import (
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	execHelper "github.com/aaronchen2k/deeptest/internal/server/modules/helper/exec"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
)

type ProcessorInterface struct {
	ID uint `json:"id" yaml:"id"`
	ProcessorEntity

	domain.Request
	domain.Response
}

func (p ProcessorInterface) Run(s *Session) (log Result, err error) {
	logUtils.Infof("interface entity")

	variableMap := GetVariableMap(p.ProcessorID)
	ReplaceAll(&p.Request, variableMap)

	p.Response, err = Invoke(p.Request)
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

	log = Result{
		Name:        p.Name,
		InterfaceId: p.InterfaceID,
		//ReqContent:   string(reqContent),
		//RespContent:  string(respContent),
		//ResultStatus: status,
		//
		//InterfaceExtractorsResult:  logExtractors,
		//InterfaceCheckpointsResult: logCheckpoints,
	}

	return
}
