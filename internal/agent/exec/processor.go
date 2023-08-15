package agentExec

import (
	"encoding/json"
	agentDomain "github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12"
)

type Processor struct {
	ProcessorBase
	Entity  IProcessorEntity `json:"entity"`
	Disable bool             `json:"disable"`
}

type ProcessorMsg struct {
	ProcessorBase
}

type ProcessorBase struct {
	ID uint `json:"id"`

	Name     string            `json:"name"`
	Comments string            `json:"comments"`
	Method   consts.HttpMethod `json:"method" yaml:"method"`

	ParentId   uint `json:"parentId"`
	ScenarioId uint `json:"scenarioId"`
	ProjectId  uint `json:"projectId"`
	UseID      uint `json:"useId"`

	EntityCategory      consts.ProcessorCategory `json:"entityCategory"`
	EntityType          consts.ProcessorType     `json:"entityType"`
	EntityId            uint                     `json:"entityId"`
	EndpointInterfaceId uint                     `json:"endpointInterfaceId"`

	Ordr      int             `json:"ordr"`
	Children  []*Processor    `json:"children"`
	Slots     iris.Map        `json:"slots"`
	IsDir     bool            `json:"isDir"`
	EntityRaw json.RawMessage `json:"entityRaw"`

	Parent                *Processor                      `json:"-"`
	Result                *agentDomain.ScenarioExecResult `json:"result"`
	ProcessorInterfaceSrc consts.ProcessorInterfaceSrc    `json:"processorInterfaceSrc"`

	Session Session `json:"-"`
}

func (p *Processor) Run(s *Session) (err error) {
	_logUtils.Infof("%d - %s %s", p.ID, p.Name, p.EntityType)
	CurrScenarioProcessorId = p.ID

	if !p.Disable && p.Entity != nil {
		p.Entity.Run(p, s)
	}

	return
}

func (p *Processor) AddResultToParent() (err error) {
	p.Parent.Result.Children = append(p.Parent.Result.Children, p.Result)
	return
}

func (p *Processor) RestoreEntity() (err error) {
	bytes, err := p.EntityRaw.MarshalJSON()

	switch p.EntityCategory {
	case consts.ProcessorInterface:
		ret := ProcessorInterface{}
		json.Unmarshal(bytes, &ret)
		p.Entity = ret

	case consts.ProcessorRoot:
		ret := ProcessorRoot{}
		json.Unmarshal(bytes, &ret)
		p.Entity = ret

	case consts.ProcessorGroup:
		ret := ProcessorGroup{}
		json.Unmarshal(bytes, &ret)
		p.Entity = ret

	case consts.ProcessorLogic:
		ret := ProcessorLogic{}
		json.Unmarshal(bytes, &ret)
		p.Entity = ret

	case consts.ProcessorLoop:
		ret := ProcessorLoop{}
		json.Unmarshal(bytes, &ret)
		p.Entity = ret

	case consts.ProcessorVariable:
		ret := ProcessorVariable{}
		json.Unmarshal(bytes, &ret)
		p.Entity = ret

	case consts.ProcessorTimer:
		ret := ProcessorTimer{}
		json.Unmarshal(bytes, &ret)
		p.Entity = ret

	case consts.ProcessorPrint:
		ret := ProcessorPrint{}
		json.Unmarshal(bytes, &ret)
		p.Entity = ret

	case consts.ProcessorCookie:
		ret := ProcessorCookie{}
		json.Unmarshal(bytes, &ret)
		p.Entity = ret

	case consts.ProcessorAssertion:
		ret := ProcessorAssertion{}
		json.Unmarshal(bytes, &ret)
		p.Entity = ret

	case consts.ProcessorData:
		ret := ProcessorData{}
		json.Unmarshal(bytes, &ret)
		p.Entity = ret

	case consts.ProcessorCustomCode:
		ret := ProcessorCustomCode{}
		json.Unmarshal(bytes, &ret)
		p.Entity = ret

	default:
	}

	return
}

func (p *Processor) AppendNewChildProcessor(category consts.ProcessorCategory, typ consts.ProcessorType) (child Processor) {
	child = Processor{
		ProcessorBase: ProcessorBase{
			EntityCategory: category,
			EntityType:     typ,
			Parent:         p,
			ParentId:       p.ID,
		},
	}

	child.Result = &agentDomain.ScenarioExecResult{
		ProcessorCategory: child.EntityCategory,
		ProcessorType:     child.EntityType,
		ParentId:          int(p.ID),
	}

	return
}
