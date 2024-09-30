package service

import (
	"fmt"
	"github.com/antchfx/jsonquery"
	v1 "github.com/deeptest-com/deeptest/cmd/server/v1/domain"
	queryUtils "github.com/deeptest-com/deeptest/internal/agent/exec/utils/query"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"strings"
)

type ParserJsonService struct {
	XPathService      *XPathService      `inject:""`
	ParserRegxService *ParserRegxService `inject:""`
	ParserService     *ParserService     `inject:""`
}

func (s *ParserJsonService) ParseJson(req *v1.ParserRequest) (ret v1.ParserResponse, err error) {
	docJson := s.updateJsonElem(req.DocContent, req.SelectContent, req.StartLine, req.EndLine,
		req.StartColumn, req.EndColumn)

	elem := s.getJsonSelectedElem(docJson, req.SelectContent)

	exprType := "xpath"
	expr, _ := s.XPathService.GetJsonXPath(elem, req.SelectContent, true)
	if expr == "" {
		ret = s.ParserService.GetRegxResponse(req)
		return
	}

	expr = strings.Replace(expr, consts.DeepestKey, req.SelectContent, 1)

	result, resultType := queryUtils.JsonQuery(req.DocContent, expr)
	fmt.Printf("%s: %v, %v", expr, result, resultType)

	ret = v1.ParserResponse{
		SelectionType: consts.NodeProp,
		Expr:          expr,
		ExprType:      exprType,
	}

	return
}

func (s *ParserJsonService) updateJsonElem(docJson, selectContent string,
	startLine, endLine, startColumn, endColumn int) (ret string) {
	lines := strings.Split(docJson, "\n")

	line := []rune(lines[startLine])

	key := fmt.Sprintf("%s", consts.DeepestKey)

	if string(line[startColumn-1:startColumn]) != "\"" {
		key = "\"" + key
	}
	if string(line[endColumn:endColumn+1]) != "\"" {
		key = key + "\""
	}

	newLine := string(line[:startColumn]) + key + string(line[endColumn:])
	lines[startLine] = newLine

	ret = strings.Join(lines, "\n")
	return
}

func (s *ParserJsonService) getJsonSelectedElem(docJson, selectContent string) (ret *jsonquery.Node) {
	doc, err := jsonquery.Parse(strings.NewReader(docJson))
	if err != nil {
		return
	}

	expr := fmt.Sprintf("//*/text()[contains(.,'%s')]", consts.DeepestKey)
	ret, err = jsonquery.Query(doc, expr)
	if err != nil {
		return
	}

	if ret == nil {
		return
	}

	if ret.Type == jsonquery.TextNode &&
		ret.Parent.Data != "" { // except the array parent node which data is empty

		ret = ret.Parent
	}

	return
}

func (s *ParserJsonService) queryElem(docJson, xpath string) (ret *jsonquery.Node) {
	doc, err := jsonquery.Parse(strings.NewReader(docJson))
	if err != nil {
		return
	}

	expr := fmt.Sprintf(xpath)
	ret, err = jsonquery.Query(doc, expr)

	return
}
