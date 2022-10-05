package agentExec

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12"
	"sync"
)

var (
	breakMap sync.Map
)

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
	Entity   IProcessorEntity `json:"entity"`

	Log Log `json:"log"`

	Session Session `json:"session"`
}

func (p *Processor) Run(s *Session) (log Log, err error) {
	logUtils.Infof("run processor %s - %s, %s", p.Name, p.EntityCategory, p.EntityType)

	log, _ = p.runEntity(s)

	if log.ProcessorCategory == consts.ProcessorLoop { // loop
		if p.EntityType == consts.ProcessorLoopUntil {
			p.runLoopUntil(s, log.Iterator)
		} else {
			p.runLoopItems(s, log.Iterator)
		}
	}

	p.Log = log

	return
}

func (p *Processor) runEntity(s *Session) (log Log, err error) {
	if p.Entity == nil {
		return
	}

	log, err = p.Entity.Run(s)

	return
}

func (p *Processor) runLoopUntil(s *Session, iterator domain.ExecIterator) (err error) {
	expression := iterator.UntilExpression

	for {
		result, err := EvaluateGovaluateExpression(expression, p.ID)
		pass, ok := result.(bool)
		if err != nil || !ok || pass {
			break
		}

		for _, child := range p.Children {
			childLog, _ := child.Run(s)
			if childLog.WillBreak {
				break
			}
		}
	}

	return
}

func (p *Processor) runLoopItems(s *Session, iterator domain.ExecIterator) (err error) {
	for _, item := range iterator.Items {
		SetVariable(p.ID, iterator.VariableName, item, false)

		for _, child := range p.Children {
			childLog, _ := child.Run(s)
			if childLog.WillBreak {
				break
			}
		}
	}

	return
}
