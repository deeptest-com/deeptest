package service

import (
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	queryHelper "github.com/aaronchen2k/deeptest/internal/agent/exec/utils/query"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/antchfx/jsonquery"
	"strings"
)

type ParserJsonService struct {
	XPathService *XPathService `inject:""`
}

func (s *ParserJsonService) ParseJson(req *v1.ParserRequest) (ret v1.ParserResponse, err error) {
	docJson := s.updateJsonElem(req.DocContent, req.SelectContent, req.StartLine, req.EndLine,
		req.StartColumn, req.EndColumn)

	req.DocContent = docJson

	elem := s.getJsonSelectedElem(req.DocContent)

	parentXpath, _ := s.XPathService.GetJsonXPath(elem, req.SelectContent, true)
	xpath := parentXpath + "/" + req.SelectContent

	result := queryHelper.JsonQuery(req.DocContent, xpath)

	fmt.Printf("%s: %v", xpath, result)

	ret = v1.ParserResponse{
		SelectionType: consts.NodeProp,
		Expr:          xpath,
	}

	return
}

func (s *ParserJsonService) updateJsonElem(docJson, selectContent string,
	startLine, endLine, startColumn, endColumn int) (ret string) {
	lines := strings.Split(docJson, "\n")

	line := lines[startLine]

	newStr := fmt.Sprintf("%s-%s", consts.DeepestKey, selectContent)
	newLine := line[:startColumn] + newStr + line[endColumn:]

	lines[startLine] = newLine

	ret = strings.Join(lines, "\n")
	return
}

func (s *ParserJsonService) getJsonSelectedElem(docJson string) (ret *jsonquery.Node) {
	doc, err := jsonquery.Parse(strings.NewReader(docJson))
	if err != nil {
		return
	}

	expr := fmt.Sprintf("//*[contains(.,'%s')]", consts.DeepestKey)
	ret, err = jsonquery.Query(doc, expr)

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
