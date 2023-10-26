package jmeterHelper

import (
	stringUtils "github.com/aaronchen2k/deeptest/pkg/lib/string"
	"github.com/beevik/etree"
	"strings"
)

func isRoot(elem *etree.Element) (ret bool) {
	ret = elem.Tag == JmeterTestPlan.String()

	return
}

func isCaredElement(elem *etree.Element) (ret bool) {
	ret = stringUtils.StrInArr(elem.Tag, ElementNames)

	return
}

func isHashTree(elem *etree.Element) (ret bool) {
	ret = elem.Tag == HashTree.String()

	return
}

func isProp(elem *etree.Element) (ret bool) {
	index := strings.Index(elem.Tag, "Prop")

	if index == len(elem.Tag)-4 {
		ret = true
	}

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
