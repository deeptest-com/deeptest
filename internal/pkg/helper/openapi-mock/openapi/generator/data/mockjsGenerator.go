package mockData

import (
	"context"
	serverDomain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	mockjsHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/mockjs"
	"github.com/getkin/kin-openapi/openapi3"
)

type MockjsGenerator struct {
	Random     randomGenerator
	Expression string
}

func (g *MockjsGenerator) GenerateDataBySchema(ctx context.Context, schema *openapi3.Schema) (value Data, err error) {
	expr := mockjsHelper.GetMockJsSchemaExpression(schema)

	value, _ = g.GenerateByMockJsExpression(expr, schema.Type)

	return
}

func (g *MockjsGenerator) GenerateByMockJsExpression(expr string, schemaType string) (ret interface{}, err error) {
	req := serverDomain.MockJsExpression{
		Expression: expr,
	}

	result, err := mockjsHelper.EvaluateExpression(req)

	ret = mockjsHelper.ConvertData(result.Result, schemaType)

	return
}
