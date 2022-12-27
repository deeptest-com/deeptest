package service

import (
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	queryHelper "github.com/aaronchen2k/deeptest/internal/agent/exec/utils/query"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"strings"
)

type ParserRegxService struct {
	XPathService  *XPathService  `inject:""`
	ParserService *ParserService `inject:""`
}

func (s *ParserRegxService) ParseRegx(req *v1.ParserRequest) (ret v1.ParserResponse, err error) {
	expr, _ := s.getRegxExpr(req.DocContent, req.SelectContent,
		req.StartLine, req.StartColumn, req.EndLine, req.EndColumn)

	result := queryHelper.RegxQuery(req.DocContent, expr)

	fmt.Printf("%s: %v", expr, result)

	ret = v1.ParserResponse{
		SelectionType: consts.NodeText,
		Expr:          expr,
	}

	return
}

func (s *ParserRegxService) getRegxExpr(docContent, selectContent string,
	startLine, startColumn, endLine, endColumn int) (ret string, err error) {

	lines := strings.Split(docContent, "\n")

	leftChars := ""
	s.ParserService.getLeftCharsInSingleLine(lines, startLine, startColumn, 6, &leftChars)

	rightChars := ""
	s.ParserService.getRightCharsInSingleLine(lines, endLine, endColumn, 6, &rightChars)

	ret = fmt.Sprintf("%s(.+)%s", leftChars, rightChars)

	return
}
