package agentExec

import (
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
)

type ProcessorGroup struct {
	ID uint `json:"id" yaml:"id"`
	ProcessorEntity
}

func (p ProcessorGroup) Run(s *Session) (log Log, err error) {
	logUtils.Infof("group")

	log = Log{
		Name: p.Name,
	}

	return
}
