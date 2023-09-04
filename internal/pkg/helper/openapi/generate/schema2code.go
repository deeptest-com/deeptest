package generate

import "github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi"

type Schema2Code struct {
	langType string
	nameRule string
}

func NewSchema2Code(langType, nameRule string) *Schema2Code {
	return &Schema2Code{langType, nameRule}
}

func (s *Schema2Code) Convert(schema openapi.SchemaRef) string {
	return ""
}
