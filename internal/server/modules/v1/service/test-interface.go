package service

import (
	serverDomain "github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
)

type TestInterfaceService struct {
	InterfaceRepo *repo.InterfaceRepo `inject:""`
}

func NewTestInterfaceService() *TestInterfaceService {
	return &TestInterfaceService{}
}

func (s *TestInterfaceService) GetTree(projectId int) (root *model.TestInterface, err error) {
	root, err = s.InterfaceRepo.GetDefInterfaceTree(projectId)
	return
}

func (s *TestInterfaceService) Get(interfId int) (interf model.TestInterface, err error) {
	interf, err = s.InterfaceRepo.Get(uint(interfId))

	return
}

func (s *TestInterfaceService) Save(interf *model.TestInterface) (err error) {
	err = s.InterfaceRepo.Save(interf)

	return
}
func (s *TestInterfaceService) Create(req serverDomain.TestInterfaceReq) (interf *model.TestInterface, err error) {
	interf = &model.TestInterface{}
	interf.Name = req.Name
	if req.Mode == "root" {
		interf.ParentID = 0
	} else {
		var target model.TestInterface

		target, err = s.InterfaceRepo.Get(uint(req.TargetId))

		if req.Mode == "child" {
			interf.ParentID = target.ID
		} else {
			interf.ParentID = target.ParentID
		}
		interf.Ordr = s.InterfaceRepo.GetMaxOrder(interf.ParentID)
	}

	err = s.InterfaceRepo.Save(interf)
	return
}

func (s *TestInterfaceService) Delete(projectId, id uint) (err error) {
	interf, _ := s.InterfaceRepo.Get(id)

	err = s.deleteInterfaceAndChildren(projectId, interf.ID)

	return
}

func (s *TestInterfaceService) Move(srcId, targetId uint, mode string) (projectId uint, srcInterface model.TestInterface, err error) {
	srcInterface, err = s.InterfaceRepo.Get(srcId)
	targetInterface, err := s.InterfaceRepo.Get(targetId)

	if "0" == mode {
		srcInterface.ParentID = targetId
		srcInterface.Ordr = s.InterfaceRepo.GetMaxOrder(srcInterface.ParentID)
	} else if "-1" == mode {
		err = s.InterfaceRepo.AddOrderForTargetAndNextCases(srcInterface.ID, targetInterface.Ordr, targetInterface.ParentID)
		if err != nil {
			return
		}

		srcInterface.ParentID = targetInterface.ParentID
		srcInterface.Ordr = targetInterface.Ordr
	} else if "1" == mode {
		err = s.InterfaceRepo.AddOrderForNextCases(srcInterface.ID, targetInterface.Ordr, targetInterface.ParentID)
		if err != nil {
			return
		}

		srcInterface.ParentID = targetInterface.ParentID
		srcInterface.Ordr = targetInterface.Ordr + 1
	}

	err = s.InterfaceRepo.UpdateOrdAndParent(srcInterface)

	return
}

func (s *TestInterfaceService) deleteInterfaceAndChildren(projectId, interfId uint) (err error) {
	err = s.InterfaceRepo.Delete(interfId)
	if err == nil {
		children, _ := s.InterfaceRepo.GetChildren(projectId, interfId)
		for _, child := range children {
			s.deleteInterfaceAndChildren(child.ProjectId, child.ID)
		}
	}

	return
}

func (s *TestInterfaceService) Update(id int, req model.TestInterface) (err error) {

	return
}
