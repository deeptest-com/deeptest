package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type SysAgentService struct {
	SysAgentRepo *repo.SysAgentRepo `inject:""`
}

func (s *SysAgentService) List() (pos []model.SysAgent, err error) {
	pos, err = s.SysAgentRepo.List()

	return
}

func (s *SysAgentService) Get(id uint) (po model.SysAgent, err error) {
	po, err = s.SysAgentRepo.Get(id)

	return
}

func (s *SysAgentService) Save(req *model.SysAgent) (err error) {
	err = s.SysAgentRepo.Save(req)
	return
}

func (s *SysAgentService) UpdateName(req v1.JslibReq) (err error) {
	err = s.SysAgentRepo.UpdateName(req)
	return
}

func (s *SysAgentService) Delete(id uint) (err error) {
	return s.SysAgentRepo.Delete(id)
}

func (s *SysAgentService) Disable(id uint) (err error) {
	return s.SysAgentRepo.Disable(id)
}
