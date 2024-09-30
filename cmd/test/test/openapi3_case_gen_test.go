package test

import (
	httpHelper "github.com/deeptest-com/deeptest/internal/agent/exec"
	"github.com/deeptest-com/deeptest/internal/pkg/domain"
	"github.com/deeptest-com/deeptest/internal/pkg/helper/openapi/convert"
	"log"
	"testing"
)

func TestOpenAPI3CaseGen(t *testing.T) {
	url := "http://127.0.0.1:8085/swagger/doc.json"

	request := domain.BaseRequest{Url: url}

	response, err := httpHelper.Get(request)
	if err != nil {
		return
	}

	data := []byte(response.Content)

	handler := convert.NewHandler(convert.SWAGGER2, data, "")
	doc3, err := handler.ToOpenapi()

	endpointSubmitExecResult := doc3.Paths["/api/v1/debugs/invoke/submitResult"]
	operationPost := endpointSubmitExecResult.Post

	log.Print(operationPost)
}
