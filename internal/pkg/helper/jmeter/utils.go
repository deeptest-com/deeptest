package jmeterHelper

import "github.com/beevik/etree"

func isRoot(elem *etree.Element) (ret bool) {
	ret = elem.Tag == JmeterTestPlan.String()

	return
}
func isHashTree(elem *etree.Element) (ret bool) {
	ret = elem.Tag == HashTree.String()

	return
}

func getElemWithHashTree() (ret []string) {
	for _, item := range ElemWithHashTree {
		ret = append(ret, item.String())
	}

	return
}

var (
	ElemWithHashTree = []JmeterElement{
		HTTPSamplerProxy,
		JmeterTestPlan,
		ResultCollector,
		TestPlan,
		ThreadGroup,
	}
)
