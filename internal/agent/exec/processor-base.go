package agentExec

import "github.com/aaronchen2k/deeptest/internal/pkg/consts"

type IProcessorEntity interface {
	Run(*Session) (string, []interface{}, error)
}

type ProcessorEntity struct {
	Name     string `gorm:"-" json:"name" yaml:"name"`
	Comments string `json:"comments" yaml:"comments"`
	Default  string `json:"default" yaml:"default"`

	ProcessorID       uint                     `json:"processorID" yaml:"processorID"`
	ProcessorCategory consts.ProcessorCategory `json:"processorCategory" yaml:"processorCategory"`
	ProcessorType     consts.ProcessorType     `json:"processorType" yaml:"processorType"`
	ParentID          uint                     `json:"parentID" yaml:"parentID"`
}
