package jmeterHelper

import (
	stringUtils "github.com/aaronchen2k/deeptest/pkg/lib/string"
	"github.com/beevik/etree"
	"log"
)

var (
	ElementNames = getElemWithHashTree()
)

func Parse(elem *etree.Element) {
	log.Println(elem.Tag)

	children := elem.ChildElements()

	for index, child := range children {
		if isHashTree(child) && !isRoot(elem) {
			continue
		}

		if stringUtils.StrInArr(child.Tag, ElementNames) {
			if index < len(children)-1 && children[index+1].Tag == "hashTree" {
				next := children[index+1]
				child.AddChild(next)

				log.Println(1)
			}
		}

		Parse(child)
	}
}
