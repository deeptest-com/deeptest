package agentExec

import (
	"encoding/json"
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
	RootProcessor *Processor        `json:"rootProcessor"`
	Variables     []domain.Variable `json:"variables"`
	ServerUrl     string            `json:"serverUrl"`
	Token         string            `json:"token"`
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

	Ordr      int              `json:"ordr"`
	Children  []*Processor     `json:"children"`
	Slots     iris.Map         `json:"slots"`
	IsDir     bool             `json:"isDir"`
	Entity    IProcessorEntity `json:"entity"`
	EntityRaw json.RawMessage  `json:"entityRaw"`

	Parent *Processor     `json:"-"`
	Result *domain.Result `json:"result"`

	Session Session `json:"-"`
}

func (p *Processor) Run(s *Session) (err error) {
	logUtils.Infof("%s - %s", p.Name, p.EntityType)

	if p.Entity != nil {
		p.Entity.Run(p, s)
	}

	return
}

func (p *Processor) AddResultToParent() (err error) {
	p.Parent.Result.Children = append(p.Parent.Result.Children, p.Result)
	return
}

func (p *Processor) UnmarshalEntity() (err error) {
	bytes, err := p.EntityRaw.MarshalJSON()

	switch p.EntityCategory {
	case consts.ProcessorRoot:
		ret := ProcessorRoot{}
		json.Unmarshal(bytes, &ret)
		p.Entity = &ret

	case consts.ProcessorInterface:
		ret := ProcessorInterface{}
		json.Unmarshal(bytes, &ret)
		p.Entity = &ret

	case consts.ProcessorGroup:
		ret := ProcessorGroup{}
		json.Unmarshal(bytes, &ret)
		p.Entity = &ret

	case consts.ProcessorLogic:
		ret := ProcessorLogic{}
		json.Unmarshal(bytes, &ret)
		p.Entity = &ret

	case consts.ProcessorLoop:
		ret := ProcessorLoop{}
		json.Unmarshal(bytes, &ret)
		p.Entity = &ret

	case consts.ProcessorVariable:
		ret := ProcessorVariable{}
		json.Unmarshal(bytes, &ret)
		p.Entity = &ret

	case consts.ProcessorTimer:
		ret := ProcessorTimer{}
		json.Unmarshal(bytes, &ret)
		p.Entity = &ret

	case consts.ProcessorPrint:
		ret := ProcessorPrint{}
		json.Unmarshal(bytes, &ret)
		p.Entity = &ret

	case consts.ProcessorCookie:
		ret := ProcessorCookie{}
		json.Unmarshal(bytes, &ret)
		p.Entity = &ret

	case consts.ProcessorAssertion:
		ret := ProcessorAssertion{}
		json.Unmarshal(bytes, &ret)
		p.Entity = &ret

	case consts.ProcessorExtractor:
		ret := ProcessorExtractor{}
		json.Unmarshal(bytes, &ret)
		p.Entity = &ret

	case consts.ProcessorData:
		ret := ProcessorData{}
		json.Unmarshal(bytes, &ret)
		p.Entity = &ret

	default:
	}

	return
}

func (p *Processor) AppendNewChildProcessor(category consts.ProcessorCategory, typ consts.ProcessorType) (child Processor) {
	child = Processor{
		EntityCategory: category,
		EntityType:     typ,
		Parent:         p,
		ParentId:       p.ID,
	}

	child.Result = &domain.Result{
		ProcessorCategory: child.EntityCategory,
		ProcessorType:     child.EntityType,
		ParentId:          int(p.ID),
	}

	return
}
