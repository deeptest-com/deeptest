package jmeterProcessor

import "github.com/beevik/etree"

func GetAttrContent(elem *etree.Element) (ret string) {
	ret = elem.Attr[0].Element().Text()

	return
}
