package runDomain

import (
	"github.com/aaronchen2k/deeptest/internal/agent/run"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"log"
	"time"
)

type ProcessorCookieStage struct {
	Stage *run.TStage
}

func (s *ProcessorCookieStage) Name() string {
	return s.Stage.Name
}

func (s *ProcessorCookieStage) Category() consts.ProcessorCategory {
	return consts.ProcessorCookie
}

func (s *ProcessorCookieStage) Struct() *run.TStage {
	return s.Stage
}

func (s *ProcessorCookieStage) Run(r *run.SessionRunner) (ret *run.StageResult, err error) {
	processor, ok := s.Stage.Processor.(ProcessorCookie)
	if ok {
		log.Println(processor)
	}

	for _, child := range s.Stage.Children {
		log.Println(child)

		child.Run(r)
	}

	return
}

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
