package domain

import "github.com/aaronchen2k/deeptest/internal/pkg/consts"

type ParserRequest struct {
	DocHtml       string `json:"docHtml" validate:"required"`
	SelectContent string `json:"selectContent" validate:"required"`
	StartLine     int    `json:"startLine"`
	EndLine       int    `json:"endLine"`
	StartColumn   int    `json:"startColumn"`
	EndColumn     int    `json:"endColumn"`
}

type ParserResponse struct {
	SelectionType consts.NodeType `json:"selectionType"`
	XPath         string          `json:"xpath"`
}

type TestXPathRequest struct {
	XPath   string                  `json:"xpath" validate:"required"`
	Content string                  `json:"content" validate:"required"`
	Type    consts.HttpRespLangType `json:"type" validate:"required"`
}

type TestXPathResponse struct {
	Result string `json:"result"`
}
