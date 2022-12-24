package service

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"golang.org/x/net/html"
)

type XPathService struct {
}

func (s *XPathService) GetXPath(node *html.Node,
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
		step := s.xpathValue(contextNode, optimized)
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

	if selectionType == consts.Prop {
		ret += "/@" + selectContent
	}

	return
}

func (s *XPathService) xpathValue(node *html.Node, optimized bool) (ret *Step) {
	var ownValue interface{}

	ownIndex := s.xpathIndex(node)
	if ownIndex == -1 {
		return nil
	} // Error.

	switch node.Type {
	case html.ElementNode:
		id := s.getAttr(node, "id")
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

func (s *XPathService) xpathIndex(node *html.Node) (ret int) {
	var siblings []*html.Node

	if node.Parent != nil {
		siblings = s.getChildren(node.Parent)
	}

	if len(siblings) == 0 {
		return 0
	} // Root node - no siblings.

	hasSameNamedElements := false

	for i := 0; i < len(siblings); i++ {
		areNodesSimilar := s.areNodesSimilar(node, siblings[i])

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
		if s.areNodesSimilar(node, siblings[i]) {
			if siblings[i] == node {
				return ownIndex
			}

			ownIndex++
		}
	}
	return -1 // An error occurred: |node| not found in parent's children.
}

func (s *XPathService) getChildren(node *html.Node) (ret []*html.Node) {
	child := node.FirstChild

	for child != nil {
		ret = append(ret, child)

		child = child.NextSibling
	}

	return
}

// Returns -1 in case of error, 0 if no siblings matching the same expression,
// <XPath index among the same expression-matching sibling nodes> otherwise.
func (s *XPathService) areNodesSimilar(left, right *html.Node) (ret bool) {
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

func (s *XPathService) reverseArray(arr []*Step) (ret []*Step) {
	for i := len(arr) - 1; i >= 0; i-- {
		ret = append(ret, arr[i])
	}

	return
}

func (s *XPathService) getAttr(node *html.Node, name string) (ret string) {
	for _, attr := range node.Attr {
		if attr.Key == name {
			return attr.Val
		}
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
