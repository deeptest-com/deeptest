package test

import (
	jmeterHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/jmeter"
	"github.com/beevik/etree"
	"log"
	"testing"
)

const (
	jmx = "/Users/aaron/rd/project/gudi/deeptest-main/xdoc/jmeter/baidu.jmx"
)

func TestParse(t *testing.T) {
	//content := fileUtils.ReadFileBuf(jmx)
	//
	//testPlan := jmeterHelper.JmeterTestPlan{}
	//xml.Unmarshal(content, &testPlan)
	//
	//log.Println(testPlan)

	doc := etree.NewDocument()
	if err := doc.ReadFromFile(jmx); err != nil {
		panic(err)
	}

	root := &etree.Element{}

	jmeterHelper.Arrange(doc.Root().ChildElements(), root)
	jmeterHelper.Parse(root)

	log.Println(1)
}
