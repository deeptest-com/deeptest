package agentExec

import (
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
)

type ProcessorEntityBase struct {
	Name     string `gorm:"-" json:"name"`
	Comments string `json:"comments"`
	Default  string `json:"default"`

	ProcessorID       uint                     `json:"processorID" yaml:"processorID"`
	ProcessorCategory consts.ProcessorCategory `json:"processorCategory" yaml:"processorCategory"`
	ProcessorType     consts.ProcessorType     `json:"processorType" yaml:"processorType"`
	ParentID          uint                     `json:"parentID" yaml:"parentID"`

	//Children []*Processor `json:"children" gorm:"-"`

	Result domain.Result `json:"result,omitempty" gorm:"-"`
}

type IProcessorEntity interface {
	Run(*Processor, *Session) (domain.Result, error)
}
