package agentExec

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
)

type IProcessorEntity interface {
	Run(*Session) (Result, error)
}

type ProcessorEntity struct {
	Name     string `gorm:"-" json:"name"`
	Comments string `json:"comments"`
	Default  string `json:"default"`

	ProcessorID       uint                     `json:"processorID" yaml:"processorID"`
	ProcessorCategory consts.ProcessorCategory `json:"processorCategory" yaml:"processorCategory"`
	ProcessorType     consts.ProcessorType     `json:"processorType" yaml:"processorType"`
	ParentID          uint                     `json:"parentID" yaml:"parentID"`

	Result Result `json:"result,omitempty" gorm:"-"`
}
