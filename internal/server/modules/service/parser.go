package service

import (
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	queryHelper "github.com/aaronchen2k/deeptest/internal/agent/exec/utils/query"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/antchfx/htmlquery"
	"github.com/antchfx/jsonquery"
	"github.com/antchfx/xmlquery"
	"golang.org/x/net/html"
	"strings"
)

type ParserService struct {
	XPathService *XPathService `inject:""`
}

func (s *ParserService) ParseHtml(req *v1.ParserRequest) (ret v1.ParserResponse, err error) {
	docHtml, selectionType := s.updateHtmlElem(req.DocHtml, req.SelectContent, req.StartLine, req.EndLine,
		req.StartColumn, req.EndColumn)

	req.DocHtml = docHtml

	elem := s.getHtmlSelectedElem(req.DocHtml, selectionType)

	xpath, _ := s.XPathService.GetHtmlXPath(elem, req.SelectContent, selectionType, true)

	result := queryHelper.HtmlQuery(req.DocHtml, xpath)

	fmt.Printf("%s - %s: %v", selectionType, xpath, result)

	ret = v1.ParserResponse{
		SelectionType: selectionType,
		XPath:         xpath,
	}

	return
}

func (s *ParserService) ParseXml(req *v1.ParserRequest) (ret v1.ParserResponse, err error) {
	docHtml, selectionType := s.updateXmlElem(req.DocHtml, req.SelectContent, req.StartLine, req.EndLine,
		req.StartColumn, req.EndColumn)

	req.DocHtml = docHtml

	elem := s.getXmlSelectedElem(req.DocHtml, selectionType)

	xpath, _ := s.XPathService.GetXmlXPath(elem, req.SelectContent, selectionType, true)

	result := queryHelper.HtmlQuery(req.DocHtml, xpath)

	fmt.Printf("%s - %s: %v", selectionType, xpath, result)

	ret = v1.ParserResponse{
		SelectionType: selectionType,
		XPath:         xpath,
	}

	return
}

func (s *ParserService) ParseJson(req *v1.ParserRequest) (ret v1.ParserResponse, err error) {
	docHtml, selectionType := s.updateJsonElem(req.DocHtml, req.SelectContent, req.StartLine, req.EndLine,
		req.StartColumn, req.EndColumn)

	req.DocHtml = docHtml

	elem := s.getJsonSelectedElem(req.DocHtml, selectionType)

	xpath, _ := s.XPathService.GetJsonXPath(elem, req.SelectContent, selectionType, true)

	result := queryHelper.JsonQuery(req.DocHtml, xpath)

	fmt.Printf("%s - %s: %v", selectionType, xpath, result)

	ret = v1.ParserResponse{
		SelectionType: selectionType,
		XPath:         xpath,
	}

	return
}

func (s *ParserService) TestXPath(req *v1.TestXPathRequest) (ret v1.TestXPathResponse, err error) {
	var result interface{}

	if req.Type == consts.LangHTML {
		result = queryHelper.HtmlQuery(req.Content, req.XPath)
	} else if req.Type == consts.LangXML {
		result = queryHelper.XmlQuery(req.Content, req.XPath)
	} else if req.Type == consts.LangJSON {
		result = queryHelper.JsonQuery(req.Content, req.XPath)
	}

	ret = v1.TestXPathResponse{
		Result: fmt.Sprintf("%v", result),
	}

	return
}

func (s *ParserService) updateHtmlElem(docHtml, selectContent string,
	startLine, endLine, startColumn, endColumn int) (ret string, selectionType consts.NodeType) {
	lines := strings.Split(docHtml, "\n")

	selectionType = s.getXmlSelectionType(lines, startLine, endLine, startColumn, endColumn)

	line := lines[startLine]
	newStr := fmt.Sprintf(" %s=\"true\" ", consts.DeepestKey)

	if selectionType == consts.Elem {
		newLine := line[:startColumn] + selectContent + newStr + line[endColumn:]

		lines[startLine] = newLine

	} else if selectionType == consts.Prop {
		newLine := line[:startColumn] + newStr + line[startColumn:]

		lines[startLine] = newLine

	} else if selectionType == consts.Content {
		newStr = fmt.Sprintf("[[%s]]", consts.DeepestKey)
		newLine := line[:endColumn] + newStr + line[endColumn:]

		lines[startLine] = newLine
	}

	ret = strings.Join(lines, "\n")

	return
}

func (s *ParserService) updateXmlElem(docHtml, selectContent string,
	startLine, endLine, startColumn, endColumn int) (ret string, selectionType consts.NodeType) {
	lines := strings.Split(docHtml, "\n")

	selectionType = s.getXmlSelectionType(lines, startLine, endLine, startColumn, endColumn)

	line := lines[startLine]
	newStr := fmt.Sprintf(" %s=\"true\" ", consts.DeepestKey)

	if selectionType == consts.Elem {
		newLine := line[:startColumn] + selectContent + newStr + line[endColumn:]

		lines[startLine] = newLine

	} else if selectionType == consts.Prop {
		newLine := line[:startColumn] + newStr + line[startColumn:]

		lines[startLine] = newLine

	} else if selectionType == consts.Content {
		newStr = fmt.Sprintf("[[%s]]", consts.DeepestKey)
		newLine := line[:endColumn] + newStr + line[endColumn:]

		lines[startLine] = newLine
	}

	ret = strings.Join(lines, "\n")
	return
}

func (s *ParserService) updateJsonElem(docHtml, selectContent string,
	startLine, endLine, startColumn, endColumn int) (ret string, selectionType consts.NodeType) {
	lines := strings.Split(docHtml, "\n")

	line := lines[startLine]

	newStr := fmt.Sprintf("%s", consts.DeepestKey)
	newLine := line[:startColumn] + newStr + line[startColumn:]

	lines[startLine] = newLine

	ret = strings.Join(lines, "\n")
	return
}

func (s *ParserService) getHtmlSelectedElem(docHtml string, selectionType consts.NodeType) (ret *html.Node) {
	doc, err := htmlquery.Parse(strings.NewReader(docHtml))
	if err != nil {
		return
	}

	expr := ""
	if selectionType == consts.Elem || selectionType == consts.Prop {
		expr = fmt.Sprintf("//*[@%s]", consts.DeepestKey)
	} else if selectionType == consts.Content {
		expr = fmt.Sprintf("//text()[contains(.,\"%s\")]", consts.DeepestKey)
	}

	ret, err = htmlquery.Query(doc, expr)

	return
}
func (s *ParserService) getXmlSelectedElem(docXml string, selectionType consts.NodeType) (ret *xmlquery.Node) {
	doc, err := xmlquery.Parse(strings.NewReader(docXml))
	if err != nil {
		return
	}

	expr := ""
	if selectionType == consts.Elem || selectionType == consts.Prop {
		expr = fmt.Sprintf("//*[@%s]", consts.DeepestKey)
	} else if selectionType == consts.Content {
		expr = fmt.Sprintf("//text()[contains(.,\"%s\")]", consts.DeepestKey)
	}

	ret, err = xmlquery.Query(doc, expr)

	return
}
func (s *ParserService) getJsonSelectedElem(docHtml string, selectionType consts.NodeType) (ret *jsonquery.Node) {
	doc, err := jsonquery.Parse(strings.NewReader(docHtml))
	if err != nil {
		return
	}

	expr := fmt.Sprintf("//*[%s]", consts.DeepestKey)
	ret, err = jsonquery.Query(doc, expr)

	return
}

func (s *ParserService) queryElem(docHtml, xpath string) (ret *html.Node) {
	doc, err := htmlquery.Parse(strings.NewReader(docHtml))
	if err != nil {
		return
	}

	expr := fmt.Sprintf(xpath)
	ret, err = htmlquery.Query(doc, expr)

	return
}

func (s *ParserService) getXmlSelectionType(lines []string, startLine, endLine, startColumn, endColumn int) (
	ret consts.NodeType) {

	leftNoSpaceChar := s.getLeftNoSpaceChar(lines, startLine, startColumn)
	rightChar := s.getRightChar(lines, endLine, endColumn)

	if leftNoSpaceChar == "<" && (rightChar == " " || rightChar == ">") {
		ret = consts.Elem
		return
	}

	leftChar := s.getLeftChar(lines, startLine, startColumn)
	rightNoSpaceChar := s.getRightNoSpaceChar(lines, endLine, endColumn)

	if leftChar == " " && rightNoSpaceChar == "=" {
		ret = consts.Prop
		return
	}

	ret = consts.Content
	return
}

func (s *ParserService) getLeftNoSpaceChar(lines []string, startLine, startColumn int) (ret string) {
	line := lines[startLine]
	if startLine == 0 && startColumn == 0 {
		return ""
	}

	leftOne := ""

	if startColumn > 0 {
		leftOne = line[startColumn-1 : startColumn]
		if s.isNotSpace(leftOne) {
			return leftOne
		}
	}

	startLine -= 1
	startColumn = len(lines[startLine])
	if startLine < 0 {
		return
	}

	return s.getLeftNoSpaceChar(lines, startLine, startColumn)
}

func (s *ParserService) getRightNoSpaceChar(lines []string, endLine, endColumn int) (ret string) {
	line := lines[endLine]

	if endLine == len(lines)-1 && endColumn == len(line)-1 {
		return
	}

	rightOne := ""

	if len(line) > endColumn {
		rightOne = line[endColumn : endColumn+1]
		if s.isNotSpace(rightOne) {
			return rightOne
		}
	}

	endLine += 1
	endColumn = -1
	if endLine >= len(lines) {
		return
	}

	return s.getLeftNoSpaceChar(lines, endLine, endColumn)
}

func (s *ParserService) getLeftChar(lines []string, startLine, startColumn int) (ret string) {
	if startLine == 0 && startColumn == 0 {
		return
	}

	line := lines[startLine]

	if startColumn > 1 {
		ret = line[startColumn-1 : startColumn]
		return
	}

	startLine -= 1
	startColumn = len(lines[startLine])
	if startLine < 0 {
		return
	}

	return s.getLeftChar(lines, startLine, startColumn)
}

func (s *ParserService) getRightChar(lines []string, endLine, endColumn int) (ret string) {
	line := lines[endLine]
	if endLine == len(lines)-1 && endColumn == len(line)-1 {
		return
	}

	if len(line) > endColumn {
		ret = line[endColumn : endColumn+1]
		return
	}

	endLine += 1
	endColumn = 0
	if endLine >= len(lines) {
		return
	}

	ret = s.getRightChar(lines, endLine, endColumn)
	return
}

func (s *ParserService) isNotSpace(str string) bool {
	temp := strings.TrimSpace(str)

	return len(temp) > 0
}
