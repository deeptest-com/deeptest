package agentExec

import (
	"github.com/aaronchen2k/deeptest/internal/agent/run"
	"time"
)

type ProcessorCookie struct {
	ID uint `json:"id" yaml:"id"`
	ProcessorEntity

	CookieName   string     `json:"cookieName" yaml:"cookieName"`
	VariableName string     `json:"variableName" yaml:"variableName"`
	RightValue   string     `json:"rightValue" yaml:"rightValue"`
	Domain       string     `json:"domain" yaml:"domain"`
	ExpireTime   *time.Time `json:"expireTime" yaml:"expireTime"`

	Children []interface{} `json:"children" yaml:"children" gorm:"-"`
}

func (s *ProcessorCookie) Run(r *run.SessionRunner) (ret *run.StageResult, err error) {
	return
}
