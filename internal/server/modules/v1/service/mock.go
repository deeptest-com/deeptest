package service

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
)

type MockService struct {
}

func NewMockService() *MockService {
	return &MockService{}
}

func (s *MockService) Exec(req model.TestRequest) (resp *model.TestResponse, err error) {

	return
}
