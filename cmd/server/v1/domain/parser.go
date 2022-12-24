package domain

type ParserRequest struct {
	DocHtml       string `json:"docHtml" validate:"required"`
	SelectContent string `json:"selectContent" validate:"required"`
	StartLine     int    `json:"startLine" validate:"required"`
	EndLine       int    `json:"endLine" validate:"required"`
	StartColumn   int    `json:"startColumn" validate:"required"`
	EndColumn     int    `json:"endColumn" validate:"required"`
}
