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
	req := serverDomain.MockJsExpression{
		Expression: expr,
	}

	ret, err := mockjsHelper.EvaluateExpression(req)

	value = ret.Result

	return
}
