package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type JslibService struct {
	JslibRepo *repo.JslibRepo `inject:""`
}

func (s *JslibService) List(keywords string) (ret []model.SysJslib, err error) {
	ret, err = s.JslibRepo.List(keywords)
	return
}

func (s *JslibService) Get(id uint) (model.SysJslib, error) {
	return s.JslibRepo.Get(id)
}

func (s *JslibService) Save(req *model.SysJslib) (err error) {
	return s.JslibRepo.Save(req)
}

func (s *JslibService) UpdateName(req v1.JslibReq) (err error) {
	return s.JslibRepo.UpdateName(req)
}

func (s *JslibService) Delete(id uint) (err error) {
	return s.JslibRepo.Delete(id)
}

func (s *JslibService) Disable(id uint) (err error) {
	return s.JslibRepo.Disable(id)
}
