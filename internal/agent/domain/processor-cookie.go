package runDomain

import (
	"github.com/aaronchen2k/deeptest/internal/agent/run"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"time"
)

type ProcessorCookie struct {
	Id uint
	model.ProcessorEntity

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
