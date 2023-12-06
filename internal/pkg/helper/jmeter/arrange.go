package jmeterHelper

import (
	"github.com/beevik/etree"
	"github.com/jinzhu/copier"
)

var (
	ElementNames = getElemWithHashTree()
)

func Arrange(children []*etree.Element, newParent *etree.Element) {
	//log.Println(elem.Tag, "-->", elem.Parent().Tag)

	for index, child := range children {
		if isHashTree(child) && !isRoot(child.Parent()) {
			continue
		}

		if isHashTree(child) && isRoot(child.Parent()) {
			Arrange(child.ChildElements(), newParent)
			return
		}

		newChild := &etree.Element{}

		copier.CopyWithOption(&newChild, child, copier.Option{DeepCopy: true})
		newChild.Child = nil

		newParent.AddChild(newChild)

		Arrange(child.ChildElements(), newChild)

		if index < len(children)-1 && children[index+1].Tag == "hashTree" {
			nextHashTree := children[index+1]

			Arrange(nextHashTree.ChildElements(), newChild)
		}
	}
}
