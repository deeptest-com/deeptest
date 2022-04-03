package serverDomain

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
)

type TestRequestReq struct {
	model.TestRequest
}

type TestRequestResp struct {
	model.TestResponse
}
