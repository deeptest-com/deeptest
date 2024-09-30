package serverDomain

import "github.com/deeptest-com/deeptest/internal/pkg/consts"

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

type TestExprRequest struct {
	Expr     string                  `json:"expr" validate:"required"`
	ExprType string                  `json:"exprType" validate:"required"`
	Content  string                  `json:"content" validate:"required"`
	Type     consts.HttpRespLangType `json:"type" validate:"required"`
}

type TestExprResponse struct {
	Result     string                     `json:"result"`
	ResultType consts.ExtractorResultType `json:"resultType"`
}
