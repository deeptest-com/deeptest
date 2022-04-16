package service

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/kataras/iris/v12"
)

type MockService struct {
}

func NewMockService() *MockService {
	return &MockService{}
}

func (s *MockService) Exec(req model.TestRequest) (resp iris.Map, err error) {
	resp["mockData"] = "mockData"

	return
}
