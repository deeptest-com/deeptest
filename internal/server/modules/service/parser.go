package service

import (
	"fmt"
	v1 "github.com/deeptest-com/deeptest/cmd/server/v1/domain"
	queryUtils "github.com/deeptest-com/deeptest/internal/agent/exec/utils/query"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"strings"
)

type ParserService struct {
	XPathService      *XPathService      `inject:""`
	ParserRegxService *ParserRegxService `inject:""`
}

func (s *ParserService) TestExpr(req *v1.TestExprRequest) (ret v1.TestExprResponse, err error) {
	var result interface{}
	var resultType consts.ExtractorResultType

	if req.ExprType == "xpath" {
		if req.Type == consts.LangHTML {
			result = queryUtils.HtmlQuery(req.Content, req.Expr)
		} else if req.Type == consts.LangXML {
			result = queryUtils.XmlQuery(req.Content, req.Expr)
		} else if req.Type == consts.LangJSON {
			result, resultType = queryUtils.JsonQuery(req.Content, req.Expr)
		}
	} else if req.ExprType == "regx" {
		result = queryUtils.RegxQuery(req.Content, req.Expr)
	}

	ret = v1.TestExprResponse{
		Result:     fmt.Sprintf("%v", result),
		ResultType: resultType,
	}

	return
}

func (s *ParserService) getLeftCharsInSingleLine(lines []string, startLine, startColumn, num int,
	ret, prefix *string) {
	line := []rune(lines[startLine])
	if startLine == 0 && startColumn == 0 {
		return
	}

	if startColumn > 0 {
		leftOne := line[startColumn-1 : startColumn]
		*ret = string(leftOne) + *ret

		startColumn -= 1

		if len(*ret) > num {
			return
		}
	} else if startColumn == 0 {
		*prefix = "^"
		return
	}

	s.getLeftCharsInSingleLine(lines, startLine, startColumn, num, ret, prefix)

	return
}

func (s *ParserService) getRightCharsInSingleLine(lines []string, endLine, endColumn int, num int, ret, postfix *string) {
	line := []rune(lines[endLine])

	if endLine == len(lines)-1 && endColumn == len(line)-1 {
		return
	}

	rightOne := ""

	if endColumn < len(line) {
		rightOne = string(line[endColumn : endColumn+1])
		*ret += rightOne

		endColumn += 1

		if len(*ret) > num {
			return
		}
	} else if endColumn == len(line) {
		*postfix = "$"
		return
	}

	s.getRightCharsInSingleLine(lines, endLine, endColumn, num, ret, postfix)

	return
}

func (s *ParserService) getLeftNoSpaceChar(lines []string, startLine, startColumn int) (ret string) {
	line := []rune(lines[startLine])
	if startLine == 0 && startColumn == 0 {
		return ""
	}

	leftOne := ""

	if startColumn > 0 {
		leftOne = string(line[startColumn-1 : startColumn])
		if s.isNotSpace(leftOne) {
			return leftOne
		}
	}

	if startColumn > 1 {
		startColumn -= 1
		return s.getLeftNoSpaceChar(lines, startLine, startColumn)
	}

	startLine -= 1
	if startLine < 0 {
		return
	}

	startColumn = len(lines[startLine])

	return s.getLeftNoSpaceChar(lines, startLine, startColumn)
}

func (s *ParserService) getRightNoSpaceChar(lines []string, endLine, endColumn int) (ret string) {
	line := []rune(lines[endLine])

	if endLine == len(lines)-1 && endColumn == len(line)-1 {
		return
	}

	rightOne := ""

	if len(line) > endColumn {
		rightOne = string(line[endColumn : endColumn+1])
		if s.isNotSpace(rightOne) {
			return rightOne
		}
	}

	if endColumn < len(line)-1 {
		endColumn += 1
		return s.getRightNoSpaceChar(lines, endLine, endColumn)
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

	line := []rune(lines[startLine])

	if startColumn > 1 {
		ret = string(line[startColumn-1 : startColumn])
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
	line := []rune(lines[endLine])
	if endLine == len(lines)-1 && endColumn == len(line)-1 {
		return
	}

	if len(line) > endColumn {
		ret = string(line[endColumn : endColumn+1])
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

func (s *ParserService) GetRegxResponse(req *v1.ParserRequest) (ret v1.ParserResponse) {
	expr, _ := s.ParserRegxService.getRegxExpr(req.DocContent, req.SelectContent,
		req.StartLine, req.StartColumn,
		req.EndLine, req.EndColumn)

	ret = v1.ParserResponse{
		SelectionType: consts.NodeContent,
		Expr:          expr,
		ExprType:      "regx",
	}

	return
}
