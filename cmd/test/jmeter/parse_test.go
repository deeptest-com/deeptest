package test

import (
	"github.com/beevik/etree"
	agentExec "github.com/deeptest-com/deeptest/internal/agent/exec"
	jmeterHelper "github.com/deeptest-com/deeptest/internal/pkg/helper/jmeter"
	"log"
	"testing"
)

const (
	jmx = "/Users/aaron/rd/project/gudi/deeptest-main/xdoc/jmeter/baidu.jmx"
)

func TestParse(t *testing.T) {
	doc := etree.NewDocument()
	if err := doc.ReadFromFile(jmx); err != nil {
		panic(err)
	}

	rootElement := &etree.Element{}
	jmeterHelper.Arrange(doc.Root().ChildElements(), rootElement)

	rootProcessor := &agentExec.Processor{}

	for _, child := range rootElement.ChildElements() {
		jmeterHelper.Parse(child, rootProcessor) // child is test plan
	}

	log.Println(rootProcessor)
}
