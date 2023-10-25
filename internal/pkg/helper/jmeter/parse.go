package jmeterHelper

import (
	"github.com/beevik/etree"
	"log"
)

func Parse(elem *etree.Element) {
	parentTag := ""
	if elem.Parent() != nil {
		parentTag = elem.Parent().Tag
	}

	log.Println(elem.Tag, " @ ", parentTag)

	for _, child := range elem.ChildElements() {
		log.Println(child.Tag, " @ ", child.Parent().Tag)

		for _, son := range child.ChildElements() {
			Parse(son)
		}
	}
}
