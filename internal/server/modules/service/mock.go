package service

import (
	"github.com/kataras/iris/v12"
)

type MockService struct {
}

func (s *MockService) Exec(req interface{}) (resp iris.Map, err error) {
	resp["req"] = req

	return
}
