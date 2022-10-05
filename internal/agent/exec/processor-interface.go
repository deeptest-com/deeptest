package agentExec

import (
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
)

type ProcessorInterface struct {
	ID uint `json:"id" yaml:"id"`
	ProcessorEntity

	InterfaceID uint `json:"interfaceID"`
}

func (p ProcessorInterface) Run(s *Session) (log Result, err error) {
	logUtils.Infof("interface entity")

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
