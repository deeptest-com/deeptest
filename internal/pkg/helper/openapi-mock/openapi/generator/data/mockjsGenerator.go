package mockData

import (
	"context"
	serverDomain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	mockjs "github.com/aaronchen2k/deeptest/internal/pkg/goja/mock-js"
	"github.com/getkin/kin-openapi/openapi3"
)

type MockjsGenerator struct {
	Random     randomGenerator
	Expression string
}

func (g *MockjsGenerator) GenerateDataBySchema(ctx context.Context, schema *openapi3.Schema) (value Data, err error) {
	expr := mockjs.GetMockJsSchemaExpression(schema)
	req := serverDomain.MockJsExpression{
		Expression: expr,
	}

	ret, err := mockjs.EvaluateExpression(req)

	value = mockjs.ConvertData(ret.Result, schema.Type)

	return
}
