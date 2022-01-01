package serverDomain

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
)

type TestStepReq struct {
	model.TestStep
	Ts string `json:"ts"`
}

type TestStepResp struct {
	model.TestStep
}
