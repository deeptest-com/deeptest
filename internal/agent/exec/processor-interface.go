package agentExec

import logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"

type ProcessorInterface struct {
	ID uint `json:"id" yaml:"id"`
	ProcessorEntity
}

func (p ProcessorInterface) Run(s *Session) (ret *Log, err error) {
	logUtils.Infof("interface")
	return
}
