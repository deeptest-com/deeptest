package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/jinzhu/copier"
)

type ServeService struct {
	ServeRepo *repo.ServeRepo `inject:""`
}

func NewServeService() *ServeService {
	return &ServeService{}
}

func (s *ServeService) Paginate(req v1.ServeReqPaginate) (ret _domain.PageData, err error) {
	ret, err = s.ServeRepo.Paginate(req)
	return
}

func (s *ServeService) Save(req v1.ServeReq) (res uint, err error) {
	var serve model.Serve
	copier.CopyWithOption(&serve, req, copier.Option{DeepCopy: true})
	err = s.ServeRepo.Save(serve.ID, &serve)
	return serve.ID, err
}

func (s *ServeService) GetById(id uint) (res model.Serve) {
	res, _ = s.ServeRepo.Get(id)
	return
}

func (s *ServeService) DeleteById(id uint) (err error) {
	err = s.ServeRepo.DeleteById(id)
	return
}

func (s *ServeService) DisableById(id uint) (err error) {
	err = s.ServeRepo.DisableById(id)
	return
}

func (s *ServeService) ListVersion(id uint) ([]model.ServeVersion, error) {
	return s.ServeRepo.ListVersion(id)
}

func (s *ServeService) SaveVersion(req v1.ServeVersionReq) (res uint, err error) {
	var serveVersion model.ServeVersion
	copier.CopyWithOption(&serveVersion, req, copier.Option{DeepCopy: true})
	err, res = s.ServeRepo.Save(serveVersion.ID, &serveVersion), serveVersion.ID
	return
}

func (s *ServeService) DeleteVersionById(id uint) (err error) {
	err = s.ServeRepo.DeleteVersionById(id)
	return
}

func (s *ServeService) DisableVersionById(id uint) (err error) {
	err = s.ServeRepo.DisableVersionById(id)
	return
}

func (s *ServeService) ListServer(serveId uint) (res []model.ServeServer, err error) {
	res, err = s.ServeRepo.ListServer(serveId)
	return
}

func (s *ServeService) SaveServer(req v1.ServeServerReq) (res uint, err error) {
	var serve model.ServeServer
	copier.CopyWithOption(&serve, req, copier.Option{DeepCopy: true})
	err = s.ServeRepo.Save(serve.ID, &serve)
	return serve.ID, err
}
