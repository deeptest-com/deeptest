package service

import (
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"strings"
)

type ParserService struct {
}

func (s *ParserService) ParseHtml(req *v1.ParserRequest) (ret string, err error) {
	req.DocHtml = s.updateElem(req.DocHtml, req.SelectContent, req.StartLine, req.EndLine,
		req.StartColumn, req.EndColumn)

	elem := s.getSelectedElem(req.DocHtml)

	ret = s.getXpath(elem)

	return
}

func (s *ParserService) updateElem(docHtml, selectContent string,
	startLine, endLine, startColumn, endColumn int) (ret string) {
	lines := strings.Split(docHtml, "\n")

	selectionType := s.getSelectionType(lines, startLine, endLine, startColumn, endColumn)

	line := lines[startLine]

	if selectionType == "elem" {
		newStr := selectContent + fmt.Sprintf(" %s=\"true\"", consts.DeepestKey)
		newLine := line[:startColumn] + newStr + line[endColumn:]

		lines[startLine] = newLine

		ret = strings.Join(lines, "\n")
		return

	} else if selectionType == "prop" {

	} else if selectionType == "content" {

	}

	return
}

func (s *ParserService) getSelectedElem(docHtml string) (ret *html.Node) {
	doc, err := htmlquery.Parse(strings.NewReader(docHtml))
	if err != nil {
		return
	}

	expr := fmt.Sprintf("//*[@%s]", consts.DeepestKey)
	ret, err = htmlquery.Query(doc, expr)

	return
}

func (s *ParserService) getXpath(elem *html.Node) (xpath string) {

	return
}

func (s *ParserService) getSelectionType(lines []string, startLine, endLine, startColumn, endColumn int) (ret string) {
	leftNoSpaceChar := s.getLeftNoSpaceChar(lines, startLine, startColumn)
	rightChar := s.getRightChar(lines, endLine, endColumn)

	if leftNoSpaceChar == "<" && (rightChar == " " || rightChar == ">") {
		ret = "elem"
		return
	}

	leftChar := s.getLeftChar(lines, startLine, startColumn)
	rightNoSpaceChar := s.getRightNoSpaceChar(lines, endLine, endColumn)

	if leftChar == " " && rightNoSpaceChar == "=" {
		ret = "prop"
		return
	}

	ret = "content"
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
	endColumn = -1
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
