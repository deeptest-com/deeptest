package jmeterHelper

import (
	"github.com/beevik/etree"
	agentExec "github.com/deeptest-com/deeptest/internal/agent/exec"
	jmeterProcessor "github.com/deeptest-com/deeptest/internal/pkg/helper/jmeter/processor"
	"log"
)

func Parse(elem *etree.Element, parentProcessor *agentExec.Processor) {
	parentTag := ""
	if elem.Parent() != nil {
		parentTag = elem.Parent().Tag
	}

	log.Println(elem.Tag, " @ ", parentTag)

	processor := GenProcessor(elem, parentProcessor)

	for _, child := range elem.ChildElements() {
		if isProp(child) {
			continue
		}

		Parse(child, processor)
	}
}

func GenProcessor(elem *etree.Element, parentProcessor *agentExec.Processor) (processor *agentExec.Processor) {
	if elem.Tag == TestPlan.String() {
		processor = jmeterProcessor.TestPlan(elem)

	} else if elem.Tag == ThreadGroup.String() {
		processor = jmeterProcessor.ThreadGroup(elem)

	} else if elem.Tag == HTTPSamplerProxy.String() {
		processor = jmeterProcessor.HTTPSamplerProxy(elem)

	} else {
		processor = jmeterProcessor.Group(elem)

	}

	if processor.Name == "" {
		processor.Name = elem.Tag
	} else {
		processor.Name += " (" + elem.Tag + ")"
	}

	parentProcessor.Children = append(parentProcessor.Children, processor)

	return
}
