package runDomain

import (
	"github.com/aaronchen2k/deeptest/internal/agent/run"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
)

type IProcessorRun interface {
	Run(*run.SessionRunner) (*run.StageResult, error)
	AddChild(BaseProcessorRun) error
}

type BaseProcessorRun struct {
	Id uint
	model.ProcessorEntity
	Children []*BaseProcessorRun `json:"children" yaml:"children"`

	Result run.StageResult
}

func (b *BaseProcessorRun) AddChild(child *BaseProcessorRun) {
	b.Children = append(b.Children, child)
}

func (b *BaseProcessorRun) Run(run *run.SessionRunner) {

}
