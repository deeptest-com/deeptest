package agentExec

import (
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12"
	"sync"
)

var (
	breakMap sync.Map
)

type ExecReq struct {
	RootProcessor Processor         `json:"rootProcessor"`
	Variables     []domain.Variable `json:"variables"`
}

type Processor struct {
	ID uint `json:"id"`

	Name     string `json:"name"`
	Comments string `json:"comments"`

	ParentId   uint `json:"parentId"`
	ScenarioId uint `json:"scenarioId"`
	ProjectId  uint `json:"projectId"`
	UseID      uint `json:"useId"`

	EntityCategory consts.ProcessorCategory `json:"entityCategory"`
	EntityType     consts.ProcessorType     `json:"entityType"`
	EntityId       uint                     `json:"entityId"`
	InterfaceId    uint                     `json:"interfaceId"`

	Ordr     int              `json:"ordr"`
	Children []*Processor     `json:"children"`
	Slots    iris.Map         `json:"slots"`
	IsDir    bool             `json:"isDir"`
	Entity   IProcessorEntity `json:"-"`

	Parent *Processor    `json:"-"`
	Result domain.Result `json:"log"`

	Session Session `json:"-"`
}

func (p *Processor) Run(s *Session) (log domain.Result, err error) {
	logUtils.Infof("%s - %s", p.Name, p.EntityType)

	if p.Entity != nil {
		log, _ = p.Entity.Run(p, s)
	}

	return
}
