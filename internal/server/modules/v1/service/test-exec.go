package service

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
)

type TestExecService struct {
}

func NewTestExecService() *TestExecService {
	return &TestExecService{}
}

func (s *TestExecService) Exec(req model.TestRequest) (resp *model.TestResponse, err error) {

	return
}
