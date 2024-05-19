package agentExec

import (
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
}

type IProcessorEntity interface {
	Run(*Processor, *ExecSession) error
}

func getPreviousBrother(processor Processor) (brother Processor, ok bool) {
	for index, item := range processor.Parent.Children {
		if item.ID == processor.ID && index > 0 {
			brother = *processor.Parent.Children[index-1]

			ok = true
		}
	}

	return
}

func getResultStatus(pass bool) (resultStatus consts.ResultStatus, desc string) {
	if pass {
		resultStatus = consts.Pass
		desc = "通过"
	} else {
		resultStatus = consts.Fail
		desc = "失败"
	}

	return
}
