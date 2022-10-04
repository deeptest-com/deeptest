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

	Ordr     int          `json:"ordr"`
	Children []*Processor `json:"children"`
	Slots    iris.Map     `json:"slots"`
	Entity   interface{}  `json:"entity"`

	Log *Log `json:"log"`
}

func (p *Processor) Run() {
	logUtils.Infof("run processor %s ", p.Name)

	for _, child := range p.Children {
		child.Run()
	}
}
