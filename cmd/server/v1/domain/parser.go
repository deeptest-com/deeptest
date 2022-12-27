package domain

import "github.com/aaronchen2k/deeptest/internal/pkg/consts"

type ParserRequest struct {
	DocContent    string `json:"docContent" validate:"required"`
	SelectContent string `json:"selectContent" validate:"required"`
	StartLine     int    `json:"startLine"`
	EndLine       int    `json:"endLine"`
	StartColumn   int    `json:"startColumn"`
	EndColumn     int    `json:"endColumn"`
}

type ParserResponse struct {
	SelectionType consts.NodeType `json:"selectionType"`
	Expr          string          `json:"expr"`
	ExprType      string          `json:"exprType"`
}

type TestXPathRequest struct {
	Expr     string                  `json:"expr" validate:"required"`
	ExprType string                  `json:"exprType" validate:"required"`
	Content  string                  `json:"content" validate:"required"`
	Type     consts.HttpRespLangType `json:"type" validate:"required"`
}

type TestXPathResponse struct {
	Result string `json:"result"`
}

type TestRegxRequest struct {
	Expr    string `json:"expr" validate:"required"`
	Content string `json:"content" validate:"required"`
}

type TestRegxResponse struct {
	Result string `json:"result"`
}
