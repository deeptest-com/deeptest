package service

import (
	"fmt"
	"github.com/antchfx/jsonquery"
	"github.com/antchfx/xmlquery"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"golang.org/x/net/html"
)

type XPathService struct {
}

func (s *XPathService) GetHtmlXPath(node *html.Node,
	selectContent string, selectionType consts.NodeType, optimized bool) (ret string, err error) {
	if node == nil {
		return
	}

	if node.Type == html.DocumentNode {
		ret = "/"
		return
	}

	steps := make([]*Step, 0)
	contextNode := node

	for contextNode != nil {
		step := s.htmlXpathValue(contextNode, optimized)
		if step == nil {
			break
		} // Error - bail out early.

		steps = append(steps, step)
		if step.Optimized {
			break
		}

		contextNode = contextNode.Parent
	}

	steps = s.reverseArray(steps)

	if len(steps) > 0 && steps[0].Optimized {
		ret = ""
	} else {
		ret = "/"
	}

	for index, step := range steps {
		ret += step.toString()
		if index < len(steps)-1 {
			ret += "/"
		}
	}

	if selectionType == consts.NodeProp {
		ret += "/@" + selectContent
	}

	return
}

func (s *XPathService) GetXmlXPath(node *xmlquery.Node,
	selectContent string, selectionType consts.NodeType, optimized bool) (ret string, err error) {
	if node == nil {
		return
	}

	if node.Type == xmlquery.DocumentNode {
		ret = "/"
		return
	}

	steps := make([]*Step, 0)
	contextNode := node

	for contextNode != nil {
		step := s.xmlXpathValue(contextNode, optimized)
		if step == nil {
			break
		} // Error - bail out early.

		steps = append(steps, step)
		if step.Optimized {
			break
		}

		contextNode = contextNode.Parent
	}

	steps = s.reverseArray(steps)

	if len(steps) > 0 && steps[0].Optimized {
		ret = ""
	} else {
		ret = "/"
	}

	for index, step := range steps {
		ret += step.toString()
		if index < len(steps)-1 {
			ret += "/"
		}
	}

	if selectionType == consts.NodeProp {
		ret += "/@" + selectContent
	}

	return
}

func (s *XPathService) GetJsonXPath(node *jsonquery.Node, selectContent string, optimized bool) (ret string, err error) {
	if node == nil {
		return
	}

	if node.Type == jsonquery.DocumentNode {
		ret = "/"
		return
	}

	steps := make([]*Step, 0)
	contextNode := node

	for contextNode != nil {
		step := s.jsonXpathValue(contextNode, optimized)
		if step == nil {
			break
		} // Error - bail out early.

		steps = append(steps, step)
		if step.Optimized {
			break
		}

		contextNode = contextNode.Parent
	}

	steps = s.reverseArray(steps)

	if len(steps) > 0 && steps[0].Optimized {
		ret = ""
	} else {
		ret = "/"
	}

	for index, step := range steps {
		ret += step.toString()
		if index < len(steps)-1 {
			ret += "/"
		}
	}

	return
}

func (s *XPathService) htmlXpathValue(node *html.Node, optimized bool) (ret *Step) {
	var ownValue interface{}

	ownIndex := s.htmlXpathIndex(node)
	if ownIndex == -1 {
		return nil
	} // Error.

	switch node.Type {
	case html.ElementNode:
		id := s.getHtmlAttr(node, "id")
		if optimized && id != "" {
			return NewStep("//*[@id=\""+id+"\"]", true)
		}

		ownValue = node.Data
		break

	case html.TextNode:
		ownValue = "text()"
		break

	case html.CommentNode:
		ownValue = "comment()"
		break

	case html.DocumentNode:
		ownValue = ""
		break

	//case html.ATTRIBUTE_NODE:
	//	ownValue = '@' + node.nodeName;
	//	break
	//
	//case html.PROCESSING_INSTRUCTION_NODE:
	//	ownValue = "processing-instruction()"
	//	break

	default:
		ownValue = ""
		break
	}

	if ownIndex > 0 {
		ownValue = ownValue.(string) + fmt.Sprintf("[%d]", ownIndex)
	}

	ret = NewStep(ownValue, node.Type == html.DocumentNode)

	return
}
func (s *XPathService) xmlXpathValue(node *xmlquery.Node, optimized bool) (ret *Step) {
	var ownValue interface{}

	ownIndex := s.xmlXpathIndex(node)
	if ownIndex == -1 {
		return nil
	} // Error.

	switch node.Type {
	case xmlquery.ElementNode:
		id := s.getXmlAttr(node, "id")
		if optimized && id != "" {
			return NewStep("//*[@id=\""+id+"\"]", true)
		}

		ownValue = node.Data
		break

	case xmlquery.TextNode:
		ownValue = "text()"
		break

	case xmlquery.CommentNode:
		ownValue = "comment()"
		break

	case xmlquery.DocumentNode:
		ownValue = ""
		break

	//case html.ATTRIBUTE_NODE:
	//	ownValue = '@' + node.nodeName;
	//	break
	//
	//case html.PROCESSING_INSTRUCTION_NODE:
	//	ownValue = "processing-instruction()"
	//	break

	default:
		ownValue = ""
		break
	}

	if ownIndex > 0 {
		ownValue = ownValue.(string) + fmt.Sprintf("[%d]", ownIndex)
	}

	ret = NewStep(ownValue, node.Type == xmlquery.DocumentNode)

	return
}
func (s *XPathService) jsonXpathValue(node *jsonquery.Node, optimized bool) (ret *Step) {
	var ownValue interface{}

	ownIndex := s.jsonXpathIndex(node)
	if ownIndex == -1 {
		return nil
	} // Error.

	switch node.Type {
	case jsonquery.ElementNode:
		id := s.getJsonAttr(node, "id")
		if optimized && id != "" {
			return NewStep("//*[@id=\""+id+"\"]", true)
		}

		ownValue = node.Data
		break

	case jsonquery.TextNode:
		ownValue = "text()"
		break

	case jsonquery.DocumentNode:
		ownValue = ""
		break

	//case html.ATTRIBUTE_NODE:
	//	ownValue = '@' + node.nodeName;
	//	break
	//
	//case html.PROCESSING_INSTRUCTION_NODE:
	//	ownValue = "processing-instruction()"
	//	break

	default:
		ownValue = ""
		break
	}

	if ownIndex > 0 {
		ownValue = ownValue.(string) + fmt.Sprintf("*[%d]", ownIndex)
	}

	ret = NewStep(ownValue, node.Type == jsonquery.DocumentNode)

	return
}

func (s *XPathService) htmlXpathIndex(node *html.Node) (ret int) {
	var siblings []*html.Node

	if node.Parent != nil {
		siblings = s.getHtmlChildren(node.Parent)
	}

	if len(siblings) == 0 {
		return 0
	} // Root node - no siblings.

	hasSameNamedElements := false

	for i := 0; i < len(siblings); i++ {
		areNodesSimilar := s.areHtmlNodesSimilar(node, siblings[i])

		if areNodesSimilar && siblings[i] != node {
			hasSameNamedElements = true
			break
		}
	}

	if !hasSameNamedElements {
		return 0
	}

	ownIndex := 1 // XPath indices start with 1.
	for i := 0; i < len(siblings); i++ {
		if s.areHtmlNodesSimilar(node, siblings[i]) {
			if siblings[i] == node {
				return ownIndex
			}

			ownIndex++
		}
	}
	return -1 // An error occurred: |node| not found in parent's children.
}
func (s *XPathService) xmlXpathIndex(node *xmlquery.Node) (ret int) {
	var siblings []*xmlquery.Node

	if node.Parent != nil {
		siblings = s.getXmlChildren(node.Parent)
	}

	if len(siblings) == 0 {
		return 0
	} // Root node - no siblings.

	hasSameNamedElements := false

	for i := 0; i < len(siblings); i++ {
		areNodesSimilar := s.areXmlNodesSimilar(node, siblings[i])

		if areNodesSimilar && siblings[i] != node {
			hasSameNamedElements = true
			break
		}
	}

	if !hasSameNamedElements {
		return 0
	}

	ownIndex := 1 // XPath indices start with 1.
	for i := 0; i < len(siblings); i++ {
		if s.areXmlNodesSimilar(node, siblings[i]) {
			if siblings[i] == node {
				return ownIndex
			}

			ownIndex++
		}
	}
	return -1 // An error occurred: |node| not found in parent's children.
}
func (s *XPathService) jsonXpathIndex(node *jsonquery.Node) (ret int) {
	var siblings []*jsonquery.Node

	if node.Parent != nil {
		siblings = s.getJsonChildren(node.Parent)
	}

	if len(siblings) == 0 {
		return 0
	} // Root node - no siblings.

	hasSameNamedElements := false

	for i := 0; i < len(siblings); i++ {
		areNodesSimilar := s.areJsonNodesSimilar(node, siblings[i])

		if areNodesSimilar && siblings[i] != node {
			hasSameNamedElements = true
			break
		}
	}

	if !hasSameNamedElements {
		return 0
	}

	ownIndex := 1 // XPath indices start with 1.
	for i := 0; i < len(siblings); i++ {
		if s.areJsonNodesSimilar(node, siblings[i]) {
			if siblings[i] == node {
				return ownIndex
			}

			ownIndex++
		}
	}
	return -1 // An error occurred: |node| not found in parent's children.
}

func (s *XPathService) getHtmlChildren(node *html.Node) (ret []*html.Node) {
	child := node.FirstChild

	for child != nil {
		ret = append(ret, child)

		child = child.NextSibling
	}

	return
}
func (s *XPathService) getXmlChildren(node *xmlquery.Node) (ret []*xmlquery.Node) {
	child := node.FirstChild

	for child != nil {
		ret = append(ret, child)

		child = child.NextSibling
	}

	return
}
func (s *XPathService) getJsonChildren(node *jsonquery.Node) (ret []*jsonquery.Node) {
	child := node.FirstChild

	for child != nil {
		ret = append(ret, child)

		child = child.NextSibling
	}

	return
}

// Returns -1 in case of error, 0 if no siblings matching the same expression,
// <XPath index among the same expression-matching sibling nodes> otherwise.
func (s *XPathService) areHtmlNodesSimilar(left, right *html.Node) (ret bool) {
	//if left == right {
	//	return true
	//}

	if left.Type == html.ElementNode && right.Type == html.ElementNode {
		return left.Data == right.Data
	}

	if left.Type == right.Type {
		return true
	}

	return false
}
func (s *XPathService) areXmlNodesSimilar(left, right *xmlquery.Node) (ret bool) {
	//if left == right {
	//	return true
	//}

	if left.Type == xmlquery.ElementNode && right.Type == xmlquery.ElementNode {
		return left.Data == right.Data
	}

	if left.Type == right.Type {
		return true
	}

	return false
}
func (s *XPathService) areJsonNodesSimilar(left, right *jsonquery.Node) (ret bool) {
	//if left == right {
	//	return true
	//}

	if left.Type == jsonquery.ElementNode && right.Type == jsonquery.ElementNode {
		return left.Data == right.Data
	}

	if left.Type == right.Type {
		return true
	}

	return false
}

func (s *XPathService) getHtmlAttr(node *html.Node, name string) (ret string) {
	for _, attr := range node.Attr {
		if attr.Key == name {
			return attr.Val
		}
	}
	return
}
func (s *XPathService) getXmlAttr(node *xmlquery.Node, name string) (ret string) {
	for _, attr := range node.Attr {
		if attr.Name.Local == name {
			return attr.Value
		}
	}
	return
}
func (s *XPathService) getJsonAttr(node *jsonquery.Node, name string) (ret string) {
	//for _, attr := range node. {
	//	if attr.Name == name {
	//		return attr.Sample
	//	}
	//}
	return
}

func (s *XPathService) reverseArray(arr []*Step) (ret []*Step) {
	for i := len(arr) - 1; i >= 0; i-- {
		ret = append(ret, arr[i])
	}

	return
}

type Step struct {
	Value     interface{}
	Optimized bool
}

func NewStep(value interface{}, optimized bool) *Step {
	return &Step{
		Value:     value,
		Optimized: optimized,
	}
}

func (step *Step) toString() string {
	return step.Value.(string)
}
