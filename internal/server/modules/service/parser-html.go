package service

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	v1 "github.com/deeptest-com/deeptest/cmd/server/v1/domain"
	queryUtils "github.com/deeptest-com/deeptest/internal/agent/exec/utils/query"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"golang.org/x/net/html"
	"strings"
)

type ParserHtmlService struct {
	ParserService     *ParserService     `inject:""`
	XPathService      *XPathService      `inject:""`
	ParserRegxService *ParserRegxService `inject:""`
}

func (s *ParserHtmlService) ParseHtml(req *v1.ParserRequest) (ret v1.ParserResponse, err error) {
	docHtml1, docHtml2, selectionType := s.updateHtmlElem(req.DocContent, req.SelectContent, req.StartLine, req.EndLine,
		req.StartColumn, req.EndColumn)

	var elem1 *html.Node
	var elem2 *html.Node

	elem1 = s.getHtmlSelectedElem(docHtml1, selectionType)
	if selectionType == consts.NodeContent {
		elem2 = s.getHtmlSelectedElem(docHtml2, selectionType)
		if elem1 == nil || elem2 == nil || elem1.Type != elem2.Type { // must both pass
			ret = s.ParserService.GetRegxResponse(req)
			return
		}
	}

	exprType := "xpath"

	var expr1 string
	var expr2 string
	expr1, _ = s.XPathService.GetHtmlXPath(elem1, req.SelectContent, selectionType, true)

	if selectionType == consts.NodeContent {
		expr2, _ = s.XPathService.GetHtmlXPath(elem2, req.SelectContent, selectionType, true)
	}

	if expr1 == "" || (selectionType == consts.NodeContent && (expr2 == "" || expr1 != expr2)) {
		ret = s.ParserService.GetRegxResponse(req)
		return
	}

	result := queryUtils.HtmlQuery(req.DocContent, expr1)
	fmt.Printf("%s - %s: %v", selectionType, expr1, result)

	ret = v1.ParserResponse{
		SelectionType: selectionType,
		Expr:          expr1,
		ExprType:      exprType,
	}

	return
}

func (s *ParserHtmlService) updateHtmlElem(docHtml, selectContent string,
	startLineNum, endLineNum, startColumn, endColumn int) (ret1, ret2 string, selectionType consts.NodeType) {
	lines := strings.Split(docHtml, "\n")

	selectionType = s.getHtmlSelectionType(lines, startLineNum, endLineNum, startColumn, endColumn)

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
		lines1 := strings.Split(docHtml, "\n")
		startLine := []rune(lines1[startLineNum])

		newStartLine := string(startLine[:startColumn]) + newStr + string(startLine[startColumn:])
		lines1[startLineNum] = newStartLine
		ret1 = strings.Join(lines1, "\n")

		//2
		lines2 := strings.Split(docHtml, "\n")
		endLine := []rune(lines2[endLineNum])

		newEndLine := string(endLine[:endColumn]) + newStr + string(endLine[endColumn:])
		lines2[endLineNum] = newEndLine
		ret2 = strings.Join(lines2, "\n")
	}

	return
}

func (s *ParserHtmlService) getHtmlSelectedElem(docHtml string, selectionType consts.NodeType) (ret *html.Node) {
	doc, err := htmlquery.Parse(strings.NewReader(docHtml))
	if err != nil {
		return
	}

	expr := ""
	if selectionType == consts.NodeElem || selectionType == consts.NodeProp {
		expr = fmt.Sprintf("//*[@%s]", consts.DeepestKey)
	} else if selectionType == consts.NodeContent {
		expr = fmt.Sprintf("//text()[contains(.,\"%s\")]", consts.DeepestKey)
	}

	ret, err = htmlquery.Query(doc, expr)

	return
}

func (s *ParserHtmlService) queryElem(docHtml, xpath string) (ret *html.Node) {
	doc, err := htmlquery.Parse(strings.NewReader(docHtml))
	if err != nil {
		return
	}

	expr := fmt.Sprintf(xpath)
	ret, err = htmlquery.Query(doc, expr)

	return
}

func (s *ParserHtmlService) getHtmlSelectionType(lines []string, startLine, endLine, startColumn, endColumn int) (
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
