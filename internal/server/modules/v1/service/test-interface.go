package service

import (
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
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
	root, err = s.InterfaceRepo.GetInterfaceTree(projectId)
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
	interf = &model.TestInterface{Name: req.Name, ProjectId: uint(req.ProjectId),
		IsDir: req.Type == serverConsts.Dir}

	target, err := s.InterfaceRepo.Get(uint(req.Target))
	if req.Mode == serverConsts.Child {
		interf.ParentId = target.ID
	} else {
		interf.ParentId = target.ParentId
	}

	interf.Ordr = s.InterfaceRepo.GetMaxOrder(interf.ParentId)

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
		srcInterface.ParentId = targetId
		srcInterface.Ordr = s.InterfaceRepo.GetMaxOrder(srcInterface.ParentId)
	} else if "-1" == mode {
		err = s.InterfaceRepo.AddOrderForTargetAndNextCases(srcInterface.ID, targetInterface.Ordr, targetInterface.ParentId)
		if err != nil {
			return
		}

		srcInterface.ParentId = targetInterface.ParentId
		srcInterface.Ordr = targetInterface.Ordr
	} else if "1" == mode {
		err = s.InterfaceRepo.AddOrderForNextCases(srcInterface.ID, targetInterface.Ordr, targetInterface.ParentId)
		if err != nil {
			return
		}

		srcInterface.ParentId = targetInterface.ParentId
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
