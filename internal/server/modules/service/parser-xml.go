package service

import (
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	queryHelper "github.com/aaronchen2k/deeptest/internal/agent/exec/utils/query"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/antchfx/xmlquery"
	"strings"
)

type ParserXmlService struct {
	ParserService     *ParserService     `inject:""`
	XPathService      *XPathService      `inject:""`
	ParserRegxService *ParserRegxService `inject:""`
}

func (s *ParserXmlService) ParseXml(req *v1.ParserRequest) (ret v1.ParserResponse, err error) {
	docXml, selectionType := s.updateXmlElem(req.DocContent, req.SelectContent, req.StartLine, req.EndLine,
		req.StartColumn, req.EndColumn)

	req.DocContent = docXml

	elem := s.getXmlSelectedElem(req.DocContent, selectionType)

	exprType := "xpath"
	expr, _ := s.XPathService.GetXmlXPath(elem, req.SelectContent, selectionType, true)

	if expr != "" {
		result := queryHelper.XmlQuery(req.DocContent, expr)
		fmt.Printf("%s - %s: %v", selectionType, expr, result)
	} else {
		expr, _ = s.ParserRegxService.getRegxExpr(req.DocContent, req.SelectContent,
			req.StartLine, req.StartColumn,
			req.EndLine, req.EndColumn)
		exprType = "regx"
	}

	ret = v1.ParserResponse{
		SelectionType: selectionType,
		Expr:          expr,
		ExprType:      exprType,
	}

	return
}

func (s *ParserXmlService) updateXmlElem(docXml, selectContent string,
	startLine, endLine, startColumn, endColumn int) (ret string, selectionType consts.NodeType) {
	lines := strings.Split(docXml, "\n")

	selectionType = s.getXmlSelectionType(lines, startLine, endLine, startColumn, endColumn)

	line := []rune(lines[startLine])
	newStr := fmt.Sprintf(" %s=\"true\" ", consts.DeepestKey)

	if selectionType == consts.NodeElem {
		newLine := string(line[:startColumn]) + selectContent + newStr + string(line[endColumn:])

		lines[startLine] = newLine

	} else if selectionType == consts.NodeProp {
		newLine := string(line[:startColumn]) + newStr + string(line[startColumn:])

		lines[startLine] = newLine

	} else if selectionType == consts.NodeContent {
		newStr = fmt.Sprintf("[[%s]]", consts.DeepestKey)
		newLine := string(line[:endColumn]) + newStr + string(line[endColumn:])

		lines[startLine] = newLine
	}

	ret = strings.Join(lines, "\n")
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
