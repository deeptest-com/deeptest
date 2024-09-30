package service

import (
	"fmt"
	"github.com/antchfx/xmlquery"
	v1 "github.com/deeptest-com/deeptest/cmd/server/v1/domain"
	queryUtils "github.com/deeptest-com/deeptest/internal/agent/exec/utils/query"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"strings"
)

type ParserXmlService struct {
	ParserService     *ParserService     `inject:""`
	XPathService      *XPathService      `inject:""`
	ParserRegxService *ParserRegxService `inject:""`
}

func (s *ParserXmlService) ParseXml(req *v1.ParserRequest) (ret v1.ParserResponse, err error) {
	docXml1, docXml2, selectionType := s.updateXmlElem(req.DocContent, req.SelectContent, req.StartLine, req.EndLine,
		req.StartColumn, req.EndColumn)

	var elem1 *xmlquery.Node
	var elem2 *xmlquery.Node

	elem1 = s.getXmlSelectedElem(docXml1, selectionType)
	if selectionType == consts.NodeContent {
		elem2 = s.getXmlSelectedElem(docXml2, selectionType)

		if elem1 == nil || elem2 == nil { // must both pass
			ret = s.ParserService.GetRegxResponse(req)
			return
		}
	}

	exprType := "xpath"

	var expr1 string
	var expr2 string

	expr1, _ = s.XPathService.GetXmlXPath(elem1, req.SelectContent, selectionType, true)
	if selectionType == consts.NodeContent {
		expr2, _ = s.XPathService.GetXmlXPath(elem2, req.SelectContent, selectionType, true)
	}

	if expr1 == "" || (selectionType == consts.NodeContent && (expr2 == "" || expr1 != expr2)) {
		ret = s.ParserService.GetRegxResponse(req)
		return
	}

	result := queryUtils.XmlQuery(req.DocContent, expr1)
	fmt.Printf("%s - %s: %v", selectionType, expr1, result)

	ret = v1.ParserResponse{
		SelectionType: selectionType,
		Expr:          expr1,
		ExprType:      exprType,
	}

	return
}

func (s *ParserXmlService) updateXmlElem(docXml, selectContent string,
	startLineNum, endLineNum, startColumn, endColumn int) (ret1, ret2 string, selectionType consts.NodeType) {
	lines := strings.Split(docXml, "\n")

	selectionType = s.getXmlSelectionType(lines, startLineNum, endLineNum, startColumn, endColumn)

	line := []rune(lines[startLineNum])
	newStr := fmt.Sprintf(" %s=\"true\" ", consts.DeepestKey)

	if selectionType == consts.NodeElem {
		newLine := string(line[:startColumn]) + selectContent + newStr + string(line[endColumn:])

		lines[startLineNum] = newLine
		ret1 = strings.Join(lines, "\n")

	} else if selectionType == consts.NodeProp {
		newLine := string(line[:startColumn]) + newStr + string(line[startColumn:])

		lines[startLineNum] = newLine
		ret1 = strings.Join(lines, "\n")

	} else if selectionType == consts.NodeContent {
		newStr := fmt.Sprintf("[[%s]]", consts.DeepestKey)

		// 1
		lines1 := strings.Split(docXml, "\n")
		startLine := []rune(lines1[startLineNum])

		newStartLine := string(startLine[:startColumn]) + newStr + string(startLine[startColumn:])
		lines1[startLineNum] = newStartLine
		ret1 = strings.Join(lines1, "\n")

		//2
		lines2 := strings.Split(docXml, "\n")
		endLine := []rune(lines2[endLineNum])

		newEndLine := string(endLine[:endColumn]) + newStr + string(endLine[endColumn:])
		lines2[endLineNum] = newEndLine
		ret2 = strings.Join(lines2, "\n")
	}

	return
}

func (s *ParserXmlService) getXmlSelectedElem(docXml string, selectionType consts.NodeType) (ret *xmlquery.Node) {
	doc, err := xmlquery.Parse(strings.NewReader(docXml))
	if err != nil {
		return
	}

	expr := ""
	if selectionType == consts.NodeElem || selectionType == consts.NodeProp {
		expr = fmt.Sprintf("//*[@%s]", consts.DeepestKey)
	} else if selectionType == consts.NodeContent {
		expr = fmt.Sprintf("//text()[contains(.,\"%s\")]", consts.DeepestKey)
	}

	ret, err = xmlquery.Query(doc, expr)

	return
}

func (s *ParserXmlService) queryElem(docXml, xpath string) (ret *xmlquery.Node) {
	doc, err := xmlquery.Parse(strings.NewReader(docXml))
	if err != nil {
		return
	}

	expr := fmt.Sprintf(xpath)
	ret, err = xmlquery.Query(doc, expr)

	return
}

func (s *ParserXmlService) getXmlSelectionType(lines []string, startLine, endLine, startColumn, endColumn int) (
	ret consts.NodeType) {

	leftNoSpaceChar := s.ParserService.getLeftNoSpaceChar(lines, startLine, startColumn)
	rightChar := s.ParserService.getRightChar(lines, endLine, endColumn)

	if leftNoSpaceChar == "<" && (rightChar == " " || rightChar == ">") {
		ret = consts.NodeElem
		return
	}

	leftChar := s.ParserService.getLeftChar(lines, startLine, startColumn)
	rightNoSpaceChar := s.ParserService.getRightNoSpaceChar(lines, endLine, endColumn)

	if leftChar == " " && rightNoSpaceChar == "=" {
		ret = consts.NodeProp
		return
	}

	ret = consts.NodeContent
	return
}
