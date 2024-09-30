package jmeterProcessor

import (
	"github.com/beevik/etree"
	agentExec "github.com/deeptest-com/deeptest/internal/agent/exec"
)

func Group(elem *etree.Element) (processor *agentExec.Processor) {
	name := elem.SelectAttrValue("testname", "")

	processor = &agentExec.Processor{
		ProcessorBase: agentExec.ProcessorBase{
			Name: name,
		},
		Entity: agentExec.ProcessorGroup{
			ProcessorEntityBase: agentExec.ProcessorEntityBase{
				Name: name,
			},
		},
	}

	return
}
