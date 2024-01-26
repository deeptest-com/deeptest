package jmeterProcessor

import (
	agentExec "github.com/aaronchen2k/deeptest/internal/agent/exec"
	"github.com/beevik/etree"
)

func ThreadGroup(elem *etree.Element) (processor *agentExec.Processor) {
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
