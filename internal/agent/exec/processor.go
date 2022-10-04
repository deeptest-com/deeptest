package agentExec

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12"
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

	Log *Log `json:"log"`

	Session Session `json:"session"`
}

func (p *Processor) Run(s *Session) {
	logUtils.Infof("run processor %s - %s, %s", p.Name, p.EntityCategory, p.EntityType)

	iterateVariableName, iterateVariableValues, err := p.runEntity(s)
	if err != nil || iterateVariableName == "" {
		return
	}

	for _, variable := range iterateVariableValues {
		SetVariable(p.ID, iterateVariableName, variable, false)

		for _, child := range p.Children {
			child.Run(s)
		}
	}
}

func (p *Processor) runEntity(s *Session) (iterateVariableName string, iterateVariableValues []interface{}, err error) {
	if p.Entity == nil {
		return
	}

	iterateVariableName, iterateVariableValues, err = p.Entity.Run(s)

	return
}
