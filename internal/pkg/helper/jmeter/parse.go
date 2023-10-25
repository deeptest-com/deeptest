package jmeterHelper

import (
	agentExec "github.com/aaronchen2k/deeptest/internal/agent/exec"
	"github.com/beevik/etree"
	"log"
)

func Parse(elem *etree.Element, processor *agentExec.Processor) {
	parentTag := ""
	if elem.Parent() != nil {
		parentTag = elem.Parent().Tag
	}

	log.Println(elem.Tag, " @ ", parentTag)

	for _, child := range elem.ChildElements() {
		log.Println(child.Tag, " @ ", child.Parent().Tag)

		childProcessor := &agentExec.Processor{}

		for _, son := range child.ChildElements() {
			Parse(son, childProcessor)
		}
	}
}
