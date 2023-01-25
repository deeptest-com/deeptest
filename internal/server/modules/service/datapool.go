package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type DatapoolService struct {
	DatapoolRepo *repo.DatapoolRepo `inject:""`
}

func NewDatapoolService() *DatapoolService {
	return &DatapoolService{}
}

func (s *DatapoolService) List(projectId uint) (ret []v1.DatapoolReq, err error) {
	ret, err = s.DatapoolRepo.List(projectId)

	return
}

func (s *DatapoolService) Get(id uint) (model.Datapool, error) {
	return s.DatapoolRepo.Get(id)
}

func (s *DatapoolService) Save(req *model.Datapool) (err error) {
	return s.DatapoolRepo.Save(req)
}

func (s *DatapoolService) SaveData(req v1.DatapoolReq) (err error) {
	return s.DatapoolRepo.SaveData(req)
}

func (s *DatapoolService) Delete(id uint) (err error) {
	return s.DatapoolRepo.Delete(id)
}
