package agentExec

import (
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
)

type ProcessorRoot struct {
	ID uint `json:"id" yaml:"id"`
	ProcessorEntity
}

func (p *ProcessorRoot) Run(s *Session) (log Log, err error) {
	logUtils.Infof("root")

	log = Log{
		Name: p.Name,
	}

	return
}
