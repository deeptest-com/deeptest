package service

import (
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	queryHelper "github.com/aaronchen2k/deeptest/internal/agent/exec/utils/query"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"strings"
)

type ParserService struct {
	XPathService *XPathService `inject:""`
}

func (s *ParserService) TestXPath(req *v1.TestXPathRequest) (ret v1.TestXPathResponse, err error) {
	var result interface{}

	if req.ExprType == "xpath" {
		if req.Type == consts.LangHTML {
			result = queryHelper.HtmlQuery(req.Content, req.Expr)
		} else if req.Type == consts.LangXML {
			result = queryHelper.XmlQuery(req.Content, req.Expr)
		} else if req.Type == consts.LangJSON {
			result = queryHelper.JsonQuery(req.Content, req.Expr)
		}
	} else if req.ExprType == "regx" {
		result = queryHelper.RegxQuery(req.Content, req.Expr)
	}

	ret = v1.TestXPathResponse{
		Result: fmt.Sprintf("%v", result),
	}

	return
}

func (s *ParserService) getLeftCharsInSingleLine(lines []string, startLine, startColumn, num int, ret *string) {
	line := lines[startLine]
	if startLine == 0 && startColumn == 0 {
		return
	}

	if startColumn > 0 {
		leftOne := line[startColumn-1 : startColumn]
		*ret = leftOne + *ret

		startColumn -= 1

		if len(*ret) > num {
			return
		}
	} else if startColumn == 0 {
		*ret = "^" + *ret
		return
	}

	s.getLeftCharsInSingleLine(lines, startLine, startColumn, num, ret)
}

func (s *ParserService) getRightCharsInSingleLine(lines []string, endLine, endColumn int, num int, ret *string) {
	line := lines[endLine]

	if endLine == len(lines)-1 && endColumn == len(line)-1 {
		return
	}

	rightOne := ""

	if endColumn < len(line) {
		rightOne = line[endColumn : endColumn+1]
		*ret += rightOne

		endColumn += 1

		if len(*ret) > num {
			return
		}
	} else if endColumn == len(line) {
		*ret += "$"
		return
	}

	s.getRightCharsInSingleLine(lines, endLine, endColumn, num, ret)
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
	if startLine < 0 {
		return
	}

	startColumn = len(lines[startLine])

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
	endColumn = 0
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
